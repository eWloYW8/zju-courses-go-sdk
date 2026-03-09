package uploads

type CreateReferenceRequest struct {
	UploadID   *int   `json:"upload_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Link       string `json:"link,omitempty"`
	ParentID   *int   `json:"parent_id,omitempty"`
	ParentType string `json:"parent_type,omitempty"`
}

type UpdateReferenceUploadRequest struct {
	UploadID int `json:"upload_id"`
}

type ShareToCoursesRequest struct {
	UploadIDs []int `json:"upload_ids,omitempty"`
	CourseIDs []int `json:"course_ids,omitempty"`
}

type DuplicateReportDownloadRequest struct {
	ReportType string `json:"report_type,omitempty"`
	DetectKey  string `json:"detect_key,omitempty"`
	Provider   string `json:"provider,omitempty"`
}
