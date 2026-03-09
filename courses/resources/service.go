package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles resource management API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

func apiPrefix(anonymous bool) string {
	if anonymous {
		return "/anonymous-api"
	}
	return "/api"
}

// --- Resource Groups ---

// ListResourceGroups returns resource groups.
func (s *Service) ListResourceGroups(ctx context.Context, opts *model.ListOptions) (*ResourceGroupsResponse, error) {
	u := addListOptions("/api/resource-groups", opts)
	result := new(ResourceGroupsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetResourceGroup returns a specific resource group.
func (s *Service) GetResourceGroup(ctx context.Context, groupID int) (*ResourceGroup, error) {
	u := fmt.Sprintf("/api/resource-groups/%d", groupID)
	result := new(ResourceGroup)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateResourceGroup creates a new resource group.
func (s *Service) CreateResourceGroup(ctx context.Context, body interface{}) (*ResourceGroup, error) {
	result := new(ResourceGroup)
	_, err := s.client.Post(ctx, "/api/resource-group", body, result)
	return result, err
}

// UpdateResourceGroup updates a resource group.
func (s *Service) UpdateResourceGroup(ctx context.Context, groupID int, body interface{}) error {
	u := fmt.Sprintf("/api/resource-group/%d", groupID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// DeleteResourceGroup deletes a resource group.
func (s *Service) DeleteResourceGroup(ctx context.Context, groupID int) error {
	u := fmt.Sprintf("/api/resource-group/%d", groupID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ListPagedResourceGroups returns resource groups using the frontend paging POST API.
func (s *Service) ListPagedResourceGroups(ctx context.Context, opts *model.ListOptions, body interface{}) (*ResourceGroupsResponse, error) {
	u := addListOptions("/api/resource-groups", opts)
	result := new(ResourceGroupsResponse)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// GetResourceGroupMembers returns members of a resource group.
func (s *Service) GetResourceGroupMembers(ctx context.Context, groupID int) (*ResourceGroupMembersResponse, error) {
	u := fmt.Sprintf("/api/resource-groups/%d/members", groupID)
	result := new(ResourceGroupMembersResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListPagedResourceGroupMembers returns paginated members of a resource group.
func (s *Service) ListPagedResourceGroupMembers(ctx context.Context, groupID int, opts *model.ListOptions) (*ResourceGroupMembersResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/resource-groups/%d/members", groupID), opts)
	result := new(ResourceGroupMembersResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// DeleteResourceGroupMembers removes members from a resource group.
func (s *Service) DeleteResourceGroupMembers(ctx context.Context, groupID int, body interface{}) error {
	u := fmt.Sprintf("/api/resource-groups/%d/member", groupID)
	_, err := s.client.DeleteWithBody(ctx, u, body, nil)
	return err
}

// LeaveResourceGroup leaves a resource group.
func (s *Service) LeaveResourceGroup(ctx context.Context, groupID int) error {
	u := fmt.Sprintf("/api/resource-groups/%d/leave", groupID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// ListResourceGroupFolders returns folders in a resource group.
func (s *Service) ListResourceGroupFolders(ctx context.Context) (*ResourceGroupFoldersResponse, error) {
	return s.ListResourceGroupFoldersWithParams(ctx, ListResourceGroupItemsParams{})
}

// ListResourceGroupFoldersWithParams returns paged folders across all resource groups.
func (s *Service) ListResourceGroupFoldersWithParams(ctx context.Context, params ListResourceGroupItemsParams) (*ResourceGroupFoldersResponse, error) {
	u := addListOptions("/api/resource-groups/folders", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if params.Conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": params.Conditions})
	}
	result := new(ResourceGroupFoldersResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListResourceGroupResources returns resources in a resource group.
func (s *Service) ListResourceGroupResources(ctx context.Context) (*ResourceGroupResourcesResponse, error) {
	return s.ListResourceGroupResourcesWithParams(ctx, ListResourceGroupItemsParams{})
}

// ListResourceGroupResourcesWithParams returns paged resources across all resource groups.
func (s *Service) ListResourceGroupResourcesWithParams(ctx context.Context, params ListResourceGroupItemsParams) (*ResourceGroupResourcesResponse, error) {
	u := addListOptions("/api/resource-groups/resources", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if params.Conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": params.Conditions})
	}
	result := new(ResourceGroupResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListPagedResourceGroupFolders returns folders in a specific resource group.
func (s *Service) ListPagedResourceGroupFolders(ctx context.Context, groupID int, opts *model.ListOptions, conditions string) (*ResourceGroupFoldersResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/resource-groups/%d/folders", groupID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(ResourceGroupFoldersResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListPagedResourceGroupResources returns resources in a specific resource group.
func (s *Service) ListPagedResourceGroupResources(ctx context.Context, groupID int, opts *model.ListOptions, conditions string) (*ResourceGroupResourcesResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/resource-groups/%d/resources", groupID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(ResourceGroupResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListPagedResourceGroupRubrics returns rubrics in a specific resource group.
func (s *Service) ListPagedResourceGroupRubrics(ctx context.Context, groupID int, opts *model.ListOptions, conditions string) (*ResourceGroupRubricsResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/resource-groups/%d/rubrics", groupID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(ResourceGroupRubricsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListPagedResourceGroupSubjectLibs returns subject libs in a specific resource group.
func (s *Service) ListPagedResourceGroupSubjectLibs(ctx context.Context, groupID int, opts *model.ListOptions, conditions string) (*ResourceGroupSubjectLibsResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/resource-groups/%d/subject-libs", groupID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(ResourceGroupSubjectLibsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// DeleteResourceGroupFolder deletes a folder from a resource group.
func (s *Service) DeleteResourceGroupFolder(ctx context.Context, groupID int, folderID int) error {
	u := fmt.Sprintf("/api/resource-groups/%d/folders/%d", groupID, folderID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// UpdateResourceGroupResource updates a shared resource record inside a resource group.
func (s *Service) UpdateResourceGroupResource(ctx context.Context, groupID int, resourceID int, body *UpdateResourceGroupResourceRequest) error {
	u := fmt.Sprintf("/api/resource-groups/%d/resource/%d", groupID, resourceID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// DeleteResourceGroupSubjectLib deletes a subject lib record referenced by a resource group.
func (s *Service) DeleteResourceGroupSubjectLib(ctx context.Context, subjectLibID int) error {
	u := fmt.Sprintf("/api/subject-libs/%d", subjectLibID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// --- Resource Folders ---

// ListResourceFolders returns resource folders.
func (s *Service) ListResourceFolders(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/resource/folders", &result)
	return result, err
}

// --- Public Resources ---

// ListPublicResources returns public resources.
func (s *Service) ListPublicResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/public-resources", &result)
	return result, err
}

// --- Shared Resources ---

// ListSharedResourcesToMe returns resources shared to the current user.
func (s *Service) ListSharedResourcesToMe(ctx context.Context) (*SharedResourcesResponse, error) {
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, "/api/shared-resources-to-me", result)
	return result, err
}

// ListSharedResources returns shared resources with pagination.
func (s *Service) ListSharedResources(ctx context.Context, opts *model.ListOptions) (*SharedResourcesResponse, error) {
	params := ListSharedResourcesParams{}
	if opts != nil {
		params.Page = opts.Page
		params.PageSize = opts.PageSize
	}
	return s.ListSharedResourcesWithParams(ctx, params)
}

// ListSharedResourcesWithParams returns shared resources with pagination and conditions.
func (s *Service) ListSharedResourcesWithParams(ctx context.Context, params ListSharedResourcesParams) (*SharedResourcesResponse, error) {
	u := addListOptions("/api/shared-resources-no-repeated", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if params.Conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": params.Conditions})
	}
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetSharedResource returns a specific shared resource.
func (s *Service) GetSharedResource(ctx context.Context, resourceID int) (*model.SharedResource, error) {
	u := fmt.Sprintf("/api/shared-resources/%d", resourceID)
	result := new(model.SharedResource)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateSharedResource creates a shared resource.
func (s *Service) CreateSharedResource(ctx context.Context, body interface{}) (*model.SharedResource, error) {
	result := new(model.SharedResource)
	_, err := s.client.Post(ctx, "/api/shared-resources", body, result)
	return result, err
}

// BatchSaveSharedResources batch saves shared resources.
func (s *Service) BatchSaveSharedResources(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/shared-resources/batch-save", body, nil)
	return err
}

// GetSharedResourceStats returns shared resource statistics.
func (s *Service) GetSharedResourceStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/shared-resources/stat", &result)
	return result, err
}

// GetSharedVideoResourceStats returns shared video resource statistics.
func (s *Service) GetSharedVideoResourceStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/shared-resources/stat/video-resources", &result)
	return result, err
}

// ExportSharedVideoResourceStats exports shared video resource statistics.
func (s *Service) ExportSharedVideoResourceStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/shared-resources/stat/video-resources/export", &result)
	return result, err
}

// GetSharedResourceClassifications returns shared resource classifications.
func (s *Service) GetSharedResourceClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/shared-resource/classifications", &result)
	return result, err
}

// ListSharedResourceClassificationsWithPrefix returns shared-resource classifications from either /api or /anonymous-api.
func (s *Service) ListSharedResourceClassificationsWithPrefix(ctx context.Context, anonymous bool) (*ResourceClassificationsResponse, error) {
	result := new(ResourceClassificationsResponse)
	_, err := s.client.Get(ctx, apiPrefix(anonymous)+"/shared-resource/classifications", result)
	return result, err
}

// ListSharedResourceManagement returns shared resource management list.
func (s *Service) ListSharedResourceManagement(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/shared-resources/management", &result)
	return result, err
}

// ShareToOtherOrgs shares resources to other organizations (admin).
func (s *Service) ShareToOtherOrgs(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/shared-resources/admin/to-other-orgs", body, nil)
	return err
}

// ListShareToOtherOrgsOpenedOrgs returns organizations available for the shared-resource admin view.
func (s *Service) ListShareToOtherOrgsOpenedOrgs(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/toggle-opened-orgs?toggle=share_resource_to_other_orgs&include_current_org=0", &result)
	return result, err
}

// GetUserSharedResources returns shared resources for a user.
func (s *Service) GetUserSharedResources(ctx context.Context, userID int) (*SharedResourcesResponse, error) {
	u := fmt.Sprintf("/api/shared-resources/user/%d", userID)
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListSharedResourcesFromMe returns shared resources created by the current user.
func (s *Service) ListSharedResourcesFromMe(ctx context.Context, opts *model.ListOptions, conditions string) (*SharedResourcesResponse, error) {
	u := addListOptions("/api/shared-resources/from-me", opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListMostLikedSharedResources returns most liked shared resources.
func (s *Service) ListMostLikedSharedResources(ctx context.Context, conditions string) (*SharedResourcesResponse, error) {
	u := "/api/shared-resources/most-liked"
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListMostLikedSharedResourcesWithPrefix returns most-liked shared resources from either /api or /anonymous-api.
func (s *Service) ListMostLikedSharedResourcesWithPrefix(ctx context.Context, anonymous bool, conditions string) (*SharedResourcesResponse, error) {
	u := apiPrefix(anonymous) + "/shared-resources/most-liked"
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListRecentUsedSharedResources returns recently used shared resources.
func (s *Service) ListRecentUsedSharedResources(ctx context.Context, classificationID string, departmentIDs string) (*SharedResourcesResponse, error) {
	params := map[string]string{}
	if classificationID != "" {
		params["classificationId"] = classificationID
	}
	if departmentIDs != "" {
		params["departmentIds"] = departmentIDs
	}
	u := addQueryParams("/api/shared-resources/recent-used", params)
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListRecentUsedSharedResourcesWithPrefix returns recently used shared resources from either /api or /anonymous-api.
func (s *Service) ListRecentUsedSharedResourcesWithPrefix(ctx context.Context, anonymous bool, params ListRecentUsedSharedResourcesParams) (*SharedResourcesResponse, error) {
	query := map[string]string{}
	if params.ClassificationID != "" {
		query["classificationId"] = params.ClassificationID
	}
	if params.DepartmentIDs != "" {
		query["departmentIds"] = params.DepartmentIDs
	}
	u := addQueryParams(apiPrefix(anonymous)+"/shared-resources/recent-used", query)
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListHomepageSharedResources returns the homepage shared-resources block using the frontend no-intercept query.
func (s *Service) ListHomepageSharedResources(ctx context.Context, anonymous bool, params ListHomepageSharedResourcesParams) (*SharedResourcesResponse, error) {
	conditions := map[string]any{
		"parent_id": 0,
		"order_by":  "view_count",
	}
	if params.DepartmentID > 0 {
		conditions["department"] = params.DepartmentID
	}
	if params.ClassificationID > 0 {
		conditions["classification"] = params.ClassificationID
	}
	u := addQueryParams(apiPrefix(anonymous)+"/shared-resources?no-intercept=true", map[string]string{
		"conditions": encodeConditions(conditions),
	})
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// SearchSharedResourcesWithPrefix searches shared resources through the /api or /anonymous-api listing endpoint.
func (s *Service) SearchSharedResourcesWithPrefix(ctx context.Context, anonymous bool, params ListSharedResourcesParams) (*SharedResourcesResponse, error) {
	u := addListOptions(apiPrefix(anonymous)+"/shared-resources", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if params.Conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": params.Conditions})
	}
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// SaveSharedResource saves a shared resource to the current user context.
func (s *Service) SaveSharedResource(ctx context.Context, resourceID int) error {
	u := fmt.Sprintf("/api/shared-resources/%d/save", resourceID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// DeleteSharedResource deletes a shared resource.
func (s *Service) DeleteSharedResource(ctx context.Context, resourceID int) error {
	u := fmt.Sprintf("/api/shared-resources/%d", resourceID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ListUserCollections returns shared-resource collections for a user.
func (s *Service) ListUserCollections(ctx context.Context, userID int, opts *model.ListOptions, conditions string) (*SharedResourcesResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/shared-resources/user/%d/collections", userID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(SharedResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// SetSharedResourceCollection marks a shared resource as collected by a user.
func (s *Service) SetSharedResourceCollection(ctx context.Context, resourceID int, userID int) error {
	u := fmt.Sprintf("/api/shared-resources/%d/user/%d/collection", resourceID, userID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// UnsetSharedResourceCollection removes collection status for a shared resource.
func (s *Service) UnsetSharedResourceCollection(ctx context.Context, resourceID int, userID int) error {
	u := fmt.Sprintf("/api/shared-resources/%d/user/%d/collection", resourceID, userID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// BuildSharedResourcePreviewURL builds the frontend preview URL for a shared resource video.
func (s *Service) BuildSharedResourcePreviewURL(resourceID int, resolution string) string {
	u := fmt.Sprintf("/shared-resources/%d/preview?preview=true", resourceID)
	if resolution != "" {
		u += "&resolution=" + url.QueryEscape(resolution)
	}
	return u
}

// BuildSharedResourceSCORMURL builds the frontend SCORM preview URL for a shared resource.
func (s *Service) BuildSharedResourceSCORMURL(resourceID int, sco string, parameters string) string {
	u := fmt.Sprintf("/shared-resources/%d/scorms?sco=%s", resourceID, url.QueryEscape(sco))
	if parameters != "" {
		u += "&para=" + url.QueryEscape(parameters)
	}
	return u
}

// --- Save Resources ---

// SaveResources saves resources.
func (s *Service) SaveResources(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/save-resources", body, nil)
	return err
}

// CheckSaveResources checks if resources can be saved.
func (s *Service) CheckSaveResources(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/save-resources/check", body, &result)
	return result, err
}

// ListSubjectLibFolders returns subject-lib folders for the save-resource picker.
func (s *Service) ListSubjectLibFolders(ctx context.Context, parentID int) (*SubjectLibFoldersResponse, error) {
	u := addQueryParams("/api/subject-libs/folders", map[string]string{"parent_id": fmt.Sprintf("%d", parentID)})
	result := new(SubjectLibFoldersResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetSubjectLibWithPrefix returns a subject library from either /api or /anonymous-api.
func (s *Service) GetSubjectLibWithPrefix(ctx context.Context, anonymous bool, subjectLibID int) (*model.SubjectLib, error) {
	u := fmt.Sprintf("%s/subject-libs/%d", apiPrefix(anonymous), subjectLibID)
	result := new(model.SubjectLib)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListDepartmentsWithPrefix returns departments from either /api or /anonymous-api with an optional field list.
func (s *Service) ListDepartmentsWithPrefix(ctx context.Context, anonymous bool, fields string) (*DepartmentsResponse, error) {
	u := apiPrefix(anonymous) + "/departments"
	if fields != "" {
		u = addQueryParams(u, map[string]string{"fields": fields})
	}
	result := new(DepartmentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListHomepageDepartmentsWithPrefix returns the departments shown on the homepage.
func (s *Service) ListHomepageDepartmentsWithPrefix(ctx context.Context, anonymous bool) (*DepartmentsResponse, error) {
	result := new(DepartmentsResponse)
	_, err := s.client.Get(ctx, apiPrefix(anonymous)+"/departments/show-on-homepage", result)
	return result, err
}

// CopyThirdPartResources copies third-party resources.
func (s *Service) CopyThirdPartResources(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/copy-third-part-resources", body, nil)
	return err
}

// --- Resource Files ---

// GetResourceFile returns a resource file.
func (s *Service) GetResourceFile(ctx context.Context, fileID int) (*model.Upload, error) {
	u := fmt.Sprintf("/api/resource-file/%d", fileID)
	result := new(model.Upload)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListResourceActivities returns activities referencing a resource.
func (s *Service) ListResourceActivities(ctx context.Context, resourceID int, opts *model.ListOptions, conditions string) (json.RawMessage, error) {
	u := addListOptions(fmt.Sprintf("/api/resources/%d/activities", resourceID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetOrgResourceFileArrears returns org resource file arrears.
func (s *Service) GetOrgResourceFileArrears(ctx context.Context, orgID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/resource-file/org/arrears/%d", orgID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Lesson Resources ---

// GetLessonResourcesSharedStat returns shared statistics for lesson resources.
func (s *Service) GetLessonResourcesSharedStat(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/lesson-resources/shared-stat", &result)
	return result, err
}

// --- Virtual Experiments ---

// GetVirtualExperiment returns a virtual experiment.
func (s *Service) GetVirtualExperiment(ctx context.Context, experimentID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/virtual-experiments/%d", experimentID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Virtual Classroom Resources ---

// ListVirtualClassroomResources returns virtual classroom resources.
func (s *Service) ListVirtualClassroomResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/virtual-classroom-resources", &result)
	return result, err
}

// --- CC License ---

// GetCCLicenseGroups returns Creative Commons license groups.
func (s *Service) GetCCLicenseGroups(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/cc-license/groups", &result)
	return result, err
}

// GetCCLicenseMap returns Creative Commons license map.
func (s *Service) GetCCLicenseMap(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/cc-license/map", &result)
	return result, err
}

// --- Slides ---

// ListPublishedSlides returns published slides.
func (s *Service) ListPublishedSlides(ctx context.Context, fields string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/slides/published?fields=%s", fields)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// SearchSlides searches slides by keyword.
func (s *Service) SearchSlides(ctx context.Context, keyword string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/slides?keyword=%s", keyword)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Video Suites ---

// GetVideoSuite returns a video suite.
func (s *Service) GetVideoSuite(ctx context.Context, suiteID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/video-suites/%d", suiteID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListVideoSuiteComments returns comments for a video suite.
func (s *Service) ListVideoSuiteComments(ctx context.Context, suiteID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/video-suite/comments/%d", suiteID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Video Quizzes ---

// GetVideoQuiz returns a video quiz.
func (s *Service) GetVideoQuiz(ctx context.Context, quizID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/video-quizzes/%d", quizID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVideoQuizOrgArrears returns video quiz org arrears.
func (s *Service) GetVideoQuizOrgArrears(ctx context.Context, orgID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/video-quizzes/org/arrears/%d", orgID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Certifications ---

// ListCertifications returns certifications.
func (s *Service) ListCertifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/certifications", &result)
	return result, err
}

// --- Course Packages ---

// ListCoursePackages returns course packages.
func (s *Service) ListCoursePackages(ctx context.Context) (*CoursePackagesResponse, error) {
	result := new(CoursePackagesResponse)
	_, err := s.client.Get(ctx, "/api/course-packages", result)
	return result, err
}

// GetCoursePackage returns a specific course package.
func (s *Service) GetCoursePackage(ctx context.Context, packageID int) (*CoursePackage, error) {
	u := fmt.Sprintf("/api/course-packages/%d", packageID)
	result := new(CoursePackage)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListCoursePackagesForCourse returns course packages under a course with pagination.
func (s *Service) ListCoursePackagesForCourse(ctx context.Context, courseID int, opts *model.ListOptions, keyword string) (*CoursePackagesResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/courses/%d/course-package", courseID), opts)
	if keyword != "" {
		u = addQueryParams(u, map[string]string{"keyword": keyword})
	}
	var raw struct {
		Data CoursePackagesResponse `json:"data"`
	}
	_, err := s.client.Get(ctx, u, &raw)
	return &raw.Data, err
}

// ExportCoursePackage starts exporting a course package.
func (s *Service) ExportCoursePackage(ctx context.Context, courseID int, body interface{}) error {
	u := fmt.Sprintf("/api/courses/%d/course-package/export", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// GetCoursePackageExportStatus returns export status for the current course package job.
func (s *Service) GetCoursePackageExportStatus(ctx context.Context, courseID int) (*CoursePackageExportStatusResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/course-package/status", courseID)
	var raw struct {
		Data CoursePackageExportStatusResponse `json:"data"`
	}
	_, err := s.client.Get(ctx, u, &raw)
	return &raw.Data, err
}

// UpdateCoursePackage updates a course package.
func (s *Service) UpdateCoursePackage(ctx context.Context, packageID int, body interface{}, noCheck bool) error {
	u := fmt.Sprintf("/api/course-packages/%d", packageID)
	if noCheck {
		u += "?no_check=true"
	}
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// DeleteCoursePackage deletes a course package.
func (s *Service) DeleteCoursePackage(ctx context.Context, packageID int) error {
	u := fmt.Sprintf("/api/course-packages/%d", packageID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// SaveCoursePackage saves a course package into resources.
func (s *Service) SaveCoursePackage(ctx context.Context, packageID int) error {
	u := fmt.Sprintf("/api/course-packages/%d/save", packageID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// --- Public Courses ---

// ListPublicCourses returns public courses.
func (s *Service) ListPublicCourses(ctx context.Context) (*PublicCoursesResponse, error) {
	result := new(PublicCoursesResponse)
	_, err := s.client.Get(ctx, "/api/courses/public", result)
	return result, err
}

// ListHotPublicCourses returns hot public courses.
func (s *Service) ListHotPublicCourses(ctx context.Context, limit int) (*PublicCoursesResponse, error) {
	u := fmt.Sprintf("/api/courses/public/hot?limit=%d", limit)
	result := new(PublicCoursesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListLatestPublicCourses returns latest public courses.
func (s *Service) ListLatestPublicCourses(ctx context.Context, limit int) (*PublicCoursesResponse, error) {
	u := fmt.Sprintf("/api/courses/public/latest?limit=%d", limit)
	result := new(PublicCoursesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Course Estimates ---

// GetCourseEstimate returns a course estimate.
func (s *Service) GetCourseEstimate(ctx context.Context, estimateID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course-estimate/%d", estimateID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListCourseEstimates returns course estimates.
func (s *Service) ListCourseEstimates(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course-estimates/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateCourseEstimate creates a course estimate.
func (s *Service) CreateCourseEstimate(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/course-estimate", body, &result)
	return result, err
}

// CreateCourseEstimateReply creates a reply to a course estimate.
func (s *Service) CreateCourseEstimateReply(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/course-estimate-reply", body, &result)
	return result, err
}

// ListCourseEstimateReplies returns replies for a course estimate.
func (s *Service) ListCourseEstimateReplies(ctx context.Context, estimateID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course-estimate-replies/%d", estimateID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListCourseEstimatesReplies returns course-level estimate replies.
func (s *Service) ListCourseEstimatesReplies(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course-estimates-replies/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// DeleteCourseEstimateReply soft-deletes a course estimate reply.
func (s *Service) DeleteCourseEstimateReply(ctx context.Context, replyID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course-estimate-reply/%d/delete", replyID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, nil, &result)
	return result, err
}

// DeleteCourseEstimate soft-deletes a course estimate.
func (s *Service) DeleteCourseEstimate(ctx context.Context, estimateID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course-estimate/%d/delete", estimateID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, nil, &result)
	return result, err
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}

func encodeConditions(v any) string {
	if v == nil {
		return ""
	}
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}
