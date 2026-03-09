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

func TestHomepageResourceHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/shared-resource/classifications":
			_, _ = w.Write([]byte(`{"classifications":[{"id":1,"name":"精选","parent_id":0}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/departments":
			if got := r.URL.Query().Get("fields"); got != "id,name,code,parent_id,stopped,short_name,is_show_on_homepage,cover" {
				t.Fatalf("unexpected department fields: %q", got)
			}
			_, _ = w.Write([]byte(`{"departments":[{"id":2,"name":"计算机学院","code":"CS","parent_id":0,"is_show_on_homepage":true}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/departments/show-on-homepage":
			_, _ = w.Write([]byte(`{"departments":[{"id":3,"name":"数学学院"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/shared-resources":
			if r.URL.Query().Get("no-intercept") == "true" {
				conditions, err := url.QueryUnescape(r.URL.Query().Get("conditions"))
				if err != nil {
					t.Fatalf("unescape homepage conditions: %v", err)
				}
				if conditions != `{"classification":9,"department":5,"order_by":"view_count","parent_id":0}` {
					t.Fatalf("unexpected homepage conditions: %s", conditions)
				}
				_, _ = w.Write([]byte(`{"resources":[{"id":8,"name":"首页资源"}]}`))
				return
			}
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected search page: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "6" {
				t.Fatalf("unexpected search page_size: %q", got)
			}
			if got := r.URL.Query().Get("fields"); got != sharedResourceSearchFields {
				t.Fatalf("unexpected search fields: %q", got)
			}
			_, _ = w.Write([]byte(`{"resources":[{"id":9,"name":"搜索资源"}],"page":2,"page_size":6,"pages":3,"total":15}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/shared-resources/from-me":
			if got := r.URL.Query().Get("page"); got != "1" {
				t.Fatalf("unexpected from-me page: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "4" {
				t.Fatalf("unexpected from-me page_size: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"mine"}` {
				t.Fatalf("unexpected from-me conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"resources":[{"id":12,"name":"我的资源"}],"page":1,"page_size":4,"pages":1,"total":1}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/shared-resources-to-me":
			if got := r.URL.Query().Get("page"); got != "3" {
				t.Fatalf("unexpected share-to-me page: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "7" {
				t.Fatalf("unexpected share-to-me page_size: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"share"}` {
				t.Fatalf("unexpected share-to-me conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"resources":[{"id":13,"name":"共享给我"}],"page":3,"page_size":7,"pages":4,"total":22}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/shared-resources/most-liked":
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"AI"}` {
				t.Fatalf("unexpected most-liked conditions: %q", got)
			}
			_, _ = w.Write([]byte(`{"resources":[{"id":10,"name":"最受欢迎"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/shared-resources/recent-used":
			if got := r.URL.Query().Get("classificationId"); got != "12" {
				t.Fatalf("unexpected classificationId: %q", got)
			}
			if got := r.URL.Query().Get("departmentIds"); got != "5,6" {
				t.Fatalf("unexpected departmentIds: %q", got)
			}
			_, _ = w.Write([]byte(`{"resources":[{"id":11,"name":"最近使用"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/subject-libs/folders":
			if got := r.URL.Query().Get("parent_id"); got != "0" {
				t.Fatalf("unexpected parent_id: %q", got)
			}
			_, _ = w.Write([]byte(`{"folders":[{"id":4,"name":"根目录","has_sub_folder":true}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/subject-libs/7":
			_, _ = w.Write([]byte(`{"id":7,"title":"期末题库","type":"folder","parent_id":0}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	classifications, err := service.ListSharedResourceClassificationsWithPrefix(ctx, true)
	if err != nil || len(classifications.Classifications) != 1 || classifications.Classifications[0].Name != "精选" {
		t.Fatalf("unexpected classifications: %#v, err=%v", classifications, err)
	}

	departments, err := service.ListDepartmentsWithPrefix(ctx, true, "id,name,code,parent_id,stopped,short_name,is_show_on_homepage,cover")
	if err != nil || len(departments.Departments) != 1 || departments.Departments[0].Code == nil || *departments.Departments[0].Code != "CS" {
		t.Fatalf("unexpected departments: %#v, err=%v", departments, err)
	}

	homepageDepartments, err := service.ListHomepageDepartmentsWithPrefix(ctx, true)
	if err != nil || len(homepageDepartments.Departments) != 1 || homepageDepartments.Departments[0].Name != "数学学院" {
		t.Fatalf("unexpected homepage departments: %#v, err=%v", homepageDepartments, err)
	}

	homepageResources, err := service.ListHomepageSharedResources(ctx, true, ListHomepageSharedResourcesParams{DepartmentID: 5, ClassificationID: 9})
	if err != nil || len(homepageResources.Resources) != 1 || homepageResources.Resources[0].Name != "首页资源" {
		t.Fatalf("unexpected homepage resources: %#v, err=%v", homepageResources, err)
	}

	searchResult, err := service.SearchSharedResourcesWithPrefix(ctx, true, ListSharedResourcesParams{Page: 2, PageSize: 6, Conditions: `{"keyword":"AI"}`})
	if err != nil || len(searchResult.Resources) != 1 || searchResult.Resources[0].ID != 9 {
		t.Fatalf("unexpected search result: %#v, err=%v", searchResult, err)
	}

	fromMe, err := service.ListSharedResourcesFromMeWithPrefix(ctx, true, ListSharedResourcesParams{Page: 1, PageSize: 4, Conditions: `{"keyword":"mine"}`})
	if err != nil || len(fromMe.Resources) != 1 || fromMe.Resources[0].ID != 12 {
		t.Fatalf("unexpected from-me result: %#v, err=%v", fromMe, err)
	}

	sharedToMe, err := service.ListSharedResourcesToMeWithParams(ctx, ListSharedResourcesParams{Page: 3, PageSize: 7, Conditions: `{"keyword":"share"}`})
	if err != nil || len(sharedToMe.Resources) != 1 || sharedToMe.Resources[0].ID != 13 {
		t.Fatalf("unexpected share-to-me result: %#v, err=%v", sharedToMe, err)
	}

	mostLiked, err := service.ListMostLikedSharedResourcesWithPrefix(ctx, true, `{"keyword":"AI"}`)
	if err != nil || len(mostLiked.Resources) != 1 || mostLiked.Resources[0].ID != 10 {
		t.Fatalf("unexpected most-liked result: %#v, err=%v", mostLiked, err)
	}

	recentUsed, err := service.ListRecentUsedSharedResourcesWithPrefix(ctx, true, ListRecentUsedSharedResourcesParams{ClassificationID: "12", DepartmentIDs: "5,6"})
	if err != nil || len(recentUsed.Resources) != 1 || recentUsed.Resources[0].ID != 11 {
		t.Fatalf("unexpected recent-used result: %#v, err=%v", recentUsed, err)
	}

	folders, err := service.ListSubjectLibFolders(context.Background(), 0)
	if err != nil || len(folders.Folders) != 1 || folders.Folders[0].Name != "根目录" || !folders.Folders[0].HasSubFolder {
		t.Fatalf("unexpected folders: %#v, err=%v", folders, err)
	}

	subjectLib, err := service.GetSubjectLibWithPrefix(ctx, true, 7)
	if err != nil || subjectLib.Title != "期末题库" {
		t.Fatalf("unexpected subject lib: %#v, err=%v", subjectLib, err)
	}
}
