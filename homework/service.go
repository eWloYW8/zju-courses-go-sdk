package homework

import (
	"context"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles homework-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Student Homework ---

// GetHomeworkScore returns the homework score for a student.
func (s *Service) GetHomeworkScore(ctx context.Context, activityID, studentID int) (*HomeworkScore, error) {
	u := fmt.Sprintf("/api/activities/%d/students/%d/homework-score", activityID, studentID)
	result := new(HomeworkScore)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListSubmissions returns submissions for a student on a homework activity.
func (s *Service) ListSubmissions(ctx context.Context, activityID, studentID int) (*SubmissionListResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/students/%d/submission_list", activityID, studentID)
	result := new(SubmissionListResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetMakeUpRecord returns the make-up record for a student.
func (s *Service) GetMakeUpRecord(ctx context.Context, activityID, studentID int) (MakeUpRecordResponse, error) {
	u := fmt.Sprintf("/api/homework/%d/students/%d/make-up-record", activityID, studentID)
	result := make(MakeUpRecordResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetResubmitRecord returns the resubmit record for a student.
func (s *Service) GetResubmitRecord(ctx context.Context, activityID, studentID int) (ResubmitRecordResponse, error) {
	u := fmt.Sprintf("/api/homework/%d/students/%d/resubmit-record", activityID, studentID)
	result := make(ResubmitRecordResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Homework Submission ---

// SubmitHomework submits homework for a student.
func (s *Service) SubmitHomework(ctx context.Context, activityID int, body *SubmitHomeworkRequest) (*Submission, error) {
	u := fmt.Sprintf("/api/homework/submission/%d", activityID)
	result := new(Submission)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateSubmission updates a homework submission.
func (s *Service) UpdateSubmission(ctx context.Context, submissionID int, body *UpdateSubmissionRequest) (*Submission, error) {
	u := fmt.Sprintf("/api/submissions/%d", submissionID)
	result := new(Submission)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// ListMarkedAttachments returns marked attachments for a submission.
func (s *Service) ListMarkedAttachments(ctx context.Context, submissionID int) (MarkedAttachmentsResponse, error) {
	u := fmt.Sprintf("/api/submissions/%d/marked_attachments", submissionID)
	result := make(MarkedAttachmentsResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetMarkedAttachment returns a marked attachment for a submission.
func (s *Service) GetMarkedAttachment(ctx context.Context, submissionID, attachmentID int) (MarkedAttachmentResponse, error) {
	u := fmt.Sprintf("/api/submissions/%d/marked_attachments/%d", submissionID, attachmentID)
	result := make(MarkedAttachmentResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// RecommendSubmission marks submissions as recommended.
func (s *Service) RecommendSubmission(ctx context.Context, submissionIDs []int) error {
	_, err := s.client.Put(ctx, "/api/submission/recommend", &RecommendSubmissionRequest{SubmissionIDs: submissionIDs}, nil)
	return err
}

// --- Homework Management (Instructor) ---

// CreateHomework creates a new homework activity.
func (s *Service) CreateHomework(ctx context.Context, courseID int, homework *CreateHomeworkRequest) (*activities.Activity, error) {
	u := fmt.Sprintf("/api/homeworks/%d", courseID)
	result := new(activities.Activity)
	_, err := s.client.Post(ctx, u, homework, result)
	return result, err
}

// UpdateHomework updates a homework activity.
func (s *Service) UpdateHomework(ctx context.Context, homeworkID int, homework *UpdateHomeworkRequest) (*activities.Activity, error) {
	u := fmt.Sprintf("/api/homework/%d", homeworkID)
	result := new(activities.Activity)
	_, err := s.client.Put(ctx, u, homework, result)
	return result, err
}

// DeleteHomework deletes a homework activity.
func (s *Service) DeleteHomework(ctx context.Context, homeworkID int) error {
	u := fmt.Sprintf("/api/homework/%d", homeworkID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ScoreSubmission scores a submission (instructor).
func (s *Service) ScoreSubmission(ctx context.Context, submissionID int, body *ScoreSubmissionRequest) error {
	u := fmt.Sprintf("/api/submissions/%d", submissionID)
	_, err := s.client.Patch(ctx, u, body, nil)
	return err
}

// ListInterScoreSubmissions returns inter-review score submissions.
func (s *Service) ListInterScoreSubmissions(ctx context.Context, activityID int) ([]InterScoreSubmission, error) {
	u := fmt.Sprintf("/api/inter-score-submissions/%d", activityID)
	var result []InterScoreSubmission
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListInterScores returns inter-review scores.
func (s *Service) ListInterScores(ctx context.Context, activityID int) ([]InterScore, error) {
	u := fmt.Sprintf("/api/inter-scores/%d", activityID)
	var result []InterScore
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetInProgressHomeworks returns in-progress homeworks.
func (s *Service) GetInProgressHomeworks(ctx context.Context) ([]InProgressHomework, error) {
	var result []InProgressHomework
	_, err := s.client.Get(ctx, "/api/in-progress-homeworks?no-intercept=true", &result)
	return result, err
}

// GetHomeworksSubmissionStatus returns homework submission statuses across courses.
func (s *Service) GetHomeworksSubmissionStatus(ctx context.Context, courseID int) (HomeworksSubmissionStatusResponse, error) {
	u := fmt.Sprintf("/api/courses/homeworks-submission-status?no-intercept=true&course_id=%d", courseID)
	result := make(HomeworksSubmissionStatusResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// DownloadHomeworkZip checks the status of a homework zip download.
func (s *Service) DownloadHomeworkZip(ctx context.Context, activityID int) (HomeworkZipStatusResponse, error) {
	u := fmt.Sprintf("/api/zip-status/homework-zip/%d", activityID)
	result := make(HomeworkZipStatusResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// MakeUpExam creates a make-up exam record.
func (s *Service) MakeUpExam(ctx context.Context, body MakeUpExamRequest) error {
	_, err := s.client.Post(ctx, "/api/make-up-exams", body, nil)
	return err
}

// MakeupExam creates a makeup exam (alternate endpoint).
func (s *Service) MakeupExam(ctx context.Context, body MakeupExamRequest) error {
	_, err := s.client.Post(ctx, "/api/makeup-exams", body, nil)
	return err
}
