package exams

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type SubjectLib struct {
	ID            int           `json:"id"`
	Title         string        `json:"title,omitempty"`
	ParentID      int           `json:"parent_id,omitempty"`
	IsFolder      bool          `json:"is_folder,omitempty"`
	Children      []*SubjectLib `json:"children,omitempty"`
	SubjectsCount int           `json:"subjects_count,omitempty"`
}

type SubjectLibsResponse struct {
	SubjectLibs []*SubjectLib `json:"subject_libs"`
}

type RubricsResponse struct {
	Rubrics []*model.Rubric `json:"rubrics"`
}

type CoursewareQuizSettings struct {
	QuizCountLimit int `json:"quiz_count_limit,omitempty"`
}

type CoursewareQuizSettingsResponse struct {
	Setting *CoursewareQuizSettings `json:"setting"`
}

type CoursewareQuizSubjectsResponse struct {
	QuizID   int            `json:"quiz_id,omitempty"`
	Subjects []*ExamSubject `json:"subjects,omitempty"`
}

type CoursewareQuizUpdateResponse struct {
	QuizID int `json:"quiz_id"`
}

// Exam represents an exam/quiz activity.
type Exam struct {
	ID                              int                `json:"id"`
	Title                           string             `json:"title,omitempty"`
	Type                            string             `json:"type,omitempty"`
	UniqueKey                       string             `json:"unique_key,omitempty"`
	CourseID                        int                `json:"course_id,omitempty"`
	ModuleID                        int                `json:"module_id,omitempty"`
	ModuleSort                      int                `json:"module_sort,omitempty"`
	SyllabusID                      int                `json:"syllabus_id,omitempty"`
	SyllabusSort                    string             `json:"syllabus_sort,omitempty"`
	Sort                            int                `json:"sort,omitempty"`
	ReferrerID                      int                `json:"referrer_id,omitempty"`
	ReferrerType                    string             `json:"referrer_type,omitempty"`
	ImportedFrom                    string             `json:"imported_from,omitempty"`
	TeachingModel                   string             `json:"teaching_model,omitempty"`
	UsingPhase                      string             `json:"using_phase,omitempty"`
	CreatedAt                       string             `json:"created_at,omitempty"`
	StartTime                       *string            `json:"start_time,omitempty"`
	EndTime                         *string            `json:"end_time,omitempty"`
	PublishTime                     string             `json:"publish_time,omitempty"`
	Published                       bool               `json:"published,omitempty"`
	EnableEdit                      bool               `json:"enable_edit,omitempty"`
	Description                     string             `json:"description,omitempty"`
	AnnounceAnswerStatus            string             `json:"announce_answer_status,omitempty"`
	AnnounceAnswerType              string             `json:"announce_answer_type,omitempty"`
	AnnounceAnswerTime              *string            `json:"announce_answer_time,omitempty"`
	AnnounceScoreStatus             string             `json:"announce_score_status,omitempty"`
	AnnounceScoreTime               *string            `json:"announce_score_time,omitempty"`
	IsAnswerAnnounced               bool               `json:"is_answer_announced,omitempty"`
	IsAnnounceAnswerTimePassed      bool               `json:"is_announce_answer_time_passed,omitempty"`
	IsAnnounceScoreTimePassed       bool               `json:"is_announce_score_time_passed,omitempty"`
	ScorePercentage                 string             `json:"score_percentage,omitempty"`
	ScoreRule                       string             `json:"score_rule,omitempty"`
	ScoreType                       string             `json:"score_type,omitempty"`
	TotalScore                      *float64           `json:"total_score,omitempty"`
	IsScorePublic                   bool               `json:"is_score_public,omitempty"`
	GroupSetID                      int                `json:"group_set_id,omitempty"`
	GroupSetName                    string             `json:"group_set_name,omitempty"`
	HasAssignGroup                  bool               `json:"has_assign_group,omitempty"`
	HasAssignStudent                bool               `json:"has_assign_student,omitempty"`
	AssignGroupIDs                  []int              `json:"assign_group_ids,omitempty"`
	AssignStudentIDs                []int              `json:"assign_student_ids,omitempty"`
	IsAssignedToAll                 bool               `json:"is_assigned_to_all,omitempty"`
	SubmitByGroup                   bool               `json:"submit_by_group,omitempty"`
	SubmitTimes                     int                `json:"submit_times,omitempty"`
	CompletionCriterion             string             `json:"completion_criterion,omitempty"`
	CompletionCriterionKey          string             `json:"completion_criterion_key,omitempty"`
	CompletionCriterionValue        string             `json:"completion_criterion_value,omitempty"`
	IsStarted                       bool               `json:"is_started,omitempty"`
	IsClosed                        bool               `json:"is_closed,omitempty"`
	IsInProgress                    bool               `json:"is_in_progress,omitempty"`
	IsOpenedCatalog                 bool               `json:"is_opened_catalog,omitempty"`
	IsPracticeMode                  bool               `json:"is_practice_mode,omitempty"`
	IsQuizControlBySubject          bool               `json:"is_quiz_control_by_subject,omitempty"`
	IsQuizPublic                    bool               `json:"is_quiz_public,omitempty"`
	IsTimed                         *bool              `json:"is_timed,omitempty"`
	LimitTime                       *int               `json:"limit_time,omitempty"`
	Duration                        *int               `json:"duration,omitempty"`
	ExamPaperTemplateID             int                `json:"exam_paper_template_id,omitempty"`
	Mode                            string             `json:"mode,omitempty"`
	HasTimeLimit                    bool               `json:"has_time_limit,omitempty"`
	EnableAutoAIGrading             bool               `json:"enable_auto_ai_grading,omitempty"`
	CheckSubmitIPConsistency        bool               `json:"check_submit_ip_consistency,omitempty"`
	IsIPConstrained                 bool               `json:"is_ip_constrained,omitempty"`
	LimitedIP                       *string            `json:"limited_ip,omitempty"`
	DisableCopyPaste                bool               `json:"disable_copy_paste,omitempty"`
	DisableDevtool                  bool               `json:"disable_devtool,omitempty"`
	DisableRightClick               bool               `json:"disable_right_click,omitempty"`
	EnableAntiCheat                 bool               `json:"enable_anti_cheat,omitempty"`
	EnableInvigilation              bool               `json:"enable_invigilation,omitempty"`
	IsFullscreenMode                bool               `json:"is_fullscreen_mode,omitempty"`
	IsLeavingWindowConstrained      bool               `json:"is_leaving_window_constrained,omitempty"`
	IsLeavingWindowTimeout          bool               `json:"is_leaving_window_timeout,omitempty"`
	LeavingWindowLimit              *int               `json:"leaving_window_limit,omitempty"`
	LeavingWindowTimeout            *int               `json:"leaving_window_timeout,omitempty"`
	LimitAnswerOnSingleClient       bool               `json:"limit_answer_on_signle_client,omitempty"`
	Platform                        string             `json:"platform,omitempty"`
	EZTest                          any                `json:"eztest,omitempty"`
	DefaultOptionsLayout            string             `json:"default_options_layout,omitempty"`
	SubjectIndexType                string             `json:"subject_index_type,omitempty"`
	KnowledgeNodeIDs                []int              `json:"knowledge_node_ids,omitempty"`
	KnowledgeNodeReference          []any              `json:"knowledge_node_reference,omitempty"`
	Prerequisites                   []any              `json:"prerequisites,omitempty"`
	ScoreItemGroupID                int                `json:"score_item_group_id,omitempty"`
	ScoreItemGroupName              *string            `json:"score_item_group_name,omitempty"`
	ScoreItemScored                 *bool              `json:"score_item_scored,omitempty"`
	SubjectsRule                    *model.SubjectsRule `json:"subjects_rule,omitempty"`
	MakeUpRecord                    any                `json:"make_up_record,omitempty"`
	ExamSubmissions                 []any              `json:"exam_submissions,omitempty"`
	SubjectsCount                   int                `json:"subjects_count,omitempty"`
	Data                            any                `json:"data,omitempty"`
}

