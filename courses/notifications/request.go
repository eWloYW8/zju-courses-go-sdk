package notifications

type CreateBulletinRequest map[string]any

type OrgBulletinRequest map[string]any

type ListNotificationsParams struct {
	Offset           int
	Limit            int
	Removed          string
	AdditionalFields string
}

type CourseBulletinOptions struct {
	OrgID        int
	IsManagement bool
}
