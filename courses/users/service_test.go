package users

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestUpdateNickname(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/user/18/nickname" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	if err := service.UpdateNickname(context.Background(), 18, UpdateNicknameRequest{Nickname: "新昵称"}); err != nil {
		t.Fatalf("UpdateNickname returned error: %v", err)
	}
}

func TestGetRecentlyVisitedCoursesUsesVisitedCoursesWrapper(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/user/recently-visited-courses" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"visited_courses":[
				{"id":3,"name":"离散数学","url":"/course/3/forum#/","teaching_unit_type":"course","current_user_is_member":true}
			]
		}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.GetRecentlyVisitedCourses(context.Background())
	if err != nil {
		t.Fatalf("GetRecentlyVisitedCourses returned error: %v", err)
	}
	if len(result.Courses) != 1 || result.Courses[0].URL != "/course/3/forum#/" || result.Courses[0].TeachingUnitType != "course" || !result.Courses[0].CurrentUserIsMember {
		t.Fatalf("unexpected recently visited courses: %#v", result)
	}
}

func TestUpdateRecentlyVisitedCoursesUsesArrayPayload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/user/recently-visited-courses" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		var body []int
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if len(body) != 3 || body[0] != 7 || body[1] != 8 || body[2] != 9 {
			t.Fatalf("unexpected body: %#v", body)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	if err := service.UpdateRecentlyVisitedCourses(context.Background(), UpdateRecentlyVisitedCoursesRequest{7, 8, 9}); err != nil {
		t.Fatalf("UpdateRecentlyVisitedCourses returned error: %v", err)
	}
}