// Classroom represents an in-class activity / classroom quiz.
type Classroom struct {
	ID                     int     `json:"id"`
	Title                  string  `json:"title,omitempty"`
	Type                   string  `json:"type,omitempty"`
	UniqueKey              string  `json:"unique_key,omitempty"`
	Status                 string  `json:"status,omitempty"`
	CourseID               int     `json:"course_id,omitempty"`
	ModuleID               int     `json:"module_id,omitempty"`
	ModuleSort             int     `json:"module_sort,omitempty"`
	SyllabusID             int     `json:"syllabus_id,omitempty"`
	SyllabusSort           string  `json:"syllabus_sort,omitempty"`
	Sort                   int     `json:"sort,omitempty"`
	ImportedFrom           string  `json:"imported_from,omitempty"`
	TeachingModel          string  `json:"teaching_model,omitempty"`
	UsingPhase             string  `json:"using_phase,omitempty"`
	CreatedAt              string  `json:"created_at,omitempty"`
	UpdatedAt              string  `json:"updated_at,omitempty"`
	UpdatedStatusAt        string  `json:"updated_status_at,omitempty"`
	StartAt                string  `json:"start_at,omitempty"`
	FinishAt               *string `json:"finish_at,omitempty"`
	Published              bool    `json:"published,omitempty"`
	EnableEdit             bool    `json:"enable_edit,omitempty"`
	Duration               *int    `json:"duration,omitempty"`
	ExamPaperTemplateID    int     `json:"exam_paper_template_id,omitempty"`
	AnnounceAnswerStatus   string  `json:"announce_answer_status,omitempty"`
	IsAnswerAnnounced      bool    `json:"is_answer_announced,omitempty"`
	IsInProgress           bool    `json:"is_in_progress,omitempty"`
	IsOpenedCatalog        bool    `json:"is_opened_catalog,omitempty"`
	IsQuizControlBySubject bool    `json:"is_quiz_control_by_subject,omitempty"`
	IsQuizPublic           bool    `json:"is_quiz_public,omitempty"`
	IsScorePublic          bool    `json:"is_score_public,omitempty"`
	IsTimed                *bool   `json:"is_timed,omitempty"`
	SubjectsCount          int     `json:"subjects_count,omitempty"`
	StartedSubjectsCount   int     `json:"started_subjects_count,omitempty"`
	FinishedSubjectsCount  int     `json:"finished_subjects_count,omitempty"`
	ScorePercentage        string  `json:"score_percentage,omitempty"`
	ScoreType              string  `json:"score_type,omitempty"`
	ScoreItemGroupID       int     `json:"score_item_group_id,omitempty"`
	ScoreItemGroupName     *string `json:"score_item_group_name,omitempty"`
	ScoreItemScored        *bool   `json:"score_item_scored,omitempty"`
	Data                   any     `json:"data,omitempty"`
}

