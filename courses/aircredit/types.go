package aircredit

import "github.com/eWloYW8/zju-courses-go-sdk/courses/model"

type AIAbilityResponse struct {
	HasAnyCourseAIAbility bool `json:"has_any_course_ai_ability"`
}

type OrgCreditStateInfo struct {
	OrgID                int `json:"org_id,omitempty"`
	CreditAssigned       int `json:"credit_assigned,omitempty"`
	CreditUsed           int `json:"credit_used,omitempty"`
	UserCreditUsed       int `json:"user_credit_used,omitempty"`
	UserCreditAssigned   int `json:"user_credit_assigned,omitempty"`
	CourseCreditUsed     int `json:"course_credit_used,omitempty"`
	CourseCreditAssigned int `json:"course_credit_assigned,omitempty"`
}

type UserCreditState struct {
	UserID            int                `json:"user_id,omitempty"`
	UserNo            string             `json:"user_no,omitempty"`
	UserName          string             `json:"user_name,omitempty"`
	Department        string             `json:"department,omitempty"`
	Role              string             `json:"role,omitempty"`
	CreditUsedPercent float64            `json:"credit_used_percent,omitempty"`
	IsLowAirCredit    bool               `json:"is_low_air_credit,omitempty"`
	CreditState       *model.CreditState `json:"credit_state,omitempty"`
}

type CourseCreditState struct {
	Department        string             `json:"department,omitempty"`
	Instructors       string             `json:"instructors,omitempty"`
	Name              string             `json:"name,omitempty"`
	Semester          string             `json:"semester,omitempty"`
	AcademicYear      string             `json:"academic_year,omitempty"`
	CourseCode        string             `json:"course_code,omitempty"`
	CourseType        int                `json:"course_type,omitempty"`
	CreditUsedPercent float64            `json:"credit_used_percent,omitempty"`
	CourseID          int                `json:"course_id,omitempty"`
	CreditState       *model.CreditState `json:"credit_state,omitempty"`
}

type UserCreditUsageStat struct {
	User             *model.User    `json:"user,omitempty"`
	UserRole         string         `json:"user_role,omitempty"`
	CreditAssigned   int            `json:"credit_assigned,omitempty"`
	CreditUsed       int            `json:"credit_used,omitempty"`
	ModuleCreditUsed map[string]int `json:"module_credit_used,omitempty"`
	UsageCount       int            `json:"usage_count,omitempty"`
}

type CourseCreditUsageStat struct {
	Course                  *model.Course `json:"course,omitempty"`
	Instructors             int           `json:"instructors,omitempty"`
	CreditUsed              int           `json:"credit_used,omitempty"`
	UsageCount              int           `json:"usage_count,omitempty"`
	StudentsCount           int           `json:"students_count,omitempty"`
	UseAirChatStudentsCount int           `json:"use_air_chat_students_count,omitempty"`
}

type CreditAudit struct {
	ID              int           `json:"id,omitempty"`
	User            *model.User   `json:"user,omitempty"`
	Course          *model.Course `json:"course,omitempty"`
	AppliedCredits  int           `json:"applied_credits,omitempty"`
	ApprovedCredits int           `json:"approved_credits,omitempty"`
	Reason          string        `json:"reason,omitempty"`
	Status          string        `json:"status,omitempty"`
	Remark          string        `json:"remark,omitempty"`
	CreatedAt       string        `json:"created_at,omitempty"`
	UpdatedAt       string        `json:"updated_at,omitempty"`
	Auditor         *model.User   `json:"auditor,omitempty"`
	Read            bool          `json:"read,omitempty"`
	TargetType      string        `json:"target_type,omitempty"`
}

type OrgAIConfig struct {
	AirCreditMode            string `json:"air_credit_mode,omitempty"`
	AirCourseGuidesEnable    bool   `json:"air_course_guides_enable,omitempty"`
	AirKBAllowedVideoType    bool   `json:"air_kb_allowed_video_type,omitempty"`
	AirKBMaxSizeOfUploadFile int    `json:"air_kb_max_size_of_upload_file,omitempty"`
	AirMaxLengthOfText       int    `json:"air_max_length_of_text,omitempty"`
	AirMaxNumOfQuizzes       int    `json:"air_max_num_of_quizzes,omitempty"`
}
