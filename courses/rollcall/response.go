package rollcall

type CourseRollcallsResponse struct {
	Rollcalls []*Rollcall `json:"rollcalls"`
}

type CreateRollcallResponse struct {
	ID      int     `json:"id"`
	Message *string `json:"message,omitempty"`
}
