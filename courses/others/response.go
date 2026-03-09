package others

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type ProjectsResponse struct {
	Items []*Project `json:"items"`
	model.Pagination
}

type ProjectApplicationsResponse struct {
	Items []*ProjectApplication `json:"items"`
	model.Pagination
}

type RoomLocationsResponse struct {
	Rooms []*RoomLocation `json:"rooms"`
}
