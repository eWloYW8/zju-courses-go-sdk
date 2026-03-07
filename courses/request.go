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

type ImportKfsCourseRequest struct {
	KFSCourseID  int `json:"kfs_course_id"`
	KFSVersionID int `json:"kfs_version_id"`
}

type TeachingObjectivesRequest struct {
	TeachingObjectives []*TeachingObjective `json:"teaching_objectives"`
}

type DeleteTeachingObjectivesRequest struct {
	TeachingObjectiveIDs []int `json:"teaching_objective_ids"`
}

type KnowledgeNodeRelationRequest struct {
	ID         int  `json:"id,omitempty"`
	Source     int  `json:"source,omitempty"`
	Target     int  `json:"target,omitempty"`
	IsDirected bool `json:"is_directed,omitempty"`
}

type DeleteKnowledgeNodeRelationsRequest struct {
	RelationIDs []int `json:"relation_ids"`
}

type DeleteKnowledgeNodesRequest struct {
	KnowledgeNodeIDs []int `json:"knowledge_node_ids"`
}

type MoveKnowledgeNodeRequest struct {
	ID       int  `json:"id,omitempty"`
	ParentID *int `json:"parent_id,omitempty"`
	Sort     int  `json:"sort,omitempty"`
}

type DeleteMediaKnowledgeReferenceRequest struct {
	ChapterID int `json:"chapter_id"`
}

type DeleteUploadKnowledgeReferenceRequest struct {
	UploadID int `json:"upload_id"`
}

type UpdateKnowledgeGraphSourceRequest struct {
	Source string `json:"source"`
}

type ImportKnowledgeNodesByCourseRequest struct {
	SourceCourseID      int  `json:"source_course_id"`
	ImportReferResource bool `json:"import_refer_resource"`
}

type SyncBlueprintRequest map[string]any
