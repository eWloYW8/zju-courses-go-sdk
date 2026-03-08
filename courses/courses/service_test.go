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
