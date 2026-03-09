package meetings

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type VTRSShareResourcesResponse struct {
	Items []*VTRSResource `json:"items"`
	model.Pagination
}

type VTRSResourcesResponse struct {
	Items []*VTRSResource `json:"items"`
	model.Pagination
}

type VTRSResourcesSummaryResponse = VTRSResourceSummary

type VTRSSubjectLibsResponse struct {
	Items []*VTRSSubjectLib `json:"items"`
	model.Pagination
}

type VTRSMemberOptionsResponse struct {
	Items []*VTRSMemberOption `json:"items"`
	model.Pagination
}

type VTRSMeetingClassificationsResponse struct {
	Classifications []*VTRSMeetingClassification `json:"classifications"`
}

type VTRSResourceClassificationsResponse struct {
	Classifications []*VTRSResourceClassification `json:"classifications"`
}

type VTRSResourceCategoryStructureResponse struct {
	Classifications []*VTRSResourceCategoryNode `json:"classifications"`
}

type VTRSResourceOperationPreCheckResponse map[string]any

type RoomLocationsResponse struct {
	Rooms []*RoomLocation `json:"rooms"`
}
