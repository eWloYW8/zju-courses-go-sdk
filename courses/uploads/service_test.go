package uploads

import (
	"context"
	"encoding/json"
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

func TestDuplicateDetectHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/duplicate-detect/file/file-a/raw":
			_, _ = w.Write([]byte("duplicate raw text"))
		case r.Method == http.MethodPost && r.URL.Path == "/api/duplicate-detect/report/download":
			var body DuplicateReportDownloadRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode report body: %v", err)
			}
			if body.ReportType != "p" || body.DetectKey != "detect-2" || body.Provider != "duplicate_lib" {
				t.Fatalf("unexpected report body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"status":"success","download_url":"https://example.com/report.pdf"}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))

	raw, err := service.CheckDuplicate(context.Background(), "file-a")
	if err != nil || string(raw) != "duplicate raw text" {
		t.Fatalf("unexpected duplicate raw response: %q, err=%v", string(raw), err)
	}

	info, err := service.DownloadDuplicateReport(context.Background(), &DuplicateReportDownloadRequest{
		ReportType: "p",
		DetectKey:  "detect-2",
		Provider:   "duplicate_lib",
	})
	if err != nil || info.Status != "success" || info.DownloadURL != "https://example.com/report.pdf" {
		t.Fatalf("unexpected report download info: %#v, err=%v", info, err)
	}
}

func TestDeleteUploadNoInterceptUsesFrontendQuery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/uploads/44" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("no-intercept"); got != "true" {
			t.Fatalf("unexpected no-intercept query: %q", got)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	if err := service.DeleteUploadNoIntercept(context.Background(), 44); err != nil {
		t.Fatalf("DeleteUploadNoIntercept returned error: %v", err)
	}
}

func TestValidateH5CoursewareUnzipUsesFrontendEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/h5-courseware/upload/31/validate-unzip" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"enable_set_h5_courseware_completion":true}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.ValidateH5CoursewareUnzip(context.Background(), 31)
	if err != nil {
		t.Fatalf("ValidateH5CoursewareUnzip returned error: %v", err)
	}
	if !result.EnableSetH5CoursewareCompletion {
		t.Fatalf("unexpected validate response: %#v", result)
	}
}

func TestStartUploadEsignUsesFrontendEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/uploads/31/start-esign" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"started"}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	resp, err := service.StartUploadEsign(context.Background(), 31)
	if err != nil || (*resp)["status"] != "started" {
		t.Fatalf("unexpected start-esign response: %#v, err=%v", resp, err)
	}
}

func TestSCORMHelpersMatchFrontendRuntime(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/activity/12/scorm/sco-1/cmi":
			_, _ = w.Write([]byte(`{"cmi":{"core":{"lesson_status":"completed"}},"suspend_data":{"p":{"1":{"visited":true,"completed":true}}},"total_pages":5,"visited_pages":4,"completed_pages":3}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/activity/12/scorm/sco-1/cmi":
			var body SCORMCMIData
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode scorm cmi body: %v", err)
			}
			if body.TotalPages != 5 || body.VisitedPages != 4 || body.CompletedPages != 3 {
				t.Fatalf("unexpected scorm cmi body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	data, err := service.GetSCORMCMI(ctx, 12, "sco-1")
	if err != nil || data.TotalPages != 5 || data.CompletedPages != 3 {
		t.Fatalf("unexpected GetSCORMCMI response: %#v, err=%v", data, err)
	}
	if err := service.UpdateSCORMCMI(ctx, 12, "sco-1", &SCORMCMIData{
		CMI:            map[string]any{"core": map[string]any{"lesson_status": "completed"}},
		SuspendData:    map[string]any{"p": map[string]any{"1": map[string]any{"visited": true, "completed": true}}},
		TotalPages:     5,
		VisitedPages:   4,
		CompletedPages: 3,
	}); err != nil {
		t.Fatalf("UpdateSCORMCMI returned error: %v", err)
	}
	if got := service.BuildSCORMPreviewURL(31, "index_lms.html", `{"attempt":1}`); got != "/api/uploads/scorm/31?sco=index_lms.html&preview=true&para=%7B%22attempt%22%3A1%7D" {
		t.Fatalf("unexpected scorm preview url: %s", got)
	}
}
