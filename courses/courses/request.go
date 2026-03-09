package courses

type ListCoursesRequest struct {
	Keyword string `json:"keyword,omitempty"`
}

type ListMyCoursesRequest struct {
	Page                  int            `json:"page,omitempty"`
	PageSize              int            `json:"page_size,omitempty"`
	Fields                string         `json:"fields,omitempty"`
	Conditions            map[string]any `json:"conditions,omitempty"`
	ShowScorePassedStatus bool           `json:"showScorePassedStatus,omitempty"`
}

type CreateCourseRequest = Course

type UpdateCourseRequest = Course

type UpdateNavSettingRequest struct {
	NavSetting []*NavSetting `json:"nav_setting"`
}

type UpsertOutlineItemRequest struct {
	ID                int    `json:"id,omitempty"`
	Key               string `json:"key,omitempty"`
	Title             string `json:"title,omitempty"`
	Description       string `json:"description,omitempty"`
	CourseID          int    `json:"course_id,omitempty"`
	SendMessage       bool   `json:"send_message,omitempty"`
	Uploads           []int  `json:"uploads,omitempty"`
	AllowDownloadData []int  `json:"allow_download_data,omitempty"`
}

type CreateModuleRequest = Module

type UpdateModuleRequest = Module

type DeleteModuleOptions struct {
	DeleteRelatedActivity bool
}

type DeleteSyllabusOptions struct {
	DeleteRelatedActivity bool
}

type ResortActivitiesRequest struct {
	CourseID    int   `json:"course_id,omitempty"`
	ActivityIDs []int `json:"activity_ids,omitempty"`
	ModuleID    *int  `json:"module_id,omitempty"`
	SyllabusID  *int  `json:"syllabus_id,omitempty"`
}

type SyncFromURPRequest struct {
	CourseIDs []int `json:"course_ids"`
}

type SyncBlueprintRequest struct {
	Sources         any   `json:"sources,omitempty"`
	TargetCourseIDs []int `json:"target_course_ids,omitempty"`
	Publish         *bool `json:"publish,omitempty"`
}

type BlueprintActivityRef struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type BindBlueprintSubCoursesRequest struct {
	SubCourseIDs []int `json:"sub_course_ids,omitempty"`
}

type ListBlueprintSubCoursesParams struct {
	Keyword    string
	SourceID   int
	SourceType string
}

type BlueprintPrerequisiteSource map[string]any

type CheckBlueprintPrerequisitesRequest struct {
	Sources []BlueprintPrerequisiteSource `json:"sources,omitempty"`
}

type RenameBlueprintSubCourseRequest struct {
	Name string `json:"name,omitempty"`
}

type ListCourseHostsParams struct {
	Type string
}

type ListInstructorEnrollmentsParams struct {
	Fields string
}

type GetEnrollmentUserParams struct {
	Fields       string
	RequestScope string
}

type ListStudentsParams struct {
	IgnoreAvatar bool
}

type ListCourseResourceAuditParams struct {
	Page       int
	PageSize   int
	Conditions any
}

type TPDOEStatStudentsParams struct {
	CourseIDs  []int
	StartDate  string
	EndDate    string
	StatType   string
	Conditions any
}

type SendMailToEnrollmentsRequest struct {
	EnrollmentIDs                 []int  `json:"enrollment_ids,omitempty"`
	EmailToInstructorAndAssistant bool   `json:"email_to_instructor_and_assistant,omitempty"`
	MailSubject                   string `json:"mail_subject,omitempty"`
	MailContent                   string `json:"mail_content,omitempty"`
}

type UpdateEnrollmentSeatNumberRequest struct {
	SeatNumber string `json:"seat_number,omitempty"`
}

type AssistantPermissions struct {
	InstructorAssistant map[string]bool `json:"instructor_assistant,omitempty"`
	StudentAssistant    map[string]bool `json:"student_assistant,omitempty"`
}

type UpdateAssistantPermissionsRequest struct {
	InstructorAssistant map[string]bool       `json:"instructor_assistant,omitempty"`
	StudentAssistant    map[string]bool       `json:"student_assistant,omitempty"`
	ChangedPermissions  *AssistantPermissions `json:"changed_permissions,omitempty"`
}

type StudentsPerformanceParams struct {
	Page             int
	PageSize         int
	Conditions       any
	OnlyStudentsName bool
	IsOriginalScore  *bool
}

type UpdateStudentPerformanceScoreRequest struct {
	StudentID int `json:"student_id"`
	Score     any `json:"score,omitempty"`
}

type CourseAdvanceSettingRequest struct {
	Params map[string]any `json:"params,omitempty"`
}

type WarningStudentsParams struct {
	Fields     string
	Conditions any
}

type ListQuestionnairesParams struct {
	Page       int
	PageSize   int
	Conditions any
}

type ListAllActivitiesByModuleIDsParams struct {
	ModuleIDs               []int
	ActivityTypes           string
	DisableLoadingAnimation bool
}

type ActivityScorePercentageItem struct {
	ID              int    `json:"id"`
	Type            string `json:"type,omitempty"`
	ScorePercentage any    `json:"score_percentage,omitempty"`
	Metrics         any    `json:"metrics,omitempty"`
}

type CustomScorePercentageItem struct {
	ID         int `json:"id"`
	Percentage any `json:"percentage,omitempty"`
	Metrics    any `json:"metrics,omitempty"`
}

type UpdateScorePercentagesRequest struct {
	ActivityScorePercentages               []*ActivityScorePercentageItem `json:"activity_score_percentages,omitempty"`
	CustomScorePercentages                 []*CustomScorePercentageItem   `json:"custom_score_percentages,omitempty"`
	RollcallScorePercentage                any                            `json:"rollcall_score_percentage,omitempty"`
	PerformanceScorePercentage             any                            `json:"performance_score_percentage,omitempty"`
	OnlineVideoCompletenessScorePercentage any                            `json:"online_video_completeness_score_percentage,omitempty"`
	OnlineVideoOBEMetrics                  any                            `json:"online_video_obe_metrics,omitempty"`
}

type CancelBlueprintActivitySyncRequest struct {
	TargetCourseID int    `json:"target_course_id"`
	ID             int    `json:"id"`
	Type           string `json:"type"`
}

type PublishActivitiesRequest map[string]any
