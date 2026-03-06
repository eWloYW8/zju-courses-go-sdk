package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// AIAbilityResponse represents the response from the AI ability check.
type AIAbilityResponse struct {
	HasAnyCourseAIAbility bool `json:"has_any_course_ai_ability"`
}

// AirCreditService handles AI credit-related API operations.
type AirCreditService struct {
	client *Client
}

// --- User AI Credit ---

// HasAIAbility checks if any of the user's courses have AI ability.
func (s *AirCreditService) HasAIAbility(ctx context.Context) (*AIAbilityResponse, error) {
	result := new(AIAbilityResponse)
	_, err := s.client.get(ctx, "/api/air-credit/user/courses/ai-ability", result)
	return result, err
}

// GetUserCreditStates returns the user's credit states.
func (s *AirCreditService) GetUserCreditStates(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/user/credit-states", &result)
	return result, err
}

// GetUserToken returns the user's AI token.
func (s *AirCreditService) GetUserToken(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/user/token", &result)
	return result, err
}

// GetUserInfo returns the user's AI credit info.
func (s *AirCreditService) GetUserInfo(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/user", &result)
	return result, err
}

// ExportUserCreditStatesStats exports the user's credit states stats.
func (s *AirCreditService) ExportUserCreditStatesStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/user/credit-states-stats/export", &result)
	return result, err
}

// --- Course AI Credit ---

// GetCourseCreditInfo returns AI credit info for a course.
func (s *AirCreditService) GetCourseCreditInfo(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/course/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetCourseCreditStates returns credit states for a course.
func (s *AirCreditService) GetCourseCreditStates(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/course/credit-states", &result)
	return result, err
}

// ExportCourseCreditStatesStats exports course credit states stats.
func (s *AirCreditService) ExportCourseCreditStatesStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/course/credit-states-stats/export", &result)
	return result, err
}

// GetCourseUsageLimit returns the AI usage limit for a course.
func (s *AirCreditService) GetCourseUsageLimit(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/course/usage-limit", &result)
	return result, err
}

// GetCourseAIAbility returns AI ability for a specific course.
func (s *AirCreditService) GetCourseAIAbility(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/courses/%d", courseID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- Admin AI Credit ---

// GetCreditsStatus returns the overall AI credits status.
func (s *AirCreditService) GetCreditsStatus(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/credits/status", &result)
	return result, err
}

// GetCredits returns AI credits information.
func (s *AirCreditService) GetCredits(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/credits", &result)
	return result, err
}

// ClearRemainingCredits clears remaining credits.
func (s *AirCreditService) ClearRemainingCredits(ctx context.Context) error {
	_, err := s.client.post(ctx, "/api/air-credit/credits/clear-remaining-credits", nil, nil)
	return err
}

// GetCreditStatesSummary returns credit states summary.
func (s *AirCreditService) GetCreditStatesSummary(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/credit-states-summary", &result)
	return result, err
}

// GetCreditStatesStats returns credit states statistics.
func (s *AirCreditService) GetCreditStatesStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/credit-states-stats", &result)
	return result, err
}

// GetOrgCreditStateInfo returns org credit state info.
func (s *AirCreditService) GetOrgCreditStateInfo(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/org/credit-state-info", &result)
	return result, err
}

// ListAudits returns AI credit audit records.
func (s *AirCreditService) ListAudits(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/air-credit/audits", opts)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetAudit returns a specific audit record.
func (s *AirCreditService) GetAudit(ctx context.Context, auditID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/audits/%d", auditID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetInstructorCurrentSemesterCourses returns instructor courses for current semester.
func (s *AirCreditService) GetInstructorCurrentSemesterCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/air-credit/instructors/current-semester-courses", &result)
	return result, err
}

// GetResourceInfo returns AI resource info.
func (s *AirCreditService) GetResourceInfo(ctx context.Context, resourceID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/resources/%d", resourceID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// GetUserAIInfo returns AI info for a user.
func (s *AirCreditService) GetUserAIInfo(ctx context.Context, userID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/users/%d", userID)
	var result json.RawMessage
	_, err := s.client.get(ctx, u, &result)
	return result, err
}

// --- AI PPT ---

// GetAIPPTUsage returns AI PPT usage.
func (s *AirCreditService) GetAIPPTUsage(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/ai-ppt/usage", &result)
	return result, err
}

// GetAIPPTUsageStats returns AI PPT usage statistics.
func (s *AirCreditService) GetAIPPTUsageStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/ai-ppt/usage/stats", &result)
	return result, err
}

// GetAIPPTUserUsageCount returns AI PPT user usage count.
func (s *AirCreditService) GetAIPPTUserUsageCount(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/ai-ppt/user/usage/count", &result)
	return result, err
}

// ExportAIPPTUserUsage exports AI PPT user usage.
func (s *AirCreditService) ExportAIPPTUserUsage(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/ai-ppt/user-usage/export", &result)
	return result, err
}

// --- Text Optimization ---

// OptimizeText optimizes text using AI.
func (s *AirCreditService) OptimizeText(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.post(ctx, "/api/text-optimization", body, &result)
	return result, err
}
