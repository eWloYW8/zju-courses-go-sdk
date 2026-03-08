package interactions

type InteractionSubjectOption struct {
	ID        int    `json:"id"`
	Content   string `json:"content,omitempty"`
	Type      string `json:"type,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	IsAnswer  bool   `json:"is_answer,omitempty"`
	IsCorrect bool   `json:"is_correct,omitempty"`
	Checked   bool   `json:"checked,omitempty"`
}

type InteractionSubject struct {
	ID                int                         `json:"id"`
	Type              string                      `json:"type,omitempty"`
	Description       string                      `json:"description,omitempty"`
	Sort              int                         `json:"sort,omitempty"`
	LastUpdatedAt     string                      `json:"last_updated_at,omitempty"`
	AnswerExplanation string                      `json:"answer_explanation,omitempty"`
	WrongExplanation  string                      `json:"wrong_explanation,omitempty"`
	Options           []*InteractionSubjectOption `json:"options,omitempty"`
}

type InteractionSubmissionSubject struct {
	ID                int    `json:"id"`
	LastUpdatedAt     string `json:"last_updated_at,omitempty"`
	AnswerExplanation string `json:"answer_explanation,omitempty"`
	WrongExplanation  string `json:"wrong_explanation,omitempty"`
}

type InteractionSubmissionSubjectsData struct {
	Subjects []*InteractionSubmissionSubject `json:"subjects,omitempty"`
}

type InteractionCorrectAnswer struct {
	AnswerOptionIDs []int `json:"answer_option_ids,omitempty"`
}

type InteractionSubmission struct {
	ID                 int                                  `json:"id"`
	IsValid            bool                                 `json:"is_valid,omitempty"`
	SubjectsData       *InteractionSubmissionSubjectsData   `json:"subjects_data,omitempty"`
	ScoreData          map[string]float64                   `json:"score_data,omitempty"`
	CorrectAnswersData map[string]*InteractionCorrectAnswer `json:"correct_answers_data,omitempty"`
	SubmissionData     map[string][]int                     `json:"submission_data,omitempty"`
}
