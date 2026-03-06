package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"
)

// StatisticsService handles statistics-related API operations.
type StatisticsService struct {
	client *Client
}

// --- User Visits ---

// GetUserVisits returns user visit statistics.
// jwt parameter is required for statistics API.
func (s *StatisticsService) GetUserVisits(ctx context.Context, jwt string, params map[string]string) (json.RawMessage, error) {
	u := fmt.Sprintf("/statistics/api/user-visits?jwt=%s", jwt)
	u = addQueryParams(u, params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Learning Activity ---

// GetLearningActivity returns learning activity statistics.
func (s *StatisticsService) GetLearningActivity(ctx context.Context, jwt string, params map[string]string) (json.RawMessage, error) {
	u := fmt.Sprintf("/statistics/api/learning-activity?jwt=%s", jwt)
	u = addQueryParams(u, params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- User Actions ---

// PostUserActions reports user actions.
func (s *StatisticsService) PostUserActions(ctx context.Context, jwt string, body interface{}) error {
	u := fmt.Sprintf("/api/user-actions?jwt=%s", jwt)
	_, err := s.client.post(ctx, u, body, nil)
	return err
}

// --- Course Statistics ---

// GetCourseStatistics returns statistics for a course.
func (s *StatisticsService) GetCourseStatistics(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/courses/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetActivitiesForCourses returns activity statistics across courses.
func (s *StatisticsService) GetActivitiesForCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/stat/activities-for-courses", &result)
	return result, err
}

// ExportCourseStats exports course statistics.
func (s *StatisticsService) ExportCourseStats(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/courses/export/to/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ExportClassHours exports class hours statistics.
func (s *StatisticsService) ExportClassHours(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/stat/courses/class-hours/export", &result)
	return result, err
}

// ExportHomeworkCorrect exports homework correction statistics.
func (s *StatisticsService) ExportHomeworkCorrect(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/stat/courses/homework-correct/export", &result)
	return result, err
}

// ExportRollcall exports rollcall statistics.
func (s *StatisticsService) ExportRollcall(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/stat/courses/rollcall/export", &result)
	return result, err
}

// ExportRollcallByClass exports rollcall statistics by class.
func (s *StatisticsService) ExportRollcallByClass(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/stat/courses/rollcall/export-by-class", &result)
	return result, err
}

// ExportAttendance exports attendance statistics.
func (s *StatisticsService) ExportAttendance(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/attendance/export/to/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetVTRSesStats returns VTRS statistics.
func (s *StatisticsService) GetVTRSesStats(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/vtrses/%d", vtrsID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetStatistic returns general statistics.
func (s *StatisticsService) GetStatistic(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/statistic", &result)
	return result, err
}

// --- Footprint ---

// RecordFootprint records a user footprint.
func (s *StatisticsService) RecordFootprint(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/footprint/tread", body, nil)
	return err
}

// --- User Completeness ---

// GetUserCompleteness returns user completeness.
func (s *StatisticsService) GetUserCompleteness(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/user-completeness", params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}
