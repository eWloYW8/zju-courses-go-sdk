package forum

type CreateTopicRequest struct {
	Title           string         `json:"title,omitempty"`
	Content         string         `json:"content,omitempty"`
	TopicCategoryID int            `json:"topic_category_id,omitempty"`
	UploadIDs       []int          `json:"upload_ids,omitempty"`
	GroupID         *int           `json:"group_id,omitempty"`
	Data            map[string]any `json:"data,omitempty"`
}

type UpdateTopicRequest struct {
	Title     string         `json:"title,omitempty"`
	Content   string         `json:"content,omitempty"`
	UploadIDs []int          `json:"upload_ids,omitempty"`
	Data      map[string]any `json:"data,omitempty"`
}

type TopTopicRequest map[string]any

type SaveForumScoreRequest struct {
	StudentID int      `json:"student_id"`
	Score     *float64 `json:"score,omitempty"`
}

type CreateReplyRequest struct {
	Content   string         `json:"content,omitempty"`
	UploadIDs []int          `json:"upload_ids,omitempty"`
	ParentID  *int           `json:"parent_id,omitempty"`
	Data      map[string]any `json:"data,omitempty"`
}

type UpdateReplyRequest struct {
	Content   string         `json:"content,omitempty"`
	UploadIDs []int          `json:"upload_ids,omitempty"`
	Data      map[string]any `json:"data,omitempty"`
}
