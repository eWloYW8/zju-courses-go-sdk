package statistics

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles statistics-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- User Visits ---

// GetUserVisits returns user visit statistics.
// jwt parameter is required for statistics API.
func (s *Service) GetUserVisits(ctx context.Context, jwt string, params StatisticsQueryParams) (UserVisitsResponse, error) {
	u := fmt.Sprintf("/statistics/api/user-visits?jwt=%s", jwt)
	u = addQueryParams(u, params)
	var result UserVisitsResponse
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Learning Activity ---

// GetLearningActivity returns learning activity statistics.
func (s *Service) GetLearningActivity(ctx context.Context, jwt string, params StatisticsQueryParams) (LearningActivityResponse, error) {
	u := fmt.Sprintf("/statistics/api/learning-activity?jwt=%s", jwt)
	u = addQueryParams(u, params)
	var result LearningActivityResponse
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- User Actions ---

// PostUserActions reports user actions.
func (s *Service) PostUserActions(ctx context.Context, jwt string, body UserActionRequest) error {
	u := fmt.Sprintf("/api/user-actions?jwt=%s", jwt)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// --- Course Statistics ---

// GetCourseStatistics returns statistics for a course.
func (s *Service) GetCourseStatistics(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/courses/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCourseTrialTeachingStats returns trial teaching statistics for a course.
func (s *Service) GetCourseTrialTeachingStats(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/courses/%d/trial-teaching", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetActivitiesForCourses returns activity statistics across courses.
func (s *Service) GetActivitiesForCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/stat/activities-for-courses", &result)
	return result, err
}

// ExportCourseStats exports course statistics.
func (s *Service) ExportCourseStats(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/courses/export/to/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ExportClassHours exports class hours statistics.
func (s *Service) ExportClassHours(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/stat/courses/class-hours/export", &result)
	return result, err
}

// ExportHomeworkCorrect exports homework correction statistics.
func (s *Service) ExportHomeworkCorrect(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/stat/courses/homework-correct/export", &result)
	return result, err
}

// ExportRollcall exports rollcall statistics.
func (s *Service) ExportRollcall(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/stat/courses/rollcall/export", &result)
	return result, err
}

// ExportRollcallByClass exports rollcall statistics by class.
func (s *Service) ExportRollcallByClass(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/stat/courses/rollcall/export-by-class", &result)
	return result, err
}

// ExportAttendance exports attendance statistics.
func (s *Service) ExportAttendance(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/attendance/export/to/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVTRSesStats returns VTRS statistics.
func (s *Service) GetVTRSesStats(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/vtrses/%d", vtrsID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVTRSOverview returns VTRS overview statistics.
func (s *Service) GetVTRSOverview(ctx context.Context, vtrsID int, dateRange string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/vtrses/%d/overview", vtrsID)
	if dateRange != "" {
		u = addQueryParams(u, map[string]string{"dateRange": dateRange})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVTRSResourcesRank returns VTRS resource ranking statistics.
func (s *Service) GetVTRSResourcesRank(ctx context.Context, vtrsID int, dateRange string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/vtrses/%d/resources/rank", vtrsID)
	if dateRange != "" {
		u = addQueryParams(u, map[string]string{"dateRange": dateRange})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVTRSTrialTeaching returns VTRS trial teaching statistics.
func (s *Service) GetVTRSTrialTeaching(ctx context.Context, vtrsID int, dateRange string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/vtrses/%d/trial-teaching", vtrsID)
	if dateRange != "" {
		u = addQueryParams(u, map[string]string{"dateRange": dateRange})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVTRSDepartment returns VTRS department statistics.
func (s *Service) GetVTRSDepartment(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/vtrses/%d/department", vtrsID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVTRSTeamActivation returns VTRS team activation statistics.
func (s *Service) GetVTRSTeamActivation(ctx context.Context, vtrsID int, opts map[string]string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/vtrses/%d/team-activation", vtrsID)
	u = addQueryParams(u, opts)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVTRSResources returns VTRS resource statistics.
func (s *Service) GetVTRSResources(ctx context.Context, vtrsID int, dateRange string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/stat/vtrses/%d/resources", vtrsID)
	if dateRange != "" {
		u = addQueryParams(u, map[string]string{"dateRange": dateRange})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetStatistic returns general statistics.
func (s *Service) GetStatistic(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/statistic", &result)
	return result, err
}

// --- Footprint ---

// RecordFootprint records a user footprint.
func (s *Service) RecordFootprint(ctx context.Context, body FootprintRequest) error {
	_, err := s.client.Post(ctx, "/api/footprint/tread", body, nil)
	return err
}

// --- User Completeness ---

// GetUserCompleteness returns user completeness.
func (s *Service) GetUserCompleteness(ctx context.Context, params UserCompletenessQueryParams) (UserCompletenessResponse, error) {
	u := addQueryParams("/api/user-completeness", params)
	var result UserCompletenessResponse
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}
