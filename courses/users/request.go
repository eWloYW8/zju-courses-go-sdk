package users

type UpdateProfileRequest map[string]any

type SearchUserParams struct {
	Keywords           string
	ExcludeStudentRole bool
	OrgID              *int
	DepartmentID       *int
}

type DepartmentUpdateRequest struct {
	DepartmentID int `json:"department_id"`
}

type UserLinkRequest struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type CheckCoursesGraduateRequest struct {
	CourseIDs []int `json:"course_ids,omitempty"`
	UserID    *int  `json:"user_id,omitempty"`
	OrgID     *int  `json:"org_id,omitempty"`
}

type NoteRequest struct {
	Content    string `json:"content,omitempty"`
	TargetType string `json:"target_type,omitempty"`
	Anchor     *int   `json:"anchor,omitempty"`
	CourseID   *int   `json:"course_id,omitempty"`
	ActivityID *int   `json:"activity_id,omitempty"`
	TargetID   *int   `json:"target_id,omitempty"`
}

type LanguageAssessmentRequest struct {
	LanguageLevel string `json:"languageLevel,omitempty"`
	LanguageScore string `json:"languageScore,omitempty"`
}

type PersonasRequest struct {
	Hobbies            []string                   `json:"hobbies,omitempty"`
	Education          string                     `json:"education,omitempty"`
	Industry           string                     `json:"industry,omitempty"`
	Career             string                     `json:"career,omitempty"`
	LanguageAssessment *LanguageAssessmentRequest `json:"languageAssessment,omitempty"`
}

type UserChatRequest struct {
	SessionID string `json:"session_id,omitempty"`
	Message   string `json:"message,omitempty"`
}

type RecalculateCourseCertificationScoresRequest struct {
	CourseIDs []int `json:"course_ids,omitempty"`
}

type ThirdPartResourcesParams struct {
	Page       int
	PageSize   int
	Conditions string
}

type OtherVideoResourcesParams struct {
	Page     int
	PageSize int
}

type AcademicLearningResourcesParams struct {
	Page       int
	PageSize   int
	CourseCode string
	Keyword    string
	Types      []string
}
