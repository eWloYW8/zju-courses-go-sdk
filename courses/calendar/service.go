package calendar

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles calendar-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Calendar Events ---

// ListCalendarEvents returns calendar events.
func (s *Service) ListCalendarEvents(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	query := map[string]string{"no-intercept": "true"}
	for k, v := range params {
		query[k] = v
	}
	u := addQueryParams("/api/calendar-events", query)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCalendarEvent returns a specific calendar event.
func (s *Service) GetCalendarEvent(ctx context.Context, eventID int) (*CalendarEvent, error) {
	u := fmt.Sprintf("/api/calendar-events/%d", eventID)
	result := new(CalendarEvent)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateCalendarEvent creates a new calendar event.
func (s *Service) CreateCalendarEvent(ctx context.Context, body interface{}) (*CalendarEvent, error) {
	result := new(CalendarEvent)
	_, err := s.client.Post(ctx, "/api/calendar-events", body, result)
	return result, err
}

// UpdateCalendarEvent updates a calendar event.
func (s *Service) UpdateCalendarEvent(ctx context.Context, eventID int, body interface{}) (*CalendarEvent, error) {
	u := fmt.Sprintf("/api/calendar-events/%d", eventID)
	result := new(CalendarEvent)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteCalendarEvent deletes a calendar event.
func (s *Service) DeleteCalendarEvent(ctx context.Context, eventID int) error {
	u := fmt.Sprintf("/api/calendar-events/%d", eventID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ListCalendarUserEvents returns calendar events for a user.
func (s *Service) ListCalendarUserEvents(ctx context.Context, userID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/calendar-events/users/%d", userID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCalendarAlerts returns calendar alerts.
func (s *Service) GetCalendarAlerts(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/calendar-alerts?no-intercept=true", &result)
	return result, err
}

// --- Timetables ---

// ListTimetables returns calendar timetables.
func (s *Service) ListTimetables(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	query := map[string]string{"no-intercept": "true"}
	for k, v := range params {
		query[k] = v
	}
	u := addQueryParams("/api/calendar-timetables", query)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetTimetable returns a specific timetable.
func (s *Service) GetTimetable(ctx context.Context, timetableID int) (*CalendarTimetable, error) {
	u := fmt.Sprintf("/api/calendar-timetables/%d", timetableID)
	result := new(CalendarTimetable)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateTimetable creates a new timetable entry.
func (s *Service) CreateTimetable(ctx context.Context, body interface{}) (*TimetableMutationResponse, error) {
	result := new(TimetableMutationResponse)
	_, err := s.client.Post(ctx, "/api/calendar-timetables", body, result)
	return result, err
}

// UpdateTimetable updates a timetable entry.
func (s *Service) UpdateTimetable(ctx context.Context, timetableID int, body interface{}) (*TimetableMutationResponse, error) {
	u := fmt.Sprintf("/api/calendar-timetables/%d", timetableID)
	result := new(TimetableMutationResponse)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteTimetable deletes a timetable entry.
func (s *Service) DeleteTimetable(ctx context.Context, timetableID int, body *DeleteTimetableRequest) (*TimetableMutationResponse, error) {
	u := fmt.Sprintf("/api/calendar-timetables/%d/delete", timetableID)
	result := new(TimetableMutationResponse)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// ListCalendarDepartments returns departments used by the calendar/team-teaching popup.
func (s *Service) ListCalendarDepartments(ctx context.Context) (*CalendarDepartmentsResponse, error) {
	result := new(CalendarDepartmentsResponse)
	_, err := s.client.Get(ctx, "/api/departments?no-intercept=true", &result)
	return result, err
}

// ListCalendarUserCourses returns course options for a user's calendar-event binding popup.
func (s *Service) ListCalendarUserCourses(ctx context.Context, userID int) (*CalendarUserCoursesResponse, error) {
	u := fmt.Sprintf("/api/calendar-events/users/%d/courses?no-intercept=true", userID)
	result := new(CalendarUserCoursesResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Teaching Calendar ---

// GetTeachingCalendars returns teaching calendars.
func (s *Service) GetTeachingCalendars(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/teaching-calendars", &result)
	return result, err
}

// GetTeachingCalendar returns a specific teaching calendar.
func (s *Service) GetTeachingCalendar(ctx context.Context, calendarID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/teaching-calendar/%d", calendarID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Calendar Meeting ---

// CreateCalendarMeeting creates a calendar meeting.
func (s *Service) CreateCalendarMeeting(ctx context.Context, body interface{}) (*CalendarMeetingMutationResponse, error) {
	result := new(CalendarMeetingMutationResponse)
	_, err := s.client.Post(ctx, "/api/calendar-meeting", body, &result)
	return result, err
}

// GetCalendarMeeting returns the calendar-meeting payload used by the frontend popup.
func (s *Service) GetCalendarMeeting(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/calendar-meeting", &result)
	return result, err
}

// UpdateCalendarMeeting updates a calendar meeting.
func (s *Service) UpdateCalendarMeeting(ctx context.Context, meetingID int, body interface{}) (*CalendarMeetingMutationResponse, error) {
	u := fmt.Sprintf("/api/calendar-meeting/%d", meetingID)
	result := new(CalendarMeetingMutationResponse)
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// DeleteCalendarMeeting deletes a calendar meeting.
func (s *Service) DeleteCalendarMeeting(ctx context.Context, meetingID int) error {
	u := fmt.Sprintf("/api/calendar-meeting/%d", meetingID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ListManagementCalendarMeetings returns management calendar meetings.
func (s *Service) ListManagementCalendarMeetings(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	return s.ListManagementCalendarMeetingsWithBody(ctx, opts, nil)
}

// ListManagementCalendarMeetingsWithBody returns management calendar meetings using the frontend POST API.
func (s *Service) ListManagementCalendarMeetingsWithBody(ctx context.Context, opts *model.ListOptions, body interface{}) (json.RawMessage, error) {
	u := addListOptions("/api/management/calendar-meeting", opts)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ExportCalendarMeetings exports calendar meetings to Excel.
func (s *Service) ExportCalendarMeetings(ctx context.Context) ([]byte, error) {
	return s.ExportCalendarMeetingsWithBody(ctx, nil)
}

// ExportCalendarMeetingsWithBody exports calendar meetings to Excel using the frontend POST API.
func (s *Service) ExportCalendarMeetingsWithBody(ctx context.Context, body interface{}) ([]byte, error) {
	req, err := s.client.NewRequest(ctx, "POST", "/api/management/calendar-meeting/excel", body)
	if err != nil {
		return nil, err
	}
	_, data, err := s.client.DoBytes(req)
	return data, err
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}
