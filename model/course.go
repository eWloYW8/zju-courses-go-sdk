package model

// CourseAttributes represents course-specific attributes.
type CourseAttributes struct {
	AudienceType          interface{} `json:"audience_type"`
	CopyStatus            *string     `json:"copy_status,omitempty"`
	Data                  interface{} `json:"data,omitempty"`
	GraduateMethod        string      `json:"graduate_method,omitempty"`
	IsDuringPublishPeriod bool        `json:"is_during_publish_period,omitempty"`
	Published             bool        `json:"published,omitempty"`
	TeachingClassName     string      `json:"teaching_class_name,omitempty"`
	Tip                   *string     `json:"tip,omitempty"`
	EducationType         int         `json:"education_type,omitempty"`
	PracticeHours         *int        `json:"practice_hours,omitempty"`
	StudentCount          int         `json:"student_count,omitempty"`
	TheoryHours           *int        `json:"theory_hours,omitempty"`
	TotalHours            *int        `json:"total_hours,omitempty"`
	ClassHours            *int        `json:"class_hours,omitempty"`
}

// Course represents a full course with all fields.
type Course struct {
	ID                        int              `json:"id"`
	Name                      string           `json:"name"`
	DisplayName               string           `json:"display_name,omitempty"`
	SecondName                *string          `json:"second_name,omitempty"`
	CourseCode                string           `json:"course_code,omitempty"`
	CourseType                int              `json:"course_type,omitempty"`
	Cover                     string           `json:"cover,omitempty"`
	SmallCover                string           `json:"small_cover,omitempty"`
	Credit                    *float64         `json:"credit,omitempty"`
	CreditState               *CreditState     `json:"credit_state,omitempty"`
	AcademicYearID            int              `json:"academic_year_id,omitempty"`
	AcademicYear              *AcademicYear    `json:"academic_year,omitempty"`
	SemesterID                int              `json:"semester_id,omitempty"`
	Semester                  *Semester        `json:"semester,omitempty"`
	DepartmentID              int              `json:"department_id,omitempty"`
	Department                *Department      `json:"department,omitempty"`
	Org                       *Org             `json:"org,omitempty"`
	OrgID                     int              `json:"org_id,omitempty"`
	CourseAttributes          *CourseAttributes `json:"course_attributes,omitempty"`
	Instructors               []*User          `json:"instructors,omitempty"`
	CreatedUser               *User            `json:"created_user,omitempty"`
	ImportedFrom              string           `json:"imported_from,omitempty"`
	StartDate                 *string          `json:"start_date,omitempty"`
	EndDate                   *string          `json:"end_date,omitempty"`
	Grade                     *string          `json:"grade,omitempty"`
	IsStarted                 bool             `json:"is_started,omitempty"`
	IsClosed                  bool             `json:"is_closed,omitempty"`
	IsInstructor              bool             `json:"is_instructor,omitempty"`
	IsMute                    bool             `json:"is_mute,omitempty"`
	IsTeamTeaching            bool             `json:"is_team_teaching,omitempty"`
	IsDefaultCourseCover      bool             `json:"is_default_course_cover,omitempty"`
	IsStudio                  bool             `json:"is_studio,omitempty"`
	IsBlocked                 *bool            `json:"is_blocked,omitempty"`
	IsBlueprintCourse         bool             `json:"is_blueprint_course,omitempty"`
	IsBlueprintSubCourse      bool             `json:"is_blueprint_sub_course,omitempty"`
	IsCombinedCourse          bool             `json:"is_combined_course,omitempty"`
	IsPublic                  bool             `json:"is_public,omitempty"`
	Archived                  bool             `json:"archived,omitempty"`
	AllowClone                bool             `json:"allow_clone,omitempty"`
	AuditStatus               string           `json:"audit_status,omitempty"`
	AuditRemark               *string          `json:"audit_remark,omitempty"`
	PublicScope               string           `json:"public_scope,omitempty"`
	Compulsory                *bool            `json:"compulsory,omitempty"`
	Klass                     *string          `json:"klass,omitempty"`
	CanWithdrawCourse         bool             `json:"can_withdraw_course,omitempty"`
	ClassroomSchedule         *string          `json:"classroom_schedule,omitempty"`
	StudyCompleteness         *float64         `json:"study_completeness,omitempty"`
	UserStickCourseRecord     interface{}      `json:"user_stick_course_record,omitempty"`
	SubjectCode               *string          `json:"subject_code,omitempty"`
	AllowAdminUpdateBasicInfo bool             `json:"allow_admin_update_basic_info,omitempty"`
	AllowUpdateBasicInfo      bool             `json:"allow_update_basic_info,omitempty"`
	AllowedToInviteAssistant  bool             `json:"allowed_to_invite_assistant,omitempty"`
	AllowedToInviteStudent    bool             `json:"allowed_to_invite_student,omitempty"`
	AllowedToJoinCourse       bool             `json:"allowed_to_join_course,omitempty"`
	StudentsCount             int              `json:"students_count,omitempty"`
	SyllabusEnabled           bool             `json:"syllabus_enabled,omitempty"`
	TeachingMode              string           `json:"teaching_mode,omitempty"`
	LearningMode              string           `json:"learning_mode,omitempty"`
	ScorePublished            bool             `json:"score_published,omitempty"`
	HasAIAbility              bool             `json:"has_ai_ability,omitempty"`
	DingtalkNotifyExam        bool             `json:"dingtalk_notify_exam,omitempty"`
	DingtalkNotifyHomework    bool             `json:"dingtalk_notify_homework,omitempty"`
	KnowledgeGraphPublishType string           `json:"knowledge_graph_publish_type,omitempty"`
	ProblemGraphPublishType   string           `json:"problem_graph_publish_type,omitempty"`
	CreatedAt                 string           `json:"created_at,omitempty"`
	UpdatedAt                 string           `json:"updated_at,omitempty"`
	Modules                   []*Module        `json:"modules,omitempty"`
	Enrollments               []*Enrollment    `json:"enrollments,omitempty"`
	Description               *string          `json:"description,omitempty"`
	AccessCode                *string          `json:"access_code,omitempty"`
	EnableForumEmailNotify    bool             `json:"enable_forum_email_notify,omitempty"`
	Locale                    string           `json:"locale,omitempty"`
}

