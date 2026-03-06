package model

// InterScoreMap represents inter-review score mapping.
type InterScoreMap struct {
	ID        int     `json:"id"`
	StartTime *string `json:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty"`
	IsStarted *bool   `json:"is_started,omitempty"`
	IsClosed  *bool   `json:"is_closed,omitempty"`
	PiecesCnt int     `json:"pieces_cnt,omitempty"`
}

// IntraScoreMap represents intra-review score mapping.
type IntraScoreMap struct {
	ID        int     `json:"id"`
	StartTime *string `json:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty"`
	IsClosed  *bool   `json:"is_closed,omitempty"`
}

// Rubric represents a rubric definition.
type Rubric struct {
	ID             int                `json:"id"`
	Name           *string            `json:"name,omitempty"`
	Conditions     []*RubricCondition `json:"conditions,omitempty"`
	CreatedAt      *string            `json:"created_at,omitempty"`
	UpdatedAt      *string            `json:"updated_at,omitempty"`
	CreatedBy      *User              `json:"created_by,omitempty"`
	EngageNumber   int                `json:"engage_number,omitempty"`
	GroupID        *int               `json:"group_id,omitempty"`
	GroupName      *string            `json:"group_name,omitempty"`
	IsSharedRubric bool               `json:"is_shared_rubric,omitempty"`
	OrgID          int                `json:"org_id,omitempty"`
}

// RubricInstance represents a rubric instance attached to an activity.
type RubricInstance struct {
	ID         int                `json:"id"`
	Name       *string            `json:"name,omitempty"`
	Conditions []*RubricCondition `json:"conditions,omitempty"`
	Rubric     *Rubric            `json:"rubric,omitempty"`
	RubricID   int                `json:"rubric_id,omitempty"`
}

