package forum

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// Service handles forum-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Response Types ---

type TopicsResponse struct {
	Topics []*model.Topic `json:"topics"`
	model.Pagination
}

// --- Forum Categories ---

// GetCategory returns a forum category with its topics.
func (s *Service) GetCategory(ctx context.Context, categoryID int, opts *model.ListOptions) (*model.ForumCategoryResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/forum/categories/%d", categoryID), opts)
	result := new(model.ForumCategoryResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Topics ---

// GetTopic returns a specific topic.
func (s *Service) GetTopic(ctx context.Context, topicID int) (*model.Topic, error) {
	u := fmt.Sprintf("/api/forum/topics/%d", topicID)
	result := new(model.Topic)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateTopic creates a new topic in a category.
func (s *Service) CreateTopic(ctx context.Context, body interface{}) (*model.Topic, error) {
	result := new(model.Topic)
	_, err := s.client.Post(ctx, "/api/topics", body, result)
	return result, err
}

// UpdateTopic updates an existing topic.
func (s *Service) UpdateTopic(ctx context.Context, topicID int, body interface{}) (*model.Topic, error) {
	u := fmt.Sprintf("/api/topics/%d", topicID)
	result := new(model.Topic)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteTopic deletes a topic.
func (s *Service) DeleteTopic(ctx context.Context, topicID int) error {
	u := fmt.Sprintf("/api/topics/%d", topicID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// LikeTopic likes a topic.
func (s *Service) LikeTopic(ctx context.Context, topicID int) error {
	u := fmt.Sprintf("/api/topics/%d/likes", topicID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// UnlikeTopic removes a like from a topic.
func (s *Service) UnlikeTopic(ctx context.Context, topicID int) error {
	u := fmt.Sprintf("/api/topics/%d/likes", topicID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// TopTopic marks a topic as topped.
func (s *Service) TopTopic(ctx context.Context, topicID int, body interface{}) error {
	u := fmt.Sprintf("/api/topics/%d/topped", topicID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// CancelTopTopic removes topped status from a topic.
func (s *Service) CancelTopTopic(ctx context.Context, topicID int, params map[string]string) error {
	u := addQueryParams(fmt.Sprintf("/api/topics/%d/topped", topicID), params)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ListLatestTopics returns the latest topics.
func (s *Service) ListLatestTopics(ctx context.Context, count int) ([]*model.Topic, error) {
	u := fmt.Sprintf("/api/topics/latest?no-intercept=true&count=%d", count)
	var wrapper struct {
		Topics []*model.Topic `json:"topics"`
	}
	_, err := s.client.Get(ctx, u, &wrapper)
	return wrapper.Topics, err
}

// --- Forum Score ---

// ForumScoreResponse wraps the forum score API response.
type ForumScoreResponse struct {
	ForumScore *model.ForumScore `json:"forum_score"`
}

// GetForumScore returns the forum score for a student.
func (s *Service) GetForumScore(ctx context.Context, activityID, studentID int) (*ForumScoreResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/students/%d/forum-score", activityID, studentID)
	result := new(ForumScoreResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetForumScores returns all forum scores for an activity.
func (s *Service) GetForumScores(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/activities/%d/forum-scores", activityID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCourseForumScores returns forum scores for a course.
func (s *Service) GetCourseForumScores(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/forum-scores", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// SaveForumScore updates a student's forum score.
func (s *Service) SaveForumScore(ctx context.Context, activityID int, studentID int, score interface{}) error {
	u := fmt.Sprintf("/api/activities/%d/forum-scores", activityID)
	_, err := s.client.Put(ctx, u, map[string]interface{}{"student_id": studentID, "score": score}, nil)
	return err
}

// IsCategoryReplied reports whether the category has replies for the current context.
func (s *Service) IsCategoryReplied(ctx context.Context, categoryID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/forum/categories/%d/is-replied", categoryID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Replies ---

// CreateReply creates a reply to a topic.
func (s *Service) CreateReply(ctx context.Context, topicID int, body interface{}) (*model.Reply, error) {
	u := fmt.Sprintf("/api/replies/%d", topicID)
	result := new(model.Reply)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateReply updates a reply.
func (s *Service) UpdateReply(ctx context.Context, replyID int, body interface{}) (*model.Reply, error) {
	u := fmt.Sprintf("/api/replies/%d", replyID)
	result := new(model.Reply)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteReply deletes a reply.
func (s *Service) DeleteReply(ctx context.Context, replyID int) error {
	u := fmt.Sprintf("/api/replies/%d", replyID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// LikeReply likes a reply.
func (s *Service) LikeReply(ctx context.Context, replyID int) error {
	u := fmt.Sprintf("/api/replies/%d/likes", replyID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// UnlikeReply removes a like from a reply.
func (s *Service) UnlikeReply(ctx context.Context, replyID int) error {
	u := fmt.Sprintf("/api/replies/%d/likes", replyID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}
