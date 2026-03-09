package homework

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles homework-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Student Homework ---

// GetHomeworkScore returns the homework score for a student.
func (s *Service) GetHomeworkScore(ctx context.Context, activityID, studentID int) (*HomeworkScore, error) {
	u := fmt.Sprintf("/api/activities/%d/students/%d/homework-score", activityID, studentID)
	result := new(HomeworkScore)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetGroupHomeworkScore returns the homework score for a group.
func (s *Service) GetGroupHomeworkScore(ctx context.Context, activityID, groupID int) (*HomeworkScore, error) {
	u := fmt.Sprintf("/api/activities/%d/groups/%d/scores", activityID, groupID)
	result := new(HomeworkScore)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListSubmissions returns submissions for a student on a homework activity.
func (s *Service) ListSubmissions(ctx context.Context, activityID, studentID int) (*SubmissionListResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/students/%d/submission_list", activityID, studentID)
	result := new(SubmissionListResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetStudentSubmission returns the current submission detail for a student on a homework activity.
func (s *Service) GetStudentSubmission(ctx context.Context, activityID, studentID int) (*Submission, error) {
	u := fmt.Sprintf("/api/course/activities/%d/students/%d/submission", activityID, studentID)
	result := new(Submission)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListGroupSubmissions returns submissions for a group on a homework activity.
func (s *Service) ListGroupSubmissions(ctx context.Context, activityID, groupID int) (*SubmissionListResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/groups/%d/submission_list", activityID, groupID)
	result := new(SubmissionListResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetGroupSubmission returns the current submission detail for a group on a homework activity.
func (s *Service) GetGroupSubmission(ctx context.Context, activityID, groupID int) (*Submission, error) {
	u := fmt.Sprintf("/api/activities/%d/groups/%d/submission", activityID, groupID)
	result := new(Submission)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListStudentSubmissionRecords returns instructor-facing student submissions for a homework activity.
func (s *Service) ListStudentSubmissionRecords(ctx context.Context, homeworkID int, params *ListSubmissionRecordsParams) (*SubmissionRecordsResponse, error) {
	u := fmt.Sprintf("/api/homework/%d/student-submissions", homeworkID)
	u = addHomeworkQueryParams(u, params)
	result := new(SubmissionRecordsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListGroupSubmissionRecords returns instructor-facing group submissions for a homework activity.
func (s *Service) ListGroupSubmissionRecords(ctx context.Context, homeworkID int, needUploadsSize bool) (*SubmissionRecordsResponse, error) {
	u := fmt.Sprintf("/api/homework/%d/group-submissions", homeworkID)
	u = addHomeworkQueryParams(u, &ListSubmissionRecordsParams{NeedUploadsSize: &needUploadsSize})
	result := new(SubmissionRecordsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetMakeUpRecord returns the make-up record for a student.
func (s *Service) GetMakeUpRecord(ctx context.Context, activityID, studentID int) (MakeUpRecordResponse, error) {
	u := fmt.Sprintf("/api/homework/%d/students/%d/make-up-record", activityID, studentID)
	result := make(MakeUpRecordResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetResubmitRecord returns the resubmit record for a student.
func (s *Service) GetResubmitRecord(ctx context.Context, activityID, studentID int) (ResubmitRecordResponse, error) {
	u := fmt.Sprintf("/api/homework/%d/students/%d/resubmit-record", activityID, studentID)
	result := make(ResubmitRecordResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListMakeUpRecords returns make-up records for a homework activity.
func (s *Service) ListMakeUpRecords(ctx context.Context, homeworkID int, userIDs []int) (*MakeUpRecordsResponse, error) {
	u := fmt.Sprintf("/api/homework/%d/make-up-records", homeworkID)
	u = addUserIDsQuery(u, userIDs)
	result := new(MakeUpRecordsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListResubmitRecords returns resubmit records for a homework activity.
func (s *Service) ListResubmitRecords(ctx context.Context, homeworkID int, userIDs []int) (*ResubmitRecordsResponse, error) {
	u := fmt.Sprintf("/api/homework/%d/resubmit-records", homeworkID)
	u = addUserIDsQuery(u, userIDs)
	result := new(ResubmitRecordsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Homework Submission ---

// SubmitHomework submits homework for a student.
func (s *Service) SubmitHomework(ctx context.Context, activityID int, body *SubmitHomeworkRequest) (*Submission, error) {
	u := fmt.Sprintf("/api/homework/submission/%d", activityID)
	result := new(Submission)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateSubmission updates a homework submission.
func (s *Service) UpdateSubmission(ctx context.Context, submissionID int, body *UpdateSubmissionRequest) (*Submission, error) {
	u := fmt.Sprintf("/api/submissions/%d", submissionID)
	result := new(Submission)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// ListMarkedAttachments returns marked attachments for a submission.
func (s *Service) ListMarkedAttachments(ctx context.Context, submissionID int) (MarkedAttachmentsResponse, error) {
	u := fmt.Sprintf("/api/submissions/%d/marked_attachments", submissionID)
	result := new(MarkedAttachmentsResponse)
	_, err := s.client.Get(ctx, u, result)
	if err != nil {
		return MarkedAttachmentsResponse{}, err
	}
	return *result, nil
}

// GetMarkedAttachment returns a marked attachment for a submission.
func (s *Service) GetMarkedAttachment(ctx context.Context, submissionID, attachmentID int) (MarkedAttachmentResponse, error) {
	u := fmt.Sprintf("/api/submissions/%d/marked_attachments/%d", submissionID, attachmentID)
	result := new(MarkedAttachmentResponse)
	_, err := s.client.Get(ctx, u, result)
	if err != nil {
		return MarkedAttachmentResponse{}, err
	}
	return *result, nil
}

// RecommendSubmission marks submissions as recommended.
func (s *Service) RecommendSubmission(ctx context.Context, submissionIDs []int) error {
	_, err := s.client.Put(ctx, "/api/submission/recommend", &RecommendSubmissionRequest{SubmissionIDs: submissionIDs}, nil)
	return err
}

// MarkSubmissionRead marks a submission as read.
func (s *Service) MarkSubmissionRead(ctx context.Context, submissionID int) error {
	u := fmt.Sprintf("/api/submissions/%d/read", submissionID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// UpdateMarkedSubmitted updates the marked-submitted state for homework submissions.
func (s *Service) UpdateMarkedSubmitted(ctx context.Context, activityID int, body *MarkedSubmittedRequest) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/activities/%d/submission/marked_submitted", activityID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, body, &result)
	return result, err
}

// --- Homework Management (Instructor) ---

// CreateHomework creates a new homework activity.
func (s *Service) CreateHomework(ctx context.Context, courseID int, homework *CreateHomeworkRequest) (*activities.Activity, error) {
	u := fmt.Sprintf("/api/homeworks/%d", courseID)
	result := new(activities.Activity)
	_, err := s.client.Post(ctx, u, homework, result)
	return result, err
}

// UpdateHomework updates a homework activity.
func (s *Service) UpdateHomework(ctx context.Context, homeworkID int, homework *UpdateHomeworkRequest) (*activities.Activity, error) {
	u := fmt.Sprintf("/api/homework/%d", homeworkID)
	result := new(activities.Activity)
	_, err := s.client.Put(ctx, u, homework, result)
	return result, err
}

// DeleteHomework deletes a homework activity.
func (s *Service) DeleteHomework(ctx context.Context, homeworkID int) error {
	u := fmt.Sprintf("/api/homework/%d", homeworkID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ScoreSubmission scores a submission (instructor).
func (s *Service) ScoreSubmission(ctx context.Context, submissionID int, body *ScoreSubmissionRequest) error {
	u := fmt.Sprintf("/api/submissions/%d", submissionID)
	_, err := s.client.Patch(ctx, u, body, nil)
	return err
}

// ListInterScoreSubmissions returns inter-review score submissions.
func (s *Service) ListInterScoreSubmissions(ctx context.Context, activityID int) ([]InterScoreSubmission, error) {
	u := fmt.Sprintf("/api/inter-score-submissions/%d", activityID)
	var result []InterScoreSubmission
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListInterScores returns inter-review scores.
func (s *Service) ListInterScores(ctx context.Context, activityID int) ([]InterScore, error) {
	u := fmt.Sprintf("/api/homework/%d/inter-scores", activityID)
	var result []InterScore
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListIntraScores returns intra-review scores.
func (s *Service) ListIntraScores(ctx context.Context, activityID int) ([]InterScore, error) {
	u := fmt.Sprintf("/api/homework/%d/intra-scores", activityID)
	var result []InterScore
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListInterScoresByUserIDs returns inter-review scores filtered by student IDs.
func (s *Service) ListInterScoresByUserIDs(ctx context.Context, activityID int, userIDs []int) ([]InterScore, error) {
	u := addUserIDsQuery(fmt.Sprintf("/api/homework/%d/inter-scores", activityID), userIDs)
	var result []InterScore
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListIntraScoreRules returns intra-review rules for an activity.
func (s *Service) ListIntraScoreRules(ctx context.Context, activityID int, userIDs []int) (*IntraScoreRulesResponse, error) {
	u := addUserIDsQuery(fmt.Sprintf("/api/activities/%d/intra-score-rules", activityID), userIDs)
	result := new(IntraScoreRulesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListHomeworkScores returns homework scores for an activity.
func (s *Service) ListHomeworkScores(ctx context.Context, activityID int, userIDs []int) (*HomeworkScoresResponse, error) {
	u := addUserIDsQuery(fmt.Sprintf("/api/activities/%d/homework-scores", activityID), userIDs)
	result := new(HomeworkScoresResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListRecommendSubmissions returns recommended submissions for an activity.
func (s *Service) ListRecommendSubmissions(ctx context.Context, activityID int) (*RecommendSubmissionsResponse, error) {
	u := fmt.Sprintf("/api/activities/%d/recommend-submissions", activityID)
	result := new(RecommendSubmissionsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetLogsByType returns homework logs of the requested type.
func (s *Service) GetLogsByType(ctx context.Context, homeworkID int, logType string) (*HomeworkLogsResponse, error) {
	u := sdk.AddQueryParams(fmt.Sprintf("/api/homeworks/%d/logs", homeworkID), map[string]string{"log_type": logType})
	result := new(HomeworkLogsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetStudentHomeworkRedoMap returns redo counts keyed by student/group ID.
func (s *Service) GetStudentHomeworkRedoMap(ctx context.Context, homeworkID int) (*RedoMapResponse, error) {
	u := fmt.Sprintf("/api/homework/%d/redo-map", homeworkID)
	result := new(RedoMapResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// MarkHomeworkSubmissionToRedo marks homework submissions for redo.
func (s *Service) MarkHomeworkSubmissionToRedo(ctx context.Context, activityID int, body *MarkHomeworkSubmissionToRedoRequest) error {
	u := fmt.Sprintf("/api/course/activities/%d/submission/redo", activityID)
	_, err := s.client.Put(ctx, u, body, nil)
	return err
}

// GetHomeworkDuplicateRate returns duplicate-detection rates for selected targets.
func (s *Service) GetHomeworkDuplicateRate(ctx context.Context, homeworkID int, targetIDs []int) (*DuplicateDetectRatesResponse, error) {
	u := sdk.AddQueryParams(fmt.Sprintf("/api/homework/%d/duplicate-detect/rate", homeworkID), map[string]string{
		"target_ids": intsToCSV(targetIDs),
	})
	result := new(DuplicateDetectRatesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetHomeworkDuplicateRateWithSubmissionID returns duplicate-detection rates for a single submission.
func (s *Service) GetHomeworkDuplicateRateWithSubmissionID(ctx context.Context, submissionID int) (*DuplicateDetectRatesResponse, error) {
	u := fmt.Sprintf("/api/homework/submission/%d/duplicate-detect-rate", submissionID)
	result := new(DuplicateDetectRatesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// AddUploadsToDuplicateLib adds uploads into a course homework duplicate library.
func (s *Service) AddUploadsToDuplicateLib(ctx context.Context, courseID int, body *AddDuplicateLibUploadsRequest) error {
	u := fmt.Sprintf("/api/course/%d/homework/duplicate-lib", courseID)
	_, err := s.client.Post(ctx, u, body, nil)
	return err
}

// DeleteUploadsFromDuplicateLib removes uploads from a course homework duplicate library.
func (s *Service) DeleteUploadsFromDuplicateLib(ctx context.Context, courseID int, uploadIDs []int) error {
	u := sdk.AddQueryParams(fmt.Sprintf("/api/course/%d/homework/duplicate-lib", courseID), map[string]string{
		"upload_ids": intsToCSV(uploadIDs),
	})
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// ListDuplicateLibUploads returns paged duplicate-library uploads with frontend-compatible defaults and condition encoding.
func (s *Service) ListDuplicateLibUploads(ctx context.Context, courseID int, params ListDuplicateLibUploadsParams) (*DuplicateLibUploadsResponse, error) {
	page := params.Page
	if page == 0 {
		page = 1
	}
	pageSize := params.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	query := map[string]string{
		"page":      strconv.Itoa(page),
		"page_size": strconv.Itoa(pageSize),
	}
	if encoded := encodeHomeworkConditions(params.Conditions); encoded != "" {
		query["conditions"] = encoded
	}
	u := sdk.AddQueryParams(fmt.Sprintf("/api/course/%d/homework/duplicate-lib", courseID), query)
	result := new(DuplicateLibUploadsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetLastDuplicateDetectTask returns the latest duplicate-detect task for a homework.
func (s *Service) GetLastDuplicateDetectTask(ctx context.Context, homeworkID int, status string) (*DuplicateDetectTask, error) {
	params := map[string]string{}
	if status != "" {
		params["status"] = status
	}
	u := sdk.AddQueryParams(fmt.Sprintf("/api/homework/%d/duplicate-detect/task", homeworkID), params)
	result := new(DuplicateDetectTask)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetDuplicateDetectRawFile returns the raw duplicate-detect source file contents.
func (s *Service) GetDuplicateDetectRawFile(ctx context.Context, fileKey string) (string, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/api/duplicate-detect/file/%s/raw", fileKey), nil)
	if err != nil {
		return "", err
	}
	_, data, err := s.client.DoBytes(req)
	return string(data), err
}

// GetDuplicateDetectResult returns duplicate-detect detail results for a homework file.
func (s *Service) GetDuplicateDetectResult(ctx context.Context, homeworkID int, fileKey string) (*DuplicateDetectTask, error) {
	u := fmt.Sprintf("/api/homework/%d/duplicate-detect-result/file/%s", homeworkID, fileKey)
	result := new(DuplicateDetectTask)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// RequestDuplicateDetectReportDownload requests a downloadable third-party duplicate report.
func (s *Service) RequestDuplicateDetectReportDownload(ctx context.Context, body *DuplicateDetectReportDownloadRequest) (*DuplicateDetectReportDownloadInfo, error) {
	result := new(DuplicateDetectReportDownloadInfo)
	_, err := s.client.Post(ctx, "/api/duplicate-detect/report/download", body, result)
	return result, err
}

// GetInProgressHomeworks returns in-progress homeworks.
func (s *Service) GetInProgressHomeworks(ctx context.Context) ([]InProgressHomework, error) {
	var result []InProgressHomework
	_, err := s.client.Get(ctx, "/api/in-progress-homeworks?no-intercept=true", &result)
	return result, err
}

// GetHomeworksSubmissionStatus returns homework submission statuses across courses.
func (s *Service) GetHomeworksSubmissionStatus(ctx context.Context, courseID int) (HomeworksSubmissionStatusResponse, error) {
	u := fmt.Sprintf("/api/courses/homeworks-submission-status?no-intercept=true&course_id=%d", courseID)
	result := new(HomeworksSubmissionStatusResponse)
	_, err := s.client.Get(ctx, u, result)
	if err != nil {
		return HomeworksSubmissionStatusResponse{}, err
	}
	return *result, nil
}

// DownloadHomeworkZip checks the status of a homework zip download.
func (s *Service) DownloadHomeworkZip(ctx context.Context, activityID int) (HomeworkZipStatusResponse, error) {
	u := fmt.Sprintf("/api/zip-status/homework-zip/%d", activityID)
	result := make(HomeworkZipStatusResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// StartHomeworkAIGenerate starts the AI homework generation SSE stream.
func (s *Service) StartHomeworkAIGenerate(ctx context.Context, courseID, homeworkID int, body *HomeworkAIGenerateRequest) (*http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("/api/courses/%d/homework/%d/ai-generate", courseID, homeworkID), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/event-stream")
	return s.client.HTTPClient().Do(req)
}

// GetSubmissionAnalysisStatus returns whether a submission can be reanalyzed.
func (s *Service) GetSubmissionAnalysisStatus(ctx context.Context, submissionID int, kind string) (SubmissionAnalysisStatusResponse, error) {
	u := fmt.Sprintf("/api/submissions/%d/%s/analysis/can-reanalyze", submissionID, kind)
	result := make(SubmissionAnalysisStatusResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetSubmissionAnalysis returns the stored AI analysis payload for a submission.
func (s *Service) GetSubmissionAnalysis(ctx context.Context, submissionID int, kind string) (SubmissionAnalysisResponse, error) {
	u := fmt.Sprintf("/api/submissions/%d/%s/analysis", submissionID, kind)
	result := make(SubmissionAnalysisResponse)
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// SaveSubmissionAnalysis saves or re-saves AI analysis content for a submission.
func (s *Service) SaveSubmissionAnalysis(ctx context.Context, submissionID int, kind string, body *SubmissionAnalysisRequest) (SubmissionAnalysisResponse, error) {
	u := fmt.Sprintf("/api/submissions/%d/%s/analysis", submissionID, kind)
	result := make(SubmissionAnalysisResponse)
	_, err := s.client.Post(ctx, u, body, &result)
	return result, err
}

func addHomeworkQueryParams(urlStr string, params *ListSubmissionRecordsParams) string {
	query := map[string]string{}
	if params != nil {
		if params.NeedUploadsSize != nil {
			query["need_uploads_size"] = strconv.FormatBool(*params.NeedUploadsSize)
		}
		if len(params.UserIDs) > 0 {
			if payload, err := json.Marshal(params.UserIDs); err == nil {
				query["user_ids"] = string(payload)
			}
		}
	}
	return sdk.AddQueryParams(urlStr, query)
}

func addUserIDsQuery(urlStr string, userIDs []int) string {
	if len(userIDs) == 0 {
		return urlStr
	}
	payload, err := json.Marshal(userIDs)
	if err != nil {
		return urlStr
	}
	return sdk.AddQueryParams(urlStr, map[string]string{"user_ids": string(payload)})
}

func intsToCSV(ids []int) string {
	if len(ids) == 0 {
		return ""
	}
	values := make([]string, len(ids))
	for i, id := range ids {
		values[i] = strconv.Itoa(id)
	}
	return strings.Join(values, ",")
}

func encodeHomeworkConditions(conditions any) string {
	switch value := conditions.(type) {
	case nil:
		return ""
	case string:
		return value
	default:
		payload, err := json.Marshal(value)
		if err != nil {
			return ""
		}
		return string(payload)
	}
}
