package meetings

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type VTRSMeetingClassification struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

type VTRSResourceClassification struct {
	ID            int    `json:"id"`
	Name          string `json:"name,omitempty"`
	Category      string `json:"category,omitempty"`
	Deletable     bool   `json:"deletable,omitempty"`
	ResourceCount int    `json:"resource_count,omitempty"`
	Sort          int    `json:"sort,omitempty"`
	VTRSID        int    `json:"vtrs_id,omitempty"`
	AllowCopy     bool   `json:"allow_copy,omitempty"`
	AllowDownload bool   `json:"allow_download,omitempty"`
	CCLicenseName string `json:"cc_license_name,omitempty"`
}

type VTRSUpload struct {
	ID                  int    `json:"id"`
	Name                string `json:"name,omitempty"`
	Size                int64  `json:"size,omitempty"`
	Type                string `json:"type,omitempty"`
	Status              string `json:"status,omitempty"`
	AllowDownload       bool   `json:"allow_download,omitempty"`
	Audio               bool   `json:"audio,omitempty"`
	IsCCVideo           bool   `json:"is_cc_video,omitempty"`
	Link                any    `json:"link,omitempty"`
	OriginAllowDownload bool   `json:"origin_allow_download,omitempty"`
	OwnerID             int    `json:"owner_id,omitempty"`
	ReferenceID         int    `json:"reference_id,omitempty"`
	ReferencedAt        string `json:"referenced_at,omitempty"`
	CreatedByID         int    `json:"created_by_id,omitempty"`
	CreatedAt           string `json:"created_at,omitempty"`
}

type VTRSResourceSummary struct {
	All          int `json:"all,omitempty"`
	Materials    int `json:"materials,omitempty"`
	ExerciseLibs int `json:"exercise_libs,omitempty"`
}

type VTRSResourceStat struct {
	SubjectLib int `json:"subject_lib,omitempty"`
	Video      int `json:"video,omitempty"`
	Audio      int `json:"audio,omitempty"`
	Image      int `json:"image,omitempty"`
	Document   int `json:"document,omitempty"`
	Link       int `json:"link,omitempty"`
	Other      int `json:"other,omitempty"`
}

type VTRSResource struct {
	ID            int         `json:"id"`
	Name          string      `json:"name,omitempty"`
	CreatedByID   int         `json:"created_by_id,omitempty"`
	Upload        *VTRSUpload `json:"upload,omitempty"`
	AllowDownload bool        `json:"allow_download,omitempty"`
	VTRSName      string      `json:"vtrs_name,omitempty"`
	VTRSCode      string      `json:"vtrs_code,omitempty"`
	AllowCopy     bool        `json:"allow_copy,omitempty"`
	CreatedBy     *model.User `json:"created_by,omitempty"`
	CCLicenseName string      `json:"cc_license_name,omitempty"`
	CreatedAt     string      `json:"created_at,omitempty"`
}

type VTRSResourceCategoryNode struct {
	ID           int                       `json:"id"`
	Name         string                    `json:"name,omitempty"`
	Sort         int                       `json:"sort,omitempty"`
	Expand       bool                      `json:"expand,omitempty"`
	HasSubFolder bool                      `json:"has_sub_folder,omitempty"`
	Children     []*VTRSResourceFolderNode `json:"children,omitempty"`
}

type VTRSResourceFolderNode struct {
	ID               int                       `json:"id"`
	Name             string                    `json:"name,omitempty"`
	ClassificationID int                       `json:"classification_id,omitempty"`
	ParentID         int                       `json:"parent_id,omitempty"`
	Expand           bool                      `json:"expand,omitempty"`
	HasSubFolder     bool                      `json:"has_sub_folder,omitempty"`
	Children         []*VTRSResourceFolderNode `json:"children,omitempty"`
	IsFolder         bool                      `json:"is_folder,omitempty"`
}

type VTRSSubjectLib struct {
	ID               int               `json:"id"`
	Title            string            `json:"title,omitempty"`
	ParentID         int               `json:"parent_id,omitempty"`
	ClassificationID int               `json:"classification_id,omitempty"`
	IsFolder         bool              `json:"is_folder,omitempty"`
	IsShared         bool              `json:"is_shared,omitempty"`
	Nums             int               `json:"nums,omitempty"`
	Type             string            `json:"type,omitempty"`
	CreatedAt        string            `json:"created_at,omitempty"`
	UpdatedAt        string            `json:"updated_at,omitempty"`
	Children         []*VTRSSubjectLib `json:"children,omitempty"`
}

type VTRSMemberOption struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id,omitempty"`
	UserNo      string `json:"user_no,omitempty"`
	Name        string `json:"name,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Email       string `json:"email,omitempty"`
	Department  string `json:"department,omitempty"`
	MobilePhone string `json:"mobile_phone,omitempty"`
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

type ZoomSettings struct {
	Mode                        string  `json:"mode,omitempty"`
	BasicDefaultRecordingType   *string `json:"basic_default_recording_type,omitempty"`
	LicenseDefaultRecordingType *string `json:"license_default_recording_type,omitempty"`
}

type OrgZoomSettingsResponse struct {
	OrgZoomSettings *ZoomSettings `json:"org_zoom_settings,omitempty"`
}

type ZoomUserInfo struct {
	Code    int     `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Type    int     `json:"type,omitempty"`
	Email   string  `json:"email,omitempty"`
}

type ZoomUserInfoResponse struct {
	ZoomInfo *ZoomUserInfo `json:"zoom_info,omitempty"`
}
