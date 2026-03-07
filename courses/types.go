package courses

import (
	"github.com/eWloYW8/zju-courses-go-sdk/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

type CourseAttributes struct {
	AudienceType          any     `json:"audience_type"`
	CopyStatus            *string `json:"copy_status,omitempty"`
	Data                  any     `json:"data,omitempty"`
	GraduateMethod        string  `json:"graduate_method,omitempty"`
	IsDuringPublishPeriod bool    `json:"is_during_publish_period,omitempty"`
	Published             bool    `json:"published,omitempty"`
	TeachingClassName     string  `json:"teaching_class_name,omitempty"`
	Tip                   *string `json:"tip,omitempty"`
	EducationType         int     `json:"education_type,omitempty"`
	PracticeHours         *int    `json:"practice_hours,omitempty"`
	StudentCount          int     `json:"student_count,omitempty"`
	TheoryHours           *int    `json:"theory_hours,omitempty"`
	TotalHours            *int    `json:"total_hours,omitempty"`
	ClassHours            *int    `json:"class_hours,omitempty"`
}

type Course struct {
	ID                        int                        `json:"id"`
	Name                      string                     `json:"name"`
	DisplayName               string                     `json:"display_name,omitempty"`
	SecondName                *string                    `json:"second_name,omitempty"`
	CourseCode                string                     `json:"course_code,omitempty"`
	CourseType                int                        `json:"course_type,omitempty"`
	Cover                     string                     `json:"cover,omitempty"`
	SmallCover                string                     `json:"small_cover,omitempty"`
	Credit                    *float64                   `json:"credit,omitempty"`
	CreditState               *model.CreditState         `json:"credit_state,omitempty"`
	AcademicYearID            int                        `json:"academic_year_id,omitempty"`
	AcademicYear              *model.AcademicYear        `json:"academic_year,omitempty"`
	SemesterID                int                        `json:"semester_id,omitempty"`
	Semester                  *model.Semester            `json:"semester,omitempty"`
	DepartmentID              int                        `json:"department_id,omitempty"`
	Department                *model.Department          `json:"department,omitempty"`
	Org                       *model.Org                 `json:"org,omitempty"`
	OrgID                     int                        `json:"org_id,omitempty"`
	CourseAttributes          *CourseAttributes          `json:"course_attributes,omitempty"`
	Instructors               []*activities.ActivityUser `json:"instructors,omitempty"`
	CreatedUser               *activities.ActivityUser   `json:"created_user,omitempty"`
	ImportedFrom              string                     `json:"imported_from,omitempty"`
	StartDate                 *string                    `json:"start_date,omitempty"`
	EndDate                   *string                    `json:"end_date,omitempty"`
	Grade                     *string                    `json:"grade,omitempty"`
	IsStarted                 bool                       `json:"is_started,omitempty"`
	IsClosed                  bool                       `json:"is_closed,omitempty"`
	IsInstructor              bool                       `json:"is_instructor,omitempty"`
	IsMute                    bool                       `json:"is_mute,omitempty"`
	IsTeamTeaching            bool                       `json:"is_team_teaching,omitempty"`
	IsDefaultCourseCover      bool                       `json:"is_default_course_cover,omitempty"`
	IsStudio                  bool                       `json:"is_studio,omitempty"`
	IsBlocked                 *bool                      `json:"is_blocked,omitempty"`
	IsBlueprintCourse         bool                       `json:"is_blueprint_course,omitempty"`
	IsBlueprintSubCourse      bool                       `json:"is_blueprint_sub_course,omitempty"`
	IsCombinedCourse          bool                       `json:"is_combined_course,omitempty"`
	IsPublic                  bool                       `json:"is_public,omitempty"`
	Archived                  bool                       `json:"archived,omitempty"`
	AllowClone                bool                       `json:"allow_clone,omitempty"`
	AuditStatus               string                     `json:"audit_status,omitempty"`
	AuditRemark               *string                    `json:"audit_remark,omitempty"`
	PublicScope               string                     `json:"public_scope,omitempty"`
	Compulsory                *bool                      `json:"compulsory,omitempty"`
	Klass                     *string                    `json:"klass,omitempty"`
	CanWithdrawCourse         bool                       `json:"can_withdraw_course,omitempty"`
	ClassroomSchedule         *string                    `json:"classroom_schedule,omitempty"`
	StudyCompleteness         *float64                   `json:"study_completeness,omitempty"`
	UserStickCourseRecord     any                        `json:"user_stick_course_record,omitempty"`
	SubjectCode               *string                    `json:"subject_code,omitempty"`
	AllowAdminUpdateBasicInfo bool                       `json:"allow_admin_update_basic_info,omitempty"`
	AllowUpdateBasicInfo      bool                       `json:"allow_update_basic_info,omitempty"`
	AllowedToInviteAssistant  bool                       `json:"allowed_to_invite_assistant,omitempty"`
	AllowedToInviteStudent    bool                       `json:"allowed_to_invite_student,omitempty"`
	AllowedToJoinCourse       bool                       `json:"allowed_to_join_course,omitempty"`
	StudentsCount             int                        `json:"students_count,omitempty"`
	SyllabusEnabled           bool                       `json:"syllabus_enabled,omitempty"`
	TeachingMode              string                     `json:"teaching_mode,omitempty"`
	LearningMode              string                     `json:"learning_mode,omitempty"`
	ScorePublished            bool                       `json:"score_published,omitempty"`
	HasAIAbility              bool                       `json:"has_ai_ability,omitempty"`
	DingtalkNotifyExam        bool                       `json:"dingtalk_notify_exam,omitempty"`
	DingtalkNotifyHomework    bool                       `json:"dingtalk_notify_homework,omitempty"`
	KnowledgeGraphPublishType string                     `json:"knowledge_graph_publish_type,omitempty"`
	ProblemGraphPublishType   string                     `json:"problem_graph_publish_type,omitempty"`
	CreatedAt                 string                     `json:"created_at,omitempty"`
	UpdatedAt                 string                     `json:"updated_at,omitempty"`
	Modules                   []*Module                  `json:"modules,omitempty"`
	Enrollments               []*Enrollment              `json:"enrollments,omitempty"`
	Description               *string                    `json:"description,omitempty"`
	AccessCode                *string                    `json:"access_code,omitempty"`
	EnableForumEmailNotify    bool                       `json:"enable_forum_email_notify,omitempty"`
	Locale                    string                     `json:"locale,omitempty"`
}

type Module struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	Sort         int                    `json:"sort"`
	CourseID     int                    `json:"course_id,omitempty"`
	IsHidden     int                    `json:"is_hidden,omitempty"`
	LessonTimeID int                    `json:"lesson_time_id,omitempty"`
	StickyTime   *string                `json:"sticky_time,omitempty"`
	ImportedFrom *string                `json:"imported_from,omitempty"`
	CreatedAt    string                 `json:"created_at,omitempty"`
	UpdatedAt    string                 `json:"updated_at,omitempty"`
	Syllabuses   []*Syllabus            `json:"syllabuses,omitempty"`
	Activities   []*activities.Activity `json:"activities,omitempty"`
}

