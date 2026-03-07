package statistics

type StatisticsQueryParams map[string]string

type UserCompletenessQueryParams map[string]string

type DepartmentUpdateRequest struct {
	DepartmentID int `json:"department_id"`
}
