package notifications

import (
	"github.com/eWloYW8/zju-courses-go-sdk/courses/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

// Notification represents a notification item.
type Notification struct {
	ID        string               `json:"id"`
	Type      string               `json:"type"`
	Top       bool                 `json:"top,omitempty"`
	Timestamp int64                `json:"timestamp,omitempty"`
	Unread    bool                 `json:"unread,omitempty"`
	CreatedAt string               `json:"created_at,omitempty"`
	Payload   *NotificationPayload `json:"payload,omitempty"`
}

// NotificationPayload represents the notification content.
type NotificationPayload struct {
	ActivityID      int     `json:"activity_id,omitempty"`
	ActivityTitle   string  `json:"activity_title,omitempty"`
	ActivityType    string  `json:"activity_type,omitempty"`
	CourseID        int     `json:"course_id,omitempty"`
	CourseName      string  `json:"course_name,omitempty"`
	CreatedAt       string  `json:"created_at,omitempty"`
	EndTime         string  `json:"end_time,omitempty"`
	GroupID         *int    `json:"group_id,omitempty"`
	GroupName       *string `json:"group_name,omitempty"`
	OrgID           int     `json:"org_id,omitempty"`
	StartTime       string  `json:"start_time,omitempty"`
	SubmissionID    *int    `json:"submission_id,omitempty"`
	TopicContent    string  `json:"topic_content,omitempty"`
	TopicID         *int    `json:"topic_id,omitempty"`
	TopicTitle      string  `json:"topic_title,omitempty"`
	UserID          int     `json:"user_id,omitempty"`
	UserName        string  `json:"user_name,omitempty"`
	BulletinID      *int    `json:"bulletin_id,omitempty"`
	BulletinTitle   string  `json:"bulletin_title,omitempty"`
	CalendarEventID *int    `json:"calendar_event_id,omitempty"`
	ExamID          *int    `json:"exam_id,omitempty"`
	MeetingID       *int    `json:"meeting_id,omitempty"`
}

// TodoItem represents a to-do item.
type TodoItem struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Type          string  `json:"type"`
	CourseID      int     `json:"course_id"`
	CourseType    int     `json:"course_type,omitempty"`
	CourseName    string  `json:"course_name"`
	CourseCode    string  `json:"course_code,omitempty"`
	StartTime     string  `json:"start_time,omitempty"`
	EndTime       string  `json:"end_time,omitempty"`
	SurveyURL     string  `json:"survey_url,omitempty"`
	IsLocked      bool    `json:"is_locked,omitempty"`
	IsStudent     bool    `json:"is_student,omitempty"`
	NotScoredNum  int     `json:"not_scored_num,omitempty"`
	SubmitRate    float64 `json:"submit_rate,omitempty"`
	Prerequisites []any   `json:"prerequisites,omitempty"`
}

// Bulletin represents a bulletin/announcement.
type Bulletin struct {
	ID        int             `json:"id"`
	Title     string          `json:"title,omitempty"`
	Content   string          `json:"content,omitempty"`
	CourseID  int             `json:"course_id,omitempty"`
	Course    *model.Course   `json:"course,omitempty"`
	CreatedAt string          `json:"created_at,omitempty"`
	UpdatedAt string          `json:"updated_at,omitempty"`
	CreatedBy *model.User     `json:"created_by,omitempty"`
	IsRead    bool            `json:"is_read,omitempty"`
	Uploads   []*model.Upload `json:"uploads,omitempty"`
}

// Announcement represents an announcement.
type Announcement struct {
	ID   int `json:"id,omitempty"`
	Data any `json:"data,omitempty"`
}

// OrgBulletin represents an organization-level bulletin.
type OrgBulletin struct {
	ID               int    `json:"id"`
	Title            string `json:"title,omitempty"`
	Content          string `json:"content,omitempty"`
	ClassificationID int    `json:"classification_id,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	IsRead           bool   `json:"is_read,omitempty"`
}

type AlertMessage struct {
	ID   int            `json:"id,omitempty"`
	Data map[string]any `json:"data,omitempty"`
}

type OrgBulletinClassification struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

type LatestActivity = activities.Activity