// Module represents a course module/chapter.
type Module struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Sort         int         `json:"sort"`
	CourseID     int         `json:"course_id,omitempty"`
	IsHidden     int         `json:"is_hidden,omitempty"`
	LessonTimeID int         `json:"lesson_time_id,omitempty"`
	StickyTime   *string     `json:"sticky_time,omitempty"`
	ImportedFrom *string     `json:"imported_from,omitempty"`
	CreatedAt    string      `json:"created_at,omitempty"`
	UpdatedAt    string      `json:"updated_at,omitempty"`
	Syllabuses   []*Syllabus `json:"syllabuses,omitempty"`
	Activities   []*Activity `json:"activities,omitempty"`
}

// Syllabus represents a syllabus item within a module.
type Syllabus struct {
	ID         int         `json:"id"`
	Title      string      `json:"title"`
	Sort       int         `json:"sort"`
	ModuleID   int         `json:"module_id,omitempty"`
	CourseID   int         `json:"course_id,omitempty"`
	CreatedAt  string      `json:"created_at,omitempty"`
	UpdatedAt  string      `json:"updated_at,omitempty"`
	Activities []*Activity `json:"activities,omitempty"`
}

// Enrollment represents a user's enrollment in a course.
type Enrollment struct {
	ID             int         `json:"id"`
	CourseID       int         `json:"course_id,omitempty"`
	UserID         int         `json:"user_id,omitempty"`
	Roles          []*Role     `json:"roles,omitempty"`
	ImportedFrom   *string     `json:"imported_from,omitempty"`
	Aliases        []string    `json:"aliases,omitempty"`
	SeatNumber     string      `json:"seat_number,omitempty"`
	RetakeStatus   bool        `json:"retake_status,omitempty"`
	MoocVideoScore *float64    `json:"mooc_video_score,omitempty"`
	Data           interface{} `json:"data,omitempty"`
	CreatedAt      string      `json:"created_at,omitempty"`
	UpdatedAt      string      `json:"updated_at,omitempty"`
	User           *User       `json:"user,omitempty"`
}

