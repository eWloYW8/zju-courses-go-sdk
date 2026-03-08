package forum

import (
	"context"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles forum-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Forum Categories ---

// GetCategory returns a forum category with its topics.
func (s *Service) GetCategory(ctx context.Context, categoryID int, opts *model.ListOptions) (*ForumCategoryResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/forum/categories/%d", categoryID), opts)
	result := new(ForumCategoryResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ExportCategoryTopics exports category topics as an Excel file.
func (s *Service) ExportCategoryTopics(ctx context.Context, categoryID int) ([]byte, error) {
	u := fmt.Sprintf("/api/categories/%d/export/excel", categoryID)
	req, err := s.client.NewRequest(ctx, "POST", u, map[string]any{})
	if err != nil {
		return nil, err
	}
	_, body, err := s.client.DoBytes(req)
	return body, err
}

// --- Topics ---

// GetTopic returns a specific topic.
func (s *Service) GetTopic(ctx context.Context, topicID int) (*Topic, error) {
	u := fmt.Sprintf("/api/forum/topics/%d", topicID)
	result := new(Topic)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateTopic creates a new topic in a category.
func (s *Service) CreateTopic(ctx context.Context, body *CreateTopicRequest) (*Topic, error) {
	result := new(Topic)
	_, err := s.client.Post(ctx, "/api/topics", body, result)
	return result, err
}

// UpdateTopic updates an existing topic.
func (s *Service) UpdateTopic(ctx context.Context, topicID int, body *UpdateTopicRequest) (*Topic, error) {
	u := fmt.Sprintf("/api/topics/%d", topicID)
	result := new(Topic)
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
func (s *Service) TopTopic(ctx context.Context, topicID int, body TopTopicRequest) error {
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
func (s *Service) ListLatestTopics(ctx context.Context, count int) ([]*Topic, error) {
	u := fmt.Sprintf("/api/topics/latest?no-intercept=true&count=%d", count)
	result := new(LatestTopicsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result.Topics, err
}

// --- Forum Score ---

// GetForumScore returns the forum score for a student.
func (s *Service) GetForumScore(ctx context.Context, activityID, studentID int) (*ForumScoreResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/students/%d/forum-score", activityID, studentID)
	result := new(ForumScoreResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetForumScores returns all forum scores for an activity.
func (s *Service) GetForumScores(ctx context.Context, activityID int) (ForumScoresResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/forum-scores", activityID)
	result := make(ForumScoresResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCourseForumScores returns forum scores for a course.
func (s *Service) GetCourseForumScores(ctx context.Context, courseID int) (CourseForumScoresResponse, error) {
	u := fmt.Sprintf("/api/course/%d/forum-scores", courseID)
	result := make(CourseForumScoresResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// SaveForumScore updates a student's forum score.
func (s *Service) SaveForumScore(ctx context.Context, activityID int, studentID int, score *float64) error {
	u := fmt.Sprintf("/api/activities/%d/forum-scores", activityID)
	_, err := s.client.Put(ctx, u, &SaveForumScoreRequest{StudentID: studentID, Score: score}, nil)
	return err
}

// IsCategoryReplied reports whether the category has replies for the current context.
func (s *Service) IsCategoryReplied(ctx context.Context, categoryID int) (CategoryRepliedResponse, error) {
	u := fmt.Sprintf("/api/forum/categories/%d/is-replied", categoryID)
	result := make(CategoryRepliedResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Replies ---

// CreateReply creates a reply to a topic.
func (s *Service) CreateReply(ctx context.Context, topicID int, body *CreateReplyRequest) (*Reply, error) {
	u := fmt.Sprintf("/api/replies/%d", topicID)
	result := new(Reply)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateReply updates a reply.
func (s *Service) UpdateReply(ctx context.Context, replyID int, body *UpdateReplyRequest) (*Reply, error) {
	u := fmt.Sprintf("/api/replies/%d", replyID)
	result := new(Reply)
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

// ListTopicCategories returns topic categories for a course.
func (s *Service) ListTopicCategories(ctx context.Context, courseID int) (*TopicCategoriesResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/topic-categories", courseID)
	result := new(TopicCategoriesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
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
