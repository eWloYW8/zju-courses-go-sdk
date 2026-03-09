package statistics

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestExportONOStatStudentsUsesFrontendEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/cooc/stat-students/export/excel" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("start_date"); got != "2026-03-01" {
			t.Fatalf("unexpected start_date: %s", got)
		}
		if got := r.URL.Query().Get("end_date"); got != "2026-03-31" {
			t.Fatalf("unexpected end_date: %s", got)
		}
		if got := r.URL.Query().Get("conditions"); got != `{"department_id":7,"keyword":"张三"}` {
			t.Fatalf("unexpected conditions: %s", got)
		}
		var body ExportONOStatStudentsRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body.CourseIDs != "12,18" {
			t.Fatalf("unexpected body: %#v", body)
		}
		w.Header().Set("Content-Type", "application/vnd.ms-excel")
		_, _ = w.Write([]byte("excel-bytes"))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	data, err := service.ExportONOStatStudents(context.Background(), "excel", ExportONOStatStudentsParams{
		StartDate: "2026-03-01",
		EndDate:   "2026-03-31",
		Conditions: map[string]any{
			"department_id": 7,
			"keyword":       "张三",
		},
	}, ExportONOStatStudentsRequest{CourseIDs: "12,18"})
	if err != nil {
		t.Fatalf("ExportONOStatStudents returned error: %v", err)
	}
	if string(data) != "excel-bytes" {
		t.Fatalf("unexpected export bytes: %q", string(data))
	}
}

func TestVTRSStatisticsUseFrontendDateRangeQueries(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/stat/vtrses/9/overview", "/api/stat/vtrses/9/resources/rank", "/api/stat/vtrses/9/trial-teaching", "/api/stat/vtrses/9/resources":
			if got := r.URL.Query().Get("date_range"); got != "2026-03-01,2026-03-09" {
				t.Fatalf("unexpected date_range query for %s: %q", r.URL.Path, got)
			}
			_, _ = w.Write([]byte(`{"ok":true}`))
		case "/api/stat/vtrses/9/team-activation":
			if got := r.URL.Query().Get("date_range"); got != "2026-03-01,2026-03-09" {
				t.Fatalf("unexpected team-activation date_range query: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "20" {
				t.Fatalf("unexpected team-activation page_size query: %q", got)
			}
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected team-activation page query: %q", got)
			}
			_, _ = w.Write([]byte(`{"ok":true}`))
		default:
			t.Fatalf("unexpected request: %s", r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if _, err := service.GetVTRSOverview(ctx, 9, "2026-03-01,2026-03-09"); err != nil {
		t.Fatalf("GetVTRSOverview returned error: %v", err)
	}
	if _, err := service.GetVTRSResourcesRank(ctx, 9, "2026-03-01,2026-03-09"); err != nil {
		t.Fatalf("GetVTRSResourcesRank returned error: %v", err)
	}
	if _, err := service.GetVTRSTrialTeaching(ctx, 9, "2026-03-01,2026-03-09"); err != nil {
		t.Fatalf("GetVTRSTrialTeaching returned error: %v", err)
	}
	if _, err := service.GetVTRSResources(ctx, 9, "2026-03-01,2026-03-09"); err != nil {
		t.Fatalf("GetVTRSResources returned error: %v", err)
	}
	if _, err := service.GetVTRSTeamActivation(ctx, 9, map[string]string{
		"page":      "2",
		"pageSize":  "20",
		"dateRange": "2026-03-01,2026-03-09",
	}); err != nil {
		t.Fatalf("GetVTRSTeamActivation returned error: %v", err)
	}
}

func TestTrackingHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodOptions && r.URL.Path == "/api/user-visits":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPost && r.URL.Path == "/api/user-visits":
			if got := r.URL.Query().Get("jwt"); got != "jwt-token" {
				t.Fatalf("unexpected user-visits jwt: %q", got)
			}
			var body TrackingRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode user-visits body: %v", err)
			}
			if body["course_id"] != float64(18) {
				t.Fatalf("unexpected user-visits body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPost && r.URL.Path == "/api/user-reply-visits":
			if got := r.URL.Query().Get("jwt"); got != "jwt-token" {
				t.Fatalf("unexpected user-reply-visits jwt: %q", got)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPost && r.URL.Path == "/api/zhiyun-visits":
			if got := r.URL.Query().Get("jwt"); got != "jwt-token" {
				t.Fatalf("unexpected zhiyun-visits jwt: %q", got)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPost && r.URL.Path == "/api/vtrs":
			if got := r.URL.Query().Get("jwt"); got != "jwt-token" {
				t.Fatalf("unexpected vtrs jwt: %q", got)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPost && r.URL.Path == "/api/vtrs-resource":
			if got := r.URL.Query().Get("jwt"); got != "jwt-token" {
				t.Fatalf("unexpected vtrs-resource jwt: %q", got)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if err := service.CheckUserVisitsAccessible(ctx); err != nil {
		t.Fatalf("CheckUserVisitsAccessible returned error: %v", err)
	}
	if err := service.TrackUserVisit(ctx, "jwt-token", TrackingRequest{"course_id": 18}); err != nil {
		t.Fatalf("TrackUserVisit returned error: %v", err)
	}
	if err := service.TrackUserReplyVisit(ctx, "jwt-token", TrackingRequest{"reply_id": 7}); err != nil {
		t.Fatalf("TrackUserReplyVisit returned error: %v", err)
	}
	if err := service.TrackZhiyunVisit(ctx, "jwt-token", TrackingRequest{"action_type": "visit_home"}); err != nil {
		t.Fatalf("TrackZhiyunVisit returned error: %v", err)
	}
	if err := service.TrackVTRSVisit(ctx, "jwt-token", TrackingRequest{"vtrs_id": 9}); err != nil {
		t.Fatalf("TrackVTRSVisit returned error: %v", err)
	}
	if err := service.TrackVTRSResourceVisit(ctx, "jwt-token", TrackingRequest{"resource_id": 11}); err != nil {
		t.Fatalf("TrackVTRSResourceVisit returned error: %v", err)
	}
}
