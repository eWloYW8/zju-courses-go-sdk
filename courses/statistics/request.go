package statistics

type UserActionRequest map[string]any

type FootprintRequest map[string]any

type TrackingRequest map[string]any

type ExportONOStatStudentsParams struct {
	StartDate  string
	EndDate    string
	Conditions any
}

type ExportONOStatStudentsRequest struct {
	CourseIDs string `json:"course_ids,omitempty"`
}
