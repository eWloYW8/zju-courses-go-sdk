package resources

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type ResourceGroupsResponse struct {
	ResourceGroups []*ResourceGroup `json:"resource_groups"`
	model.Pagination
}

type ResourceGroupFoldersResponse struct {
	Folders []*ResourceFolder `json:"folders"`
	model.Pagination
}

type ResourceGroupResourcesResponse struct {
	Resources []*model.SharedResource `json:"resources"`
	Page      int                     `json:"page"`
	Pages     int                     `json:"pages"`
	Total     int                     `json:"total"`
}

type ResourceGroupMembersResponse struct {
	Members []*model.User `json:"members"`
	model.Pagination
}

type SharedResourcesResponse struct {
	Resources []*model.SharedResource `json:"resources"`
	model.Pagination
}

type CoursePackagesResponse struct {
	Items []*CoursePackage `json:"items"`
	model.Pagination
}

type PublicCoursesResponse struct {
	Courses []*model.Course `json:"courses"`
}

type CoursePackageExportStatusResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name,omitempty"`
	Size          string `json:"size,omitempty"`
	Status        string `json:"status,omitempty"`
	URL           string `json:"url,omitempty"`
	LastUpdatedAt string `json:"last_updated_at,omitempty"`
}

type CoursePackage struct {
	ID            int    `json:"id"`
	Name          string `json:"name,omitempty"`
	Size          string `json:"size,omitempty"`
	Status        string `json:"status,omitempty"`
	URL           string `json:"url,omitempty"`
	LastUpdatedAt string `json:"last_updated_at,omitempty"`
}
