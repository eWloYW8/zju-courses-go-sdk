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

type SyncBlueprintRequest map[string]any

type BlueprintActivityRef struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
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

type CancelBlueprintActivitySyncRequest struct {
	TargetCourseID int    `json:"target_course_id"`
	ID             int    `json:"id"`
	Type           string `json:"type"`
}

type PublishActivitiesRequest map[string]any
