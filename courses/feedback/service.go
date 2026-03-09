package feedback

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles feedback-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// GetFeedbackActivity returns a feedback activity.
func (s *Service) GetFeedbackActivity(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d", activityID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetFeedbackActivityTyped returns a feedback activity with the frontend activity model.
func (s *Service) GetFeedbackActivityTyped(ctx context.Context, activityID int) (*FeedbackActivity, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d", activityID)
	result := new(FeedbackActivity)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateFeedbackActivity creates a feedback activity.
func (s *Service) CreateFeedbackActivity(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/feedback-activities", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// CreateFeedbackActivityTyped creates a feedback activity and decodes the created activity.
func (s *Service) CreateFeedbackActivityTyped(ctx context.Context, courseID int, body interface{}) (*FeedbackActivity, error) {
	u := fmt.Sprintf("/api/courses/%d/feedback-activities", courseID)
	result := new(FeedbackActivity)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateFeedbackActivity updates a feedback activity.
func (s *Service) UpdateFeedbackActivity(ctx context.Context, activityID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d", activityID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// UpdateFeedbackActivityTyped updates a feedback activity and decodes the activity payload.
func (s *Service) UpdateFeedbackActivityTyped(ctx context.Context, activityID int, body interface{}) (*FeedbackActivity, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d", activityID)
	result := new(FeedbackActivity)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// GetFeedback returns a feedback.
func (s *Service) GetFeedback(ctx context.Context, feedbackID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedbacks/%d", feedbackID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetFeedbackTyped returns a feedback record.
func (s *Service) GetFeedbackTyped(ctx context.Context, feedbackID int) (*Feedback, error) {
	u := fmt.Sprintf("/api/feedbacks/%d", feedbackID)
	result := new(Feedback)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateFeedback creates feedback under a feedback activity.
func (s *Service) CreateFeedback(ctx context.Context, activityID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d/feedbacks", activityID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// CreateFeedbackTyped creates feedback under a feedback activity.
func (s *Service) CreateFeedbackTyped(ctx context.Context, activityID int, body interface{}) (*Feedback, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d/feedbacks", activityID)
	result := new(Feedback)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateFeedback updates feedback under a feedback activity.
func (s *Service) UpdateFeedback(ctx context.Context, activityID, feedbackID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d/feedbacks/%d", activityID, feedbackID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// UpdateFeedbackTyped updates feedback under a feedback activity.
func (s *Service) UpdateFeedbackTyped(ctx context.Context, activityID, feedbackID int, body interface{}) (*Feedback, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d/feedbacks/%d", activityID, feedbackID)
	result := new(Feedback)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteFeedback deletes a feedback record.
func (s *Service) DeleteFeedback(ctx context.Context, feedbackID int) error {
	u := fmt.Sprintf("/api/feedbacks/%d", feedbackID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
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
