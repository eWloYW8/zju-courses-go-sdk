package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// AdminService handles organization and admin-related API operations.
type AdminService struct {
	client *Client
}

// --- Organization ---

// GetGlobalConfig returns the organization's global configuration.
func (s *AdminService) GetGlobalConfig(ctx context.Context) (*model.GlobalConfig, error) {
	result := new(model.GlobalConfig)
	_, err := s.client.get(ctx, "/org/global-config", result)
	return result, err
}

// GetConfig returns the file format configuration.
func (s *AdminService) GetConfig(ctx context.Context) (*model.Config, error) {
	result := new(model.Config)
	_, err := s.client.get(ctx, "/api/config?no-intercept=true", result)
	return result, err
}

// GetLangSettings returns language settings for an organization.
func (s *AdminService) GetLangSettings(ctx context.Context, orgID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/orgs/%d/lang-settings", orgID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetOrg returns organization information.
func (s *AdminService) GetOrg(ctx context.Context, orgID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/orgs/%d", orgID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListAllOrgs returns all organizations.
func (s *AdminService) ListAllOrgs(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/all-orgs", &result)
	return result, err
}

// GetPortalLogo returns the portal logo.
func (s *AdminService) GetPortalLogo(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/portal-logo", &result)
	return result, err
}

// --- Departments ---

// ListDepartments returns all departments.
func (s *AdminService) ListDepartments(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/departments", params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetDepartment returns a specific department.
func (s *AdminService) GetDepartment(ctx context.Context, deptID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/departments/%d", deptID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListDepartmentsForUser returns departments for the current user.
func (s *AdminService) ListDepartmentsForUser(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/departments-for-user", &result)
	return result, err
}

// GetSourceDepartmentCodeForUser returns the source department code.
func (s *AdminService) GetSourceDepartmentCodeForUser(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/source-department-code-for-user", &result)
	return result, err
}

// ListTopDepartments returns top-level departments.
func (s *AdminService) ListTopDepartments(ctx context.Context, fields string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/top-departments?fields=%s", fields)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Academic Years & Semesters ---

// ListAcademicYears returns all academic years.
func (s *AdminService) ListAcademicYears(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/academic-years?fields=id,name,sort,is_active", &result)
	return result, err
}

// ListSemesters returns all semesters.
func (s *AdminService) ListSemesters(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/semesters", &result)
	return result, err
}

// ListClasses returns all classes.
func (s *AdminService) ListClasses(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/classes", params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListGrades returns grades.
func (s *AdminService) ListGrades(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/grades", params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Outline Settings ---

// GetOutlineSetting returns outline setting options.
func (s *AdminService) GetOutlineSetting(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/outline-setting", &result)
	return result, err
}

// UpdateOutlineSetting updates outline settings.
func (s *AdminService) UpdateOutlineSetting(ctx context.Context, body interface{}) error {
	_, err := s.client.put(ctx, "/api/outline-setting", body, nil)
	return err
}

// ToggleOutlineSetting toggles outline setting.
func (s *AdminService) ToggleOutlineSetting(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/outline-setting/toggle", body, nil)
	return err
}

// NotifyOutline sends outline notification.
func (s *AdminService) NotifyOutline(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/outline/notify", body, nil)
	return err
}

// --- Users Management ---

// ListUsers returns users with pagination.
func (s *AdminService) ListUsers(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/users", opts)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetUserByType returns users by type.
func (s *AdminService) GetUserByType(ctx context.Context, userType, responseKey string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/user?type=%s&response_key=%s", userType, responseKey)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Enrollments Management ---

// CreateEnrollment creates an enrollment.
func (s *AdminService) CreateEnrollment(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/enrollments/%d", courseID)
	var result json.RawMessage
	_, err := s.client.post(ctx, u, body, &result)
	return result, err
}

// DeleteEnrollment deletes an enrollment.
func (s *AdminService) DeleteEnrollment(ctx context.Context, enrollmentID int) error {
	u := fmt.Sprintf("/api/enrollments/%d", enrollmentID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// --- Invites ---

// CreateInvite creates an enrollment invite.
func (s *AdminService) CreateInvite(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/invites", body, &result)
	return result, err
}

// GetInvite returns an invite.
func (s *AdminService) GetInvite(ctx context.Context, inviteID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/invites/%d", inviteID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Data Import ---

// ImportFromWord imports data from a Word document.
func (s *AdminService) ImportFromWord(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/data-import/from-word", body, &result)
	return result, err
}

// AIConvert converts data using AI.
func (s *AdminService) AIConvert(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/data-import/ai-convert", body, &result)
	return result, err
}

// ImportEnrollments imports enrollments.
func (s *AdminService) ImportEnrollments(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/data-import/enrollments", body, &result)
	return result, err
}

// ImportScores imports scores.
func (s *AdminService) ImportScores(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/scores/%d", courseID)
	var result json.RawMessage
	_, err := s.client.post(ctx, u, body, &result)
	return result, err
}

// ValidateImport validates an import.
func (s *AdminService) ValidateImport(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/data-import/validation", body, &result)
	return result, err
}

// ImportCourses imports courses.
func (s *AdminService) ImportCourses(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/data-import/courses", body, &result)
	return result, err
}

// --- Course Copy ---

// CopyCourse copies a course.
func (s *AdminService) CopyCourse(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/course-copy/copy", body, &result)
	return result, err
}

// ListCopyableCourses returns courses that can be copied.
func (s *AdminService) ListCopyableCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/course-copy/courses", &result)
	return result, err
}

// ImportMBZ imports a Moodle backup file.
func (s *AdminService) ImportMBZ(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/course/mbz/import", body, &result)
	return result, err
}

// --- Authorization ---

// GetCourseRoles returns available course roles.
func (s *AdminService) GetCourseRoles(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/authz/course-roles", &result)
	return result, err
}

// GetRoles returns available roles.
func (s *AdminService) GetRoles(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/authz/roles", &result)
	return result, err
}

// GetCourseAuthz returns authorization for a course.
func (s *AdminService) GetCourseAuthz(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/authz/courses/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Auth Code ---

// GetAuthCode generates an auth code.
func (s *AdminService) GetAuthCode(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/auth_code/get_auth_code", &result)
	return result, err
}

// ValidateAuthCode validates an auth code.
func (s *AdminService) ValidateAuthCode(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/auth_code/validate_auth_code", body, &result)
	return result, err
}

// AuthValidate validates authentication.
func (s *AdminService) AuthValidate(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/auth_code/auth_validate", body, &result)
	return result, err
}

// --- Plans ---

// GetCurrentPlan returns the current plan.
func (s *AdminService) GetCurrentPlan(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/plan/request/current", &result)
	return result, err
}

// GetChangePlanList returns the change plan list.
func (s *AdminService) GetChangePlanList(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/org/change-plan-list", &result)
	return result, err
}

// RequestChangePlan requests a plan change.
func (s *AdminService) RequestChangePlan(ctx context.Context, planID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/org/request/change-plan/%d", planID)
	var result json.RawMessage
	_, err := s.client.post(ctx, u, body, &result)
	return result, err
}

// RequestPlanPlus requests a plan plus.
func (s *AdminService) RequestPlanPlus(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/org/request/plan-plus", body, &result)
	return result, err
}

// --- Face Recognition ---

// GetFaceRecognitionConfig returns face recognition configuration.
func (s *AdminService) GetFaceRecognitionConfig(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/face-recognition/config", &result)
	return result, err
}

// --- Third-party Auth ---

// CheckGoogleAuth checks if Google auth is available.
func (s *AdminService) CheckGoogleAuth(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/google/auth/available", &result)
	return result, err
}

// CheckMicrosoftAuth checks if Microsoft auth is available.
func (s *AdminService) CheckMicrosoftAuth(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/microsoft/auth/available", &result)
	return result, err
}

// CheckLarkAuthorization checks Lark authorization.
func (s *AdminService) CheckLarkAuthorization(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/lark/authorization/check", &result)
	return result, err
}

// --- Jobs ---

// GetJob returns a background job status.
func (s *AdminService) GetJob(ctx context.Context, jobID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/jobs/%d", jobID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetLastTask returns the last background task.
func (s *AdminService) GetLastTask(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/task/last?no-intercept=true", &result)
	return result, err
}

// --- CSP ---

// DetectCSP detects CSP for a URL.
func (s *AdminService) DetectCSP(ctx context.Context, targetURL string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/csp/detect?url=%s", targetURL)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- QR Code ---

// GetQRCode generates a QR code for a URL.
func (s *AdminService) GetQRCode(ctx context.Context, targetURL string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/qrcode?url=%s", targetURL)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}
