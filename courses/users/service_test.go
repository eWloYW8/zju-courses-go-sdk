package users

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func TestChatHelpersUseFrontendModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path != "/api/user/chat" {
				t.Fatalf("unexpected GET path: %s", r.URL.Path)
			}
			_, _ = w.Write([]byte(`[
				{"message":"你好","is_reply":false,"type":"text","data":[],"session_id":"s-1"},
				{"message":"你可以试试这些问题","is_reply":true,"type":"recommend","data":["课程大纲","作业截止时间"],"session_id":"s-1"}
			]`))
		case http.MethodPost:
			if r.URL.Path != "/api/user/chat" {
				t.Fatalf("unexpected POST path: %s", r.URL.Path)
			}
			var body UserChatRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode body: %v", err)
			}
			if body.SessionID != "s-1" || body.Message != "课程大纲" {
				t.Fatalf("unexpected body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"message":"已为你打开课程大纲","is_reply":true,"type":"text","data":[],"session_id":"s-1"}`))
		default:
			t.Fatalf("unexpected method: %s", r.Method)
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	history, err := service.GetChat(context.Background())
	if err != nil {
		t.Fatalf("GetChat returned error: %v", err)
	}
	if len(history) != 2 || len(history[1].Data) != 2 || history[1].Data[0] != "课程大纲" {
		t.Fatalf("unexpected chat history: %#v", history)
	}

	reply, err := service.SendChatMessage(context.Background(), UserChatRequest{
		SessionID: "s-1",
		Message:   "课程大纲",
	})
	if err != nil {
		t.Fatalf("SendChatMessage returned error: %v", err)
	}
	if reply.SessionID != "s-1" || reply.Message != "已为你打开课程大纲" {
		t.Fatalf("unexpected chat reply: %#v", reply)
	}
}

func TestStartNotebookGradingUsesFrontendSSEEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/notebooks/12/grading" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Accept"); got != "text/event-stream" {
			t.Fatalf("unexpected accept header: %q", got)
		}
		var body NotebookGradingRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body.NotebookID != 12 || len(body.SubjectIDs) != 2 || body.SubjectIDs[0] != 3 || !body.Stream {
			t.Fatalf("unexpected notebook grading body: %#v", body)
		}
		w.Header().Set("Content-Type", "text/event-stream")
		_, _ = w.Write([]byte("data: {\"items\":[{\"notebook_id\":12,\"subject_id\":3,\"explanation\":\"ok\"}]}\n\n"))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	resp, err := service.StartNotebookGrading(context.Background(), 12, NotebookGradingRequest{
		NotebookID: 12,
		SubjectIDs: []int{3, 5},
		Stream:     true,
	})
	if err != nil {
		t.Fatalf("StartNotebookGrading returned error: %v", err)
	}
	_ = resp.Body.Close()
}

func TestSearchUsersWithPrefixUsesFrontendEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/anonymous-api/user/search" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			t.Fatalf("parse query: %v", err)
		}
		if got := values.Get("keywords"); got != "Alice" {
			t.Fatalf("unexpected keywords: %q", got)
		}
		if got := values.Get("limit"); got != "5" {
			t.Fatalf("unexpected limit: %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"id":9,"name":"Alice","user_no":"001"}]`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.SearchUsersWithPrefix(context.Background(), true, SearchUsersQuery{
		"keywords": "Alice",
		"limit":    "5",
	})
	if err != nil {
		t.Fatalf("SearchUsersWithPrefix returned error: %v", err)
	}
	if len(result) != 1 || result[0].ID != 9 || result[0].Name != "Alice" {
		t.Fatalf("unexpected search result: %#v", result)
	}
}

func TestGetUserEsignInfoUsesFrontendEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/users/18/esign-info" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"enabled":true}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	resp, err := service.GetUserEsignInfo(context.Background(), 18)
	if err != nil || (*resp)["enabled"] != true {
		t.Fatalf("unexpected esign-info response: %#v, err=%v", resp, err)
	}
}
