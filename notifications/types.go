package notifications

import (
	"github.com/eWloYW8/zju-courses-go-sdk/activities"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

type Notification = model.Notification

type NotificationPayload = model.NotificationPayload

type TodoItem = model.TodoItem

type Bulletin = model.Bulletin

type Announcement = model.Announcement

type AlertMessage struct {
	ID   int            `json:"id,omitempty"`
	Data map[string]any `json:"data,omitempty"`
}

type OrgBulletinClassification struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

type LatestActivity = activities.Activity
