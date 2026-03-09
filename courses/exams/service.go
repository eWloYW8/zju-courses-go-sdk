package exams

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

// Service handles exam-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Exam CRUD ---

// GetExam returns detailed information about an exam.
func (s *Service) GetExam(ctx context.Context, examID int) (*Exam, error) {
	u := fmt.Sprintf("/api/exams/%d", examID)
	result := new(Exam)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetExamMakeUpRecord returns the make-up record for an exam.
func (s *Service) GetExamMakeUpRecord(ctx context.Context, examID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/exam/%d/make-up-record", examID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CheckExamQualification checks whether the current user can access an exam.
func (s *Service) CheckExamQualification(ctx context.Context, examID int, checkStatus string) (json.RawMessage, error) {
	params := map[string]string{}
	if checkStatus != "" {
		params["check_status"] = checkStatus
	}
	u := addQueryParams(fmt.Sprintf("/api/exam/%d/check-exam-qualification?no-intercept=true", examID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateExam creates a new exam.
func (s *Service) CreateExam(ctx context.Context, courseID int, exam interface{}) (*Exam, error) {
	u := fmt.Sprintf("/api/exams/%d", courseID)
	result := new(Exam)
	_, err := s.client.Post(ctx, u, exam, result)
	return result, err
}

// UpdateExam updates an exam.
func (s *Service) UpdateExam(ctx context.Context, examID int, exam interface{}) (*Exam, error) {
	u := fmt.Sprintf("/api/exams/%d", examID)
	result := new(Exam)
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

// UpdateExamSubjectExplanation updates the explanation text for a specific exam subject.
func (s *Service) UpdateExamSubjectExplanation(ctx context.Context, examID int, subjectID int, body *UpdateExamSubjectExplanationRequest) (*ExamSubject, error) {
	u := fmt.Sprintf("/api/exams/%d/subjects/%d/explanation", examID, subjectID)
	result := new(ExamSubject)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// GetExamScoreDistribution returns the score-distribution payload used by exam grading views.
func (s *Service) GetExamScoreDistribution(ctx context.Context, examID int, conditions ExamScoreDistributionConditions) (ExamScoreDistributionResponse, error) {
	u := fmt.Sprintf("/api/exams/%d/score-distribution", examID)
	if conditions != nil {
		body, err := json.Marshal(conditions)
		if err != nil {
			return nil, err
		}
		u += "?conditions=" + url.QueryEscape(string(body))
	}
	result := make(ExamScoreDistributionResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Classroom / In-class Quiz ---

// GetClassroom returns a classroom quiz.
func (s *Service) GetClassroom(ctx context.Context, classroomID int) (*Classroom, error) {
	u := fmt.Sprintf("/api/classrooms/%d", classroomID)
	result := new(Classroom)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateClassroom creates a new classroom quiz.
func (s *Service) CreateClassroom(ctx context.Context, courseID int, classroom interface{}) (*Classroom, error) {
	u := fmt.Sprintf("/api/courses/%d/classroom-exams", courseID)
	result := new(Classroom)
	_, err := s.client.Post(ctx, u, classroom, result)
	return result, err
}

// UpdateClassroom updates a classroom quiz.
func (s *Service) UpdateClassroom(ctx context.Context, classroomID int, classroom interface{}) (*Classroom, error) {
	u := fmt.Sprintf("/api/classroom-exams/%d", classroomID)
	result := new(Classroom)
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

// SaveExamSubjects saves exam subjects.
func (s *Service) SaveExamSubjects(ctx context.Context, examID int, body *SaveSubjectsRequest) (*SubjectsResponse, error) {
	u := fmt.Sprintf("/api/exams/%d/subjects", examID)
	result := new(SubjectsResponse)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// SaveVideoQuizSubjects saves subjects for a video quiz.
func (s *Service) SaveVideoQuizSubjects(ctx context.Context, videoQuizID int, body *SaveSubjectsRequest) (*SubjectsResponse, error) {
	u := fmt.Sprintf("/api/video-quizzes/%d/subjects", videoQuizID)
	result := new(SubjectsResponse)
	_, err := s.client.Post(ctx, u, body, result)
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
func (s *Service) ListCoursewareQuizzes(ctx context.Context, activityID int) ([]*CoursewareQuiz, error) {
	u := fmt.Sprintf("/api/courseware-quiz/activity/%d/quizzes", activityID)
	var result []*CoursewareQuiz
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

// GetCoursewareQuizSubjects returns subjects for a specific courseware quiz.
func (s *Service) GetCoursewareQuizSubjects(ctx context.Context, quizID int) (*CoursewareQuizSubjectsResponse, error) {
	u := fmt.Sprintf("/api/courseware-quiz/quiz/%d/subjects", quizID)
	result := new(CoursewareQuizSubjectsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateCoursewareQuizSubjects updates subjects for a specific courseware quiz.
func (s *Service) UpdateCoursewareQuizSubjects(ctx context.Context, quizID int, body interface{}) (*CoursewareQuizUpdateResponse, error) {
	u := fmt.Sprintf("/api/courseware-quiz/quiz/%d/subjects", quizID)
	result := new(CoursewareQuizUpdateResponse)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// GenerateCoursewareQuizSubjects generates subjects for a courseware activity.
func (s *Service) GenerateCoursewareQuizSubjects(ctx context.Context, activityID int, body *GenerateCoursewareQuizSubjectsRequest) (json.RawMessage, error) {
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

// GenerateSubjects generates subjects from an uploaded file using AI.
func (s *Service) GenerateSubjects(ctx context.Context, body *GenerateSubjectsRequest) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/courseware-quiz/generate-subjects", body, &result)
	return result, err
}

// GenerateSubjectsByText generates subjects from text.
func (s *Service) GenerateSubjectsByText(ctx context.Context, body *GenerateSubjectsByTextRequest) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/courseware-quiz/generate-subjects-by-text", body, &result)
	return result, err
}

// GetCoursewareQuizSettings returns courseware quiz settings.
func (s *Service) GetCoursewareQuizSettings(ctx context.Context) (*CoursewareQuizSettingsResponse, error) {
	result := new(CoursewareQuizSettingsResponse)
	_, err := s.client.Get(ctx, "/api/courseware-quiz/settings", result)
	if err == nil && result.Setting != nil && result.QuizCountLimit == 0 {
		result.QuizCountLimit = result.Setting.QuizCountLimit
	}
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

// ListSubjectLibsWithFolder returns subject libraries with an optional folder tree.
func (s *Service) ListSubjectLibsWithFolder(ctx context.Context, withFolder bool) (*SubjectLibsResponse, error) {
	value := "0"
	if withFolder {
		value = "1"
	}
	return s.ListSubjectLibs(ctx, map[string]string{"with_folder": value})
}

// ListCourseSubjectLibs returns subject libraries scoped to a course.
func (s *Service) ListCourseSubjectLibs(ctx context.Context, courseID int, withFolder bool) (*SubjectLibsResponse, error) {
	value := "0"
	if withFolder {
		value = "1"
	}
	u := fmt.Sprintf("/api/course/%d/subject-libs?with_folder=%s", courseID, value)
	result := new(SubjectLibsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListQuestionnaireSubjectLibs returns questionnaire subject libraries.
func (s *Service) ListQuestionnaireSubjectLibs(ctx context.Context) (*SubjectLibsResponse, error) {
	result := new(SubjectLibsResponse)
	_, err := s.client.Get(ctx, "/api/subject-libs?lib_type=questionnaire", result)
	return result, err
}

// GetSubjectLib returns a specific subject library.
func (s *Service) GetSubjectLib(ctx context.Context, libID int) (*SubjectLib, error) {
	u := fmt.Sprintf("/api/subject-libs/%d", libID)
	result := new(SubjectLib)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateSubjectLib creates a new subject library.
func (s *Service) CreateSubjectLib(ctx context.Context, body interface{}) (*SubjectLib, error) {
	result := new(SubjectLib)
	_, err := s.client.Post(ctx, "/api/subject-libs", body, result)
	return result, err
}

// CreateCourseSubjectLib creates a course-scoped subject library or folder.
func (s *Service) CreateCourseSubjectLib(ctx context.Context, courseID int, libType string, body interface{}) (*SubjectLib, error) {
	u := fmt.Sprintf("/api/course/%d/subject-libs?lib_type=%s", courseID, libType)
	result := new(SubjectLib)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// CreateQuestionnaireSubjectLib creates a questionnaire subject library.
func (s *Service) CreateQuestionnaireSubjectLib(ctx context.Context, body interface{}) (*SubjectLib, error) {
	result := new(SubjectLib)
	_, err := s.client.Post(ctx, "/api/subject-libs?lib_type=questionnaire", body, result)
	return result, err
}

// GetSubject returns a specific subject/question.
func (s *Service) GetSubject(ctx context.Context, subjectID int) (*ExamSubject, error) {
	u := fmt.Sprintf("/api/subjects/%d", subjectID)
	result := new(ExamSubject)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetSHTVUModules returns chapter/module data from the SHTVU study platform bridge.
func (s *Service) GetSHTVUModules(ctx context.Context, courseID int) (*SHTVUModulesResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/query-shtvu-modules", courseID)
	result := new(SHTVUModulesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// SearchSHTVUSubjects returns paginated SHTVU subjects for the given course.
func (s *Service) SearchSHTVUSubjects(ctx context.Context, courseID int, params *SHTVUSearchSubjectsParams) (*SHTVUSubjectsResponse, error) {
	query := map[string]string{}
	if params != nil {
		if params.Chapters != "" {
			query["chapters"] = params.Chapters
		}
		if params.SubjectType != "" {
			query["subject_type"] = params.SubjectType
		}
		if params.Difficulties != "" {
			query["difficulties"] = params.Difficulties
		}
		if params.Keyword != "" {
			query["keyword"] = params.Keyword
		}
		if params.PageIndex > 0 {
			query["page_index"] = fmt.Sprintf("%d", params.PageIndex)
		}
		if params.PageSize > 0 {
			query["page_size"] = fmt.Sprintf("%d", params.PageSize)
		}
	}
	u := addQueryParams(fmt.Sprintf("/api/courses/%d/query-shtvu-subjects", courseID), query)
	result := new(SHTVUSubjectsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetSHTVUSubjectTypesInfo returns the available subject-type counts for random import.
func (s *Service) GetSHTVUSubjectTypesInfo(ctx context.Context, courseID int) (*SHTVUSubjectTypesInfoResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/shtvu-subject-types-info", courseID)
	result := new(SHTVUSubjectTypesInfoResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ImportRandomExamSubjectsFromSHTVU imports randomly selected SHTVU subjects into an exam.
func (s *Service) ImportRandomExamSubjectsFromSHTVU(ctx context.Context, examID int, body *ImportRandomSubjectsFromSHTVURequest) (*SubjectsResponse, error) {
	u := fmt.Sprintf("/api/exams/%d/import-random-subjects-from-shtvu", examID)
	result := new(SubjectsResponse)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// ImportRandomClassroomSubjectsFromSHTVU imports randomly selected SHTVU subjects into a classroom.
func (s *Service) ImportRandomClassroomSubjectsFromSHTVU(ctx context.Context, classroomID int, body *ImportRandomSubjectsFromSHTVURequest) (*SubjectsResponse, error) {
	u := fmt.Sprintf("/api/classrooms/%d/import-random-subjects-from-shtvu", classroomID)
	result := new(SubjectsResponse)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// ImportVideoQuizSubjectsFromSHTVU imports selected SHTVU subjects into a video quiz.
func (s *Service) ImportVideoQuizSubjectsFromSHTVU(ctx context.Context, videoQuizID int, body *ImportRandomSubjectsFromSHTVURequest) (*SubjectsResponse, error) {
	u := fmt.Sprintf("/api/video-quizzes/%d/import-subjects-from-shtvu", videoQuizID)
	result := new(SubjectsResponse)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// SearchSubjectsInLib returns subjects in a subject library filtered by keyword/type.
func (s *Service) SearchSubjectsInLib(ctx context.Context, libID int, keyword, subjectType string) ([]*ExamSubject, error) {
	u := fmt.Sprintf("/api/subject-libs/%d?keyword=%s", libID, url.QueryEscape(keyword))
	if subjectType != "" {
		u += "&subject_type=" + url.QueryEscape(subjectType)
	}
	var result struct {
		Subjects []*ExamSubject `json:"subjects"`
	}
	_, err := s.client.Get(ctx, u, &result)
	return result.Subjects, err
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

// CopySubjectLibToExam copies a subject library into an exam.
func (s *Service) CopySubjectLibToExam(ctx context.Context, examID, libID int) error {
	u := fmt.Sprintf("/api/subject-libs/%d/copy?examId=%d", libID, examID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// CopySubjectLibToClassroom copies a subject library into a classroom quiz.
func (s *Service) CopySubjectLibToClassroom(ctx context.Context, classroomID, libID int) error {
	u := fmt.Sprintf("/api/subject-libs/%d/copy?classroomId=%d", libID, classroomID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// CopySubjectLibToCoursewareQuiz copies a subject library into a courseware quiz.
func (s *Service) CopySubjectLibToCoursewareQuiz(ctx context.Context, videoQuizID, libID int) error {
	u := fmt.Sprintf("/api/subject-libs/%d/copy?videoQuizId=%d", libID, videoQuizID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// CopySubjectLibToQuestionnaire copies a subject library into a questionnaire.
func (s *Service) CopySubjectLibToQuestionnaire(ctx context.Context, questionnaireID, libID int) error {
	u := fmt.Sprintf("/api/subject-libs/%d/copy?questionnaireId=%d", libID, questionnaireID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
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

// BatchCopySubjectLibsToCoursewareQuiz copies multiple subject libraries to a courseware quiz context.
func (s *Service) BatchCopySubjectLibsToCoursewareQuiz(ctx context.Context, coursewareQuizID int, body *BatchCopySubjectLibsRequest) error {
	u := fmt.Sprintf("/api/subject-libs/batch/copy?courseware_quiz_id=%d", coursewareQuizID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// --- Rubrics ---

// ListRubrics returns rubrics.
func (s *Service) ListRubrics(ctx context.Context, params map[string]string) (*RubricsResponse, error) {
	u := addQueryParams("/api/rubrics", params)
	result := new(RubricsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListRubricsWithResource returns rubrics together with resource metadata used by the web UI.
func (s *Service) ListRubricsWithResource(ctx context.Context, opts *model.ListOptions, keyword string) (*RubricsResponse, error) {
	u := addListOptions("/api/rubrics-with-resource", opts)
	u = addQueryParams(u, map[string]string{
		"keyword":      keyword,
		"no-intercept": "true",
		"fields":       "id,name,conditions,created_by,created_at,group_id,group_name,is_shared_rubric",
	})
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
func (s *Service) CreateRubric(ctx context.Context, body interface{}) (*model.Rubric, error) {
	u := addQueryParams("/api/rubrics", map[string]string{"fields": "id,name,conditions"})
	result := new(model.Rubric)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateRubric updates a rubric.
func (s *Service) UpdateRubric(ctx context.Context, rubricID int, body interface{}) (*model.Rubric, error) {
	u := addQueryParams(fmt.Sprintf("/api/rubrics/%d", rubricID), map[string]string{
		"fields": "id,name,conditions,engage_number,created_by",
	})
	result := new(model.Rubric)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteRubric deletes a rubric.
func (s *Service) DeleteRubric(ctx context.Context, rubricID int) error {
	u := fmt.Sprintf("/api/rubric/%d?fields=id", rubricID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// UpdateSubjectGroup updates a subject group.
func (s *Service) UpdateSubjectGroup(ctx context.Context, subjectGroupID int, body interface{}) (*SubjectGroup, error) {
	u := fmt.Sprintf("/api/subject-group/%d", subjectGroupID)
	result := new(SubjectGroupResponse)
	_, err := s.client.Put(ctx, u, body, result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// DeleteSubjectGroup deletes a subject group.
func (s *Service) DeleteSubjectGroup(ctx context.Context, subjectGroupID int) error {
	u := fmt.Sprintf("/api/subject-group/%d", subjectGroupID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ListSubjectGroups returns subject groups for an exam/video-quiz/classroom target.
func (s *Service) ListSubjectGroups(ctx context.Context, targetType string, targetID int) ([]*SubjectGroup, error) {
	u := fmt.Sprintf("/api/%s/%d/subject-groups", targetType, targetID)
	result := new(SubjectGroupsResponse)
	_, err := s.client.Get(ctx, u, result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// CreateSubjectGroup creates a subject group for an exam/video-quiz/classroom target.
func (s *Service) CreateSubjectGroup(ctx context.Context, targetType string, targetID int, body *SubjectGroupRequest) (*SubjectGroup, error) {
	u := fmt.Sprintf("/api/%s/%d/subject-group", targetType, targetID)
	result := new(SubjectGroupResponse)
	_, err := s.client.Post(ctx, u, body, result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// SortSubjectGroupSubjects sorts subjects within a subject group.
func (s *Service) SortSubjectGroupSubjects(ctx context.Context, subjectGroupID int, body interface{}) error {
	u := fmt.Sprintf("/api/subject-group/%d/subjects/sort", subjectGroupID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// SortExamSubjectGroups sorts subject groups within an exam.
func (s *Service) SortExamSubjectGroups(ctx context.Context, examID int, body interface{}) error {
	u := fmt.Sprintf("/api/exam/%d/subject-groups/sort", examID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// GenerateRubric generates a rubric using AI.
func (s *Service) GenerateRubric(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/rubrics/generate", body, &result)
	return result, err
}

// --- Make-up Exams ---

// MakeUpExam creates a make-up exam record.
func (s *Service) MakeUpExam(ctx context.Context, body MakeUpExamRequest) error {
	_, err := s.client.Post(ctx, "/api/make-up-exams", body, nil)
	return err
}

// ImportMakeUpExamSubjects imports subjects into a make-up exam.
func (s *Service) ImportMakeUpExamSubjects(ctx context.Context, examID int) error {
	u := fmt.Sprintf("/api/make-up-exams/%d/subjects/import", examID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// MakeupExam creates a makeup exam (alternate endpoint).
func (s *Service) MakeupExam(ctx context.Context, body MakeupExamRequest) error {
	_, err := s.client.Post(ctx, "/api/makeup-exams", body, nil)
	return err
}

// ListCourseExams returns all exams for a course.
func (s *Service) ListCourseExams(ctx context.Context, courseID int) (*CourseExamsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/exams", courseID)
	result := new(CourseExamsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListCourseClassrooms returns classroom activities for a course.
func (s *Service) ListCourseClassrooms(ctx context.Context, courseID int) (*CourseClassroomListResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/classroom-list", courseID)
	result := new(CourseClassroomListResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListSubmittedExams returns IDs of submitted exams for a course.
func (s *Service) ListSubmittedExams(ctx context.Context, courseID int) (*SubmittedExamsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/submitted-exams", courseID)
	result := new(SubmittedExamsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}
