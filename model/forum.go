package model

// TopicCategory represents a forum topic category.
type TopicCategory struct {
	ID           int       `json:"id"`
	Title        string    `json:"title,omitempty"`
	ReferrerType string    `json:"referrer_type,omitempty"`
	Activity     *Activity `json:"activity,omitempty"`
}

// Topic represents a forum topic/post.
type Topic struct {
	ID                int         `json:"id"`
	Title             string      `json:"title,omitempty"`
	Content           string      `json:"content,omitempty"`
	Summary           string      `json:"summary,omitempty"`
	TopicType         string      `json:"topic_type,omitempty"`
	CreatedAt         string      `json:"created_at,omitempty"`
	UpdatedAt         string      `json:"updated_at,omitempty"`
	CreatedBy         *User       `json:"created_by,omitempty"`
	GroupID           *int        `json:"group_id,omitempty"`
	TopicCategoryID   int         `json:"topic_category_id,omitempty"`
	InCommonCategory  bool        `json:"in_common_category,omitempty"`
	CurrentUserLiked  bool        `json:"current_user_liked,omitempty"`
	CurrentUserRead   bool        `json:"current_user_read,omitempty"`
	IsHotTopic        bool        `json:"is_hot_topic,omitempty"`
	IsTeacherTop      bool        `json:"is_teacher_top,omitempty"`
	TopicToppedSort   *int        `json:"topic_topped_sort,omitempty"`
	LikeCount         int         `json:"like_count,omitempty"`
	ReplyCount        int         `json:"reply_count,omitempty"`
	ReadReplies       int         `json:"read_replies,omitempty"`
	UnreadReplyCount  int         `json:"unread_reply_count,omitempty"`
	HasMatchedReplies bool        `json:"has_matched_replies,omitempty"`
	UserVisitsNumber  int         `json:"user_visits_number,omitempty"`
	Enrollments       interface{} `json:"enrollments,omitempty"`
	Uploads           []*Upload   `json:"uploads,omitempty"`
}

// ForumCategoryResponse represents a forum category with paginated topics.
type ForumCategoryResponse struct {
	ID         int                  `json:"id"`
	ActivityID int                  `json:"activity_id"`
	Result     *ForumCategoryResult `json:"result,omitempty"`
}

// ForumCategoryResult represents the paginated topics result.
type ForumCategoryResult struct {
	Page     int      `json:"page"`
	PageSize int      `json:"page_size"`
	Pages    int      `json:"pages"`
	Total    int      `json:"total"`
	Start    int      `json:"start"`
	End      int      `json:"end"`
	Topics   []*Topic `json:"topics,omitempty"`
}

// Reply represents a reply to a forum topic.
type Reply struct {
	ID        int         `json:"id"`
	Content   string      `json:"content"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at,omitempty"`
	User      *User       `json:"user,omitempty"`
	CreatedBy *User       `json:"created_by,omitempty"`
	ParentID  *int        `json:"parent_id,omitempty"`
	TopicID   int         `json:"topic_id,omitempty"`
	LikeCount int         `json:"like_count,omitempty"`
	IsLiked   bool        `json:"is_liked,omitempty"`
	Uploads   []*Upload   `json:"uploads,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}
