package exams

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// --- Response Types ---

// SubjectLibsResponse represents the response from ListSubjectLibs.
type SubjectLibsResponse struct {
	SubjectLibs []json.RawMessage `json:"subject_libs"`
}

// RubricsResponse represents the response from ListRubrics.
type RubricsResponse struct {
	Rubrics []*model.Rubric `json:"rubrics"`
}

// Service handles exam-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Exam CRUD ---

// GetExam returns detailed information about an exam.
func (s *Service) GetExam(ctx context.Context, examID int) (*model.Exam, error) {
	u := fmt.Sprintf("/api/exams/%d", examID)
	result := new(model.Exam)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateExam creates a new exam.
func (s *Service) CreateExam(ctx context.Context, courseID int, exam interface{}) (*model.Exam, error) {
	u := fmt.Sprintf("/api/exams/%d", courseID)
	result := new(model.Exam)
	_, err := s.client.Post(ctx, u, exam, result)
	return result, err
}

// UpdateExam updates an exam.
func (s *Service) UpdateExam(ctx context.Context, examID int, exam interface{}) (*model.Exam, error) {
	u := fmt.Sprintf("/api/exams/%d", examID)
	result := new(model.Exam)
	_, err := s.client.Put(ctx, u, exam, result)
	return result, err
}

// DeleteExam deletes an exam.
func (s *Service) DeleteExam(ctx context.Context, examID int) error {
	u := fmt.Sprintf("/api/exams/%d", examID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// BatchDeleteExams deletes multiple exams at once.
func (s *Service) BatchDeleteExams(ctx context.Context, examIDs []int) error {
	_, err := s.client.Post(ctx, "/api/exams/batch_delete", map[string][]int{"exam_ids": examIDs}, nil)
	return err
}

// --- Exam Submissions ---

// SubmitExam submits an exam.
func (s *Service) SubmitExam(ctx context.Context, examID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/exams/submissions/%d", examID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// GetExamSubmission returns an exam submission.
func (s *Service) GetExamSubmission(ctx context.Context, submissionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/exams/submissions/%d", submissionID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Exam Scores ---

// GetExamScore returns the score for an exam.
func (s *Service) GetExamScore(ctx context.Context, examID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/exam-scores/%d", examID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// UpdateExamScore updates an exam score (instructor).
func (s *Service) UpdateExamScore(ctx context.Context, scoreID int, body interface{}) error {
	u := fmt.Sprintf("/api/exam-scores/%d", scoreID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// --- Classroom / In-class Quiz ---

// GetClassroom returns a classroom quiz.
func (s *Service) GetClassroom(ctx context.Context, classroomID int) (*model.Classroom, error) {
	u := fmt.Sprintf("/api/classrooms/%d", classroomID)
	result := new(model.Classroom)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateClassroom creates a new classroom quiz.
func (s *Service) CreateClassroom(ctx context.Context, courseID int, classroom interface{}) (*model.Classroom, error) {
	u := fmt.Sprintf("/api/classroom-exams/%d", courseID)
	result := new(model.Classroom)
	_, err := s.client.Post(ctx, u, classroom, result)
	return result, err
}

// UpdateClassroom updates a classroom quiz.
func (s *Service) UpdateClassroom(ctx context.Context, classroomID int, classroom interface{}) (*model.Classroom, error) {
	u := fmt.Sprintf("/api/classrooms/%d", classroomID)
	result := new(model.Classroom)
	_, err := s.client.Put(ctx, u, classroom, result)
	return result, err
}

// SubmitClassroomExam submits a classroom exam.
func (s *Service) SubmitClassroomExam(ctx context.Context, classroomID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/classroom-exams/%d", classroomID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// GetClassroomExam returns the classroom exam config/detail by activity id.
func (s *Service) GetClassroomExam(ctx context.Context, classroomID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/classroom-exams/%d", classroomID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// SaveClassroomSubjects saves classroom exam subjects.
func (s *Service) SaveClassroomSubjects(ctx context.Context, classroomID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/classroom-exams/%d/subjects", classroomID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// SortClassroomExamSubjects sorts classroom exam subjects.
func (s *Service) SortClassroomExamSubjects(ctx context.Context, classroomID int, body interface{}) error {
	u := fmt.Sprintf("/api/classroom-exams/%d/subject-sort", classroomID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// CreateExamFinalScore creates a final score entry.
func (s *Service) CreateExamFinalScore(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/exam-scores", body, &result)
	return result, err
}

// UpdateExamFinalScore updates a final score entry.
func (s *Service) UpdateExamFinalScore(ctx context.Context, scoreID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/exam-scores/%d", scoreID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// SortExamSubjects sorts exam subjects.
func (s *Service) SortExamSubjects(ctx context.Context, examID int, body interface{}) error {
	u := fmt.Sprintf("/api/exams/%d/subject-sort", examID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// BatchDeleteExamSubjects deletes exam subjects in batch.
func (s *Service) BatchDeleteExamSubjects(ctx context.Context, examID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/exams/%d/subjects/batch_delete", examID)
	var result json.RawMessage
	_, err := s.client.DeleteWithBody(ctx, u, body, &result)
	return result, err
}

// BatchDeleteClassroomSubjects deletes classroom exam subjects in batch.
func (s *Service) BatchDeleteClassroomSubjects(ctx context.Context, classroomID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/classroom-exams/%d/subjects/batch_delete", classroomID)
	var result json.RawMessage
	_, err := s.client.DeleteWithBody(ctx, u, body, &result)
	return result, err
}

// --- Courseware Quiz ---

// ListCoursewareQuizzes returns quizzes for a courseware activity.
func (s *Service) ListCoursewareQuizzes(ctx context.Context, activityID int) ([]interface{}, error) {
	u := fmt.Sprintf("/api/courseware-quiz/activity/%d/quizzes", activityID)
	var result []interface{}
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCoursewareQuiz returns a specific courseware quiz.
func (s *Service) GetCoursewareQuiz(ctx context.Context, quizID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courseware-quiz/quiz/%d", quizID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GenerateCoursewareQuizSubjects generates subjects for a courseware activity.
func (s *Service) GenerateCoursewareQuizSubjects(ctx context.Context, activityID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courseware-quiz/activity/%d/subjects", activityID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// FormatQuestion formats a question using AI.
func (s *Service) FormatQuestion(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/courseware-quiz/format-question", body, &result)
	return result, err
}

// GenerateSubjects generates subjects using AI.
func (s *Service) GenerateSubjects(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/courseware-quiz/generate-subjects", body, &result)
	return result, err
}

// GenerateSubjectsByText generates subjects from text.
func (s *Service) GenerateSubjectsByText(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/courseware-quiz/generate-subjects-by-text", body, &result)
	return result, err
}

// GetCoursewareQuizSettings returns courseware quiz settings.
func (s *Service) GetCoursewareQuizSettings(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/courseware-quiz/settings", &result)
	return result, err
}

// --- Subject Libraries ---

// ListSubjectLibs returns subject libraries.
func (s *Service) ListSubjectLibs(ctx context.Context, params map[string]string) (*SubjectLibsResponse, error) {
	u := addQueryParams("/api/subject-libs", params)
	result := new(SubjectLibsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetSubjectLib returns a specific subject library.
func (s *Service) GetSubjectLib(ctx context.Context, libID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/subject-libs/%d", libID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateSubjectLib creates a new subject library.
func (s *Service) CreateSubjectLib(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/subject-libs", body, &result)
	return result, err
}

// GetSubject returns a specific subject/question.
func (s *Service) GetSubject(ctx context.Context, subjectID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/subjects/%d", subjectID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// RandomSubjects gets random subjects from a library.
func (s *Service) RandomSubjects(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/subject-libs/random", body, &result)
	return result, err
}

// ListSubjectLibFolders returns subject library folders.
func (s *Service) ListSubjectLibFolders(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/subject-libs/folders", params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CopySubjectLibsToUser copies subject libraries to the current user.
func (s *Service) CopySubjectLibsToUser(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/subject-libs/copy-to-user", body, &result)
	return result, err
}

// MoveSubjectLibs moves subject libraries in the folder tree.
func (s *Service) MoveSubjectLibs(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/subject-libs/libs-move-to", body, &result)
	return result, err
}

// BatchCopySubjectLibs copies subject libraries for a courseware quiz context.
func (s *Service) BatchCopySubjectLibs(ctx context.Context, coursewareQuizID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/subject-libs/batch/copy?courseware_quiz_id=%d", coursewareQuizID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// --- Rubrics ---

// ListRubrics returns rubrics.
func (s *Service) ListRubrics(ctx context.Context, params map[string]string) (*RubricsResponse, error) {
	u := addQueryParams("/api/rubrics", params)
	result := new(RubricsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetRubricTemplate returns the default rubric template.
func (s *Service) GetRubricTemplate(ctx context.Context) (*model.RubricInstance, error) {
	result := new(model.RubricInstance)
	_, err := s.client.Get(ctx, "/api/rubrics/template?no-intercept=true", result)
	return result, err
}

// CreateRubric creates a new rubric.
func (s *Service) CreateRubric(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/rubrics", body, &result)
	return result, err
}

// UpdateRubric updates a rubric.
func (s *Service) UpdateRubric(ctx context.Context, rubricID int, body interface{}) error {
	u := fmt.Sprintf("/api/rubrics/%d", rubricID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// DeleteRubric deletes a rubric.
func (s *Service) DeleteRubric(ctx context.Context, rubricID int) error {
	u := fmt.Sprintf("/api/rubrics/%d", rubricID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// GenerateRubric generates a rubric using AI.
func (s *Service) GenerateRubric(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/rubrics/generate", body, &result)
	return result, err
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}
