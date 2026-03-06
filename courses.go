package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// CoursesService handles course-related API operations.
type CoursesService struct {
	client *Client
}

// --- List Responses ---

type MyCoursesResponse struct {
	Courses []*model.Course `json:"courses"`
	model.Pagination
}

type ModulesResponse struct {
	Modules []*model.Module `json:"modules"`
}

type EnrollmentsResponse struct {
	Enrollments []*model.Enrollment `json:"enrollments"`
}

type ActivitiesResponse struct {
	Activities []*model.Activity `json:"activities"`
}

type NavSettingResponse struct {
	NavSetting          []*model.NavSetting `json:"nav_setting"`
	CanNotDisabledItems []string            `json:"can_not_disabled_items"`
}

type OutlineResponse = model.Outline

type HomeworkScoresResponse struct {
	HomeworkActivities []*model.HomeworkActivity `json:"homework_activities"`
	Scores             []*model.HomeworkScore    `json:"scores"`
}

type HomeworkSubmissionStatusResponse struct {
	CourseID           int                               `json:"course_id"`
	HomeworkActivities []*model.HomeworkSubmissionStatus `json:"homework_activities"`
}

type ExamsResponse struct {
	Exams []*model.Exam `json:"exams"`
}

type ExamScoresResponse struct {
	ExamScores []*model.ExamScore `json:"exam_scores"`
}

type ClassroomListResponse struct {
	Classrooms []*model.Classroom `json:"classrooms"`
}

type TopicCategoriesResponse struct {
	TopicCategories []*model.TopicCategory `json:"topic_categories"`
}

type UsersSmallAvatarsResponse struct {
	Avatars map[string]string `json:"avatars"`
}

type InteractionsResponse struct {
	Interactions []*model.Interaction `json:"interactions"`
}

type LiveRecordsResponse struct {
	LiveRecords []*model.LiveRecord `json:"live_records"`
}

type RollcallsResponse struct {
	Rollcalls []*model.Rollcall `json:"rollcalls"`
}

type SubmittedExamsResponse struct {
	ExamIDs []int `json:"exam_ids"`
}

// --- Course Listing ---

