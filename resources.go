package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

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

// ResourcesService handles resource management API operations.
type ResourcesService struct {
	client *Client
}

// --- Resource Groups ---

// ListResourceGroups returns resource groups.
func (s *ResourcesService) ListResourceGroups(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/resource-groups", opts)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetResourceGroup returns a specific resource group.
func (s *ResourcesService) GetResourceGroup(ctx context.Context, groupID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/resource-groups/%d", groupID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// CreateResourceGroup creates a new resource group.
func (s *ResourcesService) CreateResourceGroup(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/resource-group", body, &result)
	return result, err
}

// UpdateResourceGroup updates a resource group.
func (s *ResourcesService) UpdateResourceGroup(ctx context.Context, groupID int, body interface{}) error {
	u := fmt.Sprintf("/api/resource-group/%d", groupID)
	_, err := s.client.put(ctx, u, body, nil)
	return err
}

// DeleteResourceGroup deletes a resource group.
func (s *ResourcesService) DeleteResourceGroup(ctx context.Context, groupID int) error {
	u := fmt.Sprintf("/api/resource-groups/%d", groupID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// ListResourceGroupFolders returns folders in a resource group.
func (s *ResourcesService) ListResourceGroupFolders(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/resource-groups/folders", &result)
	return result, err
}

// ListResourceGroupResources returns resources in a resource group.
func (s *ResourcesService) ListResourceGroupResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/resource-groups/resources", &result)
	return result, err
}

// --- Resource Folders ---

// ListResourceFolders returns resource folders.
func (s *ResourcesService) ListResourceFolders(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/resource/folders", &result)
	return result, err
}

// --- Public Resources ---

// ListPublicResources returns public resources.
func (s *ResourcesService) ListPublicResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/public-resources", &result)
	return result, err
}

// --- Shared Resources ---

// ListSharedResourcesToMe returns resources shared to the current user.
func (s *ResourcesService) ListSharedResourcesToMe(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/shared-resources-to-me", &result)
	return result, err
}

// ListSharedResources returns shared resources with pagination.
func (s *ResourcesService) ListSharedResources(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/shared-resources-no-repeated", opts)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetSharedResource returns a specific shared resource.
func (s *ResourcesService) GetSharedResource(ctx context.Context, resourceID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/shared-resources/%d", resourceID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// CreateSharedResource creates a shared resource.
func (s *ResourcesService) CreateSharedResource(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/shared-resources", body, &result)
	return result, err
}

// BatchSaveSharedResources batch saves shared resources.
func (s *ResourcesService) BatchSaveSharedResources(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/shared-resources/batch-save", body, nil)
	return err
}

// GetSharedResourceStats returns shared resource statistics.
func (s *ResourcesService) GetSharedResourceStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/shared-resources/stat", &result)
	return result, err
}

// GetSharedVideoResourceStats returns shared video resource statistics.
func (s *ResourcesService) GetSharedVideoResourceStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/shared-resources/stat/video-resources", &result)
	return result, err
}

// ExportSharedVideoResourceStats exports shared video resource statistics.
func (s *ResourcesService) ExportSharedVideoResourceStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/shared-resources/stat/video-resources/export", &result)
	return result, err
}

// GetSharedResourceClassifications returns shared resource classifications.
func (s *ResourcesService) GetSharedResourceClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/shared-resource/classifications", &result)
	return result, err
}

// ListSharedResourceManagement returns shared resource management list.
func (s *ResourcesService) ListSharedResourceManagement(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/shared-resources/management", &result)
	return result, err
}

// ShareToOtherOrgs shares resources to other organizations (admin).
func (s *ResourcesService) ShareToOtherOrgs(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/shared-resources/admin/to-other-orgs", body, nil)
	return err
}

