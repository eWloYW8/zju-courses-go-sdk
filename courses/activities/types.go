package activities

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type InterScoreMap struct {
	ID        int     `json:"id"`
	StartTime *string `json:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty"`
	IsStarted *bool   `json:"is_started,omitempty"`
	IsClosed  *bool   `json:"is_closed,omitempty"`
	PiecesCnt int     `json:"pieces_cnt,omitempty"`
}

type IntraScoreMap struct {
	ID        int     `json:"id"`
	StartTime *string `json:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty"`
	IsClosed  *bool   `json:"is_closed,omitempty"`
}

type Rubric struct {
	ID             int                `json:"id"`
	Name           *string            `json:"name,omitempty"`
	Conditions     []*RubricCondition `json:"conditions,omitempty"`
	CreatedAt      *string            `json:"created_at,omitempty"`
	UpdatedAt      *string            `json:"updated_at,omitempty"`
	CreatedBy      *ActivityUser      `json:"created_by,omitempty"`
	EngageNumber   int                `json:"engage_number,omitempty"`
	GroupID        *int               `json:"group_id,omitempty"`
	GroupName      *string            `json:"group_name,omitempty"`
	IsSharedRubric bool               `json:"is_shared_rubric,omitempty"`
	OrgID          int                `json:"org_id,omitempty"`
}

type RubricInstance struct {
	ID         int                `json:"id"`
	Name       *string            `json:"name,omitempty"`
	Conditions []*RubricCondition `json:"conditions,omitempty"`
	Rubric     *Rubric            `json:"rubric,omitempty"`
	RubricID   int                `json:"rubric_id,omitempty"`
}

