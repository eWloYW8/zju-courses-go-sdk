package homework

import (
	"github.com/eWloYW8/zju-courses-go-sdk/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

type HomeworkScore struct {
	ActivityID        int      `json:"activity_id"`
	StudentID         int      `json:"student_id,omitempty"`
	Score             *float64 `json:"score,omitempty"`
	FinalScore        *float64 `json:"final_score,omitempty"`
	InstructorComment *string  `json:"instructor_comment,omitempty"`
	InterScore        *float64 `json:"inter_score,omitempty"`
	IntraScore        *float64 `json:"intra_score,omitempty"`
}

type Submission struct {
	ID          int                     `json:"id"`
	ActivityID  int                     `json:"activity_id,omitempty"`
	StudentID   int                     `json:"student_id,omitempty"`
	GroupID     *int                    `json:"group_id,omitempty"`
	Content     string                  `json:"content,omitempty"`
	Score       *float64                `json:"score,omitempty"`
	FinalScore  *float64                `json:"final_score,omitempty"`
	IsLate      bool                    `json:"is_late,omitempty"`
	IsRecommend bool                    `json:"is_recommend,omitempty"`
	SubmittedAt string                  `json:"submitted_at,omitempty"`
	CreatedAt   string                  `json:"created_at,omitempty"`
	UpdatedAt   string                  `json:"updated_at,omitempty"`
	Uploads     []*model.Upload         `json:"uploads,omitempty"`
	Data        map[string]any          `json:"data,omitempty"`
	User        *activities.ActivityUser `json:"user,omitempty"`
}

type MakeUpRecordResponse map[string]any

type ResubmitRecordResponse map[string]any

type MarkedAttachmentsResponse map[string]any

type MarkedAttachmentResponse map[string]any

type InterScoreSubmission map[string]any

type InterScore map[string]any

type InProgressHomework map[string]any

type HomeworksSubmissionStatusResponse map[string]any

type HomeworkZipStatusResponse map[string]any
