package exams

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

func TestListCourseSubjectLibsUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/12/subject-libs" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("with_folder"); got != "1" {
			t.Fatalf("unexpected with_folder: %s", got)
		}
		return jsonResponse(`{"subject_libs":[]}`), nil
	})

	if _, err := svc.ListCourseSubjectLibs(context.Background(), 12, true); err != nil {
		t.Fatalf("ListCourseSubjectLibs error: %v", err)
	}
}

func TestCreateCourseSubjectLibUsesLibTypeQuery(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/7/subject-libs" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("lib_type"); got != "exam" {
			t.Fatalf("unexpected lib_type: %s", got)
		}
		return jsonResponse(`{"id":1,"title":"lib"}`), nil
	})

	if _, err := svc.CreateCourseSubjectLib(context.Background(), 7, "exam", &CreateSubjectLibRequest{Title: "lib"}); err != nil {
		t.Fatalf("CreateCourseSubjectLib error: %v", err)
	}
}

func TestCopySubjectLibToExamUsesCopyEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/subject-libs/99/copy" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("examId"); got != "8" {
			t.Fatalf("unexpected examId: %s", got)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.CopySubjectLibToExam(context.Background(), 8, 99); err != nil {
		t.Fatalf("CopySubjectLibToExam error: %v", err)
	}
}

func TestGetCoursewareQuizSubjectsUsesSubjectsEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courseware-quiz/quiz/11/subjects" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"subjects":[]}`), nil
	})

	if _, err := svc.GetCoursewareQuizSubjects(context.Background(), 11); err != nil {
		t.Fatalf("GetCoursewareQuizSubjects error: %v", err)
	}
}

func TestUpdateCoursewareQuizSubjectsUsesPutEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courseware-quiz/quiz/11/subjects" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"quiz_id":11}`), nil
	})

	resp, err := svc.UpdateCoursewareQuizSubjects(context.Background(), 11, &CoursewareQuizSubjectsRequest{Subjects: []any{}})
	if err != nil {
		t.Fatalf("UpdateCoursewareQuizSubjects error: %v", err)
	}
	if resp.QuizID != 11 {
		t.Fatalf("unexpected quiz_id: %d", resp.QuizID)
	}
}