type Syllabus struct {
	ID         int                    `json:"id"`
	Title      string                 `json:"title"`
	Sort       int                    `json:"sort"`
	ModuleID   int                    `json:"module_id,omitempty"`
	CourseID   int                    `json:"course_id,omitempty"`
	CreatedAt  string                 `json:"created_at,omitempty"`
	UpdatedAt  string                 `json:"updated_at,omitempty"`
	Activities []*activities.Activity `json:"activities,omitempty"`
}

type EnrollmentDetail = activities.ActivityUser

type Enrollment struct {
	ID             int                      `json:"id"`
	CourseID       int                      `json:"course_id,omitempty"`
	UserID         int                      `json:"user_id,omitempty"`
	Roles          []string                 `json:"roles,omitempty"`
	ImportedFrom   *string                  `json:"imported_from,omitempty"`
	Aliases        []*string                `json:"aliases,omitempty"`
	SeatNumber     string                   `json:"seat_number,omitempty"`
	RetakeStatus   bool                     `json:"retake_status,omitempty"`
	MoocVideoScore *float64                 `json:"mooc_video_score,omitempty"`
	Data           map[string]any           `json:"data,omitempty"`
	CreatedAt      string                   `json:"created_at,omitempty"`
	UpdatedAt      string                   `json:"updated_at,omitempty"`
	User           *activities.ActivityUser `json:"user,omitempty"`
}

