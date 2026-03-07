package courses

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

func newTestService(t *testing.T, fn roundTripFunc) *Service {
	t.Helper()
	client := sdk.NewClient(sdk.WithHTTPClient(&http.Client{Transport: fn}))
	return New(client)
}

func jsonResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestCreateCourseUsesSingularEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":1,"name":"demo"}`), nil
	})

	if _, err := svc.CreateCourse(context.Background(), &Course{Name: "demo"}); err != nil {
		t.Fatalf("CreateCourse error: %v", err)
	}
}

func TestCreateModuleUsesCourseModuleEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/123/module" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":7,"name":"module"}`), nil
	})

	if _, err := svc.CreateModule(context.Background(), 123, &Module{Name: "module"}); err != nil {
		t.Fatalf("CreateModule error: %v", err)
	}
}

func TestDeleteModuleWithOptionsUsesDeleteRelatedActivityQuery(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/module/9" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("delete_related_activity"); got != "true" {
			t.Fatalf("unexpected delete_related_activity: %s", got)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.DeleteModuleWithOptions(context.Background(), 9, &DeleteModuleOptions{DeleteRelatedActivity: true}); err != nil {
		t.Fatalf("DeleteModuleWithOptions error: %v", err)
	}
}

func TestListMyCoursesByConditionsUsesPostEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/my-courses" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"courses":[]}`), nil
	})

	if _, err := svc.ListMyCoursesByConditions(context.Background(), &ListMyCoursesRequest{Page: 1, PageSize: 10}); err != nil {
		t.Fatalf("ListMyCoursesByConditions error: %v", err)
	}
}

func TestDeleteSyllabusWithOptionsUsesSingularEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/syllabus/21" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("delete_related_activity"); got != "true" {
			t.Fatalf("unexpected delete_related_activity: %s", got)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.DeleteSyllabusWithOptions(context.Background(), 21, &DeleteSyllabusOptions{DeleteRelatedActivity: true}); err != nil {
		t.Fatalf("DeleteSyllabusWithOptions error: %v", err)
	}
}

func TestResortSyllabusUsesPutEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/syllabus/resort" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.ResortSyllabus(context.Background(), map[string]any{"syllabuses": []int{1, 2}}); err != nil {
		t.Fatalf("ResortSyllabus error: %v", err)
	}
}

func TestGetCourseCustomScoreItemsUsesCoursesEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courses/15/custom-score-items" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"custom_score_items":[]}`), nil
	})

	resp, err := svc.GetCourseCustomScoreItems(context.Background(), 15)
	if err != nil {
		t.Fatalf("GetCourseCustomScoreItems error: %v", err)
	}
	if resp.CustomScoreItems == nil {
		t.Fatal("expected custom_score_items slice")
	}
}

func TestGetGroupSetUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/group-sets/6" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":6,"name":"gset"}`), nil
	})

	resp, err := svc.GetGroupSet(context.Background(), 6)
	if err != nil {
		t.Fatalf("GetGroupSet error: %v", err)
	}
	if resp.ID != 6 {
		t.Fatalf("unexpected group set id: %d", resp.ID)
	}
}

func TestGetCourseKnowledgeGraphUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-nodes" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[]}`), nil
	})

	resp, err := svc.GetCourseKnowledgeGraph(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetCourseKnowledgeGraph error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected knowledge node items slice")
	}
}

func TestGetKnowledgeNodeTreesUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-nodes" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[{"id":1,"name":"root"}]}`), nil
	})

	resp, err := svc.GetKnowledgeNodeTrees(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetKnowledgeNodeTrees error: %v", err)
	}
	if len(resp) != 1 || resp[0].ID != 1 {
		t.Fatalf("unexpected knowledge node trees: %#v", resp)
	}
}

func TestGetKnowledgeNodeUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-node/18" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":18,"name":"node"}`), nil
	})

	resp, err := svc.GetKnowledgeNode(context.Background(), 18)
	if err != nil {
		t.Fatalf("GetKnowledgeNode error: %v", err)
	}
	if resp.ID != 18 {
		t.Fatalf("unexpected node id: %d", resp.ID)
	}
}

func TestCreateKnowledgeNodeUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-node" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":19,"name":"new-node"}`), nil
	})

	resp, err := svc.CreateKnowledgeNode(context.Background(), 9, map[string]any{"name": "new-node"})
	if err != nil {
		t.Fatalf("CreateKnowledgeNode error: %v", err)
	}
	if resp.ID != 19 {
		t.Fatalf("unexpected node id: %d", resp.ID)
	}
}

func TestBatchDeleteKnowledgeNodesUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-node/delete" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.BatchDeleteKnowledgeNodes(context.Background(), 9, &DeleteKnowledgeNodesRequest{KnowledgeNodeIDs: []int{1, 2}}); err != nil {
		t.Fatalf("BatchDeleteKnowledgeNodes error: %v", err)
	}
}

func TestGetKnowledgeNodeRecommendedCapturesUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-node/18/recommended-captures" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.GetKnowledgeNodeRecommendedCaptures(context.Background(), 18, &model.ListOptions{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("GetKnowledgeNodeRecommendedCaptures error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected capture items slice")
	}
}

func TestGetMyCapturesUsesPagedEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/my-captures" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.GetMyCaptures(context.Background(), &model.ListOptions{Page: 1, PageSize: 10}, nil)
	if err != nil {
		t.Fatalf("GetMyCaptures error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected my captures items slice")
	}
}

func TestGetPublicCapturesUsesPagedEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/public-captures" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.GetPublicCaptures(context.Background(), &model.ListOptions{Page: 1, PageSize: 10}, nil)
	if err != nil {
		t.Fatalf("GetPublicCaptures error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected public captures items slice")
	}
}

func TestGetKnowledgeNodeRecommendedResourcesUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-node/18/recommended-resource-references" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.GetKnowledgeNodeRecommendedResourceReferences(context.Background(), 18, &model.ListOptions{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("GetKnowledgeNodeRecommendedResourceReferences error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected resource items slice")
	}
}

func TestGetKnowledgeNodeStatisticsSummaryUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-nodes/statistics/summary" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"node_count":1}`), nil
	})

	resp, err := svc.GetKnowledgeNodeStatisticsSummary(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetKnowledgeNodeStatisticsSummary error: %v", err)
	}
	if resp.NodeCount != 1 {
		t.Fatalf("unexpected node_count: %d", resp.NodeCount)
	}
}

func TestGetKnowledgeNodeWeekStatsListUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-stats-timeline" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`[{"id":1,"avg_completeness_rate":66.6,"avg_mastery_rate":77.7,"stat_date":"2025-01-01","week":1}]`), nil
	})

	resp, err := svc.GetKnowledgeNodeWeekStatsList(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetKnowledgeNodeWeekStatsList error: %v", err)
	}
	if len(resp) != 1 || resp[0].ID != 1 {
		t.Fatalf("unexpected week stats list: %#v", resp)
	}
}

func TestGetKnowledgeNodeWeekStatsUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-graph-snapshot/3" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"tree":{"id":1,"name":"root"},"relations":[{"id":1,"source":1,"target":2,"is_directed":true,"color":"#000"}],"completeness":[88.8],"mastery":[77.7]}`), nil
	})

	resp, err := svc.GetKnowledgeNodeWeekStats(context.Background(), 9, 3)
	if err != nil {
		t.Fatalf("GetKnowledgeNodeWeekStats error: %v", err)
	}
	if resp.Tree == nil || resp.Tree.ID != 1 {
		t.Fatalf("unexpected snapshot tree: %#v", resp)
	}
}

func TestGetKnowledgeNodeReferenceResourceUsesNodePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-nodes/18/reference-resource" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"resources":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.GetKnowledgeNodeReferenceResource(context.Background(), 18)
	if err != nil {
		t.Fatalf("GetKnowledgeNodeReferenceResource error: %v", err)
	}
	if resp.Resources == nil {
		t.Fatal("expected reference resources slice")
	}
}

func TestGetKnowledgeNodeStudentStatUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-nodes/student/18/stats" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"resource_count":2}`), nil
	})

	resp, err := svc.GetKnowledgeNodeStudentStat(context.Background(), 9, 18)
	if err != nil {
		t.Fatalf("GetKnowledgeNodeStudentStat error: %v", err)
	}
	if resp.ResourceCount != 2 {
		t.Fatalf("unexpected resource_count: %d", resp.ResourceCount)
	}
}

func TestGetNodeBaseContentByStudentUsesNodeStudentPath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-nodes/18/student/27/stat" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"node_name":"node-a","resource_count":3}`), nil
	})

	resp, err := svc.GetNodeBaseContentByStudent(context.Background(), 18, 27)
	if err != nil {
		t.Fatalf("GetNodeBaseContentByStudent error: %v", err)
	}
	if resp.ResourceCount != 3 {
		t.Fatalf("unexpected resource_count: %d", resp.ResourceCount)
	}
}

func TestGetKnowledgeNodeReferTypeStatsByCourseUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-nodes/statistics/refer-type" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"node_count":1,"resource_ref_count":2,"activity_ref_count":3}`), nil
	})

	resp, err := svc.GetKnowledgeNodeReferTypeStatsByCourse(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetKnowledgeNodeReferTypeStatsByCourse error: %v", err)
	}
	if resp.ActivityCount != 3 {
		t.Fatalf("unexpected activity_ref_count: %d", resp.ActivityCount)
	}
}

func TestGetKnowledgeNodeStatsByStudentDimensionUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-nodes/statistics/students-dimension" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.GetKnowledgeNodeStatsByStudentDimension(context.Background(), 9, &model.ListOptions{Page: 1, PageSize: 10}, "")
	if err != nil {
		t.Fatalf("GetKnowledgeNodeStatsByStudentDimension error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected student-dimension items slice")
	}
}

func TestGetKnowledgeNodeStatisticsOverviewUsesNodePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-nodes/18/statistics/overview" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":18,"name":"node"}`), nil
	})

	resp, err := svc.GetKnowledgeNodeStatisticsOverview(context.Background(), 18)
	if err != nil {
		t.Fatalf("GetKnowledgeNodeStatisticsOverview error: %v", err)
	}
	if resp.ID != 18 {
		t.Fatalf("unexpected overview id: %d", resp.ID)
	}
}

func TestGetKnowledgeNodeStatisticsResourceDetailUsesNodePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-nodes/18/statistics/resource-detail" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.GetKnowledgeNodeStatisticsResourceDetail(context.Background(), 18, &model.ListOptions{Page: 1, PageSize: 10}, "")
	if err != nil {
		t.Fatalf("GetKnowledgeNodeStatisticsResourceDetail error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected resource detail items slice")
	}
}

func TestGetKnowledgeNodeStatisticsActivityDetailUsesNodePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-nodes/18/statistics/activity-detail" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.GetKnowledgeNodeStatisticsActivityDetail(context.Background(), 18, &model.ListOptions{Page: 1, PageSize: 10}, "")
	if err != nil {
		t.Fatalf("GetKnowledgeNodeStatisticsActivityDetail error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected activity detail items slice")
	}
}

func TestGetKnowledgeNodeStudentResourceStatUsesNodeStudentPath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-nodes/18/student/27/resource/stat" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.GetKnowledgeNodeStudentResourceStat(context.Background(), 18, 27, &model.ListOptions{Page: 1, PageSize: 10}, "")
	if err != nil {
		t.Fatalf("GetKnowledgeNodeStudentResourceStat error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected student resource stat items slice")
	}
}

func TestGetKnowledgeNodeStudentActivityStatUsesNodeStudentPath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-nodes/18/student/27/activity/stat" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"items":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.GetKnowledgeNodeStudentActivityStat(context.Background(), 18, 27, &model.ListOptions{Page: 1, PageSize: 10}, "")
	if err != nil {
		t.Fatalf("GetKnowledgeNodeStudentActivityStat error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected student activity stat items slice")
	}
}

func TestGetCurrentStudentNodeOverallStatisticsUsesCourseStudentPath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courses/9/students/18/knowledge-nodes/overall-stats" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"overall_completeness_rate":"50%","overall_mastery_rate":"40%"}`), nil
	})

	resp, err := svc.GetCurrentStudentNodeOverallStatistics(context.Background(), 9, 18)
	if err != nil {
		t.Fatalf("GetCurrentStudentNodeOverallStatistics error: %v", err)
	}
	if resp.OverallMasteryRate != "40%" {
		t.Fatalf("unexpected overall_mastery_rate: %s", resp.OverallMasteryRate)
	}
}

func TestGetCourseTeachingObjectivesUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/teaching-objective" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`[{"id":1,"content":"obj"}]`), nil
	})

	resp, err := svc.GetCourseTeachingObjectives(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetCourseTeachingObjectives error: %v", err)
	}
	if len(resp) != 1 || resp[0].ID != 1 {
		t.Fatalf("unexpected teaching objectives: %#v", resp)
	}
}

func TestCreateTeachingObjectivesUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/teaching-objective" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.CreateTeachingObjectives(context.Background(), 9, &TeachingObjectivesRequest{TeachingObjectives: []*TeachingObjective{{Content: "obj"}}})
	if err != nil {
		t.Fatalf("CreateTeachingObjectives error: %v", err)
	}
}

func TestUpdateTeachingObjectivesUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/teaching-objective" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.UpdateTeachingObjectives(context.Background(), 9, &TeachingObjectivesRequest{TeachingObjectives: []*TeachingObjective{{ID: 1, Content: "obj"}}})
	if err != nil {
		t.Fatalf("UpdateTeachingObjectives error: %v", err)
	}
}

func TestDeleteTeachingObjectivesUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/teaching-objective/delete" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.DeleteTeachingObjectives(context.Background(), 9, &DeleteTeachingObjectivesRequest{TeachingObjectiveIDs: []int{1, 2}})
	if err != nil {
		t.Fatalf("DeleteTeachingObjectives error: %v", err)
	}
}

func TestGetNodeFacetsUsesCourseNodePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-node/18/facets" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`[{"id":1,"name":"facet"}]`), nil
	})

	resp, err := svc.GetNodeFacets(context.Background(), 9, 18)
	if err != nil {
		t.Fatalf("GetNodeFacets error: %v", err)
	}
	if len(resp) != 1 || resp[0].ID != 1 {
		t.Fatalf("unexpected facets: %#v", resp)
	}
}

func TestGetFragmentsUsesCourseNodePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-node/18/fragments" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`[{"id":1,"content":"fragment"}]`), nil
	})

	resp, err := svc.GetFragments(context.Background(), 9, 18)
	if err != nil {
		t.Fatalf("GetFragments error: %v", err)
	}
	if len(resp) != 1 || resp[0].ID != 1 {
		t.Fatalf("unexpected fragments: %#v", resp)
	}
}

func TestGetKnowledgeReferencesForActivityUsesActivityPath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/activities/7/knowledge-references" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`[]`), nil
	})

	resp, err := svc.GetKnowledgeReferencesForActivity(context.Background(), 7)
	if err != nil {
		t.Fatalf("GetKnowledgeReferencesForActivity error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected knowledge references slice")
	}
}

func TestSaveKnowledgeReferencesForActivityUsesActivityPath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/activities/7/knowledge-references" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.SaveKnowledgeReferencesForActivity(context.Background(), 7, map[string]any{"knowledge_node_ids": []int{1, 2}}); err != nil {
		t.Fatalf("SaveKnowledgeReferencesForActivity error: %v", err)
	}
}

