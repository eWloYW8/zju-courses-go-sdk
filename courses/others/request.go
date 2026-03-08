package others

type ListProjectsParams struct {
	Page       int
	PageSize   int
	Conditions string
}

type AuditProjectApplicationRequest struct {
	Status string `json:"status"`
}
