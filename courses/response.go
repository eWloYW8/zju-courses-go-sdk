package courses

import (
	"github.com/eWloYW8/zju-courses-go-sdk/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

type OutlineResponse = Outline

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

type ExamsResponse struct {
	Exams []*Exam `json:"exams"`
}

type ExamScoresResponse struct {
	ExamScores []*ExamScore `json:"exam_scores"`
}

type ClassroomListResponse struct {
	Classrooms []*Classroom `json:"classrooms"`
}

type TopicCategoriesResponse struct {
	TopicCategories []*TopicCategory `json:"topic_categories"`
}

type UsersSmallAvatarsResponse struct {
	Avatars map[string]string `json:"avatars"`
}

type InteractionsResponse struct {
	Interactions []*Interaction `json:"interactions"`
}

type LiveRecordsResponse struct {
	LiveRecords []*LiveRecord `json:"live_records"`
}

type RollcallsResponse struct {
	Rollcalls []*Rollcall `json:"rollcalls"`
}

type SubmittedExamsResponse struct {
	ExamIDs []int `json:"exam_ids"`
}

type ActivityReadsForUserResponse struct {
	ActivityReads []*activities.ActivityRead `json:"activity_reads"`
}

type EntryRecordResponse map[string]any

type OnlineVideoCompletenessSettingResponse map[string]any

type CourseListResponse map[string]any

type CourseCountResponse map[string]any

type SettingsResponse map[string]any

type BlueprintSubItemsResultResponse map[string]any
