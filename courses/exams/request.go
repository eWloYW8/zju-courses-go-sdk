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

type SaveSubjectsRequest struct {
	Subjects any `json:"subjects"`
}

type UpdateCoursewareQuizSubjectsRequest struct {
	Subjects any `json:"subjects"`
}

type CoursewareQuizSubjectsRequest = UpdateCoursewareQuizSubjectsRequest

type GenerateCoursewareQuizSubjectsRequest struct {
	UploadReferenceID      int      `json:"upload_reference_id"`
	NumOfSingleSelection   int      `json:"num_of_single_selection,omitempty"`
	NumOfMultipleSelection int      `json:"num_of_multiple_selection,omitempty"`
	NumOfFillInBlank       int      `json:"num_of_fill_in_blank,omitempty"`
	NumOfTrueOrFalse       int      `json:"num_of_true_or_false,omitempty"`
	NumOfShortAnswer       int      `json:"num_of_short_answer,omitempty"`
	BloomCognitiveDomains  []string `json:"bloom_cognitive_domains,omitempty"`
	QuizKnowledgePoints    []any    `json:"quiz_knowledge_points,omitempty"`
	Locale                 string   `json:"locale,omitempty"`
	Stream                 bool     `json:"stream,omitempty"`
	PageRange              []int    `json:"page_range,omitempty"`
}

type UpdateExamSubjectExplanationRequest struct {
	AnswerExplanation string `json:"answer_explanation,omitempty"`
}

type ExamScoreDistributionConditions map[string]any

type MakeUpExamRequest map[string]any

type MakeupExamRequest map[string]any

type SHTVUSearchSubjectsParams struct {
	Chapters     string `json:"chapters,omitempty"`
	SubjectType  string `json:"subject_type,omitempty"`
	Difficulties string `json:"difficulties,omitempty"`
	Keyword      string `json:"keyword,omitempty"`
	PageIndex    int    `json:"page_index,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
}

type SHTVURandomImportItem struct {
	SubjectType string `json:"subject_type,omitempty"`
	Count       int    `json:"count,omitempty"`
	Point       any    `json:"point,omitempty"`
}

type ImportRandomSubjectsFromSHTVURequest struct {
	Items     []*SHTVURandomImportItem `json:"items"`
	Timestamp string                   `json:"timestamp,omitempty"`
}
