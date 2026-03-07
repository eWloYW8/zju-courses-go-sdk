package resources

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
	"github.com/eWloYW8/zju-courses-go-sdk/model"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

func newTestService(t *testing.T, fn roundTripFunc) *Service {
	t.Helper()
	client := sdk.NewClient(sdk.WithHTTPClient(&http.Client{Transport: fn}))
	return New(client)
}

func jsonResponse(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestListResourceGroupsUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/resource-groups" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"resource_groups":[]}`), nil
	})

	resp, err := svc.ListResourceGroups(context.Background(), &model.ListOptions{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("ListResourceGroups error: %v", err)
	}
	if resp.ResourceGroups == nil {
		t.Fatal("expected resource_groups slice")
	}
}

func TestListPagedResourceGroupFoldersUsesGroupEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/resource-groups/9/folders" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("conditions"); got != "{\"keyword\":\"abc\"}" {
			t.Fatalf("unexpected conditions: %s", got)
		}
		return jsonResponse(`{"folders":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.ListPagedResourceGroupFolders(context.Background(), 9, &model.ListOptions{Page: 1, PageSize: 10}, `{"keyword":"abc"}`)
	if err != nil {
		t.Fatalf("ListPagedResourceGroupFolders error: %v", err)
	}
	if resp.Folders == nil {
		t.Fatal("expected folders slice")
	}
}

func TestListPagedResourceGroupResourcesUsesGroupEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/resource-groups/9/resources" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"resources":[],"page":1,"pages":1,"total":0}`), nil
	})

	resp, err := svc.ListPagedResourceGroupResources(context.Background(), 9, &model.ListOptions{Page: 1, PageSize: 10}, "")
	if err != nil {
		t.Fatalf("ListPagedResourceGroupResources error: %v", err)
	}
	if resp.Resources == nil {
		t.Fatal("expected resources slice")
	}
}

func TestListSharedResourcesUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/shared-resources-no-repeated" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"resources":[]}`), nil
	})

	resp, err := svc.ListSharedResources(context.Background(), &model.ListOptions{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("ListSharedResources error: %v", err)
	}
	if resp.Resources == nil {
		t.Fatal("expected shared resources slice")
	}
}

func TestGetSharedResourceUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/shared-resources/3" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":3,"name":"res"}`), nil
	})

	resp, err := svc.GetSharedResource(context.Background(), 3)
	if err != nil {
		t.Fatalf("GetSharedResource error: %v", err)
	}
	if resp.ID != 3 {
		t.Fatalf("unexpected resource id: %d", resp.ID)
	}
}

func TestGetResourceFileUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/resource-file/7" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":7,"name":"file.pdf"}`), nil
	})

	resp, err := svc.GetResourceFile(context.Background(), 7)
	if err != nil {
		t.Fatalf("GetResourceFile error: %v", err)
	}
	if resp.ID != 7 {
		t.Fatalf("unexpected file id: %d", resp.ID)
	}
}

func TestListCoursePackagesForCourseUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courses/12/course-package" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		if got := req.URL.Query().Get("keyword"); got != "pkg" {
			t.Fatalf("unexpected keyword: %s", got)
		}
		return jsonResponse(`{"data":{"items":[],"page":1,"page_size":10,"pages":1,"total":0}}`), nil
	})

	resp, err := svc.ListCoursePackagesForCourse(context.Background(), 12, &model.ListOptions{Page: 1, PageSize: 10}, "pkg")
	if err != nil {
		t.Fatalf("ListCoursePackagesForCourse error: %v", err)
	}
	if resp.Items == nil {
		t.Fatal("expected course package items slice")
	}
}

func TestGetCoursePackageExportStatusUsesCourseEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courses/12/course-package/status" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"data":{"id":1,"status":"progressing"}}`), nil
	})

	resp, err := svc.GetCoursePackageExportStatus(context.Background(), 12)
	if err != nil {
		t.Fatalf("GetCoursePackageExportStatus error: %v", err)
	}
	if resp.ID != 1 {
		t.Fatalf("unexpected package status id: %d", resp.ID)
	}
}

func TestExportCoursePackageUsesExportEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courses/12/course-package/export" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.ExportCoursePackage(context.Background(), 12, map[string]any{"name": "pkg"}); err != nil {
		t.Fatalf("ExportCoursePackage error: %v", err)
	}
}

func TestListHotPublicCoursesUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/courses/public/hot" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"courses":[]}`), nil
	})

	resp, err := svc.ListHotPublicCourses(context.Background(), 8)
	if err != nil {
		t.Fatalf("ListHotPublicCourses error: %v", err)
	}
	if resp.Courses == nil {
		t.Fatal("expected public courses slice")
	}
}

func TestGetResourceGroupMembersUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/resource-groups/4/members" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"members":[]}`), nil
	})

	resp, err := svc.GetResourceGroupMembers(context.Background(), 4)
	if err != nil {
		t.Fatalf("GetResourceGroupMembers error: %v", err)
	}
	if resp.Members == nil {
		t.Fatal("expected members slice")
	}
}

func TestListPagedResourceGroupMembersUsesPaginationEndpoint(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/resource-groups/4/members" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"members":[],"page":1,"page_size":10,"pages":1,"total":0}`), nil
	})

	resp, err := svc.ListPagedResourceGroupMembers(context.Background(), 4, &model.ListOptions{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("ListPagedResourceGroupMembers error: %v", err)
	}
	if resp.Members == nil {
		t.Fatal("expected members slice")
	}
}

func TestDeleteResourceGroupMembersUsesDeleteBody(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodDelete {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/resource-groups/4/member" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{}`), nil
	})

	if err := svc.DeleteResourceGroupMembers(context.Background(), 4, map[string]any{"member_ids": []int{1, 2}}); err != nil {
		t.Fatalf("DeleteResourceGroupMembers error: %v", err)
	}
}
