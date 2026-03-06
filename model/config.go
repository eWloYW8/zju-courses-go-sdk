package model

// Config represents supported file format configuration.
type Config struct {
	SupportedAudioFormats      []string `json:"SUPPORTED_AUDIO_FORMATS"`
	SupportedDocumentFormats   []string `json:"SUPPORTED_DOCUMENT_FORMATS"`
	SupportedFlashFormats      []string `json:"SUPPORTED_FLASH_FORMATS"`
	SupportedImageFormats      []string `json:"SUPPORTED_IMAGE_FORMATS"`
	SupportedOpenOfficeFormats []string `json:"SUPPORTED_OPEN_OFFICE_FORMATS"`
	SupportedVideoFormats      []string `json:"SUPPORTED_VIDEO_FORMATS"`
}

// GlobalConfig represents the organization's global configuration.
type GlobalConfig struct {
	APM                             *APMConfig `json:"apm,omitempty"`
	AssetsPath                      string     `json:"assets_path,omitempty"`
	SentryClientKey                 *string    `json:"sentry_client_key,omitempty"`
	SupportedConvertDocumentFormats []string   `json:"supported_convert_document_formats,omitempty"`
	UploadExtensionFormatAllowlist  []string   `json:"upload_extension_format_allowlist,omitempty"`
}

// APMConfig represents APM (Application Performance Monitoring) configuration.
type APMConfig struct {
	Debug                 bool    `json:"DEBUG"`
	EnableAPM             bool    `json:"ENABLE_APM"`
	Environment           string  `json:"ENVIRONMENT"`
	ServerURL             string  `json:"SERVER_URL"`
	ServiceName           string  `json:"SERVICE_NAME"`
	TransactionSampleRate float64 `json:"TRANSACTION_SAMPLE_RATE"`
}

// CalendarEvent represents a calendar event.
type CalendarEvent struct {
	ID          int         `json:"id"`
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	StartTime   string      `json:"start_time,omitempty"`
	EndTime     string      `json:"end_time,omitempty"`
	CourseID    *int        `json:"course_id,omitempty"`
	CourseName  string      `json:"course_name,omitempty"`
	Type        string      `json:"type,omitempty"`
	AllDay      bool        `json:"all_day,omitempty"`
	Location    string      `json:"location,omitempty"`
	CreatedAt   string      `json:"created_at,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

// CalendarTimetable represents a timetable entry.
type CalendarTimetable struct {
	ID        int         `json:"id"`
	Title     string      `json:"title,omitempty"`
	StartTime string      `json:"start_time,omitempty"`
	EndTime   string      `json:"end_time,omitempty"`
	Location  string      `json:"location,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

// UserVisit represents user visit statistics.
type UserVisit struct {
	Date  string      `json:"date,omitempty"`
	Count int         `json:"count,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// LearningActivity represents learning activity statistics.
type LearningActivity struct {
	Data interface{} `json:"data,omitempty"`
}

// ResourceGroup represents a resource group.
type ResourceGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	OrgID       int    `json:"org_id,omitempty"`
	MemberCount int    `json:"member_count,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	CreatedBy   *User  `json:"created_by,omitempty"`
}

// ResourceFolder represents a resource folder.
type ResourceFolder struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	ParentID int    `json:"parent_id,omitempty"`
}

// SubjectLib represents a subject library.
type SubjectLib struct {
	ID        int    `json:"id"`
	Title     string `json:"title,omitempty"`
	Type      string `json:"type,omitempty"`
	Nums      int    `json:"nums,omitempty"`
	IsFolder  bool   `json:"is_folder,omitempty"`
	IsShared  bool   `json:"is_shared,omitempty"`
	ParentID  int    `json:"parent_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// SharedResource represents a shared resource.
type SharedResource struct {
	ID                   int         `json:"id"`
	Name                 string      `json:"name,omitempty"`
	OrgID                int         `json:"org_id,omitempty"`
	ParentID             int         `json:"parent_id,omitempty"`
	ResourceType         string      `json:"resource_type,omitempty"`
	AuditStatus          string      `json:"audit_status,omitempty"`
	OpenScope            string      `json:"open_scope,omitempty"`
	AllowDownload        bool        `json:"allow_download,omitempty"`
	AllowSave            bool        `json:"allow_save,omitempty"`
	IsFolder             bool        `json:"is_folder,omitempty"`
	CcLicenseName        string      `json:"cc_license_name,omitempty"`
	CcLicenseCode        string      `json:"cc_license_code,omitempty"`
	CcLicenseDescription string      `json:"cc_license_description,omitempty"`
	Upload               *Upload     `json:"upload,omitempty"`
	SubjectLibrary       *SubjectLib `json:"subject_library,omitempty"`
}

// LangSettingsResponse represents language settings response.
type LangSettingsResponse struct {
	LangSettings []string `json:"lang_settings"`
}
