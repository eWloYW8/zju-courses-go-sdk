package calendar

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestCalendarFrontendHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/calendar-events":
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected no-intercept: %q", got)
			}
			if got := r.URL.Query().Get("start"); got != "2026-03-01" {
				t.Fatalf("unexpected start: %q", got)
			}
			_, _ = w.Write([]byte(`[{"id":1}]`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/calendar-timetables":
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected no-intercept: %q", got)
			}
			if got := r.URL.Query().Get("view"); got != "month" {
				t.Fatalf("unexpected view: %q", got)
			}
			_, _ = w.Write([]byte(`[{"id":2}]`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/calendar-timetables/7/delete":
			var body DeleteTimetableRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode delete timetable body: %v", err)
			}
			if body.TargetTime != "2026-03-09T08:00:00Z" || body.UpdateScope != "current" {
				t.Fatalf("unexpected delete timetable body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"message":"deleted","timetable":{"id":7,"title":"实验课"}}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/departments":
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected departments no-intercept: %q", got)
			}
			_, _ = w.Write([]byte(`{"departments":[{"id":3,"name":"计算机学院"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/calendar-events/users/9/courses":
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected user-courses no-intercept: %q", got)
			}
			_, _ = w.Write([]byte(`{"courses":[{"id":8,"name":"算法设计","course_code":"CS101","end_date":"2026-06-30"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/calendar-meeting":
			_, _ = w.Write([]byte(`{"meeting_types":["teaching"]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/management/calendar-meeting":
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected page: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "20" {
				t.Fatalf("unexpected page_size: %q", got)
			}
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode management list body: %v", err)
			}
			if body["keyword"] != "AI" {
				t.Fatalf("unexpected management list body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"items":[{"id":5}],"page":2,"page_size":20,"pages":1,"total":1}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/management/calendar-meeting/excel":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode export body: %v", err)
			}
			if body["keyword"] != "AI" {
				t.Fatalf("unexpected export body: %#v", body)
			}
			w.Header().Set("Content-Type", "application/vnd.ms-excel")
			_, _ = w.Write([]byte("excel-bytes"))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	events, err := service.ListCalendarEvents(ctx, map[string]string{"start": "2026-03-01"})
	if err != nil || string(events) != `[{"id":1}]` {
		t.Fatalf("unexpected calendar events: %s, err=%v", string(events), err)
	}

	timetables, err := service.ListTimetables(ctx, map[string]string{"view": "month"})
	if err != nil || string(timetables) != `[{"id":2}]` {
		t.Fatalf("unexpected timetables: %s, err=%v", string(timetables), err)
	}

	deleted, err := service.DeleteTimetable(ctx, 7, &DeleteTimetableRequest{
		TargetTime:  "2026-03-09T08:00:00Z",
		UpdateScope: "current",
	})
	if err != nil || deleted.Message != "deleted" || deleted.Timetable == nil || deleted.Timetable.ID != 7 {
		t.Fatalf("unexpected delete response: %#v, err=%v", deleted, err)
	}

	departments, err := service.ListCalendarDepartments(ctx)
	if err != nil || len(departments.Departments) != 1 || departments.Departments[0].Name != "计算机学院" {
		t.Fatalf("unexpected departments: %#v, err=%v", departments, err)
	}

	courses, err := service.ListCalendarUserCourses(ctx, 9)
	if err != nil || len(courses.Courses) != 1 || courses.Courses[0].CourseCode != "CS101" {
		t.Fatalf("unexpected user courses: %#v, err=%v", courses, err)
	}
	if courses.Courses[0].EndDate == nil || *courses.Courses[0].EndDate != "2026-06-30" {
		t.Fatalf("course end_date did not decode: %#v", courses.Courses[0].EndDate)
	}

	meeting, err := service.GetCalendarMeeting(ctx)
	if err != nil || string(meeting) != `{"meeting_types":["teaching"]}` {
		t.Fatalf("unexpected calendar meeting payload: %s, err=%v", string(meeting), err)
	}

	list, err := service.ListManagementCalendarMeetingsWithBody(ctx, &model.ListOptions{Page: 2, PageSize: 20}, map[string]any{"keyword": "AI"})
	if err != nil || string(list) != `{"items":[{"id":5}],"page":2,"page_size":20,"pages":1,"total":1}` {
		t.Fatalf("unexpected management list: %s, err=%v", string(list), err)
	}

	exported, err := service.ExportCalendarMeetingsWithBody(ctx, map[string]any{"keyword": "AI"})
	if err != nil || string(exported) != "excel-bytes" {
		t.Fatalf("unexpected export payload: %q, err=%v", string(exported), err)
	}
}
