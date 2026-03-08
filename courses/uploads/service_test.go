package uploads

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestGetUploadPDFInfoDetailUsesPreviewEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/uploads/77/pdf-info" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("preview"); got != "true" {
			t.Fatalf("unexpected preview query: %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"num_pages":25}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	info, err := service.GetUploadPDFInfoDetail(context.Background(), 77)
	if err != nil {
		t.Fatalf("GetUploadPDFInfoDetail returned error: %v", err)
	}
	if info.NumPages != 25 {
		t.Fatalf("unexpected pdf info: %#v", info)
	}
}

func TestListMoodlePackagesUsesFrontendQueryAndModel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/uploads/moodle-pkg" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("page"); got != "2" {
			t.Fatalf("unexpected page query: %q", got)
		}
		if got := r.URL.Query().Get("page_size"); got != "5" {
			t.Fatalf("unexpected page_size query: %q", got)
		}
		if got := r.URL.Query().Get("conditions"); got != `{"keyword":"calc"}` {
			t.Fatalf("unexpected conditions query: %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"page":2,
			"page_size":5,
			"pages":3,
			"total":11,
			"start":6,
			"end":10,
			"items":[{"id":9,"data":{"course_name":"Calculus","course_code":"MATH100","academic_year":"2025-2026","semester":"Autumn","department":"Math","college":"Science","spoc_course_name":"Calc SPOC","instructors":["Alice","Bob"],"use_count":4}}]
		}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.ListMoodlePackages(context.Background(), &model.ListOptions{Page: 2, PageSize: 5}, `{"keyword":"calc"}`)
	if err != nil {
		t.Fatalf("ListMoodlePackages returned error: %v", err)
	}
	if result.Page != 2 || len(result.Items) != 1 || result.Items[0].Data == nil || result.Items[0].Data.SpocCourseName != "Calc SPOC" {
		t.Fatalf("unexpected moodle packages response: %#v", result)
	}
}
