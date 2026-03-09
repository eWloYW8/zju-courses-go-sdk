package homework

import "github.com/eWloYW8/zju-courses-go-sdk/courses/activities"

type SubmitHomeworkRequest struct {
	Content   string         `json:"content,omitempty"`
	UploadIDs []int          `json:"upload_ids,omitempty"`
	Data      map[string]any `json:"data,omitempty"`
	GroupID   *int           `json:"group_id,omitempty"`
}

type UpdateSubmissionRequest struct {
	Content   string         `json:"content,omitempty"`
	UploadIDs []int          `json:"upload_ids,omitempty"`
	Data      map[string]any `json:"data,omitempty"`
}

type CreateHomeworkRequest = activities.Activity

type UpdateHomeworkRequest = activities.Activity

type ScoreSubmissionRequest struct {
	Score             *float64 `json:"score,omitempty"`
	FinalScore        *float64 `json:"final_score,omitempty"`
	InstructorComment *string  `json:"instructor_comment,omitempty"`
	InterScore        *float64 `json:"inter_score,omitempty"`
	IntraScore        *float64 `json:"intra_score,omitempty"`
}

type RecommendSubmissionRequest struct {
	SubmissionIDs []int `json:"submission_ids"`
}

type MarkedSubmittedRequest struct {
	SubmissionIDs     []int    `json:"submission_ids,omitempty"`
	StudentOrGroupIDs []int    `json:"student_or_group_ids,omitempty"`
	MarkedSubmitted   int      `json:"marked_submitted"`
	SubmittedStatus   []string `json:"submittedStatus,omitempty"`
}

type ListSubmissionRecordsParams struct {
	NeedUploadsSize *bool `json:"-"`
	UserIDs         []int `json:"-"`
}

type ListDuplicateLibUploadsParams struct {
	Page       int
	PageSize   int
	Conditions any
}

type AddDuplicateLibUploadsRequest struct {
	UploadIDs []int `json:"upload_ids,omitempty"`
}

type DuplicateDetectReportDownloadRequest struct {
	ReportType string `json:"report_type,omitempty"`
	DetectKey  string `json:"detect_key,omitempty"`
	Provider   string `json:"provider,omitempty"`
}

type MarkHomeworkSubmissionToRedoRequest struct {
	SubmissionID int `json:"submission_id,omitempty"`
	StudentID    int `json:"student_id,omitempty"`
	GroupID      int `json:"group_id,omitempty"`
}

type HomeworkAIGenerateRequest struct {
	LearningGoals         string   `json:"learning_goals,omitempty"`
	BloomCognitiveDomains []string `json:"bloom_cognitive_domains,omitempty"`
	Locale                string   `json:"locale,omitempty"`
	Assignment            string   `json:"assignment,omitempty"`
	Suggestion            string   `json:"suggestion,omitempty"`
}

type SubmissionAnalysisRequest struct {
	Content string `json:"content,omitempty"`
}
