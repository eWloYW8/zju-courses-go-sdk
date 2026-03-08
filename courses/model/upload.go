package model

// Upload represents an uploaded file/resource.
type Upload struct {
	ID                              int     `json:"id"`
	Name                            string  `json:"name"`
	Key                             string  `json:"key,omitempty"`
	Size                            int64   `json:"size,omitempty"`
	Type                            string  `json:"type,omitempty"`
	ContentType                     string  `json:"content_type,omitempty"`
	Status                          string  `json:"status,omitempty"`
	Source                          *string `json:"source,omitempty"`
	Link                            *string `json:"link,omitempty"`
	CreatedAt                       string  `json:"created_at,omitempty"`
	UpdatedAt                       string  `json:"updated_at,omitempty"`
	CreatedByID                     int     `json:"created_by_id,omitempty"`
	CreatedBy                       *User   `json:"created_by,omitempty"`
	OwnerID                         int     `json:"owner_id,omitempty"`
	ReferenceID                     int     `json:"reference_id,omitempty"`
	ReferencedAt                    *string `json:"referenced_at,omitempty"`
	Deleted                         bool    `json:"deleted,omitempty"`
	Viewed                          bool    `json:"viewed,omitempty"`
	IsFolder                        bool    `json:"is_folder,omitempty"`
	IsShared                        bool    `json:"is_shared,omitempty"`
	IsCCVideo                       bool    `json:"is_cc_video,omitempty"`
	AllowDownload                   bool    `json:"allow_download,omitempty"`
	OriginAllowDownload             *bool   `json:"origin_allow_download,omitempty"`
	AllowAliyunOfficeView           bool    `json:"allow_aliyun_office_view,omitempty"`
	AllowPrivateWpsOfficeView       bool    `json:"allow_private_wps_office_view,omitempty"`
	EnableSetH5CoursewareCompletion *bool   `json:"enable_set_h5_courseware_completion,omitempty"`
	VideoSrcType                    string  `json:"video_src_type,omitempty"`
	ThirdPartReferrerID             int     `json:"third_part_referrer_id,omitempty"`
	ResourceType                    string  `json:"resource_type,omitempty"`
	ReferenceCount                  int     `json:"reference_count,omitempty"`
	ParentType                      string  `json:"parent_type,omitempty"`
	ParentID                        int     `json:"parent_id,omitempty"`
	ReferID                         int     `json:"refer_id,omitempty"`
	ReferType                       string  `json:"refer_type,omitempty"`
	IsFromKnowledgeGraph            bool    `json:"is_from_knowledge_graph,omitempty"`

	// Media
	Audio     *AudioInfo       `json:"audio,omitempty"`
	Scorm     interface{}      `json:"scorm,omitempty"`
	Thumbnail *UploadThumbnail `json:"thumbnail,omitempty"`
	Videos    []*Video         `json:"videos,omitempty"`
	VideoURLs interface{}      `json:"video_urls,omitempty"` // {"QVGA":"...","VGA":"...","HD":"..."}

	// Permissions
	CaptionPermission          interface{} `json:"caption_permission,omitempty"`
	CaptionSpeechPermission    interface{} `json:"caption_speech_permission,omitempty"`
	CaptionTranslatePermission interface{} `json:"caption_translate_permission,omitempty"`
	ChapterPermission          interface{} `json:"chapter_permission,omitempty"`
	TopicPermission            interface{} `json:"topic_permission,omitempty"`
}

type UploadThumbnail struct {
	ID int `json:"id,omitempty"`
}

// AudioInfo represents audio metadata.
type AudioInfo struct {
	Duration float64 `json:"duration,omitempty"`
}

// Video represents video information for an upload.
type Video struct {
	ID       int     `json:"id,omitempty"`
	URL      string  `json:"url,omitempty"`
	Quality  string  `json:"quality,omitempty"`
	Duration float64 `json:"duration,omitempty"`
}

// UploadReference represents a reference to an uploaded file.
type UploadReference struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Deleted    bool    `json:"deleted,omitempty"`
	OrgID      int     `json:"org_id,omitempty"`
	ParentID   int     `json:"parent_id,omitempty"`
	ParentType string  `json:"parent_type,omitempty"`
	UploadID   int     `json:"upload_id,omitempty"`
	Upload     *Upload `json:"upload,omitempty"`
}
