package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// HomeworkService handles homework-related API operations.
type HomeworkService struct {
	client *Client
}

// --- Response Types ---

type SubmissionListResponse struct {
	List []*model.Submission `json:"list"`
}

// --- Student Homework ---

// GetHomeworkScore returns the homework score for a student.
func (s *HomeworkService) GetHomeworkScore(ctx context.Context, activityID, studentID int) (*model.HomeworkScore, error) {
	u := fmt.Sprintf("/api/activities/%d/students/%d/homework-score", activityID, studentID)
	result := new(model.HomeworkScore)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListSubmissions returns submissions for a student on a homework activity.
func (s *HomeworkService) ListSubmissions(ctx context.Context, activityID, studentID int) (*SubmissionListResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/students/%d/submission_list", activityID, studentID)
	result := new(SubmissionListResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetMakeUpRecord returns the make-up record for a student.
func (s *HomeworkService) GetMakeUpRecord(ctx context.Context, activityID, studentID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/homework/%d/students/%d/make-up-record", activityID, studentID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetResubmitRecord returns the resubmit record for a student.
func (s *HomeworkService) GetResubmitRecord(ctx context.Context, activityID, studentID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/homework/%d/students/%d/resubmit-record", activityID, studentID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Homework Submission ---

// SubmitHomework submits homework for a student.
func (s *HomeworkService) SubmitHomework(ctx context.Context, activityID int, body interface{}) (*model.Submission, error) {
	u := fmt.Sprintf("/api/homework/submission/%d", activityID)
	result := new(model.Submission)
	_, err := s.client.post(ctx, u, body, result)
	return result, err
}

// UpdateSubmission updates a homework submission.
func (s *HomeworkService) UpdateSubmission(ctx context.Context, submissionID int, body interface{}) (*model.Submission, error) {
	u := fmt.Sprintf("/api/submissions/%d", submissionID)
	result := new(model.Submission)
	_, err := s.client.put(ctx, u, body, result)
	return result, err
}

// RecommendSubmission marks submissions as recommended.
func (s *HomeworkService) RecommendSubmission(ctx context.Context, submissionIDs []int) error {
	_, err := s.client.put(ctx, "/api/submission/recommend", map[string][]int{"submission_ids": submissionIDs}, nil)
	return err
}

// --- Homework Management (Instructor) ---

// CreateHomework creates a new homework activity.
func (s *HomeworkService) CreateHomework(ctx context.Context, courseID int, homework interface{}) (*model.Activity, error) {
	u := fmt.Sprintf("/api/homeworks/%d", courseID)
	result := new(model.Activity)
	_, err := s.client.post(ctx, u, homework, result)
	return result, err
}

// UpdateHomework updates a homework activity.
func (s *HomeworkService) UpdateHomework(ctx context.Context, homeworkID int, homework interface{}) (*model.Activity, error) {
	u := fmt.Sprintf("/api/homework/%d", homeworkID)
	result := new(model.Activity)
	_, err := s.client.put(ctx, u, homework, result)
	return result, err
}

// DeleteHomework deletes a homework activity.
func (s *HomeworkService) DeleteHomework(ctx context.Context, homeworkID int) error {
	u := fmt.Sprintf("/api/homework/%d", homeworkID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// ScoreSubmission scores a submission (instructor).
func (s *HomeworkService) ScoreSubmission(ctx context.Context, submissionID int, body interface{}) error {
	u := fmt.Sprintf("/api/submissions/%d", submissionID)
	_, err := s.client.patch(ctx, u, body, nil)
	return err
}

// ListInterScoreSubmissions returns inter-review score submissions.
func (s *HomeworkService) ListInterScoreSubmissions(ctx context.Context, activityID int) ([]interface{}, error) {
	u := fmt.Sprintf("/api/inter-score-submissions/%d", activityID)
	var result []interface{}
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListInterScores returns inter-review scores.
func (s *HomeworkService) ListInterScores(ctx context.Context, activityID int) ([]interface{}, error) {
	u := fmt.Sprintf("/api/inter-scores/%d", activityID)
	var result []interface{}
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetInProgressHomeworks returns in-progress homeworks.
func (s *HomeworkService) GetInProgressHomeworks(ctx context.Context) ([]interface{}, error) {
	var result []interface{}
	_, err := s.client.get(ctx, "/api/in-progress-homeworks?no-intercept=true", &result)
	return result, err
}

// GetHomeworksSubmissionStatus returns homework submission statuses across courses.
func (s *HomeworkService) GetHomeworksSubmissionStatus(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/homeworks-submission-status?no-intercept=true&course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// DownloadHomeworkZip checks the status of a homework zip download.
func (s *HomeworkService) DownloadHomeworkZip(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/zip-status/homework-zip/%d", activityID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// MakeUpExam creates a make-up exam record.
func (s *HomeworkService) MakeUpExam(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/make-up-exams", body, nil)
	return err
}

// MakeupExam creates a makeup exam (alternate endpoint).
func (s *HomeworkService) MakeupExam(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/makeup-exams", body, nil)
	return err
}
