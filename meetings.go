package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// --- Meeting & Live Operations (attached to various services) ---

// Meeting related methods on CoursesService

// GetMeeting returns a meeting.
func (s *CoursesService) GetMeeting(ctx context.Context, meetingID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/meeting/%d", meetingID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetMeetingWeekTimePeriods returns meeting week time periods.
func (s *CoursesService) GetMeetingWeekTimePeriods(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/meeting/week/time-periods", &result)
	return result, err
}

// GetShanghaiTechMeeting returns a ShanghaiTech meeting.
func (s *CoursesService) GetShanghaiTechMeeting(ctx context.Context, meetingID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/meeting/shanghaitech/%d", meetingID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Tencent Meeting ---

// GetTencentMeeting returns a Tencent meeting.
func (s *CoursesService) GetTencentMeeting(ctx context.Context, meetingID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/tencent-meeting/%d", meetingID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetTencentMeetingAuthURL returns the Tencent meeting authorization URL.
func (s *CoursesService) GetTencentMeetingAuthURL(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/tencent-meeting/authorization-url", &result)
	return result, err
}

// CheckTencentMeetingUserAuth checks Tencent meeting user auth.
func (s *CoursesService) CheckTencentMeetingUserAuth(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/tencent_meeting/check-user-auth", &result)
	return result, err
}

// ListTencentMeetingActivities returns Tencent meeting activities for a course.
func (s *CoursesService) ListTencentMeetingActivities(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/tencent-meeting/activities?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- DingTalk ---

// GetDingTalkChat returns DingTalk chat info for a course.
func (s *CoursesService) GetDingTalkChat(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/ding-talk/chat?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetDingTalkUserID returns the DingTalk user ID.
func (s *CoursesService) GetDingTalkUserID(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/ding-talk/user-id", &result)
	return result, err
}

// GetDingTalkLive returns a DingTalk live session.
func (s *CoursesService) GetDingTalkLive(ctx context.Context, liveID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/dingtalk-lives/%d", liveID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Live Activities ---

// GetLiveActivity returns a live activity.
func (s *CoursesService) GetLiveActivity(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/live-activities/%d", activityID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetLiveRecord returns a live record.
func (s *CoursesService) GetLiveRecord(ctx context.Context, recordID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/live-records/%d", recordID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetLectureLiveSchedule returns a lecture live schedule.
func (s *CoursesService) GetLectureLiveSchedule(ctx context.Context, scheduleID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/lecture-live/schedule/%d", scheduleID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetLectureLive returns a lecture live session.
func (s *CoursesService) GetLectureLive(ctx context.Context, jwt string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/lecture-live?jwt=%s", jwt)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetLectureLiveActivity returns a lecture live activity for a course.
func (s *CoursesService) GetLectureLiveActivity(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/lecture-live-activity/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- VTRS (Virtual Teaching Room System) ---

// ListVTRSes returns VTRS entries.
func (s *CoursesService) ListVTRSes(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/vtrses", &result)
	return result, err
}

// GetVTRS returns a specific VTRS entry.
func (s *CoursesService) GetVTRS(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/%d", vtrsID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetVTRSAccessCode gets the access code for a VTRS.
func (s *CoursesService) GetVTRSAccessCode(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/access-code/%d", vtrsID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListVTRSMeetingClassifications returns VTRS meeting classifications.
func (s *CoursesService) ListVTRSMeetingClassifications(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/meetings/classifications/%d", vtrsID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListVTRSResourceClassifications returns VTRS resource classifications.
func (s *CoursesService) ListVTRSResourceClassifications(ctx context.Context, vtrsID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/vtrses/resources/classifications/%d", vtrsID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ShareVTRSResources shares VTRS resources.
func (s *CoursesService) ShareVTRSResources(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/vtrses/share-resources", body, nil)
	return err
}

// ListVTRSSubjectLibs returns VTRS subject libraries.
func (s *CoursesService) ListVTRSSubjectLibs(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/vtrses/subject-libs", &result)
	return result, err
}

// --- Instruction Team Meeting ---

// GetInstructionTeamMeeting returns instruction team meeting.
func (s *CoursesService) GetInstructionTeamMeeting(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/instruction-team/meeting", &result)
	return result, err
}

// --- Combine Courses ---

// ListCombineCourses returns combined courses.
func (s *CoursesService) ListCombineCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/combine-courses", &result)
	return result, err
}

// CreateCombineCourse creates a combined course.
func (s *CoursesService) CreateCombineCourse(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/combine-courses", body, &result)
	return result, err
}

// DeleteCombineCourse deletes a combined course.
func (s *CoursesService) DeleteCombineCourse(ctx context.Context, combineID int) error {
	u := fmt.Sprintf("/api/combine-courses/%d", combineID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// --- Danmu (Bullet Screen) ---

// GetDanmu returns danmu for a course.
func (s *CoursesService) GetDanmu(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/danmu/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Interactions ---

// GetInteraction returns a specific interaction.
func (s *CoursesService) GetInteraction(ctx context.Context, interactionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/interactions/%d", interactionID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// VoteInteraction votes on an interaction.
func (s *CoursesService) VoteInteraction(ctx context.Context, interactionID int, body interface{}) error {
	u := fmt.Sprintf("/api/courses/interactions/vote/%d", interactionID)
	_, err := s.client.post(ctx, u, body, nil)
	return err
}

// --- Interaction Activities ---

// GetInteractionActivity returns an interaction activity.
func (s *CoursesService) GetInteractionActivity(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/interaction-activities/%d", activityID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetInteractionSubmission returns an interaction submission.
func (s *CoursesService) GetInteractionSubmission(ctx context.Context, submissionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/interaction-submissions/%d", submissionID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Rollcall ---

// GetRollcall returns a rollcall.
func (s *CoursesService) GetRollcall(ctx context.Context, rollcallID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/rollcall/%d", rollcallID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetMergedRollcall returns a merged rollcall.
func (s *CoursesService) GetMergedRollcall(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/rollcall/merged-rollcall", &result)
	return result, err
}

// GetMergedRollcallStudentRollcalls returns student rollcalls from merged rollcall.
func (s *CoursesService) GetMergedRollcallStudentRollcalls(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/rollcall/merged-rollcall/student-rollcalls", &result)
	return result, err
}

// GetRollcallStatus returns rollcall status for a course.
func (s *CoursesService) GetRollcallStatus(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/rollcall_status/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListTimetableRollcalls returns timetable rollcalls.
func (s *CoursesService) ListTimetableRollcalls(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/timetable_rollcalls", &result)
	return result, err
}

// --- Groups ---

// GetGroupSet returns a group set.
func (s *CoursesService) GetGroupSet(ctx context.Context, groupSetID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/group-sets/%d", groupSetID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetGroup returns a group.
func (s *CoursesService) GetGroup(ctx context.Context, groupID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/groups/%d", groupID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Course Custom Score ---

// GetCourseCustomScoreItems returns custom score items for a course.
func (s *CoursesService) GetCourseCustomScoreItems(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/custom-score-items/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Access Code ---

// GetCourseAccessCode returns the access code for a course.
func (s *CoursesService) GetCourseAccessCode(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/access-code/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Completion Criteria ---

// ListCompletionCriteria returns completion criteria.
func (s *CoursesService) ListCompletionCriteria(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/completion-criteria", &result)
	return result, err
}

// --- Syllabus ---

// GetSyllabus returns a syllabus.
func (s *CoursesService) GetSyllabus(ctx context.Context, syllabusID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/syllabus/%d", syllabusID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// CreateSyllabus creates a new syllabus.
func (s *CoursesService) CreateSyllabus(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/syllabus", body, &result)
	return result, err
}

// UpdateSyllabus updates a syllabus.
func (s *CoursesService) UpdateSyllabus(ctx context.Context, syllabusID int, body interface{}) error {
	u := fmt.Sprintf("/api/syllabus/%d", syllabusID)
	_, err := s.client.put(ctx, u, body, nil)
	return err
}

// DeleteSyllabus deletes a syllabus.
func (s *CoursesService) DeleteSyllabus(ctx context.Context, syllabusID int) error {
	u := fmt.Sprintf("/api/syllabuses/%d", syllabusID)
	_, err := s.client.delete(ctx, u, nil)
	return err
}

// ResortSyllabus resorts syllabuses.
func (s *CoursesService) ResortSyllabus(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/syllabus/resort", body, nil)
	return err
}

// --- Feedback Activities ---

// GetFeedbackActivity returns a feedback activity.
func (s *CoursesService) GetFeedbackActivity(ctx context.Context, activityID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedback-activities/%d", activityID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetFeedback returns a feedback.
func (s *CoursesService) GetFeedback(ctx context.Context, feedbackID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/feedbacks/%d", feedbackID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Questionnaires ---

// GetQuestionnaire returns a questionnaire.
func (s *CoursesService) GetQuestionnaire(ctx context.Context, questionnaireID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/questionnaire/%d", questionnaireID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListQuestionnaires returns questionnaires.
func (s *CoursesService) ListQuestionnaires(ctx context.Context, questionnaireID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/questionnaires/%d", questionnaireID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Ask Questions ---

// GetAskQuestion returns a question.
func (s *CoursesService) GetAskQuestion(ctx context.Context, questionID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/ask-questions/%d", questionID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Entries ---

// ListEntries returns entries.
func (s *CoursesService) ListEntries(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/entries", &result)
	return result, err
}

// CreateEntry creates an entry.
func (s *CoursesService) CreateEntry(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/entries", body, &result)
	return result, err
}

// GetEntry returns an entry.
func (s *CoursesService) GetEntry(ctx context.Context, entryID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/entries/%d", entryID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// BatchDeleteEntries batch deletes entries.
func (s *CoursesService) BatchDeleteEntries(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/entries/batch-delete", body, nil)
	return err
}

// --- Knowledge Graph ---

// GetCourseKnowledgeGraph returns the knowledge graph for a course.
func (s *CoursesService) GetCourseKnowledgeGraph(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/knowledge-graph/courses/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetKnowledgeNode returns a knowledge node.
func (s *CoursesService) GetKnowledgeNode(ctx context.Context, nodeID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/knowledge-node/%d", nodeID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ParseKnowledgeNodesFromDocx parses knowledge nodes from a DOCX file.
func (s *CoursesService) ParseKnowledgeNodesFromDocx(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/knowledge-nodes/parse/docx", body, &result)
	return result, err
}

// --- Lessons ---

// GetLesson returns a lesson.
func (s *CoursesService) GetLesson(ctx context.Context, lessonID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/lessons/%d", lessonID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetLessonManagement returns lesson management info.
func (s *CoursesService) GetLessonManagement(ctx context.Context, lessonID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/lessons_management/%d", lessonID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListLessonRooms returns lesson rooms.
func (s *CoursesService) ListLessonRooms(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/lesson-rooms", &result)
	return result, err
}

// ListRoomLocations returns room locations.
func (s *CoursesService) ListRoomLocations(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/room-locations", &result)
	return result, err
}

// --- Warning ---

// GetStudentWarning returns warning info for a student.
func (s *CoursesService) GetStudentWarning(ctx context.Context, studentID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/warning/student/%d", studentID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Chatrooms ---

// GetChatroom returns a chatroom.
func (s *CoursesService) GetChatroom(ctx context.Context, chatroomID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/chatrooms/%d", chatroomID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Cloud Classroom ---

// GetCloudClassroom returns cloud classroom info.
func (s *CoursesService) GetCloudClassroom(ctx context.Context, start string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/activity/cloud-classroom?start=%s", start)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Projects ---

// ListProjects returns projects.
func (s *CoursesService) ListProjects(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/projects", &result)
	return result, err
}

// GetProject returns a project.
func (s *CoursesService) GetProject(ctx context.Context, projectID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/project/%d", projectID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// CreateProject creates a project.
func (s *CoursesService) CreateProject(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/project", body, &result)
	return result, err
}

// --- Statistic Resource Audit ---

// GetStatisticResourceAudit returns statistic resource audit.
func (s *CoursesService) GetStatisticResourceAudit(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/courses/statistic/resource-audit", &result)
	return result, err
}

// --- TPDOE ---

// GetTPDOEStatStudents returns TPDOE student statistics.
func (s *CoursesService) GetTPDOEStatStudents(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/tpdoe/stat-students?course_id=%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Inspect Child ---

// InspectChild inspects a child course.
func (s *CoursesService) InspectChild(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/inspect-child/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Blueprint Sub Courses ---

// ListBlueprintSubCourses returns blueprint sub courses.
func (s *CoursesService) ListBlueprintSubCourses(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blueprint/sub-courses/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Lark Files ---

// ListLarkFiles returns Lark files.
func (s *CoursesService) ListLarkFiles(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/lark/files", &result)
	return result, err
}

// --- WeDrive ---

// GetWeDriveFile returns a WeDrive file.
func (s *CoursesService) GetWeDriveFile(ctx context.Context, fileID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/wedrive/file/%d", fileID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListWeDriveFiles returns WeDrive files with pagination.
func (s *CoursesService) ListWeDriveFiles(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/wedrive/files", opts)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- BlobStorage ---

// GetBlobStorageOpenClientURL returns blob storage open client URL.
func (s *CoursesService) GetBlobStorageOpenClientURL(ctx context.Context, parentID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/blobstorage/open-client-url?parent_id=%d", parentID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- ChinamCloud ---

// ListChinamCloudResources returns ChinamCloud resources.
func (s *CoursesService) ListChinamCloudResources(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/chinamcloud/resources", &result)
	return result, err
}

// UploadChinamCloud uploads to ChinamCloud.
func (s *CoursesService) UploadChinamCloud(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/chinamcloud/upload", body, &result)
	return result, err
}

// --- Campus Subject Lib ---

// ListCampusSubjectLibClassifications returns campus subject lib classifications.
func (s *CoursesService) ListCampusSubjectLibClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/campus-subject-lib/classifications", &result)
	return result, err
}

// GetCampusSubjectLibSubjectCount returns campus subject lib subject count by classifications.
func (s *CoursesService) GetCampusSubjectLibSubjectCount(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/campus-subject-lib/classifications/subject-count", &result)
	return result, err
}

// --- Course Classifications ---

// ListCourseClassifications returns course classifications.
func (s *CoursesService) ListCourseClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/course-classifications", &result)
	return result, err
}

// --- Curriculum Classifications ---

// ListCurriculumClassifications returns curriculum classifications.
func (s *CoursesService) ListCurriculumClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/curriculum-classifications", &result)
	return result, err
}

// --- OBE ---

// GetExistedMetrics returns existed OBE metrics.
func (s *CoursesService) GetExistedMetrics(ctx context.Context, params string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/obe/existed-metrics?params=%s", params)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Program ---

// ListCoursePrograms returns course programs.
func (s *CoursesService) ListCoursePrograms(ctx context.Context, departmentIDs string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/program/course-programs?department_ids=%s", departmentIDs)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// ListUserPrograms returns user programs.
func (s *CoursesService) ListUserPrograms(ctx context.Context, fields string) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/program/user-programs?fields=%s", fields)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Studio ---

// GetStudio returns a studio.
func (s *CoursesService) GetStudio(ctx context.Context, studioID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/studio/%d", studioID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Knowledge Capture/Resource Visit ---

// RecordKnowledgeCaptureVisit records a knowledge capture visit.
func (s *CoursesService) RecordKnowledgeCaptureVisit(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/knowledge-capture-visit", body, nil)
	return err
}

// RecordKnowledgeResourceVisit records a knowledge resource visit.
func (s *CoursesService) RecordKnowledgeResourceVisit(ctx context.Context, body interface{}) error {
	_, err := s.client.post(ctx, "/api/knowledge-resource-visit", body, nil)
	return err
}

// --- Sign In ---

// GetCourseSignIn returns the sign-in for a course.
func (s *CoursesService) GetCourseSignIn(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/courses/%d/sign-in", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}
