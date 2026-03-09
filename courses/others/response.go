package others

import (
	"encoding/json"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

type ProjectsResponse struct {
	Items []*Project `json:"items"`
	model.Pagination
}

type ProjectApplicationsResponse struct {
	Items []*ProjectApplication `json:"items"`
	model.Pagination
}

type EntriesResponse struct {
	Items []*Entry `json:"items"`
	model.Pagination
}

type EntryReferencesResponse struct {
	Items []json.RawMessage `json:"items"`
	model.Pagination
}

type RoomLocationsResponse struct {
	Rooms []*RoomLocation `json:"rooms"`
}
