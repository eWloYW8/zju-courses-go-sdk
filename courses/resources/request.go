package resources

type CreateResourceGroupRequest map[string]any
type UpdateResourceGroupRequest map[string]any
type CreateSharedResourceRequest map[string]any
type BatchSaveSharedResourcesRequest map[string]any
type SaveResourcesRequest map[string]any

type UpdateResourceGroupResourceRequest struct {
	Name          string `json:"name,omitempty"`
	AllowDownload *bool  `json:"allow_download,omitempty"`
	CCLicenseName string `json:"cc_license_name,omitempty"`
}

type ListResourceGroupItemsParams struct {
	Page       int
	PageSize   int
	Conditions string
}

type ListSharedResourcesParams struct {
	Page       int
	PageSize   int
	Conditions string
}

type ListHomepageSharedResourcesParams struct {
	DepartmentID     int
	ClassificationID int
}

type ListRecentUsedSharedResourcesParams struct {
	ClassificationID string
	DepartmentIDs    string
}
