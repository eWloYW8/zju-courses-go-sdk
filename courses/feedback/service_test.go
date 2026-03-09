package feedback

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestFeedbackTypedHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/feedback-activities/9":
			_, _ = w.Write([]byte(`{"id":9,"type":"feedback","title":"课堂反馈","data":{"is_allow_add_feedback":true}}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/courses/8/feedback-activities":
			_, _ = w.Write([]byte(`{"id":10,"type":"feedback","title":"新反馈"}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/feedback-activities/10":
			_, _ = w.Write([]byte(`{"id":10,"type":"feedback","title":"已更新反馈"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/feedbacks/7":
			_, _ = w.Write([]byte(`{"id":7,"content":"第一条反馈","created_by":{"id":2,"name":"Teacher"}}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/feedback-activities/9/feedbacks":
			_, _ = w.Write([]byte(`{"id":11,"content":"第二条反馈","user":{"id":3,"name":"Alice"}}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/feedback-activities/9/feedbacks/11":
			_, _ = w.Write([]byte(`{"id":11,"content":"已修改反馈"}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/api/feedbacks/11":
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	activity, err := service.GetFeedbackActivityTyped(ctx, 9)
	if err != nil || activity.ID != 9 || activity.Type != "feedback" {
		t.Fatalf("unexpected feedback activity: %#v, err=%v", activity, err)
	}

	createdActivity, err := service.CreateFeedbackActivityTyped(ctx, 8, map[string]any{"title": "新反馈"})
	if err != nil || createdActivity.ID != 10 || createdActivity.Title != "新反馈" {
		t.Fatalf("unexpected created feedback activity: %#v, err=%v", createdActivity, err)
	}

	updatedActivity, err := service.UpdateFeedbackActivityTyped(ctx, 10, map[string]any{"title": "已更新反馈"})
	if err != nil || updatedActivity.ID != 10 || updatedActivity.Title != "已更新反馈" {
		t.Fatalf("unexpected updated feedback activity: %#v, err=%v", updatedActivity, err)
	}

	feedbackItem, err := service.GetFeedbackTyped(ctx, 7)
	if err != nil || feedbackItem.ID != 7 || feedbackItem.CreatedBy == nil || feedbackItem.CreatedBy.ID != 2 {
		t.Fatalf("unexpected feedback item: %#v, err=%v", feedbackItem, err)
	}

	createdFeedback, err := service.CreateFeedbackTyped(ctx, 9, &FeedbackRequest{Content: "第二条反馈"})
	if err != nil || createdFeedback.ID != 11 || createdFeedback.User == nil || createdFeedback.User.ID != 3 {
		t.Fatalf("unexpected created feedback: %#v, err=%v", createdFeedback, err)
	}

	updatedFeedback, err := service.UpdateFeedbackTyped(ctx, 9, 11, &FeedbackRequest{Content: "已修改反馈"})
	if err != nil || updatedFeedback.ID != 11 || updatedFeedback.Content != "已修改反馈" {
		t.Fatalf("unexpected updated feedback: %#v, err=%v", updatedFeedback, err)
	}

	if err := service.DeleteFeedback(ctx, 11); err != nil {
		t.Fatalf("DeleteFeedback returned error: %v", err)
	}
}