type RubricCondition struct {
	ID          int            `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Weight      float64        `json:"weight,omitempty"`
	Levels      []*RubricLevel `json:"levels,omitempty"`
}

type RubricLevel struct {
	ID          int     `json:"id,omitempty"`
	Description string  `json:"description,omitempty"`
	Score       float64 `json:"score,omitempty"`
	Name        string  `json:"name,omitempty"`
}

type SubjectsRule struct {
	SelectSubjectsRandomly  bool `json:"select_subjects_randomly,omitempty"`
	ShuffleOptionsRandomly  bool `json:"shuffle_options_randomly,omitempty"`
	ShuffleSubjectsRandomly bool `json:"shuffle_subjects_randomly,omitempty"`
	SubSubjectsRandomly     bool `json:"sub_subjects_randomly,omitempty"`
}

type CompletionCriterion struct {
	ID                      int    `json:"id"`
	CompletionCriterionType string `json:"completion_criterion_type,omitempty"`
	Name                    string `json:"name,omitempty"`
	Title                   string `json:"title,omitempty"`
	Description             string `json:"description,omitempty"`
	Value                   any    `json:"value,omitempty"`
	IsDefault               bool   `json:"is_default,omitempty"`
}

type Prerequisite struct {
	ID           int    `json:"id"`
	ActivityID   int    `json:"activity_id,omitempty"`
	ActivityType string `json:"activity_type,omitempty"`
	Title        string `json:"title,omitempty"`
	Completed    bool   `json:"completed,omitempty"`
}

type ActivityUser struct {
	ID             int                   `json:"id"`
	Name           *string               `json:"name,omitempty"`
	Email          *string               `json:"email,omitempty"`
	Nickname       *string               `json:"nickname,omitempty"`
	UserNo         string                `json:"user_no,omitempty"`
	Comment        *string               `json:"comment,omitempty"`
	Grade          *model.Grade          `json:"grade,omitempty"`
	Klass          *model.Class          `json:"klass,omitempty"`
	AvatarSmallURL string                `json:"avatar_small_url,omitempty"`
	AvatarBigURL   string                `json:"avatar_big_url,omitempty"`
	PortfolioURL   string                `json:"portfolio_url,omitempty"`
	Department     *model.Department     `json:"department,omitempty"`
	LearningCenter *model.LearningCenter `json:"learning_center,omitempty"`
	Org            *model.OrgDetail      `json:"org,omitempty"`
	Program        *model.Program        `json:"program,omitempty"`
	UserAttributes *model.UserAttributes `json:"user_attributes,omitempty"`
}

type Activity struct {
	ID                                  int             `json:"id"`
	Title                               string          `json:"title"`
	Type                                string          `json:"type"`
	CourseID                            int             `json:"course_id,omitempty"`
	ModuleID                            int             `json:"module_id,omitempty"`
	SyllabusID                          int             `json:"syllabus_id,omitempty"`
	Sort                                int             `json:"sort,omitempty"`
	CreatedAt                           string          `json:"created_at,omitempty"`
	UpdatedAt                           string          `json:"updated_at,omitempty"`
	StartTime                           *string         `json:"start_time,omitempty"`
	EndTime                             *string         `json:"end_time,omitempty"`
	Description                         string          `json:"description,omitempty"`
	Data                                map[string]any  `json:"data,omitempty"`
	EnableEdit                          bool            `json:"enable_edit,omitempty"`
	Published                           bool            `json:"published,omitempty"`
	PublishTime                         *string         `json:"publish_time,omitempty"`
	UniqueKey                           string          `json:"unique_key,omitempty"`
	TeachingModel                       string          `json:"teaching_model,omitempty"`
	TeachingUnitID                      int             `json:"teaching_unit_id,omitempty"`
	UsingPhase                          string          `json:"using_phase,omitempty"`
	Version                             *int            `json:"version,omitempty"`
	ImportedFrom                        *string         `json:"imported_from,omitempty"`
	ImportedTrackID                     *string         `json:"imported_track_id,omitempty"`
	ReferrerID                          int             `json:"referrer_id,omitempty"`
	ReferrerType                        string          `json:"referrer_type,omitempty"`
	GroupSetID                          int             `json:"group_set_id,omitempty"`
	GroupSetName                        *string         `json:"group_set_name,omitempty"`
	HasAssignGroup                      bool            `json:"has_assign_group,omitempty"`
	HasAssignStudent                    bool            `json:"has_assign_student,omitempty"`
	AssignGroupIDs                      []int           `json:"assign_group_ids,omitempty"`
	AssignStudentIDs                    []int           `json:"assign_student_ids,omitempty"`
	IsAssignedToAll                     bool            `json:"is_assigned_to_all,omitempty"`
	SubmitByGroup                       bool            `json:"submit_by_group,omitempty"`
	SubmitTimes                         *int            `json:"submit_times,omitempty"`
	AssignTargets                       map[string]any  `json:"assign_targets,omitempty"`
	CompletionCriterion                 string          `json:"completion_criterion,omitempty"`
	CompletionCriterionKey              string          `json:"completion_criterion_key,omitempty"`
	CompletionCriterionValue            any             `json:"completion_criterion_value,omitempty"`
	CanShowScore                        bool            `json:"can_show_score,omitempty"`
	AverageScore                        *float64        `json:"average_score,omitempty"`
	HighestScore                        *float64        `json:"highest_score,omitempty"`
	LowestScore                         *float64        `json:"lowest_score,omitempty"`
	HasScoreCount                       int             `json:"has_score_count,omitempty"`
	ScorePercentage                     *string         `json:"score_percentage,omitempty"`
	ScorePublished                      *bool           `json:"score_published,omitempty"`
	ScoreType                           *string         `json:"score_type,omitempty"`
	ScoreItemGroupID                    int             `json:"score_item_group_id,omitempty"`
	ScoreItemGroupName                  *string         `json:"score_item_group_name,omitempty"`
	ScoreItemScored                     *bool           `json:"score_item_scored,omitempty"`
	ScoreRule                           string          `json:"score_rule,omitempty"`
	TotalScore                          *float64        `json:"total_score,omitempty"`
	IsScorePublic                       bool            `json:"is_score_public,omitempty"`
	AnnounceAnswerStatus                string          `json:"announce_answer_status,omitempty"`
	AnnounceAnswerType                  string          `json:"announce_answer_type,omitempty"`
	AnnounceAnswerTime                  *string         `json:"announce_answer_time,omitempty"`
	AnnounceScoreStatus                 string          `json:"announce_score_status,omitempty"`
	AnnounceScoreTime                   *string         `json:"announce_score_time,omitempty"`
	IsAnswerAnnounced                   bool            `json:"is_answer_announced,omitempty"`
	AnnounceAnswerAndExplanation        *bool           `json:"announce_answer_and_explanation,omitempty"`
	RubricID                            int             `json:"rubric_id,omitempty"`
	RubricInstance                      *RubricInstance `json:"rubric_instance,omitempty"`
	RubricInstanceID                    int             `json:"rubric_instance_id,omitempty"`
	IntraRubricID                       int             `json:"intra_rubric_id,omitempty"`
	IntraRubricInstance                 *RubricInstance `json:"intra_rubric_instance,omitempty"`
	IntraRubricInstanceID               int             `json:"intra_rubric_instance_id,omitempty"`
	InterReviewNamed                    *bool           `json:"inter_review_named,omitempty"`
	InterScoreMap                       *InterScoreMap  `json:"inter_score_map,omitempty"`
	IntraScoreMap                       *IntraScoreMap  `json:"intra_score_map,omitempty"`
	IsInterReviewBySubmitter            *bool           `json:"is_inter_review_by_submitter,omitempty"`
	IsReviewHomework                    *bool           `json:"is_review_homework,omitempty"`
	NeedMakeUp                          bool            `json:"need_make_up,omitempty"`
	NeedRemind                          bool            `json:"need_remind,omitempty"`
	NonSubmitTimes                      *bool           `json:"non_submit_times,omitempty"`
	IsResubmitOpen                      bool            `json:"is_resubmit_open,omitempty"`
	IsStarted                           bool            `json:"is_started,omitempty"`
	IsClosed                            bool            `json:"is_closed,omitempty"`
	IsInProgress                        bool            `json:"is_in_progress,omitempty"`
	IsOpenedCatalog                     bool            `json:"is_opened_catalog,omitempty"`
	IsAnnounceReferenceAnswerTimePassed bool            `json:"is_announce_reference_answer_time_passed,omitempty"`
	IsAnnounceScoreTimePassed           bool            `json:"is_announce_score_time_passed,omitempty"`
	IsAnnounceAnswerTimePassed          bool            `json:"is_announce_answer_time_passed,omitempty"`
	ForumCount                          int             `json:"forum_count,omitempty"`
	QuestionCount                       int             `json:"question_count,omitempty"`
	LateSubmissionCount                 int             `json:"late_submission_count,omitempty"`
	UserSubmitCount                     int             `json:"user_submit_count,omitempty"`
	SubmissionsCount                    int             `json:"submissions_count,omitempty"`
	SubjectsCount                       int             `json:"subjects_count,omitempty"`
	KnowledgeNodeIDs                    []int           `json:"knowledge_node_ids,omitempty"`
	KnowledgeNodeReference              []any           `json:"knowledge_node_reference,omitempty"`
	Prerequisites                       []*Prerequisite `json:"prerequisites,omitempty"`
	SubMeetings                         any             `json:"sub_meetings,omitempty"`
	TencentMeetingType                  *string         `json:"tencent_meeting_type,omitempty"`
	InteractionActivityAttributes       any             `json:"interaction_activity_attributes,omitempty"`
	TopicCategoryID                     int             `json:"topic_category_id,omitempty"`
	IsPracticeMode                      bool            `json:"is_practice_mode,omitempty"`
	IsQuizControlBySubject              bool            `json:"is_quiz_control_by_subject,omitempty"`
	IsQuizPublic                        bool            `json:"is_quiz_public,omitempty"`
	IsTimed                             *bool           `json:"is_timed,omitempty"`
	LimitTime                           *int            `json:"limit_time,omitempty"`
	Duration                            *int            `json:"duration,omitempty"`
	ExamPaperTemplateID                 int             `json:"exam_paper_template_id,omitempty"`
	CheckSubmitIPConsistency            bool            `json:"check_submit_ip_consistency,omitempty"`
	IsIPConstrained                     bool            `json:"is_ip_constrained,omitempty"`
	LimitedIP                           *string         `json:"limited_ip,omitempty"`
	DisableCopyPaste                    bool            `json:"disable_copy_paste,omitempty"`
	DisableDevtool                      bool            `json:"disable_devtool,omitempty"`
	DisableRightClick                   bool            `json:"disable_right_click,omitempty"`
	EnableAntiCheat                     bool            `json:"enable_anti_cheat,omitempty"`
	EnableInvigilation                  bool            `json:"enable_invigilation,omitempty"`
	IsFullscreenMode                    bool            `json:"is_fullscreen_mode,omitempty"`
	IsLeavingWindowConstrained          bool            `json:"is_leaving_window_constrained,omitempty"`
	IsLeavingWindowTimeout              bool            `json:"is_leaving_window_timeout,omitempty"`
	LeavingWindowLimit                  *int            `json:"leaving_window_limit,omitempty"`
	LeavingWindowTimeout                *int            `json:"leaving_window_timeout,omitempty"`
	LimitAnswerOnSingleClient           bool            `json:"limit_answer_on_signle_client,omitempty"`
	SubjectsRule                        *SubjectsRule   `json:"subjects_rule,omitempty"`
	DefaultOptionsLayout                string          `json:"default_options_layout,omitempty"`
	Uploads                             []*model.Upload `json:"uploads,omitempty"`
}

type ActivityRead struct {
	ID            string         `json:"id"`
	ActivityID    int            `json:"activity_id"`
	ActivityType  string         `json:"activity_type"`
	Completeness  string         `json:"completeness"`
	CreatedByID   int            `json:"created_by_id"`
	CreatedForID  int            `json:"created_for_id"`
	Data          map[string]any `json:"data,omitempty"`
	LastVisitedAt string         `json:"last_visited_at,omitempty"`
}

type Comment struct {
	ID         int             `json:"id"`
	Content    string          `json:"content,omitempty"`
	CreatedAt  string          `json:"created_at,omitempty"`
	UpdatedAt  string          `json:"updated_at,omitempty"`
	User       *ActivityUser   `json:"user,omitempty"`
	Uploads    []*model.Upload `json:"uploads,omitempty"`
	Data       map[string]any  `json:"data,omitempty"`
	LikeCount  int             `json:"like_count,omitempty"`
	IsLiked    bool            `json:"is_liked,omitempty"`
	ParentID   *int            `json:"parent_id,omitempty"`
	ReplyCount int             `json:"reply_count,omitempty"`
	ActivityID int             `json:"activity_id,omitempty"`
	Page       *int            `json:"page,omitempty"`
}

type CommentCount struct {
	Forum    int `json:"forum"`
	Question int `json:"question"`
}

type IsLockedStatus struct {
	IsLocked      bool            `json:"is_locked"`
	Prerequisites []*Prerequisite `json:"prerequisites,omitempty"`
}
