package rollcall

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

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
