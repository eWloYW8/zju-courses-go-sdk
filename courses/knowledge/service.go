package knowledge

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

// Service handles knowledge-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
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
func (s *Service) GetKnowledgeNodeTrees(ctx context.Context, courseID int) ([]*KnowledgeNode, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-nodes", courseID)
	result := new(KnowledgeNodesResponse)
	_, err := s.client.Get(ctx, u, result)
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

// GetKnowledgeNode returns a knowledge node.
func (s *Service) GetKnowledgeNode(ctx context.Context, nodeID int) (*KnowledgeNode, error) {
	u := fmt.Sprintf("/api/knowledge-node/%d", nodeID)
	result := new(KnowledgeNode)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateKnowledgeNode creates a knowledge node in a course.
func (s *Service) CreateKnowledgeNode(ctx context.Context, courseID int, body interface{}) (*KnowledgeNode, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-node", courseID)
	result := new(KnowledgeNode)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// UpdateKnowledgeNode updates a knowledge node.
func (s *Service) UpdateKnowledgeNode(ctx context.Context, nodeID int, body interface{}) (*KnowledgeNode, error) {
	u := fmt.Sprintf("/api/knowledge-node/%d", nodeID)
	result := new(KnowledgeNode)
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

// --- AI Knowledge Node Methods ---

// CheckKnowledgeNodesExist checks if knowledge nodes exist for a course.
func (s *Service) CheckKnowledgeNodesExist(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-nodes/exists", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// SyncKnowledgeNodesWithAI syncs knowledge nodes with AI for a course.
func (s *Service) SyncKnowledgeNodesWithAI(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/course/%d/knowledge-nodes/sync-air", courseID)
	var result json.RawMessage
	_, err := s.client.Post(ctx, u, nil, &result)
	return result, err
}

// GetKnowledgeGraphSettings returns knowledge graph settings.
func (s *Service) GetKnowledgeGraphSettings(ctx context.Context, graphID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/knowledge-graph/%d/settings", graphID)
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

// intsToCSV converts a slice of ints to a comma-separated string.
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
