package groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// Service handles group-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// GetGroupSet returns a group set.
func (s *Service) GetGroupSet(ctx context.Context, groupSetID int) (*model.GroupSet, error) {
	u := fmt.Sprintf("/api/group-sets/%d", groupSetID)
	result := new(model.GroupSet)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetGroup returns a group.
func (s *Service) GetGroup(ctx context.Context, groupID int) (*model.Group, error) {
	u := fmt.Sprintf("/api/groups/%d", groupID)
	result := new(model.Group)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetMyGroup returns the current user's group within a group set.
func (s *Service) GetMyGroup(ctx context.Context, groupSetID int) (*model.Group, error) {
	u := fmt.Sprintf("/api/group-sets/%d/group", groupSetID)
	result := new(model.Group)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListGroupSets returns group sets for a course.
func (s *Service) ListGroupSets(ctx context.Context, courseID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/courses/%d/group-sets", courseID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListGroups returns groups in a group set.
func (s *Service) ListGroups(ctx context.Context, groupSetID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/group-sets/%d/groups", groupSetID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
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

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}
