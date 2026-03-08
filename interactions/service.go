package interactions

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// Service handles interaction-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// GetInteraction returns a specific interaction.
func (s *Service) GetInteraction(ctx context.Context, interactionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/interactions/%d", interactionID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// VoteInteraction votes on an interaction.
func (s *Service) VoteInteraction(ctx context.Context, interactionID int, body interface{}) error {
	u := fmt.Sprintf("/api/courses/interactions/vote/%d", interactionID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// GetInteractionActivity returns an interaction activity.
func (s *Service) GetInteractionActivity(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/interaction-activities/%d", activityID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetInteractionSubmission returns an interaction submission.
func (s *Service) GetInteractionSubmission(ctx context.Context, submissionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/interaction-submissions/%d", submissionID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListCourseInteractions returns interactions for a course.
func (s *Service) ListCourseInteractions(ctx context.Context, courseID int) (*CourseInteractionsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/interactions", courseID)
	result := new(CourseInteractionsResponse)
	_, err := s.client.Get(ctx, u, result)
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
