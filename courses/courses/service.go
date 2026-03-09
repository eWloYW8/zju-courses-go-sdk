package courses

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

// ListMyCoursesByConditionsTyped returns the current user's courses using the POST-based query with typed results.
func (s *Service) ListMyCoursesByConditionsTyped(ctx context.Context, body *ListMyCoursesRequest) (*MyCoursesResponse, error) {
	result := new(MyCoursesResponse)
	_, err := s.client.Post(ctx, "/api/my-courses", body, result)
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
	return s.GetEnrollmentUserWithParams(ctx, courseID, userID, GetEnrollmentUserParams{})
}

// GetEnrollmentUserWithFields returns enrollment details using the frontend endpoint shape.
func (s *Service) GetEnrollmentUserWithFields(ctx context.Context, courseID, userID int, fields string) (*EnrollmentDetail, error) {
	return s.GetEnrollmentUserWithParams(ctx, courseID, userID, GetEnrollmentUserParams{Fields: fields})
}

// GetEnrollmentUserWithParams returns enrollment details using the frontend endpoint shape and optional request scope.
func (s *Service) GetEnrollmentUserWithParams(ctx context.Context, courseID, userID int, params GetEnrollmentUserParams) (*EnrollmentDetail, error) {
	u := fmt.Sprintf("/api/courses/%d/enrollments/users/%d", courseID, userID)
	if params.Fields != "" {
		u = addQueryParams(u, map[string]string{"fields": params.Fields})
	}
	result := new(EnrollmentDetail)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	if params.RequestScope != "" {
		req.Header.Set("Request-Scope", params.RequestScope)
	}
	_, err = s.client.Do(req, result)
	return result, err
}

// ListInstructors returns instructor enrollments for a course.
func (s *Service) ListInstructors(ctx context.Context, courseID int) (*EnrollmentsResponse, error) {
	return s.ListInstructorEnrollments(ctx, courseID, ListInstructorEnrollmentsParams{})
}

