package notifications

import (
	"context"
	"fmt"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

// Service handles notification, todo, and alert API operations.

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *sdk.Client
}

// --- Notifications ---

// ListNotifications returns notifications for a user.
func (s *Service) ListNotifications(ctx context.Context, userID int, opts *model.ListOptions) (*NotificationsResponse, error) {
	u := addListOptions(fmt.Sprintf("/ntf/users/%d/notifications", userID), opts)
	result := new(NotificationsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// --- Todos ---

// ListTodos returns the current user's to-do list.
func (s *Service) ListTodos(ctx context.Context) (*TodosResponse, error) {
	result := new(TodosResponse)
	_, err := s.client.Get(ctx, "/api/todos?exclude_questionnaire=true", result)
	return result, err
}

// --- Alert Messages ---

// ListAlertMessages returns alert messages.
func (s *Service) ListAlertMessages(ctx context.Context) (*AlertMessagesResponse, error) {
	result := new(AlertMessagesResponse)
	_, err := s.client.Get(ctx, "/api/alert/messages", result)
	return result, err
}

// MarkAlertMessagesRead marks alert messages as read.
func (s *Service) MarkAlertMessagesRead(ctx context.Context) error {
	_, err := s.client.Post(ctx, "/api/alert/messages/read", nil, nil)
	return err
}

// --- Announcements ---

// ListAnnouncements returns announcements.
func (s *Service) ListAnnouncements(ctx context.Context) (*AnnouncementsResponse, error) {
	result := new(AnnouncementsResponse)
	_, err := s.client.Get(ctx, "/api/announcement", result)
	return result, err
}

// --- Bulletins ---

// ListLatestBulletins returns the latest bulletins.
func (s *Service) ListLatestBulletins(ctx context.Context) (*BulletinsResponse, error) {
	result := new(BulletinsResponse)
	_, err := s.client.Get(ctx, "/api/bulletins/latest", result)
	return result, err
}

// GetBulletin returns a specific bulletin.
func (s *Service) GetBulletin(ctx context.Context, bulletinID int) (*Bulletin, error) {
	u := fmt.Sprintf("/api/bulletins/%d", bulletinID)
	result := new(Bulletin)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// MarkBulletinRead marks a bulletin as read.
func (s *Service) MarkBulletinRead(ctx context.Context, bulletinID int, orgID int) error {
	u := fmt.Sprintf("/api/bulletins/%d/read?org_id=%d", bulletinID, orgID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// CreateBulletin creates a new bulletin (instructor).
func (s *Service) CreateBulletin(ctx context.Context, courseID int, body CreateBulletinRequest) (*Bulletin, error) {
	u := fmt.Sprintf("/api/course/bulletins/%d", courseID)
	result := new(Bulletin)
	_, err := s.client.Post(ctx, u, body, result)
	return result, err
}

// --- Org Bulletins ---

// ListOrgBulletins returns organization bulletins.
func (s *Service) ListOrgBulletins(ctx context.Context, opts *model.ListOptions) (*BulletinsResponse, error) {
	u := addListOptions("/api/org-bulletin/bulletins", opts)
	result := new(BulletinsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListOrgBulletinsWithConditions returns organization bulletins with explicit conditions filters.
func (s *Service) ListOrgBulletinsWithConditions(ctx context.Context, opts *model.ListOptions, bulletinID *int, conditions string, fields string) (*BulletinsResponse, error) {
	params := map[string]string{}
	if bulletinID != nil {
		params["bulletin_id"] = fmt.Sprintf("%d", *bulletinID)
	}
	if conditions != "" {
		params["conditions"] = conditions
	}
	if fields != "" {
		params["fields"] = fields
	}
	u := addListOptions("/api/org-bulletin/bulletins", opts)
	u = addQueryParams(u, params)
	result := new(BulletinsResponse)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListUnreadOrgBulletins returns unread org bulletin ids or summaries.
func (s *Service) ListUnreadOrgBulletins(ctx context.Context) (*BulletinsResponse, error) {
	result := new(BulletinsResponse)
	u := addQueryParams("/api/org-bulletin/bulletins", map[string]string{"fields": "id", "conditions": "{\"unread\":true}"})
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// ListLatestOrgBulletins returns the latest organization bulletins.
func (s *Service) ListLatestOrgBulletins(ctx context.Context) (*BulletinsResponse, error) {
	result := new(BulletinsResponse)
	_, err := s.client.Get(ctx, "/api/org-bulletin/bulletins/latest", result)
	return result, err
}

// GetOrgBulletin returns a specific org bulletin.
func (s *Service) GetOrgBulletin(ctx context.Context, bulletinID int) (*Bulletin, error) {
	u := fmt.Sprintf("/api/org-bulletin/bulletins/%d", bulletinID)
	result := new(Bulletin)
	_, err := s.client.Get(ctx, u, result)
	return result, err
}

// CreateOrgBulletin creates an organization bulletin.
func (s *Service) CreateOrgBulletin(ctx context.Context, body OrgBulletinRequest) (*Bulletin, error) {
	result := new(Bulletin)
	_, err := s.client.Post(ctx, "/api/org-bulletin/bulletins", body, result)
	return result, err
}

// UpdateOrgBulletin updates an organization bulletin.
func (s *Service) UpdateOrgBulletin(ctx context.Context, bulletinID int, body OrgBulletinRequest) (*Bulletin, error) {
	u := fmt.Sprintf("/api/org-bulletin/bulletins/%d", bulletinID)
	result := new(Bulletin)
	_, err := s.client.Put(ctx, u, body, result)
	return result, err
}

// DeleteOrgBulletin deletes an organization bulletin.
func (s *Service) DeleteOrgBulletin(ctx context.Context, bulletinID int) error {
	u := fmt.Sprintf("/api/org-bulletin/bulletins/%d", bulletinID)
	_, err := s.client.Delete(ctx, u, nil)
	return err
}

// MarkOrgBulletinRead marks an organization bulletin as read.
func (s *Service) MarkOrgBulletinRead(ctx context.Context, bulletinID int) error {
	u := fmt.Sprintf("/api/org-bulletin/bulletins/%d/read", bulletinID)
	_, err := s.client.Post(ctx, u, nil, nil)
	return err
}

// ListOrgBulletinClassifications returns org bulletin classifications.
func (s *Service) ListOrgBulletinClassifications(ctx context.Context) ([]*OrgBulletinClassification, error) {
	var result []*OrgBulletinClassification
	_, err := s.client.Get(ctx, "/api/org-bulletin/classifications", &result)
	return result, err
}

// --- Latest Activities ---

// ListLatestActivities returns the latest activities (material, web_link, slide, online_video, page).
func (s *Service) ListLatestActivities(ctx context.Context) (*LatestActivitiesResponse, error) {
	result := new(LatestActivitiesResponse)
	_, err := s.client.Get(ctx, "/api/latest-activities?no-intercept=true&types=material,web_link,slide,online_video,page", result)
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
