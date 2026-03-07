package activities

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

func TestDecodeActivityFixture(t *testing.T) {
	data := readFixture(t, "api", "activities", "1087133.html")

	var activity Activity
	if err := json.Unmarshal(data, &activity); err != nil {
		t.Fatalf("unmarshal activity: %v", err)
	}
	if activity.ID != 1087133 {
		t.Fatalf("unexpected activity id: %d", activity.ID)
	}
	if activity.Type == "" {
		t.Fatal("expected activity type")
	}
}

func TestDecodeCommentsFixture(t *testing.T) {
	data := readFixture(t, "api", "activities", "1087140", "comments.html")

	var resp CommentsResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("unmarshal comments: %v", err)
	}
	if resp.PageSize != 10 {
		t.Fatalf("unexpected page size: %d", resp.PageSize)
	}
	if resp.Comments == nil {
		t.Fatal("expected comments slice")
	}
}

func TestDecodeUploadReferencesFixture(t *testing.T) {
	data := readFixture(t, "api", "activities", "1087140", "upload_references.html")

	var resp UploadReferencesResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("unmarshal upload references: %v", err)
	}
	if resp.References == nil {
		t.Fatal("expected references slice")
	}
}
