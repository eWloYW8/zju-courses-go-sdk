package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// ActivitiesService handles activity-related API operations.
type ActivitiesService struct {
	client *Client
}

// --- Response Types ---

type UploadReferencesResponse struct {
	References []*model.UploadReference `json:"references"`
}

type CommentsResponse struct {
	Comments []*model.Comment `json:"comments"`
	model.Pagination
}

type CommentPageCountResponse struct {
	PageStats []interface{} `json:"page_stats"`
}

type RecommendSubmissionsResponse struct {
	Submissions []*model.Submission `json:"submissions"`
}

// --- Activity CRUD ---

// GetActivity returns detailed information about an activity.
func (s *ActivitiesService) GetActivity(ctx context.Context, activityID int) (*model.Activity, error) {
	u := fmt.Sprintf("/api/activities/%d", activityID)
	result := new(model.Activity)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// CreateActivity creates a new activity in a course.
func (s *ActivitiesService) CreateActivity(ctx context.Context, courseID int, activity interface{}) (*model.Activity, error) {
	u := fmt.Sprintf("/api/course/activities/%d", courseID)
	result := new(model.Activity)
	_, err := s.client.post(ctx, u, activity, result)
	return result, err
}

// UpdateActivity updates an existing activity.
func (s *ActivitiesService) UpdateActivity(ctx context.Context, activityID int, activity interface{}) (*model.Activity, error) {
	u := fmt.Sprintf("/api/activities/%d", activityID)
	result := new(model.Activity)
	_, err := s.client.put(ctx, u, activity, result)
	return result, err
}

// DeleteActivity deletes an activity.
func (s *ActivitiesService) DeleteActivity(ctx context.Context, activityID int) error {
	u := fmt.Sprintf("/api/activities/%d", activityID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// DeleteCheck checks if an activity can be deleted.
func (s *ActivitiesService) DeleteCheck(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/activities/delete-check?activity_id=%d", activityID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// HaveDependents checks if activities have dependents.
func (s *ActivitiesService) HaveDependents(ctx context.Context, activityIDs []int) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/activities/have-dependents", map[string][]int{"activity_ids": activityIDs}, &result)
	return result, err
}

// --- Activity Read Status ---

// GetActivityRead returns the read status for an activity.
func (s *ActivitiesService) GetActivityRead(ctx context.Context, activityID int) (*model.ActivityRead, error) {
	u := fmt.Sprintf("/api/course/activities-read/%d", activityID)
	result := new(model.ActivityRead)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// MarkActivityRead marks an activity as read.
func (s *ActivitiesService) MarkActivityRead(ctx context.Context, activityID int) (*model.ActivityRead, error) {
	u := fmt.Sprintf("/api/course/activity-read/%d", activityID)
	result := new(model.ActivityRead)
	_, err := s.client.post(ctx, u, nil, result)
	return result, err
}

// --- Activity Lock Status ---

// CheckIsLocked checks if activities are locked.
// activityConditions format: "activityID1,activityID2,..."
func (s *ActivitiesService) CheckIsLocked(ctx context.Context, activityConditions string) (map[string]*model.IsLockedStatus, error) {
	u := fmt.Sprintf("/api/activities/is-locked?activity_conditions=%s", activityConditions)
	result := make(map[string]*model.IsLockedStatus)
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Comments ---

// ListComments returns comments for an activity.
func (s *ActivitiesService) ListComments(ctx context.Context, activityID int, opts *model.ListOptions) (*CommentsResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/activities/%d/comments", activityID), opts)
	result := new(CommentsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// CreateComment creates a comment on an activity.
func (s *ActivitiesService) CreateComment(ctx context.Context, activityID int, comment interface{}) (*model.Comment, error) {
	u := fmt.Sprintf("/api/activities/%d/comments", activityID)
	result := new(model.Comment)
	_, err := s.client.post(ctx, u, comment, result)
	return result, err
}

// OperateComment performs an operation (like/unlike) on a comment.
func (s *ActivitiesService) OperateComment(ctx context.Context, activityID int, body interface{}) error {
	u := fmt.Sprintf("/api/activities/%d/comments/operate", activityID)
	_, err := s.client.post(ctx, u, body, nil)
	return err
}

// GetCommentCount returns comment counts for an activity.
func (s *ActivitiesService) GetCommentCount(ctx context.Context, activityID int) (*model.CommentCount, error) {
	u := fmt.Sprintf("/api/activities/%d/comment/count", activityID)
	result := new(model.CommentCount)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetCommentPageCount returns page-level comment stats.
func (s *ActivitiesService) GetCommentPageCount(ctx context.Context, activityID int) (*CommentPageCountResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/comment/page-count", activityID)
	result := new(CommentPageCountResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Upload References ---

// ListUploadReferences returns upload references for an activity.
func (s *ActivitiesService) ListUploadReferences(ctx context.Context, activityID int) (*UploadReferencesResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/upload_references", activityID)
	result := new(UploadReferencesResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Recommend Submissions ---

// ListRecommendSubmissions returns recommended submissions for an activity.
func (s *ActivitiesService) ListRecommendSubmissions(ctx context.Context, activityID int) (*RecommendSubmissionsResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/recommend-submissions", activityID)
	result := new(RecommendSubmissionsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Activity Resources ---

// ListActivityResources returns resources for an activity.
func (s *ActivitiesService) ListActivityResources(ctx context.Context, activityID int) ([]interface{}, error) {
	u := fmt.Sprintf("/api/activities/resources/%d", activityID)
	var result []interface{}
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Classin Integration ---

// GetClassinJoinURL returns the Classin join URL for a course.
func (s *ActivitiesService) GetClassinJoinURL(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/activies/classin/join-url?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetClassinWebcastURL returns the Classin webcast URL.
func (s *ActivitiesService) GetClassinWebcastURL(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/activities/classin/webcast-url?activity_id=%d", activityID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}
