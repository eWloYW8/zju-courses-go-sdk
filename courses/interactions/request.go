package interactions

type CreateModuleInteractionRequest struct {
	Title             string `json:"title,omitempty"`
	Type              string `json:"type,omitempty"`
	StudentNum        *int   `json:"student_num,omitempty"`
	Countdown         *int   `json:"countdown,omitempty"`
	RecommendStudents []int  `json:"recommend_students,omitempty"`
}

type InteractionSubmissionAnswer struct {
	SubjectID       int   `json:"subject_id"`
	AnswerOptionIDs []int `json:"answer_option_ids,omitempty"`
}

type UpdateInteractionSubmissionRequest struct {
	Subjects []*InteractionSubmissionAnswer `json:"subjects,omitempty"`
}

type UpdateInteractionScoreRequest map[string]any

type UpdateInteractionVoteRequest map[string]any
