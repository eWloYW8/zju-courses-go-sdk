package meetings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles meeting and live classroom related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Meeting ---

// GetMeeting returns a meeting.
func (s *Service) GetMeeting(ctx context.Context, meetingID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/meeting/%d", meetingID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetMeetingWeekTimePeriods returns meeting week time periods.
func (s *Service) GetMeetingWeekTimePeriods(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/meeting/week/time-periods", &result)
	return result, err
}

// GetShanghaiTechMeeting returns a ShanghaiTech meeting.
func (s *Service) GetShanghaiTechMeeting(ctx context.Context, meetingID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/meeting/shanghaitech/%d", meetingID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Tencent Meeting ---

// GetTencentMeeting returns a Tencent meeting.
func (s *Service) GetTencentMeeting(ctx context.Context, meetingID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/tencent-meeting/%d", meetingID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// DeleteTencentSubMeeting deletes a Tencent recurring sub-meeting.
func (s *Service) DeleteTencentSubMeeting(ctx context.Context, subMeetingID int) error {
	u := fmt.Sprintf("/api/tencent-meeting/%d", subMeetingID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// GetTencentMeetingAuthURL returns the Tencent meeting authorization URL.
func (s *Service) GetTencentMeetingAuthURL(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/tencent-meeting/authorization-url", &result)
	return result, err
}

// CheckTencentMeetingUserAuth checks Tencent meeting user auth.
func (s *Service) CheckTencentMeetingUserAuth(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/tencent_meeting/check-user-auth", &result)
	return result, err
}

// ListTencentMeetingActivities returns Tencent meeting activities for a course.
func (s *Service) ListTencentMeetingActivities(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/tencent-meeting/activities?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- DingTalk ---

// GetDingTalkChat returns DingTalk chat info for a course.
func (s *Service) GetDingTalkChat(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/ding-talk/chat?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetDingTalkUserID returns the DingTalk user ID.
func (s *Service) GetDingTalkUserID(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/ding-talk/user-id", &result)
	return result, err
}

// GetDingTalkLive returns a DingTalk live session.
func (s *Service) GetDingTalkLive(ctx context.Context, liveID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/dingtalk-lives/%d", liveID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetZoomSettings returns organization zoom settings.
func (s *Service) GetZoomSettings(ctx context.Context, orgID int) (*OrgZoomSettingsResponse, error) {
	u := fmt.Sprintf("/api/orgs/%d/zoom-settings", orgID)
	result := new(OrgZoomSettingsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetZoomUserInfo returns zoom account dispatch information for a user.
func (s *Service) GetZoomUserInfo(ctx context.Context, userID int) (*ZoomUserInfoResponse, error) {
	u := fmt.Sprintf("/api/user/%d/zoom-info", userID)
	result := new(ZoomUserInfoResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- VTRS (Virtual Teaching Room System) ---

// ListVTRSes returns VTRS entries.
func (s *Service) ListVTRSes(ctx context.Context) (json.RawMessage, error) {
	return s.ListVTRSesWithParams(ctx, nil)
}

// ListVTRSesWithParams returns VTRS entries with frontend filter options.
func (s *Service) ListVTRSesWithParams(ctx context.Context, params *ListVTRSesParams) (json.RawMessage, error) {
	u := "/api/vtrses"
	if params != nil {
		query := map[string]string{}
		if params.Conditions != "" {
			query["conditions"] = params.Conditions
		}
		if params.NeedStat != nil {
			if *params.NeedStat {
				query["needStat"] = "true"
			} else {
				query["needStat"] = "false"
			}
		}
		if params.Fields != "" {
			query["fields"] = params.Fields
		}
		u = addQueryParams(u, query)
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVTRS returns a specific VTRS entry.
func (s *Service) GetVTRS(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/%d", vtrsID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetVTRSAccessCode gets the access code for a VTRS.
func (s *Service) GetVTRSAccessCode(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/%d/access-code", vtrsID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// RefreshVTRSAccessCode refreshes the access code for a VTRS.
func (s *Service) RefreshVTRSAccessCode(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/%d/access-code", vtrsID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, nil, &result)
	return result, err
}

// ValidateVTRSAccessCode validates a VTRS access code.
func (s *Service) ValidateVTRSAccessCode(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/access-code/%d/validate", vtrsID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListVTRSMeetingClassifications returns VTRS meeting classifications.
func (s *Service) ListVTRSMeetingClassifications(ctx context.Context, vtrsID int) (*VTRSMeetingClassificationsResponse, error) {
	u := fmt.Sprintf("/api/vtrses/%d/meetings/classifications", vtrsID)
	result := new(VTRSMeetingClassificationsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// DeleteVTRSMeetingClassification deletes a meeting classification.
func (s *Service) DeleteVTRSMeetingClassification(ctx context.Context, classificationID int) error {
	u := fmt.Sprintf("/api/vtrses/meetings/classifications/%d", classificationID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// CreateVTRSMeetingClassification creates a meeting classification.
func (s *Service) CreateVTRSMeetingClassification(ctx context.Context, vtrsID int, body *CreateVTRSMeetingClassificationRequest) (*VTRSMeetingClassification, error) {
	u := fmt.Sprintf("/api/vtrses/%d/meetings/classifications", vtrsID)
	result := new(VTRSMeetingClassification)
	var raw struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	_, err := s.client.Post(ctx, u, body, &raw)
	if err != nil {
		return nil, err
	}
	result.ID = raw.ID
	result.Name = raw.Name
	return result, nil
}

// UpdateVTRSMeetingClassification updates a meeting classification.
func (s *Service) UpdateVTRSMeetingClassification(ctx context.Context, classificationID int, name string) error {
	u := fmt.Sprintf("/api/vtrses/meetings/classifications/%d", classificationID)
	_, err := s.client.Put(ctx, u, &UpdateVTRSMeetingClassificationRequest{ID: classificationID, Name: name}, nil)
	return err
}

// ListVTRSResourceClassifications returns VTRS resource classifications.
func (s *Service) ListVTRSResourceClassifications(ctx context.Context, vtrsID int) (*VTRSResourceClassificationsResponse, error) {
	u := fmt.Sprintf("/api/vtrses/%d/resources/classifications", vtrsID)
	result := new(VTRSResourceClassificationsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// DeleteVTRSResourceClassification deletes a resource classification.
func (s *Service) DeleteVTRSResourceClassification(ctx context.Context, classificationID int) error {
	u := fmt.Sprintf("/api/vtrses/resources/classifications/%d", classificationID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// CreateVTRSResourceClassification creates a resource classification.
func (s *Service) CreateVTRSResourceClassification(ctx context.Context, vtrsID int, body *CreateVTRSResourceClassificationRequest) (*VTRSResourceClassification, error) {
	u := fmt.Sprintf("/api/vtrses/%d/resources/classifications", vtrsID)
	var raw struct {
		Classification *VTRSResourceClassification `json:"classification"`
	}
	_, err := s.client.Post(ctx, u, body, &raw)
	return raw.Classification, err
}

// UpdateVTRSResourceClassification updates a resource classification.
func (s *Service) UpdateVTRSResourceClassification(ctx context.Context, classificationID int, body *UpdateVTRSResourceClassificationRequest) error {
	u := fmt.Sprintf("/api/vtrses/resources/classifications/%d", classificationID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// SortVTRSResourceClassifications sorts resource classifications inside a VTRS.
func (s *Service) SortVTRSResourceClassifications(ctx context.Context, vtrsID int, classificationIDs []int) error {
	u := fmt.Sprintf("/api/vtrses/%d/resources/classifications/sort", vtrsID)
	_, err := s.client.Put(ctx, u, &SortVTRSResourceClassificationsRequest{Classifications: classificationIDs}, nil)
	return err
}

// ShareVTRSResources shares VTRS resources.
func (s *Service) ShareVTRSResources(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/vtrses/share-resources", body, nil)
	return err
}

// ListVTRSShareResources returns paged shared resources for VTRS contexts.
func (s *Service) ListVTRSShareResources(ctx context.Context, params *ListVTRSShareResourcesParams) (*VTRSShareResourcesResponse, error) {
	query := map[string]string{}
	if params != nil {
		if params.RefParentType != "" {
			query["ref_parent_type"] = params.RefParentType
		}
		if params.Page > 0 {
			query["page"] = fmt.Sprintf("%d", params.Page)
		}
		if params.PageSize > 0 {
			query["page_size"] = fmt.Sprintf("%d", params.PageSize)
		}
		if params.Conditions != "" {
			query["conditions"] = params.Conditions
		}
	}
	result := new(VTRSShareResourcesResponse)
	_, err := s.client.Get(ctx, addQueryParams("/api/vtrses/share-resources", query), result)
	return result, err
}

// CreateVTRSResources adds resources to a VTRS.
func (s *Service) CreateVTRSResources(ctx context.Context, vtrsID int, body *CreateVTRSResourcesRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/%d/resources", vtrsID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ListVTRSResources returns paged resources under a VTRS.
func (s *Service) ListVTRSResources(ctx context.Context, vtrsID int, params *ListVTRSResourcesParams) (*VTRSResourcesResponse, error) {
	query := map[string]string{}
	if params != nil {
		if params.ParentFolderID != nil {
			query["parent_folder_id"] = fmt.Sprintf("%d", *params.ParentFolderID)
		}
		if params.ClassificationID != nil {
			query["classification_id"] = fmt.Sprintf("%d", *params.ClassificationID)
		}
		if params.Page > 0 {
			query["page"] = fmt.Sprintf("%d", params.Page)
		}
		if params.PageSize > 0 {
			query["page_size"] = fmt.Sprintf("%d", params.PageSize)
		}
		if params.Conditions != "" {
			query["conditions"] = params.Conditions
		}
	}
	u := addQueryParams(fmt.Sprintf("/api/vtrses/%d/resources", vtrsID), query)
	result := new(VTRSResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetVTRSResourcesSummary returns resource summary counts for a VTRS.
func (s *Service) GetVTRSResourcesSummary(ctx context.Context, vtrsID int) (VTRSResourcesSummaryResponse, error) {
	u := fmt.Sprintf("/api/vtrses/%d/resources/summary", vtrsID)
	var result VTRSResourcesSummaryResponse
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// DeleteVTRSResources removes resources from a VTRS.
func (s *Service) DeleteVTRSResources(ctx context.Context, vtrsID int, body *UploadReferenceIDsRequest) error {
	u := fmt.Sprintf("/api/vtrses/%d/resources", vtrsID)
	_, err := s.client.DeleteWithBody(ctx, u, body, nil)
	return err
}

// UpdateVTRSResources updates resource references inside a VTRS.
func (s *Service) UpdateVTRSResources(ctx context.Context, vtrsID int, body *UploadReferencesRequest) error {
	u := fmt.Sprintf("/api/vtrses/%d/resources", vtrsID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// SaveVTRSResources saves selected VTRS resources into another context.
func (s *Service) SaveVTRSResources(ctx context.Context, vtrsID int, body *UploadReferenceIDsRequest) error {
	u := fmt.Sprintf("/api/vtrses/%d/save-resources", vtrsID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// ListVTRSSubjectLibs returns VTRS subject libraries.
func (s *Service) ListVTRSSubjectLibs(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/vtrses/subject-libs", &result)
	return result, err
}

// CreateVTRSSubjectLib creates a subject library under a VTRS.
func (s *Service) CreateVTRSSubjectLib(ctx context.Context, vtrsID int, body *CreateVTRSSubjectLibRequest, libType string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/%d/subject-libs", vtrsID)
	if libType != "" {
		u = addQueryParams(u, map[string]string{"lib_type": libType})
	}
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// ListVTRSSubjectLibsWithParams returns paged VTRS subject libraries.
func (s *Service) ListVTRSSubjectLibsWithParams(ctx context.Context, vtrsID int, params *ListVTRSSubjectLibsParams) (*VTRSSubjectLibsResponse, error) {
	query := map[string]string{}
	if params != nil {
		if params.Keyword != "" {
			query["keyword"] = params.Keyword
		}
		if params.ParentID != nil {
			query["parent_id"] = fmt.Sprintf("%d", *params.ParentID)
		}
		if params.ClassificationID != nil {
			query["classification_id"] = fmt.Sprintf("%d", *params.ClassificationID)
		}
		if params.Page > 0 {
			query["page"] = fmt.Sprintf("%d", params.Page)
		}
		if params.PageSize > 0 {
			query["page_size"] = fmt.Sprintf("%d", params.PageSize)
		}
		if params.Predicate != "" {
			query["predicate"] = params.Predicate
		}
		if params.Reverse != nil {
			if *params.Reverse {
				query["reverse"] = "true"
			} else {
				query["reverse"] = "false"
			}
		}
		if params.LibType != "" {
			query["lib_type"] = params.LibType
		}
	}
	u := addQueryParams(fmt.Sprintf("/api/vtrses/%d/subject-libs", vtrsID), query)
	var raw struct {
		SubjectLibs []json.RawMessage `json:"subject_libs"`
		Page        int               `json:"page"`
		PageSize    int               `json:"page_size"`
		Pages       int               `json:"pages"`
		Total       int               `json:"total"`
	}
	_, err := s.client.Get(ctx, u, &raw)
	if err != nil {
		return nil, err
	}
	return &VTRSSubjectLibsResponse{
		Items: raw.SubjectLibs,
		Pagination: model.Pagination{
			Page:     raw.Page,
			PageSize: raw.PageSize,
			Pages:    raw.Pages,
			Total:    raw.Total,
		},
	}, nil
}

// --- Instruction Team Meeting ---

// GetInstructionTeamMeeting returns instruction team meeting.
func (s *Service) GetInstructionTeamMeeting(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/instruction-team/meeting", &result)
	return result, err
}

// --- Combine Courses ---

// ListCombineCourses returns combined courses.
func (s *Service) ListCombineCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/combine-courses", &result)
	return result, err
}

// CreateCombineCourse creates a combined course.
func (s *Service) CreateCombineCourse(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/combine-courses", body, &result)
	return result, err
}

// DeleteCombineCourse deletes a combined course.
func (s *Service) DeleteCombineCourse(ctx context.Context, combineID int) error {
	u := fmt.Sprintf("/api/combine-courses/%d", combineID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// --- Lesson Rooms & Room Locations ---

// ListLessonRooms returns lesson rooms.
func (s *Service) ListLessonRooms(ctx context.Context) ([]*LessonRoom, error) {
	var result []*LessonRoom
	_, err := s.client.Get(ctx, "/api/lesson-rooms", &result)
	return result, err
}

// ListRoomLocations returns room locations for a course.
func (s *Service) ListRoomLocations(ctx context.Context, courseID int) (*RoomLocationsResponse, error) {
	u := fmt.Sprintf("/api/course/%d/room-locations", courseID)
	result := new(RoomLocationsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListEnabledRoomLocations returns room locations that are available for the given time window.
func (s *Service) ListEnabledRoomLocations(ctx context.Context, orgID int, startTime, endTime string) (*RoomLocationsResponse, error) {
	u := fmt.Sprintf("/api/org/%d/enable-room-locations", orgID)
	u = addQueryParams(u, map[string]string{
		"start_time": startTime,
		"end_time":   endTime,
	})
	result := new(RoomLocationsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Live Activities & Records ---

// GetLiveActivity returns a live activity.
func (s *Service) GetLiveActivity(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/live-activities/%d", activityID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetLiveRecord returns a live record.
func (s *Service) GetLiveRecord(ctx context.Context, recordID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/live-records/%d", recordID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetLectureLiveSchedule returns a lecture live schedule.
func (s *Service) GetLectureLiveSchedule(ctx context.Context, scheduleID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/lecture-live/schedule/%d", scheduleID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetLectureLive returns a lecture live session.
func (s *Service) GetLectureLive(ctx context.Context, jwt string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/lecture-live?jwt=%s", jwt)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Helpers ---

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}
