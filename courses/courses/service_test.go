package courses

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestListAllSections(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/courses/12/all-sections" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"sections":[{"id":1,"name":"第一教学班"},{"id":2,"name":"第二教学班"}]}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.ListAllSections(context.Background(), 12)
	if err != nil {
		t.Fatalf("ListAllSections returned error: %v", err)
	}
	if len(result.Sections) != 2 {
		t.Fatalf("unexpected section count: %d", len(result.Sections))
	}
	if result.Sections[0].Name != "第一教学班" {
		t.Fatalf("unexpected section name: %s", result.Sections[0].Name)
	}
}

func TestListMyCoursesByConditionsTyped(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/my-courses" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		var body ListMyCoursesRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if !body.ShowScorePassedStatus {
			t.Fatalf("showScorePassedStatus was not forwarded")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"courses":[
				{
					"id":1,
					"name":"程序设计基础",
					"grade":{"id":3,"name":"2023级"},
					"klass":{"id":9,"name":"计算机2301"},
					"registered":true,
					"knowledge_node_count":7,
					"team_teachings":[{"id":5,"name":"助教甲"}],
					"course_attributes":{"passing_score":60,"score_type":"percentage"}
				}
			],
			"page":1,
			"page_size":20,
			"pages":1,
			"total":1,
			"start":1,
			"end":1
		}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.ListMyCoursesByConditionsTyped(context.Background(), &ListMyCoursesRequest{
		Page:                  1,
		PageSize:              20,
		ShowScorePassedStatus: true,
	})
	if err != nil {
		t.Fatalf("ListMyCoursesByConditionsTyped returned error: %v", err)
	}
	course := result.Courses[0]
	if course.Grade == nil || course.Grade.Name != "2023级" {
		t.Fatalf("grade did not decode: %#v", course.Grade)
	}
	if course.Klass == nil || course.Klass.Name != "计算机2301" {
		t.Fatalf("klass did not decode: %#v", course.Klass)
	}
	if !course.Registered {
		t.Fatalf("registered did not decode")
	}
	if course.KnowledgeNodeCount != 7 {
		t.Fatalf("knowledge_node_count did not decode: %d", course.KnowledgeNodeCount)
	}
	if len(course.TeamTeachings) != 1 {
		t.Fatalf("team_teachings did not decode: %#v", course.TeamTeachings)
	}
	if course.CourseAttributes == nil || course.CourseAttributes.PassingScore == nil || *course.CourseAttributes.PassingScore != 60 {
		t.Fatalf("passing_score did not decode: %#v", course.CourseAttributes)
	}
	if course.CourseAttributes.ScoreType == nil || *course.CourseAttributes.ScoreType != "percentage" {
		t.Fatalf("score_type did not decode: %#v", course.CourseAttributes)
	}
}

func TestListEducatorsDecodesNestedUserModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/courses/7/educators" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"enrollments":[
				{
					"id":1,
					"user":{
					"id":99,
						"name":"张三",
						"grade":{"id":3,"name":"2023级"},
						"klass":{"id":12,"name":"计算机2301","code":"2301"},
						"learning_center":{"id":5,"name":"紫金港学习中心"},
						"user_attributes":{"portfolio_url":"https://example.com/u/99","tag":"AI助教"}
					}
				}
			]
		}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.ListEducators(context.Background(), 7, "")
	if err != nil {
		t.Fatalf("ListEducators returned error: %v", err)
	}
	user := result.Enrollments[0].User
	if user == nil || user.Grade == nil || user.Grade.Name != "2023级" {
		t.Fatalf("grade did not decode: %#v", user)
	}
	if user.Klass == nil || user.Klass.Code != "2301" {
		t.Fatalf("klass did not decode: %#v", user)
	}
	if user.LearningCenter == nil || user.LearningCenter.Name != "紫金港学习中心" {
		t.Fatalf("learning_center did not decode: %#v", user.LearningCenter)
	}
	if user.UserAttributes == nil || user.UserAttributes.Tag == nil || *user.UserAttributes.Tag != "AI助教" {
		t.Fatalf("user_attributes.tag did not decode: %#v", user.UserAttributes)
	}
}

