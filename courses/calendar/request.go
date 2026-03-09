package calendar

type CreateCalendarEventRequest map[string]any
type UpdateCalendarEventRequest map[string]any
type CreateTimetableRequest map[string]any
type UpdateTimetableRequest map[string]any
type DeleteTimetableRequest struct {
	TargetTime  string `json:"target_time,omitempty"`
	UpdateScope string `json:"update_scope,omitempty"`
}
type CreateCalendarMeetingRequest map[string]any
type UpdateCalendarMeetingRequest map[string]any
type ListManagementCalendarMeetingsRequest map[string]any
type ExportCalendarMeetingsRequest map[string]any
