package aircredit

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

// Service handles AI credit-related API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- User AI Credit ---

// HasAIAbility checks if any of the user's courses have AI ability.
func (s *Service) HasAIAbility(ctx context.Context) (*AIAbilityResponse, error) {
	result := new(AIAbilityResponse)
	_, err := s.client.Get(ctx, "/api/air-credit/user/courses/ai-ability", result)
	return result, err
}

// GetUserCreditStates returns the user's credit states.
func (s *Service) GetUserCreditStates(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/user/credit-states", &result)
	return result, err
}

// ListUserCreditStates returns paged user credit states.
func (s *Service) ListUserCreditStates(ctx context.Context, params ListCreditStatesParams) (json.RawMessage, error) {
	u := addListOptions("/api/air-credit/user/credit-states", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListUserCreditStatesTyped returns paged user credit states with typed items.
func (s *Service) ListUserCreditStatesTyped(ctx context.Context, params ListCreditStatesParams) (*UserCreditStatesResponse, error) {
	u := addListOptions("/api/air-credit/user/credit-states", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	result := new(UserCreditStatesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetUserToken returns the user's AI token.
func (s *Service) GetUserToken(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/user/token", &result)
	return result, err
}

// GetUserInfo returns the user's AI credit info.
func (s *Service) GetUserInfo(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/user", &result)
	return result, err
}

// ExportUserCreditStatesStats exports the user's credit states stats.
func (s *Service) ExportUserCreditStatesStats(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/air-credit/user/credit-states-stats/export", body, &result)
	return result, err
}

// --- Course AI Credit ---

// GetCourseCreditInfo returns AI credit info for a course.
func (s *Service) GetCourseCreditInfo(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/course/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCourseCreditStates returns credit states for a course.
func (s *Service) GetCourseCreditStates(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/course/credit-states", &result)
	return result, err
}

// ListCourseCreditStates returns paged course credit states.
func (s *Service) ListCourseCreditStates(ctx context.Context, params ListCreditStatesParams) (json.RawMessage, error) {
	u := addListOptions("/api/air-credit/course/credit-states", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListCourseCreditStatesTyped returns paged course credit states with typed items.
func (s *Service) ListCourseCreditStatesTyped(ctx context.Context, params ListCreditStatesParams) (*CourseCreditStatesResponse, error) {
	u := addListOptions("/api/air-credit/course/credit-states", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	result := new(CourseCreditStatesResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ExportCourseCreditStatesStats exports course credit states stats.
func (s *Service) ExportCourseCreditStatesStats(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/air-credit/course/credit-states-stats/export", body, &result)
	return result, err
}

// GetCourseUsageLimit returns the AI usage limit for a course.
func (s *Service) GetCourseUsageLimit(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/course/usage-limit", &result)
	return result, err
}

// UpdateCourseUsageLimit updates the AI usage limit for courses.
func (s *Service) UpdateCourseUsageLimit(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Put(ctx, "/api/air-credit/course/usage-limit", body, &result)
	return result, err
}

// GetCourseAIAbility returns AI ability for a specific course.
func (s *Service) GetCourseAIAbility(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/courses/%d", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- Admin AI Credit ---

// GetCreditsStatus returns the overall AI credits status.
func (s *Service) GetCreditsStatus(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/credits/status", &result)
	return result, err
}

// GetCredits returns AI credits information.
func (s *Service) GetCredits(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/credits", &result)
	return result, err
}

// ClearRemainingCredits clears remaining credits.
func (s *Service) ClearRemainingCredits(ctx context.Context) error {
	_, err := s.client.Put(ctx, "/api/air-credit/credits/clear-remaining-credits", nil, nil)
	return err
}

// AssignCredits creates AI credit assignments.
func (s *Service) AssignCredits(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/air-credit/credits", body, &result)
	return result, err
}

// UpdateCredits updates AI credit assignments.
func (s *Service) UpdateCredits(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Put(ctx, "/api/air-credit/credits", body, &result)
	return result, err
}

// UpdateCreditsStatus updates AI credit assignment status.
func (s *Service) UpdateCreditsStatus(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Put(ctx, "/api/air-credit/credits/status", body, &result)
	return result, err
}

// ClearRemainingCreditsWithBody clears remaining credits for a specific assignment payload.
func (s *Service) ClearRemainingCreditsWithBody(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Put(ctx, "/api/air-credit/credits/clear-remaining-credits", body, &result)
	return result, err
}

// GetCreditStatesSummary returns credit states summary.
func (s *Service) GetCreditStatesSummary(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/credit-states-summary", &result)
	return result, err
}

// GetCreditStatesSummaryByType returns credit states summary for a specific target type.
func (s *Service) GetCreditStatesSummaryByType(ctx context.Context, targetType string) (json.RawMessage, error) {
	u := "/api/air-credit/credit-states-summary"
	if targetType != "" {
		u = addQueryParams(u, map[string]string{"type": targetType})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCreditStatesStats returns credit states statistics.
func (s *Service) GetCreditStatesStats(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/credit-states-stats", &result)
	return result, err
}

// GetUserCreditStatesStats returns paged user credit statistics.
func (s *Service) GetUserCreditStatesStats(ctx context.Context, params CreditStateStatsParams) (json.RawMessage, error) {
	return s.getCreditStatesStats(ctx, "user", params)
}

// GetUserCreditStatesStatsTyped returns paged user credit usage statistics.
func (s *Service) GetUserCreditStatesStatsTyped(ctx context.Context, params CreditStateStatsParams) (*UserCreditUsageStatsResponse, error) {
	return s.getCreditStatesStatsTyped(ctx, "user", params)
}

// GetCourseCreditStatesStats returns paged course credit statistics.
func (s *Service) GetCourseCreditStatesStats(ctx context.Context, params CreditStateStatsParams) (json.RawMessage, error) {
	return s.getCreditStatesStats(ctx, "course", params)
}

// GetCourseCreditStatesStatsTyped returns paged course credit usage statistics.
func (s *Service) GetCourseCreditStatesStatsTyped(ctx context.Context, params CreditStateStatsParams) (*CourseCreditUsageStatsResponse, error) {
	return s.getCourseCreditStatesStatsTyped(ctx, "course", params)
}

// GetOrgCreditStateInfo returns org credit state info.
func (s *Service) GetOrgCreditStateInfo(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/org/credit-state-info", &result)
	return result, err
}

// GetOrgCreditStateInfoTyped returns typed org-level credit allocation info.
func (s *Service) GetOrgCreditStateInfoTyped(ctx context.Context) (*OrgCreditStateInfo, error) {
	result := new(OrgCreditStateInfo)
	_, err := s.client.Get(ctx, "/api/air-credit/org/credit-state-info", result)
	return result, err
}

// ListAudits returns AI credit audit records.
func (s *Service) ListAudits(ctx context.Context, opts *model.ListOptions) (json.RawMessage, error) {
	u := addListOptions("/api/air-credit/audits", opts)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListAuditsWithParams returns AI credit audit records with paging and filters.
func (s *Service) ListAuditsWithParams(ctx context.Context, params ListAuditsParams) (json.RawMessage, error) {
	u := addListOptions("/api/air-credit/audits", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// ListAuditsTyped returns AI credit audit records with typed items.
func (s *Service) ListAuditsTyped(ctx context.Context, params ListAuditsParams) (*CreditAuditsResponse, error) {
	u := addListOptions("/api/air-credit/audits", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	result := new(CreditAuditsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateAudit creates an AI credit audit record.
func (s *Service) CreateAudit(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/air-credit/audits", body, &result)
	return result, err
}

// GetAudit returns a specific audit record.
func (s *Service) GetAudit(ctx context.Context, auditID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/audits/%d", auditID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// MarkAuditRead marks an audit record as read or unread.
func (s *Service) MarkAuditRead(ctx context.Context, auditID int, read bool) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/audits/%d/read", auditID)
	var result json.RawMessage
	_, err := s.client.Put(ctx, u, map[string]bool{"read": read}, &result)
	return result, err
}

// GetUserAudit returns a user's AI credit audit view.
func (s *Service) GetUserAudit(ctx context.Context, userID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/users/%d/audit", userID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCourseAudit returns a course AI credit audit view.
func (s *Service) GetCourseAudit(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/courses/%d/audit", courseID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetResourceCredits returns AI credit info for a resource.
func (s *Service) GetResourceCredits(ctx context.Context, resourceID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/resources/%d/credits", resourceID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetCourseCredit returns AI credit info for a course via query endpoint.
func (s *Service) GetCourseCredit(ctx context.Context, courseID int) (json.RawMessage, error) {
	u := addQueryParams("/api/air-credit/course", map[string]string{"course_id": fmt.Sprintf("%d", courseID)})
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetInstructorCurrentSemesterCourses returns instructor courses for current semester.
func (s *Service) GetInstructorCurrentSemesterCourses(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Get(ctx, "/api/air-credit/instructors/current-semester-courses", &result)
	return result, err
}

// GetInstructorCurrentSemesterCoursesByUserIDs returns current semester courses for selected instructors.
func (s *Service) GetInstructorCurrentSemesterCoursesByUserIDs(ctx context.Context, userIDs []int) (json.RawMessage, error) {
	u := "/api/air-credit/instructors/current-semester-courses"
	if len(userIDs) > 0 {
		u = addQueryParams(u, map[string]string{"user_ids": joinInts(userIDs)})
	}
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetResourceInfo returns AI resource info.
func (s *Service) GetResourceInfo(ctx context.Context, resourceID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/resources/%d", resourceID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// GetUserAIInfo returns AI info for a user.
func (s *Service) GetUserAIInfo(ctx context.Context, userID int) (json.RawMessage, error) {
	u := fmt.Sprintf("/api/air-credit/users/%d", userID)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

// --- AI PPT ---

// GetAIPPTUsage returns AI PPT usage.
func (s *Service) GetAIPPTUsage(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/ai-ppt/usage", nil, &result)
	return result, err
}

// GetAIPPTUsageStats returns AI PPT usage statistics.
func (s *Service) GetAIPPTUsageStats(ctx context.Context, params ListAIPPTUsageStatsParams) (*AIPPTUsageStatsResponse, error) {
	u := addListOptions("/api/ai-ppt/usage/stats", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		u = addQueryParams(u, map[string]string{"conditions": encoded})
	}
	result := new(AIPPTUsageStatsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// GetAIPPTUserUsageCount returns AI PPT user usage count.
func (s *Service) GetAIPPTUserUsageCount(ctx context.Context) (*AIPPTUserUsageCountResponse, error) {
	result := new(AIPPTUserUsageCountResponse)
	_, err := s.client.Get(ctx, "/api/ai-ppt/user/usage/count", &result)
	return result, err
}

// ExportAIPPTUserUsage exports AI PPT user usage.
func (s *Service) ExportAIPPTUserUsage(ctx context.Context, body ExportAIPPTUserUsageRequest) ([]byte, error) {
	req, err := s.client.NewRequest(ctx, "POST", "/api/ai-ppt/user-usage/export", body)
	if err != nil {
		return nil, err
	}
	_, blob, err := s.client.DoBytes(req)
	return blob, err
}

// --- Text Optimization ---

// OptimizeText optimizes text using AI.
func (s *Service) OptimizeText(ctx context.Context, body interface{}) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.Post(ctx, "/api/text-optimization", body, &result)
	return result, err
}

func addListOptions(urlStr string, opts *model.ListOptions) string {
	if opts == nil {
		return urlStr
	}
	return sdk.AddListOptions(urlStr, opts.Page, opts.PageSize)
}

func addQueryParams(urlStr string, params map[string]string) string {
	return sdk.AddQueryParams(urlStr, params)
}

func encodeConditions(conditions any) string {
	switch value := conditions.(type) {
	case nil:
		return ""
	case string:
		return value
	default:
		encoded, err := json.Marshal(value)
		if err != nil {
			return ""
		}
		return string(encoded)
	}
}

func (s *Service) getCreditStatesStats(ctx context.Context, statType string, params CreditStateStatsParams) (json.RawMessage, error) {
	u := addListOptions("/api/air-credit/credit-states-stats", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	query := map[string]string{"type": statType}
	if params.StartDate != "" {
		query["start_date"] = params.StartDate
	}
	if params.EndDate != "" {
		query["end_date"] = params.EndDate
	}
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		query["conditions"] = encoded
	}
	u = addQueryParams(u, query)
	var result json.RawMessage
	_, err := s.client.Get(ctx, u, &result)
	return result, err
}

func (s *Service) getCreditStatesStatsTyped(ctx context.Context, statType string, params CreditStateStatsParams) (*UserCreditUsageStatsResponse, error) {
	u := addListOptions("/api/air-credit/credit-states-stats", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	query := map[string]string{"type": statType}
	if params.StartDate != "" {
		query["start_date"] = params.StartDate
	}
	if params.EndDate != "" {
		query["end_date"] = params.EndDate
	}
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		query["conditions"] = encoded
	}
	u = addQueryParams(u, query)
	result := new(UserCreditUsageStatsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

func (s *Service) getCourseCreditStatesStatsTyped(ctx context.Context, statType string, params CreditStateStatsParams) (*CourseCreditUsageStatsResponse, error) {
	u := addListOptions("/api/air-credit/credit-states-stats", &model.ListOptions{Page: params.Page, PageSize: params.PageSize})
	query := map[string]string{"type": statType}
	if params.StartDate != "" {
		query["start_date"] = params.StartDate
	}
	if params.EndDate != "" {
		query["end_date"] = params.EndDate
	}
	if encoded := encodeConditions(params.Conditions); encoded != "" {
		query["conditions"] = encoded
	}
	u = addQueryParams(u, query)
	result := new(CourseCreditUsageStatsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

func joinInts(values []int) string {
	if len(values) == 0 {
		return ""
	}
	out := fmt.Sprintf("%d", values[0])
	for i := 1; i < len(values); i++ {
		out += fmt.Sprintf(",%d", values[i])
	}
	return out
}
