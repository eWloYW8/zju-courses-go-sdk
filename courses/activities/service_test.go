package activities

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestFrontendActivityHelpersUseVerifiedEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/activities-read/11":
			_, _ = w.Write([]byte(`{"id":"11","activity_id":11,"activity_type":"learning_activity"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/completion-criteria":
			if got := r.URL.Query().Get("activity_type"); got != "homework" {
				t.Fatalf("unexpected activity_type: %q", got)
			}
			if got := r.URL.Query().Get("course_id"); got != "9" {
				t.Fatalf("unexpected course_id: %q", got)
			}
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected no-intercept: %q", got)
			}
			_, _ = w.Write([]byte(`{"completion_criteria":[{"id":1,"completion_criterion_type":"score","value":60}],"has_completion_criterion":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/9/activities":
			_, _ = w.Write([]byte(`{"activities":[{"id":3,"title":"A1","type":"homework"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activity/12/submitter/34/score-records":
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected page: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "20" {
				t.Fatalf("unexpected page_size: %q", got)
			}
			_, _ = w.Write([]byte(`{"records":[{"id":1,"score_method":"manual","score":88,"final_score":90,"operator_name":"TA"}],"page":2,"page_size":20,"pages":4,"total":61,"start":21,"end":40}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/9/grade-score-items":
			_, _ = w.Write([]byte(`[{"id":1,"display_name":"平时","type":"homework","percentage":0.4}]`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activies/classin/join-url":
			if got := r.URL.Query().Get("course_id"); got != "9" || r.URL.Query().Get("activity_id") != "12" || r.URL.Query().Get("user_id") != "34" {
				t.Fatalf("unexpected classin join query: %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`{"url":"https://classin.example/join"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/classin/webcast-url":
			if got := r.URL.Query().Get("course_id"); got != "9" || r.URL.Query().Get("activity_id") != "12" {
				t.Fatalf("unexpected classin webcast query: %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`{"url":"https://classin.example/webcast"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/12/comments":
			if got := r.URL.Query().Get("page"); got != "3" || r.URL.Query().Get("page_size") != "10" {
				t.Fatalf("unexpected paging: %s", r.URL.RawQuery)
			}
			if got := r.URL.Query().Get("order_key"); got != "created_at" || r.URL.Query().Get("order") != "desc" {
				t.Fatalf("unexpected ordering: %s", r.URL.RawQuery)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"type":"forum"}` {
				t.Fatalf("unexpected conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"comments":[{"id":8,"content":"hi"}],"page":3,"page_size":10}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/12/comment/count":
			if got := r.URL.Query().Get("conditions"); got != `{"type":"forum"}` {
				t.Fatalf("unexpected count conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"forum":4,"question":1}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/12/comment/page-count":
			if got := r.URL.Query().Get("conditions"); got != `{"type":"forum"}` {
				t.Fatalf("unexpected page-count conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"page_stats":[{"page":1,"forum":2}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/12/resources":
			_, _ = w.Write([]byte(`[{"id":7,"name":"slides.pdf"}]`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/activities/12/resources/7":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode update activity resource: %v", err)
			}
			if body["allow_download"] != true {
				t.Fatalf("unexpected activity resource body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodDelete && r.URL.Path == "/api/activities/12/resources/7":
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if _, err := service.MarkActivityRead(ctx, 11); err != nil {
		t.Fatalf("MarkActivityRead returned error: %v", err)
	}
	criteria, err := service.GetActivityCompletionCriteria(ctx, 9, &ActivityCriteriaQuery{CourseID: 9, ActivityType: "homework"})
	if err != nil || !criteria.HasCompletionCriterion || len(criteria.CompletionCriteria) != 1 {
		t.Fatalf("unexpected completion criteria: %#v, err=%v", criteria, err)
	}
	activities, err := service.ListCourseActivities(ctx, 9)
	if err != nil || len(activities) != 1 || activities[0].ID != 3 {
		t.Fatalf("unexpected activities: %#v, err=%v", activities, err)
	}
	records, err := service.GetScoreRecords(ctx, 12, 34, ActivityScoreRecordsParams{Page: 2, PageSize: 20})
	if err != nil || len(records.Items) != 1 || records.Items[0].OperatorName == nil || *records.Items[0].OperatorName != "TA" || records.Start != 21 || records.End != 40 {
		t.Fatalf("unexpected score records: %#v, err=%v", records, err)
	}
	items, err := service.ListGradeScoreItems(ctx, 9)
	if err != nil || len(items) != 1 || items[0].DisplayName != "平时" {
		t.Fatalf("unexpected score items: %#v, err=%v", items, err)
	}
	joinURL, err := service.GetClassinJoinURLWithParams(ctx, ClassinJoinURLParams{CourseID: 9, ActivityID: 12, UserID: 34})
	if err != nil || joinURL.URL != "https://classin.example/join" {
		t.Fatalf("unexpected classin join url: %#v, err=%v", joinURL, err)
	}
	webcastURL, err := service.GetClassinWebcastURLWithParams(ctx, ClassinWebcastURLParams{CourseID: 9, ActivityID: 12})
	if err != nil || webcastURL.URL != "https://classin.example/webcast" {
		t.Fatalf("unexpected classin webcast url: %#v, err=%v", webcastURL, err)
	}
	comments, err := service.ListCommentsWithParams(ctx, 12, CommentListParams{
		Page: 3, PageSize: 10, OrderKey: "created_at", Order: "desc", Conditions: map[string]any{"type": "forum"},
	})
	if err != nil || len(comments.Comments) != 1 || comments.Comments[0].ID != 8 {
		t.Fatalf("unexpected comments: %#v, err=%v", comments, err)
	}
	count, err := service.GetCommentCountWithConditions(ctx, 12, map[string]any{"type": "forum"})
	if err != nil || count.Forum != 4 {
		t.Fatalf("unexpected comment count: %#v, err=%v", count, err)
	}
	pageCount, err := service.GetCommentPageCountWithConditions(ctx, 12, map[string]any{"type": "forum"})
	if err != nil || len(pageCount.PageStats) != 1 || pageCount.PageStats[0].Forum != 2 {
		t.Fatalf("unexpected page count: %#v, err=%v", pageCount, err)
	}
	resources, err := service.ListActivityResources(ctx, 12)
	if err != nil || len(resources) != 1 {
		t.Fatalf("unexpected activity resources: %#v, err=%v", resources, err)
	}
	if err := service.UpdateActivityResource(ctx, 12, 7, UpdateActivityResourceRequest{"allow_download": true}); err != nil {
		t.Fatalf("UpdateActivityResource returned error: %v", err)
	}
	if err := service.DeleteActivityResource(ctx, 12, 7); err != nil {
		t.Fatalf("DeleteActivityResource returned error: %v", err)
	}
}

func TestActivityStateAndLicenseHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/activities-read/11":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode activity-read body: %v", err)
			}
			if body["progress"] != float64(80) {
				t.Fatalf("unexpected activity-read body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"logged":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/11/completion-criteria":
			if got := r.URL.Query().Get("activity_type"); got != "homework" {
				t.Fatalf("unexpected activity_type: %q", got)
			}
			if got := r.URL.Query().Get("course_id"); got != "9" {
				t.Fatalf("unexpected course_id: %q", got)
			}
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected no-intercept: %q", got)
			}
			_, _ = w.Write([]byte(`{"completion_criteria":[{"id":2}],"has_completion_criterion":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activity/11/submission-number":
			_, _ = w.Write([]byte(`{"submission_number":12}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activity/11/has-reviewed-inter-score":
			_, _ = w.Write([]byte(`{"has_reviewed_inter_score":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activity/11/first-submission-time":
			_, _ = w.Write([]byte(`{"first_submission_time":"2026-03-01 12:00:00"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activity/11/is-inter-review-started":
			_, _ = w.Write([]byte(`{"is_started":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activity/11/is-intra-review-started":
			_, _ = w.Write([]byte(`{"is_started":false}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activity/11/is-homework-expired":
			_, _ = w.Write([]byte(`{"is_expired":false}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/11/uploads-license":
			_, _ = w.Write([]byte(`{"allow_cc_license":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/online-videos/11/activity-read-count":
			_, _ = w.Write([]byte(`{"read_count":33}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/11/web-link-scores":
			_, _ = w.Write([]byte(`{"scores":[{"student_id":1,"score":90}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/11/virtual-experiment-scores":
			_, _ = w.Write([]byte(`{"scores":[{"student_id":1,"score":95}]}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if logged, err := service.LogActivityRead(ctx, 11, map[string]any{"progress": 80}); err != nil || logged["logged"] != true {
		t.Fatalf("unexpected LogActivityRead result: %#v, err=%v", logged, err)
	}
	if criteria, err := service.GetActivityCompletionCriteriaDetail(ctx, 11, &ActivityCompletionCriteriaDetailQuery{
		ActivityType: "homework",
		CourseID:     9,
	}); err != nil || !criteria.HasCompletionCriterion || len(criteria.CompletionCriteria) != 1 {
		t.Fatalf("unexpected activity completion criteria detail: %#v, err=%v", criteria, err)
	}
	if resp, err := service.GetSubmissionNumber(ctx, 11); err != nil || resp["submission_number"] != float64(12) {
		t.Fatalf("unexpected submission number: %#v, err=%v", resp, err)
	}
	if resp, err := service.HasReviewedInterScore(ctx, 11); err != nil || resp["has_reviewed_inter_score"] != true {
		t.Fatalf("unexpected reviewed inter score: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetHomeworkFirstSubmissionTime(ctx, 11); err != nil || resp["first_submission_time"] != "2026-03-01 12:00:00" {
		t.Fatalf("unexpected first submission time: %#v, err=%v", resp, err)
	}
	if resp, err := service.IsInterReviewStarted(ctx, 11); err != nil || resp["is_started"] != true {
		t.Fatalf("unexpected inter review state: %#v, err=%v", resp, err)
	}
	if resp, err := service.IsIntraReviewStarted(ctx, 11); err != nil || resp["is_started"] != false {
		t.Fatalf("unexpected intra review state: %#v, err=%v", resp, err)
	}
	if resp, err := service.IsHomeworkExpired(ctx, 11); err != nil || resp["is_expired"] != false {
		t.Fatalf("unexpected homework expiry state: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetActivityUploadsLicenseInfo(ctx, 11); err != nil || resp["allow_cc_license"] != true {
		t.Fatalf("unexpected uploads-license response: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetOnlineVideoActivityReadCount(ctx, 11); err != nil || resp["read_count"] != float64(33) {
		t.Fatalf("unexpected online-video read count: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetWebLinkScores(ctx, 11); err != nil || len(resp["scores"].([]any)) != 1 {
		t.Fatalf("unexpected web-link scores: %#v, err=%v", resp, err)
	}
	if resp, err := service.GetVirtualExperimentScores(ctx, 11); err != nil || len(resp["scores"].([]any)) != 1 {
		t.Fatalf("unexpected virtual-experiment scores: %#v, err=%v", resp, err)
	}
}

func TestCommentRepliesAndRawClassinCompatibility(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/12/comments/replies":
			if got := r.URL.Query()["comment_ids[]"]; len(got) != 2 || got[0] != "3" || got[1] != "4" {
				t.Fatalf("unexpected comment ids: %#v", r.URL.Query())
			}
			_, _ = w.Write([]byte(`[{"id":5,"content":"r1"}]`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/12/comments/3/replies":
			if got := r.URL.Query().Get("page"); got != "2" || r.URL.Query().Get("page_size") != "5" {
				t.Fatalf("unexpected reply paging: %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`[{"id":6,"content":"r2"}]`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/activities/12/comments/3/reply":
			_, _ = w.Write([]byte(`{"id":7,"content":"reply"}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/activities/12/reply/7":
			_, _ = w.Write([]byte(`{"id":7,"content":"updated"}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/api/activities/12/reply/7":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/activities/12/comments/3":
			_, _ = w.Write([]byte(`{"id":3,"content":"updated comment"}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/api/activities/12/comments/3":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/activies/classin/join-url":
			_, _ = w.Write([]byte(`{"url":"https://classin.example/raw-join"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/activities/classin/webcast-url":
			_, _ = w.Write([]byte(`{"url":"https://classin.example/raw-webcast"}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	replies, err := service.ListCommentRepliesByCommentIDs(ctx, 12, []int{3, 4})
	if err != nil || len(replies) != 1 || replies[0].ID != 5 {
		t.Fatalf("unexpected batched replies: %#v, err=%v", replies, err)
	}
	pagedReplies, err := service.ListCommentReplies(ctx, 12, 3, &model.ListOptions{Page: 2, PageSize: 5})
	if err != nil || len(pagedReplies) != 1 || pagedReplies[0].ID != 6 {
		t.Fatalf("unexpected paged replies: %#v, err=%v", pagedReplies, err)
	}
	reply, err := service.ReplyComment(ctx, 12, 3, &CreateCommentRequest{Content: "reply"})
	if err != nil || reply.ID != 7 {
		t.Fatalf("unexpected reply create: %#v, err=%v", reply, err)
	}
	updatedReply, err := service.UpdateCommentReply(ctx, 12, 7, &CreateCommentRequest{Content: "updated"})
	if err != nil || updatedReply.Content != "updated" {
		t.Fatalf("unexpected reply update: %#v, err=%v", updatedReply, err)
	}
	if err := service.DeleteCommentReply(ctx, 12, 7); err != nil {
		t.Fatalf("DeleteCommentReply returned error: %v", err)
	}
	updatedComment, err := service.UpdateComment(ctx, 12, 3, &CreateCommentRequest{Content: "updated comment"})
	if err != nil || updatedComment.Content != "updated comment" {
		t.Fatalf("unexpected comment update: %#v, err=%v", updatedComment, err)
	}
	if err := service.DeleteComment(ctx, 12, 3); err != nil {
		t.Fatalf("DeleteComment returned error: %v", err)
	}
	rawJoin, err := service.GetClassinJoinURL(ctx, 9)
	if err != nil || string(rawJoin) != `{"url":"https://classin.example/raw-join"}` {
		t.Fatalf("unexpected raw join payload: %s, err=%v", string(rawJoin), err)
	}
	rawWebcast, err := service.GetClassinWebcastURL(ctx, 12)
	if err != nil || string(rawWebcast) != `{"url":"https://classin.example/raw-webcast"}` {
		t.Fatalf("unexpected raw webcast payload: %s, err=%v", string(rawWebcast), err)
	}
}
