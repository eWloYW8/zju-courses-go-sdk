package aircredit

import (
	"context"
	"encoding/json"
	"fmt"
)

// AI Agent/Chatbot related methods are added to Service.

// --- AI Agent ---

// GetRecentAIAgentCourses returns recent courses for the AI agent.
func (s *Service) GetRecentAIAgentCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-agent/recent-courses", &result)
	return result, err
}

// GetRecentAIConversations returns recent AI conversations.
func (s *Service) GetRecentAIConversations(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-agent/recent-conversations", &result)
	return result, err
}

// GetCourseChatUsageInfo returns AI chat usage info for a course.
func (s *Service) GetCourseChatUsageInfo(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/course/%d/chat-usage-info", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCourseCreditToken returns the AI credit token for a course.
func (s *Service) GetCourseCreditToken(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/course/%d/token", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// UploadAIScreenShot uploads a screenshot for AI processing.
func (s *Service) UploadAIScreenShot(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/uploads/air-screen-shot", body, &result)
	return result, err
}

// --- Knowledge Nodes ---

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

// GetCurrentSemesterCourses returns courses for the current semester.
func (s *Service) GetCurrentSemesterCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/current-semester-courses", &result)
	return result, err
}
