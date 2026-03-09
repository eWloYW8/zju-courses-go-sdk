package resources

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

// ResourceGroup represents a resource group.
type ResourceGroup struct {
	ID          int         `json:"id"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	OrgID       int         `json:"org_id,omitempty"`
	MemberCount int         `json:"member_count,omitempty"`
	CreatedAt   string      `json:"created_at,omitempty"`
	CreatedBy   *model.User `json:"created_by,omitempty"`
}

// ResourceFolder represents a resource folder.
type ResourceFolder struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	ParentID int    `json:"parent_id,omitempty"`
}

type SharedResource = model.SharedResource

type ResourceClassification struct {
	ID       int     `json:"id"`
	Name     string  `json:"name,omitempty"`
	ParentID int     `json:"parent_id,omitempty"`
	Cover    *string `json:"cover,omitempty"`
}

type SubjectLibFolder struct {
	ID           int                 `json:"id"`
	Name         string              `json:"name,omitempty"`
	Title        string              `json:"title,omitempty"`
	ParentID     int                 `json:"parent_id,omitempty"`
	HasSubFolder bool                `json:"has_sub_folder,omitempty"`
	Level        int                 `json:"level,omitempty"`
	Selected     bool                `json:"selected,omitempty"`
	Expanded     bool                `json:"expanded,omitempty"`
	Children     []*SubjectLibFolder `json:"children,omitempty"`
}