// EnrollmentDetail represents detailed user info from the enrollment endpoint.
// This is the full user object returned by /api/courses/{id}/enrollments/users/{userId}.
type EnrollmentDetail = User

// NavSetting represents navigation setting for a course.
type NavSetting struct {
	Type        string `json:"type"`
	Parent      string `json:"parent"`
	Disabled    bool   `json:"disabled"`
	CanEnabled  bool   `json:"can_enabled"`
	CanDisabled bool   `json:"can_disabled"`
}

// ActivityPublishSetting represents publish settings for activity types.
type ActivityPublishSetting struct {
	Exam     string `json:"exam"`
	Forum    string `json:"forum"`
	Homework string `json:"homework"`
	Others   string `json:"others"`
}

// CompletedResult represents course completion status.
type CompletedResult struct {
	Completed       map[string]bool `json:"completed"`
	TotalActivities int             `json:"total_activities"`
	TotalCompleted  int             `json:"total_completed"`
}

// LastActivity represents the last visited activity info.
type LastActivity struct {
	ID                       int         `json:"id"`
	ActivityEndTime          string      `json:"activity_end_time,omitempty"`
	ActivityStartTime        string      `json:"activity_start_time,omitempty"`
	ActivityType             string      `json:"activity_type,omitempty"`
	AssignGroupIDs           []int       `json:"assign_group_ids,omitempty"`
	AssignStudentIDs         []int       `json:"assign_student_ids,omitempty"`
	AssignTargets            interface{} `json:"assign_targets,omitempty"`
	CompletionCriterionKey   string      `json:"completion_criterion_key,omitempty"`
	CompletionCriterionValue interface{} `json:"completion_criterion_value,omitempty"`
	Data                     interface{} `json:"data,omitempty"`
	EndTime                  *string     `json:"end_time,omitempty"`
	GroupSetID               *int        `json:"group_set_id,omitempty"`
	IsAssignedToAll          bool        `json:"is_assigned_to_all,omitempty"`
	IsInProgress             bool        `json:"is_in_progress,omitempty"`
	ModuleID                 int         `json:"module_id,omitempty"`
}

// CompletenessResponse represents the my-completeness API response.
type CompletenessResponse struct {
	CompletedResult            *CompletedResult `json:"completed_result"`
	LastActivity               *LastActivity    `json:"last_activity"`
	LastUpdateCompletenessTime string           `json:"last_update_completeness_time,omitempty"`
	StudyCompleteness          float64          `json:"study_completeness,omitempty"`
}

// Outline represents a course outline.
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

// OutlineField represents a course outline field.
type OutlineField struct {
	ID          int       `json:"id"`
	Key         string    `json:"key,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Uploads     []*Upload `json:"uploads,omitempty"`
}

// OutlineSettingResponse represents the outline setting API response.
type OutlineSettingResponse struct {
	ID                      int              `json:"id"`
	OrgID                   int              `json:"org_id"`
	FormattedDefaultOptions []*OutlineOption `json:"formatted_default_options,omitempty"`
	FormattedOptions        []*OutlineOption `json:"formatted_options,omitempty"`
}

// OutlineOption represents an outline setting option.
type OutlineOption struct {
	Key      string `json:"key"`
	Title    string `json:"title"`
	Required bool   `json:"required"`
}

// Rollcall represents a roll call session in a course.
type Rollcall struct {
	ID         int    `json:"id"`
	CourseID   int    `json:"course_id,omitempty"`
	ModuleID   int    `json:"module_id,omitempty"`
	Status     string `json:"status,omitempty"`
	Type       string `json:"type,omitempty"`
	StartTime  string `json:"start_time,omitempty"`
	EndTime    string `json:"end_time,omitempty"`
	Duration   int    `json:"duration,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	CreatedBy  *User  `json:"created_by,omitempty"`
	TotalCount int    `json:"total_count,omitempty"`
	SignedCount int   `json:"signed_count,omitempty"`
}

// RollcallRecord represents an individual student's roll call record.
type RollcallRecord struct {
	ID         int    `json:"id"`
	RollcallID int    `json:"rollcall_id,omitempty"`
	UserID     int    `json:"user_id,omitempty"`
	Status     string `json:"status,omitempty"`
	SignedAt   string `json:"signed_at,omitempty"`
	User       *User  `json:"user,omitempty"`
}

