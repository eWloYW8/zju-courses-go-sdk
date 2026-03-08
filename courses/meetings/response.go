package meetings

import (
	"encoding/json"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

type VTRSShareResourcesResponse struct {
	Items []json.RawMessage `json:"items"`
	model.Pagination
}

type VTRSResourcesResponse struct {
	Items []json.RawMessage `json:"items"`
	model.Pagination
}

type VTRSResourcesSummaryResponse map[string]any

type VTRSSubjectLibsResponse struct {
	Items []json.RawMessage `json:"items"`
	model.Pagination
}

type VTRSMeetingClassificationsResponse struct {
	Classifications []*VTRSMeetingClassification `json:"classifications"`
}

type VTRSResourceClassificationsResponse struct {
	Classifications []*VTRSResourceClassification `json:"classifications"`
}

type RoomLocationsResponse struct {
	Rooms []*RoomLocation `json:"rooms"`
}
