package resources

type CreateResourceGroupRequest map[string]any
type UpdateResourceGroupRequest map[string]any
type CreateSharedResourceRequest map[string]any
type BatchSaveSharedResourcesRequest map[string]any
type SaveResourcesRequest map[string]any

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
