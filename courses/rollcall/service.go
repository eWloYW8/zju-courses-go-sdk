package rollcall

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles rollcall-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// GetRollcall returns a rollcall.
func (s *Service) GetRollcall(ctx context.Context, rollcallID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d", rollcallID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateRollcall creates a module rollcall activity.
func (s *Service) CreateRollcall(ctx context.Context, courseID int, body *CreateRollcallRequest) (*CreateRollcallResponse, error) {
	u := fmt.Sprintf("/api/module/%d/rollcall", courseID)
	result := new(CreateRollcallResponse)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// StartRollcall starts a rollcall session.
func (s *Service) StartRollcall(ctx context.Context, rollcallID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d/start-rollcall", rollcallID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// StopTimetableRollcall stops a timetable rollcall session.
func (s *Service) StopTimetableRollcall(ctx context.Context, rollcallID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d/stop_time_table_rollcall", rollcallID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// GetRollcallStudentRollcalls returns student rollcalls for a rollcall.
func (s *Service) GetRollcallStudentRollcalls(ctx context.Context, rollcallID int, action string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d/student_rollcalls", rollcallID)
	if action != "" {
		u = addQueryParams(u, map[string]string{"action": action})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetRollcallStudentsPagination returns paginated student rollcalls for a rollcall.
func (s *Service) GetRollcallStudentsPagination(ctx context.Context, rollcallID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/rollcall/%d/pagination_students_rollcalls", rollcallID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetMergedRollcall returns a merged rollcall.
func (s *Service) GetMergedRollcall(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/rollcall/merged-rollcall", &result)
	return result, err
}

// GetMergedRollcallStudentRollcalls returns student rollcalls from merged rollcall.
func (s *Service) GetMergedRollcallStudentRollcalls(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/rollcall/merged-rollcall/student-rollcalls", &result)
	return result, err
}

// GetRollcallStatus returns rollcall status for a course.
func (s *Service) GetRollcallStatus(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/rollcall_status/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListTimetableRollcalls returns timetable rollcalls.
func (s *Service) ListTimetableRollcalls(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/timetable_rollcalls", &result)
	return result, err
}

// ListCourseRollcalls returns rollcall records for a course.
func (s *Service) ListCourseRollcalls(ctx context.Context, courseID int) (*CourseRollcallsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/rollcalls", courseID)
	result := new(CourseRollcallsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}