type NavSetting struct {
	Type        string `json:"type"`
	Parent      string `json:"parent,omitempty"`
	Disabled    bool   `json:"disabled"`
	CanEnabled  bool   `json:"can_enabled,omitempty"`
	CanDisabled bool   `json:"can_disabled"`
}

type ActivityPublishSetting struct {
	Exam     string `json:"exam"`
	Forum    string `json:"forum"`
	Homework string `json:"homework"`
	Others   string `json:"others"`
}

type CompletedResult struct {
	Completed       map[string][]int `json:"completed"`
	TotalActivities int              `json:"total_activities"`
	TotalCompleted  int              `json:"total_completed"`
}

type LastActivity struct {
	ID                       int             `json:"id"`
	ActivityEndTime          string          `json:"activity_end_time,omitempty"`
	ActivityStartTime        string          `json:"activity_start_time,omitempty"`
	ActivityType             string          `json:"activity_type,omitempty"`
	AssignGroupIDs           []int           `json:"assign_group_ids,omitempty"`
	AssignStudentIDs         []int           `json:"assign_student_ids,omitempty"`
	AssignTargets            map[string]any  `json:"assign_targets,omitempty"`
	CompletionCriterionKey   string          `json:"completion_criterion_key,omitempty"`
	CompletionCriterionValue any             `json:"completion_criterion_value,omitempty"`
	Data                     map[string]any  `json:"data,omitempty"`
	EndTime                  *string         `json:"end_time,omitempty"`
	GroupSetID               *int            `json:"group_set_id,omitempty"`
	IsAssignedToAll          bool            `json:"is_assigned_to_all,omitempty"`
	IsInProgress             bool            `json:"is_in_progress,omitempty"`
	ModuleID                 int             `json:"module_id,omitempty"`
	Published                bool            `json:"published,omitempty"`
	SubmitByGroup            bool            `json:"submit_by_group,omitempty"`
	SyllabusID               any             `json:"syllabus_id,omitempty"`
	TeachingUnitID           int             `json:"teaching_unit_id,omitempty"`
	Title                    string          `json:"title,omitempty"`
	Type                     string          `json:"type,omitempty"`
	Uploads                  []*model.Upload `json:"uploads,omitempty"`
}

type CompletenessResponse struct {
	CompletedResult            *CompletedResult `json:"completed_result"`
	LastActivity               *LastActivity    `json:"last_activity"`
	LastUpdateCompletenessTime string           `json:"last_update_completeness_time,omitempty"`
	StudyCompleteness          float64          `json:"study_completeness,omitempty"`
}

type Outline struct {
	ID             int             `json:"id"`
	CourseID       int             `json:"course_id,omitempty"`
	CommentChinese *OutlineField   `json:"comment_chinese,omitempty"`
	CommonFields   []*OutlineField `json:"common_fields,omitempty"`
	CustomFields   []*OutlineField `json:"custom_fields,omitempty"`
	EndDate        *string         `json:"end_date,omitempty"`
	ExternalURL    *string         `json:"external_url,omitempty"`
	IsClosed       bool            `json:"is_closed,omitempty"`
	IsImported     bool            `json:"is_imported,omitempty"`
	Status         string          `json:"status,omitempty"`
	CreatedAt      string          `json:"created_at,omitempty"`
	UpdatedAt      string          `json:"updated_at,omitempty"`
}

type OutlineField struct {
	ID          int             `json:"id"`
	Key         string          `json:"key,omitempty"`
	Title       string          `json:"title,omitempty"`
	Description string          `json:"description,omitempty"`
	Uploads     []*model.Upload `json:"uploads,omitempty"`
}

type HomeworkActivity struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
}

