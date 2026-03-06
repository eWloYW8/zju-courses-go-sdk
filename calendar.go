package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// CalendarService handles calendar-related API operations.
type CalendarService struct {
	client *Client
}

// --- Calendar Events ---

// ListCalendarEvents returns calendar events.
func (s *CalendarService) ListCalendarEvents(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/calendar-events", params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetCalendarEvent returns a specific calendar event.
func (s *CalendarService) GetCalendarEvent(ctx context.Context, eventID int) (*model.CalendarEvent, error) {
	u := fmt.Sprintf("/api/calendar-events/%d", eventID)
	result := new(model.CalendarEvent)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// CreateCalendarEvent creates a new calendar event.
func (s *CalendarService) CreateCalendarEvent(ctx context.Context, body interface{}) (*model.CalendarEvent, error) {
	result := new(model.CalendarEvent)
	_, err := s.client.post(ctx, "/api/calendar-events", body, result)
	return result, err
}

// UpdateCalendarEvent updates a calendar event.
func (s *CalendarService) UpdateCalendarEvent(ctx context.Context, eventID int, body interface{}) (*model.CalendarEvent, error) {
	u := fmt.Sprintf("/api/calendar-events/%d", eventID)
	result := new(model.CalendarEvent)
	_, err := s.client.put(ctx, u, body, result)
	return result, err
}

// DeleteCalendarEvent deletes a calendar event.
func (s *CalendarService) DeleteCalendarEvent(ctx context.Context, eventID int) error {
	u := fmt.Sprintf("/api/calendar-events/%d", eventID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// ListCalendarUserEvents returns calendar events for a user.
func (s *CalendarService) ListCalendarUserEvents(ctx context.Context, userID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/calendar-events/users/%d", userID), params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetCalendarAlerts returns calendar alerts.
func (s *CalendarService) GetCalendarAlerts(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/calendar-alerts?no-intercept=true", &result)
	return result, err
}

// --- Timetables ---

// ListTimetables returns calendar timetables.
func (s *CalendarService) ListTimetables(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/calendar-timetables", params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetTimetable returns a specific timetable.
func (s *CalendarService) GetTimetable(ctx context.Context, timetableID int) (*model.CalendarTimetable, error) {
	u := fmt.Sprintf("/api/calendar-timetables/%d", timetableID)
	result := new(model.CalendarTimetable)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// CreateTimetable creates a new timetable entry.
func (s *CalendarService) CreateTimetable(ctx context.Context, body interface{}) (*model.CalendarTimetable, error) {
	result := new(model.CalendarTimetable)
	_, err := s.client.post(ctx, "/api/calendar-timetables", body, result)
	return result, err
}

// UpdateTimetable updates a timetable entry.
func (s *CalendarService) UpdateTimetable(ctx context.Context, timetableID int, body interface{}) (*model.CalendarTimetable, error) {
	u := fmt.Sprintf("/api/calendar-timetables/%d", timetableID)
	result := new(model.CalendarTimetable)
	_, err := s.client.put(ctx, u, body, result)
	return result, err
}

// DeleteTimetable deletes a timetable entry.
func (s *CalendarService) DeleteTimetable(ctx context.Context, timetableID int) error {
	u := fmt.Sprintf("/api/calendar-timetables/%d", timetableID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// --- Teaching Calendar ---

// GetTeachingCalendars returns teaching calendars.
func (s *CalendarService) GetTeachingCalendars(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/teaching-calendars", &result)
	return result, err
}

// GetTeachingCalendar returns a specific teaching calendar.
func (s *CalendarService) GetTeachingCalendar(ctx context.Context, calendarID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/teaching-calendar/%d", calendarID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Calendar Meeting ---

// CreateCalendarMeeting creates a calendar meeting.
func (s *CalendarService) CreateCalendarMeeting(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/calendar-meeting", body, &result)
	return result, err
}

// GetCalendarMeeting returns a specific calendar meeting.
func (s *CalendarService) GetCalendarMeeting(ctx context.Context, meetingID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/calendar-meeting/%d", meetingID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// UpdateCalendarMeeting updates a calendar meeting.
func (s *CalendarService) UpdateCalendarMeeting(ctx context.Context, meetingID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/calendar-meeting/%d", meetingID)
	var result json.RawMessage
	_, err := s.client.put(ctx, u, body, &result)
	return result, err
}

// DeleteCalendarMeeting deletes a calendar meeting.
func (s *CalendarService) DeleteCalendarMeeting(ctx context.Context, meetingID int) error {
	u := fmt.Sprintf("/api/calendar-meeting/%d", meetingID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// ListManagementCalendarMeetings returns management calendar meetings.
func (s *CalendarService) ListManagementCalendarMeetings(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/management/calendar-meeting", opts)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ExportCalendarMeetings exports calendar meetings to Excel.
func (s *CalendarService) ExportCalendarMeetings(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/management/calendar-meeting/excel", &result)
	return result, err
}
