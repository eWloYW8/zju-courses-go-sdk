package interactions

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
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

// DeleteInteraction deletes a course interaction.
func (s *Service) DeleteInteraction(ctx context.Context, interactionID int) error {
	u := fmt.Sprintf("/api/interactions/%d", interactionID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
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

// ListInteractionSubjects returns the subjects configured for an interaction activity.
func (s *Service) ListInteractionSubjects(ctx context.Context, activityID int) (*InteractionSubjectsResponse, error) {
	u := fmt.Sprintf("/api/interaction-activities/%d/subjects", activityID)
	result := new(InteractionSubjectsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetCurrentUserInteractionSubmission returns the current user's submission for an interaction activity.
func (s *Service) GetCurrentUserInteractionSubmission(ctx context.Context, activityID int) (*InteractionSubmission, error) {
	u := fmt.Sprintf("/api/interaction-activities/%d/submission", activityID)
	result := new(InteractionSubmission)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListInteractionSubmissions returns submissions for an interaction activity.
func (s *Service) ListInteractionSubmissions(ctx context.Context, activityID int) (*InteractionSubmissionsResponse, error) {
	u := fmt.Sprintf("/api/interaction-activities/%d/submissions", activityID)
	result := new(InteractionSubmissionsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateInteractionSubmission creates a submission for an interaction activity.
func (s *Service) CreateInteractionSubmission(ctx context.Context, activityID int) (*InteractionSubmission, error) {
	u := fmt.Sprintf("/api/interaction-activities/%d/submissions", activityID)
	result := new(InteractionSubmission)
	_, err := s.client.Post(ctx, u, nil, result)
	return result, err
}

// GetInteractionSubmission returns an interaction submission.
func (s *Service) GetInteractionSubmission(ctx context.Context, submissionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/interaction-submissions/%d", submissionID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// UpdateInteractionSubmission updates answers in an interaction submission.
func (s *Service) UpdateInteractionSubmission(ctx context.Context, submissionID int, body *UpdateInteractionSubmissionRequest) (*InteractionSubmission, error) {
	u := fmt.Sprintf("/api/interaction-submissions/%d", submissionID)
	result := new(InteractionSubmission)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// SaveInteractionActivityAsResource saves an interaction activity as a reusable resource.
func (s *Service) SaveInteractionActivityAsResource(ctx context.Context, activityID int) error {
	u := fmt.Sprintf("/api/interaction-activities/%d/save-as-resource", activityID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
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
