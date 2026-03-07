package activities

import "github.com/eWloYW8/zju-courses-go-sdk/model"

type UploadReferencesResponse struct {
	References []*model.UploadReference `json:"references"`
}

type CommentsResponse struct {
	Comments []*Comment `json:"comments"`
	model.Pagination
}

type CommentPageCountResponse struct {
	PageStats []any `json:"page_stats"`
}

type RecommendSubmissionsResponse struct {
	Submissions []*model.Submission `json:"submissions"`
}

type ActivityResourcesResponse struct {
	Resources []map[string]any `json:"resources"`
}

type DeleteCheckResponse map[string]any

type HaveDependentsResponse map[string]any

type ExamActivityReadLogResponse map[string]any

type ClassinJoinURLResponse map[string]any

type ClassinWebcastURLResponse map[string]any
