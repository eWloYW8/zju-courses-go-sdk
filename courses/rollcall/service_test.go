package rollcall

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestCreateRollcall(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/module/23/rollcall" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":61,"message":"ok"}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.CreateRollcall(context.Background(), 23, &CreateRollcallRequest{
		Status:   "waiting",
		CourseID: 23,
		ModuleID: 5,
		Title:    "点名",
		IsNumber: true,
	})
	if err != nil {
		t.Fatalf("CreateRollcall returned error: %v", err)
	}
	if result.ID != 61 || result.Message == nil || *result.Message != "ok" {
		t.Fatalf("unexpected result: %#v", result)
	}
}

func TestListCourseRollcallsUsesFrontendEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/course/23/rollcalls" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"rollcalls":[{"id":8,"title":"第1次点名","course_id":23}]}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.ListCourseRollcalls(context.Background(), 23)
	if err != nil {
		t.Fatalf("ListCourseRollcalls returned error: %v", err)
	}
	if len(result.Rollcalls) != 1 || result.Rollcalls[0].CourseID != 23 {
		t.Fatalf("unexpected rollcalls: %#v", result)
	}
}

func TestRollcallStudentHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/rollcall/15/student_rollcalls":
			if got := r.URL.Query().Get("action"); got != "manual_refresh" {
				t.Fatalf("unexpected action: %q", got)
			}
			_, _ = w.Write([]byte(`{"id":15,"title":"2026-03-09 09:00","status":"in_progress","type":"number","comment":"note","is_number":true,"student_rollcalls":[{"student_id":1,"status":"on_call_fine","status_detail":"ok","rollcall_status":"on_call_fine","temperature_status":"normal"}],"children":[{"id":16,"type":"qr_rollcall"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/timetable_rollcalls":
			query := r.URL.Query()
			if got := query["course_ids"]; len(got) != 2 || got[0] != "7" || got[1] != "9" {
				t.Fatalf("unexpected course_ids: %#v", got)
			}
			if got := query.Get("rollcall_date"); got != "2026-03-09" {
				t.Fatalf("unexpected rollcall_date: %q", got)
			}
			_, _ = w.Write([]byte(`[{"id":7,"isInstructor":true,"rollcall":{"id":51,"status":"in_progress","end_time":"2026-03-09T10:00:00Z"}}]`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/rollcall_status/23/result":
			_, _ = w.Write([]byte(`{"status":"in_progress"}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/rollcall/9/answer_self_registration_rollcall":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode self-registration body: %v", err)
			}
			if len(body) != 0 {
				t.Fatalf("unexpected self-registration body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"ok":true}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/rollcall/9/answer_number_rollcall":
			var body AnswerNumberRollcallRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode number body: %v", err)
			}
			if body.NumberCode != "123456" {
				t.Fatalf("unexpected number code: %#v", body)
			}
			_, _ = w.Write([]byte(`{"ok":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/23/student/88/rollcalls":
			if got := r.URL.RawQuery; got != "page=1&page_size=10" {
				t.Fatalf("unexpected query: %s", got)
			}
			_, _ = w.Write([]byte(`{"rollcalls":[{"student_rollcall_id":4,"student_status":"absent","status":"on_personal_leave","student_status_detail":"late","title":"2026-03-09 09:00","scored":true}]}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/course/23/student/88/rollcalls":
			var body UpdateCourseStudentRollcallsRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode update body: %v", err)
			}
			if len(body.StudentRollcalls) != 1 || body.StudentRollcalls[0].StudentRollcallID != 4 || body.StudentRollcalls[0].StudentStatus != "absent" {
				t.Fatalf("unexpected update body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"updated":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/23/leave-record":
			if got := r.URL.Query().Get("timestamp"); got != "2026-03-09T00:00:00Z" {
				t.Fatalf("unexpected timestamp: %q", got)
			}
			_, _ = w.Write([]byte(`{"user_nos":["20230001"]}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))

	rollcallDetail, err := service.GetRollcallStudentRollcallsDetail(context.Background(), 15, "manual_refresh")
	if err != nil {
		t.Fatalf("GetRollcallStudentRollcallsDetail returned error: %v", err)
	}
	if rollcallDetail.ID != 15 || len(rollcallDetail.StudentRollcalls) != 1 || len(rollcallDetail.Children) != 1 {
		t.Fatalf("unexpected rollcall detail: %#v", rollcallDetail)
	}

	timetable, err := service.ListTimetableRollcallsWithParams(context.Background(), &ListTimetableRollcallsParams{
		CourseIDs:    []int{7, 9},
		RollcallDate: "2026-03-09",
	})
	if err != nil {
		t.Fatalf("ListTimetableRollcallsWithParams returned error: %v", err)
	}
	if len(timetable) != 1 || !timetable[0].IsInstructor || timetable[0].Rollcall == nil || timetable[0].Rollcall.ID != 51 {
		t.Fatalf("unexpected timetable rollcalls: %#v", timetable)
	}

	statusResult, err := service.GetRollcallStatusResult(context.Background(), 23)
	if err != nil {
		t.Fatalf("GetRollcallStatusResult returned error: %v", err)
	}
	if statusResult["status"] != "in_progress" {
		t.Fatalf("unexpected status result: %#v", statusResult)
	}

	if _, err := service.AnswerSelfRegistrationRollcall(context.Background(), 9); err != nil {
		t.Fatalf("AnswerSelfRegistrationRollcall returned error: %v", err)
	}
	if _, err := service.AnswerNumberRollcall(context.Background(), 9, &AnswerNumberRollcallRequest{NumberCode: "123456"}); err != nil {
		t.Fatalf("AnswerNumberRollcall returned error: %v", err)
	}

	studentRollcalls, err := service.ListCourseStudentRollcalls(context.Background(), 23, 88, &model.ListOptions{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("ListCourseStudentRollcalls returned error: %v", err)
	}
	if len(studentRollcalls.Rollcalls) != 1 || studentRollcalls.Rollcalls[0].StudentRollcallID != 4 {
		t.Fatalf("unexpected student rollcalls: %#v", studentRollcalls)
	}

	updateResult, err := service.UpdateCourseStudentRollcalls(context.Background(), 23, 88, &UpdateCourseStudentRollcallsRequest{
		StudentRollcalls: []*UpdateCourseStudentRollcall{
			{StudentRollcallID: 4, StudentStatus: "absent"},
		},
	})
	if err != nil {
		t.Fatalf("UpdateCourseStudentRollcalls returned error: %v", err)
	}
	if updated, _ := updateResult["updated"].(bool); !updated {
		t.Fatalf("unexpected update result: %#v", updateResult)
	}

	leaveRecord, err := service.GetLeaveRecord(context.Background(), 23, "2026-03-09T00:00:00Z")
	if err != nil {
		t.Fatalf("GetLeaveRecord returned error: %v", err)
	}
	if len(leaveRecord.UserNos) != 1 || leaveRecord.UserNos[0] != "20230001" {
		t.Fatalf("unexpected leave record: %#v", leaveRecord)
	}
}
