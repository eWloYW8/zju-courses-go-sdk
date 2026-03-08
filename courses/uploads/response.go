package uploads

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type UploadsListResponse struct {
	Uploads []*Upload `json:"uploads"`
	model.Pagination
}

type UploadReferencesResponse struct {
	References []*UploadReference `json:"references,omitempty"`
	model.Pagination
}

type UploadURLResponse struct {
	URL string `json:"url"`
}

type MoodlePackage struct {
	ID   int                `json:"id"`
	Data *MoodlePackageData `json:"data,omitempty"`
}

type MoodlePackageData struct {
	CourseName     string   `json:"course_name,omitempty"`
	CourseCode     string   `json:"course_code,omitempty"`
	AcademicYear   string   `json:"academic_year,omitempty"`
	Semester       string   `json:"semester,omitempty"`
	Department     string   `json:"department,omitempty"`
	College        string   `json:"college,omitempty"`
	SpocCourseName string   `json:"spoc_course_name,omitempty"`
	Instructors    []string `json:"instructors,omitempty"`
	UseCount       int      `json:"use_count,omitempty"`
}

type MoodlePackagesResponse struct {
	Items []*MoodlePackage `json:"items"`
	model.Pagination
}
