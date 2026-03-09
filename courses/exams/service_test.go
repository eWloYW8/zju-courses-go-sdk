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
		if body.ModuleID != 18 || body.ModuleType != "subject_lib" {
			t.Fatalf("unexpected module payload: %#v", body)
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
		ModuleID:               18,
		ModuleType:             "subject_lib",
		PageRange:              []int{4, 9},
	})
	if err != nil {
		t.Fatalf("GenerateCoursewareQuizSubjects returned error: %v", err)
	}
}

func TestGenerateSubjectsAndGenerateSubjectsByTextUseFrontendPayloads(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/courseware-quiz/generate-subjects":
			var body GenerateSubjectsRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode generate-subjects body: %v", err)
			}
			if body.UploadID != 77 || body.ModuleID != 9 || body.ModuleType != "exam" || body.GroupID != "" {
				t.Fatalf("unexpected generate-subjects body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"ok":true}`))
		case "/api/courseware-quiz/generate-subjects-by-text":
			var body GenerateSubjectsByTextRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode generate-subjects-by-text body: %v", err)
			}
			if body.TextContent != "chapter summary" || body.ModuleID != 9 || body.ModuleType != "exam" || !body.Stream {
				t.Fatalf("unexpected generate-subjects-by-text body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"ok":true}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if _, err := service.GenerateSubjects(ctx, &GenerateSubjectsRequest{
		UploadID:               77,
		ModuleID:               9,
		ModuleType:             "exam",
		GroupID:                "",
		NumOfSingleSelection:   2,
		NumOfMultipleSelection: 1,
		PageRange:              []int{1, 3},
	}); err != nil {
		t.Fatalf("GenerateSubjects returned error: %v", err)
	}

	if _, err := service.GenerateSubjectsByText(ctx, &GenerateSubjectsByTextRequest{
		TextContent: "chapter summary",
		GenerateCoursewareQuizSubjectsRequest: GenerateCoursewareQuizSubjectsRequest{
			ModuleID:              9,
			ModuleType:            "exam",
			NumOfSingleSelection:  1,
			NumOfTrueOrFalse:      1,
			Stream:                true,
			BloomCognitiveDomains: []string{"apply"},
			QuizKnowledgePoints:   []any{"kp-2"},
		},
	}); err != nil {
		t.Fatalf("GenerateSubjectsByText returned error: %v", err)
	}
}

