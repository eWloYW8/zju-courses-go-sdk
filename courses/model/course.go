package model

// CourseAttributes represents course-specific attributes.
type CourseAttributes struct {
	AudienceType          interface{} `json:"audience_type"`
	CopyStatus            *string     `json:"copy_status,omitempty"`
	Data                  interface{} `json:"data,omitempty"`
	GraduateMethod        string      `json:"graduate_method,omitempty"`
	IsDuringPublishPeriod bool        `json:"is_during_publish_period,omitempty"`
	PassingScore          *float64    `json:"passing_score,omitempty"`
	Published             bool        `json:"published,omitempty"`
	ScoreType             *string     `json:"score_type,omitempty"`
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
	ID                        int               `json:"id"`
	Name                      string            `json:"name"`
	DisplayName               string            `json:"display_name,omitempty"`
	SecondName                *string           `json:"second_name,omitempty"`
	CourseCode                string            `json:"course_code,omitempty"`
	CourseType                int               `json:"course_type,omitempty"`
	Cover                     string            `json:"cover,omitempty"`
	SmallCover                string            `json:"small_cover,omitempty"`
	Credit                    *float64          `json:"credit,omitempty"`
	CreditState               *CreditState      `json:"credit_state,omitempty"`
	AcademicYearID            int               `json:"academic_year_id,omitempty"`
	AcademicYear              *AcademicYear     `json:"academic_year,omitempty"`
	SemesterID                int               `json:"semester_id,omitempty"`
	Semester                  *Semester         `json:"semester,omitempty"`
	DepartmentID              int               `json:"department_id,omitempty"`
	Department                *Department       `json:"department,omitempty"`
	Org                       *Org              `json:"org,omitempty"`
	OrgID                     int               `json:"org_id,omitempty"`
	CourseAttributes          *CourseAttributes `json:"course_attributes,omitempty"`
	Instructors               []*User           `json:"instructors,omitempty"`
	CreatedUser               *User             `json:"created_user,omitempty"`
	ImportedFrom              string            `json:"imported_from,omitempty"`
	StartDate                 *string           `json:"start_date,omitempty"`
	EndDate                   *string           `json:"end_date,omitempty"`
	Grade                     *Grade            `json:"grade,omitempty"`
	IsStarted                 bool              `json:"is_started,omitempty"`
	IsClosed                  bool              `json:"is_closed,omitempty"`
	IsInstructor              bool              `json:"is_instructor,omitempty"`
	IsMute                    bool              `json:"is_mute,omitempty"`
	Registered                bool              `json:"registered,omitempty"`
	IsTeamTeaching            bool              `json:"is_team_teaching,omitempty"`
	IsDefaultCourseCover      bool              `json:"is_default_course_cover,omitempty"`
	IsStudio                  bool              `json:"is_studio,omitempty"`
	IsBlocked                 *bool             `json:"is_blocked,omitempty"`
	IsBlueprintCourse         bool              `json:"is_blueprint_course,omitempty"`
	IsBlueprintSubCourse      bool              `json:"is_blueprint_sub_course,omitempty"`
	IsCombinedCourse          bool              `json:"is_combined_course,omitempty"`
	IsPublic                  bool              `json:"is_public,omitempty"`
	Archived                  bool              `json:"archived,omitempty"`
	AllowClone                bool              `json:"allow_clone,omitempty"`
	AuditStatus               string            `json:"audit_status,omitempty"`
	AuditRemark               *string           `json:"audit_remark,omitempty"`
	PublicScope               string            `json:"public_scope,omitempty"`
	Compulsory                *bool             `json:"compulsory,omitempty"`
	Klass                     *Class            `json:"klass,omitempty"`
	CanWithdrawCourse         bool              `json:"can_withdraw_course,omitempty"`
	ClassroomSchedule         *string           `json:"classroom_schedule,omitempty"`
	KnowledgeNodeCount        int               `json:"knowledge_node_count,omitempty"`
	StudyCompleteness         *float64          `json:"study_completeness,omitempty"`
	UserStickCourseRecord     interface{}       `json:"user_stick_course_record,omitempty"`
	SubjectCode               *string           `json:"subject_code,omitempty"`
	AllowAdminUpdateBasicInfo bool              `json:"allow_admin_update_basic_info,omitempty"`
	AllowUpdateBasicInfo      bool              `json:"allow_update_basic_info,omitempty"`
	AllowedToInviteAssistant  bool              `json:"allowed_to_invite_assistant,omitempty"`
	AllowedToInviteStudent    bool              `json:"allowed_to_invite_student,omitempty"`
	AllowedToJoinCourse       bool              `json:"allowed_to_join_course,omitempty"`
	StudentsCount             int               `json:"students_count,omitempty"`
	SyllabusEnabled           bool              `json:"syllabus_enabled,omitempty"`
	TeachingMode              string            `json:"teaching_mode,omitempty"`
	LearningMode              string            `json:"learning_mode,omitempty"`
	ScorePublished            bool              `json:"score_published,omitempty"`
	HasAIAbility              bool              `json:"has_ai_ability,omitempty"`
	DingtalkNotifyExam        bool              `json:"dingtalk_notify_exam,omitempty"`
	DingtalkNotifyHomework    bool              `json:"dingtalk_notify_homework,omitempty"`
	KnowledgeGraphPublishType string            `json:"knowledge_graph_publish_type,omitempty"`
	ProblemGraphPublishType   string            `json:"problem_graph_publish_type,omitempty"`
	CreatedAt                 string            `json:"created_at,omitempty"`
	UpdatedAt                 string            `json:"updated_at,omitempty"`
	URL                       string            `json:"url,omitempty"`
	TeachingUnitType          string            `json:"teaching_unit_type,omitempty"`
	CurrentUserIsMember       bool              `json:"current_user_is_member,omitempty"`
	TeamTeachings             []*User           `json:"team_teachings,omitempty"`
	Modules                   []*Module         `json:"modules,omitempty"`
	Enrollments               []*Enrollment     `json:"enrollments,omitempty"`
	Description               *string           `json:"description,omitempty"`
	AccessCode                *string           `json:"access_code,omitempty"`
	EnableForumEmailNotify    bool              `json:"enable_forum_email_notify,omitempty"`
	Locale                    string            `json:"locale,omitempty"`
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

// Prerequisite represents a prerequisite activity for unlocking.
type Prerequisite struct {
	ID           int    `json:"id"`
	ActivityID   int    `json:"activity_id,omitempty"`
	ActivityType string `json:"activity_type,omitempty"`
	Title        string `json:"title,omitempty"`
	Completed    bool   `json:"completed,omitempty"`
}
