package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
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
	return s.GetActivityWithOptions(ctx, activityID, nil)
}

// GetActivityWithOptions returns detailed information about an activity with optional query fields.
func (s *Service) GetActivityWithOptions(ctx context.Context, activityID int, opts *GetActivityOptions) (*Activity, error) {
	u := fmt.Sprintf("/api/activities/%d", activityID)
	if opts != nil {
		values := url.Values{}
		if opts.Fields != "" {
			values.Set("fields", opts.Fields)
		}
		if opts.SubCourseID != nil {
			values.Set("sub_course_id", strconv.Itoa(*opts.SubCourseID))
		}
		if encoded := values.Encode(); encoded != "" {
			u += "?" + encoded
		}
	}
	result := new(Activity)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateActivity creates a new activity in a course.
func (s *Service) CreateActivity(ctx context.Context, courseID int, activity interface{}) (*Activity, error) {
	u := fmt.Sprintf("/api/courses/%d/activities", courseID)
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
	return s.DeleteActivityWithOptions(ctx, activityID, nil)
}

// DeleteActivityWithOptions deletes an activity with optional query parameters.
func (s *Service) DeleteActivityWithOptions(ctx context.Context, activityID int, opts *DeleteActivityOptions) error {
	u := fmt.Sprintf("/api/activities/%d", activityID)
	if opts != nil {
		values := url.Values{}
		if opts.ActivityType != "" {
			values.Set("activity_type", opts.ActivityType)
		}
		if opts.KeepOriginal {
			values.Set("keep_original", "true")
		}
		if encoded := values.Encode(); encoded != "" {
			u += "?" + encoded
		}
	}
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// BatchDeleteActivities deletes multiple activities in one request.
func (s *Service) BatchDeleteActivities(ctx context.Context, activityIDs []int) error {
	_, err := s.client.DeleteWithBody(ctx, "/api/activities", &BatchDeleteActivitiesRequest{ActivityIDs: activityIDs}, nil)
	return err
}

// DeleteCheck checks if an activity can be deleted.
func (s *Service) DeleteCheck(ctx context.Context, activityID int) (*DeleteCheckResponse, error) {
	return s.DeleteCheckWithType(ctx, activityID, "")
}

// DeleteCheckWithType checks if an activity of a specific type can be deleted.
func (s *Service) DeleteCheckWithType(ctx context.Context, activityID int, activityType string) (*DeleteCheckResponse, error) {
	values := url.Values{}
	values.Set("activity_id", strconv.Itoa(activityID))
	if activityType != "" {
		values.Set("activity_type", activityType)
	}
	u := "/api/activities/delete-check?" + values.Encode()
	result := new(DeleteCheckResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// HaveDependents checks if activities have dependents.
func (s *Service) HaveDependents(ctx context.Context, activityIDs []int) (*HaveDependentsResponse, error) {
	return s.HaveDependentsWithType(ctx, activityIDs, "")
}

// HaveDependentsWithType checks if activities of a specific type have dependents.
func (s *Service) HaveDependentsWithType(ctx context.Context, activityIDs []int, activityType string) (*HaveDependentsResponse, error) {
	values := url.Values{}
	for _, activityID := range activityIDs {
		values.Add("activity_ids", strconv.Itoa(activityID))
	}
	if activityType != "" {
		values.Set("activity_type", activityType)
	}
	u := "/api/activities/have-dependents"
	if encoded := values.Encode(); encoded != "" {
		u += "?" + encoded
	}
	result := new(HaveDependentsResponse)
	_, err := s.client.Get(ctx, u, result)
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
	u := fmt.Sprintf("/api/course/activities-read/%d", activityID)
	result := new(ActivityRead)
	_, err := s.client.Post(ctx, u, nil, result)
	return result, err
}

// LogActivityRead logs activity read progress using the frontend JSON body shape.
func (s *Service) LogActivityRead(ctx context.Context, activityID int, body interface{}) (ActivityReadLogResponse, error) {
	u := fmt.Sprintf("/api/course/activities-read/%d", activityID)
	result := make(ActivityReadLogResponse)
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// GetStudentSubmission returns a student's current submission for a course activity.
func (s *Service) GetStudentSubmission(ctx context.Context, activityID, studentID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/activities/%d/students/%d/submission", activityID, studentID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateCourseActivitySubmission creates a submission draft or final submission for a course activity.
func (s *Service) CreateCourseActivitySubmission(ctx context.Context, activityID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/activities/%d/submissions", activityID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// UpdateCourseActivitySubmission updates a submission draft for a course activity.
func (s *Service) UpdateCourseActivitySubmission(ctx context.Context, activityID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/activities/%d/submissions", activityID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// ScoreCourseActivitySubmission scores a course activity submission.
func (s *Service) ScoreCourseActivitySubmission(ctx context.Context, activityID int, body interface{}) (json.RawMessage, error) {
	values := url.Values{}
	values.Set("fields", "id,score,instructor_comment,rubric_score,final_score")
	values.Set("need_submission_correct", "true")
	u := fmt.Sprintf("/api/course/activities/%d/submission/score?%s", activityID, values.Encode())
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// SaveWebLinkScore updates a student's score for a web link activity.
func (s *Service) SaveWebLinkScore(ctx context.Context, activityID, studentID int, score interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/activities/%d/web_link_score", activityID)
	body := map[string]interface{}{
		"student_id": studentID,
		"score":      score,
	}
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// LogExamActivityRead logs exam activity read progress.
func (s *Service) LogExamActivityRead(ctx context.Context, activityID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/activities-read/exam/%d", activityID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// GetActivityCompletionCriteria returns available completion criteria for a course/activity-type pair.
func (s *Service) GetActivityCompletionCriteria(ctx context.Context, courseID int, query *ActivityCriteriaQuery) (*ActivityCompletionCriteriaResponse, error) {
	values := url.Values{}
	if query != nil {
		if query.ActivityType != "" {
			values.Set("activity_type", query.ActivityType)
		}
		if query.CourseID > 0 {
			values.Set("course_id", strconv.Itoa(query.CourseID))
		}
	} else if courseID > 0 {
		values.Set("course_id", strconv.Itoa(courseID))
	}
	values.Set("no-intercept", "true")
	u := "/api/completion-criteria"
	if encoded := values.Encode(); encoded != "" {
		u += "?" + encoded
	}
	result := new(ActivityCompletionCriteriaResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetActivityCompletionCriteriaDetail returns completion criteria for a specific activity.
func (s *Service) GetActivityCompletionCriteriaDetail(ctx context.Context, activityID int, query *ActivityCompletionCriteriaDetailQuery) (*ActivityCompletionCriteriaResponse, error) {
	values := url.Values{}
	if query != nil {
		if query.ActivityType != "" {
			values.Set("activity_type", query.ActivityType)
		}
		if query.CourseID > 0 {
			values.Set("course_id", strconv.Itoa(query.CourseID))
		}
	}
	values.Set("no-intercept", "true")
	u := fmt.Sprintf("/api/activities/%d/completion-criteria", activityID)
	if encoded := values.Encode(); encoded != "" {
		u += "?" + encoded
	}
	result := new(ActivityCompletionCriteriaResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListCourseActivities returns activities for a course.
func (s *Service) ListCourseActivities(ctx context.Context, courseID int) ([]*Activity, error) {
	u := fmt.Sprintf("/api/courses/%d/activities", courseID)
	result := new(CourseActivitiesResponse)
	_, err := s.client.Get(ctx, u, result)
	if err != nil {
		return nil, err
	}
	return result.Activities, nil
}

// ListPrerequisites returns prerequisite activities for a specific activity.
func (s *Service) ListPrerequisites(ctx context.Context, activityID int, activityType string) (*ActivityPrerequisitesResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/prerequisites", activityID)
	if activityType != "" {
		u += "?activity_type=" + url.QueryEscape(activityType) + "&no-intercept=true"
	}
	result := new(ActivityPrerequisitesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// DeletePrerequisites removes prerequisites for a specific activity.
func (s *Service) DeletePrerequisites(ctx context.Context, activityID int, activityType string) error {
	u := fmt.Sprintf("/api/activities/%d/prerequisites", activityID)
	if activityType != "" {
		u += "?activity_type=" + url.QueryEscape(activityType)
	}
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// GetUnavailablePrerequisites returns unavailable prerequisites for an activity.
func (s *Service) GetUnavailablePrerequisites(ctx context.Context, activityID int, activityType string) (*HaveDependentsResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/unavailable-prerequisites", activityID)
	if activityType != "" {
		u += "?activity_type=" + url.QueryEscape(activityType)
	}
	result := new(HaveDependentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetActivityHasDependents returns whether a single activity has dependents.
func (s *Service) GetActivityHasDependents(ctx context.Context, activityID int, activityType string) (*HaveDependentsResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/has-dependents", activityID)
	if activityType != "" {
		u += "?activity_type=" + url.QueryEscape(activityType)
	}
	result := new(HaveDependentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// RemoveAllDependants removes all dependent relationships for an activity.
func (s *Service) RemoveAllDependants(ctx context.Context, activityID int, activityType string) error {
	u := fmt.Sprintf("/api/activities/%d/remove-all-dependants", activityID)
	body := map[string]string{}
	if activityType != "" {
		body["activity_type"] = activityType
	}
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// GetSubmissionNumber returns submission-number stats for an activity.
func (s *Service) GetSubmissionNumber(ctx context.Context, activityID int) (ActivitySubmissionNumberResponse, error) {
	u := fmt.Sprintf("/api/activity/%d/submission-number", activityID)
	result := make(ActivitySubmissionNumberResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// HasReviewedInterScore returns inter-review status for an activity.
func (s *Service) HasReviewedInterScore(ctx context.Context, activityID int) (ActivityReviewStateResponse, error) {
	u := fmt.Sprintf("/api/activity/%d/has-reviewed-inter-score", activityID)
	result := make(ActivityReviewStateResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetHomeworkFirstSubmissionTime returns first-submission-time data for a homework activity.
func (s *Service) GetHomeworkFirstSubmissionTime(ctx context.Context, activityID int) (HomeworkFirstSubmissionTimeResponse, error) {
	u := fmt.Sprintf("/api/activity/%d/first-submission-time", activityID)
	result := make(HomeworkFirstSubmissionTimeResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// IsInterReviewStarted returns whether inter-review has started.
func (s *Service) IsInterReviewStarted(ctx context.Context, activityID int) (ActivityReviewStateResponse, error) {
	u := fmt.Sprintf("/api/activity/%d/is-inter-review-started", activityID)
	result := make(ActivityReviewStateResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// IsIntraReviewStarted returns whether intra-review has started.
func (s *Service) IsIntraReviewStarted(ctx context.Context, activityID int) (ActivityReviewStateResponse, error) {
	u := fmt.Sprintf("/api/activity/%d/is-intra-review-started", activityID)
	result := make(ActivityReviewStateResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// IsHomeworkExpired returns whether a homework activity is expired.
func (s *Service) IsHomeworkExpired(ctx context.Context, activityID int) (ActivityReviewStateResponse, error) {
	u := fmt.Sprintf("/api/activity/%d/is-homework-expired", activityID)
	result := make(ActivityReviewStateResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetActivityUploadsLicenseInfo returns uploads-license info for an activity.
func (s *Service) GetActivityUploadsLicenseInfo(ctx context.Context, activityID int) (ActivityUploadsLicenseInfoResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/uploads-license", activityID)
	result := make(ActivityUploadsLicenseInfoResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetOnlineVideoActivityReadCount returns read-count stats for an online-video activity.
func (s *Service) GetOnlineVideoActivityReadCount(ctx context.Context, activityID int) (OnlineVideoActivityReadCountResponse, error) {
	u := fmt.Sprintf("/api/online-videos/%d/activity-read-count", activityID)
	result := make(OnlineVideoActivityReadCountResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetWebLinkScores returns score data for a web-link activity.
func (s *Service) GetWebLinkScores(ctx context.Context, activityID int) (ActivityScoresResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/web-link-scores", activityID)
	result := make(ActivityScoresResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVirtualExperimentScores returns score data for a virtual-experiment activity.
func (s *Service) GetVirtualExperimentScores(ctx context.Context, activityID int) (ActivityScoresResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/virtual-experiment-scores", activityID)
	result := make(ActivityScoresResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Activity Lock Status ---

// CheckIsLocked checks if activities are locked.
// activityConditions should be the JSON-encoded frontend payload used by
// `/api/activities/is-locked`, for example:
// `[{"activity_id":1,"course_id":2,"activity_type":"homework"}]`.
func (s *Service) CheckIsLocked(ctx context.Context, activityConditions string) (map[string]*IsLockedStatus, error) {
	u := fmt.Sprintf("/api/activities/is-locked?activity_conditions=%s", activityConditions)
	result := make(map[string]*IsLockedStatus)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CheckIsLockedWithConditions checks lock state using structured activity conditions.
func (s *Service) CheckIsLockedWithConditions(ctx context.Context, conditions []*ActivityLockCondition) (map[string]*IsLockedStatus, error) {
	body, err := json.Marshal(conditions)
	if err != nil {
		return nil, err
	}
	return s.CheckIsLocked(ctx, string(body))
}

// ConvertOnlineVideoToInteraction converts an online-video activity into an interaction activity.
func (s *Service) ConvertOnlineVideoToInteraction(ctx context.Context, activityID int) (*ConvertedInteractionResponse, error) {
	u := fmt.Sprintf("/api/online-videos/%d/convert-to-interaction", activityID)
	result := new(ConvertedInteractionResponse)
	_, err := s.client.Post(ctx, u, nil, result)
	return result, err
}

// --- Comments ---

// ListComments returns comments for an activity.
func (s *Service) ListComments(ctx context.Context, activityID int, opts *model.ListOptions) (*CommentsResponse, error) {
	params := CommentListParams{}
	if opts != nil {
		params.Page = opts.Page
		params.PageSize = opts.PageSize
	}
	return s.ListCommentsWithParams(ctx, activityID, params)
}

// ListCommentsWithParams returns comments for an activity with frontend ordering and condition filters.
func (s *Service) ListCommentsWithParams(ctx context.Context, activityID int, params CommentListParams) (*CommentsResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/activities/%d/comments", activityID), &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	query := map[string]string{}
	if params.OrderKey != "" {
		query["order_key"] = params.OrderKey
	}
	if params.Order != "" {
		query["order"] = params.Order
	}
	if params.Conditions != nil {
		body, err := json.Marshal(params.Conditions)
		if err != nil {
			return nil, err
		}
		query["conditions"] = string(body)
	}
	u = addQueryParams(u, query)
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

// UpdateComment updates an activity comment.
func (s *Service) UpdateComment(ctx context.Context, activityID, commentID int, body *CreateCommentRequest) (*Comment, error) {
	u := fmt.Sprintf("/api/activities/%d/comments/%d", activityID, commentID)
	result := new(Comment)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteComment deletes an activity comment.
func (s *Service) DeleteComment(ctx context.Context, activityID, commentID int) error {
	u := fmt.Sprintf("/api/activities/%d/comments/%d", activityID, commentID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ListCommentRepliesByCommentIDs returns replies for a batch of comment ids.
func (s *Service) ListCommentRepliesByCommentIDs(ctx context.Context, activityID int, commentIDs []int) ([]*Comment, error) {
	values := url.Values{}
	for _, id := range commentIDs {
		values.Add("comment_ids[]", strconv.Itoa(id))
	}
	u := fmt.Sprintf("/api/activities/%d/comments/replies", activityID)
	if encoded := values.Encode(); encoded != "" {
		u += "?" + encoded
	}
	var result []*Comment
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListCommentReplies returns paged replies for a single comment.
func (s *Service) ListCommentReplies(ctx context.Context, activityID, commentID int, opts *model.ListOptions) ([]*Comment, error) {
	u := addListOptions(fmt.Sprintf("/api/activities/%d/comments/%d/replies", activityID, commentID), opts)
	var result []*Comment
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ReplyComment creates a reply under a comment.
func (s *Service) ReplyComment(ctx context.Context, activityID, commentID int, body *CreateCommentRequest) (*Comment, error) {
	u := fmt.Sprintf("/api/activities/%d/comments/%d/reply", activityID, commentID)
	result := new(Comment)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateCommentReply updates a comment reply.
func (s *Service) UpdateCommentReply(ctx context.Context, activityID, replyID int, body *CreateCommentRequest) (*Comment, error) {
	u := fmt.Sprintf("/api/activities/%d/reply/%d", activityID, replyID)
	result := new(Comment)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteCommentReply deletes a comment reply.
func (s *Service) DeleteCommentReply(ctx context.Context, activityID, replyID int) error {
	u := fmt.Sprintf("/api/activities/%d/reply/%d", activityID, replyID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// OperateComment performs an operation (like/unlike) on a comment.
func (s *Service) OperateComment(ctx context.Context, activityID int, body *OperateCommentRequest) error {
	u := fmt.Sprintf("/api/activities/%d/comments/operate", activityID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// GetCommentCount returns comment counts for an activity.
func (s *Service) GetCommentCount(ctx context.Context, activityID int) (*CommentCount, error) {
	return s.GetCommentCountWithConditions(ctx, activityID, nil)
}

// GetCommentCountWithConditions returns comment counts for an activity with frontend conditions.
func (s *Service) GetCommentCountWithConditions(ctx context.Context, activityID int, conditions any) (*CommentCount, error) {
	u := fmt.Sprintf("/api/activities/%d/comment/count", activityID)
	if conditions != nil {
		body, err := json.Marshal(conditions)
		if err != nil {
			return nil, err
		}
		u += "?conditions=" + url.QueryEscape(string(body))
	}
	result := new(CommentCount)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetCommentPageCount returns page-level comment stats.
func (s *Service) GetCommentPageCount(ctx context.Context, activityID int) (*CommentPageCountResponse, error) {
	return s.GetCommentPageCountWithConditions(ctx, activityID, nil)
}

// GetCommentPageCountWithConditions returns page-level comment stats with frontend conditions.
func (s *Service) GetCommentPageCountWithConditions(ctx context.Context, activityID int, conditions any) (*CommentPageCountResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/comment/page-count", activityID)
	if conditions != nil {
		body, err := json.Marshal(conditions)
		if err != nil {
			return nil, err
		}
		u += "?conditions=" + url.QueryEscape(string(body))
	}
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
	u := fmt.Sprintf("/api/activities/%d/resources", activityID)
	var result []interface{}
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// DeleteActivityResource deletes a resource under an activity.
func (s *Service) DeleteActivityResource(ctx context.Context, activityID, resourceID int) error {
	u := fmt.Sprintf("/api/activities/%d/resources/%d", activityID, resourceID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// UpdateActivityResource updates a resource under an activity.
func (s *Service) UpdateActivityResource(ctx context.Context, activityID, resourceID int, body UpdateActivityResourceRequest) error {
	u := fmt.Sprintf("/api/activities/%d/resources/%d", activityID, resourceID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// SaveActivityResource saves resources under an activity.
func (s *Service) SaveActivityResource(ctx context.Context, activityID int) error {
	u := fmt.Sprintf("/api/activities/resources/%d/save", activityID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// --- Classin Integration ---

// GetClassinJoinURL returns the Classin join URL for a course.
func (s *Service) GetClassinJoinURL(ctx context.Context, courseID int) (json.RawMessage, error) {
	result, err := s.GetClassinJoinURLWithParams(ctx, ClassinJoinURLParams{CourseID: courseID})
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(body), nil
}

// GetClassinJoinURLWithParams returns the typed Classin join URL payload.
func (s *Service) GetClassinJoinURLWithParams(ctx context.Context, params ClassinJoinURLParams) (*ClassinURLResponse, error) {
	values := url.Values{}
	if params.CourseID > 0 {
		values.Set("course_id", strconv.Itoa(params.CourseID))
	}
	if params.ActivityID > 0 {
		values.Set("activity_id", strconv.Itoa(params.ActivityID))
	}
	if params.UserID > 0 {
		values.Set("user_id", strconv.Itoa(params.UserID))
	}
	u := "/api/activies/classin/join-url"
	if encoded := values.Encode(); encoded != "" {
		u += "?" + encoded
	}
	result := new(ClassinURLResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetClassinWebcastURL returns the raw Classin webcast URL payload.
func (s *Service) GetClassinWebcastURL(ctx context.Context, activityID int) (json.RawMessage, error) {
	result, err := s.GetClassinWebcastURLWithParams(ctx, ClassinWebcastURLParams{ActivityID: activityID})
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(body), nil
}

// GetClassinWebcastURLWithParams returns the typed Classin webcast URL payload.
func (s *Service) GetClassinWebcastURLWithParams(ctx context.Context, params ClassinWebcastURLParams) (*ClassinURLResponse, error) {
	values := url.Values{}
	if params.CourseID > 0 {
		values.Set("course_id", strconv.Itoa(params.CourseID))
	}
	if params.ActivityID > 0 {
		values.Set("activity_id", strconv.Itoa(params.ActivityID))
	}
	u := "/api/activities/classin/webcast-url"
	if encoded := values.Encode(); encoded != "" {
		u += "?" + encoded
	}
	result := new(ClassinURLResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetScoreRecords returns score-change records for a submitter.
func (s *Service) GetScoreRecords(ctx context.Context, activityID, submitterID int, params ActivityScoreRecordsParams) (*ScoreRecordsPage, error) {
	values := url.Values{}
	if params.Page > 0 {
		values.Set("page", strconv.Itoa(params.Page))
	}
	if params.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(params.PageSize))
	}
	u := fmt.Sprintf("/api/activity/%d/submitter/%d/score-records", activityID, submitterID)
	if encoded := values.Encode(); encoded != "" {
		u += "?" + encoded
	}
	result := new(ScoreRecordsResponse)
	_, err := s.client.Get(ctx, u, result)
	if err != nil {
		return nil, err
	}
	return &ScoreRecordsPage{
		Items:      result.Records,
		Pagination: result.Pagination,
		Start:      result.Start,
		End:        result.End,
	}, nil
}

// ListGradeScoreItems returns frontend grade-score items for a course.
func (s *Service) ListGradeScoreItems(ctx context.Context, courseID int) ([]*GradeScoreItem, error) {
	u := fmt.Sprintf("/api/courses/%d/grade-score-items", courseID)
	var result []*GradeScoreItem
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
