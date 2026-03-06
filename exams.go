package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

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

// ExamsService handles exam-related API operations.
type ExamsService struct {
	client *Client
}

// --- Exam CRUD ---

// GetExam returns detailed information about an exam.
func (s *ExamsService) GetExam(ctx context.Context, examID int) (*model.Exam, error) {
	u := fmt.Sprintf("/api/exams/%d", examID)
	result := new(model.Exam)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// CreateExam creates a new exam.
func (s *ExamsService) CreateExam(ctx context.Context, courseID int, exam interface{}) (*model.Exam, error) {
	u := fmt.Sprintf("/api/exams/%d", courseID)
	result := new(model.Exam)
	_, err := s.client.post(ctx, u, exam, result)
	return result, err
}

// UpdateExam updates an exam.
func (s *ExamsService) UpdateExam(ctx context.Context, examID int, exam interface{}) (*model.Exam, error) {
	u := fmt.Sprintf("/api/exams/%d", examID)
	result := new(model.Exam)
	_, err := s.client.put(ctx, u, exam, result)
	return result, err
}

// DeleteExam deletes an exam.
func (s *ExamsService) DeleteExam(ctx context.Context, examID int) error {
	u := fmt.Sprintf("/api/exams/%d", examID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// BatchDeleteExams deletes multiple exams at once.
func (s *ExamsService) BatchDeleteExams(ctx context.Context, examIDs []int) error {
	_, err := s.client.post(ctx, "/api/exams/batch_delete", map[string][]int{"exam_ids": examIDs}, nil)
	return err
}

// --- Exam Submissions ---

// SubmitExam submits an exam.
func (s *ExamsService) SubmitExam(ctx context.Context, examID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/exams/submissions/%d", examID)
	var result json.RawMessage
	_, err := s.client.post(ctx, u, body, &result)
	return result, err
}

// GetExamSubmission returns an exam submission.
func (s *ExamsService) GetExamSubmission(ctx context.Context, submissionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/exams/submissions/%d", submissionID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Exam Scores ---

// GetExamScore returns the score for an exam.
func (s *ExamsService) GetExamScore(ctx context.Context, examID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/exam-scores/%d", examID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// UpdateExamScore updates an exam score (instructor).
func (s *ExamsService) UpdateExamScore(ctx context.Context, scoreID int, body interface{}) error {
	u := fmt.Sprintf("/api/exam-scores/%d", scoreID)
	_, err := s.client.put(ctx, u, body, nil)
	return err
}

// --- Classroom / In-class Quiz ---

// GetClassroom returns a classroom quiz.
func (s *ExamsService) GetClassroom(ctx context.Context, classroomID int) (*model.Classroom, error) {
	u := fmt.Sprintf("/api/classrooms/%d", classroomID)
	result := new(model.Classroom)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// CreateClassroom creates a new classroom quiz.
func (s *ExamsService) CreateClassroom(ctx context.Context, courseID int, classroom interface{}) (*model.Classroom, error) {
	u := fmt.Sprintf("/api/classroom-exams/%d", courseID)
	result := new(model.Classroom)
	_, err := s.client.post(ctx, u, classroom, result)
	return result, err
}

// UpdateClassroom updates a classroom quiz.
func (s *ExamsService) UpdateClassroom(ctx context.Context, classroomID int, classroom interface{}) (*model.Classroom, error) {
	u := fmt.Sprintf("/api/classrooms/%d", classroomID)
	result := new(model.Classroom)
	_, err := s.client.put(ctx, u, classroom, result)
	return result, err
}

// SubmitClassroomExam submits a classroom exam.
func (s *ExamsService) SubmitClassroomExam(ctx context.Context, classroomID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/classroom-exams/%d", classroomID)
	var result json.RawMessage
	_, err := s.client.post(ctx, u, body, &result)
	return result, err
}

// --- Courseware Quiz ---

// ListCoursewareQuizzes returns quizzes for a courseware activity.
func (s *ExamsService) ListCoursewareQuizzes(ctx context.Context, activityID int) ([]interface{}, error) {
	u := fmt.Sprintf("/api/courseware-quiz/activity/%d/quizzes", activityID)
	var result []interface{}
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetCoursewareQuiz returns a specific courseware quiz.
func (s *ExamsService) GetCoursewareQuiz(ctx context.Context, quizID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courseware-quiz/quiz/%d", quizID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// FormatQuestion formats a question using AI.
func (s *ExamsService) FormatQuestion(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/courseware-quiz/format-question", body, &result)
	return result, err
}

// GenerateSubjects generates subjects using AI.
func (s *ExamsService) GenerateSubjects(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/courseware-quiz/generate-subjects", body, &result)
	return result, err
}

// GenerateSubjectsByText generates subjects from text.
func (s *ExamsService) GenerateSubjectsByText(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/courseware-quiz/generate-subjects-by-text", body, &result)
	return result, err
}

// GetCoursewareQuizSettings returns courseware quiz settings.
func (s *ExamsService) GetCoursewareQuizSettings(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/courseware-quiz/settings", &result)
	return result, err
}

// --- Subject Libraries ---

// ListSubjectLibs returns subject libraries.
func (s *ExamsService) ListSubjectLibs(ctx context.Context, params map[string]string) (*SubjectLibsResponse, error) {
	u := addQueryParams("/api/subject-libs", params)
	result := new(SubjectLibsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetSubjectLib returns a specific subject library.
func (s *ExamsService) GetSubjectLib(ctx context.Context, libID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/subject-libs/%d", libID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// CreateSubjectLib creates a new subject library.
func (s *ExamsService) CreateSubjectLib(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/subject-libs", body, &result)
	return result, err
}

// GetSubject returns a specific subject/question.
func (s *ExamsService) GetSubject(ctx context.Context, subjectID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/subjects/%d", subjectID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// RandomSubjects gets random subjects from a library.
func (s *ExamsService) RandomSubjects(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/subject-libs/random", body, &result)
	return result, err
}

// --- Rubrics ---

// ListRubrics returns rubrics.
func (s *ExamsService) ListRubrics(ctx context.Context, params map[string]string) (*RubricsResponse, error) {
	u := addQueryParams("/api/rubrics", params)
	result := new(RubricsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetRubricTemplate returns the default rubric template.
func (s *ExamsService) GetRubricTemplate(ctx context.Context) (*model.RubricInstance, error) {
	result := new(model.RubricInstance)
	_, err := s.client.get(ctx, "/api/rubrics/template?no-intercept=true", result)
	return result, err
}

// CreateRubric creates a new rubric.
func (s *ExamsService) CreateRubric(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/rubrics", body, &result)
	return result, err
}

// UpdateRubric updates a rubric.
func (s *ExamsService) UpdateRubric(ctx context.Context, rubricID int, body interface{}) error {
	u := fmt.Sprintf("/api/rubrics/%d", rubricID)
	_, err := s.client.put(ctx, u, body, nil)
	return err
}

// DeleteRubric deletes a rubric.
func (s *ExamsService) DeleteRubric(ctx context.Context, rubricID int) error {
	u := fmt.Sprintf("/api/rubrics/%d", rubricID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// GenerateRubric generates a rubric using AI.
func (s *ExamsService) GenerateRubric(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/rubrics/generate", body, &result)
	return result, err
}