// RubricCondition represents a rubric condition with levels.
type RubricCondition struct {
	ID          int            `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Weight      float64        `json:"weight,omitempty"`
	Levels      []*RubricLevel `json:"levels,omitempty"`
}

// RubricLevel represents a rubric scoring level.
type RubricLevel struct {
	ID          int     `json:"id,omitempty"`
	Description string  `json:"description,omitempty"`
	Score       float64 `json:"score,omitempty"`
	Name        string  `json:"name,omitempty"`
}

// Activity represents a course activity (homework, forum, material, exam, etc.)
type Activity struct {
	ID                              int         `json:"id"`
	Title                           string      `json:"title"`
	Type                            string      `json:"type"`
	CourseID                        int         `json:"course_id,omitempty"`
	ModuleID                        int         `json:"module_id,omitempty"`
	SyllabusID                      int         `json:"syllabus_id,omitempty"`
	Sort                            int         `json:"sort,omitempty"`
	CreatedAt                       string      `json:"created_at,omitempty"`
	UpdatedAt                       string      `json:"updated_at,omitempty"`
	StartTime                       *string     `json:"start_time,omitempty"`
	EndTime                         *string     `json:"end_time,omitempty"`
	Description                     string      `json:"description,omitempty"`
	Data                            interface{} `json:"data,omitempty"`
	EnableEdit                      bool        `json:"enable_edit,omitempty"`
	Published                       bool        `json:"published,omitempty"`
	PublishTime                     *string     `json:"publish_time,omitempty"`
	UniqueKey                       string      `json:"unique_key,omitempty"`
	TeachingModel                   string      `json:"teaching_model,omitempty"`
	TeachingUnitID                  int         `json:"teaching_unit_id,omitempty"`
	UsingPhase                      string      `json:"using_phase,omitempty"`
	Version                         *int        `json:"version,omitempty"`
	ImportedFrom                    *string     `json:"imported_from,omitempty"`
	ImportedTrackID                 *string     `json:"imported_track_id,omitempty"`
	ReferrerID                      int         `json:"referrer_id,omitempty"`
	ReferrerType                    string      `json:"referrer_type,omitempty"`

	// Group assignment
	GroupSetID       int         `json:"group_set_id,omitempty"`
	GroupSetName     *string     `json:"group_set_name,omitempty"`
	HasAssignGroup   bool        `json:"has_assign_group,omitempty"`
	HasAssignStudent bool        `json:"has_assign_student,omitempty"`
	AssignGroupIDs   []int       `json:"assign_group_ids,omitempty"`
	AssignStudentIDs []int       `json:"assign_student_ids,omitempty"`
	IsAssignedToAll  bool        `json:"is_assigned_to_all,omitempty"`
	SubmitByGroup    bool        `json:"submit_by_group,omitempty"`
	SubmitTimes      *int        `json:"submit_times,omitempty"`
	AssignTargets    interface{} `json:"assign_targets,omitempty"`

	// Completion
	CompletionCriterion      string      `json:"completion_criterion,omitempty"`
	CompletionCriterionKey   string      `json:"completion_criterion_key,omitempty"`
	CompletionCriterionValue interface{} `json:"completion_criterion_value,omitempty"`

	// Score
	CanShowScore       bool     `json:"can_show_score,omitempty"`
	AverageScore       *float64 `json:"average_score,omitempty"`
	HighestScore       *float64 `json:"highest_score,omitempty"`
	LowestScore        *float64 `json:"lowest_score,omitempty"`
	HasScoreCount      int      `json:"has_score_count,omitempty"`
	ScorePercentage    *string  `json:"score_percentage,omitempty"`
	ScorePublished     *bool    `json:"score_published,omitempty"`
	ScoreType          *string  `json:"score_type,omitempty"`
	ScoreItemGroupID   int      `json:"score_item_group_id,omitempty"`
	ScoreItemGroupName *string  `json:"score_item_group_name,omitempty"`
	ScoreItemScored    *bool    `json:"score_item_scored,omitempty"`
	ScoreRule          string   `json:"score_rule,omitempty"`
	TotalScore         *float64 `json:"total_score,omitempty"`
	IsScorePublic      bool     `json:"is_score_public,omitempty"`

	// Answer/Score announcement
	AnnounceAnswerStatus             string  `json:"announce_answer_status,omitempty"`
	AnnounceAnswerType               string  `json:"announce_answer_type,omitempty"`
	AnnounceAnswerTime               *string `json:"announce_answer_time,omitempty"`
	AnnounceScoreStatus              string  `json:"announce_score_status,omitempty"`
	AnnounceScoreTime                *string `json:"announce_score_time,omitempty"`
	IsAnswerAnnounced                bool    `json:"is_answer_announced,omitempty"`
	AnnounceAnswerAndExplanation     *bool   `json:"announce_answer_and_explanation,omitempty"`

	// Rubric
	RubricID              int             `json:"rubric_id,omitempty"`
	RubricInstance        *RubricInstance `json:"rubric_instance,omitempty"`
	RubricInstanceID      int             `json:"rubric_instance_id,omitempty"`
	IntraRubricID         int             `json:"intra_rubric_id,omitempty"`
	IntraRubricInstance   *RubricInstance `json:"intra_rubric_instance,omitempty"`
	IntraRubricInstanceID int             `json:"intra_rubric_instance_id,omitempty"`

	// Review
	InterReviewNamed         *bool          `json:"inter_review_named,omitempty"`
	InterScoreMap            *InterScoreMap `json:"inter_score_map,omitempty"`
	IntraScoreMap            *IntraScoreMap `json:"intra_score_map,omitempty"`
	IsInterReviewBySubmitter *bool          `json:"is_inter_review_by_submitter,omitempty"`
	IsReviewHomework         *bool          `json:"is_review_homework,omitempty"`
	NeedMakeUp               bool           `json:"need_make_up,omitempty"`
	NeedRemind               bool           `json:"need_remind,omitempty"`
	NonSubmitTimes           *bool          `json:"non_submit_times,omitempty"`
	IsResubmitOpen           bool           `json:"is_resubmit_open,omitempty"`

	// Status flags
	IsStarted                               bool `json:"is_started,omitempty"`
	IsClosed                                bool `json:"is_closed,omitempty"`
	IsInProgress                            bool `json:"is_in_progress,omitempty"`
	IsOpenedCatalog                         bool `json:"is_opened_catalog,omitempty"`
	IsAnnounceReferenceAnswerTimePassed     bool `json:"is_announce_reference_answer_time_passed,omitempty"`
	IsAnnounceScoreTimePassed               bool `json:"is_announce_score_time_passed,omitempty"`
	IsAnnounceAnswerTimePassed              bool `json:"is_announce_answer_time_passed,omitempty"`

	// Counts
	ForumCount          int `json:"forum_count,omitempty"`
	QuestionCount       int `json:"question_count,omitempty"`
	LateSubmissionCount int `json:"late_submission_count,omitempty"`
	UserSubmitCount     int `json:"user_submit_count,omitempty"`
	SubmissionsCount    int `json:"submissions_count,omitempty"`
	SubjectsCount       int `json:"subjects_count,omitempty"`

	// Knowledge
	KnowledgeNodeIDs       []int           `json:"knowledge_node_ids,omitempty"`
	KnowledgeNodeReference []interface{}   `json:"knowledge_node_reference,omitempty"`
	Prerequisites          []*Prerequisite `json:"prerequisites,omitempty"`

	// Meeting
	SubMeetings                   interface{} `json:"sub_meetings,omitempty"`
	TencentMeetingType            *string     `json:"tencent_meeting_type,omitempty"`
	InteractionActivityAttributes interface{} `json:"interaction_activity_attributes,omitempty"`

	// Forum-specific
	TopicCategoryID int `json:"topic_category_id,omitempty"`

	// Quiz settings
	IsPracticeMode         bool  `json:"is_practice_mode,omitempty"`
	IsQuizControlBySubject bool  `json:"is_quiz_control_by_subject,omitempty"`
	IsQuizPublic           bool  `json:"is_quiz_public,omitempty"`
	IsTimed                *bool `json:"is_timed,omitempty"`
	LimitTime              *int  `json:"limit_time,omitempty"`
	Duration               *int  `json:"duration,omitempty"`
	ExamPaperTemplateID    int   `json:"exam_paper_template_id,omitempty"`

	// Anti-cheat
	CheckSubmitIPConsistency   bool    `json:"check_submit_ip_consistency,omitempty"`
	IsIPConstrained            bool    `json:"is_ip_constrained,omitempty"`
	LimitedIP                  *string `json:"limited_ip,omitempty"`
	DisableCopyPaste           bool    `json:"disable_copy_paste,omitempty"`
	DisableDevtool             bool    `json:"disable_devtool,omitempty"`
	DisableRightClick          bool    `json:"disable_right_click,omitempty"`
	EnableAntiCheat            bool    `json:"enable_anti_cheat,omitempty"`
	EnableInvigilation         bool    `json:"enable_invigilation,omitempty"`
	IsFullscreenMode           bool    `json:"is_fullscreen_mode,omitempty"`
	IsLeavingWindowConstrained bool    `json:"is_leaving_window_constrained,omitempty"`
	IsLeavingWindowTimeout     bool    `json:"is_leaving_window_timeout,omitempty"`
	LeavingWindowLimit         *int    `json:"leaving_window_limit,omitempty"`
	LeavingWindowTimeout       *int    `json:"leaving_window_timeout,omitempty"`
	LimitAnswerOnSingleClient  bool    `json:"limit_answer_on_signle_client,omitempty"`

	// Subject rules
	SubjectsRule         *SubjectsRule `json:"subjects_rule,omitempty"`
	DefaultOptionsLayout string        `json:"default_options_layout,omitempty"`

	// Uploads
	Uploads []*Upload `json:"uploads,omitempty"`
}

// SubjectsRule represents quiz subject display/shuffle rules.
type SubjectsRule struct {
	SelectSubjectsRandomly  bool `json:"select_subjects_randomly,omitempty"`
	ShuffleOptionsRandomly  bool `json:"shuffle_options_randomly,omitempty"`
	ShuffleSubjectsRandomly bool `json:"shuffle_subjects_randomly,omitempty"`
	SubSubjectsRandomly     bool `json:"sub_subjects_randomly,omitempty"`
}

// ActivityRead represents activity read/completion status.
type ActivityRead struct {
	ID            string      `json:"id"`
	ActivityID    int         `json:"activity_id"`
	ActivityType  string      `json:"activity_type"`
	Completeness  string      `json:"completeness"`
	CreatedByID   int         `json:"created_by_id"`
	CreatedForID  int         `json:"created_for_id"`
	Data          interface{} `json:"data,omitempty"`
	LastVisitedAt string      `json:"last_visited_at,omitempty"`
}

// Comment represents a comment on an activity.
type Comment struct {
	ID         int         `json:"id"`
	Content    string      `json:"content,omitempty"`
	CreatedAt  string      `json:"created_at,omitempty"`
	UpdatedAt  string      `json:"updated_at,omitempty"`
	User       *User       `json:"user,omitempty"`
	Uploads    []*Upload   `json:"uploads,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	LikeCount  int         `json:"like_count,omitempty"`
	IsLiked    bool        `json:"is_liked,omitempty"`
	ParentID   *int        `json:"parent_id,omitempty"`
	ReplyCount int         `json:"reply_count,omitempty"`
	ActivityID int         `json:"activity_id,omitempty"`
	Page       *int        `json:"page,omitempty"`
}

// CommentCount represents comment count by type.
type CommentCount struct {
	Forum    int `json:"forum"`
	Question int `json:"question"`
}

// KnowledgeNode represents a knowledge node in the course knowledge tree.
type KnowledgeNode struct {
	ID         int              `json:"id"`
	Name       string           `json:"name,omitempty"`
	ParentID   *int             `json:"parent_id,omitempty"`
	CourseID   int              `json:"course_id,omitempty"`
	Sort       int              `json:"sort,omitempty"`
	Level      int              `json:"level,omitempty"`
	Children   []*KnowledgeNode `json:"children,omitempty"`
	Activities []*Activity      `json:"activities,omitempty"`
}
