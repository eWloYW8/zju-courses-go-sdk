package meetings

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
