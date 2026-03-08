package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"net/url"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

// Service handles organization and admin-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Organization ---

// GetGlobalConfig returns the organization's global configuration.
func (s *Service) GetGlobalConfig(ctx context.Context) (*GlobalConfig, error) {
	result := new(GlobalConfig)
	_, err := s.client.Get(ctx, "/org/global-config", result)
	return result, err
}

// GetConfig returns the file format configuration.
func (s *Service) GetConfig(ctx context.Context) (*Config, error) {
	result := new(Config)
	_, err := s.client.Get(ctx, "/api/config?no-intercept=true", result)
	return result, err
}

// GetLangSettings returns language settings for an organization.
func (s *Service) GetLangSettings(ctx context.Context, orgID int) (*LangSettingsResponse, error) {
	u := fmt.Sprintf("/api/orgs/%d/lang-settings", orgID)
	result := new(LangSettingsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetOrg returns organization information.
func (s *Service) GetOrg(ctx context.Context, orgID int) (*model.OrgDetail, error) {
	u := fmt.Sprintf("/api/orgs/%d", orgID)
	result := new(model.OrgDetail)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetCurrentOrg returns the current organization information.
func (s *Service) GetCurrentOrg(ctx context.Context) (*model.OrgDetail, error) {
	result := new(model.OrgDetail)
	_, err := s.client.Get(ctx, "/api/org", result)
	return result, err
}

// GetOrgDepartments returns departments for an organization.
func (s *Service) GetOrgDepartments(ctx context.Context, orgID int) (*DepartmentsResponse, error) {
	u := fmt.Sprintf("/api/orgs/%d/departments", orgID)
	result := new(DepartmentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetOrgAssociatedPartnerTypes returns associated partner types for an organization.
func (s *Service) GetOrgAssociatedPartnerTypes(ctx context.Context, orgID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/orgs/%d/associated-partner-types", orgID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListAllOrgs returns all organizations.
func (s *Service) ListAllOrgs(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/all-orgs", &result)
	return result, err
}

// ListOrgsByIDs returns organizations by repeated id query parameters.
func (s *Service) ListOrgsByIDs(ctx context.Context, orgIDs []int) (json.RawMessage, error) {
	values := url.Values{}
	for _, id := range orgIDs {
		values.Add("id", fmt.Sprintf("%d", id))
	}
	u := "/api/orgs"
	if encoded := values.Encode(); encoded != "" {
		u += "?" + encoded
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetPortalLogo returns the portal logo.
func (s *Service) GetPortalLogo(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/portal-logo", &result)
	return result, err
}

// --- Departments ---

// ListDepartments returns all departments.
func (s *Service) ListDepartments(ctx context.Context, params map[string]string) (*DepartmentsResponse, error) {
	u := addQueryParams("/api/departments", params)
	result := new(DepartmentsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetDepartment returns a specific department.
func (s *Service) GetDepartment(ctx context.Context, deptID int) (*model.Department, error) {
	u := fmt.Sprintf("/api/departments/%d", deptID)
	result := new(model.Department)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListDepartmentsForUser returns departments for the current user.
func (s *Service) ListDepartmentsForUser(ctx context.Context) (*DepartmentsResponse, error) {
	result := new(DepartmentsResponse)
	_, err := s.client.Get(ctx, "/api/departments-for-user", result)
	return result, err
}

// GetSourceDepartmentCodeForUser returns the source department code.
func (s *Service) GetSourceDepartmentCodeForUser(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/source-department-code-for-user", &result)
	return result, err
}

// ListTopDepartments returns top-level departments.
func (s *Service) ListTopDepartments(ctx context.Context, fields string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/top-departments?fields=%s", fields)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Academic Years & Semesters ---

// ListAcademicYears returns all academic years.
func (s *Service) ListAcademicYears(ctx context.Context) (*AcademicYearsResponse, error) {
	result := new(AcademicYearsResponse)
	_, err := s.client.Get(ctx, "/api/academic-years?fields=id,name,sort,is_active", result)
	return result, err
}

// ListSemesters returns all semesters.
func (s *Service) ListSemesters(ctx context.Context) (*SemestersResponse, error) {
	result := new(SemestersResponse)
	_, err := s.client.Get(ctx, "/api/semesters", result)
	return result, err
}

// ListClasses returns all classes.
func (s *Service) ListClasses(ctx context.Context, params map[string]string) (*ClassesResponse, error) {
	u := addQueryParams("/api/classes", params)
	result := new(ClassesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListGrades returns grades.
func (s *Service) ListGrades(ctx context.Context, params map[string]string) (*GradesResponse, error) {
	u := addQueryParams("/api/grades", params)
	result := new(GradesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Outline Settings ---

// GetOutlineSetting returns outline setting options.
func (s *Service) GetOutlineSetting(ctx context.Context) (*OutlineSettingResponse, error) {
	result := new(OutlineSettingResponse)
	_, err := s.client.Get(ctx, "/api/outline-setting", result)
	return result, err
}

// UpdateOutlineSetting updates outline settings.
func (s *Service) UpdateOutlineSetting(ctx context.Context, body interface{}) error {
	_, err := s.client.Put(ctx, "/api/outline-setting", body, nil)
	return err
}

// ToggleOutlineSetting toggles outline setting.
func (s *Service) ToggleOutlineSetting(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/outline-setting/toggle", body, nil)
	return err
}

// NotifyOutline sends outline notification.
func (s *Service) NotifyOutline(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/outline/notify", body, nil)
	return err
}

// --- Users Management ---

// ListUsers returns users with pagination.
func (s *Service) ListUsers(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/users", opts)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListUsersForManagement returns users for management views with filters.
func (s *Service) ListUsersForManagement(ctx context.Context, opts *model.ListOptions, body interface{}) (json.RawMessage, error) {
	return s.ListUsersForManagementWithParams(ctx, opts, body, ListUsersForManagementParams{
		ForManagement:   true,
		NeedAIActivated: true,
	})
}

// ListUsersForManagementWithParams returns users for management views with filters and query options.
func (s *Service) ListUsersForManagementWithParams(ctx context.Context, opts *model.ListOptions, body interface{}, params ListUsersForManagementParams) (json.RawMessage, error) {
	u := addListOptions("/api/users", opts)
	query := map[string]string{}
	if params.ForManagement {
		query["for_management"] = "true"
	}
	if params.NeedAIActivated {
		query["need_ai_activated"] = "true"
	}
	if params.IgnoreAvatar {
		query["ignore_avatar"] = "true"
	}
	u = addQueryParams(u, query)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// GetUserByType returns users by type.
func (s *Service) GetUserByType(ctx context.Context, userType, responseKey string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/user?type=%s&response_key=%s", userType, responseKey)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListInstructorUsers returns instructor users using the frontend endpoint shape.
func (s *Service) ListInstructorUsers(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/user?type=instructor", &result)
	return result, err
}

// --- Enrollments Management ---

// CreateEnrollment creates an enrollment.
func (s *Service) CreateEnrollment(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/enrollments", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// DeleteEnrollment deletes an enrollment.
func (s *Service) DeleteEnrollment(ctx context.Context, enrollmentID int) error {
	u := fmt.Sprintf("/api/course/enrollments/%d", enrollmentID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ChangeEnrollmentRole updates the role of an enrollment using the frontend endpoint shape.
func (s *Service) ChangeEnrollmentRole(ctx context.Context, enrollmentID int, body ChangeEnrollmentRoleRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/enrollments/%d", enrollmentID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// --- Invites ---

// CreateInvite creates an enrollment invite.
func (s *Service) CreateInvite(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/invites", body, &result)
	return result, err
}

// GetInvite returns an invite.
func (s *Service) GetInvite(ctx context.Context, inviteID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/invites/%d", inviteID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Data Import ---

// ImportFromWord imports data from a Word document.
func (s *Service) ImportFromWord(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/data-import/from-word", body, &result)
	return result, err
}

// AIConvert converts data using AI.
func (s *Service) AIConvert(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/data-import/ai-convert", body, &result)
	return result, err
}

// ImportEnrollments imports enrollments.
func (s *Service) ImportEnrollments(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/data-import/enrollments", body, &result)
	return result, err
}

// ImportScores imports scores.
func (s *Service) ImportScores(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/scores/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ValidateImport validates an import.
func (s *Service) ValidateImport(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/data-import/validation", body, &result)
	return result, err
}

// ImportCourses imports courses.
func (s *Service) ImportCourses(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/data-import/courses", body, &result)
	return result, err
}

// ImportEditCourses imports edited course data.
func (s *Service) ImportEditCourses(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/data-import/edit-courses", body, &result)
	return result, err
}

// ImportEnrollmentsForCourse imports enrollments for a specific course.
func (s *Service) ImportEnrollmentsForCourse(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/enrollments/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ImportCourseGroups imports course groups.
func (s *Service) ImportCourseGroups(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/data-import/course-groups", body, &result)
	return result, err
}

// ImportItemScores imports item scores for a course.
func (s *Service) ImportItemScores(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/item_scores/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ImportSeatNumbers imports seat numbers for a course.
func (s *Service) ImportSeatNumbers(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/seat-number/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ImportCourseRollcalls imports rollcall data for a course.
func (s *Service) ImportCourseRollcalls(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/course/%d/rollcall", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ImportChaoxingScore imports Chaoxing scores for a course.
func (s *Service) ImportChaoxingScore(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/chaoxing-score/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ImportSubjectLibFromZip imports subject library data from a ZIP file.
func (s *Service) ImportSubjectLibFromZip(ctx context.Context, subjectLibID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/subject-libs/%d/from-zip", subjectLibID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ImportQuestionnaireSubjectLibFromZip imports questionnaire subject library data from a ZIP file.
func (s *Service) ImportQuestionnaireSubjectLibFromZip(ctx context.Context, subjectLibID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/subject-libs/%d/from-zip/questionnaire", subjectLibID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ImportExamSubjectsFromZip imports exam subjects from a ZIP file.
func (s *Service) ImportExamSubjectsFromZip(ctx context.Context, examID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/exams/%d/subjects/from-zip", examID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ImportClassroomExamSubjectsFromZip imports classroom exam subjects from a ZIP file.
func (s *Service) ImportClassroomExamSubjectsFromZip(ctx context.Context, classroomExamID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/data-import/classroom-exams/%d/subjects/from-zip", classroomExamID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// --- Course Copy ---

// CopyCourse copies a course.
func (s *Service) CopyCourse(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/course-copy/copy", body, &result)
	return result, err
}

// ListCopyableCourses returns courses that can be copied.
func (s *Service) ListCopyableCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/course-copy/courses", &result)
	return result, err
}

// ImportMBZ imports a Moodle backup file.
func (s *Service) ImportMBZ(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/course/mbz/import", body, &result)
	return result, err
}

// --- Authorization ---

// GetCourseRoles returns available course roles.
func (s *Service) GetCourseRoles(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/authz/course-roles", &result)
	return result, err
}

// GetRoles returns available roles.
func (s *Service) GetRoles(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/authz/roles", &result)
	return result, err
}

// GetCourseAuthz returns authorization for a course.
func (s *Service) GetCourseAuthz(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/authz/courses/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Auth Code ---

// GetAuthCode generates an auth code.
func (s *Service) GetAuthCode(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/auth_code/get_auth_code", &result)
	return result, err
}

// ValidateAuthCode validates an auth code.
func (s *Service) ValidateAuthCode(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/auth_code/validate_auth_code", body, &result)
	return result, err
}

// AuthValidate validates authentication.
func (s *Service) AuthValidate(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/auth_code/auth_validate", body, &result)
	return result, err
}

// --- Plans ---

// GetCurrentPlan returns the current plan.
func (s *Service) GetCurrentPlan(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/plan/request/current", &result)
	return result, err
}

// GetChangePlanList returns the change plan list.
func (s *Service) GetChangePlanList(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/org/change-plan-list", &result)
	return result, err
}

// RequestChangePlan requests a plan change.
func (s *Service) RequestChangePlan(ctx context.Context, planID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/org/request/change-plan/%d", planID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// RequestPlanPlus requests a plan plus.
func (s *Service) RequestPlanPlus(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/org/request/plan-plus", body, &result)
	return result, err
}

// ListWGAdminOrgRequests returns working-group admin organization requests.
func (s *Service) ListWGAdminOrgRequests(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/wg-admin/orgs/requests", params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// UpdateWGAdminOrgRequest updates a working-group admin organization request.
func (s *Service) UpdateWGAdminOrgRequest(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Put(ctx, "/api/wg-admin/orgs/request", body, &result)
	return result, err
}

// --- Face Recognition ---

// GetFaceRecognitionConfig returns face recognition configuration.
func (s *Service) GetFaceRecognitionConfig(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/face-recognition/config", &result)
	return result, err
}

// --- Third-party Auth ---

// CheckGoogleAuth checks if Google auth is available.
func (s *Service) CheckGoogleAuth(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/google/auth/available", &result)
	return result, err
}

// CheckMicrosoftAuth checks if Microsoft auth is available.
func (s *Service) CheckMicrosoftAuth(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/microsoft/auth/available", &result)
	return result, err
}

// CheckLarkAuthorization checks Lark authorization.
func (s *Service) CheckLarkAuthorization(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/lark/authorization/check", &result)
	return result, err
}

// --- Jobs ---

// GetJob returns a background job status.
func (s *Service) GetJob(ctx context.Context, jobID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/jobs/%d", jobID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetLastTask returns the last background task.
func (s *Service) GetLastTask(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/task/last?no-intercept=true", &result)
	return result, err
}

// --- CSP ---

// DetectCSP detects CSP for a URL.
func (s *Service) DetectCSP(ctx context.Context, targetURL string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/csp/detect?url=%s", targetURL)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- QR Code ---

// GetQRCode generates a QR code for a URL.
func (s *Service) GetQRCode(ctx context.Context, targetURL string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/qrcode?url=%s", targetURL)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}
