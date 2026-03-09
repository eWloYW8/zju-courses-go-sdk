package exams

type CreateSubjectLibRequest struct {
	Title    string `json:"title"`
	ParentID int    `json:"parent_id,omitempty"`
}

type ListCourseExamListParams struct {
	Page       int
	PageSize   int
	Conditions any
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

type CreateCoursewareQuizRequest struct {
	UploadReferenceID int `json:"upload_reference_id,omitempty"`
	Subjects          any `json:"subjects,omitempty"`
}

type GenerateCoursewareQuizSubjectsRequest struct {
	UploadID               int      `json:"upload_id,omitempty"`
	UploadReferenceID      int      `json:"upload_reference_id,omitempty"`
	ModuleID               int      `json:"module_id,omitempty"`
	ModuleType             string   `json:"module_type,omitempty"`
	GroupID                string   `json:"group_id,omitempty"`
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

type GenerateSubjectsRequest = GenerateCoursewareQuizSubjectsRequest

type GenerateSubjectsByTextRequest struct {
	TextContent string `json:"text_content"`
	GenerateCoursewareQuizSubjectsRequest
}

type UpdateExamSubjectExplanationRequest struct {
	AnswerExplanation string `json:"answer_explanation,omitempty"`
}

type ListExamRetakeRecordsParams struct {
	Page     int
	PageSize int
}

type ExamScoreDistributionConditions map[string]any

type MakeUpExamRequest map[string]any

type MakeupExamRequest map[string]any

type ExamSubjectsStatParams struct {
	ExamPaperType string
	Conditions    any
}

type ExamScoreListParams struct {
	Conditions any
}

type ExamExamineesParams struct {
	Conditions any
}

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

type SubjectGroupRequest struct {
	SubjectType  string `json:"subject_type,omitempty"`
	ReferrerType string `json:"referrer_type,omitempty"`
	ReferrerID   int    `json:"referrer_id,omitempty"`
	Sort         int    `json:"sort,omitempty"`
}

type ClassroomStatusRequest struct {
	Status int `json:"status"`
}

type UpdateClassroomSubjectStatusRequest struct {
	Status int `json:"status"`
}

type ClassroomScoreListParams struct {
	IgnoreAvatar bool
	ExamineeIDs  []int
}

type CoursewareQuizSubjectStatisticParams struct {
	Page       int
	PageSize   int
	Conditions any
}

type SearchExamSubjectsParams struct {
	IsMakeupExam bool
	Keyword      string
	SubjectType  string
}

type UpdateExamSubmissionCommentRequest struct {
	Comment string `json:"comment,omitempty"`
}

type CampusSubjectSelection struct {
	ID    int    `json:"id"`
	Count int    `json:"count"`
	Point string `json:"point,omitempty"`
}

type ImportSubjectsFromCampusConditions struct {
	ExamSubjectTypes            []string `json:"exam_subject_types,omitempty"`
	ExamSubjectDifficultyLevels []string `json:"exam_subject_difficulty_levels,omitempty"`
}

type ImportSubjectsFromCampusRequest struct {
	Items                       []*CampusSubjectSelection           `json:"items"`
	Settings                    any                                 `json:"settings,omitempty"`
	ExamSubjectTypes            []string                            `json:"exam_subject_types,omitempty"`
	ExamSubjectDifficultyLevels []string                            `json:"exam_subject_difficulty_levels,omitempty"`
	Conditions                  *ImportSubjectsFromCampusConditions `json:"conditions,omitempty"`
}

type ExamPointRule struct {
	RuleName            string `json:"rule_name,omitempty"`
	RulePoint           string `json:"rule_point,omitempty"`
	RuleNumber          string `json:"rule_number,omitempty"`
	RuleDifficultyLevel string `json:"rule_difficulty_level,omitempty"`
}

type ExamSubjectPointsAndRules struct {
	SubjectIndex int              `json:"subject_index"`
	PointRules   []*ExamPointRule `json:"point_rules"`
}

type UpdateExamPointsAndRulesRequest struct {
	SubjectsPointsAndRules []*ExamSubjectPointsAndRules `json:"subjects_points_and_rules"`
}