func TestGetEnrollmentUserUsesFrontendEndpointAndModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/courses/18/enrollments/users/42" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("fields") != "id,name,user_addresses,user_auth_externals,learning_center" {
			t.Fatalf("unexpected fields query: %s", r.URL.RawQuery)
		}
		if got := r.Header.Get("Request-Scope"); got != "teaching-team" {
			t.Fatalf("unexpected Request-Scope header: %s", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"id":42,
			"name":"李四",
			"email":"lisi@example.edu.cn",
			"mobile_phone":"13800001111",
			"learning_center":{"id":6,"name":"紫金港学习中心"},
			"user_addresses":[
				{"id":1,"type":"home","name":"杭州市西湖区","post_code":"310000"}
			],
			"user_auth_externals":[
				{"id":7,"type":"qq","uid":"123456"},
				{"id":8,"type":"wechat","uid":"wx_abc"}
			]
		}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.GetEnrollmentUserWithParams(context.Background(), 18, 42, GetEnrollmentUserParams{
		Fields:       "id,name,user_addresses,user_auth_externals,learning_center",
		RequestScope: "teaching-team",
	})
	if err != nil {
		t.Fatalf("GetEnrollmentUserWithParams returned error: %v", err)
	}
	if result.Email == nil || *result.Email != "lisi@example.edu.cn" {
		t.Fatalf("email did not decode: %#v", result.Email)
	}
	if result.MobilePhone != "13800001111" {
		t.Fatalf("mobile_phone did not decode: %s", result.MobilePhone)
	}
	if result.LearningCenter == nil || result.LearningCenter.Name != "紫金港学习中心" {
		t.Fatalf("learning_center did not decode: %#v", result.LearningCenter)
	}
	if len(result.UserAddresses) != 1 || result.UserAddresses[0].Name != "杭州市西湖区" || result.UserAddresses[0].PostCode != "310000" {
		t.Fatalf("user_addresses did not decode: %#v", result.UserAddresses)
	}
	if len(result.UserAuthExternals) != 2 || result.UserAuthExternals[0].Type != "qq" || result.UserAuthExternals[0].UID != "123456" {
		t.Fatalf("user_auth_externals did not decode: %#v", result.UserAuthExternals)
	}
}

func TestGetBlueprintSubItemsCountEncodesActivities(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/blueprint/11/sub-items-count" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		got := r.URL.Query().Get("activities")
		want := `[{"id":21,"type":"homework"}]`
		if got != want {
			t.Fatalf("unexpected activities query: %s", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"items":[{"course_id":8,"count":2}]}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.GetBlueprintSubItemsCount(context.Background(), 11, []BlueprintActivityRef{{ID: 21, Type: "homework"}})
	if err != nil {
		t.Fatalf("GetBlueprintSubItemsCount returned error: %v", err)
	}
	if len(result.Items) != 1 || result.Items[0].Count != 2 {
		t.Fatalf("unexpected response: %#v", result.Items)
	}
}

func TestCancelBlueprintActivitySyncUsesDeleteBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/blueprint/9/activities/5/cancel-sync" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("read body: %v", err)
		}
		want := `{"target_course_id":18,"id":33,"type":"questionnaire"}`
		if string(bytes.TrimSpace(body)) != want {
			t.Fatalf("unexpected body: %s", string(body))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	err := service.CancelBlueprintActivitySync(context.Background(), 9, 5, &CancelBlueprintActivitySyncRequest{
		TargetCourseID: 18,
		ID:             33,
		Type:           "questionnaire",
	})
	if err != nil {
		t.Fatalf("CancelBlueprintActivitySync returned error: %v", err)
	}
}