type HomeworkScore struct {
	ActivityID        int      `json:"activity_id"`
	StudentID         int      `json:"student_id,omitempty"`
	Score             *float64 `json:"score,omitempty"`
	FinalScore        *float64 `json:"final_score,omitempty"`
	InstructorComment *string  `json:"instructor_comment,omitempty"`
	InterScore        *float64 `json:"inter_score,omitempty"`
	IntraScore        *float64 `json:"intra_score,omitempty"`
}

type HomeworkSubmissionStatus struct {
	ID                        int      `json:"id"`
	HomeworkType              string   `json:"homework_type,omitempty"`
	Score                     *float64 `json:"score,omitempty"`
	Status                    string   `json:"status,omitempty"`
	StatusCode                string   `json:"status_code,omitempty"`
	IsAnnounceScoreTimePassed bool     `json:"is_announce_score_time_passed,omitempty"`
}

type Exam struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
}

type ExamScore struct {
	ExamID int      `json:"exam_id"`
	Score  *float64 `json:"score,omitempty"`
}

type Classroom struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
}

type TopicCategory struct {
	ID           int                  `json:"id"`
	Title        string               `json:"title,omitempty"`
	ReferrerType string               `json:"referrer_type,omitempty"`
	Activity     *activities.Activity `json:"activity,omitempty"`
}

type Interaction struct {
	ID         int            `json:"id"`
	Title      string         `json:"title,omitempty"`
	Type       string         `json:"type,omitempty"`
	Status     string         `json:"status,omitempty"`
	CourseID   int            `json:"course_id,omitempty"`
	ActivityID int            `json:"activity_id,omitempty"`
	CreatedAt  string         `json:"created_at,omitempty"`
	Data       map[string]any `json:"data,omitempty"`
}

type LiveRecord struct {
	ID         int    `json:"id"`
	Title      string `json:"title,omitempty"`
	URL        string `json:"url,omitempty"`
	Duration   int    `json:"duration,omitempty"`
	StartTime  string `json:"start_time,omitempty"`
	EndTime    string `json:"end_time,omitempty"`
	Status     string `json:"status,omitempty"`
	CourseID   int    `json:"course_id,omitempty"`
	ActivityID int    `json:"activity_id,omitempty"`
}

type Rollcall struct {
	ID          int                      `json:"id"`
	CourseID    int                      `json:"course_id,omitempty"`
	ModuleID    int                      `json:"module_id,omitempty"`
	Status      string                   `json:"status,omitempty"`
	Type        string                   `json:"type,omitempty"`
	StartTime   string                   `json:"start_time,omitempty"`
	EndTime     string                   `json:"end_time,omitempty"`
	Duration    int                      `json:"duration,omitempty"`
	CreatedAt   string                   `json:"created_at,omitempty"`
	UpdatedAt   string                   `json:"updated_at,omitempty"`
	CreatedBy   *activities.ActivityUser `json:"created_by,omitempty"`
	TotalCount  int                      `json:"total_count,omitempty"`
	SignedCount int                      `json:"signed_count,omitempty"`
}