func TestRemoveMediaKnowledgeReferenceForActivityUsesActivityPath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/activities/7/knowledge-references/11" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("chapter_id"); got != "13" {
			t.Fatalf("unexpected chapter_id: %s", got)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.RemoveMediaKnowledgeReferenceForActivity(context.Background(), 7, 11, &DeleteMediaKnowledgeReferenceRequest{ChapterID: 13})
	if err != nil {
		t.Fatalf("RemoveMediaKnowledgeReferenceForActivity error: %v", err)
	}
}

func TestRemoveUploadKnowledgeReferenceForActivityUsesActivityPath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/activities/7/knowledge-references/11" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.RemoveUploadKnowledgeReferenceForActivity(context.Background(), 7, 11, &DeleteUploadKnowledgeReferenceRequest{UploadID: 17})
	if err != nil {
		t.Fatalf("RemoveUploadKnowledgeReferenceForActivity error: %v", err)
	}
}

func TestUpdateKnowledgeNodeUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-node/18" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":18,"name":"updated"}`), nil
	})

	resp, err := svc.UpdateKnowledgeNode(context.Background(), 18, map[string]any{"name": "updated"})
	if err != nil {
		t.Fatalf("UpdateKnowledgeNode error: %v", err)
	}
	if resp.ID != 18 {
		t.Fatalf("unexpected node id: %d", resp.ID)
	}
}

func TestAddKnowledgeNodeRelationUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-node/relation" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":1,"source":2,"target":3,"is_directed":true}`), nil
	})

	resp, err := svc.AddKnowledgeNodeRelation(context.Background(), 9, &KnowledgeNodeRelationRequest{Source: 2, Target: 3, IsDirected: true})
	if err != nil {
		t.Fatalf("AddKnowledgeNodeRelation error: %v", err)
	}
	if resp.ID != 1 {
		t.Fatalf("unexpected relation id: %d", resp.ID)
	}
}

func TestUpdateKnowledgeNodeRelationUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-node/relation" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":1,"source":2,"target":3,"is_directed":false}`), nil
	})

	resp, err := svc.UpdateKnowledgeNodeRelation(context.Background(), 9, &KnowledgeNodeRelationRequest{ID: 1, Source: 2, Target: 3})
	if err != nil {
		t.Fatalf("UpdateKnowledgeNodeRelation error: %v", err)
	}
	if resp.ID != 1 {
		t.Fatalf("unexpected relation id: %d", resp.ID)
	}
}

func TestDeleteKnowledgeNodeRelationsUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-node/relation/delete" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.DeleteKnowledgeNodeRelations(context.Background(), 9, &DeleteKnowledgeNodeRelationsRequest{RelationIDs: []int{1, 2}})
	if err != nil {
		t.Fatalf("DeleteKnowledgeNodeRelations error: %v", err)
	}
}

func TestListCoursewareActivitiesUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/courseware-activities" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`[{"id":1,"title":"act","type":"lesson"}]`), nil
	})

	resp, err := svc.ListCoursewareActivities(context.Background(), 9)
	if err != nil {
		t.Fatalf("ListCoursewareActivities error: %v", err)
	}
	if len(resp) != 1 || resp[0].ID != 1 {
		t.Fatalf("unexpected courseware activities: %#v", resp)
	}
}

func TestMoveKnowledgeNodeUsesCoursePath(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-node/move" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.MoveKnowledgeNode(context.Background(), 9, &MoveKnowledgeNodeRequest{ID: 1, Sort: 2})
	if err != nil {
		t.Fatalf("MoveKnowledgeNode error: %v", err)
	}
}

func TestGetKfsSubjectsUsesKnowledgeGraphEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-graph/kfs-subjects" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`[{"id":1,"name":"subject","courses":[]}]`), nil
	})

	resp, err := svc.GetKfsSubjects(context.Background())
	if err != nil {
		t.Fatalf("GetKfsSubjects error: %v", err)
	}
	if len(resp) != 1 || resp[0].ID != 1 {
		t.Fatalf("unexpected KFS subjects: %#v", resp)
	}
}

func TestGetKfsImportInfoUsesKnowledgeGraphEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-graph/courses/9/kfs-import-info" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"last_import_time":1700000000,"import_course_ids":[1,2],"server":"kfs"}`), nil
	})

	resp, err := svc.GetKfsImportInfo(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetKfsImportInfo error: %v", err)
	}
	if resp.Server != "kfs" {
		t.Fatalf("unexpected server: %s", resp.Server)
	}
}

