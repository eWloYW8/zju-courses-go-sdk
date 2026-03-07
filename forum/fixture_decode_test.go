package forum

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

func TestDecodeForumCategoryFixture(t *testing.T) {
	data := readFixture(t, "api", "forum", "categories", "167810.html")

	var resp ForumCategoryResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("unmarshal category: %v", err)
	}
	if resp.ID != 167810 {
		t.Fatalf("unexpected category id: %d", resp.ID)
	}
	if resp.Result == nil || len(resp.Result.Topics) == 0 {
		t.Fatal("expected topics in category result")
	}
	if len(resp.Result.Topics[0].Enrollments) == 0 {
		t.Fatal("expected topic enrollments")
	}
}

func TestDecodeForumScoreFixture(t *testing.T) {
	data := readFixture(t, "api", "activities", "1087136", "students", "280566", "forum-score.html")

	var resp ForumScoreResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("unmarshal forum score: %v", err)
	}
	if resp.ForumScore == nil {
		t.Fatal("expected forum score")
	}
}
