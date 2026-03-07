package admin

import "github.com/eWloYW8/zju-courses-go-sdk/model"

type AcademicYearsResponse struct {
	AcademicYears []*model.AcademicYear `json:"academic_years"`
}

type SemestersResponse struct {
	Semesters []*model.Semester `json:"semesters"`
}

type DepartmentsResponse struct {
	Departments []*model.Department `json:"departments"`
}

type ClassesResponse struct {
	Classes []*model.Class `json:"classes"`
}

type GradesResponse struct {
	Grades []*model.Grade `json:"grades"`
}
