package model

// HomeworkActivity represents a homework in the scores list.
type HomeworkActivity struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
}

// HomeworkSubmissionStatus represents submission status for a homework.
type HomeworkSubmissionStatus struct {
	ID                        int      `json:"id"`
	HomeworkType              string   `json:"homework_type,omitempty"`
	Score                     *float64 `json:"score,omitempty"`
	Status                    string   `json:"status,omitempty"`
	StatusCode                string   `json:"status_code,omitempty"`
	IsAnnounceScoreTimePassed bool     `json:"is_announce_score_time_passed,omitempty"`
}

// HomeworkScore represents homework score info.
type HomeworkScore struct {
	ActivityID        int      `json:"activity_id"`
	StudentID         int      `json:"student_id,omitempty"`
	Score             *float64 `json:"score,omitempty"`
	FinalScore        *float64 `json:"final_score,omitempty"`
	InstructorComment *string  `json:"instructor_comment,omitempty"`
	InterScore        *float64 `json:"inter_score,omitempty"`
	IntraScore        *float64 `json:"intra_score,omitempty"`
}

// Submission represents a homework submission.
type Submission struct {
	ID           int         `json:"id"`
	ActivityID   int         `json:"activity_id,omitempty"`
	StudentID    int         `json:"student_id,omitempty"`
	GroupID      *int        `json:"group_id,omitempty"`
	Content      string      `json:"content,omitempty"`
	Score        *float64    `json:"score,omitempty"`
	FinalScore   *float64    `json:"final_score,omitempty"`
	IsLate       bool        `json:"is_late,omitempty"`
	IsRecommend  bool        `json:"is_recommend,omitempty"`
	SubmittedAt  string      `json:"submitted_at,omitempty"`
	CreatedAt    string      `json:"created_at,omitempty"`
	UpdatedAt    string      `json:"updated_at,omitempty"`
	Uploads      []*Upload   `json:"uploads,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	User         *User       `json:"user,omitempty"`
}

// ForumScore represents a student's forum score.
type ForumScore struct {
	ActivityID int      `json:"activity_id"`
	Score      *float64 `json:"score,omitempty"`
	StudentID  int      `json:"student_id,omitempty"`
}

// HomeworkScorePercentage represents score weight configuration for homework.
type HomeworkScorePercentage struct {
	HomeworkScore    float64 `json:"homework_score,omitempty"`
	InterReviewScore float64 `json:"inter_review_score,omitempty"`
	IntraReviewScore float64 `json:"intra_review_score,omitempty"`
}

// AutoComputeRules represents forum auto-compute scoring rules.
type AutoComputeRules struct {
	EachReplyScore  string `json:"each_reply_score,omitempty"`
	MaxReplyScore   string `json:"max_reply_score,omitempty"`
	EachPraiseScore string `json:"each_praise_score,omitempty"`
	MaxPraiseScore  string `json:"max_praise_score,omitempty"`
}
