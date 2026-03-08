package meetings

type ListVTRSesParams struct {
	Conditions string
	NeedStat   *bool
	Fields     string
}

type ListVTRSShareResourcesParams struct {
	RefParentType string
	Page          int
	PageSize      int
	Conditions    string
}

type ListVTRSResourcesParams struct {
	ParentFolderID   *int
	ClassificationID *int
	Page             int
	PageSize         int
	Conditions       string
}

type CreateVTRSResourcesRequest struct {
	Uploads          any  `json:"uploads,omitempty"`
	Resources        any  `json:"resources,omitempty"`
	ParentFolderID   *int `json:"parent_folder_id,omitempty"`
	ClassificationID *int `json:"classification_id,omitempty"`
}

type UploadReferenceIDsRequest struct {
	UploadReferenceIDs []int `json:"upload_reference_ids,omitempty"`
}

type UploadReferencesRequest struct {
	UploadReferences []int `json:"upload_references,omitempty"`
}

type CreateVTRSSubjectLibRequest struct {
	Title            string `json:"title,omitempty"`
	ParentID         *int   `json:"parent_id,omitempty"`
	ClassificationID *int   `json:"classification_id,omitempty"`
}

type ListVTRSSubjectLibsParams struct {
	Keyword          string
	ParentID         *int
	ClassificationID *int
	Page             int
	PageSize         int
	Predicate        string
	Reverse          *bool
	LibType          string
}

type CreateVTRSMeetingClassificationRequest struct {
	Name     string `json:"name,omitempty"`
	Category string `json:"category,omitempty"`
}

type UpdateVTRSMeetingClassificationRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

type CreateVTRSResourceClassificationRequest struct {
	Name          string `json:"name,omitempty"`
	Category      string `json:"category,omitempty"`
	AllowCopy     *bool  `json:"allow_copy,omitempty"`
	AllowDownload *bool  `json:"allow_download,omitempty"`
	CCLicenseName string `json:"cc_license_name,omitempty"`
}

type UpdateVTRSResourceClassificationRequest struct {
	Name          string `json:"name,omitempty"`
	AllowCopy     *bool  `json:"allow_copy,omitempty"`
	AllowDownload *bool  `json:"allow_download,omitempty"`
	CCLicenseName string `json:"cc_license_name,omitempty"`
}

type SortVTRSResourceClassificationsRequest struct {
	Classifications []int `json:"classifications,omitempty"`
}
