package activities

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// Service handles activity-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Activity CRUD ---

// GetActivity returns detailed information about an activity.
func (s *Service) GetActivity(ctx context.Context, activityID int) (*Activity, error) {
	u := fmt.Sprintf("/api/activities/%d", activityID)
	result := new(Activity)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateActivity creates a new activity in a course.
func (s *Service) CreateActivity(ctx context.Context, courseID int, activity interface{}) (*Activity, error) {
	u := fmt.Sprintf("/api/course/activities/%d", courseID)
	result := new(Activity)
	_, err := s.client.Post(ctx, u, activity, result)
	return result, err
}

// UpdateActivity updates an existing activity.
func (s *Service) UpdateActivity(ctx context.Context, activityID int, activity interface{}) (*Activity, error) {
	u := fmt.Sprintf("/api/activities/%d", activityID)
	result := new(Activity)
	_, err := s.client.Put(ctx, u, activity, result)
	return result, err
}

// DeleteActivity deletes an activity.
func (s *Service) DeleteActivity(ctx context.Context, activityID int) error {
	u := fmt.Sprintf("/api/activities/%d", activityID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// DeleteCheck checks if an activity can be deleted.
func (s *Service) DeleteCheck(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/activities/delete-check?activity_id=%d", activityID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// HaveDependents checks if activities have dependents.
func (s *Service) HaveDependents(ctx context.Context, activityIDs []int) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/activities/have-dependents", map[string][]int{"activity_ids": activityIDs}, &result)
	return result, err
}

// --- Activity Read Status ---

// GetActivityRead returns the read status for an activity.
func (s *Service) GetActivityRead(ctx context.Context, activityID int) (*ActivityRead, error) {
	u := fmt.Sprintf("/api/course/activities-read/%d", activityID)
	result := new(ActivityRead)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// MarkActivityRead marks an activity as read.
func (s *Service) MarkActivityRead(ctx context.Context, activityID int) (*ActivityRead, error) {
	u := fmt.Sprintf("/api/course/activity-read/%d", activityID)
	result := new(ActivityRead)
	_, err := s.client.Post(ctx, u, nil, result)
	return result, err
}

// LogExamActivityRead logs exam activity read progress.
func (s *Service) LogExamActivityRead(ctx context.Context, activityID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/activities-read/exam/%d", activityID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// --- Activity Lock Status ---

// CheckIsLocked checks if activities are locked.
// activityConditions format: "activityID1,activityID2,..."
func (s *Service) CheckIsLocked(ctx context.Context, activityConditions string) (map[string]*IsLockedStatus, error) {
	u := fmt.Sprintf("/api/activities/is-locked?activity_conditions=%s", activityConditions)
	result := make(map[string]*IsLockedStatus)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Comments ---

// ListComments returns comments for an activity.
func (s *Service) ListComments(ctx context.Context, activityID int, opts *model.ListOptions) (*CommentsResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/activities/%d/comments", activityID), opts)
	result := new(CommentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateComment creates a comment on an activity.
func (s *Service) CreateComment(ctx context.Context, activityID int, comment *CreateCommentRequest) (*Comment, error) {
	u := fmt.Sprintf("/api/activities/%d/comments", activityID)
	result := new(Comment)
	_, err := s.client.Post(ctx, u, comment, result)
	return result, err
}

// OperateComment performs an operation (like/unlike) on a comment.
func (s *Service) OperateComment(ctx context.Context, activityID int, body *OperateCommentRequest) error {
	u := fmt.Sprintf("/api/activities/%d/comments/operate", activityID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// GetCommentCount returns comment counts for an activity.
func (s *Service) GetCommentCount(ctx context.Context, activityID int) (*CommentCount, error) {
	u := fmt.Sprintf("/api/activities/%d/comment/count", activityID)
	result := new(CommentCount)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetCommentPageCount returns page-level comment stats.
func (s *Service) GetCommentPageCount(ctx context.Context, activityID int) (*CommentPageCountResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/comment/page-count", activityID)
	result := new(CommentPageCountResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Upload References ---

// ListUploadReferences returns upload references for an activity.
func (s *Service) ListUploadReferences(ctx context.Context, activityID int) (*UploadReferencesResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/upload_references", activityID)
	result := new(UploadReferencesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Recommend Submissions ---

// ListRecommendSubmissions returns recommended submissions for an activity.
func (s *Service) ListRecommendSubmissions(ctx context.Context, activityID int) (*RecommendSubmissionsResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/recommend-submissions", activityID)
	result := new(RecommendSubmissionsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Activity Resources ---

// ListActivityResources returns resources for an activity.
func (s *Service) ListActivityResources(ctx context.Context, activityID int) ([]interface{}, error) {
	u := fmt.Sprintf("/api/activities/resources/%d", activityID)
	var result []interface{}
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Classin Integration ---

// GetClassinJoinURL returns the Classin join URL for a course.
func (s *Service) GetClassinJoinURL(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/activies/classin/join-url?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetClassinWebcastURL returns the Classin webcast URL.
func (s *Service) GetClassinWebcastURL(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/activities/classin/webcast-url?activity_id=%d", activityID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
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
