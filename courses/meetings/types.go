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
