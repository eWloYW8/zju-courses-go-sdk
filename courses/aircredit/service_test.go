package aircredit

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestTypedCreditStateHelpersUseFrontendModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/air-credit/user/credit-states":
			conditions, err := url.QueryUnescape(r.URL.Query().Get("conditions"))
			if err != nil {
				t.Fatalf("unescape user credit-state conditions: %v", err)
			}
			if conditions != `{"keyword":"alice"}` {
				t.Fatalf("unexpected user credit-state conditions: %s", conditions)
			}
			_, _ = w.Write([]byte(`{
				"items":[{"user_id":7,"user_no":"2025001","user_name":"Alice","department":"CS","role":"student","credit_used_percent":12.5,"is_low_air_credit":true,"credit_state":{"credit_assigned":100,"credit_used":25,"credit_remaining":75,"credit_limit":200,"status":"normal","has_credit_limit":true}}],
				"page":2,
				"page_size":3,
				"pages":4,
				"total":10
			}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/air-credit/course/credit-states":
			conditions, err := url.QueryUnescape(r.URL.Query().Get("conditions"))
			if err != nil {
				t.Fatalf("unescape course credit-state conditions: %v", err)
			}
			if conditions != `{"keyword":"math"}` {
				t.Fatalf("unexpected course credit-state conditions: %s", conditions)
			}
			_, _ = w.Write([]byte(`{
				"items":[{"course_id":9,"name":"Calculus","department":"Math","instructors":"Prof. Li","semester":"秋","academic_year":"2025-2026","course_code":"MATH100","course_type":1,"credit_used_percent":30,"credit_state":{"credit_assigned":50,"credit_used":15,"credit_remaining":35,"status":"normal"}}],
				"page":1,
				"page_size":5,
				"pages":1,
				"total":1
			}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/air-credit/credit-states-stats":
			if got := r.URL.Query().Get("type"); got == "user" {
				conditions, err := url.QueryUnescape(r.URL.Query().Get("conditions"))
				if err != nil {
					t.Fatalf("unescape user stats conditions: %v", err)
				}
				if conditions != `{"keyword":"alice"}` {
					t.Fatalf("unexpected user stats conditions: %s", conditions)
				}
				_, _ = w.Write([]byte(`{
					"items":[{"user":{"id":7,"name":"Alice","user_no":"2025001","department":{"id":4,"name":"CS"}},"user_role":"student","credit_assigned":100,"credit_used":25,"module_credit_used":{"material":3,"Chat":4},"usage_count":6}],
					"page":1,
					"page_size":10,
					"pages":1,
					"total":1
				}`))
				return
			}
			if got := r.URL.Query().Get("type"); got == "course" {
				_, _ = w.Write([]byte(`{
					"items":[{"course":{"id":9,"name":"Calculus","course_code":"MATH100","course_type":1,"academic_year":{"id":1,"name":"2025-2026"},"semester":{"id":2,"name":"秋"},"klass":{"id":3,"name":"甲班"},"department":{"id":4,"name":"Math"},"instructors":[{"id":2,"name":"Prof. Li"}],"credit_state":{"credit_assigned":50,"credit_used":15,"credit_remaining":35,"status":"normal"},"ai_activation":"enabled"},"instructors":1,"credit_used":15,"usage_count":20,"students_count":60,"use_air_chat_students_count":8}],
					"page":1,
					"page_size":10,
					"pages":1,
					"total":1
				}`))
				return
			}
			t.Fatalf("unexpected stats type: %s", r.URL.Query().Get("type"))
		case r.Method == http.MethodGet && r.URL.Path == "/api/air-credit/org/credit-state-info":
			_, _ = w.Write([]byte(`{"org_id":9,"credit_assigned":300,"credit_used":120,"user_credit_used":50,"user_credit_assigned":100,"course_credit_used":70,"course_credit_assigned":150}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/air-credit/audits":
			conditions, err := url.QueryUnescape(r.URL.Query().Get("conditions"))
			if err != nil {
				t.Fatalf("unescape audit conditions: %v", err)
			}
			if conditions != `{"status":"waiting"}` {
				t.Fatalf("unexpected audit conditions: %s", conditions)
			}
			_, _ = w.Write([]byte(`{
				"items":[{"id":4,"target_type":"course","status":"waiting","applied_credits":10,"approved_credits":0,"reason":"need credits","remark":"","created_at":"2026-03-01","updated_at":"2026-03-02","read":false,"user":{"id":7,"name":"Alice","user_no":"2025001"},"course":{"id":9,"name":"Calculus","course_code":"MATH100"},"auditor":{"id":2,"name":"Admin"}}],
				"page":1,
				"page_size":5,
				"pages":1,
				"total":1
			}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	userStates, err := service.ListUserCreditStatesTyped(ctx, ListCreditStatesParams{Page: 2, PageSize: 3, Conditions: map[string]any{"keyword": "alice"}})
	if err != nil || len(userStates.Items) != 1 || userStates.Items[0].CreditState == nil {
		t.Fatalf("unexpected user credit states: %#v, err=%v", userStates, err)
	}
	if userStates.Items[0].CreditState.CreditAssigned != 100 || userStates.Items[0].CreditState.CreditRemaining != 75 {
		t.Fatalf("user credit state did not decode: %#v", userStates.Items[0].CreditState)
	}

	courseStates, err := service.ListCourseCreditStatesTyped(ctx, ListCreditStatesParams{Page: 1, PageSize: 5, Conditions: map[string]any{"keyword": "math"}})
	if err != nil || len(courseStates.Items) != 1 || courseStates.Items[0].CourseCode != "MATH100" {
		t.Fatalf("unexpected course credit states: %#v, err=%v", courseStates, err)
	}

	userStats, err := service.GetUserCreditStatesStatsTyped(ctx, CreditStateStatsParams{Page: 1, PageSize: 10, Conditions: map[string]any{"keyword": "alice"}})
	if err != nil || len(userStats.Items) != 1 || userStats.Items[0].User == nil || userStats.Items[0].ModuleCreditUsed["Chat"] != 4 {
		t.Fatalf("unexpected user credit stats: %#v, err=%v", userStats, err)
	}

	courseStats, err := service.GetCourseCreditStatesStatsTyped(ctx, CreditStateStatsParams{Page: 1, PageSize: 10})
	if err != nil || len(courseStats.Items) != 1 || courseStats.Items[0].Course == nil || courseStats.Items[0].Course.CreditState == nil {
		t.Fatalf("unexpected course credit stats: %#v, err=%v", courseStats, err)
	}

	orgInfo, err := service.GetOrgCreditStateInfoTyped(ctx)
	if err != nil || orgInfo.OrgID != 9 || orgInfo.CourseCreditAssigned != 150 {
		t.Fatalf("unexpected org credit state info: %#v, err=%v", orgInfo, err)
	}

	audits, err := service.ListAuditsTyped(ctx, ListAuditsParams{Page: 1, PageSize: 5, Conditions: map[string]any{"status": "waiting"}})
	if err != nil || len(audits.Items) != 1 || audits.Items[0].Course == nil || audits.Items[0].TargetType != "course" {
		t.Fatalf("unexpected audits response: %#v, err=%v", audits, err)
	}
}
