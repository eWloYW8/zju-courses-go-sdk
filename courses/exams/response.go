package exams

import (
	"encoding/json"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

type ExamSubmissionResponse = json.RawMessage
type ExamScoreResponse = json.RawMessage
type ClassroomExamResponse = json.RawMessage
type ExamScoreDistributionResponse map[string]any
type CourseExamListResponse map[string]any
type ExamSubjectsStatResponse map[string]any
type ExamSubmissionsResponse map[string]any
type ExamScoreListResponse map[string]any
type ExamExamineeResponse map[string]any
type ExamSubjectsSummaryResponse map[string]any
type ExamExamineesResponse map[string]any
type ExamGroupsResponse map[string]any
type ExamSubmissionCountStatusResponse map[string]any
type ExamSubmissionSyncResponse map[string]any
type ExamSubmissionSyncTaskProgressResponse map[string]any
type ExamSubjectExamineesResponse map[string]any
type ExamSubjectGroupsResponse map[string]any
type ExamSubjectSubmissionsResponse map[string]any
type ExamScoreOperationResponse map[string]any

type ExamRetakeRecord map[string]any

type ExamRetakeRecordsResponse struct {
	Items []*ExamRetakeRecord `json:"items,omitempty"`
	model.Pagination
}

type ExamPointsAndRulesResponse struct {
	Message                    string `json:"message,omitempty"`
	SelectSubjectsRandomlyRule any    `json:"select_subjects_randomly_rule,omitempty"`
}

type CourseExam struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
}

type CourseExamsResponse struct {
	Exams []*CourseExam `json:"exams"`
}

type CourseClassroom struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
}

type CourseClassroomListResponse struct {
	Classrooms []*CourseClassroom `json:"classrooms"`
}

type SubmittedExamsResponse struct {
	ExamIDs []int `json:"exam_ids"`
}

type SubjectGroupsResponse struct {
	Data []*SubjectGroup `json:"data,omitempty"`
}

type SubjectGroupResponse struct {
	Data *SubjectGroup `json:"data,omitempty"`
}

type ClassroomMySubmissionsResponse struct {
	Submissions []*ClassroomSubmission `json:"submissions,omitempty"`
}

type ClassroomExamineesResponse struct {
	Examinees []*ClassroomExaminee `json:"examinees,omitempty"`
}

type ClassroomScoreListResponse struct {
	Examinees []*ClassroomExaminee `json:"examinees,omitempty"`
}

type ExamPaperZipStatusResponse struct {
	PaperZip *ExamPaperZip `json:"paper_zip,omitempty"`
}

type CoursewareQuizSubjectStatisticResponse struct {
	Items []*UserAnswerSubjectStatistic `json:"items,omitempty"`
	model.Pagination
}
