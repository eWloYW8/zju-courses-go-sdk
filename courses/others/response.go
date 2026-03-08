package others

import (
	"encoding/json"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

type ProjectsResponse struct {
	Items []json.RawMessage `json:"items"`
	model.Pagination
}

type ProjectApplicationsResponse struct {
	Items []json.RawMessage `json:"items"`
	model.Pagination
}
