package notifications

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestHomepageNotificationModelsDecodeFrontendPayloads(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/todos":
			if r.URL.Query().Get("no-intercept") != "true" {
				t.Fatalf("unexpected todo query: %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`{
				"todo_list":[
					{"id":3,"title":"问卷待完成","type":"urp_pending_survey","course_id":8,"course_type":8,"course_name":"通识课","survey_url":"https://example.com/survey"}
				]
			}`))
		case "/api/bulletins/latest":
			_, _ = w.Write([]byte(`{
				"bulletins":[
					{"id":5,"title":"课程公告","course":{"id":8,"name":"高等数学"}}
				]
			}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))

	todos, err := service.ListTodosNoIntercept(context.Background())
	if err != nil {
		t.Fatalf("ListTodosNoIntercept returned error: %v", err)
	}
	if len(todos.TodoList) != 1 || todos.TodoList[0].CourseType != 8 || todos.TodoList[0].SurveyURL != "https://example.com/survey" {
		t.Fatalf("unexpected todos: %#v", todos)
	}

	bulletins, err := service.ListLatestBulletins(context.Background())
	if err != nil {
		t.Fatalf("ListLatestBulletins returned error: %v", err)
	}
	if len(bulletins.Bulletins) != 1 || bulletins.Bulletins[0].Course == nil || bulletins.Bulletins[0].Course.Name != "高等数学" {
		t.Fatalf("unexpected bulletins: %#v", bulletins)
	}
}
