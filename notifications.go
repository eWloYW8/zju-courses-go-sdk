package zjucourses

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// NotificationsService handles notification, todo, and alert API operations.
type NotificationsService struct {
	client *Client
}

// --- Response Types ---

type NotificationsResponse struct {
	Notifications []*model.Notification `json:"notifications"`
	UnreadCount   int                   `json:"unread_count,omitempty"`
	TotalCount    int                   `json:"total_count,omitempty"`
}

type TodosResponse struct {
	TodoList []*model.TodoItem `json:"todo_list"`
}

type AlertMessagesResponse struct {
	Data []interface{} `json:"data"`
}

type AnnouncementsResponse struct {
	Announcements []*model.Announcement `json:"announcements"`
}

type BulletinsResponse struct {
	Bulletins []*model.Bulletin `json:"bulletins"`
}

// --- Notifications ---

// ListNotifications returns notifications for a user.
func (s *NotificationsService) ListNotifications(ctx context.Context, userID int, opts *model.ListOptions) (*NotificationsResponse, error) {
	u := addListOptions(fmt.Sprintf("/ntf/users/%d/notifications", userID), opts)
	result := new(NotificationsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// --- Todos ---

// ListTodos returns the current user's to-do list.
func (s *NotificationsService) ListTodos(ctx context.Context) (*TodosResponse, error) {
	result := new(TodosResponse)
	_, err := s.client.get(ctx, "/api/todos?exclude_questionnaire=true", result)
	return result, err
}

// --- Alert Messages ---

// ListAlertMessages returns alert messages.
func (s *NotificationsService) ListAlertMessages(ctx context.Context) (*AlertMessagesResponse, error) {
	result := new(AlertMessagesResponse)
	_, err := s.client.get(ctx, "/api/alert/messages", result)
	return result, err
}

// MarkAlertMessagesRead marks alert messages as read.
func (s *NotificationsService) MarkAlertMessagesRead(ctx context.Context) error {
	_, err := s.client.post(ctx, "/api/alert/messages/read", nil, nil)
	return err
}

// --- Announcements ---

// ListAnnouncements returns announcements.
func (s *NotificationsService) ListAnnouncements(ctx context.Context) (*AnnouncementsResponse, error) {
	result := new(AnnouncementsResponse)
	_, err := s.client.get(ctx, "/api/announcement", result)
	return result, err
}

// --- Bulletins ---

// ListLatestBulletins returns the latest bulletins.
func (s *NotificationsService) ListLatestBulletins(ctx context.Context) (*BulletinsResponse, error) {
	result := new(BulletinsResponse)
	_, err := s.client.get(ctx, "/api/bulletins/latest", result)
	return result, err
}

// GetBulletin returns a specific bulletin.
func (s *NotificationsService) GetBulletin(ctx context.Context, bulletinID int) (*model.Bulletin, error) {
	u := fmt.Sprintf("/api/bulletins/%d", bulletinID)
	result := new(model.Bulletin)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// CreateBulletin creates a new bulletin (instructor).
func (s *NotificationsService) CreateBulletin(ctx context.Context, courseID int, body interface{}) (*model.Bulletin, error) {
	u := fmt.Sprintf("/api/course/bulletins/%d", courseID)
	result := new(model.Bulletin)
	_, err := s.client.post(ctx, u, body, result)
	return result, err
}

// --- Org Bulletins ---

// ListOrgBulletins returns organization bulletins.
func (s *NotificationsService) ListOrgBulletins(ctx context.Context, opts *model.ListOptions) (*BulletinsResponse, error) {
	u := addListOptions("/api/org-bulletin/bulletins", opts)
	result := new(BulletinsResponse)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListLatestOrgBulletins returns the latest organization bulletins.
func (s *NotificationsService) ListLatestOrgBulletins(ctx context.Context) (*BulletinsResponse, error) {
	result := new(BulletinsResponse)
	_, err := s.client.get(ctx, "/api/org-bulletin/bulletins/latest", result)
	return result, err
}

// GetOrgBulletin returns a specific org bulletin.
func (s *NotificationsService) GetOrgBulletin(ctx context.Context, bulletinID int) (*model.Bulletin, error) {
	u := fmt.Sprintf("/api/org-bulletin/bulletins/%d", bulletinID)
	result := new(model.Bulletin)
	_, err := s.client.get(ctx, u, result)
	return result, err
}

// ListOrgBulletinClassifications returns org bulletin classifications.
func (s *NotificationsService) ListOrgBulletinClassifications(ctx context.Context) (json.RawMessage, error) {
	var result json.RawMessage
	_, err := s.client.get(ctx, "/api/org-bulletin/classifications", &result)
	return result, err
}

// --- Latest Activities ---

// LatestActivitiesResponse represents the latest activities response.
type LatestActivitiesResponse struct {
	Activities []*model.Activity `json:"activities,omitempty"`
}

// ListLatestActivities returns the latest activities (material, web_link, slide, online_video, page).
func (s *NotificationsService) ListLatestActivities(ctx context.Context) (*LatestActivitiesResponse, error) {
	result := new(LatestActivitiesResponse)
	_, err := s.client.get(ctx, "/api/latest-activities?no-intercept=true&types=material,web_link,slide,online_video,page", result)
	return result, err
}
