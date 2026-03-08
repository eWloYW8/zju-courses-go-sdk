package resources

import "github.com/eWloYW8/zju-courses-go-sdk/model"

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
