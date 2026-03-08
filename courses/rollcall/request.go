package rollcall

type CreateRollcallRequest struct {
	Status   string `json:"status,omitempty"`
	CourseID int    `json:"course_id,omitempty"`
	ModuleID int    `json:"module_id,omitempty"`
	Title    string `json:"title,omitempty"`
	IsNumber bool   `json:"is_number,omitempty"`
	Type     string `json:"type,omitempty"`
}
