package model

// User represents a complete user in the system.
type User struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	Email             *string            `json:"email,omitempty"`
	Nickname          *string            `json:"nickname,omitempty"`
	UserNo            string             `json:"user_no,omitempty"`
	AvatarSmallURL    string             `json:"avatar_small_url,omitempty"`
	AvatarBigURL      string             `json:"avatar_big_url,omitempty"`
	PortfolioURL      string             `json:"portfolio_url,omitempty"`
	Comment           *string            `json:"comment,omitempty"`
	Grade             *Grade             `json:"grade,omitempty"`
	Klass             *Class             `json:"klass,omitempty"`
	MobilePhone       string             `json:"mobile_phone,omitempty"`
	Language          string             `json:"language,omitempty"`
	CreatedAt         string             `json:"created_at,omitempty"`
	UpdatedAt         string             `json:"updated_at,omitempty"`
	ImportedFrom      string             `json:"imported_from,omitempty"`
	IsImportedData    bool               `json:"is_imported_data,omitempty"`
	ProgramID         int                `json:"program_id,omitempty"`
	Remarks           *string            `json:"remarks,omitempty"`
	StorageAssigned   int                `json:"storage_assigned,omitempty"`
	StorageUsed       int                `json:"storage_used,omitempty"`
	WebexAuth         bool               `json:"webex_auth,omitempty"`
	Department        *Department        `json:"department,omitempty"`
	LearningCenter    *LearningCenter    `json:"learning_center,omitempty"`
	Org               *OrgDetail         `json:"org,omitempty"`
	Program           *Program           `json:"program,omitempty"`
	UserAttributes    *UserAttributes    `json:"user_attributes,omitempty"`
	UserPersonas      *UserPersonas      `json:"user_personas,omitempty"`
	UserAddresses     []UserAddress      `json:"user_addresses,omitempty"`
	UserAuthExternals []UserAuthExternal `json:"user_auth_externals,omitempty"`
	Roles             []*Role            `json:"roles,omitempty"`
	Audit             *Audit             `json:"audit,omitempty"`
}

// UserAttributes represents extended user attribute fields.
type UserAttributes struct {
	ID             int     `json:"id"`
	AddressName    *string `json:"address_name"`
	AdmissionDate  *string `json:"admission_date"`
	Birthday       *string `json:"birthday"`
	Education      *string `json:"education"`
	FirstJobDate   *string `json:"first_job_date"`
	GraduatedFrom  *string `json:"graduated_from"`
	IdentityNumber *string `json:"identity_number"`
	IncomeName     *string `json:"income_name"`
	JobName        *string `json:"job_name"`
	JobType        *string `json:"job_type"`
	Nation         *string `json:"nation"`
	NativePlace    *string `json:"native_place"`
	OccupationType *string `json:"occupation_type"`
	PortfolioURL   *string `json:"portfolio_url"`
	Tag            *string `json:"tag"`
}

// UserPersonas represents user persona wrapper.
type UserPersonas struct {
	Data *UserPersonaData `json:"data,omitempty"`
}

// UserPersonaData represents the data inside user personas.
type UserPersonaData struct {
	AreaCodeForCompanyPhone *string `json:"area_code_for_company_phone"`
	AreaCodeForFaxNumber    *string `json:"area_code_for_fax_number"`
	CompanyName             *string `json:"company_name"`
	CompanyPhone            *string `json:"company_phone"`
	CountryCode             *string `json:"country_code"`
	Desc                    *string `json:"desc"`
	Direction               *string `json:"direction"`
	FaxNumber               *string `json:"fax_number"`
	IsTTBA                  *bool   `json:"is_ttba"`
	LineID                  *string `json:"line_id"`
	Society                 *string `json:"society"`
	Title                   *string `json:"title"`
	UploadID                *int    `json:"upload_id"`
	UploadURL               *string `json:"upload_url"`
}

// UserAddress represents a user's address entry.
type UserAddress struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Address  string `json:"address,omitempty"`
	PostCode string `json:"post_code,omitempty"`
	Type     string `json:"type,omitempty"`
}

// UserAuthExternal represents an external authentication provider link.
type UserAuthExternal struct {
	ID       int    `json:"id"`
	Type     string `json:"type,omitempty"`
	Provider string `json:"provider,omitempty"`
	UID      string `json:"uid,omitempty"`
}

// Department represents an academic department.
type Department struct {
	ID               int           `json:"id"`
	Code             *string       `json:"code,omitempty"`
	Name             string        `json:"name"`
	ShortName        *string       `json:"short_name,omitempty"`
	Stopped          *bool         `json:"stopped,omitempty"`
	ParentID         int           `json:"parent_id,omitempty"`
	Sort             int           `json:"sort,omitempty"`
	Cover            *string       `json:"cover,omitempty"`
	IsShowOnHomepage *bool         `json:"is_show_on_homepage,omitempty"`
	StorageAssigned  int           `json:"storage_assigned,omitempty"`
	StorageUsed      int           `json:"storage_used,omitempty"`
	CreatedAt        *string       `json:"created_at,omitempty"`
	UpdatedAt        *string       `json:"updated_at,omitempty"`
	CreatedUser      *User         `json:"created_user,omitempty"`
	UpdatedUser      *User         `json:"updated_user,omitempty"`
	Children         []*Department `json:"departments,omitempty"`
}

