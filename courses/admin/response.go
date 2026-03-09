package admin

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type AcademicYearsResponse struct {
	AcademicYears []*model.AcademicYear `json:"academic_years"`
}

type SemestersResponse struct {
	Semesters []*model.Semester `json:"semesters"`
}

type DepartmentsResponse struct {
	Departments []*model.Department `json:"departments"`
}

type SourceDepartmentCodeResponse struct {
	DepartmentCode string `json:"department_code,omitempty"`
}

type ClassesResponse struct {
	Classes []*model.Class `json:"classes"`
}

type GradesResponse struct {
	Grades []*model.Grade `json:"grades"`
}

type AssistantRolesResponse []*AssistantRole

type JobStatusResponse struct {
	Status string         `json:"status,omitempty"`
	Result map[string]any `json:"result,omitempty"`
}
