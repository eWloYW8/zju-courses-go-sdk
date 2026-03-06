package zjucourses

import (
	"context"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// ForumService handles forum-related API operations.
type ForumService struct {
	client *Client
}

// --- Response Types ---

type TopicsResponse struct {
	Topics []*model.Topic `json:"topics"`
	model.Pagination
}

// --- Forum Categories ---

// GetCategory returns a forum category with its topics.
func (s *ForumService) GetCategory(ctx context.Context, categoryID int, opts *model.ListOptions) (*model.ForumCategoryResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/forum/categories/%d", categoryID), opts)
	result := new(model.ForumCategoryResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Topics ---

// GetTopic returns a specific topic.
func (s *ForumService) GetTopic(ctx context.Context, topicID int) (*model.Topic, error) {
	u := fmt.Sprintf("/api/forum/topics/%d", topicID)
	result := new(model.Topic)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// CreateTopic creates a new topic in a category.
func (s *ForumService) CreateTopic(ctx context.Context, body interface{}) (*model.Topic, error) {
	result := new(model.Topic)
	_, err := s.client.post(ctx, "/api/topics", body, result)
	return result, err
}

// UpdateTopic updates an existing topic.
func (s *ForumService) UpdateTopic(ctx context.Context, topicID int, body interface{}) (*model.Topic, error) {
	u := fmt.Sprintf("/api/topics/%d", topicID)
	result := new(model.Topic)
	_, err := s.client.put(ctx, u, body, result)
	return result, err
}

// DeleteTopic deletes a topic.
func (s *ForumService) DeleteTopic(ctx context.Context, topicID int) error {
	u := fmt.Sprintf("/api/topics/%d", topicID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// ListLatestTopics returns the latest topics.
func (s *ForumService) ListLatestTopics(ctx context.Context, count int) ([]*model.Topic, error) {
	u := fmt.Sprintf("/api/topics/latest?no-intercept=true&count=%d", count)
	var wrapper struct {
		Topics []*model.Topic `json:"topics"`
	}
	_, err := s.client.get(ctx, u, &wrapper)
	return wrapper.Topics, err
}

// --- Forum Score ---

// ForumScoreResponse wraps the forum score API response.
type ForumScoreResponse struct {
	ForumScore *model.ForumScore `json:"forum_score"`
}

// GetForumScore returns the forum score for a student.
func (s *ForumService) GetForumScore(ctx context.Context, activityID, studentID int) (*ForumScoreResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/students/%d/forum-score", activityID, studentID)
	result := new(ForumScoreResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Replies ---

// CreateReply creates a reply to a topic.
func (s *ForumService) CreateReply(ctx context.Context, topicID int, body interface{}) (*model.Reply, error) {
	u := fmt.Sprintf("/api/replies/%d", topicID)
	result := new(model.Reply)
	_, err := s.client.post(ctx, u, body, result)
	return result, err
}

// UpdateReply updates a reply.
func (s *ForumService) UpdateReply(ctx context.Context, replyID int, body interface{}) (*model.Reply, error) {
	u := fmt.Sprintf("/api/replies/%d", replyID)
	result := new(model.Reply)
	_, err := s.client.put(ctx, u, body, result)
	return result, err
}

// DeleteReply deletes a reply.
func (s *ForumService) DeleteReply(ctx context.Context, replyID int) error {
	u := fmt.Sprintf("/api/replies/%d", replyID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}
