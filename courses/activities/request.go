package activities

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type CreateActivityRequest = Activity

type UpdateActivityRequest = Activity

type GetActivityOptions struct {
	Fields      string
	SubCourseID *int
}

type DeleteActivityOptions struct {
	ActivityType string
	KeepOriginal bool
}

type BatchDeleteActivitiesRequest struct {
	ActivityIDs []int `json:"activity_ids"`
}

type HaveDependentsRequest struct {
	ActivityIDs  []int
	ActivityType string
}

type ActivityCriteriaQuery struct {
	ActivityType string
	CourseID     int
}

type ActivityCompletionCriteriaDetailQuery struct {
	ActivityType string
	CourseID     int
}

type ActivityPrerequisiteQuery struct {
	ActivityType string
}

type ActivityLockCondition struct {
	ActivityID   int    `json:"activity_id"`
	CourseID     int    `json:"course_id,omitempty"`
	ActivityType string `json:"activity_type,omitempty"`
}

type LogExamActivityReadRequest struct {
	Data map[string]any `json:"data,omitempty"`
}

type CreateCommentRequest struct {
	Content   string         `json:"content,omitempty"`
	UploadIDs []int          `json:"upload_ids,omitempty"`
	ParentID  *int           `json:"parent_id,omitempty"`
	Data      map[string]any `json:"data,omitempty"`
}

type OperateCommentRequest struct {
	CommentID int    `json:"comment_id"`
	Action    string `json:"action"`
}

type CommentListParams struct {
	Page       int
	PageSize   int
	OrderKey   string
	Order      string
	Conditions any
}

type ActivityScoreRecordsParams struct {
	Page     int
	PageSize int
}

type ClassinJoinURLParams struct {
	CourseID   int
	ActivityID int
	UserID     int
}

type ClassinWebcastURLParams struct {
	CourseID   int
	ActivityID int
}

type UpdateActivityResourceRequest map[string]any

type ScoreRecordsPage struct {
	Items []*ScoreRecord
	model.Pagination
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}
