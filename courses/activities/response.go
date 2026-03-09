package activities

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type UploadReferencesResponse struct {
	References []*model.UploadReference `json:"references"`
}

type CourseActivitiesResponse struct {
	Activities []*Activity `json:"activities"`
}

type CommentsResponse struct {
	Comments []*Comment `json:"comments"`
	model.Pagination
}

type CommentPageCountResponse struct {
	PageStats []CommentPageStat `json:"page_stats"`
}

type ScoreRecordsResponse struct {
	Records []*ScoreRecord `json:"records"`
	model.Pagination
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

type RecommendSubmissionsResponse struct {
	Submissions []*model.Submission `json:"submissions"`
}

type ActivityResourcesResponse struct {
	Resources []map[string]any `json:"resources"`
}

type DeleteCheckResponse struct {
	SafeDelete bool `json:"safe_delete"`
}

type HaveDependentsResponse struct {
	HasDependents            bool            `json:"has_dependents"`
	UnavailablePrerequisites []*Prerequisite `json:"unavailable_prerequisites,omitempty"`
}

type ActivityCompletionCriteriaResponse struct {
	CompletionCriteria     []*CompletionCriterion `json:"completion_criteria"`
	HasCompletionCriterion bool                   `json:"has_completion_criterion"`
}

type ActivityPrerequisitesResponse struct {
	Prerequisites []*Prerequisite `json:"prerequisites"`
}

type CommentPageStat struct {
	Page     int `json:"page"`
	Forum    int `json:"forum,omitempty"`
	Question int `json:"question,omitempty"`
}

type ExamActivityReadLogResponse map[string]any

type ConvertedInteractionResponse struct {
	ID int `json:"id"`
}
