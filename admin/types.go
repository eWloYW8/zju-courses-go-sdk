package admin

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
