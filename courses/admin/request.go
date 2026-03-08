package admin

type CreateEnrollmentRequest map[string]any
type CreateInviteRequest map[string]any
type ImportRequest map[string]any
type CopyCourseRequest map[string]any

type ChangeEnrollmentRoleRequest struct {
	Role   string `json:"role"`
	RoleID *int   `json:"role_id,omitempty"`
}

type ListUsersForManagementParams struct {
	IgnoreAvatar    bool
	ForManagement   bool
	NeedAIActivated bool
}
