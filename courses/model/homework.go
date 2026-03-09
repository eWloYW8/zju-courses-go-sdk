package model

// Submission represents a homework submission.
type Submission struct {
	ID          int         `json:"id"`
	ActivityID  int         `json:"activity_id,omitempty"`
	StudentID   int         `json:"student_id,omitempty"`
	GroupID     *int        `json:"group_id,omitempty"`
	Content     string      `json:"content,omitempty"`
	Score       *float64    `json:"score,omitempty"`
	FinalScore  *float64    `json:"final_score,omitempty"`
	IsLate      bool        `json:"is_late,omitempty"`
	IsRecommend bool        `json:"is_recommend,omitempty"`
	SubmittedAt string      `json:"submitted_at,omitempty"`
	CreatedAt   string      `json:"created_at,omitempty"`
	UpdatedAt   string      `json:"updated_at,omitempty"`
	Uploads     []*Upload   `json:"uploads,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	User        *User       `json:"user,omitempty"`
	CreatedBy   *User       `json:"created_by,omitempty"`
}
