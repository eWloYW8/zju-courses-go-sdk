package rollcall

type Rollcall struct {
	ID           int     `json:"id"`
	Title        string  `json:"title,omitempty"`
	Type         string  `json:"type,omitempty"`
	Source       string  `json:"source,omitempty"`
	Status       string  `json:"status,omitempty"`
	ModuleID     int     `json:"module_id,omitempty"`
	CourseID     int     `json:"course_id,omitempty"`
	RollcallTime *string `json:"rollcall_time,omitempty"`
	StartTime    *string `json:"start_time,omitempty"`
	UniqueKey    string  `json:"unique_key,omitempty"`
}