// LiveRecord represents a live recording in a course.
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

// Interaction represents a course interaction.
type Interaction struct {
	ID         int         `json:"id"`
	Title      string      `json:"title,omitempty"`
	Type       string      `json:"type,omitempty"`
	Status     string      `json:"status,omitempty"`
	CourseID   int         `json:"course_id,omitempty"`
	ActivityID int         `json:"activity_id,omitempty"`
	CreatedAt  string      `json:"created_at,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// InteractionActivity represents an interaction activity within a course.
type InteractionActivity struct {
	ID         int         `json:"id"`
	Title      string      `json:"title,omitempty"`
	Type       string      `json:"type,omitempty"`
	CourseID   int         `json:"course_id,omitempty"`
	StartTime  *string     `json:"start_time,omitempty"`
	EndTime    *string     `json:"end_time,omitempty"`
	Status     string      `json:"status,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Attributes interface{} `json:"interaction_activity_attributes,omitempty"`
}

// InteractionSubmission represents a student's submission for an interaction.
type InteractionSubmission struct {
	ID         int         `json:"id"`
	UserID     int         `json:"user_id,omitempty"`
	ActivityID int         `json:"activity_id,omitempty"`
	Content    string      `json:"content,omitempty"`
	Score      *float64    `json:"score,omitempty"`
	CreatedAt  string      `json:"created_at,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// IsLockedStatus represents the lock status of an activity.
type IsLockedStatus struct {
	IsLocked      bool            `json:"is_locked"`
	Prerequisites []*Prerequisite `json:"prerequisites,omitempty"`
}

// Prerequisite represents a prerequisite activity for unlocking.
type Prerequisite struct {
	ID           int    `json:"id"`
	ActivityID   int    `json:"activity_id,omitempty"`
	ActivityType string `json:"activity_type,omitempty"`
	Title        string `json:"title,omitempty"`
	Completed    bool   `json:"completed,omitempty"`
}

// BlueprintSubItemsResponse represents blueprint sub-items count response.
type BlueprintSubItemsResponse struct {
	Items []BlueprintSubItem `json:"items"`
}

// BlueprintSubItem represents a single blueprint sub-item entry.
type BlueprintSubItem struct {
	CourseID int `json:"course_id"`
	Count    int `json:"count"`
}

// GroupSet represents a group set containing multiple groups in a course.
type GroupSet struct {
	ID         int      `json:"id"`
	Name       string   `json:"name,omitempty"`
	CourseID   int      `json:"course_id,omitempty"`
	GroupCount int      `json:"group_count,omitempty"`
	Groups     []*Group `json:"groups,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty"`
	UpdatedAt  string   `json:"updated_at,omitempty"`
}

// Group represents a student group within a group set.
type Group struct {
	ID         int     `json:"id"`
	Name       string  `json:"name,omitempty"`
	GroupSetID int     `json:"group_set_id,omitempty"`
	LeaderID   *int    `json:"leader_id,omitempty"`
	Members    []*User `json:"members,omitempty"`
}

// CombineCourse represents a combined course linking multiple courses.
type CombineCourse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name,omitempty"`
	MasterID  int       `json:"master_id,omitempty"`
	CourseIDs []int     `json:"course_ids,omitempty"`
	Courses   []*Course `json:"courses,omitempty"`
	CreatedAt string    `json:"created_at,omitempty"`
}

// CustomScoreItem represents a custom score item for course grading.
type CustomScoreItem struct {
	ID       int     `json:"id"`
	Name     string  `json:"name,omitempty"`
	CourseID int     `json:"course_id,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	Sort     int     `json:"sort,omitempty"`
}

// CompletionCriterion represents a completion criterion definition.
type CompletionCriterion struct {
	Key         string `json:"key"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// AccessCode represents a course access code.
type AccessCode struct {
	Code      string `json:"code"`
	CourseID  int    `json:"course_id,omitempty"`
	IsEnabled bool   `json:"is_enabled,omitempty"`
}

// SignIn represents a sign-in session for a course.
type SignIn struct {
	ID        int    `json:"id"`
	CourseID  int    `json:"course_id,omitempty"`
	Status    string `json:"status,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Type      string `json:"type,omitempty"`
}