type KnowledgeCapture struct {
	ID        int     `json:"id"`
	Code      string  `json:"code,omitempty"`
	Title     string  `json:"title,omitempty"`
	Name      string  `json:"name,omitempty"`
	CourseID  *int    `json:"course_id,omitempty"`
	UploadID  *int    `json:"upload_id,omitempty"`
	URL       string  `json:"url,omitempty"`
	Cover     *string `json:"cover,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
}

type KnowledgeNodeRecommendedResourceReference struct {
	ID     int           `json:"id"`
	Name   string        `json:"name,omitempty"`
	Upload *model.Upload `json:"upload,omitempty"`
}

type KnowledgeNodeCompletenessItem struct {
	ActivityID   int `json:"activity_id,omitempty"`
	ActivityType int `json:"activity_type,omitempty"`
	Completeness int `json:"completeness,omitempty"`
}

type KnowledgeNodeSummaryNode struct {
	ID                      int                         `json:"id"`
	ParentID                *int                        `json:"parent_id,omitempty"`
	Name                    string                      `json:"name,omitempty"`
	AverageMasteryRate      string                      `json:"average_mastery_rate,omitempty"`
	AverageCompletenessRate string                      `json:"average_completeness_rate,omitempty"`
	CognitiveDimension      string                      `json:"cognitive_dimension,omitempty"`
	Children                []*KnowledgeNodeSummaryNode `json:"children,omitempty"`
}

type KnowledgeNodeStatisticsSummary struct {
	NodeCount                    int                         `json:"node_count,omitempty"`
	NodeWithReferenceCount       int                         `json:"node_with_reference_count,omitempty"`
	AverageMasteryRate           string                      `json:"average_mastery_rate,omitempty"`
	AverageCompletenessRate      string                      `json:"average_completeness_rate,omitempty"`
	RelationCount                int                         `json:"relation_count,omitempty"`
	CompletenessRateDistribution any                         `json:"completeness_rate_distribution,omitempty"`
	MasteryRateDistribution      any                         `json:"mastery_rate_distribution,omitempty"`
	Nodes                        []*KnowledgeNodeSummaryNode `json:"nodes,omitempty"`
}

type KnowledgeNodeOverview struct {
	ID                           int    `json:"id"`
	Name                         string `json:"name,omitempty"`
	Description                  string `json:"description,omitempty"`
	AverageMasteryRate           string `json:"avg_mastery_rate,omitempty"`
	AverageCompletenessRate      string `json:"avg_completeness_rate,omitempty"`
	UploadCount                  int    `json:"upload_count,omitempty"`
	ActivityCount                int    `json:"activity_count,omitempty"`
	CompletenessRateDistribution any    `json:"completeness_rate_distribution,omitempty"`
	MasteryRateDistribution      any    `json:"mastery_rate_distribution,omitempty"`
}

type KnowledgeNodeStudentDimensionItem struct {
	StudentID        int    `json:"student_id,omitempty"`
	Name             string `json:"name,omitempty"`
	UserNo           string `json:"user_no,omitempty"`
	MasteryRate      string `json:"mastery_rate,omitempty"`
	CompletenessRate string `json:"completeness_rate,omitempty"`
	ImportedFrom     string `json:"imported_from,omitempty"`
}

type KnowledgeNodeStudentDetail struct {
	StudentID            int                         `json:"student_id,omitempty"`
	Name                 string                      `json:"name,omitempty"`
	UserNo               string                      `json:"user_no,omitempty"`
	MasteryRate          string                      `json:"mastery_rate,omitempty"`
	CompletenessRate     string                      `json:"completeness_rate,omitempty"`
	ImportedFrom         string                      `json:"imported_from,omitempty"`
	MasteryRateRank      string                      `json:"mastery_rate_rank,omitempty"`
	CompletenessRateRank string                      `json:"completeness_rate_rank,omitempty"`
	Nodes                []*KnowledgeNodeSummaryNode `json:"nodes,omitempty"`
}

type KnowledgeNodeResourceDetail struct {
	ID                int    `json:"id"`
	Name              string `json:"name,omitempty"`
	Type              string `json:"type,omitempty"`
	AllowDownload     bool   `json:"allow_download,omitempty"`
	Extension         string `json:"extension,omitempty"`
	CompletionStudent string `json:"completion_student,omitempty"`
	Completeness      string `json:"completeness,omitempty"`
	Visits            int    `json:"visits,omitempty"`
	Source            string `json:"source,omitempty"`
	Status            string `json:"status,omitempty"`
}

type KnowledgeNodeActivityDetail struct {
	ID                int    `json:"id"`
	Name              string `json:"name,omitempty"`
	Type              string `json:"type,omitempty"`
	CompletionStudent string `json:"completion_student,omitempty"`
	Completeness      string `json:"completeness,omitempty"`
	MasteryRate       string `json:"mastery_rate,omitempty"`
}

type KnowledgeNodeStudentResourceStat struct {
	ID     int    `json:"id"`
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Viewed bool   `json:"viewed,omitempty"`
	Visits int    `json:"visits,omitempty"`
	Status string `json:"status,omitempty"`
	Source string `json:"source,omitempty"`
}

type KnowledgeNodeStudentActivityStat struct {
	ID           int    `json:"id"`
	Title        string `json:"title,omitempty"`
	Type         string `json:"type,omitempty"`
	Completeness string `json:"completeness,omitempty"`
	Score        any    `json:"score,omitempty"`
}

type KnowledgeNodeStudentOverallStatistics struct {
	OverallCompletenessRate string `json:"overall_completeness_rate,omitempty"`
	OverallMasteryRate      string `json:"overall_mastery_rate,omitempty"`
}

type KnowledgeGraphUser struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

type KnowledgeGraphKFSCourse struct {
	ID              int                   `json:"id"`
	Name            string                `json:"name,omitempty"`
	TopicCount      int                   `json:"topic_count,omitempty"`
	DependencyCount int                   `json:"dependency_count,omitempty"`
	FacetCount      int                   `json:"facet_count,omitempty"`
	FragmentCount   int                   `json:"fragment_count,omitempty"`
	Builders        []*KnowledgeGraphUser `json:"builders,omitempty"`
	Principals      []*KnowledgeGraphUser `json:"principals,omitempty"`
	LatestVersionID int                   `json:"latest_version_id,omitempty"`
}

type KnowledgeGraphKFSSubject struct {
	Name     string                     `json:"name,omitempty"`
	ID       int                        `json:"id"`
	DataType string                     `json:"data_type,omitempty"`
	Courses  []*KnowledgeGraphKFSCourse `json:"courses,omitempty"`
}

type KnowledgeGraphKFSImportInfo struct {
	LastImportTime  int64  `json:"last_import_time,omitempty"`
	ImportCourseIDs []int  `json:"import_course_ids,omitempty"`
	Server          string `json:"server,omitempty"`
}

type KnowledgeGraphForestVersionStats struct {
	TopicCount      int `json:"topic_count,omitempty"`
	DependencyCount int `json:"dependency_count,omitempty"`
	FacetCount      int `json:"facet_count,omitempty"`
	FragmentCount   int `json:"fragment_count,omitempty"`
}

type KnowledgeGraphForestVersionStatsItem struct {
	ID int `json:"id"`
	KnowledgeGraphForestVersionStats
}

type KnowledgeGraphPublishedForestVersion struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	CourseID  int    `json:"course_id,omitempty"`
	Published bool   `json:"published,omitempty"`
}

type KnowledgeGraphDiff struct {
	Action   string `json:"action,omitempty"`
	Type     string `json:"type,omitempty"`
	NodeID   int    `json:"node_id,omitempty"`
	NodeName string `json:"node_name,omitempty"`
	Before   any    `json:"before,omitempty"`
	After    any    `json:"after,omitempty"`
	Data     any    `json:"data,omitempty"`
}

type KnowledgeGraphSimilarity struct {
	ID               int     `json:"id"`
	Name             string  `json:"name,omitempty"`
	Similarity       float64 `json:"similarity,omitempty"`
	SimilarityFormat string  `json:"similarity_format,omitempty"`
	Checked          bool    `json:"checked,omitempty"`
}

type KnowledgeNodeWeekStat struct {
	ID               int     `json:"id"`
	CompletenessRate float64 `json:"avg_completeness_rate,omitempty"`
	MasteryRate      float64 `json:"avg_mastery_rate,omitempty"`
	StatDate         string  `json:"stat_date,omitempty"`
	Week             int     `json:"week,omitempty"`
}

type KnowledgeGraphSnapshotRelation struct {
	ID         int    `json:"id"`
	Source     int    `json:"source,omitempty"`
	Target     int    `json:"target,omitempty"`
	IsDirected bool   `json:"is_directed,omitempty"`
	Color      string `json:"color,omitempty"`
}

type KnowledgeGraphSnapshot struct {
	Tree         *model.KnowledgeNode              `json:"tree,omitempty"`
	Relations    []*KnowledgeGraphSnapshotRelation `json:"relations,omitempty"`
	Completeness []float64                         `json:"completeness,omitempty"`
	Mastery      []float64                         `json:"mastery,omitempty"`
}

type KnowledgeNodeRelation struct {
	ID         int  `json:"id"`
	Source     int  `json:"source,omitempty"`
	Target     int  `json:"target,omitempty"`
	IsDirected bool `json:"is_directed,omitempty"`
}

type CoursewareActivity struct {
	ID             int    `json:"id"`
	Title          string `json:"title,omitempty"`
	Type           string `json:"type,omitempty"`
	PublishType    string `json:"publish_type,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
	EndTime        string `json:"end_time,omitempty"`
	IsClosed       bool   `json:"is_closed,omitempty"`
	Score          any    `json:"score,omitempty"`
	Completeness   any    `json:"completeness,omitempty"`
	ScorePublished bool   `json:"score_published,omitempty"`
	Published      bool   `json:"published,omitempty"`
	IsLocked       bool   `json:"is_locked,omitempty"`
	Prerequisites  any    `json:"prerequisites,omitempty"`
	ModuleID       int    `json:"module_id,omitempty"`
	SyllabusID     int    `json:"syllabus_id,omitempty"`
	ModuleName     string `json:"module_name,omitempty"`
	SyllabusName   string `json:"syllabus_name,omitempty"`
}

