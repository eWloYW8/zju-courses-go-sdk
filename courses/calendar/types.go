package calendar

// CalendarEvent represents a calendar event.
type CalendarEvent struct {
	ID          int    `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	CourseID    *int   `json:"course_id,omitempty"`
	CourseName  string `json:"course_name,omitempty"`
	Type        string `json:"type,omitempty"`
	AllDay      bool   `json:"all_day,omitempty"`
	Location    string `json:"location,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	Data        any    `json:"data,omitempty"`
}

// CalendarTimetable represents a timetable entry.
type CalendarTimetable struct {
	ID        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Location  string `json:"location,omitempty"`
	Data      any    `json:"data,omitempty"`
}
