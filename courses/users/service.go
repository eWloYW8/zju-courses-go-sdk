package users

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

// Service handles user-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- User Profile ---

// GetProfile returns the current user's profile.
func (s *Service) GetProfile(ctx context.Context) (*UserProfile, error) {
	result := new(UserProfile)
	_, err := s.client.Get(ctx, "/api/user", result)
	return result, err
}

// UpdateProfile updates the current user's profile.
func (s *Service) UpdateProfile(ctx context.Context, body UpdateProfileRequest) (*UserProfile, error) {
	result := new(UserProfile)
	_, err := s.client.Put(ctx, "/api/user", body, result)
	return result, err
}

// SearchUser searches for users.
func (s *Service) SearchUser(ctx context.Context, params SearchUserParams) ([]*UserSearchResult, error) {
	query := map[string]string{}
	if params.Keywords != "" {
		query["keywords"] = params.Keywords
	}
	if params.ExcludeStudentRole {
		query["exclude_student_role"] = "true"
	}
	if params.OrgID != nil {
		query["org_id"] = fmt.Sprintf("%d", *params.OrgID)
	}
	if params.DepartmentID != nil {
		query["department_id"] = fmt.Sprintf("%d", *params.DepartmentID)
	}
	var result []*UserSearchResult
	_, err := s.client.Get(ctx, addQueryParams("/api/user/search", query), &result)
	return result, err
}