// ListInstructorEnrollments returns instructor enrollments for a course with optional field selection.
func (s *Service) ListInstructorEnrollments(ctx context.Context, courseID int, params ListInstructorEnrollmentsParams) (*EnrollmentsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/instructors", courseID)
	if params.Fields != "" {
		u = addQueryParams(u, map[string]string{"fields": params.Fields})
	}
	result := new(EnrollmentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListCourseInstructors returns the flattened instructor user list for a course.
func (s *Service) ListCourseInstructors(ctx context.Context, courseID int) (*CourseInstructorsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/instructors", courseID)
	result := new(CourseInstructorsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListEducators returns educators/enrollments for a course using the management endpoint.
func (s *Service) ListEducators(ctx context.Context, courseID int, fields string) (*EnrollmentsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/educators", courseID)
	if fields != "" {
		u = addQueryParams(u, map[string]string{"fields": fields})
	}
	result := new(EnrollmentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListHosts returns host users for a course.
func (s *Service) ListHosts(ctx context.Context, courseID int, params ListCourseHostsParams) (*CourseHostsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/hosts", courseID)
	if params.Type != "" {
		u = addQueryParams(u, map[string]string{"type": params.Type})
	}
	result := new(CourseHostsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListGroupUsers returns users by group name for a course.
func (s *Service) ListGroupUsers(ctx context.Context, courseID int, groupName string) (*CourseGroupUsersResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/group-users", courseID)
	if groupName != "" {
		u = addQueryParams(u, map[string]string{"group_name": groupName})
	}
	result := new(CourseGroupUsersResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListAvailableUsers returns users that can be added to the course teaching team.
func (s *Service) ListAvailableUsers(ctx context.Context, courseID int, keyword string, withoutStudent bool) (*AvailableUsersResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/available-users", courseID)
	params := map[string]string{}
	if keyword != "" {
		params["conditions"] = encodeConditions(map[string]string{"keyword": keyword})
	}
	if withoutStudent {
		params["without_student"] = "true"
	}
	u = addQueryParams(u, params)
	result := new(AvailableUsersResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListStudents returns course students with optional avatar suppression.
func (s *Service) ListStudents(ctx context.Context, courseID int, params ListStudentsParams) (*StudentsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/students", courseID)
	query := map[string]string{}
	if params.IgnoreAvatar {
		query["ignore_avatar"] = "true"
	}
	u = addQueryParams(u, query)
	result := new(StudentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListAllSections returns all sections for a course.
func (s *Service) ListAllSections(ctx context.Context, courseID int) (*SectionsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/all-sections", courseID)
	result := new(SectionsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListEntryRefers returns course-scoped entry references used by the rich-text editor.
func (s *Service) ListEntryRefers(ctx context.Context, courseID int) (*EntryRefersResponse, error) {
	u := fmt.Sprintf("/api/course/%d/entry-refers", courseID)
	result := new(EntryRefersResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListScoreItemGroups returns score item groups for a course.
func (s *Service) ListScoreItemGroups(ctx context.Context, courseID int, withoutScoreItem bool) (*ScoreItemGroupsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/score-item-groups", courseID)
	if withoutScoreItem {
		u = addQueryParams(u, map[string]string{"without_score_item": "true"})
	}
	result := new(ScoreItemGroupsResponse)
	_, err := s.client.Get(ctx, u, result)
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

// CreateOrUpdateOutlineItem creates or updates a course outline item.
func (s *Service) CreateOrUpdateOutlineItem(ctx context.Context, courseOutlineID int, body *UpsertOutlineItemRequest) (*OutlineItemResponse, error) {
	u := fmt.Sprintf("/api/course/%d/outline-item", courseOutlineID)
	result := new(OutlineItemResponse)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// DeleteOutlineItem deletes a course outline item.
func (s *Service) DeleteOutlineItem(ctx context.Context, courseOutlineID int, itemID int) error {
	u := fmt.Sprintf("/api/course/%d/outline-item/%d", courseOutlineID, itemID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// GetOutlineDownloadPrintPermissions returns outline download/print permissions for a given type.
func (s *Service) GetOutlineDownloadPrintPermissions(ctx context.Context, courseID int, outlineType string) (*OutlineDownloadPrintPermissions, error) {
	u := fmt.Sprintf("/api/courses/%d/download-print-permissions", courseID)
	if outlineType != "" {
		u = addQueryParams(u, map[string]string{"type": outlineType})
	}
	result := new(OutlineDownloadPrintPermissions)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateOutlineDownloadPrintPermissions updates outline download/print permissions.
func (s *Service) UpdateOutlineDownloadPrintPermissions(ctx context.Context, courseID int, body *OutlineDownloadPrintPermissions) error {
	u := fmt.Sprintf("/api/courses/%d/download-print-permissions", courseID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// GetActivityPublishSetting returns the activity publish settings.
func (s *Service) GetActivityPublishSetting(ctx context.Context, courseID int) (*ActivityPublishSetting, error) {
	u := fmt.Sprintf("/api/course/%d/activity-publish-setting", courseID)
	result := new(ActivityPublishSetting)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetTemplateSetting returns the course template setting.
func (s *Service) GetTemplateSetting(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/template-setting", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
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

// CreateLiveRecordAsrTask starts an ASR task for a live record.
func (s *Service) CreateLiveRecordAsrTask(ctx context.Context, liveRecordID int) (LiveRecordAsrTaskResponse, error) {
	u := fmt.Sprintf("/api/live-records/%d/task/asr", liveRecordID)
	result := make(LiveRecordAsrTaskResponse)
	_, err := s.client.Post(ctx, u, nil, &result)
	return result, err
}

// GetLiveRecordAsrTaskStatus returns ASR task status with the frontend no-intercept query.
func (s *Service) GetLiveRecordAsrTaskStatus(ctx context.Context, liveRecordID int) (LiveRecordAsrTaskStatusResponse, error) {
	u := fmt.Sprintf("/api/live-records/%d/task/asr?no-intercept=true", liveRecordID)
	result := make(LiveRecordAsrTaskStatusResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetLiveRecordCaptions returns captions for a live record.
func (s *Service) GetLiveRecordCaptions(ctx context.Context, liveRecordID int) (LiveRecordCaptionsResponse, error) {
	u := fmt.Sprintf("/api/live-records/%d/caption", liveRecordID)
	result := make(LiveRecordCaptionsResponse)
	_, err := s.client.Get(ctx, u, &result)
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

// RecordEntry posts the course entry record used by the frontend route tracker.
func (s *Service) RecordEntry(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/entry/record", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, nil, &result)
	return result, err
}

// GetEntryRecord is kept for compatibility and uses the frontend POST endpoint.
func (s *Service) GetEntryRecord(ctx context.Context, courseID int) (json.RawMessage, error) {
	return s.RecordEntry(ctx, courseID)
}

// GetOnlineVideoCompletenessSetting returns the online video completeness setting.
func (s *Service) GetOnlineVideoCompletenessSetting(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/online-video-completeness/setting?no-loading-animation=true", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetOnlineVideoCompletenessSettingDetail returns typed online-video completeness settings.
func (s *Service) GetOnlineVideoCompletenessSettingDetail(ctx context.Context, courseID int) (*OnlineVideoCompletenessSettingResponse, error) {
	u := fmt.Sprintf("/api/course/%d/online-video-completeness/setting?no-loading-animation=true", courseID)
	result := new(OnlineVideoCompletenessSettingResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateOnlineVideoCompletenessSetting updates online-video completeness settings.
func (s *Service) UpdateOnlineVideoCompletenessSetting(ctx context.Context, courseID int, body interface{}) (*OnlineVideoCompletenessSettingResponse, error) {
	u := fmt.Sprintf("/api/course/%d/online-video-completeness/setting", courseID)
	result := new(OnlineVideoCompletenessSettingResponse)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteOnlineVideoCompletenessSetting deletes online-video completeness settings.
func (s *Service) DeleteOnlineVideoCompletenessSetting(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/course/%d/online-video-completeness/setting", courseID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// GetOnlineVideoCompletenessScores returns course-level online-video completeness scores.
func (s *Service) GetOnlineVideoCompletenessScores(ctx context.Context, courseID int) (*OnlineVideoCompletenessScoresResponse, error) {
	u := fmt.Sprintf("/api/course/%d/online-video-completeness/scores", courseID)
	result := new(OnlineVideoCompletenessScoresResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetOnlineVideoCompletenessScore returns the current student's online-video completeness score.
func (s *Service) GetOnlineVideoCompletenessScore(ctx context.Context, courseID int) (*OnlineVideoCompletenessScoreResponse, error) {
	u := fmt.Sprintf("/api/course/%d/online-video-completeness/score", courseID)
	result := new(OnlineVideoCompletenessScoreResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetPerformanceScoreSetting returns course performance score settings.
func (s *Service) GetPerformanceScoreSetting(ctx context.Context, courseID int) (*PerformanceScoreSettingResponse, error) {
	u := fmt.Sprintf("/api/course/%d/performance/score-setting", courseID)
	result := new(PerformanceScoreSettingResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdatePerformanceScoreSetting updates course performance score settings.
func (s *Service) UpdatePerformanceScoreSetting(ctx context.Context, courseID int, body interface{}) error {
	u := fmt.Sprintf("/api/course/%d/performance/score-setting", courseID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// ListStudentsPerformance returns course students-performance data using the frontend query shape.
func (s *Service) ListStudentsPerformance(ctx context.Context, courseID int, params *StudentsPerformanceParams) (*StudentsPerformanceResponse, error) {
	values := url.Values{}
	if params == nil || params.IsOriginalScore == nil {
		values.Set("isOriginalScore", "true")
	} else {
		values.Set("isOriginalScore", strconv.FormatBool(*params.IsOriginalScore))
	}
	if params != nil {
		if params.Page > 0 {
			values.Set("page", strconv.Itoa(params.Page))
		}
		if params.PageSize > 0 {
			values.Set("page_size", strconv.Itoa(params.PageSize))
		}
		if params.OnlyStudentsName {
			values.Set("onlyStudentsName", "true")
		}
		if encoded := encodeConditions(params.Conditions); encoded != "" {
			values.Set("conditions", encoded)
		}
	}
	u := fmt.Sprintf("/api/course/%d/students-performance", courseID)
	if encoded := values.Encode(); encoded != "" {
		u += "?" + encoded
	}
	result := new(StudentsPerformanceResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateStudentPerformanceScore updates one student's performance score.
func (s *Service) UpdateStudentPerformanceScore(ctx context.Context, courseID int, body *UpdateStudentPerformanceScoreRequest) error {
	u := fmt.Sprintf("/api/course/%d/performance/score", courseID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// GetScorePercentages returns the remaining/allocated score percentages for a course.
func (s *Service) GetScorePercentages(ctx context.Context, courseID int) (*ScorePercentagesResponse, error) {
	u := fmt.Sprintf("/api/course/%d/score-percentages", courseID)
	result := new(ScorePercentagesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetHomeworkSubmissionNumber returns the aggregated homework submission number payload for a course.
func (s *Service) GetHomeworkSubmissionNumber(ctx context.Context, courseID int) (*HomeworkSubmissionNumberResponse, error) {
	u := fmt.Sprintf("/api/course/%d/homework-submission-number", courseID)
	result := new(HomeworkSubmissionNumberResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetClassroomExamScores returns classroom score data for a course.
func (s *Service) GetClassroomExamScores(ctx context.Context, courseID int) (*ClassroomExamScoresResponse, error) {
	u := fmt.Sprintf("/api/course/%d/classroom-exam-scores", courseID)
	result := new(ClassroomExamScoresResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListExamSubmissions returns exam-submissions data for a course.
func (s *Service) ListExamSubmissions(ctx context.Context, courseID int) (*CourseExamSubmissionsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/exam-submissions", courseID)
	result := new(CourseExamSubmissionsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateSHTVUGroups creates SHTVU groups for a course.
func (s *Service) CreateSHTVUGroups(ctx context.Context, courseID int) (*SHTVUGroupsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/create-groups", courseID)
	result := new(SHTVUGroupsResponse)
	_, err := s.client.Post(ctx, u, nil, result)
	return result, err
}

// ListCatalogActivities returns catalog activities for a course.
func (s *Service) ListCatalogActivities(ctx context.Context, courseID int) (*CatalogActivitiesResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/catalog-activities", courseID)
	result := new(CatalogActivitiesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListCatalogActivitiesAnonymous returns visitor catalog activities for a course.
func (s *Service) ListCatalogActivitiesAnonymous(ctx context.Context, courseID int) (*CatalogActivitiesResponse, error) {
	u := fmt.Sprintf("/anonymous-api/courses/%d/catalog-activities", courseID)
	result := new(CatalogActivitiesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListQuestionnaires returns questionnaires for a course.
func (s *Service) ListQuestionnaires(ctx context.Context, courseID int) (*CourseQuestionnairesResponse, error) {
	u := fmt.Sprintf("/api/course/%d/questionnaires", courseID)
	result := new(CourseQuestionnairesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListQuestionnairesWithParams returns questionnaire-list data for a course.
func (s *Service) ListQuestionnairesWithParams(ctx context.Context, courseID int, params *ListQuestionnairesParams) (*QuestionnaireListResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/questionnaire-list", courseID)
	query := map[string]string{}
	if params != nil {
		if params.Page > 0 {
			query["page"] = strconv.Itoa(params.Page)
		}
		if params.PageSize > 0 {
			query["page_size"] = strconv.Itoa(params.PageSize)
		}
		if params.Conditions != nil {
			body, err := json.Marshal(params.Conditions)
			if err != nil {
				return nil, err
			}
			query["conditions"] = string(body)
		}
	}
	u = addQueryParams(u, query)
	result := new(QuestionnaireListResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetQuestionnaireScores returns questionnaire scores for a course.
func (s *Service) GetQuestionnaireScores(ctx context.Context, courseID int) (*QuestionnaireScoresResponse, error) {
	u := fmt.Sprintf("/api/course/%d/questionnaire-scores", courseID)
	result := new(QuestionnaireScoresResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetCourseWebLinkScores returns course-level web-link scores.
func (s *Service) GetCourseWebLinkScores(ctx context.Context, courseID int) (*CourseWebLinkScoresResponse, error) {
	u := fmt.Sprintf("/api/course/%d/web-link-scores", courseID)
	result := new(CourseWebLinkScoresResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetCourseVirtualExperimentScores returns course-level virtual-experiment scores.
func (s *Service) GetCourseVirtualExperimentScores(ctx context.Context, courseID int) (*CourseVirtualExperimentScoresResponse, error) {
	u := fmt.Sprintf("/api/course/%d/virtual-experiment-scores", courseID)
	result := new(CourseVirtualExperimentScoresResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListCourseVirtualExperiments returns virtual-experiment activities for a course.
func (s *Service) ListCourseVirtualExperiments(ctx context.Context, courseID int) (*CourseVirtualExperimentsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/virtual-experiments", courseID)
	result := new(CourseVirtualExperimentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListAllActivitiesByModuleIDs returns all activities for selected modules using the frontend query shape.
func (s *Service) ListAllActivitiesByModuleIDs(ctx context.Context, courseID int, params *ListAllActivitiesByModuleIDsParams) (*AllActivitiesResponse, error) {
	query := map[string]string{}
	if params != nil {
		if len(params.ModuleIDs) > 0 {
			ids := make([]string, 0, len(params.ModuleIDs))
			for _, id := range params.ModuleIDs {
				ids = append(ids, strconv.Itoa(id))
			}
			query["module_ids"] = "[" + strings.Join(ids, ",") + "]"
		}
		if params.ActivityTypes != "" {
			query["activity_types"] = params.ActivityTypes
		}
		if params.DisableLoadingAnimation {
			query["no-loading-animation"] = "true"
		}
	}
	u := addQueryParams(fmt.Sprintf("/api/courses/%d/all-activities", courseID), query)
	result := new(AllActivitiesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateCourseAdvanceSetting updates course advance settings used by the score-percentage UI.
func (s *Service) UpdateCourseAdvanceSetting(ctx context.Context, courseID int, body *CourseAdvanceSettingRequest) error {
	u := fmt.Sprintf("/api/course/%d/course-advance-setting", courseID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// UpdateScorePercentages updates activity/custom/rollcall/performance score percentages.
func (s *Service) UpdateScorePercentages(ctx context.Context, courseID int, body *UpdateScorePercentagesRequest) error {
	u := fmt.Sprintf("/api/courses/%d/score-percentages", courseID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// GetStudentSelfScore returns the current student's overall score breakdown.
func (s *Service) GetStudentSelfScore(ctx context.Context, courseID int) (*StudentSelfScoreResponse, error) {
	u := fmt.Sprintf("/api/course/%d/student-self-score", courseID)
	result := new(StudentSelfScoreResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetRollcallScore returns rollcall scoring data for a course.
func (s *Service) GetRollcallScore(ctx context.Context, courseID int) (*RollcallScoreResponse, error) {
	u := fmt.Sprintf("/api/course/%d/rollcall-score", courseID)
	result := new(RollcallScoreResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetPerformanceScore returns performance-score data for a course.
func (s *Service) GetPerformanceScore(ctx context.Context, courseID int, isOriginalScore bool) (*PerformanceScoreResponse, error) {
	u := fmt.Sprintf("/api/course/%d/performance-score", courseID)
	u = addQueryParams(u, map[string]string{"isOriginalScore": strconv.FormatBool(isOriginalScore)})
	result := new(PerformanceScoreResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetAnnounceScoreSettings returns the frontend announce-score-settings payload.
func (s *Service) GetAnnounceScoreSettings(ctx context.Context, courseID int) (*AnnounceScoreSettingsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/announce-score-settings", courseID)
	result := new(AnnounceScoreSettingsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateAnnounceScoreSettings updates the frontend announce-score-settings payload.
func (s *Service) UpdateAnnounceScoreSettings(ctx context.Context, courseID int, body interface{}) error {
	u := fmt.Sprintf("/api/courses/%d/announce-score-settings", courseID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// GetPerformanceScorePercentage returns performance score percentage settings for a course.
func (s *Service) GetPerformanceScorePercentage(ctx context.Context, courseID int) (*PerformanceScorePercentageResponse, error) {
	u := fmt.Sprintf("/api/course/%d/performance/score-percentage", courseID)
	result := new(PerformanceScorePercentageResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListWarnings returns the warning list payload for a course.
func (s *Service) ListWarnings(ctx context.Context, courseID int, fields string) (*WarningListResponse, error) {
	u := fmt.Sprintf("/api/course/%d/warnings", courseID)
	if fields != "" {
		u = addQueryParams(u, map[string]string{"fields": fields})
	}
	result := new(WarningListResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetWarning returns a specific warning payload.
func (s *Service) GetWarning(ctx context.Context, courseID, warningID int) (*WarningResponse, error) {
	u := fmt.Sprintf("/api/course/%d/warnings/%d", courseID, warningID)
	result := new(WarningResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetWarningThreshold returns the warning threshold payload.
func (s *Service) GetWarningThreshold(ctx context.Context, courseID, warningID int) (*WarningThresholdResponse, error) {
	u := fmt.Sprintf("/api/course/%d/warnings/%d/threshold", courseID, warningID)
	result := new(WarningThresholdResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateWarningThreshold creates a warning threshold configuration.
func (s *Service) CreateWarningThreshold(ctx context.Context, courseID, warningID int, body interface{}) error {
	u := fmt.Sprintf("/api/course/%d/warnings/%d/threshold", courseID, warningID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// UpdateWarningThreshold updates a warning threshold configuration.
func (s *Service) UpdateWarningThreshold(ctx context.Context, courseID, warningID int, body interface{}) error {
	u := fmt.Sprintf("/api/course/%d/warnings/%d/threshold", courseID, warningID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// ListWarningStudents returns the warning student list payload.
func (s *Service) ListWarningStudents(ctx context.Context, courseID, warningID int, params *WarningStudentsParams) (*WarningStudentsResponse, error) {
	query := map[string]string{}
	if params != nil {
		if params.Fields != "" {
			query["fields"] = params.Fields
		}
		if params.Conditions != nil {
			body, err := json.Marshal(params.Conditions)
			if err != nil {
				return nil, err
			}
			query["conditions"] = string(body)
		}
	}
	u := addQueryParams(fmt.Sprintf("/api/course/%d/warnings/%d/students", courseID, warningID), query)
	result := new(WarningStudentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateWarningStudentComment updates a warning-student comment.
func (s *Service) UpdateWarningStudentComment(ctx context.Context, warningStudentID int, body interface{}) error {
	u := fmt.Sprintf("/api/warning/student/%d/comment", warningStudentID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// DeleteWarningStudent removes a student from a warning scope.
func (s *Service) DeleteWarningStudent(ctx context.Context, courseID, warningStudentID int) error {
	u := fmt.Sprintf("/api/course/%d/warning-students/%d", courseID, warningStudentID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// GetScoreRanks returns score-ranks data for a course.
func (s *Service) GetScoreRanks(ctx context.Context, courseID int) (*ScoreRanksResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/score-ranks", courseID)
	result := new(ScoreRanksResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetScoreItemSettings returns score-item-settings data for a course.
func (s *Service) GetScoreItemSettings(ctx context.Context, courseID int) (*ScoreItemSettingsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/score-item-settings", courseID)
	result := new(ScoreItemSettingsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetEnrollmentRawScore returns raw enrollment score data for a course.
func (s *Service) GetEnrollmentRawScore(ctx context.Context, courseID int) (*EnrollmentRawScoreResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/enrollment-raw-score", courseID)
	result := new(EnrollmentRawScoreResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetScoreDistribution returns the frontend score-distribution payload.
func (s *Service) GetScoreDistribution(ctx context.Context, courseID int) (*ScoreDistributionResponse, error) {
	u := fmt.Sprintf("/api/course/%d/score-distribution?no-intercept=true", courseID)
	result := new(ScoreDistributionResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetScoreTypeSettings returns score-type settings for a course.
func (s *Service) GetScoreTypeSettings(ctx context.Context, courseID int) (*ScoreTypeSettingsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/score-type-settings", courseID)
	result := new(ScoreTypeSettingsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetCourseScoreType returns the score type for a specific course activity category.
func (s *Service) GetCourseScoreType(ctx context.Context, courseID int, activityType string) (*CourseScoreTypeResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/%s/score-type", courseID, url.PathEscape(activityType))
	result := new(CourseScoreTypeResponse)
	_, err := s.client.Get(ctx, u, result)
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

// AssignActivityToModule assigns an activity to a module within a course.
func (s *Service) AssignActivityToModule(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/assign-activity-to-module", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// AddEnrollments adds course enrollments using the teaching-team endpoint.
func (s *Service) AddEnrollments(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/add-enrollments", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// GetCombineCourseActivityReadSubCourseIDs returns the sub-course IDs that already contain read/submission data for the activity.
func (s *Service) GetCombineCourseActivityReadSubCourseIDs(ctx context.Context, courseID, activityID int) ([]int, error) {
	u := addQueryParams(fmt.Sprintf("/api/combine-courses/%d/activity-read-sub-course-ids", courseID), map[string]string{
		"activity_id": strconv.Itoa(activityID),
	})
	var result []int
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// UpdateEnrollment updates a single course enrollment.
func (s *Service) UpdateEnrollment(ctx context.Context, enrollmentID int, body interface{}) error {
	u := fmt.Sprintf("/api/course/enrollments/%d", enrollmentID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// UpdateEnrollmentSeatNumber updates a course enrollment seat number.
func (s *Service) UpdateEnrollmentSeatNumber(ctx context.Context, enrollmentID int, body *UpdateEnrollmentSeatNumberRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/enrollments/%d", enrollmentID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// DeleteStudentEnrollment removes a user enrollment from a course.
func (s *Service) DeleteStudentEnrollment(ctx context.Context, courseID int, userID int) error {
	u := fmt.Sprintf("/api/course/%d/students/%d", courseID, userID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// MailToEnrollments sends an email to selected enrollments.
func (s *Service) MailToEnrollments(ctx context.Context, courseID int, body *SendMailToEnrollmentsRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/mail-to-enrollments", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// GetAssistantPermissions returns course assistant permissions.
func (s *Service) GetAssistantPermissions(ctx context.Context, courseID int) (*AssistantPermissions, error) {
	u := fmt.Sprintf("/api/course/%d/assistant-permissions", courseID)
	result := new(AssistantPermissions)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateAssistantPermissions updates course assistant permissions.
func (s *Service) UpdateAssistantPermissions(ctx context.Context, courseID int, body *UpdateAssistantPermissionsRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/assistant-permissions", courseID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// UpdateMainSetting updates main teaching-team settings such as course leader.
func (s *Service) UpdateMainSetting(ctx context.Context, courseID int, body interface{}) error {
	u := fmt.Sprintf("/api/courses/%d/main-setting", courseID)
	_, err := s.client.Put(ctx, u, body, nil)
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
func (s *Service) GetBlueprintSubItemsCount(ctx context.Context, courseID int, activities []BlueprintActivityRef) (*BlueprintSubItemsResponse, error) {
	u := fmt.Sprintf("/api/blueprint/%d/sub-items-count", courseID)
	if len(activities) > 0 {
		encoded, err := json.Marshal(activities)
		if err != nil {
			return nil, err
		}
		u = addQueryParams(u, map[string]string{"activities": string(encoded)})
	}
	result := new(BlueprintSubItemsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetBlueprintAllSubActivitiesCount returns the aggregated sub-activity count for a blueprint course.
func (s *Service) GetBlueprintAllSubActivitiesCount(ctx context.Context, courseID int) (*BlueprintAllSubActivitiesCountResponse, error) {
	u := fmt.Sprintf("/api/blueprint/%d/all-sub-activities-count", courseID)
	result := new(BlueprintAllSubActivitiesCountResponse)
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

// PublishActivities publishes activities to target courses.
func (s *Service) PublishActivities(ctx context.Context, courseID int, body PublishActivitiesRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/publish-activities", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ListBlueprintSubCourses returns blueprint sub courses.
func (s *Service) ListBlueprintSubCourses(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blueprint/%d/sub-courses", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListBlueprintSubCoursesWithParams returns blueprint sub courses using the frontend query shape.
func (s *Service) ListBlueprintSubCoursesWithParams(ctx context.Context, courseID int, params *ListBlueprintSubCoursesParams) (*BlueprintSubCoursesResponse, error) {
	query := map[string]string{}
	if params != nil {
		if params.Keyword != "" {
			query["keyword"] = params.Keyword
		}
		if params.SourceID > 0 {
			query["source_id"] = strconv.Itoa(params.SourceID)
		}
		if params.SourceType != "" {
			query["source_type"] = params.SourceType
		}
	}
	u := addQueryParams(fmt.Sprintf("/api/blueprint/%d/sub-courses", courseID), query)
	result := new(BlueprintSubCoursesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// BindBlueprintSubCourses binds sub-courses to a blueprint course.
func (s *Service) BindBlueprintSubCourses(ctx context.Context, courseID int, body *BindBlueprintSubCoursesRequest) error {
	u := fmt.Sprintf("/api/blueprint/%d/sub-courses", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// CheckBlueprintPrerequisites checks whether sources have prerequisites before blueprint sync.
func (s *Service) CheckBlueprintPrerequisites(ctx context.Context, courseID int, body *CheckBlueprintPrerequisitesRequest) (*BlueprintCheckPrerequisitesResponse, error) {
	u := fmt.Sprintf("/api/blueprint/%d/check-prerequisites", courseID)
	result := new(BlueprintCheckPrerequisitesResponse)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// DeleteBlueprintSubCourse removes a bound sub-course from a blueprint course.
func (s *Service) DeleteBlueprintSubCourse(ctx context.Context, courseID, subCourseID int) error {
	u := fmt.Sprintf("/api/blueprint/%d/sub-courses/%d", courseID, subCourseID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// RenameBlueprintSubCourse updates the frontend-visible name for a blueprint sub-course.
func (s *Service) RenameBlueprintSubCourse(ctx context.Context, subCourseID int, body *RenameBlueprintSubCourseRequest) error {
	u := fmt.Sprintf("/api/blueprint/sub-courses/%d/name", subCourseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// DeleteBlueprint deletes a blueprint course mapping/configuration.
func (s *Service) DeleteBlueprint(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/blueprint/%d", courseID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// CancelBlueprintActivitySync cancels blueprint sync for an activity.
func (s *Service) CancelBlueprintActivitySync(ctx context.Context, courseID, activityID int, body *CancelBlueprintActivitySyncRequest) error {
	u := fmt.Sprintf("/api/blueprint/%d/activities/%d/cancel-sync", courseID, activityID)
	_, err := s.client.DeleteWithBody(ctx, u, body, nil)
	return err
}

// GetBlueprintSubmittedInfo returns blueprint submitted sync info for a target object.
func (s *Service) GetBlueprintSubmittedInfo(ctx context.Context, courseID int, resourceType string, resourceID int) (*BlueprintSubmittedInfoResponse, error) {
	u := fmt.Sprintf("/api/blueprint/%d/%s/%d/submitted-info", courseID, resourceType, resourceID)
	result := new(BlueprintSubmittedInfoResponse)
	_, err := s.client.Get(ctx, u, result)
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

// GetTextAnalyticsConfig returns course forum text analytics configuration.
func (s *Service) GetTextAnalyticsConfig(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/text-analytics-config", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListCourseAuditReferences returns paged resource audit references for a specific course.
func (s *Service) ListCourseAuditReferences(ctx context.Context, courseID int, params ListCourseResourceAuditParams) (*CourseAuditReferencesResponse, error) {
	u := addListOptions(
		fmt.Sprintf("/api/courses/%d/resource-audit", courseID),
		&model.ListOptions{Page: params.Page, PageSize: params.PageSize},
	)
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	result := new(CourseAuditReferencesResponse)
	_, err := s.client.Get(ctx, u, result)
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

// GetTPDOEStatStudentsWithParams returns TPDOE student statistics using the newer frontend query shape.
func (s *Service) GetTPDOEStatStudentsWithParams(ctx context.Context, params TPDOEStatStudentsParams) (json.RawMessage, error) {
	values := url.Values{}
	for _, courseID := range params.CourseIDs {
		values.Add("courseIds", fmt.Sprintf("%d", courseID))
	}
	if params.StartDate != "" {
		values.Set("startDate", params.StartDate)
	}
	if params.EndDate != "" {
		values.Set("endDate", params.EndDate)
	}
	if params.StatType != "" {
		values.Set("statType", params.StatType)
	}
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		values.Set("conditions", encoded)
	}
	u := "/api/courses/tpdoe/stat-students"
	if encoded := values.Encode(); encoded != "" {
		u += "?" + encoded
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// InspectCourse triggers the backend inspection workflow for a course.
func (s *Service) InspectCourse(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/courses/%d/inspect", courseID)
	_, err := s.client.Put(ctx, u, nil, nil)
	return err
}

// UpdateAudit updates a course audit record.
func (s *Service) UpdateAudit(ctx context.Context, courseID int, auditID int, body interface{}) error {
	u := fmt.Sprintf("/api/courses/%d/audit/%d", courseID, auditID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
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

// ListCourseClassificationsWithPrefix returns course classifications from either /api or /anonymous-api.
func (s *Service) ListCourseClassificationsWithPrefix(ctx context.Context, anonymous bool) (*CourseClassificationsResponse, error) {
	prefix := "/api"
	if anonymous {
		prefix = "/anonymous-api"
	}
	result := new(CourseClassificationsResponse)
	_, err := s.client.Get(ctx, prefix+"/course-classifications", result)
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

func encodeConditions(conditions any) string {
	switch value := conditions.(type) {
	case nil:
		return ""
	case string:
		return value
	default:
		encoded, err := json.Marshal(value)
		if err != nil {
			return ""
		}
		return string(encoded)
	}
}