func TestImportKfsCourseUsesKnowledgeGraphEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-graph/courses/9/kfs-course-import" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.ImportKfsCourse(context.Background(), 9, &ImportKfsCourseRequest{KFSCourseID: 11, KFSVersionID: 22})
	if err != nil {
		t.Fatalf("ImportKfsCourse error: %v", err)
	}
}

func TestGetKfsCourseForestVersionStatsUsesKnowledgeGraphEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-graph/kfs-courses/11/forest-versions/22/stats" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"topic_count":1,"dependency_count":2,"facet_count":3,"fragment_count":4}`), nil
	})

	resp, err := svc.GetKfsCourseForestVersionStats(context.Background(), 11, 22)
	if err != nil {
		t.Fatalf("GetKfsCourseForestVersionStats error: %v", err)
	}
	if resp.FragmentCount != 4 {
		t.Fatalf("unexpected fragment_count: %d", resp.FragmentCount)
	}
}

func TestGetChinamCloudKnowledgeGraphDiffUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-graphs/diff" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`[{"action":"update","type":"node","node_id":1,"node_name":"n1"}]`), nil
	})

	resp, err := svc.GetChinamCloudKnowledgeGraphDiff(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetChinamCloudKnowledgeGraphDiff error: %v", err)
	}
	if len(resp) != 1 || resp[0].NodeID != 1 {
		t.Fatalf("unexpected knowledge graph diff: %#v", resp)
	}
}

func TestBatchGetPublishedForestVersionByKFSCourseIDsUsesKnowledgeGraphEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-graph/kfs-courses/-/published-forest-versions:batchGet" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("ids"); got != "1,2" {
			t.Fatalf("unexpected ids query: %s", got)
		}
		return jsonResponse(`[{"id":10,"name":"v1","course_id":1,"published":true}]`), nil
	})

	resp, err := svc.BatchGetPublishedForestVersionByKFSCourseIDs(context.Background(), []int{1, 2})
	if err != nil {
		t.Fatalf("BatchGetPublishedForestVersionByKFSCourseIDs error: %v", err)
	}
	if len(resp) != 1 || resp[0].ID != 10 {
		t.Fatalf("unexpected published forest versions: %#v", resp)
	}
}

func TestBatchGetForestVersionStatsByKFSVersionIDsUsesKnowledgeGraphEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/knowledge-graph/forest-versions/-/stats:batchGet" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("ids"); got != "3,4" {
			t.Fatalf("unexpected ids query: %s", got)
		}
		return jsonResponse(`[{"id":3,"topic_count":1,"dependency_count":2,"facet_count":3,"fragment_count":4}]`), nil
	})

	resp, err := svc.BatchGetForestVersionStatsByKFSVersionIDs(context.Background(), []int{3, 4})
	if err != nil {
		t.Fatalf("BatchGetForestVersionStatsByKFSVersionIDs error: %v", err)
	}
	if len(resp) != 1 || resp[0].ID != 3 {
		t.Fatalf("unexpected forest version stats: %#v", resp)
	}
}

func TestGetKnowledgeGraphSimilarityUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-graphs/similarity" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"similarities":[{"id":1,"name":"node-a","similarity":0.8,"similarity_format":"80%","checked":true}]}`), nil
	})

	resp, err := svc.GetKnowledgeGraphSimilarity(context.Background(), 9, map[string]any{"node_ids": []int{1, 2}})
	if err != nil {
		t.Fatalf("GetKnowledgeGraphSimilarity error: %v", err)
	}
	if len(resp) != 1 || resp[0].ID != 1 {
		t.Fatalf("unexpected knowledge graph similarities: %#v", resp)
	}
}