// GetUserByID returns a user by their ID.
func (s *Service) GetUserByID(ctx context.Context, userID int) (*UserProfile, error) {
	u := fmt.Sprintf("/api/users/%d", userID)
	result := new(UserProfile)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- User Resources ---

// ListResources returns the current user's uploaded resources.
func (s *Service) ListResources(ctx context.Context, opts *model.ListOptions) (*UserResourcesResponse, error) {
	u := addListOptions("/api/user/resources", opts)
	result := new(UserResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetStorageUsed returns the user's storage usage.
func (s *Service) GetStorageUsed(ctx context.Context) (*StorageUsedResponse, error) {
	result := new(StorageUsedResponse)
	_, err := s.client.Get(ctx, "/api/user/storage-used", result)
	return result, err
}

// --- Academic Years & Semesters ---

// ListMyAcademicYears returns the user's academic years.
func (s *Service) ListMyAcademicYears(ctx context.Context) (*AcademicYearsResponse, error) {
	result := new(AcademicYearsResponse)
	_, err := s.client.Get(ctx, "/api/my-academic-years", result)
	return result, err
}

// ListMyCurriculumAcademicYears returns the user's curriculum academic years.
func (s *Service) ListMyCurriculumAcademicYears(ctx context.Context) (*AcademicYearsResponse, error) {
	result := new(AcademicYearsResponse)
	_, err := s.client.Get(ctx, "/api/my-curriculum-academic-years?fields=id,name,sort,is_active", result)
	return result, err
}

// ListMySemesters returns the user's semesters.
func (s *Service) ListMySemesters(ctx context.Context) (*SemestersResponse, error) {
	result := new(SemestersResponse)
	_, err := s.client.Get(ctx, "/api/my-semesters", result)
	return result, err
}

// ListMyCurriculumSemesters returns the user's curriculum semesters.
func (s *Service) ListMyCurriculumSemesters(ctx context.Context, academicYearIDs []int) (*SemestersResponse, error) {
	params := map[string]string{}
	for i, id := range academicYearIDs {
		params[fmt.Sprintf("academic_year_ids[%d]", i)] = fmt.Sprintf("%d", id)
	}
	u := addQueryParams("/api/my-curriculum-semesters", params)
	result := new(SemestersResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListMyAllSemesters returns all semesters for the user.
func (s *Service) ListMyAllSemesters(ctx context.Context) (*SemestersResponse, error) {
	result := new(SemestersResponse)
	_, err := s.client.Get(ctx, "/api/my-semesters-all", result)
	return result, err
}

// ListMyDepartments returns the user's departments.
func (s *Service) ListMyDepartments(ctx context.Context) (*DepartmentsResponse, error) {
	result := new(DepartmentsResponse)
	_, err := s.client.Get(ctx, "/api/my-departments", result)
	return result, err
}

// ListMyClasses returns the user's classes.
func (s *Service) ListMyClasses(ctx context.Context) (*ClassesResponse, error) {
	result := new(ClassesResponse)
	_, err := s.client.Get(ctx, "/api/my-classes", result)
	return result, err
}

// GetUserClasses returns classes associated with the current user.
func (s *Service) GetUserClasses(ctx context.Context) (*ClassesResponse, error) {
	result := new(ClassesResponse)
	_, err := s.client.Get(ctx, "/api/user/classes", result)
	return result, err
}

// ListMyGrades returns the user's grades.
func (s *Service) ListMyGrades(ctx context.Context) (*GradesResponse, error) {
	result := new(GradesResponse)
	_, err := s.client.Get(ctx, "/api/my-grades", result)
	return result, err
}

// --- Recently Visited ---

// GetRecentlyVisitedCourses returns recently visited courses.
func (s *Service) GetRecentlyVisitedCourses(ctx context.Context) (*RecentlyVisitedCoursesResponse, error) {
	result := new(RecentlyVisitedCoursesResponse)
	_, err := s.client.Get(ctx, "/api/user/recently-visited-courses", result)
	return result, err
}

// --- User Preferences ---

// SetLanguage sets the user's language preference.
func (s *Service) SetLanguage(ctx context.Context, lang string) error {
	_, err := s.client.Put(ctx, "/api/user/language", map[string]string{"language": lang}, nil)
	return err
}

// GetLinks returns user's custom links.
func (s *Service) GetLinks(ctx context.Context) ([]*UserLink, error) {
	var result []*UserLink
	_, err := s.client.Get(ctx, "/api/user/links", &result)
	return result, err
}

// CreateLink creates a custom link.
func (s *Service) CreateLink(ctx context.Context, body UserLinkRequest) (*UserLink, error) {
	result := new(UserLink)
	_, err := s.client.Post(ctx, "/api/user/links", body, result)
	return result, err
}

// UpdateLink updates a custom link.
func (s *Service) UpdateLink(ctx context.Context, linkID int, body UserLinkRequest) error {
	u := fmt.Sprintf("/api/user/links/%d", linkID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// DeleteLink deletes a custom link.
func (s *Service) DeleteLink(ctx context.Context, linkID int) error {
	u := fmt.Sprintf("/api/user/links/%d", linkID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// --- User Course Operations ---

// CheckCourseGraduate checks if a course is graduated.
func (s *Service) CheckCourseGraduate(ctx context.Context, courseID int) (CourseGraduateCheckResponse, error) {
	u := fmt.Sprintf("/api/user/check-course-graduate?course_id=%d", courseID)
	var result CourseGraduateCheckResponse
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CheckCoursesGraduate checks graduation status for multiple courses and a user context.
func (s *Service) CheckCoursesGraduate(ctx context.Context, body CheckCoursesGraduateRequest) (CourseGraduateCheckResponse, error) {
	var result CourseGraduateCheckResponse
	_, err := s.client.Post(ctx, "/api/user/check-course-graduate", body, &result)
	return result, err
}

// CheckExpiredPassword checks if the user's password is expired.
func (s *Service) CheckExpiredPassword(ctx context.Context) (*ExpiredPasswordResponse, error) {
	result := new(ExpiredPasswordResponse)
	_, err := s.client.Get(ctx, "/api/user/check-expired-password", result)
	return result, err
}

// GetFirstTimeLogin returns if this is the user's first login.
func (s *Service) GetFirstTimeLogin(ctx context.Context) (FirstTimeLoginResponse, error) {
	var result FirstTimeLoginResponse
	_, err := s.client.Get(ctx, "/api/user/first-time-login", &result)
	return result, err
}

// UpdateFirstTimeLogin updates first-time-login state.
func (s *Service) UpdateFirstTimeLogin(ctx context.Context) error {
	_, err := s.client.Put(ctx, "/api/user/first-time-login", nil, nil)
	return err
}

// GetPreTask returns the user's pre-task information.
func (s *Service) GetPreTask(ctx context.Context) (*PreTaskResponse, error) {
	result := new(PreTaskResponse)
	_, err := s.client.Get(ctx, "/api/user/pre-task", result)
	return result, err
}

// GetPersonas returns the user's personas.
func (s *Service) GetPersonas(ctx context.Context) (*PersonasResponse, error) {
	result := new(PersonasResponse)
	_, err := s.client.Get(ctx, "/api/user/personas", result)
	return result, err
}

// CreatePersonas submits the user's personas information.
func (s *Service) CreatePersonas(ctx context.Context, body PersonasRequest) error {
	_, err := s.client.Post(ctx, "/api/user/personas", body, nil)
	return err
}

// GetDepartment returns the user's department.
func (s *Service) GetDepartment(ctx context.Context) (*Department, error) {
	result := new(Department)
	_, err := s.client.Get(ctx, "/api/user/department", result)
	return result, err
}

// UpdateDepartment updates the user's department.
func (s *Service) UpdateDepartment(ctx context.Context, body DepartmentUpdateRequest) (*Department, error) {
	result := new(Department)
	_, err := s.client.Put(ctx, "/api/user/department", body, result)
	return result, err
}

// GetAssociationCode returns the user's association code.
func (s *Service) GetAssociationCode(ctx context.Context) (*AssociationCodeResponse, error) {
	result := new(AssociationCodeResponse)
	_, err := s.client.Get(ctx, "/api/user/association-code", result)
	return result, err
}

// ResetAssociationCode resets the user's association code.
func (s *Service) ResetAssociationCode(ctx context.Context) (*AssociationCodeResponse, error) {
	result := new(AssociationCodeResponse)
	_, err := s.client.Put(ctx, "/api/user/association-code", nil, result)
	return result, err
}

// GetChat returns the user's chat history.
func (s *Service) GetChat(ctx context.Context) ([]*ChatMessage, error) {
	var result []*ChatMessage
	_, err := s.client.Get(ctx, "/api/user/chat", &result)
	return result, err
}

// SendChatMessage sends a chat message.
func (s *Service) SendChatMessage(ctx context.Context, body UserChatRequest) (*ChatMessage, error) {
	result := new(ChatMessage)
	_, err := s.client.Post(ctx, "/api/user/chat", body, result)
	return result, err
}

// GetFailedCourses returns the user's failed courses.
func (s *Service) GetFailedCourses(ctx context.Context) (*FailedCoursesResponse, error) {
	result := new(FailedCoursesResponse)
	_, err := s.client.Get(ctx, "/api/user/failed-courses", result)
	return result, err
}

// GetCourseCertificationScores returns course certification scores.
func (s *Service) GetCourseCertificationScores(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/user/course-certification/scores", &result)
	return result, err
}

// RecalculateCourseCertificationScores recalculates course certification scores.
func (s *Service) RecalculateCourseCertificationScores(ctx context.Context, body RecalculateCourseCertificationScoresRequest) error {
	_, err := s.client.Put(ctx, "/api/user/course-certification/scores", body, nil)
	return err
}

// GetAcademicLearningResources returns academic learning resources.
func (s *Service) GetAcademicLearningResources(ctx context.Context, params AcademicLearningResourcesParams) (*AcademicLearningResourcesResponse, error) {
	query := map[string]string{}
	if params.Page > 0 {
		query["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.PageSize > 0 {
		query["page_size"] = fmt.Sprintf("%d", params.PageSize)
	}
	if params.CourseCode != "" {
		query["course_code"] = params.CourseCode
	}
	if params.Keyword != "" {
		query["keyword"] = params.Keyword
	}
	for i, typ := range params.Types {
		query[fmt.Sprintf("types[%d]", i)] = typ
	}
	result := new(AcademicLearningResourcesResponse)
	_, err := s.client.Get(ctx, addQueryParams("/api/user/academic-learning-resources", query), result)
	return result, err
}

// GetThirdPartResources returns third-party resources.
func (s *Service) GetThirdPartResources(ctx context.Context, params ThirdPartResourcesParams) (*ThirdPartResourcesResponse, error) {
	query := map[string]string{}
	if params.Page > 0 {
		query["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.PageSize > 0 {
		query["page_size"] = fmt.Sprintf("%d", params.PageSize)
	}
	if params.Conditions != "" {
		query["conditions"] = params.Conditions
	}
	result := new(ThirdPartResourcesResponse)
	_, err := s.client.Get(ctx, addQueryParams("/api/user/third-part-resources", query), result)
	return result, err
}

// GetOtherVideoResources returns other video resources.
func (s *Service) GetOtherVideoResources(ctx context.Context, params OtherVideoResourcesParams) (*OtherVideoResourcesResponse, error) {
	query := map[string]string{}
	if params.Page > 0 {
		query["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.PageSize > 0 {
		query["page_size"] = fmt.Sprintf("%d", params.PageSize)
	}
	result := new(OtherVideoResourcesResponse)
	_, err := s.client.Get(ctx, addQueryParams("/api/user/other-video-resources", query), result)
	return result, err
}

// --- Notes ---

// ListNotes returns the user's notes.
func (s *Service) ListNotes(ctx context.Context, params map[string]string) (NotesResponse, error) {
	u := addQueryParams("/api/notes", params)
	var result NotesResponse
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateNote creates a note.
func (s *Service) CreateNote(ctx context.Context, body NoteRequest) (*Note, error) {
	result := new(Note)
	_, err := s.client.Post(ctx, "/api/notes", body, result)
	return result, err
}

// UpdateNote updates a note.
func (s *Service) UpdateNote(ctx context.Context, noteID int, body NoteRequest) (*Note, error) {
	u := fmt.Sprintf("/api/notes/%d", noteID)
	result := new(Note)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteNote deletes a note.
func (s *Service) DeleteNote(ctx context.Context, noteID int) error {
	u := fmt.Sprintf("/api/notes/%d", noteID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// --- Notebooks ---

// GetNotebook returns a notebook.
func (s *Service) GetNotebook(ctx context.Context, notebookID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/notebooks/%d", notebookID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Sign In ---

// SignIn performs a sign-in to the platform.
func (s *Service) SignIn(ctx context.Context) (*SignIn, error) {
	result := new(SignIn)
	_, err := s.client.Post(ctx, "/api/sign-in", nil, result)
	return result, err
}

// GetSignInStats returns sign-in statistics.
func (s *Service) GetSignInStats(ctx context.Context) (*SignInStatsResponse, error) {
	result := new(SignInStatsResponse)
	_, err := s.client.Get(ctx, "/api/sign-in/stats", result)
	return result, err
}

// --- Email Verification ---

// SendVerificationEmail sends a verification email.
func (s *Service) SendVerificationEmail(ctx context.Context) error {
	_, err := s.client.Post(ctx, "/api/user/send_verification_email", struct{}{}, nil)
	return err
}

// SendOrgSignUpVerificationEmail sends an org sign-up verification email.
func (s *Service) SendOrgSignUpVerificationEmail(ctx context.Context) error {
	_, err := s.client.Post(ctx, "/api/user/send_org_sing_up_verification_email", struct{}{}, nil)
	return err
}

// --- Captures ---

// ListMyCaptures returns the user's captures.
func (s *Service) ListMyCaptures(ctx context.Context) (*MyCapturesResponse, error) {
	result := new(MyCapturesResponse)
	_, err := s.client.Get(ctx, "/api/my-captures", result)
	return result, err
}

// ListPublicCaptures returns public captures.
func (s *Service) ListPublicCaptures(ctx context.Context) (*PublicCapturesResponse, error) {
	result := new(PublicCapturesResponse)
	_, err := s.client.Get(ctx, "/api/public-captures", result)
	return result, err
}

// GetCapture returns a specific capture.
func (s *Service) GetCapture(ctx context.Context, captureID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/captures/%d", captureID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- User Index Stats ---

// GetCoursesInfoStatus returns course info status for the user index.
func (s *Service) GetCoursesInfoStatus(ctx context.Context) (*CoursesInfoStatusResponse, error) {
	result := new(CoursesInfoStatusResponse)
	_, err := s.client.Get(ctx, "/api/user-index-stat/courses/info-status", result)
	return result, err
}

// GetOrgSummary returns the organization summary for the user index.
func (s *Service) GetOrgSummary(ctx context.Context) (OrgSummaryResponse, error) {
	var result OrgSummaryResponse
	_, err := s.client.Get(ctx, "/api/user-index-stat/org-summary", &result)
	return result, err
}

// GetCoursesIdentities returns course identities.
func (s *Service) GetCoursesIdentities(ctx context.Context) (*CoursesIdentitiesResponse, error) {
	result := new(CoursesIdentitiesResponse)
	_, err := s.client.Get(ctx, "/api/courses-identities?no-intercept=true", result)
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
