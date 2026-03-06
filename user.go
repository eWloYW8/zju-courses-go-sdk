package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// UserService handles user-related API operations.
type UserService struct {
	client *Client
}

// --- Response Types ---

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
	Classes []interface{} `json:"classes"`
}

type GradesResponse struct {
	Grades []*model.Grade `json:"grades"`
}

type UserResourcesResponse struct {
	Uploads []*model.Upload `json:"uploads"`
	model.Pagination
}

type RecentlyVisitedCoursesResponse struct {
	Courses []*model.Course `json:"courses,omitempty"`
}

// --- User Profile ---

// GetProfile returns the current user's profile.
func (s *UserService) GetProfile(ctx context.Context) (*model.UserProfile, error) {
	result := new(model.UserProfile)
	_, err := s.client.get(ctx, "/api/user", result)
	return result, err
}

// UpdateProfile updates the current user's profile.
func (s *UserService) UpdateProfile(ctx context.Context, body interface{}) (*model.UserProfile, error) {
	result := new(model.UserProfile)
	_, err := s.client.put(ctx, "/api/user", body, result)
	return result, err
}

// SearchUser searches for users.
func (s *UserService) SearchUser(ctx context.Context, params map[string]string) ([]interface{}, error) {
	u := addQueryParams("/api/user/search", params)
	var result []interface{}
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetUserByID returns a user by their ID.
func (s *UserService) GetUserByID(ctx context.Context, userID int) (*model.UserProfile, error) {
	u := fmt.Sprintf("/api/users/%d", userID)
	result := new(model.UserProfile)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

type StorageUsedResponse struct {
	StorageUsed     int64 `json:"storage_used"`
	StorageAssigned int64 `json:"storage_assigned"`
}

// --- User Resources ---

// ListResources returns the current user's uploaded resources.
func (s *UserService) ListResources(ctx context.Context, opts *model.ListOptions) (*UserResourcesResponse, error) {
	u := addListOptions("/api/user/resources", opts)
	result := new(UserResourcesResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// GetStorageUsed returns the user's storage usage.
func (s *UserService) GetStorageUsed(ctx context.Context) (*StorageUsedResponse, error) {
	result := new(StorageUsedResponse)
	_, err := s.client.get(ctx, "/api/user/storage-used", result)
	return result, err
}

// --- Academic Years & Semesters ---

// ListMyAcademicYears returns the user's academic years.
func (s *UserService) ListMyAcademicYears(ctx context.Context) (*AcademicYearsResponse, error) {
	result := new(AcademicYearsResponse)
	_, err := s.client.get(ctx, "/api/my-academic-years", result)
	return result, err
}

// ListMySemesters returns the user's semesters.
func (s *UserService) ListMySemesters(ctx context.Context) (*SemestersResponse, error) {
	result := new(SemestersResponse)
	_, err := s.client.get(ctx, "/api/my-semesters", result)
	return result, err
}

// ListMyAllSemesters returns all semesters for the user.
func (s *UserService) ListMyAllSemesters(ctx context.Context) (*SemestersResponse, error) {
	result := new(SemestersResponse)
	_, err := s.client.get(ctx, "/api/my-semesters-all", result)
	return result, err
}

// ListMyDepartments returns the user's departments.
func (s *UserService) ListMyDepartments(ctx context.Context) (*DepartmentsResponse, error) {
	result := new(DepartmentsResponse)
	_, err := s.client.get(ctx, "/api/my-departments", result)
	return result, err
}

// ListMyClasses returns the user's classes.
func (s *UserService) ListMyClasses(ctx context.Context) (*ClassesResponse, error) {
	result := new(ClassesResponse)
	_, err := s.client.get(ctx, "/api/my-classes", result)
	return result, err
}

// ListMyGrades returns the user's grades.
func (s *UserService) ListMyGrades(ctx context.Context) (*GradesResponse, error) {
	result := new(GradesResponse)
	_, err := s.client.get(ctx, "/api/my-grades", result)
	return result, err
}

// --- Recently Visited ---

// GetRecentlyVisitedCourses returns recently visited courses.
func (s *UserService) GetRecentlyVisitedCourses(ctx context.Context) (*RecentlyVisitedCoursesResponse, error) {
	result := new(RecentlyVisitedCoursesResponse)
	_, err := s.client.get(ctx, "/api/user/recently-visited-courses", result)
	return result, err
}

// --- User Preferences ---

// SetLanguage sets the user's language preference.
func (s *UserService) SetLanguage(ctx context.Context, lang string) error {
	_, err := s.client.put(ctx, "/api/user/language", map[string]string{"language": lang}, nil)
	return err
}

// GetLinks returns user's custom links.
func (s *UserService) GetLinks(ctx context.Context) ([]interface{}, error) {
	var result []interface{}
	_, err := s.client.get(ctx, "/api/user/links", &result)
	return result, err
}

// CreateLink creates a custom link.
func (s *UserService) CreateLink(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/user/links", body, &result)
	return result, err
}

// UpdateLink updates a custom link.
func (s *UserService) UpdateLink(ctx context.Context, linkID int, body interface{}) error {
	u := fmt.Sprintf("/api/user/links/%d", linkID)
	_, err := s.client.put(ctx, u, body, nil)
	return err
}

// DeleteLink deletes a custom link.
func (s *UserService) DeleteLink(ctx context.Context, linkID int) error {
	u := fmt.Sprintf("/api/user/links/%d", linkID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// --- User Course Operations ---

// CheckCourseGraduate checks if a course is graduated.
func (s *UserService) CheckCourseGraduate(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/user/check-course-graduate?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// CheckExpiredPassword checks if the user's password is expired.
func (s *UserService) CheckExpiredPassword(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/check-expired-password", &result)
	return result, err
}

// GetFirstTimeLogin returns if this is the user's first login.
func (s *UserService) GetFirstTimeLogin(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/first-time-login", &result)
	return result, err
}

// GetPreTask returns the user's pre-task information.
func (s *UserService) GetPreTask(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/pre-task", &result)
	return result, err
}

// GetPersonas returns the user's personas.
func (s *UserService) GetPersonas(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/personas", &result)
	return result, err
}

// GetDepartment returns the user's department.
func (s *UserService) GetDepartment(ctx context.Context) (*model.Department, error) {
	result := new(model.Department)
	_, err := s.client.get(ctx, "/api/user/department", result)
	return result, err
}

// GetAssociationCode returns the user's association code.
func (s *UserService) GetAssociationCode(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/association-code", &result)
	return result, err
}

// GetChat returns the user's chat information.
func (s *UserService) GetChat(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/chat", &result)
	return result, err
}

// GetFailedCourses returns the user's failed courses.
func (s *UserService) GetFailedCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/failed-courses", &result)
	return result, err
}

// GetCourseCertificationScores returns course certification scores.
func (s *UserService) GetCourseCertificationScores(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/course-certification/scores", &result)
	return result, err
}

// GetAcademicLearningResources returns academic learning resources.
func (s *UserService) GetAcademicLearningResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/academic-learning-resources", &result)
	return result, err
}

// GetThirdPartResources returns third-party resources.
func (s *UserService) GetThirdPartResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/third-part-resources", &result)
	return result, err
}

// GetOtherVideoResources returns other video resources.
func (s *UserService) GetOtherVideoResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user/other-video-resources", &result)
	return result, err
}

// --- Notes ---

// ListNotes returns the user's notes.
func (s *UserService) ListNotes(ctx context.Context, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams("/api/notes", params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// CreateNote creates a note.
func (s *UserService) CreateNote(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/notes", body, &result)
	return result, err
}

// UpdateNote updates a note.
func (s *UserService) UpdateNote(ctx context.Context, noteID int, body interface{}) error {
	u := fmt.Sprintf("/api/notes/%d", noteID)
	_, err := s.client.put(ctx, u, body, nil)
	return err
}

// DeleteNote deletes a note.
func (s *UserService) DeleteNote(ctx context.Context, noteID int) error {
	u := fmt.Sprintf("/api/notes/%d", noteID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// --- Notebooks ---

// GetNotebook returns a notebook.
func (s *UserService) GetNotebook(ctx context.Context, notebookID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/notebooks/%d", notebookID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Sign In ---

// SignIn performs a sign-in to the platform.
func (s *UserService) SignIn(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/sign-in", body, &result)
	return result, err
}

// GetSignInStats returns sign-in statistics.
func (s *UserService) GetSignInStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/sign-in/stats", &result)
	return result, err
}

// --- Email Verification ---

// SendVerificationEmail sends a verification email.
func (s *UserService) SendVerificationEmail(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/user/send_verification_email", body, nil)
	return err
}

// SendOrgSignUpVerificationEmail sends an org sign-up verification email.
func (s *UserService) SendOrgSignUpVerificationEmail(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/user/send_org_sing_up_verification_email", body, nil)
	return err
}

// --- Captures ---

// ListMyCaptures returns the user's captures.
func (s *UserService) ListMyCaptures(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/my-captures", &result)
	return result, err
}

// ListPublicCaptures returns public captures.
func (s *UserService) ListPublicCaptures(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/public-captures", &result)
	return result, err
}

// GetCapture returns a specific capture.
func (s *UserService) GetCapture(ctx context.Context, captureID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/captures/%d", captureID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- User Index Stats ---

// GetCoursesInfoStatus returns course info status for the user index.
func (s *UserService) GetCoursesInfoStatus(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user-index-stat/courses/info-status", &result)
	return result, err
}

// GetOrgSummary returns the organization summary for the user index.
func (s *UserService) GetOrgSummary(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/user-index-stat/org-summary", &result)
	return result, err
}

// GetCoursesIdentities returns course identities.
func (s *UserService) GetCoursesIdentities(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/courses-identities?no-intercept=true", &result)
	return result, err
}
