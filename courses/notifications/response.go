package notifications

type NotificationsResponse struct {
	Notifications []*Notification `json:"notifications"`
	UnreadCount   int             `json:"unread_count,omitempty"`
	TotalCount    int             `json:"total_count,omitempty"`
}

type TodosResponse struct {
	TodoList []*TodoItem `json:"todo_list"`
}

type AlertMessagesResponse struct {
	Data []AlertMessage `json:"data"`
}

type AnnouncementsResponse struct {
	Announcements []*Announcement `json:"announcements"`
}

type BulletinsResponse struct {
	Bulletins []*Bulletin `json:"bulletins"`
}

type LatestActivitiesResponse struct {
	Activities []*LatestActivity `json:"activities,omitempty"`
}
