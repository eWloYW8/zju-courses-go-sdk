package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"net/url"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// ResourceGroupFoldersResponse represents the response from ListResourceGroupFolders.
type ResourceGroupFoldersResponse struct {
	Folders    []interface{}    `json:"folders"`
	Pagination model.Pagination `json:"pagination"`
}

// ResourceGroupResourcesResponse represents the response from ListResourceGroupResources.
type ResourceGroupResourcesResponse struct {
	Resources []interface{} `json:"resources"`
	Page      int           `json:"page"`
	Pages     int           `json:"pages"`
	Total     int           `json:"total"`
}

// Service handles resource management API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Resource Groups ---

// ListResourceGroups returns resource groups.
func (s *Service) ListResourceGroups(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/resource-groups", opts)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetResourceGroup returns a specific resource group.
func (s *Service) GetResourceGroup(ctx context.Context, groupID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/resource-groups/%d", groupID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateResourceGroup creates a new resource group.
func (s *Service) CreateResourceGroup(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/resource-group", body, &result)
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
func (s *Service) ListPagedResourceGroups(ctx context.Context, opts *model.ListOptions, body interface{}) (json.RawMessage, error) {
	u := addListOptions("/api/resource-groups", opts)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// GetResourceGroupMembers returns members of a resource group.
func (s *Service) GetResourceGroupMembers(ctx context.Context, groupID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/resource-groups/%d/members", groupID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// LeaveResourceGroup leaves a resource group.
func (s *Service) LeaveResourceGroup(ctx context.Context, groupID int) error {
	u := fmt.Sprintf("/api/resource-groups/%d/leave", groupID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// ListResourceGroupFolders returns folders in a resource group.
func (s *Service) ListResourceGroupFolders(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/resource-groups/folders", &result)
	return result, err
}

// ListResourceGroupResources returns resources in a resource group.
func (s *Service) ListResourceGroupResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/resource-groups/resources", &result)
	return result, err
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
func (s *Service) ListSharedResourcesToMe(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/shared-resources-to-me", &result)
	return result, err
}

// ListSharedResources returns shared resources with pagination.
func (s *Service) ListSharedResources(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/shared-resources-no-repeated", opts)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetSharedResource returns a specific shared resource.
func (s *Service) GetSharedResource(ctx context.Context, resourceID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/shared-resources/%d", resourceID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateSharedResource creates a shared resource.
func (s *Service) CreateSharedResource(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/shared-resources", body, &result)
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

// GetUserSharedResources returns shared resources for a user.
func (s *Service) GetUserSharedResources(ctx context.Context, userID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/shared-resources/user/%d", userID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListSharedResourcesFromMe returns shared resources created by the current user.
func (s *Service) ListSharedResourcesFromMe(ctx context.Context, opts *model.ListOptions, conditions string) (json.RawMessage, error) {
	u := addListOptions("/api/shared-resources/from-me", opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListMostLikedSharedResources returns most liked shared resources.
func (s *Service) ListMostLikedSharedResources(ctx context.Context, conditions string) (json.RawMessage, error) {
	u := "/api/shared-resources/most-liked"
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListRecentUsedSharedResources returns recently used shared resources.
func (s *Service) ListRecentUsedSharedResources(ctx context.Context, classificationID string, departmentIDs string) (json.RawMessage, error) {
	params := map[string]string{}
	if classificationID != "" {
		params["classificationId"] = classificationID
	}
	if departmentIDs != "" {
		params["departmentIds"] = departmentIDs
	}
	u := addQueryParams("/api/shared-resources/recent-used", params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
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
func (s *Service) ListUserCollections(ctx context.Context, userID int, opts *model.ListOptions, conditions string) (json.RawMessage, error) {
	u := addListOptions(fmt.Sprintf("/api/shared-resources/user/%d/collections", userID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
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

// CopyThirdPartResources copies third-party resources.
func (s *Service) CopyThirdPartResources(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/copy-third-part-resources", body, nil)
	return err
}

// --- Resource Files ---

// GetResourceFile returns a resource file.
func (s *Service) GetResourceFile(ctx context.Context, fileID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/resource-file/%d", fileID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
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
func (s *Service) ListCoursePackages(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/course-packages", &result)
	return result, err
}

// GetCoursePackage returns a specific course package.
func (s *Service) GetCoursePackage(ctx context.Context, packageID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course-packages/%d", packageID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Public Courses ---

// ListPublicCourses returns public courses.
func (s *Service) ListPublicCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/courses/public", &result)
	return result, err
}

// ListHotPublicCourses returns hot public courses.
func (s *Service) ListHotPublicCourses(ctx context.Context, limit int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/public/hot?limit=%d", limit)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListLatestPublicCourses returns latest public courses.
func (s *Service) ListLatestPublicCourses(ctx context.Context, limit int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/public/latest?limit=%d", limit)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
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

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}
