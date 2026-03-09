package others

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles miscellaneous API operations that don't fit neatly into other modules.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Projects ---

// ListProjects returns projects.
func (s *Service) ListProjects(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/projects", &result)
	return result, err
}

// ListProjectsWithParams returns projects with frontend paging/filter parameters.
func (s *Service) ListProjectsWithParams(ctx context.Context, params ListProjectsParams) (*ProjectsResponse, error) {
	u := addListOptions("/api/projects", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	result := new(ProjectsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetProject returns a project.
func (s *Service) GetProject(ctx context.Context, projectID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/project/%d", projectID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateProject creates a project.
func (s *Service) CreateProject(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/project", body, &result)
	return result, err
}

// CreateProjectTyped creates a project and decodes the created project payload.
func (s *Service) CreateProjectTyped(ctx context.Context, body *CreateProjectRequest) (*Project, error) {
	result := new(Project)
	_, err := s.client.Post(ctx, "/api/project", body, result)
	return result, err
}

// ApplyProject submits a project application.
func (s *Service) ApplyProject(ctx context.Context, projectID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/projects/%d/apply", projectID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, nil, &result)
	return result, err
}

// ListProjectApplications returns paged applications for a project.
func (s *Service) ListProjectApplications(ctx context.Context, projectID int, params ListProjectsParams) (*ProjectApplicationsResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/projects/%d/apply", projectID), &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	result := new(ProjectApplicationsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateProject updates a project.
func (s *Service) UpdateProject(ctx context.Context, projectID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/project/%d", projectID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// UpdateProjectTyped updates a project and decodes the updated project payload.
func (s *Service) UpdateProjectTyped(ctx context.Context, projectID int, body *UpdateProjectRequest) (*Project, error) {
	u := fmt.Sprintf("/api/project/%d", projectID)
	result := new(Project)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// AuditProjectApplication audits a project application.
func (s *Service) AuditProjectApplication(ctx context.Context, projectID, applicationID int, status string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/projects/%d/audit/%d", projectID, applicationID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, AuditProjectApplicationRequest{Status: status}, &result)
	return result, err
}

// ListProjectSharedResources returns project shared resources using the frontend conditions payload.
func (s *Service) ListProjectSharedResources(ctx context.Context, projectID int, conditions ProjectSharedResourceConditions) ([]*ProjectSharedResource, error) {
	u := fmt.Sprintf("/api/project/%d/share-resource", projectID)
	if encoded := encodeConditions(conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	var result []*ProjectSharedResource
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateProjectSharedResources creates project shared-resource references.
func (s *Service) CreateProjectSharedResources(ctx context.Context, projectID int, body *ProjectSharedResourceRequest) error {
	u := fmt.Sprintf("/api/project/%d/share-resource", projectID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// UpdateProjectSharedResources updates project shared-resource references.
func (s *Service) UpdateProjectSharedResources(ctx context.Context, projectID int, body *ProjectSharedResourceRequest) error {
	u := fmt.Sprintf("/api/project/%d/share-resource", projectID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// DeleteProjectSharedResource deletes a project shared-resource reference.
func (s *Service) DeleteProjectSharedResource(ctx context.Context, projectID int, params *DeleteProjectSharedResourceRequest) error {
	u := fmt.Sprintf("/api/project/%d/share-resource", projectID)
	query := map[string]string{
		"reference_id": fmt.Sprintf("%d", params.ReferenceID),
		"upload_id":    fmt.Sprintf("%d", params.UploadID),
	}
	if len(params.NodeIDs) > 0 {
		query["node_ids"] = encodeConditions(params.NodeIDs)
	}
	u = addQueryParams(u, query)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// --- Studio ---

// GetStudio returns a studio.
func (s *Service) GetStudio(ctx context.Context, studioID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/studio/%d", studioID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Chatrooms ---

// GetChatroom returns a chatroom.
func (s *Service) GetChatroom(ctx context.Context, chatroomID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/chatrooms/%d", chatroomID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Cloud Classroom ---

// GetCloudClassroom returns cloud classroom info.
func (s *Service) GetCloudClassroom(ctx context.Context, start string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/activity/cloud-classroom?start=%s", start)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Classroom Subjects ---

// GetClassroomSubjectsRule returns classroom subject rules.
func (s *Service) GetClassroomSubjectsRule(ctx context.Context, classroomID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/classroom/%d/subjects-rule", classroomID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetClassroomSubjects returns classroom subjects.
func (s *Service) GetClassroomSubjects(ctx context.Context, classroomID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/classroom/%d/subject", classroomID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Questionnaires ---

// GetQuestionnaire returns a questionnaire.
func (s *Service) GetQuestionnaire(ctx context.Context, questionnaireID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/questionnaire/%d", questionnaireID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ExportQuestionnaireExcel exports questionnaire results as Excel.
func (s *Service) ExportQuestionnaireExcel(ctx context.Context, questionnaireID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/questionnaire/%d/export/excel", questionnaireID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ExportQuestionnaireCSV exports questionnaire results as CSV.
func (s *Service) ExportQuestionnaireCSV(ctx context.Context, questionnaireID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/questionnaire/%d/export/csv", questionnaireID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// SortQuestionnaireSubjects sorts questionnaire subjects.
func (s *Service) SortQuestionnaireSubjects(ctx context.Context, questionnaireID int, body interface{}) error {
	u := fmt.Sprintf("/api/questionnaire/%d/subject-sort", questionnaireID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// GetQuestionnaireSubjects returns questionnaire subjects.
func (s *Service) GetQuestionnaireSubjects(ctx context.Context, questionnaireID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/questionnaire/%d/subjects", questionnaireID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListQuestionnaires returns questionnaires.
func (s *Service) ListQuestionnaires(ctx context.Context, questionnaireID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/questionnaires/%d", questionnaireID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Ask Questions ---

// GetAskQuestion returns a question.
func (s *Service) GetAskQuestion(ctx context.Context, questionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/ask-questions/%d", questionID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Entries ---

// ListEntries returns entries.
func (s *Service) ListEntries(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/entries", &result)
	return result, err
}

// ListEntriesWithParams returns the frontend paged entry list with typed items.
func (s *Service) ListEntriesWithParams(ctx context.Context, params ListEntriesParams) (*EntriesResponse, error) {
	u := addListOptions("/api/entries", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	fields := params.Fields
	if fields == "" {
		fields = "id,org_id,name,created_at,updated_at,created_by_id,updated_by_id,reference_count"
	}
	query := map[string]string{"fields": fields}
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		query["conditions"] = encoded
	}
	u = addQueryParams(u, query)
	result := new(EntriesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateEntry creates an entry.
func (s *Service) CreateEntry(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/entries", body, &result)
	return result, err
}

// GetEntry returns an entry.
func (s *Service) GetEntry(ctx context.Context, entryID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/entries/%d", entryID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetEntryTyped returns the frontend entry detail payload with default fields.
func (s *Service) GetEntryTyped(ctx context.Context, entryID int) (*Entry, error) {
	u := fmt.Sprintf("/api/entries/%d", entryID)
	u = addQueryParams(u, map[string]string{
		"fields": "id,org_id,name,definition,uploads,keywords,created_at,updated_at,created_by_id,updated_by_id,reference_count",
	})
	result := new(Entry)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListEntryReferences returns paged entry references.
func (s *Service) ListEntryReferences(ctx context.Context, entryID int, params ListEntryReferencesParams) (*EntryReferencesResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/entries/%d/references", entryID), &model.ListOptions{
		Page:     params.Page,
		PageSize: params.PageSize,
	})
	result := new(EntryReferencesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateEntry updates an entry and decodes the updated payload.
func (s *Service) UpdateEntry(ctx context.Context, entryID int, body interface{}) (*Entry, error) {
	u := fmt.Sprintf("/api/entries/%d", entryID)
	result := new(Entry)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteEntry deletes an entry.
func (s *Service) DeleteEntry(ctx context.Context, entryID int) error {
	u := fmt.Sprintf("/api/entries/%d", entryID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// BatchDeleteEntries batch deletes entries.
func (s *Service) BatchDeleteEntries(ctx context.Context, body interface{}) error {
	_, err := s.client.DeleteWithBody(ctx, "/api/entries/batch-delete", body, nil)
	return err
}

// --- Warning ---

// GetStudentWarning returns warning info for a student.
func (s *Service) GetStudentWarning(ctx context.Context, studentID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/warning/student/%d", studentID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Lark Files ---

// ListLarkFiles returns Lark files.
func (s *Service) ListLarkFiles(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/lark/files", &result)
	return result, err
}

// --- WeDrive ---

// GetWeDriveFile returns a WeDrive file.
func (s *Service) GetWeDriveFile(ctx context.Context, fileID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/wedrive/file/%d", fileID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListWeDriveFiles returns WeDrive files with pagination.
func (s *Service) ListWeDriveFiles(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/wedrive/files", opts)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- BlobStorage ---

// GetBlobStorageOpenClientURL returns blob storage open client URL.
func (s *Service) GetBlobStorageOpenClientURL(ctx context.Context, parentID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blobstorage/open-client-url?parent_id=%d", parentID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- ChinamCloud ---

// ListChinamCloudResources returns ChinamCloud resources.
func (s *Service) ListChinamCloudResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/chinamcloud/resources", &result)
	return result, err
}

// UploadChinamCloud uploads to ChinamCloud.
func (s *Service) UploadChinamCloud(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/chinamcloud/upload", body, &result)
	return result, err
}

// --- Campus Subject Lib ---

// ListCampusSubjectLibClassifications returns campus subject lib classifications.
func (s *Service) ListCampusSubjectLibClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/campus-subject-lib/classifications", &result)
	return result, err
}

// GetCampusSubjectLibSubjectCount returns campus subject lib subject count by classifications.
func (s *Service) GetCampusSubjectLibSubjectCount(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/campus-subject-lib/classifications/subject-count", &result)
	return result, err
}

// --- OBE ---

// GetExistedMetrics returns existed OBE metrics.
func (s *Service) GetExistedMetrics(ctx context.Context, params string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/obe/existed-metrics?params=%s", params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Program ---

// ListCoursePrograms returns course programs.
func (s *Service) ListCoursePrograms(ctx context.Context, departmentIDs string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/program/course-programs?department_ids=%s", departmentIDs)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListUserPrograms returns user programs.
func (s *Service) ListUserPrograms(ctx context.Context, fields string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/program/user-programs?fields=%s", fields)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Lessons ---

// GetLesson returns a lesson.
func (s *Service) GetLesson(ctx context.Context, lessonID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/lessons/%d", lessonID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetLessonManagement returns lesson management info.
func (s *Service) GetLessonManagement(ctx context.Context, lessonID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/lessons_management/%d", lessonID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetOrgDownloadCapture returns whether lesson capture download is enabled for an org.
func (s *Service) GetOrgDownloadCapture(ctx context.Context, orgID int) (*FeatureEnabledResponse, error) {
	u := fmt.Sprintf("/api/orgs/%d/download-capture", orgID)
	result := new(FeatureEnabledResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListLessonRooms returns lesson rooms.
func (s *Service) ListLessonRooms(ctx context.Context) ([]*LessonRoom, error) {
	var result []*LessonRoom
	_, err := s.client.Get(ctx, "/api/lesson-rooms", &result)
	return result, err
}

// ListRoomLocations returns the frontend global room-location list.
func (s *Service) ListRoomLocations(ctx context.Context, courseID int) (*RoomLocationsResponse, error) {
	result := new(RoomLocationsResponse)
	_, err := s.client.Get(ctx, "/api/room-locations", result)
	return result, err
}

// ListGlobalRoomLocations returns the frontend lesson room-location list.
func (s *Service) ListGlobalRoomLocations(ctx context.Context) (*RoomLocationsResponse, error) {
	result := new(RoomLocationsResponse)
	_, err := s.client.Get(ctx, "/api/room-locations", result)
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

func encodeConditions(conditions any) string {
	switch value := conditions.(type) {
	case nil:
		return ""
	case string:
		return value
	default:
		encoded, err := json.Marshal(value)
		if err != nil {
			return ""
		}
		return string(encoded)
	}
}
