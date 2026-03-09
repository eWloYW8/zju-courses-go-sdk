package courses

import (
	"context"
	"encoding/json"
	"fmt"
)

// --- Course Custom Score ---

// GetCourseCustomScoreItems returns custom score items for a course.
func (s *Service) GetCourseCustomScoreItems(ctx context.Context, courseID int) (*CourseCustomScoreItemsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/custom-score-items", courseID)
	result := new(CourseCustomScoreItemsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateCourseCustomScoreItems updates custom score items for a course.
func (s *Service) UpdateCourseCustomScoreItems(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/custom-score-items/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// CreateCourseCustomScoreItem creates a custom score item for a course.
func (s *Service) CreateCourseCustomScoreItem(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/custom-score-item", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// --- Access Code ---

// GetCourseAccessCode returns the access code for a course.
func (s *Service) GetCourseAccessCode(ctx context.Context, courseID int) (*CourseAccessCodeResponse, error) {
	u := fmt.Sprintf("/api/course/%d/access_code", courseID)
	result := new(CourseAccessCodeResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ResetCourseAccessCode regenerates the access code for a course.
func (s *Service) ResetCourseAccessCode(ctx context.Context, courseID int) (*CourseAccessCodeResponse, error) {
	u := fmt.Sprintf("/api/course/%d/access_code", courseID)
	result := new(CourseAccessCodeResponse)
	_, err := s.client.Put(ctx, u, nil, result)
	return result, err
}

// ValidateCourseAccessCode validates a join-course access code through the frontend popup endpoint.
func (s *Service) ValidateCourseAccessCode(ctx context.Context, accessCode string) (*CourseAccessCodeValidationResponse, error) {
	u := fmt.Sprintf("/api/course/access-code/%s/validate", accessCode)
	result := new(CourseAccessCodeValidationResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListTeamTeachingOpenedOrgs returns organizations available in the team-teaching invite popup.
func (s *Service) ListTeamTeachingOpenedOrgs(ctx context.Context) (*OpenedOrgsResponse, error) {
	result := new(OpenedOrgsResponse)
	_, err := s.client.Get(ctx, "/api/toggle-opened-orgs?toggle=org_team_teaching", result)
	return result, err
}

// --- Completion Criteria ---

// ListCompletionCriteria returns completion criteria.
func (s *Service) ListCompletionCriteria(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/completion-criteria", &result)
	return result, err
}
