package knowledge

import (
	"context"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestKnowledgeModelsDecodeFrontendPayloads(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/knowledge-node/7":
			_, _ = w.Write([]byte(`{
				"id":7,
				"name":"Limits",
				"description":"Intro node",
				"source":"manual",
				"parent_id":1,
				"mastery_rate":88.5,
				"completeness_rate":77,
				"resource_ref_count":2,
				"activity_ref_count":3,
				"subject_ref_count":1,
				"capture_ref_count":4,
				"cognitive_dimension":"apply",
				"prev_relation":[{"id":1,"name":"prev","relation_id":11}],
				"labels":[{"key":"difficulty","value":"easy"}],
				"teaching_objectives":[{"id":9,"content":"understand limit","refer_count":2}],
				"uploads":[{"id":10,"name":"slides.pdf","viewed":true,"refer_id":99,"refer_type":"knowledge_node","is_from_knowledge_graph":true}],
				"external_resources":[{"id":5,"external_id":8,"title":"MOOC","type":"url","viewed":true,"visits":3}],
				"captures":[{"id":6,"code":"CAP1","name":"Capture","org_code":"ORG","viewed":true,"refer_id":4}],
				"subjects":[{"id":12,"type":"single_selection","description":"Q1","source":{"id":2,"type":"exam","name":"Quiz 1"},"knowledge_node_reference":[{"id":7,"name":"Limits"}]}],
				"activities":[{"id":13,"title":"Quiz","type":"exam","module_name":"Week 1","syllabus_name":"Lesson A"}]
			}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/3/knowledge-nodes/statistics/summary":
			_, _ = w.Write([]byte(`{
				"node_count":2,
				"node_with_reference_count":1,
				"average_mastery_rate":84.5,
				"average_completeness_rate":71,
				"relation_count":3,
				"nodes":[{"id":1,"name":"Limits","average_mastery_rate":84.5,"average_completeness_rate":71,"cognitive_dimension":"apply"}]
			}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/knowledge-nodes/4/statistics/student-detail":
			if got := r.URL.Query().Get("page"); got != "1" {
				t.Fatalf("unexpected page query: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "10" {
				t.Fatalf("unexpected page_size query: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"alice"}` {
				t.Fatalf("unexpected conditions query: %q", got)
			}
			_, _ = w.Write([]byte(`{
				"items":[{"student_id":2,"name":"Alice","user_no":"001","mastery_rate":95,"completeness_rate":88.5,"mastery_rate_rank":1,"completeness_rate_rank":2}]
			}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	node, err := service.GetKnowledgeNode(ctx, 7)
	if err != nil {
		t.Fatalf("GetKnowledgeNode returned error: %v", err)
	}
	if node.MasteryRate == nil || *node.MasteryRate != 88.5 || node.CognitiveDimension != "apply" {
		t.Fatalf("unexpected knowledge node metrics: %#v", node)
	}
	if len(node.PrevRelation) != 1 || node.PrevRelation[0].RelationID != 11 {
		t.Fatalf("unexpected knowledge relations: %#v", node.PrevRelation)
	}
	if len(node.Uploads) != 1 || !node.Uploads[0].Viewed || node.Uploads[0].ReferID != 99 {
		t.Fatalf("unexpected knowledge uploads: %#v", node.Uploads)
	}
	if len(node.Activities) != 1 || node.Activities[0].ModuleName != "Week 1" || node.Activities[0].SyllabusName != "Lesson A" {
		t.Fatalf("unexpected knowledge activities: %#v", node.Activities)
	}
	if len(node.Subjects) != 1 || node.Subjects[0].Source == nil || node.Subjects[0].Source.Name != "Quiz 1" {
		t.Fatalf("unexpected knowledge subjects: %#v", node.Subjects)
	}

	summary, err := service.GetKnowledgeNodeStatisticsSummary(ctx, 3)
	if err != nil {
		t.Fatalf("GetKnowledgeNodeStatisticsSummary returned error: %v", err)
	}
	if summary.AverageMasteryRate == nil || *summary.AverageMasteryRate != 84.5 {
		t.Fatalf("unexpected summary mastery rate: %#v", summary)
	}
	if len(summary.Nodes) != 1 || summary.Nodes[0].AverageCompletenessRate == nil || *summary.Nodes[0].AverageCompletenessRate != 71 {
		t.Fatalf("unexpected summary nodes: %#v", summary.Nodes)
	}

	details, err := service.GetKnowledgeNodeStatisticsStudentDetail(ctx, 4, &model.ListOptions{Page: 1, PageSize: 10}, `{"keyword":"alice"}`)
	if err != nil {
		t.Fatalf("GetKnowledgeNodeStatisticsStudentDetail returned error: %v", err)
	}
	if len(details.Items) != 1 || details.Items[0].MasteryRate == nil || *details.Items[0].MasteryRate != 95 {
		t.Fatalf("unexpected student details: %#v", details.Items)
	}
	if details.Items[0].MasteryRateRank == nil || *details.Items[0].MasteryRateRank != 1 {
		t.Fatalf("unexpected mastery rank: %#v", details.Items[0])
	}
}

func TestKnowledgeFrontendHelpersUseAnonymousMultipartAndExportEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/course/3/knowledge-nodes":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"items":[{"id":1,"name":"Limits"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/course/3/knowledge-nodes/statistics/summary":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"node_count":1}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/courses/3/extension-apps":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"data":[{"name":"knowledge-graph"}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/knowledge-nodes/parse/docx":
			assertMultipartFields(t, r, map[string]string{}, true)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"items":[{"name":"Parsed"}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/courses/3/knowledge-nodes/import":
			if strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
				assertMultipartFields(t, r, map[string]string{"import_type": "markdown"}, true)
			} else {
				if ct := r.Header.Get("Content-Type"); ct != "application/json" {
					t.Fatalf("unexpected import content-type: %s", ct)
				}
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/courses/5/knowledge-nodes/import":
			assertMultipartFields(t, r, map[string]string{
				"import_type":           "course",
				"source_course_id":      "3",
				"import_refer_resource": "true",
			}, false)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/courses/3/knowledge-nodes/export":
			if r.URL.Query().Get("format") != "csv" {
				t.Fatalf("unexpected export format: %q", r.URL.Query().Get("format"))
			}
			w.Header().Set("Content-Type", "text/csv")
			_, _ = w.Write([]byte("id,name\n1,Limits\n"))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/3/knowledge-graph/cluster":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"nodes":[{"id":1}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/3/knowledge-graph/sync-status":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"status":"running"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/knowledge-graph/kfs-courses/-/published-forest-versions:batchGet":
			if got := r.URL.Query()["ids[]"]; len(got) != 2 || got[0] != "1" || got[1] != "2" {
				t.Fatalf("unexpected published-forest ids query: %#v", r.URL.Query())
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`[{"id":9,"course_id":1}]`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/knowledge-graph/forest-versions/-/stats:batchGet":
			if got := r.URL.Query()["ids[]"]; len(got) != 2 || got[0] != "7" || got[1] != "8" {
				t.Fatalf("unexpected forest-stats ids query: %#v", r.URL.Query())
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`[{"id":7,"topic_count":2}]`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/3/knowledge-nodes/ai-parse":
			if accept := r.Header.Get("Accept"); accept != "text/event-stream" {
				t.Fatalf("unexpected ai-parse accept header: %q", accept)
			}
			w.Header().Set("Content-Type", "text/event-stream")
			_, _ = io.WriteString(w, "data: ok\n\n")
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/3/knowledge-base":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"id":7,"name":"KB"}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/3/knowledge-base/7/resources/uploads":
			assertMultipartFields(t, r, map[string]string{}, true)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"id":15,"status":"processing"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/course/3/knowledge-base/7/resources":
			if got := r.URL.Query().Get("page"); got != "1" {
				t.Fatalf("unexpected kb page query: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "10" {
				t.Fatalf("unexpected kb page_size query: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"pdf"}` {
				t.Fatalf("unexpected kb conditions query: %q", got)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"page":1,"page_size":10,"pages":1,"total":1,"items":[{"id":15,"name":"doc.pdf","status":"ready"}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/3/knowledge-base/7/resources/remove":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"removed":true}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/3/knowledge-base/7/resources/retry":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"retried":1}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	nodes, err := service.GetKnowledgeNodeTreesWithPrefix(ctx, 3, true)
	if err != nil || len(nodes) != 1 || nodes[0].Name != "Limits" {
		t.Fatalf("unexpected anonymous knowledge nodes: %#v, err=%v", nodes, err)
	}

	summary, err := service.GetKnowledgeNodeStatisticsSummaryWithPrefix(ctx, 3, true)
	if err != nil || summary.NodeCount != 1 {
		t.Fatalf("unexpected anonymous summary: %#v, err=%v", summary, err)
	}

	apps, err := service.GetCourseExtensionApps(ctx, 3, true)
	if err != nil || len(apps) != 1 || apps[0]["name"] != "knowledge-graph" {
		t.Fatalf("unexpected extension apps: %#v, err=%v", apps, err)
	}

	if _, err := service.ParseKnowledgeNodesFromDocxReader(ctx, strings.NewReader("docx"), "outline.docx"); err != nil {
		t.Fatalf("ParseKnowledgeNodesFromDocxReader returned error: %v", err)
	}

	if err := service.ImportKnowledgeNodes(ctx, 3, &ImportKnowledgeNodesRequest{Data: []map[string]any{{"name": "Limits"}}}); err != nil {
		t.Fatalf("ImportKnowledgeNodes returned error: %v", err)
	}

	if err := service.ImportKnowledgeNodesByFileReader(ctx, 3, strings.NewReader("data"), "outline.md", "markdown"); err != nil {
		t.Fatalf("ImportKnowledgeNodesByFileReader returned error: %v", err)
	}

	if err := service.ImportKnowledgeNodesByCourse(ctx, 5, &ImportKnowledgeNodesByCourseRequest{SourceCourseID: 3, ImportReferResource: true}); err != nil {
		t.Fatalf("ImportKnowledgeNodesByCourse returned error: %v", err)
	}

	exportData, err := service.ExportKnowledgeNodes(ctx, 3, "csv")
	if err != nil || !strings.Contains(string(exportData), "Limits") {
		t.Fatalf("unexpected export data: %q, err=%v", string(exportData), err)
	}

	cluster, err := service.GetKnowledgeGraphCluster(ctx, 3)
	if err != nil || cluster["nodes"] == nil {
		t.Fatalf("unexpected cluster: %#v, err=%v", cluster, err)
	}

	syncStatus, err := service.GetChinamCloudKnowledgeGraphSyncStatus(ctx, 3)
	if err != nil || syncStatus["status"] != "running" {
		t.Fatalf("unexpected sync status: %#v, err=%v", syncStatus, err)
	}

	published, err := service.BatchGetPublishedForestVersionByKFSCourseIDs(ctx, []int{1, 2})
	if err != nil || len(published) != 1 || published[0].ID != 9 {
		t.Fatalf("unexpected published versions: %#v, err=%v", published, err)
	}

	stats, err := service.BatchGetForestVersionStatsByKFSVersionIDs(ctx, []int{7, 8})
	if err != nil || len(stats) != 1 || stats[0].TopicCount != 2 {
		t.Fatalf("unexpected forest stats: %#v, err=%v", stats, err)
	}

	resp, err := service.StartKnowledgeNodesAIParse(ctx, 3, &AIParseKnowledgeNodesRequest{UploadID: 12})
	if err != nil {
		t.Fatalf("StartKnowledgeNodesAIParse returned error: %v", err)
	}
	_ = resp.Body.Close()

	kb, err := service.GetKnowledgeBase(ctx, 3)
	if err != nil || kb["id"] != float64(7) {
		t.Fatalf("unexpected knowledge base: %#v, err=%v", kb, err)
	}

	uploaded, err := service.UploadKnowledgeBaseResource(ctx, 3, 7, strings.NewReader("pdf"), "doc.pdf")
	if err != nil || uploaded["status"] != "processing" {
		t.Fatalf("unexpected uploaded knowledge base resource: %#v, err=%v", uploaded, err)
	}

	resources, err := service.ListKnowledgeBaseResources(ctx, 3, 7, ListKnowledgeBaseResourcesParams{
		Page:       1,
		PageSize:   10,
		Conditions: map[string]any{"keyword": "pdf"},
	})
	if err != nil || len(resources.Items) != 1 || resources.Items[0]["name"] != "doc.pdf" {
		t.Fatalf("unexpected knowledge base resources: %#v, err=%v", resources, err)
	}

	removed, err := service.RemoveKnowledgeBaseResource(ctx, 3, 7, 15)
	if err != nil || removed["removed"] != true {
		t.Fatalf("unexpected remove knowledge base resource response: %#v, err=%v", removed, err)
	}

	retried, err := service.RetryKnowledgeBaseResource(ctx, 3, 7, []int{15})
	if err != nil || retried["retried"] != float64(1) {
		t.Fatalf("unexpected retry knowledge base resource response: %#v, err=%v", retried, err)
	}
}

func assertMultipartFields(t *testing.T, r *http.Request, expectedFields map[string]string, expectFile bool) {
	t.Helper()

	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		t.Fatalf("parse media type: %v", err)
	}
	if mediaType != "multipart/form-data" {
		t.Fatalf("unexpected media type: %s", mediaType)
	}

	reader := multipart.NewReader(r.Body, params["boundary"])
	seenFile := false
	fields := map[string]string{}
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("next part: %v", err)
		}
		data, err := io.ReadAll(part)
		if err != nil {
			t.Fatalf("read part: %v", err)
		}
		if part.FileName() != "" {
			seenFile = true
			continue
		}
		fields[part.FormName()] = string(data)
	}
	if expectFile != seenFile {
		t.Fatalf("unexpected file presence: expect=%v got=%v", expectFile, seenFile)
	}
	for key, value := range expectedFields {
		if fields[key] != value {
			t.Fatalf("unexpected multipart field %s: %q", key, fields[key])
		}
	}
}
