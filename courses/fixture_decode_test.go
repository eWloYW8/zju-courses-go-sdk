package courses

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func readFixture(t *testing.T, parts ...string) []byte {
	t.Helper()
	pathParts := append([]string{"..", "courses.zju.edu.cn"}, parts...)
	data, err := os.ReadFile(filepath.Join(pathParts...))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	return data
}

func TestDecodeCourseFixtures(t *testing.T) {
	tests := []struct {
		name string
		path []string
		target any
		assert func(t *testing.T, target any)
	}{
		{
			name: "course detail",
			path: []string{"api", "courses", "93119.html"},
			target: &Course{},
			assert: func(t *testing.T, target any) {
				course := target.(*Course)
				if course.ID != 93119 {
					t.Fatalf("unexpected course id: %d", course.ID)
				}
			},
		},
		{
			name: "modules",
			path: []string{"api", "courses", "93119", "modules.html"},
			target: &ModulesResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*ModulesResponse)
				if len(resp.Modules) == 0 {
					t.Fatal("expected modules")
				}
			},
		},
		{
			name: "enrollments",
			path: []string{"api", "course", "93119", "enrollments.html"},
			target: &EnrollmentsResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*EnrollmentsResponse)
				if len(resp.Enrollments) == 0 {
					t.Fatal("expected enrollments")
				}
				if len(resp.Enrollments[0].Roles) == 0 {
					t.Fatal("expected enrollment roles")
				}
			},
		},
		{
			name: "activity reads",
			path: []string{"api", "course", "93119", "activity-reads-for-user.html"},
			target: &ActivityReadsForUserResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*ActivityReadsForUserResponse)
				if len(resp.ActivityReads) == 0 {
					t.Fatal("expected activity reads")
				}
			},
		},
		{
			name: "homework scores",
			path: []string{"api", "course", "93119", "homework-scores.html"},
			target: &HomeworkScoresResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*HomeworkScoresResponse)
				if len(resp.HomeworkActivities) == 0 {
					t.Fatal("expected homework activities")
				}
			},
		},
		{
			name: "submission status",
			path: []string{"api", "course", "93119", "homework", "submission-status.html"},
			target: &HomeworkSubmissionStatusResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*HomeworkSubmissionStatusResponse)
				if resp.CourseID != 93119 {
					t.Fatalf("unexpected course id: %d", resp.CourseID)
				}
			},
		},
		{
			name: "nav setting",
			path: []string{"api", "courses", "93119", "nav-setting.html"},
			target: &NavSettingResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*NavSettingResponse)
				if len(resp.NavSetting) == 0 {
					t.Fatal("expected nav settings")
				}
			},
		},
		{
			name: "outline",
			path: []string{"api", "courses", "93119", "outline.html"},
			target: &OutlineResponse{},
			assert: func(t *testing.T, target any) {
				outline := target.(*OutlineResponse)
				if outline.ID == 0 {
					t.Fatal("expected outline id")
				}
			},
		},
		{
			name: "completeness",
			path: []string{"api", "course", "93119", "my-completeness.html"},
			target: &CompletenessResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*CompletenessResponse)
				if resp.CompletedResult == nil {
					t.Fatal("expected completed result")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := readFixture(t, tt.path...)
			if err := json.Unmarshal(data, tt.target); err != nil {
				t.Fatalf("unmarshal %s: %v", tt.name, err)
			}
			tt.assert(t, tt.target)
		})
	}
}
