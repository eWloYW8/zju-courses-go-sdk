package model

// SubjectLib represents a subject library.
type SubjectLib struct {
	ID        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Type      string `json:"type,omitempty"`
	Nums      int    `json:"nums,omitempty"`
	IsFolder  bool   `json:"is_folder,omitempty"`
	IsShared  bool   `json:"is_shared,omitempty"`
	ParentID  int    `json:"parent_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// SharedResource represents a shared resource.
type SharedResource struct {
	ID                   int         `json:"id"`
	Name                 string      `json:"name,omitempty"`
	OrgID                int         `json:"org_id,omitempty"`
	ParentID             int         `json:"parent_id,omitempty"`
	ResourceType         string      `json:"resource_type,omitempty"`
	AuditStatus          string      `json:"audit_status,omitempty"`
	OpenScope            string      `json:"open_scope,omitempty"`
	AllowDownload        bool        `json:"allow_download,omitempty"`
	AllowSave            bool        `json:"allow_save,omitempty"`
	IsFolder             bool        `json:"is_folder,omitempty"`
	CcLicenseName        string      `json:"cc_license_name,omitempty"`
	CcLicenseCode        string      `json:"cc_license_code,omitempty"`
	CcLicenseDescription string      `json:"cc_license_description,omitempty"`
	Upload               *Upload     `json:"upload,omitempty"`
	SubjectLibrary       *SubjectLib `json:"subject_library,omitempty"`
}