// ListMyCourses returns the current user's courses.
func (s *CoursesService) ListMyCourses(ctx context.Context, opts *model.ListOptions) (*MyCoursesResponse, error) {
	u := addListOptions("/api/my-courses", opts)
	result := new(MyCoursesResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListMyCoursesWithFilters returns courses with additional filters.
func (s *CoursesService) ListMyCoursesWithFilters(ctx context.Context, opts *model.ListOptions, params map[string]string) (*MyCoursesResponse, error) {
	u := addListOptions("/api/my-courses", opts)
	u = addQueryParams(u, params)
	result := new(MyCoursesResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetCourse returns detailed information about a specific course.
func (s *CoursesService) GetCourse(ctx context.Context, courseID int) (*model.Course, error) {
	u := fmt.Sprintf("/api/courses/%d", courseID)
	result := new(model.Course)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Course Modules ---

// ListModules returns all modules for a course.
func (s *CoursesService) ListModules(ctx context.Context, courseID int) (*ModulesResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/modules", courseID)
	result := new(ModulesResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Course Activities ---

// ListActivities returns all activities for a course.
func (s *CoursesService) ListActivities(ctx context.Context, courseID int) (*ActivitiesResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/activities", courseID)
	result := new(ActivitiesResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Enrollments ---

// ListEnrollments returns all enrollments for a course.
func (s *CoursesService) ListEnrollments(ctx context.Context, courseID int) (*EnrollmentsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/enrollments", courseID)
	result := new(EnrollmentsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetEnrollmentUser returns enrollment details for a specific user.
func (s *CoursesService) GetEnrollmentUser(ctx context.Context, courseID, userID int) (*model.EnrollmentDetail, error) {
	u := fmt.Sprintf("/api/courses/%d/enrollments/users/%d", courseID, userID)
	result := new(model.EnrollmentDetail)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListInstructors returns instructors for a course.
func (s *CoursesService) ListInstructors(ctx context.Context, courseID int) (*EnrollmentsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/instructors", courseID)
	result := new(EnrollmentsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Navigation & Settings ---

// GetNavSetting returns navigation settings for a course.
func (s *CoursesService) GetNavSetting(ctx context.Context, courseID int) (*NavSettingResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/nav-setting", courseID)
	result := new(NavSettingResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetOutline returns the course outline.
func (s *CoursesService) GetOutline(ctx context.Context, courseID int) (*OutlineResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/outline", courseID)
	result := new(OutlineResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetActivityPublishSetting returns the activity publish settings.
func (s *CoursesService) GetActivityPublishSetting(ctx context.Context, courseID int) (*model.ActivityPublishSetting, error) {
	u := fmt.Sprintf("/api/course/%d/activity-publish-setting", courseID)
	result := new(model.ActivityPublishSetting)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Scores ---

// ListHomeworkScores returns homework score entries for a course.
func (s *CoursesService) ListHomeworkScores(ctx context.Context, courseID int) (*HomeworkScoresResponse, error) {
	u := fmt.Sprintf("/api/course/%d/homework-scores", courseID)
	result := new(HomeworkScoresResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetHomeworkSubmissionStatus returns homework submission statuses for a course.
func (s *CoursesService) GetHomeworkSubmissionStatus(ctx context.Context, courseID int) (*HomeworkSubmissionStatusResponse, error) {
	u := fmt.Sprintf("/api/course/%d/homework/submission-status", courseID)
	result := new(HomeworkSubmissionStatusResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListExamScores returns exam scores for a course.
func (s *CoursesService) ListExamScores(ctx context.Context, courseID int) (*ExamScoresResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/exam-scores", courseID)
	result := new(ExamScoresResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Exams & Classroom ---

// ListExams returns all exams for a course.
func (s *CoursesService) ListExams(ctx context.Context, courseID int) (*ExamsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/exams", courseID)
	result := new(ExamsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListClassrooms returns classroom activities for a course.
func (s *CoursesService) ListClassrooms(ctx context.Context, courseID int) (*ClassroomListResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/classroom-list", courseID)
	result := new(ClassroomListResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListSubmittedExams returns IDs of submitted exams.
func (s *CoursesService) ListSubmittedExams(ctx context.Context, courseID int) (*SubmittedExamsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/submitted-exams", courseID)
	result := new(SubmittedExamsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Forum ---

// ListTopicCategories returns topic categories for a course.
func (s *CoursesService) ListTopicCategories(ctx context.Context, courseID int) (*TopicCategoriesResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/topic-categories", courseID)
	result := new(TopicCategoriesResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Other ---

// GetUsersSmallAvatars returns small avatar URLs for all users in a course.
func (s *CoursesService) GetUsersSmallAvatars(ctx context.Context, courseID int) (*UsersSmallAvatarsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/users-small-avatars", courseID)
	result := new(UsersSmallAvatarsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListInteractions returns interactions for a course.
func (s *CoursesService) ListInteractions(ctx context.Context, courseID int) (*InteractionsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/interactions", courseID)
	result := new(InteractionsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListLiveRecords returns live records for a course.
func (s *CoursesService) ListLiveRecords(ctx context.Context, courseID int) (*LiveRecordsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/live-record", courseID)
	result := new(LiveRecordsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListRollcalls returns rollcall records for a course.
func (s *CoursesService) ListRollcalls(ctx context.Context, courseID int) (*RollcallsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/modules/rollcalls", courseID)
	result := new(RollcallsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetCompleteness returns the user's course completion status.
func (s *CoursesService) GetCompleteness(ctx context.Context, courseID int) (*model.CompletenessResponse, error) {
	u := fmt.Sprintf("/api/course/%d/my-completeness", courseID)
	result := new(model.CompletenessResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetActivityReadsForUser returns activity read statuses for the current user.
func (s *CoursesService) GetActivityReadsForUser(ctx context.Context, courseID int) ([]*model.ActivityRead, error) {
	u := fmt.Sprintf("/api/course/%d/activity-reads-for-user", courseID)
	var wrapper struct {
		ActivityReads []*model.ActivityRead `json:"activity_reads"`
	}
	_, err := s.client.get(ctx, u, &wrapper)
	return wrapper.ActivityReads, err
}

// GetEntryRecord returns the course entry record.
func (s *CoursesService) GetEntryRecord(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/entry/record", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetOnlineVideoCompletenessSetting returns the online video completeness setting.
func (s *CoursesService) GetOnlineVideoCompletenessSetting(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/online-video-completeness/setting", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Course Management (Instructor) ---

// CreateCourse creates a new course.
func (s *CoursesService) CreateCourse(ctx context.Context, course *model.Course) (*model.Course, error) {
	result := new(model.Course)
	_, err := s.client.post(ctx, "/api/courses", course, result)
	return result, err
}

// UpdateCourse updates a course.
func (s *CoursesService) UpdateCourse(ctx context.Context, courseID int, course *model.Course) (*model.Course, error) {
	u := fmt.Sprintf("/api/courses/%d", courseID)
	result := new(model.Course)
	_, err := s.client.put(ctx, u, course, result)
	return result, err
}

// DeleteCourse deletes a course.
func (s *CoursesService) DeleteCourse(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/courses/%d", courseID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// UpdateNavSetting updates navigation settings for a course.
func (s *CoursesService) UpdateNavSetting(ctx context.Context, courseID int, settings []*model.NavSetting) error {
	u := fmt.Sprintf("/api/courses/%d/nav-setting", courseID)
	_, err := s.client.put(ctx, u, map[string][]*model.NavSetting{"nav_setting": settings}, nil)
	return err
}

// CreateModule creates a new module in a course.
func (s *CoursesService) CreateModule(ctx context.Context, courseID int, module *model.Module) (*model.Module, error) {
	u := fmt.Sprintf("/api/courses/%d/modules", courseID)
	result := new(model.Module)
	_, err := s.client.post(ctx, u, module, result)
	return result, err
}

// UpdateModule updates a module.
func (s *CoursesService) UpdateModule(ctx context.Context, moduleID int, module *model.Module) (*model.Module, error) {
	u := fmt.Sprintf("/api/module/%d", moduleID)
	result := new(model.Module)
	_, err := s.client.put(ctx, u, module, result)
	return result, err
}

// DeleteModule deletes a module.
func (s *CoursesService) DeleteModule(ctx context.Context, moduleID int) error {
	u := fmt.Sprintf("/api/module/%d", moduleID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// ResortActivities resorts activities in a course.
func (s *CoursesService) ResortActivities(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/activity-resort", body, nil)
	return err
}

// SyncFromURP syncs courses from URP.
func (s *CoursesService) SyncFromURP(ctx context.Context, courseIDs []int) error {
	_, err := s.client.post(ctx, "/api/courses/sync_from_urp", map[string][]int{"course_ids": courseIDs}, nil)
	return err
}

// GetCourseCount returns the total number of courses matching filters.
func (s *CoursesService) GetCourseCount(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/courses/count", params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetSettings returns course settings.
func (s *CoursesService) GetSettings(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/courses/settings", &result)
	return result, err
}

// GetBlueprintSubItemsCount returns blueprint sub-items count.
func (s *CoursesService) GetBlueprintSubItemsCount(ctx context.Context, courseID int) (*model.BlueprintSubItemsResponse, error) {
	u := fmt.Sprintf("/api/blueprint/%d/sub-items-count", courseID)
	result := new(model.BlueprintSubItemsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}
