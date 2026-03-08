package resources

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestListSharedResourcesWithParamsUsesFrontendQueryAndModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/shared-resources-no-repeated" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("page"); got != "2" {
			t.Fatalf("unexpected page: %q", got)
		}
		if got := r.URL.Query().Get("page_size"); got != "5" {
			t.Fatalf("unexpected page_size: %q", got)
		}
		conditions, err := url.QueryUnescape(r.URL.Query().Get("conditions"))
		if err != nil {
			t.Fatalf("unescape conditions: %v", err)
		}
		if conditions != `{"keyword":"AI"}` {
			t.Fatalf("unexpected conditions: %s", conditions)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"resources":[
				{
					"id":10,
					"name":"共享课件",
					"org_id":2,
					"parent_id":0,
					"cc_license_id":8,
					"referrer_id":88,
					"cc_license_name":"CC BY",
					"cc_license_link":"https://creativecommons.org/licenses/by/4.0/",
					"cc_license_code":"BY",
					"cc_license_description":"署名",
					"referrer_type":"lesson_resource",
					"resource_type":"file",
					"audit_status":"approved",
					"open_scope":"org",
					"allow_download":true,
					"allow_save":true,
					"is_folder":false,
					"reported":false,
					"selected":true,
					"_checked":true,
					"slide":{"id":7,"title":"Week 1","video_id":11,"demand_id":12,"template_id":13},
					"course_package":{"id":6,"name":"pkg","is_folder":false,"parent_id":1},
					"lesson_resource":{"id":5,"name":"lesson.pdf","mimetype":"application/pdf","app_id":3},
					"percentage":42.5,
					"course_code":"(2025-2026-2)-MATH1147G-0000123-1"
				}
			],
			"page":2,
			"page_size":5,
			"pages":4,
			"total":16
		}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.ListSharedResourcesWithParams(context.Background(), ListSharedResourcesParams{
		Page:       2,
		PageSize:   5,
		Conditions: `{"keyword":"AI"}`,
	})
	if err != nil {
		t.Fatalf("ListSharedResourcesWithParams returned error: %v", err)
	}
	if len(result.Resources) != 1 {
		t.Fatalf("unexpected resources count: %#v", result.Resources)
	}
	resource := result.Resources[0]
	if resource.CcLicenseID != 8 || resource.ReferrerID != 88 || resource.CcLicenseLink == "" || resource.ReferrerType != "lesson_resource" {
		t.Fatalf("shared-resource metadata did not decode: %#v", resource)
	}
	if !resource.Selected || !resource.Checked {
		t.Fatalf("selection state did not decode: %#v", resource)
	}
	if resource.Slide == nil || resource.Slide.TemplateID != 13 {
		t.Fatalf("slide did not decode: %#v", resource.Slide)
	}
	if resource.CoursePackage == nil || resource.CoursePackage.ParentID != 1 {
		t.Fatalf("course package did not decode: %#v", resource.CoursePackage)
	}
	if resource.LessonResource == nil || resource.LessonResource.Mimetype != "application/pdf" {
		t.Fatalf("lesson resource did not decode: %#v", resource.LessonResource)
	}
	if resource.Percentage == nil || *resource.Percentage != 42.5 || resource.CourseCode == "" {
		t.Fatalf("percentage/course_code did not decode: %#v", resource)
	}
}
