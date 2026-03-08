package admin

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

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

// LangSettingsResponse represents language settings response.
type LangSettingsResponse struct {
	LangSettings []string `json:"lang_settings"`
}

// OutlineSettingResponse represents the outline setting API response.
type OutlineSettingResponse struct {
	ID                      int              `json:"id"`
	OrgID                   int              `json:"org_id"`
	FormattedDefaultOptions []*OutlineOption `json:"formatted_default_options,omitempty"`
	FormattedOptions        []*OutlineOption `json:"formatted_options,omitempty"`
}

// OutlineOption represents an outline setting option.
type OutlineOption struct {
	Key      string `json:"key"`
	Title    string `json:"title"`
	Required bool   `json:"required"`
}

// AssistantRolePermission represents a permission toggle on an assistant role.
type AssistantRolePermission struct {
	Code     string `json:"code"`
	Enabled  bool   `json:"enabled,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

// AssistantRole represents a configurable assistant role in course authz settings.
type AssistantRole struct {
	ID          int                        `json:"id"`
	Name        string                     `json:"name,omitempty"`
	Alias       string                     `json:"alias,omitempty"`
	Permissions []*AssistantRolePermission `json:"permissions,omitempty"`
}

// AllOrgsResponse represents the frontend /api/all-orgs wrapper.
type AllOrgsResponse struct {
	Orgs []*model.OrgDetail `json:"orgs"`
}

// OrgSettings represents organization settings returned by /api/orgs/{id}/settings.
type OrgSettings map[string]any

// LiveRecordOrgSettings represents live-record settings returned by /api/orgs/{id}/live-record-settings.
type LiveRecordOrgSettings map[string]any

// PortalLogo represents the portal-logo payload returned by /api/portal-logo.
type PortalLogo map[string]any

// AlertPopupSettings represents the alert-popup payload returned by /api/orgs/{id}/alert-popup.
type AlertPopupSettings map[string]any

// UpdateOrgSettingRequest represents a generic organization-settings form submission.
type UpdateOrgSettingRequest map[string]any

// UpdateAlertPopupSettingRequest represents a generic alert-popup form submission.
type UpdateAlertPopupSettingRequest map[string]any

// CopyableCourse represents the minimal course item returned by /api/course-copy/courses.
type CopyableCourse struct {
	ID         int    `json:"id"`
	Name       string `json:"name,omitempty"`
	CourseCode string `json:"course_code,omitempty"`
}

// CopyableCoursesResponse represents the /api/course-copy/courses response wrapper.
type CopyableCoursesResponse struct {
	Courses []*CopyableCourse `json:"courses"`
}

// CopyableCoursesQuery represents the frontend search params for /api/course-copy/courses.
type CopyableCoursesQuery struct {
	Keyword string
	Fields  string
}

// MoodleImportRequest represents the JSON body for importing a listed Moodle package into a course.
type MoodleImportRequest struct {
	UploadID int `json:"upload_id"`
}

// MBZImportRequest represents the multipart form body for /api/course/mbz/import.
type MBZImportRequest struct {
	UploadID int `json:"upload_id"`
}

// LastTask represents one entry returned by /api/task/last.
type LastTask struct {
	ID        int             `json:"id,omitempty"`
	Type      string          `json:"type,omitempty"`
	Status    string          `json:"status,omitempty"`
	Output    *LastTaskOutput `json:"output,omitempty"`
	CreatedAt string          `json:"created_at,omitempty"`
	UpdatedAt string          `json:"updated_at,omitempty"`
}

// LastTaskOutput represents the task output payload.
type LastTaskOutput struct {
	Imported *LastTaskImported `json:"imported,omitempty"`
}

// LastTaskImported represents the imported payload nested under task output.
type LastTaskImported struct {
	Course *LastTaskCourse `json:"course,omitempty"`
}

// LastTaskCourse represents the imported course metadata nested in task output.
type LastTaskCourse struct {
	ID int `json:"id,omitempty"`
}
