package courses

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// --- Meeting & Live Operations (attached to various services) ---

// Meeting related methods on Service

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

// --- Live Activities ---

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

// GetLectureLiveActivity returns a lecture live activity for a course.
func (s *Service) GetLectureLiveActivity(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/lecture-live-activity/%d", courseID)
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

// --- Danmu (Bullet Screen) ---

// GetDanmu returns danmu for a course.
func (s *Service) GetDanmu(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/danmu/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Interactions ---

// GetInteraction returns a specific interaction.
func (s *Service) GetInteraction(ctx context.Context, interactionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/interactions/%d", interactionID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// VoteInteraction votes on an interaction.
func (s *Service) VoteInteraction(ctx context.Context, interactionID int, body interface{}) error {
	u := fmt.Sprintf("/api/courses/interactions/vote/%d", interactionID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// --- Interaction Activities ---

// GetInteractionActivity returns an interaction activity.
func (s *Service) GetInteractionActivity(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/interaction-activities/%d", activityID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetInteractionSubmission returns an interaction submission.
func (s *Service) GetInteractionSubmission(ctx context.Context, submissionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/interaction-submissions/%d", submissionID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Rollcall ---

// GetRollcall returns a rollcall.
func (s *Service) GetRollcall(ctx context.Context, rollcallID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d", rollcallID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// StartRollcall starts a rollcall session.
func (s *Service) StartRollcall(ctx context.Context, rollcallID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d/start-rollcall", rollcallID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// StopTimetableRollcall stops a timetable rollcall session.
func (s *Service) StopTimetableRollcall(ctx context.Context, rollcallID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d/stop_time_table_rollcall", rollcallID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// GetRollcallStudentRollcalls returns student rollcalls for a rollcall.
func (s *Service) GetRollcallStudentRollcalls(ctx context.Context, rollcallID int, action string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d/student_rollcalls", rollcallID)
	if action != "" {
		u = addQueryParams(u, map[string]string{"action": action})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetRollcallStudentsPagination returns paginated student rollcalls for a rollcall.
func (s *Service) GetRollcallStudentsPagination(ctx context.Context, rollcallID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/rollcall/%d/pagination_students_rollcalls", rollcallID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetMergedRollcall returns a merged rollcall.
func (s *Service) GetMergedRollcall(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/rollcall/merged-rollcall", &result)
	return result, err
}

// GetMergedRollcallStudentRollcalls returns student rollcalls from merged rollcall.
func (s *Service) GetMergedRollcallStudentRollcalls(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/rollcall/merged-rollcall/student-rollcalls", &result)
	return result, err
}

// GetRollcallStatus returns rollcall status for a course.
func (s *Service) GetRollcallStatus(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/rollcall_status/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListTimetableRollcalls returns timetable rollcalls.
func (s *Service) ListTimetableRollcalls(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/timetable_rollcalls", &result)
	return result, err
}

// --- Groups ---

// GetGroupSet returns a group set.
func (s *Service) GetGroupSet(ctx context.Context, groupSetID int) (*model.GroupSet, error) {
	u := fmt.Sprintf("/api/group-sets/%d", groupSetID)
	result := new(model.GroupSet)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetGroup returns a group.
func (s *Service) GetGroup(ctx context.Context, groupID int) (*model.Group, error) {
	u := fmt.Sprintf("/api/groups/%d", groupID)
	result := new(model.Group)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetMyGroup returns the current user's group within a group set.
func (s *Service) GetMyGroup(ctx context.Context, groupSetID int) (*model.Group, error) {
	u := fmt.Sprintf("/api/group-sets/%d/group", groupSetID)
	result := new(model.Group)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListGroupSets returns group sets for a course.
func (s *Service) ListGroupSets(ctx context.Context, courseID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/courses/%d/group-sets", courseID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListGroups returns groups in a group set.
func (s *Service) ListGroups(ctx context.Context, groupSetID int, params map[string]string) (json.RawMessage, error) {
	u := addQueryParams(fmt.Sprintf("/api/group-sets/%d/groups", groupSetID), params)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetGroupOtherUsers returns users in the group set but outside the current group.
func (s *Service) GetGroupOtherUsers(ctx context.Context, groupSetID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d/other-users", groupSetID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetGroupHomeworkActivities returns grouped homework activities for a group set.
func (s *Service) GetGroupHomeworkActivities(ctx context.Context, groupSetID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d/group-homework-activities", groupSetID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetGroupExams returns grouped exams for a group set.
func (s *Service) GetGroupExams(ctx context.Context, groupSetID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d/group-exams", groupSetID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// UpdateGroup updates a group.
func (s *Service) UpdateGroup(ctx context.Context, groupID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/groups/%d", groupID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// AddGroupMembers adds members to a group.
func (s *Service) AddGroupMembers(ctx context.Context, groupID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/groups/%d/members", groupID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// --- Course Custom Score ---

// GetCourseCustomScoreItems returns custom score items for a course.
func (s *Service) GetCourseCustomScoreItems(ctx context.Context, courseID int) (*CourseCustomScoreItemsResponse, error) {
	u := fmt.Sprintf("/api/courses/%d/custom-score-items", courseID)
	result := new(CourseCustomScoreItemsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// UpdateCourseCustomScoreItems updates custom score items for a course.
func (s *Service) UpdateCourseCustomScoreItems(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/custom-score-items/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// CreateCourseCustomScoreItem creates a custom score item for a course.
func (s *Service) CreateCourseCustomScoreItem(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/custom-score-item", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// --- Access Code ---

// GetCourseAccessCode returns the access code for a course.
func (s *Service) GetCourseAccessCode(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/access-code/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Completion Criteria ---

// ListCompletionCriteria returns completion criteria.
func (s *Service) ListCompletionCriteria(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/completion-criteria", &result)
	return result, err
}

// --- Syllabus ---

// GetSyllabus returns a syllabus.
func (s *Service) GetSyllabus(ctx context.Context, syllabusID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/syllabus/%d", syllabusID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateSyllabus creates a new syllabus.
func (s *Service) CreateSyllabus(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/syllabus", body, &result)
	return result, err
}

// UpdateSyllabus updates a syllabus.
func (s *Service) UpdateSyllabus(ctx context.Context, syllabusID int, body interface{}) error {
	u := fmt.Sprintf("/api/syllabus/%d", syllabusID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// DeleteSyllabus deletes a syllabus.
func (s *Service) DeleteSyllabus(ctx context.Context, syllabusID int) error {
	return s.DeleteSyllabusWithOptions(ctx, syllabusID, nil)
}

// DeleteSyllabusWithOptions deletes a syllabus with optional query parameters.
func (s *Service) DeleteSyllabusWithOptions(ctx context.Context, syllabusID int, opts *DeleteSyllabusOptions) error {
	u := fmt.Sprintf("/api/syllabus/%d", syllabusID)
	if opts != nil && opts.DeleteRelatedActivity {
		u += "?delete_related_activity=true"
	}
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ResortSyllabus resorts syllabuses.
func (s *Service) ResortSyllabus(ctx context.Context, body interface{}) error {
	_, err := s.client.Put(ctx, "/api/syllabus/resort", body, nil)
	return err
}

// --- Feedback Activities ---

// GetFeedbackActivity returns a feedback activity.
func (s *Service) GetFeedbackActivity(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d", activityID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateFeedbackActivity creates a feedback activity.
func (s *Service) CreateFeedbackActivity(ctx context.Context, courseID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/feedback-activities", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// UpdateFeedbackActivity updates a feedback activity.
func (s *Service) UpdateFeedbackActivity(ctx context.Context, activityID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d", activityID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// GetFeedback returns a feedback.
func (s *Service) GetFeedback(ctx context.Context, feedbackID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedbacks/%d", feedbackID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// CreateFeedback creates feedback under a feedback activity.
func (s *Service) CreateFeedback(ctx context.Context, activityID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d/feedbacks", activityID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

// UpdateFeedback updates feedback under a feedback activity.
func (s *Service) UpdateFeedback(ctx context.Context, activityID, feedbackID int, body interface{}) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d/feedbacks/%d", activityID, feedbackID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// DeleteFeedback deletes a feedback record.
func (s *Service) DeleteFeedback(ctx context.Context, feedbackID int) error {
	u := fmt.Sprintf("/api/feedbacks/%d", feedbackID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
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

// SortQuestionnaireSubjects sorts questionnaire subjects.
func (s *Service) SortQuestionnaireSubjects(ctx context.Context, questionnaireID int, body interface{}) error {
	u := fmt.Sprintf("/api/questionnaire/%d/subject-sort", questionnaireID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
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

// BatchDeleteEntries batch deletes entries.
func (s *Service) BatchDeleteEntries(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/entries/batch-delete", body, nil)
	return err
}

// --- Knowledge Graph ---

// GetCourseKnowledgeGraph returns the knowledge graph for a course.
func (s *Service) GetCourseKnowledgeGraph(ctx context.Context, courseID int) (*KnowledgeNodesResponse, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-nodes", courseID)
	result := new(KnowledgeNodesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeTrees returns knowledge-node trees for a course.
func (s *Service) GetKnowledgeNodeTrees(ctx context.Context, courseID int) ([]*model.KnowledgeNode, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-nodes", courseID)
	result := new(KnowledgeNodesResponse)
	_, err := s.client.Get(ctx, u, result)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

// GetKnowledgeNode returns a knowledge node.
func (s *Service) GetKnowledgeNode(ctx context.Context, nodeID int) (*model.KnowledgeNode, error) {
	u := fmt.Sprintf("/api/knowledge-node/%d", nodeID)
	result := new(model.KnowledgeNode)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateKnowledgeNode creates a knowledge node in a course.
func (s *Service) CreateKnowledgeNode(ctx context.Context, courseID int, body interface{}) (*model.KnowledgeNode, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-node", courseID)
	result := new(model.KnowledgeNode)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateKnowledgeNode updates a knowledge node.
func (s *Service) UpdateKnowledgeNode(ctx context.Context, nodeID int, body interface{}) (*model.KnowledgeNode, error) {
	u := fmt.Sprintf("/api/knowledge-node/%d", nodeID)
	result := new(model.KnowledgeNode)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteKnowledgeNode deletes a knowledge node.
func (s *Service) DeleteKnowledgeNode(ctx context.Context, nodeID int) error {
	u := fmt.Sprintf("/api/knowledge-node/%d", nodeID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// BatchDeleteKnowledgeNodes deletes multiple knowledge nodes in a course.
func (s *Service) BatchDeleteKnowledgeNodes(ctx context.Context, courseID int, body *DeleteKnowledgeNodesRequest) error {
	u := fmt.Sprintf("/api/course/%d/knowledge-node/delete", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// AddKnowledgeNodeReferences adds references to a knowledge node.
func (s *Service) AddKnowledgeNodeReferences(ctx context.Context, nodeID int, body interface{}) error {
	u := fmt.Sprintf("/api/knowledge-node/%d/reference", nodeID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// DeleteKnowledgeNodeReferences deletes references from a knowledge node.
func (s *Service) DeleteKnowledgeNodeReferences(ctx context.Context, nodeID int, body interface{}) error {
	u := fmt.Sprintf("/api/knowledge-node/%d/reference/delete", nodeID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// CreateTeachingObjectives creates teaching objectives for a course.
func (s *Service) CreateTeachingObjectives(ctx context.Context, courseID int, body *TeachingObjectivesRequest) error {
	u := fmt.Sprintf("/api/course/%d/teaching-objective", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// UpdateTeachingObjectives updates teaching objectives for a course.
func (s *Service) UpdateTeachingObjectives(ctx context.Context, courseID int, body *TeachingObjectivesRequest) error {
	u := fmt.Sprintf("/api/course/%d/teaching-objective", courseID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// DeleteTeachingObjectives deletes teaching objectives for a course.
func (s *Service) DeleteTeachingObjectives(ctx context.Context, courseID int, body *DeleteTeachingObjectivesRequest) error {
	u := fmt.Sprintf("/api/course/%d/teaching-objective/delete", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// AddKnowledgeNodeRelation creates a relation for course knowledge nodes.
func (s *Service) AddKnowledgeNodeRelation(ctx context.Context, courseID int, body *KnowledgeNodeRelationRequest) (*KnowledgeNodeRelation, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-node/relation", courseID)
	result := new(KnowledgeNodeRelation)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateKnowledgeNodeRelation updates a relation for course knowledge nodes.
func (s *Service) UpdateKnowledgeNodeRelation(ctx context.Context, courseID int, body *KnowledgeNodeRelationRequest) (*KnowledgeNodeRelation, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-node/relation", courseID)
	result := new(KnowledgeNodeRelation)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteKnowledgeNodeRelations deletes relations for course knowledge nodes.
func (s *Service) DeleteKnowledgeNodeRelations(ctx context.Context, courseID int, body *DeleteKnowledgeNodeRelationsRequest) error {
	u := fmt.Sprintf("/api/course/%d/knowledge-node/relation/delete", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// ListCoursewareActivities returns courseware activities for knowledge graph association.
func (s *Service) ListCoursewareActivities(ctx context.Context, courseID int) ([]*CoursewareActivity, error) {
	u := fmt.Sprintf("/api/course/%d/courseware-activities", courseID)
	var result []*CoursewareActivity
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// MoveKnowledgeNode moves a knowledge node within a course graph.
func (s *Service) MoveKnowledgeNode(ctx context.Context, courseID int, body *MoveKnowledgeNodeRequest) error {
	u := fmt.Sprintf("/api/course/%d/knowledge-node/move", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// GetKnowledgeNodeRecommendedCaptures returns recommended captures for a knowledge node.
func (s *Service) GetKnowledgeNodeRecommendedCaptures(ctx context.Context, nodeID int, opts *model.ListOptions) (*KnowledgeNodeCapturesResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/knowledge-node/%d/recommended-captures", nodeID), opts)
	result := new(KnowledgeNodeCapturesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetMyCaptures returns current-user captures with pagination.
func (s *Service) GetMyCaptures(ctx context.Context, opts *model.ListOptions, params map[string]string) (*KnowledgeNodeCapturesResponse, error) {
	u := addListOptions("/api/my-captures", opts)
	if len(params) > 0 {
		u = addQueryParams(u, params)
	}
	result := new(KnowledgeNodeCapturesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetPublicCaptures returns public captures with pagination.
func (s *Service) GetPublicCaptures(ctx context.Context, opts *model.ListOptions, params map[string]string) (*KnowledgeNodeCapturesResponse, error) {
	u := addListOptions("/api/public-captures", opts)
	if len(params) > 0 {
		u = addQueryParams(u, params)
	}
	result := new(KnowledgeNodeCapturesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeRecommendedResourceReferences returns recommended resource references for a knowledge node.
func (s *Service) GetKnowledgeNodeRecommendedResourceReferences(ctx context.Context, nodeID int, opts *model.ListOptions) (*KnowledgeNodeResourcesResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/knowledge-node/%d/recommended-resource-references", nodeID), opts)
	result := new(KnowledgeNodeResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeStatisticsOverview returns statistics overview for a knowledge node.
func (s *Service) GetKnowledgeNodeStatisticsOverview(ctx context.Context, nodeID int) (*KnowledgeNodeOverview, error) {
	u := fmt.Sprintf("/api/knowledge-nodes/%d/statistics/overview", nodeID)
	result := new(KnowledgeNodeOverview)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeWeekStatsList returns timeline stats for a course knowledge graph.
func (s *Service) GetKnowledgeNodeWeekStatsList(ctx context.Context, courseID int) ([]*KnowledgeNodeWeekStat, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-stats-timeline", courseID)
	var result []*KnowledgeNodeWeekStat
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetKnowledgeNodeWeekStats returns a knowledge graph snapshot for a course.
func (s *Service) GetKnowledgeNodeWeekStats(ctx context.Context, courseID int, snapshotID int) (*KnowledgeGraphSnapshot, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-graph-snapshot/%d", courseID, snapshotID)
	result := new(KnowledgeGraphSnapshot)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeStatisticsSummary returns statistics summary for all knowledge nodes in a course.
func (s *Service) GetKnowledgeNodeStatisticsSummary(ctx context.Context, courseID int) (*KnowledgeNodeStatisticsSummary, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-nodes/statistics/summary", courseID)
	result := new(KnowledgeNodeStatisticsSummary)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeReferenceResource returns reference resources for a knowledge node.
func (s *Service) GetKnowledgeNodeReferenceResource(ctx context.Context, nodeID int) (*KnowledgeNodeReferenceResourcesResponse, error) {
	u := fmt.Sprintf("/api/knowledge-nodes/%d/reference-resource", nodeID)
	result := new(KnowledgeNodeReferenceResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeStatisticsStudentDetail returns student statistics detail for a knowledge node.
func (s *Service) GetKnowledgeNodeStatisticsStudentDetail(ctx context.Context, nodeID int, opts *model.ListOptions, conditions string) (*KnowledgeNodeStudentDetailsResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/knowledge-nodes/%d/statistics/student-detail", nodeID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(KnowledgeNodeStudentDetailsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeStatisticsResourceDetail returns resource statistics detail for a knowledge node.
func (s *Service) GetKnowledgeNodeStatisticsResourceDetail(ctx context.Context, nodeID int, opts *model.ListOptions, conditions string) (*KnowledgeNodeResourceDetailsResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/knowledge-nodes/%d/statistics/resource-detail", nodeID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(KnowledgeNodeResourceDetailsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeStatisticsActivityDetail returns activity statistics detail for a knowledge node.
func (s *Service) GetKnowledgeNodeStatisticsActivityDetail(ctx context.Context, nodeID int, opts *model.ListOptions, conditions string) (*KnowledgeNodeActivityDetailsResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/knowledge-nodes/%d/statistics/activity-detail", nodeID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(KnowledgeNodeActivityDetailsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeStudentResourceStat returns a student's resource statistics for a knowledge node.
func (s *Service) GetKnowledgeNodeStudentResourceStat(ctx context.Context, nodeID int, studentID int, opts *model.ListOptions, keyword string) (*KnowledgeNodeStudentResourcesResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/knowledge-nodes/%d/student/%d/resource/stat", nodeID, studentID), opts)
	if keyword != "" {
		u = addQueryParams(u, map[string]string{"keyword": keyword})
	}
	result := new(KnowledgeNodeStudentResourcesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeStudentActivityStat returns a student's activity statistics for a knowledge node.
func (s *Service) GetKnowledgeNodeStudentActivityStat(ctx context.Context, nodeID int, studentID int, opts *model.ListOptions, keyword string) (*KnowledgeNodeStudentActivitiesResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/knowledge-nodes/%d/student/%d/activity/stat", nodeID, studentID), opts)
	if keyword != "" {
		u = addQueryParams(u, map[string]string{"keyword": keyword})
	}
	result := new(KnowledgeNodeStudentActivitiesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeStudentStat returns overall knowledge-node statistics for a student in a course.
func (s *Service) GetKnowledgeNodeStudentStat(ctx context.Context, courseID int, studentID int) (*KnowledgeNodeStudentReferenceStat, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-nodes/student/%d/stats", courseID, studentID)
	result := new(KnowledgeNodeStudentReferenceStat)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetNodeBaseContentByStudent returns base node content statistics for a student.
func (s *Service) GetNodeBaseContentByStudent(ctx context.Context, nodeID int, studentID int) (*KnowledgeNodeStudentReferenceStat, error) {
	u := fmt.Sprintf("/api/knowledge-nodes/%d/student/%d/stat", nodeID, studentID)
	result := new(KnowledgeNodeStudentReferenceStat)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetCurrentStudentNodeOverallStatistics returns current student overall knowledge-node statistics in a course.
func (s *Service) GetCurrentStudentNodeOverallStatistics(ctx context.Context, courseID int, studentID int) (*KnowledgeNodeStudentOverallStatistics, error) {
	u := fmt.Sprintf("/api/courses/%d/students/%d/knowledge-nodes/overall-stats", courseID, studentID)
	result := new(KnowledgeNodeStudentOverallStatistics)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeReferTypeStatsByCourse returns knowledge node reference-type statistics for a course.
func (s *Service) GetKnowledgeNodeReferTypeStatsByCourse(ctx context.Context, courseID int) (*KnowledgeNodeReferTypeStats, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-nodes/statistics/refer-type", courseID)
	result := new(KnowledgeNodeReferTypeStats)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetKnowledgeNodeStatsByStudentDimension returns paginated student-dimension statistics for knowledge nodes in a course.
func (s *Service) GetKnowledgeNodeStatsByStudentDimension(ctx context.Context, courseID int, opts *model.ListOptions, conditions string) (*KnowledgeNodeStudentDimensionResponse, error) {
	u := addListOptions(fmt.Sprintf("/api/course/%d/knowledge-nodes/statistics/students-dimension", courseID), opts)
	if conditions != "" {
		u = addQueryParams(u, map[string]string{"conditions": conditions})
	}
	result := new(KnowledgeNodeStudentDimensionResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetCourseTeachingObjectives returns teaching objectives for a course.
func (s *Service) GetCourseTeachingObjectives(ctx context.Context, courseID int) ([]*TeachingObjective, error) {
	u := fmt.Sprintf("/api/course/%d/teaching-objective", courseID)
	var result []*TeachingObjective
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetNodeFacets returns facets for a knowledge node in a course.
func (s *Service) GetNodeFacets(ctx context.Context, courseID int, nodeID int) ([]*KnowledgeFacet, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-node/%d/facets", courseID, nodeID)
	var result []*KnowledgeFacet
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetFragments returns fragments for a knowledge node in a course.
func (s *Service) GetFragments(ctx context.Context, courseID int, nodeID int) ([]*KnowledgeFragment, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-node/%d/fragments", courseID, nodeID)
	var result []*KnowledgeFragment
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetKnowledgeReferencesForActivity returns knowledge references for an activity.
func (s *Service) GetKnowledgeReferencesForActivity(ctx context.Context, activityID int) ([]*ActivityKnowledgeReference, error) {
	u := fmt.Sprintf("/api/activities/%d/knowledge-references", activityID)
	var result []*ActivityKnowledgeReference
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// SaveKnowledgeReferencesForActivity saves knowledge references for an activity.
func (s *Service) SaveKnowledgeReferencesForActivity(ctx context.Context, activityID int, body interface{}) error {
	u := fmt.Sprintf("/api/activities/%d/knowledge-references", activityID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// RemoveMediaKnowledgeReferenceForActivity removes a media-fragment knowledge reference from an activity.
func (s *Service) RemoveMediaKnowledgeReferenceForActivity(ctx context.Context, activityID int, referenceID int, body *DeleteMediaKnowledgeReferenceRequest) error {
	u := fmt.Sprintf("/api/activities/%d/knowledge-references/%d", activityID, referenceID)
	if body != nil && body.ChapterID != 0 {
		u = addQueryParams(u, map[string]string{"chapter_id": strconv.Itoa(body.ChapterID)})
	}
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// RemoveUploadKnowledgeReferenceForActivity removes an upload knowledge reference from an activity.
func (s *Service) RemoveUploadKnowledgeReferenceForActivity(ctx context.Context, activityID int, referenceID int, body *DeleteUploadKnowledgeReferenceRequest) error {
	u := fmt.Sprintf("/api/activities/%d/knowledge-references/%d", activityID, referenceID)
	_, err := s.client.DeleteWithBody(ctx, u, body, nil)
	return err
}

// ParseKnowledgeNodesFromDocx parses knowledge nodes from a DOCX file.
func (s *Service) ParseKnowledgeNodesFromDocx(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/knowledge-nodes/parse/docx", body, &result)
	return result, err
}

// UpdateKnowledgeGraphStatus updates the publish status for a course knowledge graph.
func (s *Service) UpdateKnowledgeGraphStatus(ctx context.Context, courseID int, publishType string) error {
	u := fmt.Sprintf("/api/course/%d/knowledge-graph-status", courseID)
	_, err := s.client.Post(ctx, u, map[string]string{"publish_type": publishType}, nil)
	return err
}

// GetKfsSubjects returns available KFS subjects for knowledge-graph import.
func (s *Service) GetKfsSubjects(ctx context.Context) ([]*KnowledgeGraphKFSSubject, error) {
	var result []*KnowledgeGraphKFSSubject
	_, err := s.client.Get(ctx, "/api/knowledge-graph/kfs-subjects", &result)
	return result, err
}

// GetKfsImportInfo returns KFS import information for a course.
func (s *Service) GetKfsImportInfo(ctx context.Context, courseID int) (*KnowledgeGraphKFSImportInfo, error) {
	u := fmt.Sprintf("/api/knowledge-graph/courses/%d/kfs-import-info", courseID)
	result := new(KnowledgeGraphKFSImportInfo)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ImportKfsCourse imports a KFS course version into the target course knowledge graph.
func (s *Service) ImportKfsCourse(ctx context.Context, courseID int, body *ImportKfsCourseRequest) error {
	u := fmt.Sprintf("/api/knowledge-graph/courses/%d/kfs-course-import", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// GetKfsCourseForestVersionStats returns stats for a KFS forest version.
func (s *Service) GetKfsCourseForestVersionStats(ctx context.Context, kfsCourseID int, versionID int) (*KnowledgeGraphForestVersionStats, error) {
	u := fmt.Sprintf("/api/knowledge-graph/kfs-courses/%d/forest-versions/%d/stats", kfsCourseID, versionID)
	result := new(KnowledgeGraphForestVersionStats)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// BatchGetPublishedForestVersionByKFSCourseIDs returns published forest versions for KFS courses.
func (s *Service) BatchGetPublishedForestVersionByKFSCourseIDs(ctx context.Context, ids []int) ([]*KnowledgeGraphPublishedForestVersion, error) {
	u := addQueryParams("/api/knowledge-graph/kfs-courses/-/published-forest-versions:batchGet", map[string]string{
		"ids": intsToCSV(ids),
	})
	var result []*KnowledgeGraphPublishedForestVersion
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// BatchGetForestVersionStatsByKFSVersionIDs returns forest-version stats for KFS version IDs.
func (s *Service) BatchGetForestVersionStatsByKFSVersionIDs(ctx context.Context, ids []int) ([]*KnowledgeGraphForestVersionStatsItem, error) {
	u := addQueryParams("/api/knowledge-graph/forest-versions/-/stats:batchGet", map[string]string{
		"ids": intsToCSV(ids),
	})
	var result []*KnowledgeGraphForestVersionStatsItem
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetChinamCloudKnowledgeGraphDiff returns diff entries between local and ChinamCloud knowledge graphs.
func (s *Service) GetChinamCloudKnowledgeGraphDiff(ctx context.Context, courseID int) ([]*KnowledgeGraphDiff, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-graphs/diff", courseID)
	var result []*KnowledgeGraphDiff
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetKnowledgeGraphSimilarity returns similarity candidates for a course knowledge graph.
func (s *Service) GetKnowledgeGraphSimilarity(ctx context.Context, courseID int, body interface{}) ([]*KnowledgeGraphSimilarity, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-graphs/similarity", courseID)
	var payload struct {
		Similarities []*KnowledgeGraphSimilarity `json:"similarities"`
	}
	_, err := s.client.Post(ctx, u, body, &payload)
	return payload.Similarities, err
}

// SyncChinamCloudKnowledgeGraph starts syncing the course knowledge graph with ChinamCloud.
func (s *Service) SyncChinamCloudKnowledgeGraph(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/course/%d/knowledge-graph/sync", courseID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// EditKnowledgeGraphSource updates the source of a course knowledge graph.
func (s *Service) EditKnowledgeGraphSource(ctx context.Context, courseID int, body *UpdateKnowledgeGraphSourceRequest) error {
	u := fmt.Sprintf("/api/course/%d/knowledge-graph/source", courseID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// SyncChinamCloudKnowledgeGraphDiff syncs selected diff entries back to the course knowledge graph.
func (s *Service) SyncChinamCloudKnowledgeGraphDiff(ctx context.Context, courseID int, body interface{}) error {
	u := fmt.Sprintf("/api/course/%d/knowledge-graphs/diff/sync", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// ReplaceChinamKnowledgeGraph replaces the course knowledge graph from ChinamCloud.
func (s *Service) ReplaceChinamKnowledgeGraph(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/course/%d/knowledge-graph/replace", courseID)
	_, err := s.client.Put(ctx, u, nil, nil)
	return err
}

// ImportKnowledgeNodes posts structured knowledge-node data into a course.
func (s *Service) ImportKnowledgeNodes(ctx context.Context, courseID int, body interface{}) error {
	u := fmt.Sprintf("/api/courses/%d/knowledge-nodes/import", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// ImportKnowledgeNodesByCourse imports knowledge nodes from another course.
func (s *Service) ImportKnowledgeNodesByCourse(ctx context.Context, courseID int, body *ImportKnowledgeNodesByCourseRequest) error {
	u := fmt.Sprintf("/api/courses/%d/knowledge-nodes/import", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// CancelPreview cancels a course preview.
func (s *Service) CancelPreview(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/courses/%d/preview", courseID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// GetKnowledgeGraphEmbedURL returns the embedded cluster-graph URL for a course.
func (s *Service) GetKnowledgeGraphEmbedURL(ctx context.Context, courseID int) (string, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-graph/embed/cluster-graph", courseID)
	var result struct {
		URL string `json:"url"`
	}
	_, err := s.client.Get(ctx, u, &result)
	return result.URL, err
}

// GetChinamCloudGraphEditURL returns the edit URL for a course knowledge graph.
func (s *Service) GetChinamCloudGraphEditURL(ctx context.Context, courseID int) (string, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-graph/edit-url", courseID)
	var result struct {
		URL string `json:"url"`
	}
	_, err := s.client.Get(ctx, u, &result)
	return result.URL, err
}

// GetChinamCloudResourceViewURL returns the preview URL for an external resource.
func (s *Service) GetChinamCloudResourceViewURL(ctx context.Context, courseID int, resourceID int) (string, error) {
	u := fmt.Sprintf("/api/course/%d/external-resource/%d/preview-url", courseID, resourceID)
	var result struct {
		URL string `json:"url"`
	}
	_, err := s.client.Get(ctx, u, &result)
	return result.URL, err
}

// GetFileStatus returns the processing status of an upload file.
func (s *Service) GetFileStatus(ctx context.Context, uploadID int) (string, error) {
	u := fmt.Sprintf("/api/uploads/%d", uploadID)
	var result struct {
		Status string `json:"status"`
	}
	_, err := s.client.Get(ctx, u, &result)
	return result.Status, err
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

// --- Warning ---

// GetStudentWarning returns warning info for a student.
func (s *Service) GetStudentWarning(ctx context.Context, studentID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/warning/student/%d", studentID)
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

// GetQuestionnaireSubjects returns questionnaire subjects.
func (s *Service) GetQuestionnaireSubjects(ctx context.Context, questionnaireID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/questionnaire/%d/subjects", questionnaireID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Projects ---

// ListProjects returns projects.
func (s *Service) ListProjects(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/projects", &result)
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

// --- Statistic Resource Audit ---

// GetStatisticResourceAudit returns statistic resource audit.
func (s *Service) GetStatisticResourceAudit(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/courses/statistic/resource-audit", &result)
	return result, err
}

// --- TPDOE ---

// GetTPDOEStatStudents returns TPDOE student statistics.
func (s *Service) GetTPDOEStatStudents(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/tpdoe/stat-students?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Inspect Child ---

// InspectChild inspects a child course.
func (s *Service) InspectChild(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/inspect-child/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Blueprint Sub Courses ---

// ListBlueprintSubCourses returns blueprint sub courses.
func (s *Service) ListBlueprintSubCourses(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blueprint/sub-courses/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// DeleteBlueprint deletes a blueprint course mapping/configuration.
func (s *Service) DeleteBlueprint(ctx context.Context, courseID int) error {
	u := fmt.Sprintf("/api/blueprint/%d", courseID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// CancelBlueprintActivitySync cancels blueprint sync for an activity.
func (s *Service) CancelBlueprintActivitySync(ctx context.Context, courseID, activityID int, body interface{}) error {
	u := fmt.Sprintf("/api/blueprint/%d/activities/%d/cancel-sync", courseID, activityID)
	_, err := s.client.Delete(ctx, u, body)
	return err
}

// GetBlueprintSubmittedInfo returns blueprint submitted sync info for a target object.
func (s *Service) GetBlueprintSubmittedInfo(ctx context.Context, courseID int, resourceType string, resourceID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blueprint/%d/%s/%d/submitted-info", courseID, resourceType, resourceID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// SyncBlueprintSubject syncs a blueprint subject item.
func (s *Service) SyncBlueprintSubject(ctx context.Context, courseID int, resourceType string, resourceID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blueprint/%d/%s/%d/sync-subject", courseID, resourceType, resourceID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, nil, &result)
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

// --- Course Classifications ---

// ListCourseClassifications returns course classifications.
func (s *Service) ListCourseClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/course-classifications", &result)
	return result, err
}

// --- Curriculum Classifications ---

// ListCurriculumClassifications returns curriculum classifications.
func (s *Service) ListCurriculumClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/curriculum-classifications", &result)
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

// --- Studio ---

// GetStudio returns a studio.
func (s *Service) GetStudio(ctx context.Context, studioID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/studio/%d", studioID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Knowledge Capture/Resource Visit ---

// RecordKnowledgeCaptureVisit records a knowledge capture visit.
func (s *Service) RecordKnowledgeCaptureVisit(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/knowledge-capture-visit", body, nil)
	return err
}

// RecordKnowledgeResourceVisit records a knowledge resource visit.
func (s *Service) RecordKnowledgeResourceVisit(ctx context.Context, body interface{}) error {
	_, err := s.client.Post(ctx, "/api/knowledge-resource-visit", body, nil)
	return err
}

// --- Sign In ---

// GetCourseSignIn returns the sign-in for a course.
func (s *Service) GetCourseSignIn(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/sign-in", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

func intsToCSV(ids []int) string {
	if len(ids) == 0 {
		return ""
	}
	parts := make([]string, len(ids))
	for i, id := range ids {
		parts[i] = strconv.Itoa(id)
	}
	return strings.Join(parts, ",")
}
