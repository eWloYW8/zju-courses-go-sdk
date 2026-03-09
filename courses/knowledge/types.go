package knowledge

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

// KnowledgeNode represents a knowledge node in the course knowledge tree.
type KnowledgeNode struct {
	ID                 int                          `json:"id"`
	Name               string                       `json:"name,omitempty"`
	ParentID           *int                         `json:"parent_id,omitempty"`
	CourseID           int                          `json:"course_id,omitempty"`
	Sort               int                          `json:"sort,omitempty"`
	Level              int                          `json:"level,omitempty"`
	Description        string                       `json:"description,omitempty"`
	Source             string                       `json:"source,omitempty"`
	States             *KnowledgeNodeState          `json:"states,omitempty"`
	MasteryRate        *float64                     `json:"mastery_rate,omitempty"`
	CompletenessRate   *float64                     `json:"completeness_rate,omitempty"`
	ResourceRefCount   int                          `json:"resource_ref_count,omitempty"`
	ActivityRefCount   int                          `json:"activity_ref_count,omitempty"`
	SubjectRefCount    int                          `json:"subject_ref_count,omitempty"`
	CaptureRefCount    int                          `json:"capture_ref_count,omitempty"`
	CognitiveDimension string                       `json:"cognitive_dimension,omitempty"`
	PrevRelation       []*KnowledgeNodeRelationMeta `json:"prev_relation,omitempty"`
	PostRelation       []*KnowledgeNodeRelationMeta `json:"post_relation,omitempty"`
	UndirectedRelation []*KnowledgeNodeRelationMeta `json:"undirected_relation,omitempty"`
	Children           []*KnowledgeNode             `json:"children,omitempty"`
	Activities         []*model.Activity            `json:"activities,omitempty"`
	TeachingObjectives []*TeachingObjective         `json:"teaching_objectives,omitempty"`
	Uploads            []*model.Upload              `json:"uploads,omitempty"`
	ExternalResources  []*KnowledgeExternalResource `json:"external_resources,omitempty"`
	Captures           []*KnowledgeCapture          `json:"captures,omitempty"`
	Subjects           []*KnowledgeSubject          `json:"subjects,omitempty"`
	Labels             []*KnowledgeLabel            `json:"labels,omitempty"`
}

