package activities

type CreateActivityRequest = Activity

type UpdateActivityRequest = Activity

type LogExamActivityReadRequest struct {
	Data map[string]any `json:"data,omitempty"`
}

type CreateCommentRequest struct {
	Content   string   `json:"content,omitempty"`
	UploadIDs []int    `json:"upload_ids,omitempty"`
	ParentID  *int     `json:"parent_id,omitempty"`
	Data      map[string]any `json:"data,omitempty"`
}

type OperateCommentRequest struct {
	CommentID int    `json:"comment_id"`
	Action    string `json:"action"`
}
