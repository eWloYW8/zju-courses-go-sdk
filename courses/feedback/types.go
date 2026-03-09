package feedback

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type FeedbackActivity = model.Activity

type Feedback struct {
	ID        int            `json:"id"`
	Content   string         `json:"content,omitempty"`
	CreatedAt *string        `json:"created_at,omitempty"`
	UpdatedAt *string        `json:"updated_at,omitempty"`
	CreatedBy *model.User    `json:"created_by,omitempty"`
	User      *model.User    `json:"user,omitempty"`
	Data      map[string]any `json:"data,omitempty"`
}
