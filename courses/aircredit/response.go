package aircredit

import (
	"encoding/json"

	"github.com/eWloYW8/zju-courses-go-sdk/courses/model"
)

type CreditStatesResponse = json.RawMessage
type UserTokenResponse = json.RawMessage
type CreditsResponse = json.RawMessage

type AIPPTUserUsageCountResponse struct {
	Counts int `json:"counts"`
}

type AIPPTUsageStatItem struct {
	UserID     int    `json:"user_id"`
	UserNo     string `json:"user_no,omitempty"`
	UserName   string `json:"user_name,omitempty"`
	Department string `json:"department,omitempty"`
	Role       string `json:"role,omitempty"`
	UsageCount int    `json:"usage_count,omitempty"`
}

type AIPPTUsageStatsResponse struct {
	Items []*AIPPTUsageStatItem `json:"items"`
	model.Pagination
}

type UserCreditStatesResponse struct {
	Items []*UserCreditState `json:"items"`
	model.Pagination
}

type CourseCreditStatesResponse struct {
	Items []*CourseCreditState `json:"items"`
	model.Pagination
}

type UserCreditUsageStatsResponse struct {
	Items []*UserCreditUsageStat `json:"items"`
	model.Pagination
}

type CourseCreditUsageStatsResponse struct {
	Items []*CourseCreditUsageStat `json:"items"`
	model.Pagination
}

type CreditAuditsResponse struct {
	Items []*CreditAudit `json:"items"`
	model.Pagination
}
