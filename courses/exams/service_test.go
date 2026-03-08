package exams

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestCreateClassroomUsesCourseEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/courses/13/classroom-exams" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":7,"title":"随堂测验"}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.CreateClassroom(context.Background(), 13, map[string]any{"title": "随堂测验"})
	if err != nil {
		t.Fatalf("CreateClassroom returned error: %v", err)
	}
	if result.ID != 7 {
		t.Fatalf("unexpected classroom: %#v", result)
	}
}

func TestUpdateClassroomUsesClassroomExamEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/classroom-exams/88" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":88,"announce_answer_time":"2026-03-09T08:00:00Z","subjects_rule":{"select_subjects_randomly":true}}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.UpdateClassroom(context.Background(), 88, map[string]any{"title": "更新"})
	if err != nil {
		t.Fatalf("UpdateClassroom returned error: %v", err)
	}
	if result.ID != 88 || result.AnnounceAnswerTime == nil || *result.AnnounceAnswerTime != "2026-03-09T08:00:00Z" {
		t.Fatalf("unexpected classroom: %#v", result)
	}
	if result.SubjectsRule == nil || !result.SubjectsRule.SelectSubjectsRandomly {
		t.Fatalf("subjects_rule did not decode: %#v", result.SubjectsRule)
	}
}

func TestGenerateCoursewareQuizSubjectsUsesFrontendPayload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/courseware-quiz/activity/18/subjects" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		var body GenerateCoursewareQuizSubjectsRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body.UploadReferenceID != 66 || body.NumOfSingleSelection != 3 || body.NumOfMultipleSelection != 2 {
			t.Fatalf("unexpected request body: %#v", body)
		}
		if len(body.BloomCognitiveDomains) != 1 || body.BloomCognitiveDomains[0] != "understand" {
			t.Fatalf("unexpected bloom domains: %#v", body.BloomCognitiveDomains)
		}
		if len(body.PageRange) != 2 || body.PageRange[0] != 4 || body.PageRange[1] != 9 || !body.Stream {
			t.Fatalf("unexpected stream/page_range: %#v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	_, err := service.GenerateCoursewareQuizSubjects(context.Background(), 18, &GenerateCoursewareQuizSubjectsRequest{
		UploadReferenceID:      66,
		NumOfSingleSelection:   3,
		NumOfMultipleSelection: 2,
		NumOfFillInBlank:       1,
		NumOfTrueOrFalse:       0,
		NumOfShortAnswer:       1,
		BloomCognitiveDomains:  []string{"understand"},
		QuizKnowledgePoints:    []any{"kp-1"},
		Locale:                 "",
		Stream:                 true,
		PageRange:              []int{4, 9},
	})
	if err != nil {
		t.Fatalf("GenerateCoursewareQuizSubjects returned error: %v", err)
	}
}

func TestGetCoursewareQuizSettingsDecodesDirectPayload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/courseware-quiz/settings" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"quiz_count_limit":12}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.GetCoursewareQuizSettings(context.Background())
	if err != nil {
		t.Fatalf("GetCoursewareQuizSettings returned error: %v", err)
	}
	if result.QuizCountLimit != 12 {
		t.Fatalf("unexpected settings: %#v", result)
	}
}

func TestSHTVUAndSubjectSaveHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/5/query-shtvu-modules":
			_, _ = w.Write([]byte(`{"chapters":[{"id":1,"name":"Chapter 1"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/5/query-shtvu-subjects":
			if got := r.URL.Query().Get("chapters"); got != "10,11" {
				t.Fatalf("unexpected chapters query: %q", got)
			}
			if got := r.URL.Query().Get("subject_type"); got != "single_selection,multiple_selection" {
				t.Fatalf("unexpected subject_type query: %q", got)
			}
			if got := r.URL.Query().Get("difficulties"); got != "1,2,3" {
				t.Fatalf("unexpected difficulties query: %q", got)
			}
			if got := r.URL.Query().Get("keyword"); got != "integral" {
				t.Fatalf("unexpected keyword query: %q", got)
			}
			if got := r.URL.Query().Get("page_index"); got != "2" {
				t.Fatalf("unexpected page_index query: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "10" {
				t.Fatalf("unexpected page_size query: %q", got)
			}
			_, _ = w.Write([]byte(`{"pages":3,"subjects":[{"id":8,"type":"single_selection","description":"Q1","timestamp":"ts-1"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/5/shtvu-subject-types-info":
			_, _ = w.Write([]byte(`{"subject_types_info":[{"type":"single_selection","subject_count":12}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/exams/6/subjects":
			var body SaveSubjectsRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode exam subjects body: %v", err)
			}
			subjects, ok := body.Subjects.([]any)
			if !ok || len(subjects) != 1 {
				t.Fatalf("unexpected exam subjects body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"subjects":[{"id":1,"type":"single_selection"}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/video-quizzes/7/subjects":
			var body SaveSubjectsRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode video quiz subjects body: %v", err)
			}
			if body.Subjects == nil {
				t.Fatalf("unexpected empty video quiz subjects body")
			}
			_, _ = w.Write([]byte(`{"subjects":[{"id":2,"type":"multiple_selection"}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/exams/6/import-random-subjects-from-shtvu":
			var body ImportRandomSubjectsFromSHTVURequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode exam import body: %v", err)
			}
			if len(body.Items) != 1 || body.Items[0].SubjectType != "single_selection" || body.Items[0].Point != "2" {
				t.Fatalf("unexpected exam import body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"subjects":[{"id":3,"type":"single_selection"}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/classrooms/9/import-random-subjects-from-shtvu":
			_, _ = w.Write([]byte(`{"subjects":[{"id":4,"type":"true_or_false"}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/video-quizzes/7/import-subjects-from-shtvu":
			var body ImportRandomSubjectsFromSHTVURequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode video quiz import body: %v", err)
			}
			if body.Timestamp != "vts-1" {
				t.Fatalf("unexpected video quiz timestamp: %#v", body)
			}
			_, _ = w.Write([]byte(`{"subjects":[{"id":5,"type":"single_selection"}]}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	modules, err := service.GetSHTVUModules(ctx, 5)
	if err != nil || len(modules.Chapters) != 1 || modules.Chapters[0].Name != "Chapter 1" {
		t.Fatalf("unexpected modules: %#v, err=%v", modules, err)
	}

	subjects, err := service.SearchSHTVUSubjects(ctx, 5, &SHTVUSearchSubjectsParams{
		Chapters:     "10,11",
		SubjectType:  "single_selection,multiple_selection",
		Difficulties: "1,2,3",
		Keyword:      "integral",
		PageIndex:    2,
		PageSize:     10,
	})
	if err != nil || subjects.Pages != 3 || len(subjects.Subjects) != 1 {
		t.Fatalf("unexpected shtvu subjects: %#v, err=%v", subjects, err)
	}
	if got, ok := subjects.Subjects[0].Timestamp.(string); !ok || got != "ts-1" {
		t.Fatalf("unexpected subject timestamp: %#v", subjects.Subjects[0].Timestamp)
	}

	info, err := service.GetSHTVUSubjectTypesInfo(ctx, 5)
	if err != nil || len(info.SubjectTypesInfo) != 1 || info.SubjectTypesInfo[0].SubjectCount != 12 {
		t.Fatalf("unexpected subject types info: %#v, err=%v", info, err)
	}

	savedExam, err := service.SaveExamSubjects(ctx, 6, &SaveSubjectsRequest{Subjects: []any{map[string]any{"id": 1}}})
	if err != nil || len(savedExam.Subjects) != 1 || savedExam.Subjects[0].ID != 1 {
		t.Fatalf("unexpected saved exam subjects: %#v, err=%v", savedExam, err)
	}

	savedVideoQuiz, err := service.SaveVideoQuizSubjects(ctx, 7, &SaveSubjectsRequest{Subjects: []any{map[string]any{"id": 2}}})
	if err != nil || len(savedVideoQuiz.Subjects) != 1 || savedVideoQuiz.Subjects[0].ID != 2 {
		t.Fatalf("unexpected saved video quiz subjects: %#v, err=%v", savedVideoQuiz, err)
	}

	importExam, err := service.ImportRandomExamSubjectsFromSHTVU(ctx, 6, &ImportRandomSubjectsFromSHTVURequest{
		Items: []*SHTVURandomImportItem{{SubjectType: "single_selection", Count: 2, Point: "2"}},
	})
	if err != nil || len(importExam.Subjects) != 1 || importExam.Subjects[0].ID != 3 {
		t.Fatalf("unexpected imported exam subjects: %#v, err=%v", importExam, err)
	}

	importClassroom, err := service.ImportRandomClassroomSubjectsFromSHTVU(ctx, 9, &ImportRandomSubjectsFromSHTVURequest{
		Items: []*SHTVURandomImportItem{{SubjectType: "true_or_false", Count: 1}},
	})
	if err != nil || len(importClassroom.Subjects) != 1 || importClassroom.Subjects[0].ID != 4 {
		t.Fatalf("unexpected imported classroom subjects: %#v, err=%v", importClassroom, err)
	}

	importVideoQuiz, err := service.ImportVideoQuizSubjectsFromSHTVU(ctx, 7, &ImportRandomSubjectsFromSHTVURequest{
		Items:     []*SHTVURandomImportItem{{SubjectType: "single_selection", Count: 1}},
		Timestamp: "vts-1",
	})
	if err != nil || len(importVideoQuiz.Subjects) != 1 || importVideoQuiz.Subjects[0].ID != 5 {
		t.Fatalf("unexpected imported video quiz subjects: %#v, err=%v", importVideoQuiz, err)
	}
}
