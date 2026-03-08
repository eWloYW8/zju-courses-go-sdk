package groups

import "encoding/json"

// ListGroupSetsParams contains optional query params for listing group sets.
type ListGroupSetsParams struct {
	PreloadID *int `json:"-"`
}

// GetGroupStudentsParams contains optional query params for listing students.
type GetGroupStudentsParams struct {
	IgnoreAvatar bool   `json:"-"`
	Fields       string `json:"-"`
}

// ListGroupsParams contains optional query params for listing groups.
type ListGroupsParams struct {
	Fields string `json:"-"`
}

// CopyGroupSetRequest is the request body used by the frontend when copying a group set.
type CopyGroupSetRequest struct {
	Name string `json:"name"`
}

// SortGroupsRequest is the request body used by the frontend when sorting groups.
type SortGroupsRequest struct {
	GroupsSorting json.RawMessage `json:"groupsSorting"`
}

// UploadGroupFilesRequest is the request body used by the frontend when attaching uploads to a group.
type UploadGroupFilesRequest struct {
	Uploads json.RawMessage `json:"uploads"`
}

// BatchDownloadUploadsRequest is the request body used by the frontend when requesting a batch blob.
type BatchDownloadUploadsRequest struct {
	ID json.RawMessage `json:"id"`
}