type KnowledgeNodeStudentReferenceStat struct {
	NodeName         string `json:"node_name,omitempty"`
	NodeDesc         string `json:"node_desc,omitempty"`
	Name             string `json:"name,omitempty"`
	UserNo           string `json:"user_no,omitempty"`
	MasteryRank      string `json:"mastery_rank,omitempty"`
	MasteryRate      string `json:"mastery_rate,omitempty"`
	CompletenessRank string `json:"completeness_rank,omitempty"`
	CompletenessRate string `json:"completeness_rate,omitempty"`
	ResourceCount    int    `json:"resource_count,omitempty"`
	ActivityCount    int    `json:"activity_count,omitempty"`
}

type KnowledgeNodeReferTypeStats struct {
	NodeCount     int `json:"node_count,omitempty"`
	ResourceCount int `json:"resource_ref_count,omitempty"`
	ActivityCount int `json:"activity_ref_count,omitempty"`
}

type TeachingObjective struct {
	ID         int    `json:"id"`
	Content    string `json:"content,omitempty"`
	ReferCount int    `json:"refer_count,omitempty"`
}

type KnowledgeFragmentDetail struct {
	Name        string `json:"name,omitempty"`
	Book        string `json:"book,omitempty"`
	ISBN        string `json:"isbn,omitempty"`
	Page        int    `json:"page,omitempty"`
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

type KnowledgeFragment struct {
	ID      int                      `json:"id"`
	FacetID int                      `json:"facet_id,omitempty"`
	Content string                   `json:"content,omitempty"`
	Source  string                   `json:"source,omitempty"`
	Detail  *KnowledgeFragmentDetail `json:"detail,omitempty"`
}

type KnowledgeFacet struct {
	ID        int                  `json:"id"`
	Name      string               `json:"name,omitempty"`
	ParentID  int                  `json:"parent_id,omitempty"`
	TopicID   int                  `json:"topic_id,omitempty"`
	Children  []*KnowledgeFacet    `json:"children,omitempty"`
	Fragments []*KnowledgeFragment `json:"fragments,omitempty"`
}

type KnowledgeReferenceData struct {
	MaterialUploads []model.Upload `json:"material_uploads,omitempty"`
	MediaFragments  []any          `json:"media_fragments,omitempty"`
}

type ActivityKnowledgeReference struct {
	ID              int                     `json:"id"`
	CourseID        int                     `json:"course_id,omitempty"`
	KnowledgeNodeID int                     `json:"knowledge_node_id,omitempty"`
	ReferID         int                     `json:"refer_id,omitempty"`
	ReferType       string                  `json:"refer_type,omitempty"`
	Data            *KnowledgeReferenceData `json:"data,omitempty"`
}

type BlueprintSubItemsResponse struct {
	Items []BlueprintSubItem `json:"items"`
}

type BlueprintSubItem struct {
	CourseID int `json:"course_id"`
	Count    int `json:"count"`
}
