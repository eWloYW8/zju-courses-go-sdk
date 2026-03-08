package model

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

// SharedResourceSlide represents a linked slide reference on a shared resource.
type SharedResourceSlide struct {
	ID         int    `json:"id"`
	Title      string `json:"title,omitempty"`
	VideoID    int    `json:"video_id,omitempty"`
	DemandID   int    `json:"demand_id,omitempty"`
	TemplateID int    `json:"template_id,omitempty"`
}

// SharedResourceCoursePackage represents a linked course-package reference.
type SharedResourceCoursePackage struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	IsFolder bool   `json:"is_folder,omitempty"`
	ParentID int    `json:"parent_id,omitempty"`
}

// SharedResourceLessonResource represents a linked lesson-resource reference.
type SharedResourceLessonResource struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Mimetype string `json:"mimetype,omitempty"`
	AppID    int    `json:"app_id,omitempty"`
}

// SharedResource represents a shared resource.
type SharedResource struct {
	ID                   int                           `json:"id"`
	Name                 string                        `json:"name,omitempty"`
	OrgID                int                           `json:"org_id,omitempty"`
	ParentID             int                           `json:"parent_id,omitempty"`
	CcLicenseID          int                           `json:"cc_license_id,omitempty"`
	ReferrerID           int                           `json:"referrer_id,omitempty"`
	ResourceType         string                        `json:"resource_type,omitempty"`
	AuditStatus          string                        `json:"audit_status,omitempty"`
	OpenScope            string                        `json:"open_scope,omitempty"`
	ReferrerType         string                        `json:"referrer_type,omitempty"`
	AllowDownload        bool                          `json:"allow_download,omitempty"`
	AllowSave            bool                          `json:"allow_save,omitempty"`
	IsFolder             bool                          `json:"is_folder,omitempty"`
	Reported             bool                          `json:"reported,omitempty"`
	Selected             bool                          `json:"selected,omitempty"`
	Checked              bool                          `json:"_checked,omitempty"`
	CcLicenseName        string                        `json:"cc_license_name,omitempty"`
	CcLicenseLink        string                        `json:"cc_license_link,omitempty"`
	CcLicenseCode        string                        `json:"cc_license_code,omitempty"`
	CcLicenseDescription string                        `json:"cc_license_description,omitempty"`
	Upload               *Upload                       `json:"upload,omitempty"`
	SubjectLibrary       *SubjectLib                   `json:"subject_library,omitempty"`
	Slide                *SharedResourceSlide          `json:"slide,omitempty"`
	CoursePackage        *SharedResourceCoursePackage  `json:"course_package,omitempty"`
	LessonResource       *SharedResourceLessonResource `json:"lesson_resource,omitempty"`
	Percentage           *float64                      `json:"percentage,omitempty"`
	CourseCode           string                        `json:"course_code,omitempty"`
}
