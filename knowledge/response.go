package knowledge

import "github.com/eWloYW8/zju-courses-go-sdk/model"

type KnowledgeNodesResponse struct {
	Items []*KnowledgeNode `json:"items"`
}

type KnowledgeNodeCapturesResponse struct {
	Items []*KnowledgeCapture `json:"items"`
	model.Pagination
}

type KnowledgeNodeResourcesResponse struct {
	Items []*KnowledgeNodeRecommendedResourceReference `json:"items"`
	model.Pagination
}

type KnowledgeNodeReferenceResourcesResponse struct {
	Resources []*model.SharedResource `json:"resources"`
	model.Pagination
}

type KnowledgeNodeStudentDimensionResponse struct {
	Items []*KnowledgeNodeStudentDimensionItem `json:"items"`
	model.Pagination
}

type KnowledgeNodeStudentDetailsResponse struct {
	Items []*KnowledgeNodeStudentDetail `json:"items"`
	model.Pagination
}

type KnowledgeNodeResourceDetailsResponse struct {
	Items []*KnowledgeNodeResourceDetail `json:"items"`
	model.Pagination
}

type KnowledgeNodeActivityDetailsResponse struct {
	Items []*KnowledgeNodeActivityDetail `json:"items"`
	model.Pagination
}

type KnowledgeNodeStudentResourcesResponse struct {
	Items []*KnowledgeNodeStudentResourceStat `json:"items"`
	model.Pagination
}

type KnowledgeNodeStudentActivitiesResponse struct {
	Items []*KnowledgeNodeStudentActivityStat `json:"items"`
	model.Pagination
}
