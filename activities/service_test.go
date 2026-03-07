package activities

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

func newTestService(t *testing.T, fn roundTripFunc) *Service {
	t.Helper()
	client := sdk.NewClient(sdk.WithHTTPClient(&http.Client{Transport: fn}))
	return New(client)
}

func jsonResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestCreateActivityUsesCoursesActivitiesPath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courses/123/activities" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":1,"title":"demo","type":"page"}`), nil
	})

	if _, err := svc.CreateActivity(context.Background(), 123, &CreateActivityRequest{Title: "demo", Type: "page"}); err != nil {
		t.Fatalf("CreateActivity error: %v", err)
	}
}

func TestDeleteCheckWithTypeUsesActivityTypeQuery(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/activities/delete-check" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("activity_id"); got != "12" {
			t.Fatalf("unexpected activity_id: %s", got)
		}
		if got := req.URL.Query().Get("activity_type"); got != "homework" {
			t.Fatalf("unexpected activity_type: %s", got)
		}
		return jsonResponse(`{"safe_delete":true}`), nil
	})

	resp, err := svc.DeleteCheckWithType(context.Background(), 12, "homework")
	if err != nil {
		t.Fatalf("DeleteCheckWithType error: %v", err)
	}
	if !resp.SafeDelete {
		t.Fatal("expected safe_delete=true")
	}
}

func TestHaveDependentsWithTypeUsesGetQuery(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/activities/have-dependents" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		values := req.URL.Query()["activity_ids"]
		if len(values) != 2 || values[0] != "3" || values[1] != "5" {
			t.Fatalf("unexpected activity_ids: %#v", values)
		}
		if got := req.URL.Query().Get("activity_type"); got != "forum" {
			t.Fatalf("unexpected activity_type: %s", got)
		}
		return jsonResponse(`{"has_dependents":true}`), nil
	})

	resp, err := svc.HaveDependentsWithType(context.Background(), []int{3, 5}, "forum")
	if err != nil {
		t.Fatalf("HaveDependentsWithType error: %v", err)
	}
	if !resp.HasDependents {
		t.Fatal("expected has_dependents=true")
	}
}

func TestDeleteActivityWithOptionsUsesQuery(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/activities/88" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("activity_type"); got != "exam" {
			t.Fatalf("unexpected activity_type: %s", got)
		}
		if got := req.URL.Query().Get("keep_original"); got != "true" {
			t.Fatalf("unexpected keep_original: %s", got)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.DeleteActivityWithOptions(context.Background(), 88, &DeleteActivityOptions{
		ActivityType: "exam",
		KeepOriginal: true,
	})
	if err != nil {
		t.Fatalf("DeleteActivityWithOptions error: %v", err)
	}
}

func TestBatchDeleteActivitiesUsesDeleteBody(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/activities" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		body, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("read body: %v", err)
		}
		if string(body) != `{"activity_ids":[1,2,3]}` {
			t.Fatalf("unexpected body: %s", string(body))
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.BatchDeleteActivities(context.Background(), []int{1, 2, 3}); err != nil {
		t.Fatalf("BatchDeleteActivities error: %v", err)
	}
}
