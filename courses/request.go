package courses

type ListCoursesRequest struct {
	Keyword string `json:"keyword,omitempty"`
}

type ListMyCoursesRequest struct {
	Page       int            `json:"page,omitempty"`
	PageSize   int            `json:"page_size,omitempty"`
	Fields     string         `json:"fields,omitempty"`
	Conditions map[string]any `json:"conditions,omitempty"`
}

type CreateCourseRequest = Course

type UpdateCourseRequest = Course

type UpdateNavSettingRequest struct {
	NavSetting []*NavSetting `json:"nav_setting"`
}

type CreateModuleRequest = Module

type UpdateModuleRequest = Module

type DeleteModuleOptions struct {
	DeleteRelatedActivity bool
}

type DeleteSyllabusOptions struct {
	DeleteRelatedActivity bool
}

type ResortActivitiesRequest struct {
	CourseID    int   `json:"course_id,omitempty"`
	ActivityIDs []int `json:"activity_ids,omitempty"`
	ModuleID    *int  `json:"module_id,omitempty"`
	SyllabusID  *int  `json:"syllabus_id,omitempty"`
}

type SyncFromURPRequest struct {
	CourseIDs []int `json:"course_ids"`
}

type SyncBlueprintRequest map[string]any
