package forum

import (
	"github.com/eWloYW8/zju-courses-go-sdk/courses/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

type Topic struct {
	ID                int                      `json:"id"`
	Title             string                   `json:"title,omitempty"`
	Content           string                   `json:"content,omitempty"`
	Summary           string                   `json:"summary,omitempty"`
	TopicType         string                   `json:"topic_type,omitempty"`
	CreatedAt         string                   `json:"created_at,omitempty"`
	UpdatedAt         *string                  `json:"updated_at,omitempty"`
	CreatedBy         *activities.ActivityUser  `json:"created_by,omitempty"`
	GroupID           *int                     `json:"group_id,omitempty"`
	TopicCategoryID   int                      `json:"topic_category_id,omitempty"`
	InCommonCategory  bool                     `json:"in_common_category,omitempty"`
	CurrentUserLiked  bool                     `json:"current_user_liked,omitempty"`
	CurrentUserRead   bool                     `json:"current_user_read,omitempty"`
	IsHotTopic        bool                     `json:"is_hot_topic,omitempty"`
	IsTeacherTop      bool                     `json:"is_teacher_top,omitempty"`
	TopicToppedSort   *int                     `json:"topic_topped_sort,omitempty"`
	LikeCount         int                      `json:"like_count,omitempty"`
	ReplyCount        int                      `json:"reply_count,omitempty"`
	ReadReplies       []int                    `json:"read_replies,omitempty"`
	UnreadReplyCount  int                      `json:"unread_reply_count,omitempty"`
	HasMatchedReplies bool                     `json:"has_matched_replies,omitempty"`
	UserVisitsNumber  int                      `json:"user_visits_number,omitempty"`
	Enrollments       []*TopicEnrollment       `json:"enrollments,omitempty"`
	Uploads           []*model.Upload          `json:"uploads,omitempty"`
}

type TopicEnrollment struct {
	Aliases []*string `json:"aliases,omitempty"`
	Roles   []string  `json:"roles,omitempty"`
}

type ForumCategoryResponse struct {
	ID         int                  `json:"id"`
	ActivityID int                  `json:"activity_id"`
	Result     *ForumCategoryResult `json:"result,omitempty"`
}

type ForumCategoryResult struct {
	Page     int      `json:"page"`
	PageSize int      `json:"page_size"`
	Pages    int      `json:"pages"`
	Total    int      `json:"total"`
	Start    int      `json:"start"`
	End      int      `json:"end"`
	Topics   []*Topic `json:"topics,omitempty"`
}

type Reply struct {
	ID        int                      `json:"id"`
	Content   string                   `json:"content"`
	CreatedAt string                   `json:"created_at"`
	UpdatedAt *string                  `json:"updated_at,omitempty"`
	User      *activities.ActivityUser  `json:"user,omitempty"`
	CreatedBy *activities.ActivityUser  `json:"created_by,omitempty"`
	ParentID  *int                     `json:"parent_id,omitempty"`
	TopicID   int                      `json:"topic_id,omitempty"`
	LikeCount int                      `json:"like_count,omitempty"`
	IsLiked   bool                     `json:"is_liked,omitempty"`
	Uploads   []*model.Upload          `json:"uploads,omitempty"`
	Data      map[string]any           `json:"data,omitempty"`
}

type ForumScore struct {
	ActivityID int      `json:"activity_id"`
	Score      *float64 `json:"score,omitempty"`
	StudentID  int      `json:"student_id,omitempty"`
}

type ForumScoresResponse map[string]any

type CourseForumScoresResponse map[string]any

type CategoryRepliedResponse map[string]any

type TopicCategory struct {
	ID           int    `json:"id"`
	Title        string `json:"title,omitempty"`
	ReferrerType string `json:"referrer_type,omitempty"`
}