func TestGetAssistantPermissions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/course/15/assistant-permissions" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"instructor_assistant":{"edit_course_catlog":true},
			"student_assistant":{"view_course_rollcall":false}
		}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.GetAssistantPermissions(context.Background(), 15)
	if err != nil {
		t.Fatalf("GetAssistantPermissions returned error: %v", err)
	}
	if !result.InstructorAssistant["edit_course_catlog"] || result.StudentAssistant["view_course_rollcall"] {
		t.Fatalf("unexpected permissions: %#v", result)
	}
}

func TestMailToEnrollments(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/courses/6/mail-to-enrollments" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		var body SendMailToEnrollmentsRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if len(body.EnrollmentIDs) != 2 || body.MailSubject != "通知" {
			t.Fatalf("unexpected body: %#v", body)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message":"sent"}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	_, err := service.MailToEnrollments(context.Background(), 6, &SendMailToEnrollmentsRequest{
		EnrollmentIDs: []int{1, 2},
		MailSubject:   "通知",
		MailContent:   "内容",
	})
	if err != nil {
		t.Fatalf("MailToEnrollments returned error: %v", err)
	}
}

func TestOutlineHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/outline":
			_, _ = w.Write([]byte(`{
				"id":14,
				"common_fields":[{"id":1,"key":"course_objective","title":"课程目标"}],
				"teaching_schedule_and_others_fields":[{"id":2,"key":"teaching_schedule","title":"教学进度"}],
				"comment_chinese":{"id":3,"key":"comment_chinese","attachments":[{"id":9,"name":"导学视频.mp4","type":"video"}]}
			}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/14/outline-item":
			var body UpsertOutlineItemRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode outline-item body: %v", err)
			}
			if body.CourseID != 14 || !body.SendMessage || len(body.Uploads) != 2 || len(body.AllowDownloadData) != 1 {
				t.Fatalf("unexpected outline-item body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"course_outline_item":{"id":5,"key":"comment_chinese","attachments":[{"id":9,"name":"导学视频.mp4","type":"video"}]}}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/api/course/14/outline-item/5":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/download-print-permissions":
			if got := r.URL.Query().Get("type"); got != "outline" {
				t.Fatalf("unexpected download-print type: %q", got)
			}
			_, _ = w.Write([]byte(`{"allow_download":true,"allow_print":false}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/courses/14/download-print-permissions":
			var body OutlineDownloadPrintPermissions
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode download-print body: %v", err)
			}
			if !body.AllowDownload || !body.AllowPrint {
				t.Fatalf("unexpected download-print body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	outline, err := service.GetOutline(ctx, 14)
	if err != nil {
		t.Fatalf("GetOutline returned error: %v", err)
	}
	if len(outline.TeachingScheduleAndOthersFields) != 1 || outline.CommentChinese == nil || len(outline.CommentChinese.Attachments) != 1 {
		t.Fatalf("unexpected outline decode: %#v", outline)
	}

	item, err := service.CreateOrUpdateOutlineItem(ctx, 14, &UpsertOutlineItemRequest{
		CourseID:          14,
		Key:               "comment_chinese",
		SendMessage:       true,
		Uploads:           []int{7, 9},
		AllowDownloadData: []int{7},
	})
	if err != nil {
		t.Fatalf("CreateOrUpdateOutlineItem returned error: %v", err)
	}
	if item.CourseOutlineItem == nil || item.CourseOutlineItem.ID != 5 || len(item.CourseOutlineItem.Attachments) != 1 {
		t.Fatalf("unexpected outline-item response: %#v", item)
	}

	if err := service.DeleteOutlineItem(ctx, 14, 5); err != nil {
		t.Fatalf("DeleteOutlineItem returned error: %v", err)
	}

	permissions, err := service.GetOutlineDownloadPrintPermissions(ctx, 14, "outline")
	if err != nil {
		t.Fatalf("GetOutlineDownloadPrintPermissions returned error: %v", err)
	}
	if !permissions.AllowDownload || permissions.AllowPrint {
		t.Fatalf("unexpected download-print permissions: %#v", permissions)
	}

	if err := service.UpdateOutlineDownloadPrintPermissions(ctx, 14, &OutlineDownloadPrintPermissions{
		AllowDownload: true,
		AllowPrint:    true,
	}); err != nil {
		t.Fatalf("UpdateOutlineDownloadPrintPermissions returned error: %v", err)
	}
}

func TestScoreAndCompletenessHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/performance/score-setting":
			_, _ = w.Write([]byte(`{"score_percentage":15}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/course/14/performance/score-setting":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/students-performance":
			if got := r.URL.Query().Get("isOriginalScore"); got != "true" {
				t.Fatalf("unexpected isOriginalScore: %q", got)
			}
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected page: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "10" {
				t.Fatalf("unexpected page_size: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"user_ids":[3,4]}` {
				t.Fatalf("unexpected conditions: %q", got)
			}
			if got := r.URL.Query().Get("onlyStudentsName"); got != "true" {
				t.Fatalf("unexpected onlyStudentsName: %q", got)
			}
			_, _ = w.Write([]byte(`{"items":[{"student_id":3,"score":88}]}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/course/14/performance/score":
			var body UpdateStudentPerformanceScoreRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode performance/score body: %v", err)
			}
			if body.StudentID != 3 || body.Score != "88.5" {
				t.Fatalf("unexpected performance/score body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/score-percentages":
			_, _ = w.Write([]byte(`{"score_percentage_total":100,"unpublished_percentage":20}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/course/14/course-advance-setting":
			var body CourseAdvanceSettingRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode course-advance-setting body: %v", err)
			}
			if body.Params["master_score_percent"] != float64(0) {
				t.Fatalf("unexpected course-advance-setting body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/courses/14/score-percentages":
			var body UpdateScorePercentagesRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode score-percentages body: %v", err)
			}
			if len(body.ActivityScorePercentages) != 1 || len(body.CustomScorePercentages) != 1 || body.RollcallScorePercentage != "10" {
				t.Fatalf("unexpected score-percentages body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/student-self-score":
			_, _ = w.Write([]byte(`{"score":90}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/rollcall-score":
			_, _ = w.Write([]byte(`{"score_percentage":10}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/performance-score":
			if got := r.URL.Query().Get("isOriginalScore"); got != "true" {
				t.Fatalf("unexpected performance-score query: %q", got)
			}
			_, _ = w.Write([]byte(`{"score":12}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/score-distribution":
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected no-intercept: %q", got)
			}
			_, _ = w.Write([]byte(`{"items":[{"range":"90-100","count":3}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/score-type-settings":
			_, _ = w.Write([]byte(`{"score_type":"percentage"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/online-video-completeness/setting":
			if got := r.URL.Query().Get("no-loading-animation"); got != "true" {
				t.Fatalf("unexpected no-loading-animation: %q", got)
			}
			_, _ = w.Write([]byte(`{"id":1,"score_percentage":5}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/course/14/online-video-completeness/setting":
			_, _ = w.Write([]byte(`{"id":1,"score_percentage":8}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/api/course/14/online-video-completeness/setting":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/online-video-completeness/scores":
			_, _ = w.Write([]byte(`{"items":[{"student_id":3,"score":5}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/online-video-completeness/score":
			_, _ = w.Write([]byte(`{"score":5,"complete_rate":0.8}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if setting, err := service.GetPerformanceScoreSetting(ctx, 14); err != nil || (*setting)["score_percentage"] != float64(15) {
		t.Fatalf("unexpected performance score setting: %#v, err=%v", setting, err)
	}
	if err := service.UpdatePerformanceScoreSetting(ctx, 14, map[string]any{"score_percentage": 18}); err != nil {
		t.Fatalf("UpdatePerformanceScoreSetting returned error: %v", err)
	}
	if students, err := service.ListStudentsPerformance(ctx, 14, &StudentsPerformanceParams{
		Page:             2,
		PageSize:         10,
		Conditions:       map[string]any{"user_ids": []int{3, 4}},
		OnlyStudentsName: true,
	}); err != nil || len((*students)["items"].([]any)) != 1 {
		t.Fatalf("unexpected students-performance: %#v, err=%v", students, err)
	}
	if err := service.UpdateStudentPerformanceScore(ctx, 14, &UpdateStudentPerformanceScoreRequest{StudentID: 3, Score: "88.5"}); err != nil {
		t.Fatalf("UpdateStudentPerformanceScore returned error: %v", err)
	}
	if left, err := service.GetScorePercentages(ctx, 14); err != nil || (*left)["score_percentage_total"] != float64(100) {
		t.Fatalf("unexpected score-percentages: %#v, err=%v", left, err)
	}
	if err := service.UpdateCourseAdvanceSetting(ctx, 14, &CourseAdvanceSettingRequest{Params: map[string]any{"master_score_percent": 0}}); err != nil {
		t.Fatalf("UpdateCourseAdvanceSetting returned error: %v", err)
	}
	if err := service.UpdateScorePercentages(ctx, 14, &UpdateScorePercentagesRequest{
		ActivityScorePercentages: []*ActivityScorePercentageItem{{ID: 1, Type: "homework", ScorePercentage: "40"}},
		CustomScorePercentages:   []*CustomScorePercentageItem{{ID: 2, Percentage: "5"}},
		RollcallScorePercentage:  "10",
	}); err != nil {
		t.Fatalf("UpdateScorePercentages returned error: %v", err)
	}
	if self, err := service.GetStudentSelfScore(ctx, 14); err != nil || (*self)["score"] != float64(90) {
		t.Fatalf("unexpected student-self-score: %#v, err=%v", self, err)
	}
	if rollcall, err := service.GetRollcallScore(ctx, 14); err != nil || (*rollcall)["score_percentage"] != float64(10) {
		t.Fatalf("unexpected rollcall-score: %#v, err=%v", rollcall, err)
	}
	if performance, err := service.GetPerformanceScore(ctx, 14, true); err != nil || (*performance)["score"] != float64(12) {
		t.Fatalf("unexpected performance-score: %#v, err=%v", performance, err)
	}
	if distribution, err := service.GetScoreDistribution(ctx, 14); err != nil || len((*distribution)["items"].([]any)) != 1 {
		t.Fatalf("unexpected score-distribution: %#v, err=%v", distribution, err)
	}
	if scoreTypes, err := service.GetScoreTypeSettings(ctx, 14); err != nil || (*scoreTypes)["score_type"] != "percentage" {
		t.Fatalf("unexpected score-type-settings: %#v, err=%v", scoreTypes, err)
	}
	if videoSetting, err := service.GetOnlineVideoCompletenessSettingDetail(ctx, 14); err != nil || (*videoSetting)["score_percentage"] != float64(5) {
		t.Fatalf("unexpected online-video-completeness setting: %#v, err=%v", videoSetting, err)
	}
	if updated, err := service.UpdateOnlineVideoCompletenessSetting(ctx, 14, map[string]any{"score_percentage": 8}); err != nil || (*updated)["score_percentage"] != float64(8) {
		t.Fatalf("unexpected updated online-video-completeness setting: %#v, err=%v", updated, err)
	}
	if err := service.DeleteOnlineVideoCompletenessSetting(ctx, 14); err != nil {
		t.Fatalf("DeleteOnlineVideoCompletenessSetting returned error: %v", err)
	}
	if scores, err := service.GetOnlineVideoCompletenessScores(ctx, 14); err != nil || len((*scores)["items"].([]any)) != 1 {
		t.Fatalf("unexpected online-video-completeness scores: %#v, err=%v", scores, err)
	}
	if score, err := service.GetOnlineVideoCompletenessScore(ctx, 14); err != nil || (*score)["score"] != float64(5) {
		t.Fatalf("unexpected online-video-completeness score: %#v, err=%v", score, err)
	}
}