// ExamSubject represents a question/subject within an exam.
type ExamSubject struct {
	ID                int              `json:"id"`
	UUID              string           `json:"uuid,omitempty"`
	Type              string           `json:"type"`
	Description       string           `json:"description,omitempty"`
	Point             any              `json:"point,omitempty"`
	Score             any              `json:"score,omitempty"`
	Sort              int              `json:"sort,omitempty"`
	ParentID          *int             `json:"parent_id,omitempty"`
	GroupID           *int             `json:"group_id,omitempty"`
	DifficultyLevel   string           `json:"difficulty_level,omitempty"`
	AnswerExplanation string           `json:"answer_explanation,omitempty"`
	WrongExplanation  string           `json:"wrong_explanation,omitempty"`
	CorrectAnswers    []any            `json:"correct_answers,omitempty"`
	Answers           []any            `json:"answers,omitempty"`
	Options           []*SubjectOption `json:"options,omitempty"`
	SubSubjects       []*ExamSubject   `json:"sub_subjects,omitempty"`
	Attachments       []any            `json:"attachments,omitempty"`
	OptionsLayout     string           `json:"options_layout,omitempty"`
	OptionType        string           `json:"option_type,omitempty"`
	ScoreRule         string           `json:"score_rule,omitempty"`
	CaseSensitive     bool             `json:"case_sensitive,omitempty"`
	Unordered         bool             `json:"unordered,omitempty"`
	PlayLimit         bool             `json:"play_limit,omitempty"`
	PlayLimitTimes    int              `json:"play_limit_times,omitempty"`
	HasAudio          bool             `json:"has_audio,omitempty"`
	KnowledgeNodeIDs  []int            `json:"knowledge_node_ids,omitempty"`
	Duration          *int             `json:"duration,omitempty"`
	IsTimed           bool             `json:"is_timed,omitempty"`
}

// SubjectOption represents an option for an exam subject.
type SubjectOption struct {
	ID       int    `json:"id,omitempty"`
	Content  string `json:"content,omitempty"`
	IsAnswer bool   `json:"is_answer,omitempty"`
	Type     string `json:"type,omitempty"`
	Sort     int    `json:"sort,omitempty"`
}

// SubjectGroup represents a group of subjects in an exam.
type SubjectGroup struct {
	ID           int            `json:"id,omitempty"`
	SubjectType  string         `json:"subject_type,omitempty"`
	ReferrerType string         `json:"referrer_type,omitempty"`
	ReferrerID   int            `json:"referrer_id,omitempty"`
	Sort         int            `json:"sort,omitempty"`
	Subjects     []*ExamSubject `json:"subjects,omitempty"`
}

// CoursewareQuiz represents a courseware quiz.
type CoursewareQuiz struct {
	ID                int `json:"id,omitempty"`
	SubjectsCount     int `json:"subjects_count,omitempty"`
	UploadReferenceID int `json:"upload_reference_id,omitempty"`
	SubmittedTimes    int `json:"submitted_times,omitempty"`
}