// LearningCenter represents a user's learning center.
type LearningCenter struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

// Org represents the organization in brief form.
type Org struct {
	IsEnterpriseOrOrganization bool `json:"is_enterprise_or_organization"`
	IsTransferArrears          bool `json:"is_transfer_arrears,omitempty"`
}

// OrgDetail represents full organization details.
type OrgDetail struct {
	ID                         int               `json:"id"`
	Code                       string            `json:"code,omitempty"`
	Name                       string            `json:"name,omitempty"`
	RawName                    string            `json:"raw_name,omitempty"`
	Domain                     string            `json:"domain,omitempty"`
	Flag                       int               `json:"flag,omitempty"`
	ParentID                   int               `json:"parent_id,omitempty"`
	Copyright                  *string           `json:"copyright,omitempty"`
	CurrentOrgPlanID           int               `json:"current_org_plan_id,omitempty"`
	EnableExternalDomain       bool              `json:"enable_external_domain,omitempty"`
	ExternalDomain             *string           `json:"external_domain,omitempty"`
	FilingNumber               *string           `json:"filing_number,omitempty"`
	IsEnterpriseOrOrganization bool              `json:"is_enterprise_or_organization,omitempty"`
	IsTransferArrears          bool              `json:"is_transfer_arrears,omitempty"`
	OrgNames                   map[string]string `json:"org_names,omitempty"`
	Logo                       *string           `json:"logo,omitempty"`
	SmallLogo                  *string           `json:"small_logo,omitempty"`
	Favicon                    *string           `json:"favicon,omitempty"`
	StorageAssigned            int64             `json:"storage_assigned,omitempty"`
	StorageUsed                int64             `json:"storage_used,omitempty"`
}

// Program represents an academic program.
type Program struct {
	ID           int         `json:"id"`
	Code         *string     `json:"code,omitempty"`
	Name         *string     `json:"name,omitempty"`
	EnglishName  *string     `json:"english_name,omitempty"`
	Description  *string     `json:"description,omitempty"`
	ExternalID   *string     `json:"external_id,omitempty"`
	Level        *string     `json:"level,omitempty"`
	Status       *string     `json:"status,omitempty"`
	DepartmentID int         `json:"department_id,omitempty"`
	Discipline   *Discipline `json:"discipline,omitempty"`
}

// Discipline represents an academic discipline.
type Discipline struct {
	ID   int     `json:"id"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Role represents a user role in the system.
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
}

// CreditState represents AI credit state information.
type CreditState struct {
	CreditRemaining int    `json:"credit_remaining"`
	Status          string `json:"status"`
}

// AcademicYear represents an academic year.
type AcademicYear struct {
	ID       int    `json:"id"`
	Code     string `json:"code,omitempty"`
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
	IsActive bool   `json:"is_active,omitempty"`
}

// Semester represents a semester within an academic year.
type Semester struct {
	ID             int    `json:"id"`
	AcademicYearID int    `json:"academic_year_id,omitempty"`
	Code           string `json:"code,omitempty"`
	Name           string `json:"name"`
	RealName       string `json:"real_name,omitempty"`
	Sort           int    `json:"sort"`
	IsActive       bool   `json:"is_active,omitempty"`
	StartDate      string `json:"start_date,omitempty"`
	EndDate        string `json:"end_date,omitempty"`
}

// Audit represents an audit record for approval workflows.
type Audit struct {
	ID        int     `json:"id"`
	Status    *string `json:"status,omitempty"`
	Comment   *string `json:"comment,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	AuditBy   *User   `json:"audit_by,omitempty"`
}

// Pagination represents common pagination metadata in list responses.
type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Pages    int `json:"pages"`
	Total    int `json:"total"`
	Start    int `json:"start"`
	End      int `json:"end"`
}

// ListOptions specifies pagination parameters for list API requests.
type ListOptions struct {
	Page     int `url:"page,omitempty"`
	PageSize int `url:"page_size,omitempty"`
}

// Grade represents a grade entry for a student in a course.
type Grade struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	CourseID int      `json:"course_id,omitempty"`
	Score    *float64 `json:"score,omitempty"`
	OrgID    int      `json:"org_id,omitempty"`
}

// Class represents an administrative class grouping of students.
type Class struct {
	ID           int    `json:"id"`
	Name         string `json:"name,omitempty"`
	Code         string `json:"code,omitempty"`
	OrgID        int    `json:"org_id,omitempty"`
	DepartmentID int    `json:"department_id,omitempty"`
	Grade        string `json:"grade,omitempty"`
}
