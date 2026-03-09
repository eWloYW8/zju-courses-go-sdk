package others

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type Project struct {
	ID                 int                 `json:"id"`
	Name               string              `json:"name,omitempty"`
	Description        *string             `json:"classroom_schedule,omitempty"`
	KnowledgeNodeCount int                 `json:"knowledge_node_count,omitempty"`
	Instructors        []*model.User       `json:"instructors,omitempty"`
	EnrolledProject    bool                `json:"enrolled_project,omitempty"`
	Audit              *ProjectApplication `json:"audit,omitempty"`
}

type ProjectApplication struct {
	ID        int               `json:"id"`
	Status    string            `json:"status,omitempty"`
	CreatedAt string            `json:"created_at,omitempty"`
	User      *ProjectApplicant `json:"user,omitempty"`
}

type ProjectApplicant struct {
	ID         int               `json:"id"`
	Name       string            `json:"name,omitempty"`
	UserNo     string            `json:"user_no,omitempty"`
	Department *model.Department `json:"department,omitempty"`
}

type ProjectSharedResource struct {
	ID             int              `json:"id"`
	Name           string           `json:"name,omitempty"`
	CreatedAt      string           `json:"created_at,omitempty"`
	CreatedByID    int              `json:"created_by_id,omitempty"`
	RefParentID    *int             `json:"ref_parent_id,omitempty"`
	Upload         *model.Upload    `json:"upload,omitempty"`
	AllowDownload  bool             `json:"allow_download,omitempty"`
	KnowledgeCount int              `json:"knowledge_count,omitempty"`
	KnowledgeNodes []map[string]any `json:"knowledge_nodes,omitempty"`
}

type Entry struct {
	ID             int             `json:"id"`
	OrgID          int             `json:"org_id,omitempty"`
	Name           string          `json:"name,omitempty"`
	Definition     string          `json:"definition,omitempty"`
	Uploads        []*model.Upload `json:"uploads,omitempty"`
	Keywords       []*EntryKeyword `json:"keywords,omitempty"`
	CreatedAt      string          `json:"created_at,omitempty"`
	UpdatedAt      string          `json:"updated_at,omitempty"`
	CreatedByID    int             `json:"created_by_id,omitempty"`
	UpdatedByID    int             `json:"updated_by_id,omitempty"`
	ReferenceCount int             `json:"reference_count,omitempty"`
}

type FeatureEnabledResponse struct {
	Enabled bool `json:"enabled"`
}

type LessonRoom struct {
	RoomCode string `json:"room_code,omitempty"`
	RoomName string `json:"room_name,omitempty"`
	AppID    any    `json:"app_id,omitempty"`
}

type RoomLocation struct {
	ID       int    `json:"id"`
	Building string `json:"building,omitempty"`
	RoomName string `json:"room_name,omitempty"`
	RoomCode string `json:"room_code,omitempty"`
	Seats    int    `json:"seats,omitempty"`
}
