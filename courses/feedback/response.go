package feedback

type FeedbacksResponse struct {
	Feedbacks []*Feedback `json:"feedbacks,omitempty"`
}
