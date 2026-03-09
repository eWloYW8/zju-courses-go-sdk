package others

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestProjectHelpersUseFrontendEndpointsAndModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/projects":
			conditions, err := url.QueryUnescape(r.URL.Query().Get("conditions"))
			if err != nil {
				t.Fatalf("unescape project conditions: %v", err)
			}
			var decoded map[string]any
			if err := json.Unmarshal([]byte(conditions), &decoded); err != nil {
				t.Fatalf("decode project conditions: %v", err)
			}
			if decoded["keyword"] != "proj" || decoded["applyStatus"] != "waiting" || decoded["hasKnowledgeNode"] != true {
				t.Fatalf("unexpected project conditions: %#v", decoded)
			}
			_, _ = w.Write([]byte(`{
				"items":[{"id":9,"name":"AI Project","classroom_schedule":"shared desc","knowledge_node_count":5,"enrolled_project":true,"instructors":[{"id":2,"name":"Teacher","user_no":"001"}],"audit":{"id":3,"status":"waiting","created_at":"2026-03-01","user":{"id":7,"name":"Alice","user_no":"2025001","department":{"id":4,"name":"CS"}}}}],
				"page":1,
				"page_size":10,
				"pages":1,
				"total":1
			}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/project":
			var body CreateProjectRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode create project body: %v", err)
			}
			if body.Name != "New Project" {
				t.Fatalf("unexpected create project body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"id":12,"name":"New Project"}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/project/12":
			var body UpdateProjectRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode update project body: %v", err)
			}
			if body.Name != "Renamed" || body.Description != "updated desc" {
				t.Fatalf("unexpected update project body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"id":12,"name":"Renamed","classroom_schedule":"updated desc"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/projects/12/apply":
			conditions, err := url.QueryUnescape(r.URL.Query().Get("conditions"))
			if err != nil {
				t.Fatalf("unescape apply conditions: %v", err)
			}
			if conditions != `{"keyword":"alice"}` {
				t.Fatalf("unexpected apply conditions: %s", conditions)
			}
			_, _ = w.Write([]byte(`{
				"items":[{"id":5,"status":"waiting","created_at":"2026-03-02","user":{"id":7,"name":"Alice","user_no":"2025001","department":{"id":4,"name":"CS"}}}],
				"page":1,
				"page_size":10,
				"pages":1,
				"total":1
			}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/projects/12/audit/5":
			var body AuditProjectApplicationRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode audit body: %v", err)
			}
			if body.Status != "accepted" {
				t.Fatalf("unexpected audit body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"ok":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/project/12/share-resource":
			conditions, err := url.QueryUnescape(r.URL.Query().Get("conditions"))
			if err != nil {
				t.Fatalf("unescape share-resource conditions: %v", err)
			}
			var decoded map[string]any
			if err := json.Unmarshal([]byte(conditions), &decoded); err != nil {
				t.Fatalf("decode share-resource conditions: %v", err)
			}
			if decoded["keyword"] != "ppt" || decoded["ref_parent_id"] != float64(6) {
				t.Fatalf("unexpected share-resource conditions: %#v", decoded)
			}
			_, _ = w.Write([]byte(`[{"id":7,"name":"讲义","created_at":"2026-03-01","created_by_id":5,"ref_parent_id":6,"upload":{"id":88,"name":"slides.pdf","type":"file","allow_download":true},"allow_download":false,"knowledge_count":2,"knowledge_nodes":[{"id":3,"name":"极限"}]}]`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/project/12/share-resource":
			var body ProjectSharedResourceRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode create shared resource body: %v", err)
			}
			if body.ResourceNewName != "slides" || body.AllowDownload != "false" || len(body.Uploads) != 2 {
				t.Fatalf("unexpected create shared resource body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/project/12/share-resource":
			var body ProjectSharedResourceRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode update shared resource body: %v", err)
			}
			if body.ReferenceID == nil || *body.ReferenceID != 7 || body.AllowDownload != "true" {
				t.Fatalf("unexpected update shared resource body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodDelete && r.URL.Path == "/api/project/12/share-resource":
			if got := r.URL.Query().Get("reference_id"); got != "7" {
				t.Fatalf("unexpected reference_id: %q", got)
			}
			if got := r.URL.Query().Get("upload_id"); got != "88" {
				t.Fatalf("unexpected upload_id: %q", got)
			}
			nodeIDs, err := url.QueryUnescape(r.URL.Query().Get("node_ids"))
			if err != nil {
				t.Fatalf("unescape node_ids: %v", err)
			}
			if nodeIDs != `[3,4]` {
				t.Fatalf("unexpected node_ids: %s", nodeIDs)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	projects, err := service.ListProjectsWithParams(ctx, ListProjectsParams{
		Page:     1,
		PageSize: 10,
		Conditions: map[string]any{
			"keyword":          "proj",
			"applyStatus":      "waiting",
			"hasKnowledgeNode": true,
		},
	})
	if err != nil {
		t.Fatalf("ListProjectsWithParams returned error: %v", err)
	}
	if len(projects.Items) != 1 || projects.Items[0].Audit == nil || projects.Items[0].Audit.User == nil {
		t.Fatalf("unexpected projects response: %#v", projects)
	}
	if projects.Items[0].Description == nil || *projects.Items[0].Description != "shared desc" {
		t.Fatalf("project description did not decode: %#v", projects.Items[0])
	}

	created, err := service.CreateProjectTyped(ctx, &CreateProjectRequest{Name: "New Project"})
	if err != nil || created.ID != 12 {
		t.Fatalf("CreateProjectTyped returned %#v, err=%v", created, err)
	}

	updated, err := service.UpdateProjectTyped(ctx, 12, &UpdateProjectRequest{Name: "Renamed", Description: "updated desc"})
	if err != nil || updated.Description == nil || *updated.Description != "updated desc" {
		t.Fatalf("UpdateProjectTyped returned %#v, err=%v", updated, err)
	}

	applications, err := service.ListProjectApplications(ctx, 12, ListProjectsParams{Page: 1, PageSize: 10, Conditions: map[string]any{"keyword": "alice"}})
	if err != nil || len(applications.Items) != 1 || applications.Items[0].User == nil || applications.Items[0].User.UserNo != "2025001" {
		t.Fatalf("unexpected project applications: %#v, err=%v", applications, err)
	}

	if _, err := service.AuditProjectApplication(ctx, 12, 5, "accepted"); err != nil {
		t.Fatalf("AuditProjectApplication returned error: %v", err)
	}

	refParentID := 6
	resources, err := service.ListProjectSharedResources(ctx, 12, ProjectSharedResourceConditions{Keyword: "ppt", RefParentID: &refParentID})
	if err != nil || len(resources) != 1 || resources[0].Upload == nil || resources[0].Upload.ID != 88 {
		t.Fatalf("unexpected project shared resources: %#v, err=%v", resources, err)
	}

	if err := service.CreateProjectSharedResources(ctx, 12, &ProjectSharedResourceRequest{
		ResourceNewName: "slides",
		AllowDownload:   "false",
		Uploads:         []int{88, 89},
		NodeIDs:         []any{3, 4},
	}); err != nil {
		t.Fatalf("CreateProjectSharedResources returned error: %v", err)
	}

	referenceID := 7
	if err := service.UpdateProjectSharedResources(ctx, 12, &ProjectSharedResourceRequest{
		ResourceNewName: "slides-v2",
		AllowDownload:   "true",
		Uploads:         []int{88},
		NodeIDs:         []any{"3", "4"},
		ReferenceID:     &referenceID,
	}); err != nil {
		t.Fatalf("UpdateProjectSharedResources returned error: %v", err)
	}

	if err := service.DeleteProjectSharedResource(ctx, 12, &DeleteProjectSharedResourceRequest{
		ReferenceID: 7,
		UploadID:    88,
		NodeIDs:     []int{3, 4},
	}); err != nil {
		t.Fatalf("DeleteProjectSharedResource returned error: %v", err)
	}
}
