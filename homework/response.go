package homework

import "github.com/eWloYW8/zju-courses-go-sdk/activities"

type SubmissionListResponse struct {
	List []*Submission `json:"list"`
}

type HomeworkActivityResponse = activities.Activity