// GetUserSharedResources returns shared resources for a user.
func (s *ResourcesService) GetUserSharedResources(ctx context.Context, userID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/shared-resources/user/%d", userID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Save Resources ---

// SaveResources saves resources.
func (s *ResourcesService) SaveResources(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/save-resources", body, nil)
	return err
}

// CheckSaveResources checks if resources can be saved.
func (s *ResourcesService) CheckSaveResources(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/save-resources/check", body, &result)
	return result, err
}

// CopyThirdPartResources copies third-party resources.
func (s *ResourcesService) CopyThirdPartResources(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/copy-third-part-resources", body, nil)
	return err
}

// --- Resource Files ---

// GetResourceFile returns a resource file.
func (s *ResourcesService) GetResourceFile(ctx context.Context, fileID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/resource-file/%d", fileID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetOrgResourceFileArrears returns org resource file arrears.
func (s *ResourcesService) GetOrgResourceFileArrears(ctx context.Context, orgID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/resource-file/org/arrears/%d", orgID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Lesson Resources ---

// GetLessonResourcesSharedStat returns shared statistics for lesson resources.
func (s *ResourcesService) GetLessonResourcesSharedStat(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/lesson-resources/shared-stat", &result)
	return result, err
}

// --- Virtual Experiments ---

// GetVirtualExperiment returns a virtual experiment.
func (s *ResourcesService) GetVirtualExperiment(ctx context.Context, experimentID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/virtual-experiments/%d", experimentID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Virtual Classroom Resources ---

// ListVirtualClassroomResources returns virtual classroom resources.
func (s *ResourcesService) ListVirtualClassroomResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/virtual-classroom-resources", &result)
	return result, err
}

// --- CC License ---

// GetCCLicenseGroups returns Creative Commons license groups.
func (s *ResourcesService) GetCCLicenseGroups(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/cc-license/groups", &result)
	return result, err
}

// GetCCLicenseMap returns Creative Commons license map.
func (s *ResourcesService) GetCCLicenseMap(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/cc-license/map", &result)
	return result, err
}

// --- Slides ---

// ListPublishedSlides returns published slides.
func (s *ResourcesService) ListPublishedSlides(ctx context.Context, fields string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/slides/published?fields=%s", fields)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// SearchSlides searches slides by keyword.
func (s *ResourcesService) SearchSlides(ctx context.Context, keyword string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/slides?keyword=%s", keyword)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Video Suites ---

// GetVideoSuite returns a video suite.
func (s *ResourcesService) GetVideoSuite(ctx context.Context, suiteID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/video-suites/%d", suiteID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListVideoSuiteComments returns comments for a video suite.
func (s *ResourcesService) ListVideoSuiteComments(ctx context.Context, suiteID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/video-suite/comments/%d", suiteID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Video Quizzes ---

// GetVideoQuiz returns a video quiz.
func (s *ResourcesService) GetVideoQuiz(ctx context.Context, quizID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/video-quizzes/%d", quizID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetVideoQuizOrgArrears returns video quiz org arrears.
func (s *ResourcesService) GetVideoQuizOrgArrears(ctx context.Context, orgID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/video-quizzes/org/arrears/%d", orgID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Certifications ---

// ListCertifications returns certifications.
func (s *ResourcesService) ListCertifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/certifications", &result)
	return result, err
}

// --- Course Packages ---

// ListCoursePackages returns course packages.
func (s *ResourcesService) ListCoursePackages(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/course-packages", &result)
	return result, err
}

// GetCoursePackage returns a specific course package.
func (s *ResourcesService) GetCoursePackage(ctx context.Context, packageID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course-packages/%d", packageID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Public Courses ---

// ListPublicCourses returns public courses.
func (s *ResourcesService) ListPublicCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/courses/public", &result)
	return result, err
}

// ListHotPublicCourses returns hot public courses.
func (s *ResourcesService) ListHotPublicCourses(ctx context.Context, limit int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/public/hot?limit=%d", limit)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListLatestPublicCourses returns latest public courses.
func (s *ResourcesService) ListLatestPublicCourses(ctx context.Context, limit int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/public/latest?limit=%d", limit)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Course Estimates ---

// GetCourseEstimate returns a course estimate.
func (s *ResourcesService) GetCourseEstimate(ctx context.Context, estimateID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course-estimate/%d", estimateID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListCourseEstimates returns course estimates.
func (s *ResourcesService) ListCourseEstimates(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course-estimates/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// CreateCourseEstimate creates a course estimate.
func (s *ResourcesService) CreateCourseEstimate(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/course-estimate", body, &result)
	return result, err
}

// CreateCourseEstimateReply creates a reply to a course estimate.
func (s *ResourcesService) CreateCourseEstimateReply(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/course-estimate-reply", body, &result)
	return result, err
}
