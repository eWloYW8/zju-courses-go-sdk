package knowledge

type ImportKfsCourseRequest struct {
	KFSCourseID  int `json:"kfs_course_id"`
	KFSVersionID int `json:"kfs_version_id"`
}

type ImportKnowledgeNodesRequest struct {
	Data any `json:"data"`
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

type AIParseKnowledgeNodesRequest struct {
	UploadID int `json:"upload_id"`
}
