package courses

import (
	"context"
	"fmt"
)

// --- Scores ---

// ListHomeworkScores returns homework score entries for a course.
func (s *Service) ListHomeworkScores(ctx context.Context, courseID int) (*HomeworkScoresResponse, error) {
	u := fmt.Sprintf("/api/course/%d/homework-scores", courseID)
	result := new(HomeworkScoresResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetHomeworkSubmissionStatus returns homework submission statuses for a course.
func (s *Service) GetHomeworkSubmissionStatus(ctx context.Context, courseID int) (*HomeworkSubmissionStatusResponse, error) {
	u := fmt.Sprintf("/api/course/%d/homework/submission-status?no-intercept=true", courseID)
	result := new(HomeworkSubmissionStatusResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListExamScores returns exam scores for a course.
func (s *Service) ListExamScores(ctx context.Context, courseID int) (*ExamScoresResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/exam-scores?no-intercept=true", courseID)
	result := new(ExamScoresResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}
