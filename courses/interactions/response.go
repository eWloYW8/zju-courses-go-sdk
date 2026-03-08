package interactions

type Interaction struct {
	ID         int            `json:"id"`
	Title      string         `json:"title,omitempty"`
	Type       string         `json:"type,omitempty"`
	Status     string         `json:"status,omitempty"`
	CourseID   int            `json:"course_id,omitempty"`
	ActivityID int            `json:"activity_id,omitempty"`
	CreatedAt  string         `json:"created_at,omitempty"`
	Data       map[string]any `json:"data,omitempty"`
}

type CourseInteractionsResponse struct {
	Interactions []*Interaction `json:"interactions"`
}

type InteractionSubjectsResponse struct {
	Subjects []*InteractionSubject `json:"subjects"`
}

type InteractionSubmissionsResponse struct {
	Submissions []*InteractionSubmission `json:"submissions"`
}
