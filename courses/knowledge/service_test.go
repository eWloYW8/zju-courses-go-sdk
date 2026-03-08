package knowledge

import (
	"context"
	"net/http"
	"net/http/httptest"
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
