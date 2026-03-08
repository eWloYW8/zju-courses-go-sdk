package users

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type AcademicYearsResponse struct {
	AcademicYears []*AcademicYear `json:"academic_years"`
}

type SemestersResponse struct {
	Semesters []*Semester `json:"semesters"`
}

type DepartmentsResponse struct {
	Departments []*Department `json:"departments"`
}

type ClassesResponse struct {
	Classes []*Class `json:"classes"`
}

type GradesResponse struct {
	Grades []*Grade `json:"grades"`
}

type UserResourcesResponse struct {
	Uploads []*Upload `json:"uploads"`
	model.Pagination
}

type AcademicLearningResourcesResponse struct {
	Uploads []*Upload `json:"uploads"`
	model.Pagination
}

type ThirdPartResourcesResponse struct {
	Uploads []*Upload `json:"uploads"`
	model.Pagination
}

type OtherVideoResourcesResponse struct {
	Uploads []*Upload `json:"uploads"`
	model.Pagination
}

type RecentlyVisitedCoursesResponse struct {
	Courses []*Course `json:"courses,omitempty"`
}

type NotesResponse []*Note

type SignInStatsResponse = SignIn

type FailedCoursesResponse struct {
	Courses []*Course `json:"courses,omitempty"`
}

type MyCapturesResponse struct {
	Items []*Capture `json:"items"`
	model.Pagination
}

type PublicCapturesResponse struct {
	Items []*Capture `json:"items"`
	model.Pagination
}

type CoursesIdentitiesResponse struct {
	CoursesIdentities []*Course `json:"courses_identities"`
}
