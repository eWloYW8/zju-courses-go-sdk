package meetings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
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

// --- VTRS (Virtual Teaching Room System) ---

// ListVTRSes returns VTRS entries.
func (s *Service) ListVTRSes(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/vtrses", &result)
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
	u := fmt.Sprintf("/api/vtrses/access-code/%d", vtrsID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListVTRSMeetingClassifications returns VTRS meeting classifications.
func (s *Service) ListVTRSMeetingClassifications(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/meetings/classifications/%d", vtrsID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListVTRSResourceClassifications returns VTRS resource classifications.
func (s *Service) ListVTRSResourceClassifications(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/resources/classifications/%d", vtrsID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ShareVTRSResources shares VTRS resources.
func (s *Service) ShareVTRSResources(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/vtrses/share-resources", body, nil)
	return err
}

// ListVTRSSubjectLibs returns VTRS subject libraries.
func (s *Service) ListVTRSSubjectLibs(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/vtrses/subject-libs", &result)
	return result, err
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
func (s *Service) ListLessonRooms(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/lesson-rooms", &result)
	return result, err
}

// ListRoomLocations returns room locations.
func (s *Service) ListRoomLocations(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/room-locations", &result)
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
