package homework

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestHomepageHomeworkModelsDecodeFrontendPayloads(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/in-progress-homeworks":
			if r.URL.Query().Get("no-intercept") != "true" {
				t.Fatalf("unexpected query: %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`[
				{"id":5,"course_id":9,"course_type":8,"title":"章节作业","type":"homework","is_locked":true}
			]`))
		case "/api/courses/homeworks-submission-status":
			if r.URL.Query().Get("no-intercept") != "true" {
				t.Fatalf("unexpected query: %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`{
				"homework_statuses":[
					{"id":5,"score":95,"status":"submitted","status_code":"scored","is_announce_score_time_passed":true}
				]
			}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))

	homeworks, err := service.GetInProgressHomeworks(context.Background())
	if err != nil {
		t.Fatalf("GetInProgressHomeworks returned error: %v", err)
	}
	if len(homeworks) != 1 || homeworks[0].CourseID != 9 || !homeworks[0].IsLocked || homeworks[0].CourseType != 8 {
		t.Fatalf("unexpected in-progress homeworks: %#v", homeworks)
	}

	status, err := service.GetHomeworksSubmissionStatus(context.Background(), 0)
	if err != nil {
		t.Fatalf("GetHomeworksSubmissionStatus returned error: %v", err)
	}
	if len(status.HomeworkStatuses) != 1 || status.HomeworkStatuses[0].Score == nil || *status.HomeworkStatuses[0].Score != 95 || !status.HomeworkStatuses[0].IsAnnounceScoreTimePassed {
		t.Fatalf("unexpected homework statuses: %#v", status)
	}
}

func TestHomeworkSubmissionHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/8/groups/6/submission_list":
			_, _ = w.Write([]byte(`{"list":[{"id":10,"group_id":6}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/homework/8/student-submissions":
			if r.URL.Query().Get("need_uploads_size") != "true" {
				t.Fatalf("unexpected need_uploads_size query: %q", r.URL.Query().Get("need_uploads_size"))
			}
			if r.URL.Query().Get("user_ids") != `[2,3]` {
				t.Fatalf("unexpected user_ids query: %q", r.URL.Query().Get("user_ids"))
			}
			_, _ = w.Write([]byte(`{"submissions":[{"id":11,"student_id":2}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/homework/8/group-submissions":
			if r.URL.Query().Get("need_uploads_size") != "false" {
				t.Fatalf("unexpected group need_uploads_size query: %q", r.URL.Query().Get("need_uploads_size"))
			}
			_, _ = w.Write([]byte(`{"submissions":[{"id":12,"group_id":6}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/activities/8/students/2/submission":
			_, _ = w.Write([]byte(`{"id":13,"student_id":2}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/8/groups/6/submission":
			_, _ = w.Write([]byte(`{"id":14,"group_id":6}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/submissions/10/marked_attachments":
			_, _ = w.Write([]byte(`{"marked_attachment_infos":[{"origin_upload":{"id":2,"upload":{"id":3,"name":"origin.pdf"}},"marked_attachment":{"id":4,"name":"marked.pdf"}}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/submissions/10/marked_attachments/4":
			_, _ = w.Write([]byte(`{"marked_attachment":{"id":4,"name":"marked.pdf"}}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/8/homework-scores":
			if r.URL.Query().Get("user_ids") != `[2,3]` {
				t.Fatalf("unexpected homework-scores user_ids query: %q", r.URL.Query().Get("user_ids"))
			}
			_, _ = w.Write([]byte(`{"homework_scores":[{"student_id":2,"score":91}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/8/groups/6/scores":
			_, _ = w.Write([]byte(`{"group_id":6,"score":86}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/8/recommend-submissions":
			_, _ = w.Write([]byte(`{"submissions":[{"id":15,"is_recommend":true}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/homework/8/make-up-records":
			if r.URL.Query().Get("user_ids") != `[2,3]` {
				t.Fatalf("unexpected make-up-records user_ids query: %q", r.URL.Query().Get("user_ids"))
			}
			_, _ = w.Write([]byte(`{"make_up_records":[{"student_id":2}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/homework/8/resubmit-records":
			if r.URL.Query().Get("user_ids") != `[2,3]` {
				t.Fatalf("unexpected resubmit-records user_ids query: %q", r.URL.Query().Get("user_ids"))
			}
			_, _ = w.Write([]byte(`{"resubmit_records":[{"student_id":3}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/homework/8/inter-scores":
			_, _ = w.Write([]byte(`[{"score":88}]`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/homework/8/intra-scores":
			_, _ = w.Write([]byte(`[{"score":77}]`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/8/intra-score-rules":
			if r.URL.Query().Get("user_ids") != `[2,3]` {
				t.Fatalf("unexpected intra-score-rules user_ids query: %q", r.URL.Query().Get("user_ids"))
			}
			_, _ = w.Write([]byte(`{"intra_scores":[{"student_id":2,"score":66}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/homeworks/8/logs":
			if r.URL.Query().Get("log_type") != "submit" {
				t.Fatalf("unexpected log_type query: %q", r.URL.Query().Get("log_type"))
			}
			_, _ = w.Write([]byte(`{"logs":[{"id":1,"type":"submit"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/homework/8/redo-map":
			_, _ = w.Write([]byte(`{"redo_map":{"2":1}}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/course/activities/8/submission/redo":
			var body MarkHomeworkSubmissionToRedoRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode redo body: %v", err)
			}
			if body.SubmissionID != 10 || body.StudentID != 2 {
				t.Fatalf("unexpected redo body: %#v", body)
			}
			_, _ = w.Write([]byte(`{}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/homework/8/duplicate-detect/rate":
			if r.URL.Query().Get("target_ids") != "10,11" {
				t.Fatalf("unexpected target_ids query: %q", r.URL.Query().Get("target_ids"))
			}
			_, _ = w.Write([]byte(`{"items":[{"submission_id":10,"rate":0.8}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/homework/submission/10/duplicate-detect-rate":
			_, _ = w.Write([]byte(`{"items":[{"submission_id":10,"rate":0.9}]}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	groupSubs, err := service.ListGroupSubmissions(ctx, 8, 6)
	if err != nil || len(groupSubs.List) != 1 || groupSubs.List[0].ID != 10 {
		t.Fatalf("unexpected group submissions: %#v, err=%v", groupSubs, err)
	}

	needUploadsSize := true
	studentSubs, err := service.ListStudentSubmissionRecords(ctx, 8, &ListSubmissionRecordsParams{NeedUploadsSize: &needUploadsSize, UserIDs: []int{2, 3}})
	if err != nil || len(studentSubs.Submissions) != 1 || studentSubs.Submissions[0].ID != 11 {
		t.Fatalf("unexpected student submission records: %#v, err=%v", studentSubs, err)
	}

	groupRecords, err := service.ListGroupSubmissionRecords(ctx, 8, false)
	if err != nil || len(groupRecords.Submissions) != 1 || groupRecords.Submissions[0].ID != 12 {
		t.Fatalf("unexpected group submission records: %#v, err=%v", groupRecords, err)
	}

	studentSubmission, err := service.GetStudentSubmission(ctx, 8, 2)
	if err != nil || studentSubmission.ID != 13 {
		t.Fatalf("unexpected student submission: %#v, err=%v", studentSubmission, err)
	}

	groupSubmission, err := service.GetGroupSubmission(ctx, 8, 6)
	if err != nil || groupSubmission.ID != 14 {
		t.Fatalf("unexpected group submission: %#v, err=%v", groupSubmission, err)
	}

	attachments, err := service.ListMarkedAttachments(ctx, 10)
	if err != nil || len(attachments.MarkedAttachmentInfos) != 1 || attachments.MarkedAttachmentInfos[0].OriginUpload == nil || attachments.MarkedAttachmentInfos[0].OriginUpload.Upload == nil || attachments.MarkedAttachmentInfos[0].OriginUpload.Upload.Name != "origin.pdf" {
		t.Fatalf("unexpected marked attachments: %#v, err=%v", attachments, err)
	}

	attachment, err := service.GetMarkedAttachment(ctx, 10, 4)
	if err != nil || attachment.MarkedAttachment == nil || attachment.MarkedAttachment.Name != "marked.pdf" {
		t.Fatalf("unexpected marked attachment: %#v, err=%v", attachment, err)
	}

	interScores, err := service.ListInterScores(ctx, 8)
	if err != nil || len(interScores) != 1 {
		t.Fatalf("unexpected inter scores: %#v, err=%v", interScores, err)
	}

	intraScores, err := service.ListIntraScores(ctx, 8)
	if err != nil || len(intraScores) != 1 {
		t.Fatalf("unexpected intra scores: %#v, err=%v", intraScores, err)
	}

	scores, err := service.ListHomeworkScores(ctx, 8, []int{2, 3})
	if err != nil || len(scores.HomeworkScores) != 1 || scores.HomeworkScores[0].Score == nil || *scores.HomeworkScores[0].Score != 91 {
		t.Fatalf("unexpected homework scores: %#v, err=%v", scores, err)
	}

	groupScore, err := service.GetGroupHomeworkScore(ctx, 8, 6)
	if err != nil || groupScore.GroupID == nil || *groupScore.GroupID != 6 || groupScore.Score == nil || *groupScore.Score != 86 {
		t.Fatalf("unexpected group score: %#v, err=%v", groupScore, err)
	}

	recommendations, err := service.ListRecommendSubmissions(ctx, 8)
	if err != nil || len(recommendations.Submissions) != 1 || !recommendations.Submissions[0].IsRecommend {
		t.Fatalf("unexpected recommendations: %#v, err=%v", recommendations, err)
	}

	makeUps, err := service.ListMakeUpRecords(ctx, 8, []int{2, 3})
	if err != nil || len(makeUps.MakeUpRecords) != 1 {
		t.Fatalf("unexpected make-up records: %#v, err=%v", makeUps, err)
	}

	resubmits, err := service.ListResubmitRecords(ctx, 8, []int{2, 3})
	if err != nil || len(resubmits.ResubmitRecords) != 1 {
		t.Fatalf("unexpected resubmit records: %#v, err=%v", resubmits, err)
	}

	intraRules, err := service.ListIntraScoreRules(ctx, 8, []int{2, 3})
	if err != nil || len(intraRules.IntraScores) != 1 {
		t.Fatalf("unexpected intra-score rules: %#v, err=%v", intraRules, err)
	}

	logs, err := service.GetLogsByType(ctx, 8, "submit")
	if err != nil || len(logs.Logs) != 1 {
		t.Fatalf("unexpected logs: %#v, err=%v", logs, err)
	}

	redoMap, err := service.GetStudentHomeworkRedoMap(ctx, 8)
	if err != nil || redoMap.RedoMap["2"] != 1 {
		t.Fatalf("unexpected redo map: %#v, err=%v", redoMap, err)
	}

	if err := service.MarkHomeworkSubmissionToRedo(ctx, 8, &MarkHomeworkSubmissionToRedoRequest{SubmissionID: 10, StudentID: 2}); err != nil {
		t.Fatalf("MarkHomeworkSubmissionToRedo returned error: %v", err)
	}

	duplicateRates, err := service.GetHomeworkDuplicateRate(ctx, 8, []int{10, 11})
	if err != nil || len(duplicateRates.Items) != 1 {
		t.Fatalf("unexpected duplicate rates: %#v, err=%v", duplicateRates, err)
	}

	duplicateRatesBySubmission, err := service.GetHomeworkDuplicateRateWithSubmissionID(ctx, 10)
	if err != nil || len(duplicateRatesBySubmission.Items) != 1 {
		t.Fatalf("unexpected duplicate rates by submission: %#v, err=%v", duplicateRatesBySubmission, err)
	}
}
