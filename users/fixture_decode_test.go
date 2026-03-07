package users

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

func TestDecodeCourseGraduateCheckResponse(t *testing.T) {
	data := []byte(`{"completed_course_id":[1,2],"uncompleted_course_id":[3,4]}`)
	var resp CourseGraduateCheckResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("unmarshal course graduate response: %v", err)
	}
	if len(resp.CompletedCourseIDs) != 2 || resp.CompletedCourseIDs[0] != 1 || resp.CompletedCourseIDs[1] != 2 {
		t.Fatal("expected completed course ids")
	}
	if len(resp.UncompletedCourseIDs) != 2 || resp.UncompletedCourseIDs[0] != 3 || resp.UncompletedCourseIDs[1] != 4 {
		t.Fatal("expected uncompleted course ids")
	}
}

func TestDecodeUserFixtures(t *testing.T) {
	tests := []struct {
		name   string
		path   []string
		target any
		assert func(t *testing.T, target any)
	}{
		{
			name:   "academic years",
			path:   []string{"api", "my-academic-years.html"},
			target: &AcademicYearsResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*AcademicYearsResponse)
				if len(resp.AcademicYears) == 0 {
					t.Fatal("expected academic years")
				}
			},
		},
		{
			name:   "semesters",
			path:   []string{"api", "my-semesters.html"},
			target: &SemestersResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*SemestersResponse)
				if len(resp.Semesters) == 0 {
					t.Fatal("expected semesters")
				}
			},
		},
		{
			name:   "classes",
			path:   []string{"api", "my-classes.html"},
			target: &ClassesResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*ClassesResponse)
				if resp.Classes == nil {
					t.Fatal("expected classes slice")
				}
			},
		},
		{
			name:   "grades",
			path:   []string{"api", "my-grades.html"},
			target: &GradesResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*GradesResponse)
				if resp.Grades == nil {
					t.Fatal("expected grades slice")
				}
			},
		},
		{
			name:   "departments",
			path:   []string{"api", "my-departments.html"},
			target: &DepartmentsResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*DepartmentsResponse)
				if len(resp.Departments) == 0 {
					t.Fatal("expected departments")
				}
			},
		},
		{
			name:   "user resources",
			path:   []string{"api", "user", "resources.html"},
			target: &UserResourcesResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*UserResourcesResponse)
				if len(resp.Uploads) == 0 {
					t.Fatal("expected uploads")
				}
				if resp.Uploads[7].Thumbnail == nil || resp.Uploads[7].Thumbnail.ID == 0 {
					t.Fatal("expected thumbnail object")
				}
			},
		},
		{
			name:   "academic learning resources wrapper",
			path:   []string{"api", "user", "resources.html"},
			target: &AcademicLearningResourcesResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*AcademicLearningResourcesResponse)
				if len(resp.Uploads) == 0 || resp.Pages == 0 {
					t.Fatal("expected paginated uploads response")
				}
			},
		},
		{
			name:   "third part resources wrapper",
			path:   []string{"api", "user", "resources.html"},
			target: &ThirdPartResourcesResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*ThirdPartResourcesResponse)
				if len(resp.Uploads) == 0 || resp.Pages == 0 {
					t.Fatal("expected paginated uploads response")
				}
			},
		},
		{
			name:   "other video resources wrapper",
			path:   []string{"api", "user", "resources.html"},
			target: &OtherVideoResourcesResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*OtherVideoResourcesResponse)
				if len(resp.Uploads) == 0 || resp.Pages == 0 {
					t.Fatal("expected paginated uploads response")
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
