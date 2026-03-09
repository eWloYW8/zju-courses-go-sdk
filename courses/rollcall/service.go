package rollcall

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

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

// GetRollcallStudentRollcallsDetail returns rollcall details with student rollcalls.
func (s *Service) GetRollcallStudentRollcallsDetail(ctx context.Context, rollcallID int, action string) (*Rollcall, error) {
	u := fmt.Sprintf("/api/rollcall/%d/student_rollcalls", rollcallID)
	if action != "" {
		u = addQueryParams(u, map[string]string{"action": action})
	}
	result := new(Rollcall)
	_, err := s.client.Get(ctx, u, result)
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

// GetRollcallStatusResult returns the resolved rollcall status result for a course.
func (s *Service) GetRollcallStatusResult(ctx context.Context, courseID int) (RollcallStatusResultResponse, error) {
	u := fmt.Sprintf("/api/courses/rollcall_status/%d/result", courseID)
	result := make(RollcallStatusResultResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListTimetableRollcalls returns timetable rollcalls.
func (s *Service) ListTimetableRollcalls(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/timetable_rollcalls", &result)
	return result, err
}

// ListTimetableRollcallsWithParams returns timetable rollcalls filtered like the frontend.
func (s *Service) ListTimetableRollcallsWithParams(ctx context.Context, params *ListTimetableRollcallsParams) ([]*TimetableRollcallCourse, error) {
	u := "/api/timetable_rollcalls"
	if params != nil {
		values := url.Values{}
		for _, courseID := range params.CourseIDs {
			values.Add("course_ids", strconv.Itoa(courseID))
		}
		if params.RollcallDate != "" {
			values.Set("rollcall_date", params.RollcallDate)
		}
		u = addQueryValues(u, values)
	}
	var result []*TimetableRollcallCourse
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListCourseRollcalls returns rollcall records for a course.
func (s *Service) ListCourseRollcalls(ctx context.Context, courseID int) (*CourseRollcallsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/rollcalls", courseID)
	result := new(CourseRollcallsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// AnswerSelfRegistrationRollcall answers a self-registration rollcall.
func (s *Service) AnswerSelfRegistrationRollcall(ctx context.Context, rollcallID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d/answer_self_registration_rollcall", rollcallID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, struct{}{}, &result)
	return result, err
}

// AnswerNumberRollcall answers a number-code rollcall.
func (s *Service) AnswerNumberRollcall(ctx context.Context, rollcallID int, body *AnswerNumberRollcallRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d/answer_number_rollcall", rollcallID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// ListCourseStudentRollcalls returns a student's rollcall records for a course.
func (s *Service) ListCourseStudentRollcalls(ctx context.Context, courseID, studentID int, opts *model.ListOptions) (*CourseStudentRollcallsResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/course/%d/student/%d/rollcalls", courseID, studentID), opts)
	result := new(CourseStudentRollcallsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateCourseStudentRollcalls updates a student's rollcall statuses.
func (s *Service) UpdateCourseStudentRollcalls(ctx context.Context, courseID, studentID int, body *UpdateCourseStudentRollcallsRequest) (UpdateCourseStudentRollcallsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/student/%d/rollcalls", courseID, studentID)
	result := make(UpdateCourseStudentRollcallsResponse)
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// GetLeaveRecord returns leave users for a course at a timestamp.
func (s *Service) GetLeaveRecord(ctx context.Context, courseID int, timestamp string) (*LeaveRecordResponse, error) {
	u := addQueryParams(fmt.Sprintf("/api/course/%d/leave-record", courseID), map[string]string{
		"timestamp": timestamp,
	})
	result := new(LeaveRecordResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}

func addQueryValues(urlStr string, values url.Values) string {
	if len(values) == 0 {
		return urlStr
	}
	sep := "?"
	if containsQuery(urlStr) {
		sep = "&"
	}
	return urlStr + sep + values.Encode()
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}

func containsQuery(urlStr string) bool {
	for i := 0; i < len(urlStr); i++ {
		if urlStr[i] == '?' {
			return true
		}
	}
	return false
}