func TestRubricHelpersUseFrontendFieldsQueries(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/rubrics":
			if got := r.URL.Query().Get("fields"); got != "id,name,conditions" {
				t.Fatalf("unexpected create rubric fields query: %q", got)
			}
			_, _ = w.Write([]byte(`{"id":3,"name":"Rubric","conditions":[{"name":"Criterion","levels":[{"score":5,"description":"good"}]}]}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/rubrics/3":
			if got := r.URL.Query().Get("fields"); got != "id,name,conditions,engage_number,created_by" {
				t.Fatalf("unexpected update rubric fields query: %q", got)
			}
			_, _ = w.Write([]byte(`{"id":3,"name":"Rubric","engage_number":2,"created_by":{"id":9,"name":"Teacher"}}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	created, err := service.CreateRubric(ctx, map[string]any{"name": "Rubric"})
	if err != nil || created.ID != 3 || created.Name == nil || *created.Name != "Rubric" {
		t.Fatalf("unexpected created rubric: %#v, err=%v", created, err)
	}

	updated, err := service.UpdateRubric(ctx, 3, map[string]any{"name": "Rubric"})
	if err != nil || updated.EngageNumber != 2 || updated.CreatedBy == nil || updated.CreatedBy.ID != 9 {
		t.Fatalf("unexpected updated rubric: %#v, err=%v", updated, err)
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

func TestCoursewareQuizHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/courseware-quiz/activity/18/quizzes":
			var body CreateCoursewareQuizRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode create quiz body: %v", err)
			}
			subjects, ok := body.Subjects.([]any)
			if !ok || body.UploadReferenceID != 66 || len(subjects) != 1 {
				t.Fatalf("unexpected create quiz body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"quiz_id":91}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courseware-quiz/quiz/91/my-submission":
			_, _ = w.Write([]byte(`{"submit_times":3,"submission":{"id":7,"submitted_times":1,"non_custom":false,"announce_answer_and_explanation":true}}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/courseware-quiz/quiz/91/submissions":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode submit quiz body: %v", err)
			}
			answers, ok := body["subjects_answers"].([]any)
			if !ok || len(answers) != 1 {
				t.Fatalf("unexpected submit quiz body: %#v", body)
			}
			_, _ = w.Write([]byte(`{}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courseware-quiz/quiz/91/statistic":
			_, _ = w.Write([]byte(`{
				"summary":{"correct_rate":0.5,"students_count":40,"submission_rate":0.8,"submitter_count":32},
				"subjects_statistic":[{"subject_id":3,"correct_count":16,"correct_rate":"50%","wrong_count":16,"wrong_rate":"50%","submitted_count":32,"unsubmitted_count":8,"options_statistic":[{"option_id":1,"sort":1,"chosen_count":18}],"answers_statistic":[{"sort":1,"correct_count":10}]}],
				"analysis":{"last_analyzed_at":"2026-03-09T08:00:00Z","reanalysis":true}
			}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/18/courseware-quiz/quiz/91/subject/3/statistic":
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected page query: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "20" {
				t.Fatalf("unexpected page_size query: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"2025"}` {
				t.Fatalf("unexpected conditions query: %q", got)
			}
			_, _ = w.Write([]byte(`{
				"page":2,
				"page_size":20,
				"pages":3,
				"total":41,
				"items":[{"id":3,"type":"single_selection","description":"Q1","correct_count":16,"correct_rate":"50%","wrong_count":16,"wrong_rate":"50%","submitted_count":32,"unsubmitted_count":8,"options":[{"id":1,"content":"A","chosen_count":18}]}]
			}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/courseware-quiz/quiz/91/analyze":
			_, _ = w.Write([]byte(`{"status":"queued"}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	created, err := service.CreateCoursewareQuiz(ctx, 18, &CreateCoursewareQuizRequest{
		UploadReferenceID: 66,
		Subjects: []any{
			map[string]any{"description": "Q1", "type": "single_selection"},
		},
	})
	if err != nil || created.QuizID != 91 {
		t.Fatalf("unexpected created quiz: %#v, err=%v", created, err)
	}

	submission, err := service.GetCoursewareQuizSubmission(ctx, 91)
	if err != nil || submission.SubmitTimes != 3 || submission.Submission == nil || submission.Submission.SubmittedTimes != 1 || !submission.Submission.AnnounceAnswerAndExplanation {
		t.Fatalf("unexpected quiz submission: %#v, err=%v", submission, err)
	}

	if err := service.SubmitCoursewareQuiz(ctx, 91, map[string]any{
		"subjects_answers": []any{
			map[string]any{"subject_id": 3, "type": "single_selection", "answer_option_ids": []int{1}},
		},
	}); err != nil {
		t.Fatalf("SubmitCoursewareQuiz returned error: %v", err)
	}

	stat, err := service.GetCoursewareQuizStatistic(ctx, 91)
	if err != nil || stat.Summary == nil || stat.Summary.SubmitterCount != 32 || len(stat.SubjectsStatistic) != 1 || stat.AnalysisStatus == nil || !stat.AnalysisStatus.Reanalysis {
		t.Fatalf("unexpected quiz statistic: %#v, err=%v", stat, err)
	}

	subjectStat, err := service.GetCoursewareQuizSubjectStatistic(ctx, 18, 91, 3, CoursewareQuizSubjectStatisticParams{
		Page:       2,
		PageSize:   20,
		Conditions: map[string]any{"keyword": "2025"},
	})
	if err != nil || subjectStat.Page != 2 || len(subjectStat.Items) != 1 || len(subjectStat.Items[0].Options) != 1 || subjectStat.Items[0].Options[0].ChosenCount != 18 {
		t.Fatalf("unexpected subject statistic: %#v, err=%v", subjectStat, err)
	}

	analyzeResult, err := service.AnalyzeCoursewareQuiz(ctx, 91)
	if err != nil || string(analyzeResult) != `{"status":"queued"}` {
		t.Fatalf("unexpected analyze result: %s, err=%v", string(analyzeResult), err)
	}
}

func TestExamPaperZipHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/exams/18/zip-papers":
			_, _ = w.Write([]byte(`{}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/18/zip-status":
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected no-intercept query: %q", got)
			}
			_, _ = w.Write([]byte(`{"paper_zip":{"id":7,"key":"paper-zip","status":"ready","created_at":"2026-03-09T08:00:00Z"}}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if err := service.StartExamPaperZip(ctx, 18); err != nil {
		t.Fatalf("StartExamPaperZip returned error: %v", err)
	}

	status, err := service.GetExamPaperZipStatus(ctx, 18)
	if err != nil {
		t.Fatalf("GetExamPaperZipStatus returned error: %v", err)
	}
	if status.PaperZip == nil || status.PaperZip.ID != 7 || status.PaperZip.Status != "ready" || status.PaperZip.Key != "paper-zip" {
		t.Fatalf("unexpected paper zip status: %#v", status)
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

func TestSubjectGroupHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/exam/6/subject-groups":
			_, _ = w.Write([]byte(`{"data":[{"id":1,"subject_type":"single_selection","sort":1}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/exam/6/subject-group":
			var body SubjectGroupRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode create subject group body: %v", err)
			}
			if body.SubjectType != "multiple_selection" || body.ReferrerType != "exam" || body.ReferrerID != 6 {
				t.Fatalf("unexpected subject group body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"data":{"id":2,"subject_type":"multiple_selection","referrer_type":"exam","referrer_id":6}}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/subject-group/2":
			_, _ = w.Write([]byte(`{"data":{"id":2,"subject_type":"multiple_selection","sort":2}}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	groups, err := service.ListSubjectGroups(ctx, "exam", 6)
	if err != nil || len(groups) != 1 || groups[0].ID != 1 {
		t.Fatalf("unexpected subject groups: %#v, err=%v", groups, err)
	}

	created, err := service.CreateSubjectGroup(ctx, "exam", 6, &SubjectGroupRequest{SubjectType: "multiple_selection", ReferrerType: "exam", ReferrerID: 6})
	if err != nil || created.ID != 2 || created.ReferrerID != 6 {
		t.Fatalf("unexpected created subject group: %#v, err=%v", created, err)
	}

	updated, err := service.UpdateSubjectGroup(ctx, 2, &SubjectGroupRequest{Sort: 2})
	if err != nil || updated.Sort != 2 {
		t.Fatalf("unexpected updated subject group: %#v, err=%v", updated, err)
	}
}

func TestSubjectLibHelpersUseFrontendEndpointsAndModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/subject-libs":
			if got := r.URL.Query().Get("lib_type"); got == "questionnaire" {
				_, _ = w.Write([]byte(`{"subject_libs":[{"id":3,"title":"问卷题库","is_folder":false,"type":"questionnaire"}]}`))
				return
			}
			if got := r.URL.Query().Get("with_folder"); got != "1" {
				t.Fatalf("unexpected with_folder query: %q", got)
			}
			_, _ = w.Write([]byte(`{"subject_libs":[{"id":1,"title":"公共题库","is_folder":true,"is_shared":true,"nums":5,"type":"folder","created_at":"2026-03-09","updated_at":"2026-03-10"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/9/subject-libs":
			if got := r.URL.Query().Get("with_folder"); got != "1" {
				t.Fatalf("unexpected course with_folder query: %q", got)
			}
			_, _ = w.Write([]byte(`{"subject_libs":[{"id":2,"title":"课程题库","is_folder":false,"nums":3}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/subject-libs":
			if got := r.URL.Query().Get("lib_type"); got == "questionnaire" {
				var body CreateSubjectLibRequest
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					t.Fatalf("decode questionnaire body: %v", err)
				}
				if body.Title != "问卷新题库" {
					t.Fatalf("unexpected questionnaire body: %#v", body)
				}
				_, _ = w.Write([]byte(`{"id":4,"title":"问卷新题库","type":"questionnaire"}`))
				return
			}
			var body CreateSubjectLibRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode subject-lib body: %v", err)
			}
			if body.Title != "个人题库" {
				t.Fatalf("unexpected subject-lib body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"id":5,"title":"个人题库","nums":1}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/9/subject-libs":
			if got := r.URL.Query().Get("lib_type"); got != "folder" {
				t.Fatalf("unexpected course lib_type query: %q", got)
			}
			var body CreateSubjectLibRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode course subject-lib body: %v", err)
			}
			if body.Title != "课程文件夹" || body.ParentID != 2 {
				t.Fatalf("unexpected course subject-lib body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"id":6,"title":"课程文件夹","parent_id":2,"is_folder":true}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/subject-libs/6/copy":
			if got := r.URL.Query().Get("questionnaireId"); got != "11" {
				t.Fatalf("unexpected questionnaireId query: %q", got)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPost && r.URL.Path == "/api/subject-libs/batch/copy":
			if got := r.URL.Query().Get("courseware_quiz_id"); got != "15" {
				t.Fatalf("unexpected courseware_quiz_id query: %q", got)
			}
			var body BatchCopySubjectLibsRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode batch copy body: %v", err)
			}
			if len(body.LibIDs) != 2 || body.SubjectID != 8 || body.CourseID != 9 {
				t.Fatalf("unexpected batch copy body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	personal, err := service.ListSubjectLibsWithFolder(ctx, true)
	if err != nil || len(personal.SubjectLibs) != 1 || !personal.SubjectLibs[0].IsShared || personal.SubjectLibs[0].Nums != 5 {
		t.Fatalf("unexpected personal subject libs: %#v, err=%v", personal, err)
	}

	course, err := service.ListCourseSubjectLibs(ctx, 9, true)
	if err != nil || len(course.SubjectLibs) != 1 || course.SubjectLibs[0].Nums != 3 {
		t.Fatalf("unexpected course subject libs: %#v, err=%v", course, err)
	}

	questionnaire, err := service.ListQuestionnaireSubjectLibs(ctx)
	if err != nil || len(questionnaire.SubjectLibs) != 1 || questionnaire.SubjectLibs[0].Type != "questionnaire" {
		t.Fatalf("unexpected questionnaire subject libs: %#v, err=%v", questionnaire, err)
	}

	created, err := service.CreateSubjectLib(ctx, &CreateSubjectLibRequest{Title: "个人题库"})
	if err != nil || created.ID != 5 || created.Nums != 1 {
		t.Fatalf("unexpected created personal subject lib: %#v, err=%v", created, err)
	}

	courseCreated, err := service.CreateCourseSubjectLib(ctx, 9, "folder", &CreateSubjectLibRequest{Title: "课程文件夹", ParentID: 2})
	if err != nil || courseCreated.ID != 6 || !courseCreated.IsFolder {
		t.Fatalf("unexpected created course subject lib: %#v, err=%v", courseCreated, err)
	}

	questionnaireCreated, err := service.CreateQuestionnaireSubjectLib(ctx, &CreateSubjectLibRequest{Title: "问卷新题库"})
	if err != nil || questionnaireCreated.ID != 4 || questionnaireCreated.Type != "questionnaire" {
		t.Fatalf("unexpected created questionnaire subject lib: %#v, err=%v", questionnaireCreated, err)
	}

	if err := service.CopySubjectLibToQuestionnaire(ctx, 11, 6); err != nil {
		t.Fatalf("CopySubjectLibToQuestionnaire returned error: %v", err)
	}

	if err := service.BatchCopySubjectLibsToCoursewareQuiz(ctx, 15, &BatchCopySubjectLibsRequest{
		LibIDs:    []int{6, 7},
		SubjectID: 8,
		CourseID:  9,
	}); err != nil {
		t.Fatalf("BatchCopySubjectLibsToCoursewareQuiz returned error: %v", err)
	}
}

func TestClassroomTypedHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodDelete && r.URL.Path == "/api/classrooms/9":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/classrooms/9/status":
			var body ClassroomStatusRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode classroom status body: %v", err)
			}
			if body.Status != 1 && body.Status != 2 {
				t.Fatalf("unexpected classroom status body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"id":9,"status":"start","subjects_count":3}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/classroom-exams/9":
			_, _ = w.Write([]byte(`{"id":9,"title":"课堂测验","announce_answer_status":"immediate_announce","subjects_rule":{"select_subjects_randomly":true}}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/classroom-exams/9/subjects":
			_, _ = w.Write([]byte(`{"subjects":[{"id":1,"type":"single_selection"}]}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/classrooms/9/subjects/1/status":
			var body UpdateClassroomSubjectStatusRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode classroom subject status body: %v", err)
			}
			if body.Status != 2 {
				t.Fatalf("unexpected classroom subject status body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"id":1,"type":"single_selection","description":"Q1"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/classroom-exams/9/my-submissions":
			_, _ = w.Write([]byte(`{"submissions":[{"id":21,"submitted_times":1}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/classroom-exams/9/submissions/21":
			_, _ = w.Write([]byte(`{"id":21,"submitted_times":1,"score":95}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/classroom-exams/9/score-list":
			if got := r.URL.Query().Get("ignore_avatar"); got != "true" {
				t.Fatalf("unexpected ignore_avatar query: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"examinee_ids":[2,3]}` {
				t.Fatalf("unexpected conditions query: %q", got)
			}
			_, _ = w.Write([]byte(`{"examinees":[{"id":2,"name":"Alice","user_no":"2025001","submitted":true,"score":88}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/classroom-exams/9/examinees":
			_, _ = w.Write([]byte(`{"examinees":[{"id":3,"name":"Bob","department":{"id":4,"name":"CS"}}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/classroom-exams/9/submission-count-status":
			_, _ = w.Write([]byte(`{"submitted_count":8,"unsubmitted_count":2}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if err := service.DeleteClassroom(ctx, 9); err != nil {
		t.Fatalf("DeleteClassroom returned error: %v", err)
	}

	started, err := service.StartClassroom(ctx, 9)
	if err != nil || started.ID != 9 || started.SubjectsCount != 3 {
		t.Fatalf("unexpected started classroom: %#v, err=%v", started, err)
	}

	finished, err := service.FinishClassroom(ctx, 9)
	if err != nil || finished.ID != 9 {
		t.Fatalf("unexpected finished classroom: %#v, err=%v", finished, err)
	}

	classroom, err := service.GetClassroomExamTyped(ctx, 9)
	if err != nil || classroom.ID != 9 || classroom.SubjectsRule == nil || !classroom.SubjectsRule.SelectSubjectsRandomly {
		t.Fatalf("unexpected classroom exam: %#v, err=%v", classroom, err)
	}

	saved, err := service.SaveClassroomSubjectsTyped(ctx, 9, &SaveSubjectsRequest{Subjects: []any{map[string]any{"id": 1}}})
	if err != nil || len(saved.Subjects) != 1 || saved.Subjects[0].ID != 1 {
		t.Fatalf("unexpected saved classroom subjects: %#v, err=%v", saved, err)
	}

	subject, err := service.UpdateClassroomSubjectStatus(ctx, 9, 1, 2)
	if err != nil || subject.ID != 1 {
		t.Fatalf("unexpected classroom subject update: %#v, err=%v", subject, err)
	}

	mySubs, err := service.ListClassroomMySubmissions(ctx, 9)
	if err != nil || len(mySubs.Submissions) != 1 || (*mySubs.Submissions[0])["id"].(float64) != 21 {
		t.Fatalf("unexpected classroom my submissions: %#v, err=%v", mySubs, err)
	}

	submission, err := service.GetClassroomSubmission(ctx, 9, 21)
	if err != nil || (*submission)["id"].(float64) != 21 {
		t.Fatalf("unexpected classroom submission: %#v, err=%v", submission, err)
	}

	scoreList, err := service.ListClassroomScoreList(ctx, 9, &ClassroomScoreListParams{
		IgnoreAvatar: true,
		ExamineeIDs:  []int{2, 3},
	})
	if err != nil || len(scoreList.Examinees) != 1 || scoreList.Examinees[0].Name != "Alice" {
		t.Fatalf("unexpected classroom score list: %#v, err=%v", scoreList, err)
	}

	examinees, err := service.ListClassroomExaminees(ctx, 9)
	if err != nil || len(examinees.Examinees) != 1 || examinees.Examinees[0].Department == nil || examinees.Examinees[0].Department.Name != "CS" {
		t.Fatalf("unexpected classroom examinees: %#v, err=%v", examinees, err)
	}

	countStatus, err := service.GetClassroomSubmissionCountStatus(ctx, 9)
	if err != nil || (*countStatus)["submitted_count"].(float64) != 8 {
		t.Fatalf("unexpected classroom submission count status: %#v, err=%v", countStatus, err)
	}
}

func TestSearchSubjectsInLibOmitsAllSubjectType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/subject-libs/6" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("keyword"); got != "calc" {
			t.Fatalf("unexpected keyword query: %q", got)
		}
		if got := r.URL.Query().Get("subject_type"); got != "" {
			t.Fatalf("subject_type=all should be omitted, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"subjects":[{"id":1,"type":"single_selection"}]}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	subjects, err := service.SearchSubjectsInLib(context.Background(), 6, "calc", "all")
	if err != nil || len(subjects) != 1 || subjects[0].ID != 1 {
		t.Fatalf("unexpected searched subjects: %#v, err=%v", subjects, err)
	}
}

func TestCourseExamAndSyncHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/courses/14/exams":
			_, _ = w.Write([]byte(`{"id":9,"title":"Midterm"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/14/exam-list":
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected exam-list page: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "20" {
				t.Fatalf("unexpected exam-list page_size: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"mid"}` {
				t.Fatalf("unexpected exam-list conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"items":[{"id":9,"title":"Midterm"}],"page":2,"page_size":20,"pages":1,"total":1}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/9/subjects-stat":
			if got := r.URL.Query().Get("exam_paper_type"); got != "practice" {
				t.Fatalf("unexpected exam_paper_type: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"subject_ids":[1,2]}` {
				t.Fatalf("unexpected subjects-stat conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"subjects":[{"id":1}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/9/submissions":
			_, _ = w.Write([]byte(`{"submissions":[{"id":2}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/9/score-list":
			if got := r.URL.Query().Get("conditions"); got != `{"examinee_ids":[3]}` {
				t.Fatalf("unexpected score-list conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"examinees":[{"id":3}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/9/examinees/3":
			_, _ = w.Write([]byte(`{"id":3,"name":"Alice"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/9/subjects-summary":
			if got := r.URL.Query().Get("forAllSubjects"); got == "" {
				_, _ = w.Write([]byte(`{"subjects":[{"id":1}]}`))
				return
			}
			if got := r.URL.Query().Get("forAllSubjects"); got != "false" {
				t.Fatalf("unexpected forAllSubjects: %q", got)
			}
			_, _ = w.Write([]byte(`{"subjects":[{"id":2}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exam/9/examinees":
			if got := r.URL.Query().Get("conditions"); got != `{"group_ids":[4]}` {
				t.Fatalf("unexpected exam examinees conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"examinees":[{"id":3}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exam/9/groups":
			_, _ = w.Write([]byte(`{"groups":[{"id":4}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exam/9/submission-count-status":
			_, _ = w.Write([]byte(`{"submitted_count":8}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/exams/9/subjects/sync-to-platform":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/9/submission/sync-from-platform":
			_, _ = w.Write([]byte(`{"task_id":"task-1"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/9/submission/sync-task-progress":
			_, _ = w.Write([]byte(`{"status":"FINISHED"}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	created, err := service.CreateExam(ctx, 14, map[string]any{"title": "Midterm"})
	if err != nil || created.ID != 9 {
		t.Fatalf("unexpected created exam: %#v, err=%v", created, err)
	}

	examList, err := service.ListCourseExamList(ctx, 14, &ListCourseExamListParams{
		Page:       2,
		PageSize:   20,
		Conditions: map[string]any{"keyword": "mid"},
	})
	if err != nil || len(examList["items"].([]any)) != 1 {
		t.Fatalf("unexpected course exam list: %#v, err=%v", examList, err)
	}

	if stat, err := service.GetExamSubjectsStat(ctx, 9, &ExamSubjectsStatParams{
		ExamPaperType: "practice",
		Conditions:    map[string]any{"subject_ids": []int{1, 2}},
	}); err != nil || len(stat["subjects"].([]any)) != 1 {
		t.Fatalf("unexpected subjects stat: %#v, err=%v", stat, err)
	}

	if submissions, err := service.ListExamSubmissions(ctx, 9); err != nil || len(submissions["submissions"].([]any)) != 1 {
		t.Fatalf("unexpected exam submissions: %#v, err=%v", submissions, err)
	}

	if scoreList, err := service.ListExamScoreList(ctx, 9, &ExamScoreListParams{
		Conditions: map[string]any{"examinee_ids": []int{3}},
	}); err != nil || len(scoreList["examinees"].([]any)) != 1 {
		t.Fatalf("unexpected exam score list: %#v, err=%v", scoreList, err)
	}

	if examinee, err := service.GetExamExaminee(ctx, 9, 3); err != nil || examinee["name"] != "Alice" {
		t.Fatalf("unexpected exam examinee: %#v, err=%v", examinee, err)
	}

	if summary, err := service.GetExamSubjectsSummary(ctx, 9, nil); err != nil || len(summary["subjects"].([]any)) != 1 {
		t.Fatalf("unexpected exam subjects summary: %#v, err=%v", summary, err)
	}

	forAll := false
	if summary, err := service.GetExamSubjectsSummary(ctx, 9, &forAll); err != nil || summary["subjects"].([]any)[0].(map[string]any)["id"] != float64(2) {
		t.Fatalf("unexpected all-subjects summary: %#v, err=%v", summary, err)
	}

	if examinees, err := service.GetExamExaminees(ctx, 9, &ExamExamineesParams{
		Conditions: map[string]any{"group_ids": []int{4}},
	}); err != nil || len(examinees["examinees"].([]any)) != 1 {
		t.Fatalf("unexpected exam examinees: %#v, err=%v", examinees, err)
	}

	if groups, err := service.GetExamGroups(ctx, 9); err != nil || len(groups["groups"].([]any)) != 1 {
		t.Fatalf("unexpected exam groups: %#v, err=%v", groups, err)
	}

	if countStatus, err := service.GetExamSubmissionCountStatus(ctx, 9); err != nil || countStatus["submitted_count"] != float64(8) {
		t.Fatalf("unexpected exam submission count status: %#v, err=%v", countStatus, err)
	}

	if err := service.SyncExamSubjectsToPlatform(ctx, 9); err != nil {
		t.Fatalf("SyncExamSubjectsToPlatform returned error: %v", err)
	}

	if syncResp, err := service.SyncExamSubmissionsFromPlatform(ctx, 9); err != nil || syncResp["task_id"] != "task-1" {
		t.Fatalf("unexpected exam submission sync response: %#v, err=%v", syncResp, err)
	}

	if progress, err := service.GetExamSubmissionSyncTaskProgress(ctx, 9); err != nil || progress["status"] != "FINISHED" {
		t.Fatalf("unexpected exam submission sync progress: %#v, err=%v", progress, err)
	}
}

func TestListSubmittedExamsUsesNoInterceptQuery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/courses/14/submitted-exams" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("no-intercept"); got != "true" {
			t.Fatalf("unexpected no-intercept: %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"exam_ids":[9,10]}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	resp, err := service.ListSubmittedExams(context.Background(), 14)
	if err != nil || len(resp.ExamIDs) != 2 {
		t.Fatalf("unexpected submitted exams response: %#v, err=%v", resp, err)
	}
}

func TestExamGradingHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/12/retake-record":
			if got := r.URL.Query().Get("page"); got != "3" {
				t.Fatalf("unexpected retake page: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "15" {
				t.Fatalf("unexpected retake page_size: %q", got)
			}
			if got := r.URL.Query().Get("include_make_up"); got != "true" {
				t.Fatalf("unexpected include_make_up: %q", got)
			}
			_, _ = w.Write([]byte(`{"page":3,"page_size":15,"pages":4,"total":42,"items":[{"id":8,"exam_paper_type":"makeup_exam","status":"finished"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/12/subjective-questions":
			_, _ = w.Write([]byte(`{"subjects":[{"id":5,"type":"short_answer","description":"Q1"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/subjects/5":
			if got := r.URL.Query().Get("can_select_sub_subject"); got != "true" {
				t.Fatalf("unexpected can_select_sub_subject: %q", got)
			}
			if got := r.URL.Query().Get("exam_id"); got != "12" {
				t.Fatalf("unexpected exam_id: %q", got)
			}
			_, _ = w.Write([]byte(`{"id":5,"type":"single_selection","description":"Q1","sub_subjects":[{"id":6,"type":"short_answer"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/12/subjects":
			if got := r.URL.Query().Get("keyword"); got != "matrix" {
				t.Fatalf("unexpected subject keyword: %q", got)
			}
			if got := r.URL.Query().Get("is_makeup_exam"); got != "true" {
				t.Fatalf("unexpected is_makeup_exam: %q", got)
			}
			if got := r.URL.Query().Get("subject_type"); got != "short_answer" {
				t.Fatalf("unexpected subject_type: %q", got)
			}
			_, _ = w.Write([]byte(`{"subjects":[{"id":7,"type":"short_answer","description":"Q2"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/12/subjects/5/examinees":
			if got := r.URL.Query().Get("is_makeup_exam"); got != "true" {
				t.Fatalf("unexpected examinees is_makeup_exam: %q", got)
			}
			_, _ = w.Write([]byte(`{"examinees":[{"id":3,"name":"Alice"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/12/subjects/5/groups":
			if got := r.URL.Query().Get("is_makeup_exam"); got != "false" {
				t.Fatalf("unexpected groups is_makeup_exam: %q", got)
			}
			_, _ = w.Write([]byte(`{"groups":[{"id":4,"name":"G1"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/12/subjects/5/examinees/3/submissions":
			if got := r.URL.Query().Get("is_makeup_exam"); got != "true" {
				t.Fatalf("unexpected examinee submissions is_makeup_exam: %q", got)
			}
			_, _ = w.Write([]byte(`{"submissions":[{"id":21}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/exams/12/subjects/5/groups/4/submissions":
			if got := r.URL.Query().Get("is_makeup_exam"); got != "false" {
				t.Fatalf("unexpected group submissions is_makeup_exam: %q", got)
			}
			_, _ = w.Write([]byte(`{"submissions":[{"id":22}]}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/exams/12/submissions/21/comment":
			var body UpdateExamSubmissionCommentRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode submission comment body: %v", err)
			}
			if body.Comment != "well done" {
				t.Fatalf("unexpected submission comment body: %#v", body)
			}
			_, _ = w.Write([]byte(`{}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/exams/12/give-score":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode give-score body: %v", err)
			}
			if body["score"] != float64(95) {
				t.Fatalf("unexpected give-score body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"message":"ok"}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/exams/12/give-scores":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode give-scores body: %v", err)
			}
			scores, ok := body["scores"].([]any)
			if !ok || len(scores) != 2 {
				t.Fatalf("unexpected give-scores body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"message":"batch-ok"}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	retakes, err := service.ListExamRetakeRecords(ctx, 12, &ListExamRetakeRecordsParams{Page: 3, PageSize: 15})
	if err != nil || retakes.Page != 3 || len(retakes.Items) != 1 || (*retakes.Items[0])["exam_paper_type"] != "makeup_exam" {
		t.Fatalf("unexpected retake records: %#v, err=%v", retakes, err)
	}

	subjective, err := service.ListExamSubjectiveQuestions(ctx, 12)
	if err != nil || len(subjective.Subjects) != 1 || subjective.Subjects[0].Type != "short_answer" {
		t.Fatalf("unexpected subjective questions: %#v, err=%v", subjective, err)
	}

	subject, err := service.GetSubjectForExam(ctx, 5, 12)
	if err != nil || subject.ID != 5 || len(subject.SubSubjects) != 1 {
		t.Fatalf("unexpected exam subject: %#v, err=%v", subject, err)
	}

	subjects, err := service.SearchExamSubjects(ctx, 12, &SearchExamSubjectsParams{
		IsMakeupExam: true,
		Keyword:      "matrix",
		SubjectType:  "short_answer",
	})
	if err != nil || len(subjects) != 1 || subjects[0].ID != 7 {
		t.Fatalf("unexpected exam subjects search: %#v, err=%v", subjects, err)
	}

	if examinees, err := service.ListExamSubjectExaminees(ctx, 12, 5, true); err != nil || len(examinees["examinees"].([]any)) != 1 {
		t.Fatalf("unexpected subject examinees: %#v, err=%v", examinees, err)
	}

	if groups, err := service.ListExamSubjectGroups(ctx, 12, 5, false); err != nil || len(groups["groups"].([]any)) != 1 {
		t.Fatalf("unexpected subject groups: %#v, err=%v", groups, err)
	}

	if submissions, err := service.ListExamSubjectExamineeSubmissions(ctx, 12, 5, 3, true); err != nil || len(submissions["submissions"].([]any)) != 1 {
		t.Fatalf("unexpected examinee submissions: %#v, err=%v", submissions, err)
	}

	if submissions, err := service.ListExamSubjectGroupSubmissions(ctx, 12, 5, 4, false); err != nil || len(submissions["submissions"].([]any)) != 1 {
		t.Fatalf("unexpected group submissions: %#v, err=%v", submissions, err)
	}

	if err := service.UpdateExamSubmissionComment(ctx, 12, 21, &UpdateExamSubmissionCommentRequest{Comment: "well done"}); err != nil {
		t.Fatalf("UpdateExamSubmissionComment returned error: %v", err)
	}

	if result, err := service.GiveExamScore(ctx, 12, map[string]any{"score": 95}); err != nil || result["message"] != "ok" {
		t.Fatalf("unexpected give-score response: %#v, err=%v", result, err)
	}

	if result, err := service.GiveExamScores(ctx, 12, map[string]any{"scores": []map[string]any{{"user_id": 3, "score": 95}, {"user_id": 4, "score": 88}}}); err != nil || result["message"] != "batch-ok" {
		t.Fatalf("unexpected give-scores response: %#v, err=%v", result, err)
	}
}

func TestExamPointRuleAndCampusImportHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPut && r.URL.Path == "/api/exams/12/points-and-rules":
			var body UpdateExamPointsAndRulesRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode exam points body: %v", err)
			}
			if len(body.SubjectsPointsAndRules) != 1 || body.SubjectsPointsAndRules[0].SubjectIndex != 0 {
				t.Fatalf("unexpected exam points body: %#v", body)
			}
			if len(body.SubjectsPointsAndRules[0].PointRules) != 1 || body.SubjectsPointsAndRules[0].PointRules[0].RuleDifficultyLevel != "easy" {
				t.Fatalf("unexpected exam point rules: %#v", body)
			}
			_, _ = w.Write([]byte(`{"message":"saved","select_subjects_randomly_rule":{"0":{"single_selection":{"subjects_count":"2","subject_point":"2.0"}}}}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/classroom-exams/9/points":
			var body UpdateExamPointsAndRulesRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode classroom points body: %v", err)
			}
			if len(body.SubjectsPointsAndRules) != 1 || body.SubjectsPointsAndRules[0].PointRules[0].RuleName != "single_selection" {
				t.Fatalf("unexpected classroom points body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"message":"classroom-saved"}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/exams/12/imported-subjects-from-campus":
			var body ImportSubjectsFromCampusRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode exam campus import body: %v", err)
			}
			if len(body.Items) != 1 || body.Items[0].ID != 5 || body.Items[0].Count != 2 || body.Items[0].Point != "2.0" {
				t.Fatalf("unexpected exam campus import items: %#v", body)
			}
			if len(body.ExamSubjectTypes) != 1 || body.ExamSubjectTypes[0] != "single_selection" {
				t.Fatalf("unexpected exam subject types: %#v", body)
			}
			if len(body.ExamSubjectDifficultyLevels) != 1 || body.ExamSubjectDifficultyLevels[0] != "easy" {
				t.Fatalf("unexpected exam difficulty levels: %#v", body)
			}
			_, _ = w.Write([]byte(`{"subjects":[{"id":31,"type":"single_selection"}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/classrooms/9/imported-subjects-from-campus":
			var body ImportSubjectsFromCampusRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode classroom campus import body: %v", err)
			}
			if body.Settings == nil {
				t.Fatalf("unexpected empty classroom campus import settings")
			}
			_, _ = w.Write([]byte(`{"subjects":[{"id":32,"type":"multiple_selection"}]}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	examRule, err := service.UpdateExamPointsAndRules(ctx, 12, &UpdateExamPointsAndRulesRequest{
		SubjectsPointsAndRules: []*ExamSubjectPointsAndRules{
			{
				SubjectIndex: 0,
				PointRules: []*ExamPointRule{
					{RuleName: "single_selection", RulePoint: "2.0", RuleNumber: "2", RuleDifficultyLevel: "easy"},
				},
			},
		},
	})
	if err != nil || examRule.Message != "saved" || examRule.SelectSubjectsRandomlyRule == nil {
		t.Fatalf("unexpected exam point rule response: %#v, err=%v", examRule, err)
	}

	classroomRule, err := service.UpdateClassroomPoints(ctx, 9, &UpdateExamPointsAndRulesRequest{
		SubjectsPointsAndRules: []*ExamSubjectPointsAndRules{
			{
				SubjectIndex: 0,
				PointRules: []*ExamPointRule{
					{RuleName: "single_selection", RulePoint: "3.0", RuleNumber: "1"},
				},
			},
		},
	})
	if err != nil || classroomRule.Message != "classroom-saved" {
		t.Fatalf("unexpected classroom point rule response: %#v, err=%v", classroomRule, err)
	}

	examImport, err := service.ImportExamSubjectsFromCampus(ctx, 12, &ImportSubjectsFromCampusRequest{
		Items: []*CampusSubjectSelection{
			{ID: 5, Count: 2, Point: "2.0"},
		},
		Settings:                    map[string]any{"play_time_limit": 0},
		ExamSubjectTypes:            []string{"single_selection"},
		ExamSubjectDifficultyLevels: []string{"easy"},
	})
	if err != nil || len(examImport.Subjects) != 1 || examImport.Subjects[0].ID != 31 {
		t.Fatalf("unexpected exam campus import response: %#v, err=%v", examImport, err)
	}

	classroomImport, err := service.ImportClassroomSubjectsFromCampus(ctx, 9, &ImportSubjectsFromCampusRequest{
		Items: []*CampusSubjectSelection{
			{ID: 6, Count: 1, Point: "3.0"},
		},
		Settings:                    map[string]any{"play_time_limit": 10},
		ExamSubjectTypes:            []string{"multiple_selection"},
		ExamSubjectDifficultyLevels: []string{"medium"},
	})
	if err != nil || len(classroomImport.Subjects) != 1 || classroomImport.Subjects[0].ID != 32 {
		t.Fatalf("unexpected classroom campus import response: %#v, err=%v", classroomImport, err)
	}
}
