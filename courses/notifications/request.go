package notifications

type CreateBulletinRequest map[string]any

type OrgBulletinRequest map[string]any

type CourseBulletinOptions struct {
	OrgID        int
	IsManagement bool
}
