package uploads

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type Upload = model.Upload

type UploadReference = model.UploadReference

// UploadPDFInfo represents the preview PDF metadata used by the AI quiz flow.
type UploadPDFInfo struct {
	NumPages int `json:"num_pages,omitempty"`
}

// SCORMCMIData represents SCORM CMI learner data.
type SCORMCMIData struct {
	CMI            any `json:"cmi"`
	SuspendData    any `json:"suspend_data,omitempty"`
	TotalPages     int `json:"total_pages,omitempty"`
	VisitedPages   int `json:"visited_pages,omitempty"`
	CompletedPages int `json:"completed_pages,omitempty"`
}