func TestSyncChinamCloudKnowledgeGraphUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-graph/sync" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.SyncChinamCloudKnowledgeGraph(context.Background(), 9); err != nil {
		t.Fatalf("SyncChinamCloudKnowledgeGraph error: %v", err)
	}
}

func TestEditKnowledgeGraphSourceUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-graph/source" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.EditKnowledgeGraphSource(context.Background(), 9, &UpdateKnowledgeGraphSourceRequest{Source: "chinamcloud"})
	if err != nil {
		t.Fatalf("EditKnowledgeGraphSource error: %v", err)
	}
}

func TestSyncChinamCloudKnowledgeGraphDiffUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-graphs/diff/sync" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.SyncChinamCloudKnowledgeGraphDiff(context.Background(), 9, map[string]any{"node_ids": []int{1}})
	if err != nil {
		t.Fatalf("SyncChinamCloudKnowledgeGraphDiff error: %v", err)
	}
}

func TestReplaceChinamKnowledgeGraphUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPut {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-graph/replace" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.ReplaceChinamKnowledgeGraph(context.Background(), 9); err != nil {
		t.Fatalf("ReplaceChinamKnowledgeGraph error: %v", err)
	}
}

func TestImportKnowledgeNodesUsesCoursesEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courses/9/knowledge-nodes/import" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.ImportKnowledgeNodes(context.Background(), 9, map[string]any{"data": []map[string]any{{"name": "n1"}}})
	if err != nil {
		t.Fatalf("ImportKnowledgeNodes error: %v", err)
	}
}

func TestImportKnowledgeNodesByCourseUsesCoursesEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courses/9/knowledge-nodes/import" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	err := svc.ImportKnowledgeNodesByCourse(context.Background(), 9, &ImportKnowledgeNodesByCourseRequest{
		SourceCourseID:      11,
		ImportReferResource: true,
	})
	if err != nil {
		t.Fatalf("ImportKnowledgeNodesByCourse error: %v", err)
	}
}

func TestCancelPreviewUsesCoursesEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courses/9/preview" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.CancelPreview(context.Background(), 9); err != nil {
		t.Fatalf("CancelPreview error: %v", err)
	}
}

func TestGetKnowledgeGraphEmbedURLUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-graph/embed/cluster-graph" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"url":"https://example.com/embed"}`), nil
	})

	resp, err := svc.GetKnowledgeGraphEmbedURL(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetKnowledgeGraphEmbedURL error: %v", err)
	}
	if resp != "https://example.com/embed" {
		t.Fatalf("unexpected embed url: %s", resp)
	}
}

func TestGetChinamCloudGraphEditURLUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/knowledge-graph/edit-url" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"url":"https://example.com/edit"}`), nil
	})

	resp, err := svc.GetChinamCloudGraphEditURL(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetChinamCloudGraphEditURL error: %v", err)
	}
	if resp != "https://example.com/edit" {
		t.Fatalf("unexpected edit url: %s", resp)
	}
}

func TestGetChinamCloudResourceViewURLUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/course/9/external-resource/12/preview-url" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"url":"https://example.com/view"}`), nil
	})

	resp, err := svc.GetChinamCloudResourceViewURL(context.Background(), 9, 12)
	if err != nil {
		t.Fatalf("GetChinamCloudResourceViewURL error: %v", err)
	}
	if resp != "https://example.com/view" {
		t.Fatalf("unexpected preview url: %s", resp)
	}
}

func TestGetFileStatusUsesUploadsEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/uploads/33" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"status":"ready"}`), nil
	})

	resp, err := svc.GetFileStatus(context.Background(), 33)
	if err != nil {
		t.Fatalf("GetFileStatus error: %v", err)
	}
	if resp != "ready" {
		t.Fatalf("unexpected file status: %s", resp)
	}
}
