package rollcall

type CreateRollcallRequest struct {
	Status   string `json:"status,omitempty"`
	CourseID int    `json:"course_id,omitempty"`
	ModuleID int    `json:"module_id,omitempty"`
	Title    string `json:"title,omitempty"`
	IsNumber bool   `json:"is_number,omitempty"`
	Type     string `json:"type,omitempty"`
}

type AnswerNumberRollcallRequest struct {
	NumberCode string `json:"numberCode,omitempty"`
}

type ListTimetableRollcallsParams struct {
	CourseIDs    []int
	RollcallDate string
}

type UpdateCourseStudentRollcallsRequest struct {
	StudentRollcalls []*UpdateCourseStudentRollcall `json:"student_rollcalls"`
}

type UpdateCourseStudentRollcall struct {
	StudentRollcallID int    `json:"student_rollcall_id"`
	StudentStatus     string `json:"student_status"`
}
