package activities

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

type ActivityPrerequisiteQuery struct {
	ActivityType string
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
