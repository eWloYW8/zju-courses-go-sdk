package rollcall

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type CourseRollcallsResponse struct {
	Rollcalls []*Rollcall `json:"rollcalls"`
}

type CreateRollcallResponse struct {
	ID      int     `json:"id"`
	Message *string `json:"message,omitempty"`
}

type RollcallStatusResultResponse map[string]any

type CourseStudentRollcallsResponse struct {
	Rollcalls []*CourseStudentRollcall `json:"rollcalls"`
}

type UpdateCourseStudentRollcallsResponse map[string]any

type LeaveRecordResponse struct {
	UserNos []string `json:"user_nos"`
}

type TimetableRollcallsResponse struct {
	Courses []*TimetableRollcallCourse `json:"courses,omitempty"`
	model.Pagination
}
