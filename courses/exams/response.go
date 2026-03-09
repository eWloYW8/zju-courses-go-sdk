package exams

import "encoding/json"

type ExamSubmissionResponse = json.RawMessage
type ExamScoreResponse = json.RawMessage
type ClassroomExamResponse = json.RawMessage
type ExamScoreDistributionResponse map[string]any

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
