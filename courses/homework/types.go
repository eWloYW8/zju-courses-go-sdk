package homework

import (
	"github.com/eWloYW8/zju-courses-go-sdk/courses/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

type HomeworkScore struct {
	ActivityID        int      `json:"activity_id"`
	StudentID         int      `json:"student_id,omitempty"`
	GroupID           *int     `json:"group_id,omitempty"`
	Score             *float64 `json:"score,omitempty"`
	FinalScore        *float64 `json:"final_score,omitempty"`
	InstructorComment *string  `json:"instructor_comment,omitempty"`
	InterScore        *float64 `json:"inter_score,omitempty"`
	IntraScore        *float64 `json:"intra_score,omitempty"`
}

type Submission struct {
	ID          int                      `json:"id"`
	ActivityID  int                      `json:"activity_id,omitempty"`
	StudentID   int                      `json:"student_id,omitempty"`
	GroupID     *int                     `json:"group_id,omitempty"`
	Content     string                   `json:"content,omitempty"`
	Score       *float64                 `json:"score,omitempty"`
	FinalScore  *float64                 `json:"final_score,omitempty"`
	IsLate      bool                     `json:"is_late,omitempty"`
	IsRecommend bool                     `json:"is_recommend,omitempty"`
	SubmittedAt string                   `json:"submitted_at,omitempty"`
	CreatedAt   string                   `json:"created_at,omitempty"`
	UpdatedAt   string                   `json:"updated_at,omitempty"`
	Uploads     []*model.Upload          `json:"uploads,omitempty"`
	Data        map[string]any           `json:"data,omitempty"`
	User        *activities.ActivityUser `json:"user,omitempty"`
	CreatedBy   *activities.ActivityUser `json:"created_by,omitempty"`
}

type MakeUpRecord map[string]any

type MakeUpRecordResponse = MakeUpRecord

type ResubmitRecord map[string]any

type ResubmitRecordResponse = ResubmitRecord

type MarkedAttachmentInfo struct {
	OriginUpload     *model.UploadReference `json:"origin_upload,omitempty"`
	MarkedAttachment *model.Upload          `json:"marked_attachment,omitempty"`
}

type MarkedAttachmentsResponse struct {
	MarkedAttachmentInfos []*MarkedAttachmentInfo `json:"marked_attachment_infos,omitempty"`
}

type MarkedAttachmentResponse struct {
	MarkedAttachment *model.Upload `json:"marked_attachment,omitempty"`
}

type InterScoreSubmission map[string]any

type InterScore map[string]any

type HomeworkLog map[string]any

type DuplicateDetectRateItem map[string]any

type InProgressHomework struct {
	ID         int     `json:"id"`
	CourseID   int     `json:"course_id,omitempty"`
	CourseType int     `json:"course_type,omitempty"`
	Title      string  `json:"title,omitempty"`
	Type       string  `json:"type,omitempty"`
	StartTime  *string `json:"start_time,omitempty"`
	EndTime    *string `json:"end_time,omitempty"`
	IsLocked   bool    `json:"is_locked,omitempty"`
}

type HomeworkSubmissionStatus struct {
	ID                        int      `json:"id"`
	Score                     *float64 `json:"score,omitempty"`
	Status                    string   `json:"status,omitempty"`
	StatusCode                string   `json:"status_code,omitempty"`
	IsAnnounceScoreTimePassed bool     `json:"is_announce_score_time_passed,omitempty"`
}

type HomeworksSubmissionStatusResponse struct {
	HomeworkStatuses []*HomeworkSubmissionStatus `json:"homework_statuses"`
}

type HomeworkZipStatusResponse map[string]any
