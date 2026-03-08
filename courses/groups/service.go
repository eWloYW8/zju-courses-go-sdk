package groups

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles group-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// GetGroupSet returns a group set.
func (s *Service) GetGroupSet(ctx context.Context, groupSetID int) (*GroupSet, error) {
	u := fmt.Sprintf("/api/group-sets/%d", groupSetID)
	result := new(GroupSet)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetGroup returns a group.
func (s *Service) GetGroup(ctx context.Context, groupID int) (*Group, error) {
	u := fmt.Sprintf("/api/groups/%d", groupID)
	result := new(Group)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetMyGroup returns the current user's group within a group set.
func (s *Service) GetMyGroup(ctx context.Context, groupSetID int) (*Group, error) {
	u := fmt.Sprintf("/api/group-sets/%d/group", groupSetID)
	result := new(Group)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateGroupSet creates a group set for a course.
func (s *Service) CreateGroupSet(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/group-sets", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// UpdateGroupSet updates a group set.
func (s *Service) UpdateGroupSet(ctx context.Context, groupSetID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d", groupSetID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// DeleteGroupSet deletes a group set.
func (s *Service) DeleteGroupSet(ctx context.Context, groupSetID int) error {
	u := fmt.Sprintf("/api/group-sets/%d", groupSetID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// CopyGroupSet copies a group set to the given course.
func (s *Service) CopyGroupSet(ctx context.Context, courseID, groupSetID int, body *CopyGroupSetRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/group-sets/%d/copy", courseID, groupSetID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ListGroupSets returns group sets for a course.
func (s *Service) ListGroupSets(ctx context.Context, courseID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/courses/%d/group-sets", courseID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListGroupSetsWithParams returns group sets for a course with frontend query params.
func (s *Service) ListGroupSetsWithParams(ctx context.Context, courseID int, params *ListGroupSetsParams) (*GroupSetsResponse, error) {
	query := map[string]string{}
	if params != nil && params.PreloadID != nil {
		query["preload_id"] = strconv.Itoa(*params.PreloadID)
	}
	u := addQueryParams(fmt.Sprintf("/api/courses/%d/group-sets", courseID), query)
	result := new(GroupSetsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetStudents returns students for a course using the same query params as the frontend.
func (s *Service) GetStudents(ctx context.Context, courseID int, params *GetGroupStudentsParams) (*CourseStudentsResponse, error) {
	query := map[string]string{}
	if params != nil {
		if params.IgnoreAvatar {
			query["ignore_avatar"] = strconv.FormatBool(true)
		}
		if params.Fields != "" {
			query["fields"] = params.Fields
		}
	}
	u := addQueryParams(fmt.Sprintf("/api/course/%d/students", courseID), query)
	result := new(CourseStudentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetEnrollments returns course enrollments for group management pages.
func (s *Service) GetEnrollments(ctx context.Context, courseID int) (*CourseEnrollmentsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/enrollments", courseID)
	result := new(CourseEnrollmentsResponse)
	_, err := s.client.Post(ctx, u, nil, result)
	return result, err
}

// ListGroups returns groups in a group set.
func (s *Service) ListGroups(ctx context.Context, groupSetID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/group-sets/%d/groups", groupSetID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListGroupsWithParams returns groups in a group set with frontend query params.
func (s *Service) ListGroupsWithParams(ctx context.Context, groupSetID int, params *ListGroupsParams) (*GroupsResponse, error) {
	query := map[string]string{}
	if params != nil && params.Fields != "" {
		query["fields"] = params.Fields
	}
	u := addQueryParams(fmt.Sprintf("/api/group-sets/%d/groups", groupSetID), query)
	result := new(GroupsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetGroupOtherUsers returns users in the group set but outside the current group.
func (s *Service) GetGroupOtherUsers(ctx context.Context, groupSetID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d/other-users", groupSetID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetGroupHomeworkActivities returns grouped homework activities for a group set.
func (s *Service) GetGroupHomeworkActivities(ctx context.Context, groupSetID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d/group-homework-activities", groupSetID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetGroupExams returns grouped exams for a group set.
func (s *Service) GetGroupExams(ctx context.Context, groupSetID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d/group-exams", groupSetID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// UpdateGroup updates a group.
func (s *Service) UpdateGroup(ctx context.Context, groupID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/groups/%d", groupID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// AddGroupMembers adds members to a group.
func (s *Service) AddGroupMembers(ctx context.Context, groupID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/groups/%d/members", groupID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// DeleteGroupMember removes a member from a group.
func (s *Service) DeleteGroupMember(ctx context.Context, groupID, memberID int) error {
	u := fmt.Sprintf("/api/groups/%d/members/%d", groupID, memberID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// UpdateGroupMember updates a group member.
func (s *Service) UpdateGroupMember(ctx context.Context, groupID, memberID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/groups/%d/members/%d", groupID, memberID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// UpdateGroupInfo updates the information section of a group.
func (s *Service) UpdateGroupInfo(ctx context.Context, groupID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/groups/%d/info", groupID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// CreateGroup creates a group inside a group set.
func (s *Service) CreateGroup(ctx context.Context, groupSetID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d/groups", groupSetID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// DeleteGroup deletes a group.
func (s *Service) DeleteGroup(ctx context.Context, groupID int) error {
	u := fmt.Sprintf("/api/groups/%d", groupID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// SortGroups updates the order of groups inside a group set.
func (s *Service) SortGroups(ctx context.Context, groupSetID int, body *SortGroupsRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d/sort", groupSetID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// UploadGroupFiles binds uploads to a group.
func (s *Service) UploadGroupFiles(ctx context.Context, groupID int, body *UploadGroupFilesRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/groups/%d/upload", groupID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// DeleteGroupFile removes a file from a group.
func (s *Service) DeleteGroupFile(ctx context.Context, groupID, referenceID int) error {
	u := fmt.Sprintf("/api/groups/%d/upload/%d", groupID, referenceID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// BatchDownloadUploads creates a batch blob task for the given upload selection.
func (s *Service) BatchDownloadUploads(ctx context.Context, body *BatchDownloadUploadsRequest) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/uploads/batch/blob", body, &result)
	return result, err
}

// RandomGrouping groups students automatically within a group set.
func (s *Service) RandomGrouping(ctx context.Context, groupSetID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d/random-grouping", groupSetID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ExportCourseGroupSetInfoToExcel exports all group-set info for a course as an Excel file.
func (s *Service) ExportCourseGroupSetInfoToExcel(ctx context.Context, courseID int) ([]byte, error) {
	u := fmt.Sprintf("/api/courses/%d/group-sets/export/excel", courseID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	_, body, err := s.client.DoBytes(req)
	return body, err
}

// ExportGroupSetInfoToExcel exports a group set as an Excel file.
func (s *Service) ExportGroupSetInfoToExcel(ctx context.Context, groupSetID int) ([]byte, error) {
	u := fmt.Sprintf("/api/group-sets/%d/export/excel", groupSetID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	_, body, err := s.client.DoBytes(req)
	return body, err
}

// GetGroupsSubmissionStatusOfUser returns group submission status for a course.
func (s *Service) GetGroupsSubmissionStatusOfUser(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/groups/submission-status", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetActivitiesByGroupSet returns activities related to a group set.
func (s *Service) GetActivitiesByGroupSet(ctx context.Context, groupSetID int) (*GroupActivitiesResponse, error) {
	u := fmt.Sprintf("/api/group-sets/%d/activities", groupSetID)
	result := new(GroupActivitiesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetTeachingTeamGroups returns the teaching-team groups for a course.
func (s *Service) GetTeachingTeamGroups(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/teaching-team/groups", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}
