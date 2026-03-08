package groups

import (
	"encoding/json"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

type GroupSetsResponse struct {
	GroupSets []*GroupSet `json:"group_sets"`
}

type CourseStudentsResponse struct {
	Students []*model.User `json:"students"`
}

type CourseEnrollmentsResponse struct {
	Enrollments []*model.Enrollment `json:"enrollments"`
}

type GroupsResponse struct {
	Groups []*Group `json:"groups"`
}

type GroupActivitiesResponse struct {
	Activities []*model.Activity `json:"activities"`
}

type GroupSubmissionStatusResponse struct {
	Statuses json.RawMessage `json:"statuses"`
}
