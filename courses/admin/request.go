package admin

type CreateEnrollmentRequest map[string]any
type CreateInviteRequest map[string]any
type ImportRequest map[string]any

type CopyCourseRequest struct {
	CourseIDs []int `json:"course_ids"`
}

type RequestChangePlanRequest struct {
	Period string `json:"period"`
}

type ChangeEnrollmentRoleRequest struct {
	Role   string `json:"role"`
	RoleID *int   `json:"role_id,omitempty"`
}

type UpdateUserDepartmentRequest struct {
	DepartmentID int `json:"department_id"`
}

type UpdateAssistantRolePermissionsRequest struct {
	RolePermissions []*AssistantRolePermissionUpdateItem `json:"role_permissions,omitempty"`
}

type AssistantRolePermissionUpdateItem struct {
	RoleID      int                              `json:"role_id"`
	Permissions []*AssistantRolePermissionUpdate `json:"permissions,omitempty"`
}

type AssistantRolePermissionUpdate struct {
	Code   string `json:"code"`
	Enable bool   `json:"enable"`
}

type ListUsersForManagementParams struct {
	IgnoreAvatar    bool
	ForManagement   bool
	NeedAIActivated bool
}

type SaveOutlineRequiredOptionsRequest struct {
	RequiredOptions []string `json:"required_options,omitempty"`
}
