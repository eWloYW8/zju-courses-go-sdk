package homework

import "github.com/eWloYW8/zju-courses-go-sdk/courses/activities"

type SubmissionListResponse struct {
	List []*Submission `json:"list"`
}

type HomeworkActivityResponse = activities.Activity

type SubmissionRecordsResponse struct {
	Submissions []*Submission `json:"submissions,omitempty"`
}

type HomeworkScoresResponse struct {
	HomeworkScores []*HomeworkScore `json:"homework_scores,omitempty"`
}

type RecommendSubmissionsResponse struct {
	Submissions []*Submission `json:"submissions,omitempty"`
}

type MakeUpRecordsResponse struct {
	MakeUpRecords []MakeUpRecord `json:"make_up_records,omitempty"`
}

type ResubmitRecordsResponse struct {
	ResubmitRecords []ResubmitRecord `json:"resubmit_records,omitempty"`
}

type HomeworkLogsResponse struct {
	Logs []HomeworkLog `json:"logs,omitempty"`
}

type IntraScoreRulesResponse struct {
	IntraScores []InterScore `json:"intra_scores,omitempty"`
}

type RedoMapResponse struct {
	RedoMap map[string]int `json:"redo_map,omitempty"`
}

type DuplicateDetectRatesResponse struct {
	Items []DuplicateDetectRateItem `json:"items,omitempty"`
}
