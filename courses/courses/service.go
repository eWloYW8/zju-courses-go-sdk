package courses

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles course-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Course Listing ---

// ListMyCourses returns the current user's courses.
func (s *Service) ListMyCourses(ctx context.Context, opts *model.ListOptions) (*MyCoursesResponse, error) {
	u := addListOptions("/api/my-courses", opts)
	result := new(MyCoursesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListMyCoursesWithFilters returns courses with additional filters.
func (s *Service) ListMyCoursesWithFilters(ctx context.Context, opts *model.ListOptions, params map[string]string) (*MyCoursesResponse, error) {
	u := addListOptions("/api/my-courses", opts)
	u = addQueryParams(u, params)
	result := new(MyCoursesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListMyCoursesByConditions returns the current user's courses using the POST-based query used by the web app.
func (s *Service) ListMyCoursesByConditions(ctx context.Context, body *ListMyCoursesRequest) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/my-courses", body, &result)
	return result, err
}

// GetCourse returns detailed information about a specific course.
func (s *Service) GetCourse(ctx context.Context, courseID int) (*Course, error) {
	return s.GetCourseWithFields(ctx, courseID, "")
}

// GetCourseWithFields returns detailed information about a specific course with an optional field list.
func (s *Service) GetCourseWithFields(ctx context.Context, courseID int, fields string) (*Course, error) {
	u := fmt.Sprintf("/api/courses/%d", courseID)
	if fields != "" {
		u += "?fields=" + url.QueryEscape(fields)
	}
	result := new(Course)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListCourses returns courses for management/search views.
func (s *Service) ListCourses(ctx context.Context, opts *model.ListOptions, body interface{}) (json.RawMessage, error) {
	u := addListOptions("/api/courses", opts)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// --- Course Modules ---

// ListModules returns all modules for a course.
func (s *Service) ListModules(ctx context.Context, courseID int) (*ModulesResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/modules", courseID)
	result := new(ModulesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Course Activities ---

// ListActivities returns all activities for a course.
func (s *Service) ListActivities(ctx context.Context, courseID int) (*ActivitiesResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/activities", courseID)
	result := new(ActivitiesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Enrollments ---

// ListEnrollments returns all enrollments for a course.
func (s *Service) ListEnrollments(ctx context.Context, courseID int) (*EnrollmentsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/enrollments", courseID)
	result := new(EnrollmentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetEnrollmentUser returns enrollment details for a specific user.
func (s *Service) GetEnrollmentUser(ctx context.Context, courseID, userID int) (*EnrollmentDetail, error) {
	return s.GetEnrollmentUserWithFields(ctx, courseID, userID, "")
}

// GetEnrollmentUserWithFields returns enrollment details using the frontend endpoint shape.
func (s *Service) GetEnrollmentUserWithFields(ctx context.Context, courseID, userID int, fields string) (*EnrollmentDetail, error) {
	u := fmt.Sprintf("/api/course/%d/enrollment/%d", courseID, userID)
	if fields != "" {
		u = addQueryParams(u, map[string]string{"fields": fields})
	}
	result := new(EnrollmentDetail)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListInstructors returns instructors for a course.
func (s *Service) ListInstructors(ctx context.Context, courseID int) (*EnrollmentsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/instructors", courseID)
	result := new(EnrollmentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListAllSections returns all sections for a course.
func (s *Service) ListAllSections(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/all-sections", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Navigation & Settings ---

// GetNavSetting returns navigation settings for a course.
func (s *Service) GetNavSetting(ctx context.Context, courseID int) (*NavSettingResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/nav-setting", courseID)
	result := new(NavSettingResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetOutline returns the course outline.
func (s *Service) GetOutline(ctx context.Context, courseID int) (*OutlineResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/outline", courseID)
	result := new(OutlineResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetActivityPublishSetting returns the activity publish settings.
func (s *Service) GetActivityPublishSetting(ctx context.Context, courseID int) (*ActivityPublishSetting, error) {
	u := fmt.Sprintf("/api/course/%d/activity-publish-setting", courseID)
	result := new(ActivityPublishSetting)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Other ---

// GetUsersSmallAvatars returns small avatar URLs for all users in a course.
func (s *Service) GetUsersSmallAvatars(ctx context.Context, courseID int) (*UsersSmallAvatarsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/users-small-avatars", courseID)
	result := new(UsersSmallAvatarsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListLiveRecords returns live records for a course.
func (s *Service) ListLiveRecords(ctx context.Context, courseID int) (*LiveRecordsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/live-record", courseID)
	result := new(LiveRecordsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetCompleteness returns the user's course completion status.
func (s *Service) GetCompleteness(ctx context.Context, courseID int) (*CompletenessResponse, error) {
	u := fmt.Sprintf("/api/course/%d/my-completeness", courseID)
	result := new(CompletenessResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetActivityReadsForUser returns activity read statuses for the current user.
func (s *Service) GetActivityReadsForUser(ctx context.Context, courseID int) ([]*activities.ActivityRead, error) {
	u := fmt.Sprintf("/api/course/%d/activity-reads-for-user", courseID)
	result := new(ActivityReadsForUserResponse)
	_, err := s.client.Get(ctx, u, result)
	return result.ActivityReads, err
}

// GetEntryRecord returns the course entry record.
func (s *Service) GetEntryRecord(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/entry/record", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetOnlineVideoCompletenessSetting returns the online video completeness setting.
func (s *Service) GetOnlineVideoCompletenessSetting(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/online-video-completeness/setting", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Course Management (Instructor) ---

// CreateCourse creates a new course.
func (s *Service) CreateCourse(ctx context.Context, course *Course) (*Course, error) {
	result := new(Course)
	_, err := s.client.Post(ctx, "/api/course", course, result)
	return result, err
}

// UpdateCourse updates a course.
func (s *Service) UpdateCourse(ctx context.Context, courseID int, course *Course) (*Course, error) {
	u := fmt.Sprintf("/api/courses/%d", courseID)
	result := new(Course)
	_, err := s.client.Put(ctx, u, course, result)
	return result, err
}

// DeleteCourse deletes a course.
func (s *Service) DeleteCourse(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/courses/%d", courseID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// UpdateNavSetting updates navigation settings for a course.
func (s *Service) UpdateNavSetting(ctx context.Context, courseID int, settings []*NavSetting) error {
	u := fmt.Sprintf("/api/courses/%d/nav-setting", courseID)
	_, err := s.client.Put(ctx, u, &UpdateNavSettingRequest{NavSetting: settings}, nil)
	return err
}

// CreateModule creates a new module in a course.
func (s *Service) CreateModule(ctx context.Context, courseID int, module *Module) (*Module, error) {
	u := fmt.Sprintf("/api/course/%d/module", courseID)
	result := new(Module)
	_, err := s.client.Post(ctx, u, module, result)
	return result, err
}

// UpdateModule updates a module.
func (s *Service) UpdateModule(ctx context.Context, moduleID int, module *Module) (*Module, error) {
	u := fmt.Sprintf("/api/module/%d", moduleID)
	result := new(Module)
	_, err := s.client.Put(ctx, u, module, result)
	return result, err
}

// DeleteModule deletes a module.
func (s *Service) DeleteModule(ctx context.Context, moduleID int) error {
	return s.DeleteModuleWithOptions(ctx, moduleID, nil)
}

// DeleteModuleWithOptions deletes a module with optional query parameters.
func (s *Service) DeleteModuleWithOptions(ctx context.Context, moduleID int, opts *DeleteModuleOptions) error {
	u := fmt.Sprintf("/api/module/%d", moduleID)
	if opts != nil && opts.DeleteRelatedActivity {
		u += "?delete_related_activity=true"
	}
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// GetModuleHasDependents returns whether a module has dependent activities.
func (s *Service) GetModuleHasDependents(ctx context.Context, moduleID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/modules/%d/has-dependents", moduleID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// UpdateModuleActivitiesSort updates the activity order within a module.
func (s *Service) UpdateModuleActivitiesSort(ctx context.Context, moduleID int, body interface{}) error {
	u := fmt.Sprintf("/api/modules/%d/activity-sort", moduleID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// CreateModuleInteraction creates an interaction activity directly under a module.
func (s *Service) CreateModuleInteraction(ctx context.Context, moduleID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/modules/%d/interaction", moduleID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ResortActivities resorts activities in a course.
func (s *Service) ResortActivities(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/activity-resort", body, nil)
	return err
}

// SyncFromURP syncs courses from URP.
func (s *Service) SyncFromURP(ctx context.Context, courseIDs []int) error {
	_, err := s.client.Post(ctx, "/api/courses/sync_from_urp", map[string][]int{"course_ids": courseIDs}, nil)
	return err
}

// GetCourseCount returns the total number of courses matching filters.
func (s *Service) GetCourseCount(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/courses/count", params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetSettings returns course settings.
func (s *Service) GetSettings(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/courses/settings", &result)
	return result, err
}

// --- Blueprint ---

// GetBlueprintSubItems returns blueprint sub-items for a specific entity.
func (s *Service) GetBlueprintSubItems(ctx context.Context, courseID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/blueprint/%d/sub-items", courseID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetBlueprintSubItemsCount returns blueprint sub-items count.
func (s *Service) GetBlueprintSubItemsCount(ctx context.Context, courseID int) (*BlueprintSubItemsResponse, error) {
	u := fmt.Sprintf("/api/blueprint/%d/sub-items-count", courseID)
	result := new(BlueprintSubItemsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// SyncBlueprint syncs blueprint content to target courses.
func (s *Service) SyncBlueprint(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blueprint/%d/sync", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ListBlueprintSubCourses returns blueprint sub courses.
func (s *Service) ListBlueprintSubCourses(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blueprint/sub-courses/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// DeleteBlueprint deletes a blueprint course mapping/configuration.
func (s *Service) DeleteBlueprint(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/blueprint/%d", courseID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// CancelBlueprintActivitySync cancels blueprint sync for an activity.
func (s *Service) CancelBlueprintActivitySync(ctx context.Context, courseID, activityID int, body interface{}) error {
	u := fmt.Sprintf("/api/blueprint/%d/activities/%d/cancel-sync", courseID, activityID)
	_, err := s.client.Delete(ctx, u, body)
	return err
}

// GetBlueprintSubmittedInfo returns blueprint submitted sync info for a target object.
func (s *Service) GetBlueprintSubmittedInfo(ctx context.Context, courseID int, resourceType string, resourceID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blueprint/%d/%s/%d/submitted-info", courseID, resourceType, resourceID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// SyncBlueprintSubject syncs a blueprint subject item.
func (s *Service) SyncBlueprintSubject(ctx context.Context, courseID int, resourceType string, resourceID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blueprint/%d/%s/%d/sync-subject", courseID, resourceType, resourceID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, nil, &result)
	return result, err
}

// --- Danmu (Bullet Screen) ---

// GetDanmu returns danmu for a course.
func (s *Service) GetDanmu(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/danmu/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Statistic Resource Audit ---

// GetStatisticResourceAudit returns statistic resource audit.
func (s *Service) GetStatisticResourceAudit(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/courses/statistic/resource-audit", &result)
	return result, err
}

// --- TPDOE ---

// GetTPDOEStatStudents returns TPDOE student statistics.
func (s *Service) GetTPDOEStatStudents(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/tpdoe/stat-students?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Inspect Child ---

// InspectChild inspects a child course.
func (s *Service) InspectChild(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/inspect-child/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Course Classifications ---

// ListCourseClassifications returns course classifications.
func (s *Service) ListCourseClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/course-classifications", &result)
	return result, err
}

// --- Curriculum Classifications ---

// ListCurriculumClassifications returns curriculum classifications.
func (s *Service) ListCurriculumClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/curriculum-classifications", &result)
	return result, err
}

// --- Sign In ---

// GetCourseSignIn returns the sign-in for a course.
func (s *Service) GetCourseSignIn(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/sign-in", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Current Semester Courses ---

// GetCurrentSemesterCourses returns courses for the current semester.
func (s *Service) GetCurrentSemesterCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/current-semester-courses", &result)
	return result, err
}

// --- Tencent Meeting Activities ---

// ListTencentMeetingActivities returns Tencent meeting activities for a course.
func (s *Service) ListTencentMeetingActivities(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/tencent-meeting/activities?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Lecture Live Activity ---

// GetLectureLiveActivity returns a lecture live activity for a course.
func (s *Service) GetLectureLiveActivity(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/lecture-live-activity/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Preview ---

// CancelPreview cancels a course preview.
func (s *Service) CancelPreview(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/courses/%d/preview", courseID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// --- Helpers ---

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}
