package homework

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

func TestDecodeHomeworkFixtures(t *testing.T) {
	tests := []struct {
		name string
		path []string
		target any
		assert func(t *testing.T, target any)
	}{
		{
			name: "homework score",
			path: []string{"api", "activities", "1087133", "students", "280566", "homework-score.html"},
			target: &HomeworkScore{},
			assert: func(t *testing.T, target any) {
				score := target.(*HomeworkScore)
				if score.ActivityID != 0 && score.ActivityID != 1087133 {
					t.Fatalf("unexpected activity id: %d", score.ActivityID)
				}
			},
		},
		{
			name: "submission list",
			path: []string{"api", "activities", "1087133", "students", "280566", "submission_list.html"},
			target: &SubmissionListResponse{},
			assert: func(t *testing.T, target any) {
				resp := target.(*SubmissionListResponse)
				if resp.List == nil {
					t.Fatal("expected submission list slice")
				}
			},
		},
		{
			name: "make-up record",
			path: []string{"api", "homework", "1087133", "students", "280566", "make-up-record.html"},
			target: &MakeUpRecordResponse{},
			assert: func(t *testing.T, target any) {
				if *target.(*MakeUpRecordResponse) == nil {
					t.Fatal("expected make-up record map")
				}
			},
		},
		{
			name: "resubmit record",
			path: []string{"api", "homework", "1087133", "students", "280566", "resubmit-record.html"},
			target: &ResubmitRecordResponse{},
			assert: func(t *testing.T, target any) {
				if *target.(*ResubmitRecordResponse) == nil {
					t.Fatal("expected resubmit record map")
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
