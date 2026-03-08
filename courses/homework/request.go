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