type KnowledgeCapture struct {
	ID        int     `json:"id"`
	Code      string  `json:"code,omitempty"`
	Title     string  `json:"title,omitempty"`
	Name      string  `json:"name,omitempty"`
	OrgCode   string  `json:"org_code,omitempty"`
	CourseID  *int    `json:"course_id,omitempty"`
	UploadID  *int    `json:"upload_id,omitempty"`
	URL       string  `json:"url,omitempty"`
	Cover     *string `json:"cover,omitempty"`
	Deleted   bool    `json:"deleted,omitempty"`
	Viewed    bool    `json:"viewed,omitempty"`
	ReferID   int     `json:"refer_id,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
}

type KnowledgeNodeRecommendedResourceReference struct {
	ID     int           `json:"id"`
	Name   string        `json:"name,omitempty"`
	Upload *model.Upload `json:"upload,omitempty"`
}

type KnowledgeNodeCompletenessItem struct {
	ActivityID   int `json:"activity_id,omitempty"`
	ActivityType int `json:"activity_type,omitempty"`
	Completeness int `json:"completeness,omitempty"`
}

type KnowledgeNodeSummaryNode struct {
	ID                      int                         `json:"id"`
	ParentID                *int                        `json:"parent_id,omitempty"`
	Name                    string                      `json:"name,omitempty"`
	AverageMasteryRate      *float64                    `json:"average_mastery_rate,omitempty"`
	AverageCompletenessRate *float64                    `json:"average_completeness_rate,omitempty"`
	CognitiveDimension      string                      `json:"cognitive_dimension,omitempty"`
	Children                []*KnowledgeNodeSummaryNode `json:"children,omitempty"`
}

type KnowledgeNodeStatisticsSummary struct {
	NodeCount                    int                         `json:"node_count,omitempty"`
	NodeWithReferenceCount       int                         `json:"node_with_reference_count,omitempty"`
	AverageMasteryRate           *float64                    `json:"average_mastery_rate,omitempty"`
	AverageCompletenessRate      *float64                    `json:"average_completeness_rate,omitempty"`
	RelationCount                int                         `json:"relation_count,omitempty"`
	CompletenessRateDistribution any                         `json:"completeness_rate_distribution,omitempty"`
	MasteryRateDistribution      any                         `json:"mastery_rate_distribution,omitempty"`
	Nodes                        []*KnowledgeNodeSummaryNode `json:"nodes,omitempty"`
}

type KnowledgeNodeOverview struct {
	ID                           int      `json:"id"`
	Name                         string   `json:"name,omitempty"`
	Description                  string   `json:"description,omitempty"`
	AverageMasteryRate           *float64 `json:"avg_mastery_rate,omitempty"`
	AverageCompletenessRate      *float64 `json:"avg_completeness_rate,omitempty"`
	UploadCount                  int      `json:"upload_count,omitempty"`
	ActivityCount                int      `json:"activity_count,omitempty"`
	CompletenessRateDistribution any      `json:"completeness_rate_distribution,omitempty"`
	MasteryRateDistribution      any      `json:"mastery_rate_distribution,omitempty"`
}

type KnowledgeNodeStudentDimensionItem struct {
	StudentID        int      `json:"student_id,omitempty"`
	Name             string   `json:"name,omitempty"`
	UserNo           string   `json:"user_no,omitempty"`
	MasteryRate      *float64 `json:"mastery_rate,omitempty"`
	CompletenessRate *float64 `json:"completeness_rate,omitempty"`
	ImportedFrom     string   `json:"imported_from,omitempty"`
}

type KnowledgeNodeStudentDetail struct {
	StudentID            int                         `json:"student_id,omitempty"`
	Name                 string                      `json:"name,omitempty"`
	UserNo               string                      `json:"user_no,omitempty"`
	MasteryRate          *float64                    `json:"mastery_rate,omitempty"`
	CompletenessRate     *float64                    `json:"completeness_rate,omitempty"`
	ImportedFrom         string                      `json:"imported_from,omitempty"`
	MasteryRateRank      *int                        `json:"mastery_rate_rank,omitempty"`
	CompletenessRateRank *int                        `json:"completeness_rate_rank,omitempty"`
	Nodes                []*KnowledgeNodeSummaryNode `json:"nodes,omitempty"`
}

type KnowledgeNodeResourceDetail struct {
	ID                int      `json:"id"`
	Name              string   `json:"name,omitempty"`
	Type              string   `json:"type,omitempty"`
	AllowDownload     bool     `json:"allow_download,omitempty"`
	Extension         string   `json:"extension,omitempty"`
	CompletionStudent string   `json:"completion_student,omitempty"`
	Completeness      *float64 `json:"completeness,omitempty"`
	Visits            int      `json:"visits,omitempty"`
	Source            string   `json:"source,omitempty"`
	Status            string   `json:"status,omitempty"`
}

type KnowledgeNodeActivityDetail struct {
	ID                int      `json:"id"`
	Name              string   `json:"name,omitempty"`
	Type              string   `json:"type,omitempty"`
	CompletionStudent string   `json:"completion_student,omitempty"`
	Completeness      *float64 `json:"completeness,omitempty"`
	MasteryRate       *float64 `json:"mastery_rate,omitempty"`
}

type KnowledgeNodeStudentResourceStat struct {
	ID     int    `json:"id"`
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Viewed bool   `json:"viewed,omitempty"`
	Visits int    `json:"visits,omitempty"`
	Status string `json:"status,omitempty"`
	Source string `json:"source,omitempty"`
}

type KnowledgeNodeStudentActivityStat struct {
	ID           int    `json:"id"`
	Title        string `json:"title,omitempty"`
	Type         string `json:"type,omitempty"`
	Completeness string `json:"completeness,omitempty"`
	Score        any    `json:"score,omitempty"`
}

type KnowledgeNodeStudentOverallStatistics struct {
	OverallCompletenessRate *float64 `json:"overall_completeness_rate,omitempty"`
	OverallMasteryRate      *float64 `json:"overall_mastery_rate,omitempty"`
}

type KnowledgeGraphUser struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

type KnowledgeGraphKFSCourse struct {
	ID              int                   `json:"id"`
	Name            string                `json:"name,omitempty"`
	TopicCount      int                   `json:"topic_count,omitempty"`
	DependencyCount int                   `json:"dependency_count,omitempty"`
	FacetCount      int                   `json:"facet_count,omitempty"`
	FragmentCount   int                   `json:"fragment_count,omitempty"`
	Builders        []*KnowledgeGraphUser `json:"builders,omitempty"`
	Principals      []*KnowledgeGraphUser `json:"principals,omitempty"`
	LatestVersionID int                   `json:"latest_version_id,omitempty"`
}

type KnowledgeGraphKFSSubject struct {
	Name     string                     `json:"name,omitempty"`
	ID       int                        `json:"id"`
	DataType string                     `json:"data_type,omitempty"`
	Courses  []*KnowledgeGraphKFSCourse `json:"courses,omitempty"`
}

type KnowledgeGraphKFSImportInfo struct {
	LastImportTime  int64  `json:"last_import_time,omitempty"`
	ImportCourseIDs []int  `json:"import_course_ids,omitempty"`
	Server          string `json:"server,omitempty"`
}

type KnowledgeGraphForestVersionStats struct {
	TopicCount      int `json:"topic_count,omitempty"`
	DependencyCount int `json:"dependency_count,omitempty"`
	FacetCount      int `json:"facet_count,omitempty"`
	FragmentCount   int `json:"fragment_count,omitempty"`
}

type KnowledgeGraphForestVersionStatsItem struct {
	ID int `json:"id"`
	KnowledgeGraphForestVersionStats
}

type KnowledgeGraphPublishedForestVersion struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	CourseID  int    `json:"course_id,omitempty"`
	Published bool   `json:"published,omitempty"`
}

type KnowledgeGraphDiff struct {
	Action   string `json:"action,omitempty"`
	Type     string `json:"type,omitempty"`
	NodeID   int    `json:"node_id,omitempty"`
	NodeName string `json:"node_name,omitempty"`
	Before   any    `json:"before,omitempty"`
	After    any    `json:"after,omitempty"`
	Data     any    `json:"data,omitempty"`
}

type KnowledgeGraphSimilarity struct {
	ID               int     `json:"id"`
	Name             string  `json:"name,omitempty"`
	Similarity       float64 `json:"similarity,omitempty"`
	SimilarityFormat string  `json:"similarity_format,omitempty"`
	Checked          bool    `json:"checked,omitempty"`
}

type KnowledgeExtensionApp map[string]any

type KnowledgeGraphCluster map[string]any

type KnowledgeGraphSyncStatus map[string]any

type KnowledgeNodeWeekStat struct {
	ID               int     `json:"id"`
	CompletenessRate float64 `json:"avg_completeness_rate,omitempty"`
	MasteryRate      float64 `json:"avg_mastery_rate,omitempty"`
	StatDate         string  `json:"stat_date,omitempty"`
	Week             int     `json:"week,omitempty"`
}

type KnowledgeGraphSnapshotRelation struct {
	ID         int    `json:"id"`
	Source     int    `json:"source,omitempty"`
	Target     int    `json:"target,omitempty"`
	IsDirected bool   `json:"is_directed,omitempty"`
	Color      string `json:"color,omitempty"`
}

type KnowledgeGraphSnapshot struct {
	Tree         *KnowledgeNode                    `json:"tree,omitempty"`
	Relations    []*KnowledgeGraphSnapshotRelation `json:"relations,omitempty"`
	Completeness []float64                         `json:"completeness,omitempty"`
	Mastery      []float64                         `json:"mastery,omitempty"`
}

type KnowledgeNodeRelation struct {
	ID         int  `json:"id"`
	Source     int  `json:"source,omitempty"`
	Target     int  `json:"target,omitempty"`
	IsDirected bool `json:"is_directed,omitempty"`
}

type CoursewareActivity struct {
	ID             int    `json:"id"`
	Title          string `json:"title,omitempty"`
	Type           string `json:"type,omitempty"`
	PublishType    string `json:"publish_type,omitempty"`
	StartTime      string `json:"start_time,omitempty"`
	EndTime        string `json:"end_time,omitempty"`
	IsClosed       bool   `json:"is_closed,omitempty"`
	Score          any    `json:"score,omitempty"`
	Completeness   any    `json:"completeness,omitempty"`
	ScorePublished bool   `json:"score_published,omitempty"`
	Published      bool   `json:"published,omitempty"`
	IsLocked       bool   `json:"is_locked,omitempty"`
	Prerequisites  any    `json:"prerequisites,omitempty"`
	ModuleID       int    `json:"module_id,omitempty"`
	SyllabusID     int    `json:"syllabus_id,omitempty"`
	ModuleName     string `json:"module_name,omitempty"`
	SyllabusName   string `json:"syllabus_name,omitempty"`
}

type KnowledgeNodeStudentReferenceStat struct {
	NodeName         string   `json:"node_name,omitempty"`
	NodeDesc         string   `json:"node_desc,omitempty"`
	Name             string   `json:"name,omitempty"`
	UserNo           string   `json:"user_no,omitempty"`
	MasteryRank      *int     `json:"mastery_rank,omitempty"`
	MasteryRate      *float64 `json:"mastery_rate,omitempty"`
	CompletenessRank *int     `json:"completeness_rank,omitempty"`
	CompletenessRate *float64 `json:"completeness_rate,omitempty"`
	ResourceCount    int      `json:"resource_count,omitempty"`
	ActivityCount    int      `json:"activity_count,omitempty"`
}

type KnowledgeBase map[string]any

type KnowledgeBaseResource map[string]any

type KnowledgeBaseResourcesResponse struct {
	Items []KnowledgeBaseResource `json:"items,omitempty"`
	model.Pagination
}

type KnowledgeBaseMutationResponse map[string]any

type KnowledgeNodeReferTypeStats struct {
	NodeCount     int `json:"node_count,omitempty"`
	ResourceCount int `json:"resource_ref_count,omitempty"`
	ActivityCount int `json:"activity_ref_count,omitempty"`
}

type TeachingObjective struct {
	ID         int    `json:"id"`
	Content    string `json:"content,omitempty"`
	ReferCount int    `json:"refer_count,omitempty"`
}

type KnowledgeFragmentDetail struct {
	Name        string `json:"name,omitempty"`
	Book        string `json:"book,omitempty"`
	ISBN        string `json:"isbn,omitempty"`
	Page        int    `json:"page,omitempty"`
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
}

type KnowledgeFragment struct {
	ID      int                      `json:"id"`
	FacetID int                      `json:"facet_id,omitempty"`
	Content string                   `json:"content,omitempty"`
	Source  string                   `json:"source,omitempty"`
	Detail  *KnowledgeFragmentDetail `json:"detail,omitempty"`
}

type KnowledgeFacet struct {
	ID        int                  `json:"id"`
	Name      string               `json:"name,omitempty"`
	ParentID  int                  `json:"parent_id,omitempty"`
	TopicID   int                  `json:"topic_id,omitempty"`
	Children  []*KnowledgeFacet    `json:"children,omitempty"`
	Fragments []*KnowledgeFragment `json:"fragments,omitempty"`
}

type KnowledgeReferenceData struct {
	MaterialUploads []*model.Upload           `json:"material_uploads,omitempty"`
	MediaFragments  []*KnowledgeMediaFragment `json:"media_fragments,omitempty"`
}

type ActivityKnowledgeReference struct {
	ID              int                     `json:"id"`
	CourseID        int                     `json:"course_id,omitempty"`
	KnowledgeNodeID int                     `json:"knowledge_node_id,omitempty"`
	ReferID         int                     `json:"refer_id,omitempty"`
	ReferType       string                  `json:"refer_type,omitempty"`
	Data            *KnowledgeReferenceData `json:"data,omitempty"`
}

type KnowledgeNodeState struct {
	Expanded      bool `json:"expanded,omitempty"`
	Selected      bool `json:"selected,omitempty"`
	Checked       bool `json:"checked,omitempty"`
	Indeterminate bool `json:"indeterminate,omitempty"`
	Disabled      bool `json:"disabled,omitempty"`
	Matched       bool `json:"matched,omitempty"`
	Dragging      bool `json:"dragging,omitempty"`
	Draggable     bool `json:"draggable,omitempty"`
	Immutable     bool `json:"immutable,omitempty"`
}

type KnowledgeNodeRelationMeta struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	RelationID int    `json:"relation_id,omitempty"`
	Data       any    `json:"data,omitempty"`
}

type KnowledgeLabel struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type KnowledgeExternalResource struct {
	ID         int    `json:"id,omitempty"`
	ExternalID int    `json:"external_id,omitempty"`
	Title      string `json:"title,omitempty"`
	Type       string `json:"type,omitempty"`
	Viewed     bool   `json:"viewed,omitempty"`
	Visits     int    `json:"visits,omitempty"`
	Data       any    `json:"data,omitempty"`
}

type KnowledgeSubjectSource struct {
	ID           int    `json:"id,omitempty"`
	Type         string `json:"type,omitempty"`
	Name         string `json:"name,omitempty"`
	IsMakeupExam bool   `json:"is_makeup_exam,omitempty"`
}

type KnowledgeNodeReferenceItem struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type KnowledgeSubject struct {
	ID                     int                           `json:"id,omitempty"`
	Type                   string                        `json:"type,omitempty"`
	Description            string                        `json:"description,omitempty"`
	Point                  any                           `json:"point,omitempty"`
	ShowAnswer             bool                          `json:"show_answer,omitempty"`
	ShowExplanation        bool                          `json:"show_explanation,omitempty"`
	Source                 *KnowledgeSubjectSource       `json:"source,omitempty"`
	SubSubjects            []*KnowledgeSubject           `json:"sub_subjects,omitempty"`
	KnowledgeNodeReference []*KnowledgeNodeReferenceItem `json:"knowledge_node_reference,omitempty"`
}

type KnowledgePoint struct {
	ID      int    `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type KnowledgeMediaFragment struct {
	ChapterID          int               `json:"chapter_id,omitempty"`
	ChapterName        string            `json:"chapter_name,omitempty"`
	StartTS            int               `json:"start_ts,omitempty"`
	EndTS              int               `json:"end_ts,omitempty"`
	MediaID            int               `json:"media_id,omitempty"`
	KnowledgeNodeID    int               `json:"knowledge_node_id,omitempty"`
	KnowledgeNodeName  string            `json:"knowledge_node_name,omitempty"`
	KnowledgePointList []*KnowledgePoint `json:"knowledge_point_list,omitempty"`
	Duration           *float64          `json:"duration,omitempty"`
	Completeness       *float64          `json:"completeness,omitempty"`
}
