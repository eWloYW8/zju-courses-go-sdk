package exams

type CreateSubjectLibRequest struct {
	Title    string `json:"title"`
	ParentID int    `json:"parent_id,omitempty"`
}

type BatchCopySubjectLibsRequest struct {
	LibIDs    []int `json:"lib_ids"`
	SubjectID int   `json:"subject_id,omitempty"`
	CourseID  int   `json:"course_id,omitempty"`
}

type CoursewareQuizSubjectsRequest struct {
	Subjects any `json:"subjects"`
}

type MakeUpExamRequest map[string]any

type MakeupExamRequest map[string]any
