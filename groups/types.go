package groups

import "github.com/eWloYW8/zju-courses-go-sdk/model"

// GroupSet represents a group set containing multiple groups in a course.
type GroupSet struct {
	ID         int      `json:"id"`
	Name       string   `json:"name,omitempty"`
	CourseID   int      `json:"course_id,omitempty"`
	GroupCount int      `json:"group_count,omitempty"`
	Groups     []*Group `json:"groups,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty"`
	UpdatedAt  string   `json:"updated_at,omitempty"`
}

// Group represents a student group within a group set.
type Group struct {
	ID         int           `json:"id"`
	Name       string        `json:"name,omitempty"`
	GroupSetID int           `json:"group_set_id,omitempty"`
	LeaderID   *int          `json:"leader_id,omitempty"`
	Members    []*model.User `json:"members,omitempty"`
}
