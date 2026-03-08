package users

import "github.com/eWloYW8/zju-courses-go-sdk/model"

// UserProfile represents a user's profile information.
type UserProfile struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Email          string            `json:"email"`
	Nickname       *string           `json:"nickname"`
	UserNo         string            `json:"user_no"`
	AvatarSmallURL string            `json:"avatar_small_url"`
	AvatarBigURL   string            `json:"avatar_big_url"`
	Department     *model.Department `json:"department"`
	Org            *model.Org        `json:"org"`
	Language       string            `json:"language,omitempty"`
}

type AcademicYear = model.AcademicYear

type Semester = model.Semester

type Department = model.Department

type Class = model.Class

type Grade = model.Grade

type Upload = model.Upload

type Course = model.Course

// SignIn represents a sign-in session for a course.
type SignIn struct {
	ID        int    `json:"id"`
	CourseID  int    `json:"course_id,omitempty"`
	Status    string `json:"status,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Type      string `json:"type,omitempty"`
}

type UserLink struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type UserSearchResult struct {
	ID     int    `json:"id"`
	Name   string `json:"name,omitempty"`
	UserNo string `json:"user_no,omitempty"`
}

type FirstTimeLoginResponse map[string]any

type CourseGraduateCheckResponse struct {
	CompletedCourseIDs   []int `json:"completed_course_id,omitempty"`
	UncompletedCourseIDs []int `json:"uncompleted_course_id,omitempty"`
}

type ExpiredPasswordResponse struct {
	IsPasswordExpired bool `json:"is_password_expired"`
}

type PreTaskResponse struct {
	HasPreTask bool   `json:"has_pre_task"`
	URL        string `json:"url,omitempty"`
}

type PersonasResponse struct {
	HasPersonas bool `json:"has_personas"`
}

type AssociationCodeResponse struct {
	AssociationCode string `json:"association_code,omitempty"`
}

type ChatMessage struct {
	Message   string         `json:"message,omitempty"`
	IsReply   bool           `json:"is_reply,omitempty"`
	Type      string         `json:"type,omitempty"`
	Data      map[string]any `json:"data,omitempty"`
	SessionID string         `json:"session_id,omitempty"`
}

type Note struct {
	ID         int    `json:"id"`
	Content    string `json:"content,omitempty"`
	TargetType string `json:"target_type,omitempty"`
	Anchor     *int   `json:"anchor,omitempty"`
	CourseID   *int   `json:"course_id,omitempty"`
	ActivityID *int   `json:"activity_id,omitempty"`
	TargetID   *int   `json:"target_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	Order      string `json:"order,omitempty"`
}

type StorageUsedResponse struct {
	StorageUsed     int64 `json:"storage_used"`
	StorageAssigned int64 `json:"storage_assigned"`
}

type Capture struct {
	ID        int     `json:"id"`
	Code      string  `json:"code,omitempty"`
	Title     string  `json:"title,omitempty"`
	Name      string  `json:"name,omitempty"`
	CourseID  *int    `json:"course_id,omitempty"`
	UploadID  *int    `json:"upload_id,omitempty"`
	URL       string  `json:"url,omitempty"`
	Cover     *string `json:"cover,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
}

type CoursesInfoStatusResponse struct {
	OutlineStatus string `json:"outline_status,omitempty"`
	ScoreStatus   string `json:"score_status,omitempty"`
}

type OrgSummaryResponse map[string]any
