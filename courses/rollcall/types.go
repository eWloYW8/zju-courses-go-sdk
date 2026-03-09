package rollcall

import "encoding/json"

type Rollcall struct {
	ID               int                `json:"id"`
	Title            string             `json:"title,omitempty"`
	Type             string             `json:"type,omitempty"`
	Source           string             `json:"source,omitempty"`
	Status           string             `json:"status,omitempty"`
	ModuleID         int                `json:"module_id,omitempty"`
	CourseID         int                `json:"course_id,omitempty"`
	RollcallTime     *string            `json:"rollcall_time,omitempty"`
	StartTime        *string            `json:"start_time,omitempty"`
	EndTime          *string            `json:"end_time,omitempty"`
	UniqueKey        string             `json:"unique_key,omitempty"`
	Comment          *string            `json:"comment,omitempty"`
	IsRadar          bool               `json:"is_radar,omitempty"`
	IsNumber         bool               `json:"is_number,omitempty"`
	ExternalAPIKeyID int                `json:"external_api_key_id,omitempty"`
	Children         []*Rollcall        `json:"children,omitempty"`
	StudentRollcalls []*RollcallStudent `json:"student_rollcalls,omitempty"`
}

type RollcallStudent struct {
	StudentID         int     `json:"student_id"`
	Status            string  `json:"status,omitempty"`
	StatusDetail      *string `json:"status_detail,omitempty"`
	RollcallStatus    string  `json:"rollcall_status,omitempty"`
	TemperatureStatus *string `json:"temperature_status,omitempty"`
}

type TimetableRollcallCourse struct {
	ID           int       `json:"id"`
	IsInstructor bool      `json:"-"`
	Rollcall     *Rollcall `json:"rollcall,omitempty"`
}

func (c *TimetableRollcallCourse) UnmarshalJSON(data []byte) error {
	type alias TimetableRollcallCourse
	aux := struct {
		*alias
		IsInstructorSnake *bool `json:"is_instructor,omitempty"`
		IsInstructorCamel *bool `json:"isInstructor,omitempty"`
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch {
	case aux.IsInstructorSnake != nil:
		c.IsInstructor = *aux.IsInstructorSnake
	case aux.IsInstructorCamel != nil:
		c.IsInstructor = *aux.IsInstructorCamel
	}
	return nil
}

type CourseStudentRollcall struct {
	StudentRollcallID   int     `json:"student_rollcall_id"`
	StudentStatus       string  `json:"student_status,omitempty"`
	Status              string  `json:"status,omitempty"`
	StudentStatusDetail *string `json:"student_status_detail,omitempty"`
	Title               string  `json:"title,omitempty"`
	Scored              bool    `json:"scored,omitempty"`
}
