package others

type ListProjectsParams struct {
	Page       int
	PageSize   int
	Conditions any
}

type ListEntriesParams struct {
	Page       int
	PageSize   int
	Conditions any
	Fields     string
}

type ListEntryReferencesParams struct {
	Page     int
	PageSize int
}

type CreateProjectRequest struct {
	Name string `json:"name"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type AuditProjectApplicationRequest struct {
	Status string `json:"status"`
}

type ProjectSharedResourceConditions struct {
	Keyword     string `json:"keyword,omitempty"`
	RefParentID *int   `json:"ref_parent_id,omitempty"`
}

type ProjectSharedResourceRequest struct {
	ResourceNewName string `json:"resource_new_name,omitempty"`
	AllowDownload   string `json:"allow_download,omitempty"`
	Uploads         []int  `json:"uploads,omitempty"`
	NodeIDs         []any  `json:"node_ids,omitempty"`
	ReferenceID     *int   `json:"reference_id,omitempty"`
}

type DeleteProjectSharedResourceRequest struct {
	ReferenceID int   `json:"reference_id"`
	UploadID    int   `json:"upload_id"`
	NodeIDs     []int `json:"node_ids,omitempty"`
}

type EntryKeyword struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
