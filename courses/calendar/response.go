package calendar

import (
	"encoding/json"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

type CalendarEventsResponse = json.RawMessage
type CalendarAlertsResponse = json.RawMessage
type TimetablesResponse = json.RawMessage
type TeachingCalendarsResponse = json.RawMessage
type CalendarMeetingResponse = json.RawMessage

type CalendarDepartmentsResponse struct {
	Departments []*model.Department `json:"departments"`
}

type CalendarUserCoursesResponse struct {
	Courses []*model.Course `json:"courses"`
}

type TimetableMutationResponse struct {
	Message   string             `json:"message,omitempty"`
	Timetable *CalendarTimetable `json:"timetable,omitempty"`
}

type CalendarMeetingMutationResponse struct {
	Message string `json:"message,omitempty"`
}

type TeamTeachingOpenedOrgsResponse struct {
	Orgs []*model.OrgDetail `json:"orgs"`
}
