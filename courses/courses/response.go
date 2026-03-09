package courses

import (
	"encoding/json"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

type OutlineResponse = Outline

type OutlineItemResponse struct {
	CourseOutlineItem *OutlineField `json:"course_outline_item,omitempty"`
}

type MyCoursesResponse struct {
	Courses []*Course `json:"courses"`
	model.Pagination
}

type ModulesResponse struct {
	Modules []*Module `json:"modules"`
}

type EnrollmentsResponse struct {
	Enrollments []*Enrollment `json:"enrollments"`
}

type CourseInstructorsResponse struct {
	Instructors []*model.User `json:"instructors"`
}

type CourseHostsResponse struct {
	Hosts []*model.User `json:"hosts"`
}

type CourseGroupUsersResponse struct {
	Result []*model.User `json:"result"`
}

type AvailableUsersResponse struct {
	Users []*model.User `json:"users"`
}

type StudentsResponse struct {
	Students []*model.User `json:"students"`
}

type SectionsResponse struct {
	Sections []*Section `json:"sections"`
}

type EntryRefersResponse struct {
	EntryRefers []*EntryRefer `json:"entry_refers"`
}

type ScoreItemGroupsResponse struct {
	Items []*ScoreItemGroup `json:"items"`
}

type ActivitiesResponse struct {
	Activities []*activities.Activity `json:"activities"`
}

type NavSettingResponse struct {
	NavSetting          []*NavSetting `json:"nav_setting"`
	CanNotDisabledItems []string      `json:"can_not_disabled_items"`
}

type HomeworkScoresResponse struct {
	HomeworkActivities []*HomeworkActivity `json:"homework_activities"`
	Scores             []*HomeworkScore    `json:"scores"`
}

type HomeworkSubmissionStatusResponse struct {
	CourseID           int                         `json:"course_id"`
	HomeworkActivities []*HomeworkSubmissionStatus `json:"homework_activities"`
}

type CatalogActivitiesResponse map[string]any

type HomeworkSubmissionNumberResponse map[string]any

type ExamScoresResponse struct {
	ExamScores []*ExamScore `json:"exam_scores"`
}

type CourseExamSubmissionsResponse map[string]any

type ClassroomExamScoresResponse map[string]any

type UsersSmallAvatarsResponse struct {
	Avatars map[string]string `json:"avatars"`
}

type LiveRecordsResponse struct {
	LiveRecords []*LiveRecord `json:"live_records"`
}

type LiveRecordAsrTaskResponse map[string]any

type LiveRecordAsrTaskStatusResponse map[string]any

type LiveRecordCaptionsResponse map[string]any

type ActivityReadsForUserResponse struct {
	ActivityReads []*activities.ActivityRead `json:"activity_reads"`
}

type CourseCustomScoreItemsResponse struct {
	CustomScoreItems []*CustomScoreItem `json:"custom_score_items"`
}

type CourseAccessCodeResponse struct {
	AccessCode string `json:"access_code"`
	Resetable  bool   `json:"resetable"`
}

type CourseAccessCodeValidationResponse struct {
	Message string `json:"message,omitempty"`
}

type OpenedOrgsResponse struct {
	Orgs []*model.OrgDetail `json:"orgs"`
}

type EntryRecordResponse map[string]any

type OnlineVideoCompletenessSettingResponse map[string]any

type PerformanceScoreSettingResponse map[string]any

type AnnounceScoreSettingsResponse map[string]any

type PerformanceScorePercentageResponse map[string]any

type WarningListResponse map[string]any

type WarningResponse map[string]any

type WarningThresholdResponse map[string]any

type WarningStudentsResponse map[string]any

type ScoreRanksResponse map[string]any

type ScoreItemSettingsResponse map[string]any

type EnrollmentRawScoreResponse map[string]any

type CourseScoreTypeResponse struct {
	ScoreType string `json:"score_type,omitempty"`
}

type StudentsPerformanceResponse map[string]any

type ScorePercentagesResponse map[string]any

type StudentSelfScoreResponse map[string]any

type RollcallScoreResponse map[string]any

type PerformanceScoreResponse map[string]any

type CourseQuestionnairesResponse map[string]any

type QuestionnaireListResponse map[string]any

type QuestionnaireScoresResponse map[string]any

type CourseWebLinkScoresResponse map[string]any

type CourseVirtualExperimentScoresResponse map[string]any

type CourseVirtualExperimentsResponse map[string]any

type AllActivitiesResponse map[string]any

type SHTVUGroupsResponse map[string]any

type ScoreDistributionResponse map[string]any

type ScoreTypeSettingsResponse map[string]any

type OnlineVideoCompletenessScoresResponse map[string]any

type OnlineVideoCompletenessScoreResponse map[string]any

type CourseListResponse map[string]any

type CourseCountResponse map[string]any

type SettingsResponse map[string]any

type BlueprintSubItemsResultResponse map[string]any

type BlueprintAllSubActivitiesCountResponse map[string]any

type BlueprintSubCoursesResponse struct {
	Courses []*Course `json:"courses"`
}

type BlueprintCheckPrerequisitesResponse struct {
	HasPrerequisites bool `json:"has_prerequisites"`
}

type BlueprintSubmittedInfoResponse struct {
	UnableSync []*BlueprintSubmittedInfo `json:"unable_sync,omitempty"`
	NeedSync   []*BlueprintSubmittedInfo `json:"need_sync,omitempty"`
}

type CourseClassificationsResponse struct {
	Classifications []*CourseClassification `json:"classifications"`
}

type CourseAuditReferencesResponse struct {
	AuditReferences []json.RawMessage `json:"audit_references"`
	model.Pagination
}
