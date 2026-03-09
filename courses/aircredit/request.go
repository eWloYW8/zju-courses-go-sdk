package aircredit

type AssignCreditsRequest map[string]any
type UpdateCreditsRequest map[string]any
type OptimizeTextRequest map[string]any
type ExportAIPPTUserUsageRequest map[string]any
type LessonPlanStreamRequest map[string]any

type UserTokenParams struct {
	ModuleID     int
	ModuleType   string
	ResourceName string
	UploadID     int
}

type ExportLessonPlanRequest struct {
	Chapters   any `json:"chapters,omitempty"`
	TemplateID int `json:"template_id,omitempty"`
}

type ListCreditStatesParams struct {
	Page       int
	PageSize   int
	Conditions any
}

type CreditStateStatsParams struct {
	StartDate  string
	EndDate    string
	Page       int
	PageSize   int
	Conditions any
}

type ListAuditsParams struct {
	Page       int
	PageSize   int
	Conditions any
}

type UpdateCreditsStatusRequest struct {
	AssignIDs  []int  `json:"assign_ids,omitempty"`
	AssignType string `json:"assign_type,omitempty"`
	Status     string `json:"status,omitempty"`
	Refunded   bool   `json:"refunded,omitempty"`
}

type ClearRemainingCreditsRequest struct {
	AssignID   int    `json:"assign_id"`
	AssignType string `json:"assign_type"`
}

type ModifyUsageLimitRequest struct {
	UsageLimit int `json:"usage_limit"`
}

type CreateCreditAuditRequest struct {
	TargetID   int    `json:"target_id"`
	TargetType string `json:"target_type"`
	Credits    int    `json:"credits"`
	Reason     string `json:"reason,omitempty"`
}

type UpdateCreditAuditRequest struct {
	Action          string `json:"action"`
	ApprovedCredits *int   `json:"approved_credits,omitempty"`
	Remark          string `json:"remark,omitempty"`
}

type UpdateCreditAuditsRequest struct {
	AuditIDs []int  `json:"audit_ids,omitempty"`
	Action   string `json:"action"`
	Remark   string `json:"remark,omitempty"`
}

type SwitchAuditReadRequest struct {
	Read bool `json:"read"`
}

type ListAIPPTUsageStatsParams struct {
	Page       int
	PageSize   int
	Conditions any
}
