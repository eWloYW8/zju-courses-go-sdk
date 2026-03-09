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

func TestRecordEntryUsesFrontendPostEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/course/18/entry/record" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"recorded":true}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.RecordEntry(context.Background(), 18)
	if err != nil {
		t.Fatalf("RecordEntry returned error: %v", err)
	}
	if string(result) != `{"recorded":true}` {
		t.Fatalf("unexpected record entry payload: %s", string(result))
	}

	compat, err := service.GetEntryRecord(context.Background(), 18)
	if err != nil {
		t.Fatalf("GetEntryRecord returned error: %v", err)
	}
	if string(compat) != `{"recorded":true}` {
		t.Fatalf("unexpected compatibility payload: %s", string(compat))
	}
}

func TestGetCombineCourseActivityReadSubCourseIDsUsesFrontendEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/combine-courses/11/activity-read-sub-course-ids" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("activity_id"); got != "23" {
			t.Fatalf("unexpected activity_id query: %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[5,8]`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.GetCombineCourseActivityReadSubCourseIDs(context.Background(), 11, 23)
	if err != nil {
		t.Fatalf("GetCombineCourseActivityReadSubCourseIDs returned error: %v", err)
	}
	if len(result) != 2 || result[0] != 5 || result[1] != 8 {
		t.Fatalf("unexpected sub-course ids: %#v", result)
	}
}

func TestBlueprintSubCourseHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/blueprint/11/sub-courses":
			if got := r.URL.Query().Get("keyword"); got != "算法" {
				t.Fatalf("unexpected keyword: %s", got)
			}
			if got := r.URL.Query().Get("source_id"); got != "6" {
				t.Fatalf("unexpected source_id: %s", got)
			}
			if got := r.URL.Query().Get("source_type"); got != "course" {
				t.Fatalf("unexpected source_type: %s", got)
			}
			_, _ = w.Write([]byte(`{
				"courses":[
					{
						"id":21,
						"name":"算法设计与分析",
						"course_code":"211G0210",
						"status":"ongoing",
						"archived":false,
						"is_instructor":true,
						"course_attributes":{"teaching_class_name":"算法2301"},
						"department":{"id":7,"name":"计算机学院"},
						"academic_year":{"id":4,"name":"2025-2026"},
						"semester":{"id":2,"name":"春","real_name":"春夏"},
						"instructors":[{"id":3,"name":"王老师"}]
					}
				]
			}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/blueprint/11/sub-courses":
			var body BindBlueprintSubCoursesRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode bind body: %v", err)
			}
			if len(body.SubCourseIDs) != 2 || body.SubCourseIDs[0] != 21 || body.SubCourseIDs[1] != 22 {
				t.Fatalf("unexpected bind body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPost && r.URL.Path == "/api/blueprint/11/check-prerequisites":
			var body CheckBlueprintPrerequisitesRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode prerequisites body: %v", err)
			}
			if len(body.Sources) != 1 || body.Sources[0]["type"] != "homework" {
				t.Fatalf("unexpected prerequisites body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"has_prerequisites":true}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/api/blueprint/11/sub-courses/21":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPost && r.URL.Path == "/api/blueprint/sub-courses/21/name":
			var body RenameBlueprintSubCourseRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode rename body: %v", err)
			}
			if body.Name != "蓝图子课程" {
				t.Fatalf("unexpected rename body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	list, err := service.ListBlueprintSubCoursesWithParams(context.Background(), 11, &ListBlueprintSubCoursesParams{
		Keyword:    "算法",
		SourceID:   6,
		SourceType: "course",
	})
	if err != nil {
		t.Fatalf("ListBlueprintSubCoursesWithParams returned error: %v", err)
	}
	if len(list.Courses) != 1 || list.Courses[0].CourseAttributes == nil || list.Courses[0].CourseAttributes.TeachingClassName != "算法2301" {
		t.Fatalf("unexpected sub-courses: %#v", list.Courses)
	}
	if list.Courses[0].Department == nil || list.Courses[0].Department.Name != "计算机学院" {
		t.Fatalf("department did not decode: %#v", list.Courses[0].Department)
	}
	if list.Courses[0].Semester == nil || list.Courses[0].Semester.RealName != "春夏" {
		t.Fatalf("semester did not decode: %#v", list.Courses[0].Semester)
	}

	if err := service.BindBlueprintSubCourses(context.Background(), 11, &BindBlueprintSubCoursesRequest{SubCourseIDs: []int{21, 22}}); err != nil {
		t.Fatalf("BindBlueprintSubCourses returned error: %v", err)
	}

	check, err := service.CheckBlueprintPrerequisites(context.Background(), 11, &CheckBlueprintPrerequisitesRequest{
		Sources: []BlueprintPrerequisiteSource{{"id": 9, "type": "homework"}},
	})
	if err != nil {
		t.Fatalf("CheckBlueprintPrerequisites returned error: %v", err)
	}
	if !check.HasPrerequisites {
		t.Fatalf("expected has_prerequisites to decode")
	}

	if err := service.DeleteBlueprintSubCourse(context.Background(), 11, 21); err != nil {
		t.Fatalf("DeleteBlueprintSubCourse returned error: %v", err)
	}
	if err := service.RenameBlueprintSubCourse(context.Background(), 21, &RenameBlueprintSubCourseRequest{Name: "蓝图子课程"}); err != nil {
		t.Fatalf("RenameBlueprintSubCourse returned error: %v", err)
	}
}

func TestSyncBlueprintUsesFrontendBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/blueprint/11/sync" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		var body SyncBlueprintRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if len(body.TargetCourseIDs) != 2 || body.TargetCourseIDs[0] != 21 || body.Publish == nil || !*body.Publish {
			t.Fatalf("unexpected body: %#v", body)
		}
		sources, ok := body.Sources.([]any)
		if !ok || len(sources) != 1 {
			t.Fatalf("unexpected sources: %#v", body.Sources)
		}
		_, _ = w.Write([]byte(`{"message":"ok"}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	publish := true
	_, err := service.SyncBlueprint(context.Background(), 11, SyncBlueprintRequest{
		Sources:         []map[string]any{{"id": 9, "type": "homework"}},
		TargetCourseIDs: []int{21, 22},
		Publish:         &publish,
	})
	if err != nil {
		t.Fatalf("SyncBlueprint returned error: %v", err)
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

func TestCourseJoinPopupHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/access-code/ABCDEFG/validate":
			_, _ = w.Write([]byte(`{"message":"ok"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/toggle-opened-orgs":
			if got := r.URL.Query().Get("toggle"); got != "org_team_teaching" {
				t.Fatalf("unexpected toggle query: %q", got)
			}
			_, _ = w.Write([]byte(`{"orgs":[{"id":3,"name":"紫金港校区","code":"ZJG"}]}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	validation, err := service.ValidateCourseAccessCode(ctx, "ABCDEFG")
	if err != nil || validation.Message != "ok" {
		t.Fatalf("unexpected access-code validation: %#v, err=%v", validation, err)
	}

	orgs, err := service.ListTeamTeachingOpenedOrgs(ctx)
	if err != nil || len(orgs.Orgs) != 1 || orgs.Orgs[0].Name != "紫金港校区" || orgs.Orgs[0].Code != "ZJG" {
		t.Fatalf("unexpected opened orgs: %#v, err=%v", orgs, err)
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

func TestWarningAndScoreHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/announce-score-settings":
			_, _ = w.Write([]byte(`{"enabled":true}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/courses/14/announce-score-settings":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode announce-score-settings body: %v", err)
			}
			if body["enabled"] != true {
				t.Fatalf("unexpected announce-score-settings body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/performance/score-percentage":
			_, _ = w.Write([]byte(`{"score_percentage":15}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/warnings":
			if got := r.URL.Query().Get("fields"); got != "id,title" {
				t.Fatalf("unexpected warnings fields: %q", got)
			}
			_, _ = w.Write([]byte(`{"warnings":[{"id":5,"title":"预警"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/warnings/5":
			_, _ = w.Write([]byte(`{"id":5,"title":"预警"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/warnings/5/threshold":
			_, _ = w.Write([]byte(`{"min_score":60}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/14/warnings/5/threshold":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode create warning threshold body: %v", err)
			}
			if body["min_score"] != float64(60) {
				t.Fatalf("unexpected create warning threshold body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/course/14/warnings/5/threshold":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode update warning threshold body: %v", err)
			}
			if body["min_score"] != float64(70) {
				t.Fatalf("unexpected update warning threshold body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/warnings/5/students":
			if got := r.URL.Query().Get("fields"); got != "id,name" {
				t.Fatalf("unexpected warning students fields: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"alice"}` {
				t.Fatalf("unexpected warning students conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"students":[{"id":9,"name":"Alice"}]}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/warning/student/9/comment":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode warning student comment body: %v", err)
			}
			if body["comment"] != "follow up" {
				t.Fatalf("unexpected warning student comment body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodDelete && r.URL.Path == "/api/course/14/warning-students/9":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/score-ranks":
			_, _ = w.Write([]byte(`{"items":[{"label":"A","count":1}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/score-item-settings":
			_, _ = w.Write([]byte(`{"items":[{"id":1}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/enrollment-raw-score":
			_, _ = w.Write([]byte(`{"items":[{"student_id":3,"score":95}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/exam/score-type":
			_, _ = w.Write([]byte(`{"score_type":"percentage"}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if settings, err := service.GetAnnounceScoreSettings(ctx, 14); err != nil || (*settings)["enabled"] != true {
		t.Fatalf("unexpected announce-score-settings: %#v, err=%v", settings, err)
	}

	if err := service.UpdateAnnounceScoreSettings(ctx, 14, map[string]any{"enabled": true}); err != nil {
		t.Fatalf("UpdateAnnounceScoreSettings returned error: %v", err)
	}

	if percentage, err := service.GetPerformanceScorePercentage(ctx, 14); err != nil || (*percentage)["score_percentage"] != float64(15) {
		t.Fatalf("unexpected performance score percentage: %#v, err=%v", percentage, err)
	}

	if warnings, err := service.ListWarnings(ctx, 14, "id,title"); err != nil || len((*warnings)["warnings"].([]any)) != 1 {
		t.Fatalf("unexpected warnings list: %#v, err=%v", warnings, err)
	}

	if warning, err := service.GetWarning(ctx, 14, 5); err != nil || (*warning)["id"] != float64(5) {
		t.Fatalf("unexpected warning: %#v, err=%v", warning, err)
	}

	if threshold, err := service.GetWarningThreshold(ctx, 14, 5); err != nil || (*threshold)["min_score"] != float64(60) {
		t.Fatalf("unexpected warning threshold: %#v, err=%v", threshold, err)
	}

	if err := service.CreateWarningThreshold(ctx, 14, 5, map[string]any{"min_score": 60}); err != nil {
		t.Fatalf("CreateWarningThreshold returned error: %v", err)
	}

	if err := service.UpdateWarningThreshold(ctx, 14, 5, map[string]any{"min_score": 70}); err != nil {
		t.Fatalf("UpdateWarningThreshold returned error: %v", err)
	}

	if students, err := service.ListWarningStudents(ctx, 14, 5, &WarningStudentsParams{
		Fields:     "id,name",
		Conditions: map[string]any{"keyword": "alice"},
	}); err != nil || len((*students)["students"].([]any)) != 1 {
		t.Fatalf("unexpected warning students: %#v, err=%v", students, err)
	}

	if err := service.UpdateWarningStudentComment(ctx, 9, map[string]any{"comment": "follow up"}); err != nil {
		t.Fatalf("UpdateWarningStudentComment returned error: %v", err)
	}

	if err := service.DeleteWarningStudent(ctx, 14, 9); err != nil {
		t.Fatalf("DeleteWarningStudent returned error: %v", err)
	}

	if ranks, err := service.GetScoreRanks(ctx, 14); err != nil || len((*ranks)["items"].([]any)) != 1 {
		t.Fatalf("unexpected score-ranks: %#v, err=%v", ranks, err)
	}

	if settings, err := service.GetScoreItemSettings(ctx, 14); err != nil || len((*settings)["items"].([]any)) != 1 {
		t.Fatalf("unexpected score-item-settings: %#v, err=%v", settings, err)
	}

	if rawScore, err := service.GetEnrollmentRawScore(ctx, 14); err != nil || len((*rawScore)["items"].([]any)) != 1 {
		t.Fatalf("unexpected enrollment-raw-score: %#v, err=%v", rawScore, err)
	}

	if scoreType, err := service.GetCourseScoreType(ctx, 14, "exam"); err != nil || scoreType.ScoreType != "percentage" {
		t.Fatalf("unexpected course score type: %#v, err=%v", scoreType, err)
	}
}

func TestLegacyActivityAggregateHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/homework/submission-status":
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected homework submission status query: %q", got)
			}
			_, _ = w.Write([]byte(`{"course_id":14,"homework_activities":[{"id":1}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/exam-scores":
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected exam scores query: %q", got)
			}
			_, _ = w.Write([]byte(`{"exam_scores":[{"id":2}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/catalog-activities":
			_, _ = w.Write([]byte(`{"items":[{"id":1}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/courses/14/catalog-activities":
			_, _ = w.Write([]byte(`{"items":[{"id":2}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/homework-submission-number":
			_, _ = w.Write([]byte(`{"submitted":30}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/classroom-exam-scores":
			_, _ = w.Write([]byte(`{"scores":[{"student_id":3,"score":88}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/exam-submissions":
			_, _ = w.Write([]byte(`{"submissions":[{"id":4}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/courses/14/create-groups":
			_, _ = w.Write([]byte(`{"created":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/questionnaires":
			_, _ = w.Write([]byte(`{"questionnaires":[{"id":5}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/questionnaire-list":
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected questionnaire-list page: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "10" {
				t.Fatalf("unexpected questionnaire-list page_size: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"survey"}` {
				t.Fatalf("unexpected questionnaire-list conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"items":[{"id":6}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/questionnaire-scores":
			_, _ = w.Write([]byte(`{"scores":[{"student_id":3,"score":90}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/web-link-scores":
			_, _ = w.Write([]byte(`{"scores":[{"student_id":3,"score":91}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/virtual-experiment-scores":
			_, _ = w.Write([]byte(`{"scores":[{"student_id":3,"score":92}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/virtual-experiments":
			_, _ = w.Write([]byte(`{"items":[{"id":7}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/14/entry-refers":
			_, _ = w.Write([]byte(`{"entry_refers":[{"id":9,"name":"矩阵","definition":"线性变换"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/all-activities":
			if got := r.URL.Query().Get("module_ids"); got != "[1,2]" {
				t.Fatalf("unexpected module_ids: %q", got)
			}
			if got := r.URL.Query().Get("activity_types"); got != "homework,exams" {
				t.Fatalf("unexpected activity_types: %q", got)
			}
			if got := r.URL.Query().Get("no-loading-animation"); got != "true" {
				t.Fatalf("unexpected no-loading-animation: %q", got)
			}
			_, _ = w.Write([]byte(`{"activities":[{"id":8}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/live-records/21/task/asr":
			_, _ = w.Write([]byte(`{"task_id":"asr-1"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/live-records/21/task/asr":
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected live-record asr status query: %q", got)
			}
			_, _ = w.Write([]byte(`{"status":"running"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/live-records/21/caption":
			_, _ = w.Write([]byte(`{"captions":[{"id":1}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/blueprint/14/all-sub-activities-count":
			_, _ = w.Write([]byte(`{"count":12}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if resp, err := service.GetHomeworkSubmissionStatus(ctx, 14); err != nil || len(resp.HomeworkActivities) != 1 {
		t.Fatalf("unexpected homework submission status: %#v, err=%v", resp, err)
	}
	if resp, err := service.ListExamScores(ctx, 14); err != nil || len(resp.ExamScores) != 1 {
		t.Fatalf("unexpected exam scores: %#v, err=%v", resp, err)
	}
	if resp, err := service.ListCatalogActivities(ctx, 14); err != nil || len((*resp)["items"].([]any)) != 1 {
		t.Fatalf("unexpected catalog activities: %#v, err=%v", resp, err)
	}
	if resp, err := service.ListCatalogActivitiesAnonymous(ctx, 14); err != nil || len((*resp)["items"].([]any)) != 1 {
		t.Fatalf("unexpected anonymous catalog activities: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetHomeworkSubmissionNumber(ctx, 14); err != nil || (*resp)["submitted"] != float64(30) {
		t.Fatalf("unexpected homework submission number: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetClassroomExamScores(ctx, 14); err != nil || len((*resp)["scores"].([]any)) != 1 {
		t.Fatalf("unexpected classroom exam scores: %#v, err=%v", resp, err)
	}
	if resp, err := service.ListExamSubmissions(ctx, 14); err != nil || len((*resp)["submissions"].([]any)) != 1 {
		t.Fatalf("unexpected exam submissions: %#v, err=%v", resp, err)
	}
	if resp, err := service.CreateSHTVUGroups(ctx, 14); err != nil || (*resp)["created"] != true {
		t.Fatalf("unexpected create groups response: %#v, err=%v", resp, err)
	}
	if resp, err := service.ListQuestionnaires(ctx, 14); err != nil || len((*resp)["questionnaires"].([]any)) != 1 {
		t.Fatalf("unexpected questionnaires: %#v, err=%v", resp, err)
	}
	if resp, err := service.ListQuestionnairesWithParams(ctx, 14, &ListQuestionnairesParams{
		Page:       2,
		PageSize:   10,
		Conditions: map[string]any{"keyword": "survey"},
	}); err != nil || len((*resp)["items"].([]any)) != 1 {
		t.Fatalf("unexpected questionnaire list: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetQuestionnaireScores(ctx, 14); err != nil || len((*resp)["scores"].([]any)) != 1 {
		t.Fatalf("unexpected questionnaire scores: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetCourseWebLinkScores(ctx, 14); err != nil || len((*resp)["scores"].([]any)) != 1 {
		t.Fatalf("unexpected web-link scores: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetCourseVirtualExperimentScores(ctx, 14); err != nil || len((*resp)["scores"].([]any)) != 1 {
		t.Fatalf("unexpected virtual-experiment scores: %#v, err=%v", resp, err)
	}
	if resp, err := service.ListCourseVirtualExperiments(ctx, 14); err != nil || len((*resp)["items"].([]any)) != 1 {
		t.Fatalf("unexpected virtual experiments: %#v, err=%v", resp, err)
	}
	if resp, err := service.ListEntryRefers(ctx, 14); err != nil || len(resp.EntryRefers) != 1 || resp.EntryRefers[0].Name != "矩阵" {
		t.Fatalf("unexpected entry refers: %#v, err=%v", resp, err)
	}
	if resp, err := service.ListAllActivitiesByModuleIDs(ctx, 14, &ListAllActivitiesByModuleIDsParams{
		ModuleIDs:               []int{1, 2},
		ActivityTypes:           "homework,exams",
		DisableLoadingAnimation: true,
	}); err != nil || len((*resp)["activities"].([]any)) != 1 {
		t.Fatalf("unexpected all activities response: %#v, err=%v", resp, err)
	}
	if resp, err := service.CreateLiveRecordAsrTask(ctx, 21); err != nil || resp["task_id"] != "asr-1" {
		t.Fatalf("unexpected live-record asr task response: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetLiveRecordAsrTaskStatus(ctx, 21); err != nil || resp["status"] != "running" {
		t.Fatalf("unexpected live-record asr status: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetLiveRecordCaptions(ctx, 21); err != nil || len(resp["captions"].([]any)) != 1 {
		t.Fatalf("unexpected live-record captions: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetBlueprintAllSubActivitiesCount(ctx, 14); err != nil || (*resp)["count"] != float64(12) {
		t.Fatalf("unexpected blueprint all-sub-activities count: %#v, err=%v", resp, err)
	}
}
