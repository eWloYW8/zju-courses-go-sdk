package forum

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestListLatestTopicsDecodesCourseID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/topics/latest" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("no-intercept") != "true" || r.URL.Query().Get("count") != "2" {
			t.Fatalf("unexpected query: %s", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"topics":[{"id":9,"course_id":18,"group_id":3,"title":"讨论帖"}]}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	topics, err := service.ListLatestTopics(context.Background(), 2)
	if err != nil {
		t.Fatalf("ListLatestTopics returned error: %v", err)
	}
	if len(topics) != 1 || topics[0].CourseID != 18 || topics[0].GroupID == nil || *topics[0].GroupID != 3 {
		t.Fatalf("unexpected topics: %#v", topics)
	}
}
