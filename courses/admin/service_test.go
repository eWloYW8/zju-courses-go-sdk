package admin

import (
	"context"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestGetJobUsesStatusEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/jobs/42/status" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"finished","result":{"job_id":42}}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.GetJob(context.Background(), 42)
	if err != nil {
		t.Fatalf("GetJob returned error: %v", err)
	}
	if result.Status != "finished" || result.Result["job_id"].(float64) != 42 {
		t.Fatalf("unexpected result: %#v", result)
	}
}

func TestOrgSettingHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/orgs/9/settings":
			_, _ = w.Write([]byte(`{"watermark_enabled":true}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/orgs/9/live-record-settings":
			_, _ = w.Write([]byte(`{"auto_publish":false}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/portal-logo":
			_, _ = w.Write([]byte(`{"logo_url":"https://example.com/logo.png"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/orgs/9/alert-popup":
			_, _ = w.Write([]byte(`{"enabled":true,"content":"notice"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/all-orgs":
			_, _ = w.Write([]byte(`{"orgs":[{"id":1,"name":"ZJU"},{"id":2,"name":"CAD"}]}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/orgs/9/settings":
			if got := r.URL.Query().Get("form_type"); got != "2" {
				t.Fatalf("unexpected form_type: %q", got)
			}
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode body: %v", err)
			}
			if body["watermark_enabled"] != true {
				t.Fatalf("unexpected settings body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/orgs/9/portal-logo":
			if got := r.URL.Query().Get("upload_id"); got != "33" {
				t.Fatalf("unexpected upload_id: %q", got)
			}
			data, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("read body: %v", err)
			}
			if string(data) != "null" && len(data) != 0 {
				t.Fatalf("unexpected portal-logo body: %q", string(data))
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/orgs/9/alert-popup":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode alert body: %v", err)
			}
			if body["enabled"] != true {
				t.Fatalf("unexpected alert body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	settings, err := service.GetOrgSetting(ctx, 9)
	if err != nil || !(*settings)["watermark_enabled"].(bool) {
		t.Fatalf("unexpected org settings: %#v, err=%v", settings, err)
	}

	liveSettings, err := service.GetLiveRecordOrgSetting(ctx, 9)
	if err != nil || (*liveSettings)["auto_publish"].(bool) {
		t.Fatalf("unexpected live-record settings: %#v, err=%v", liveSettings, err)
	}

	logo, err := service.GetOrgPortalLogo(ctx)
	if err != nil || (*logo)["logo_url"] != "https://example.com/logo.png" {
		t.Fatalf("unexpected portal logo: %#v, err=%v", logo, err)
	}

	alert, err := service.GetAlertPopupSettings(ctx, 9)
	if err != nil || (*alert)["content"] != "notice" {
		t.Fatalf("unexpected alert settings: %#v, err=%v", alert, err)
	}

	orgs, err := service.ListAllOrgDetails(ctx)
	if err != nil || len(orgs.Orgs) != 2 || orgs.Orgs[1].Name != "CAD" {
		t.Fatalf("unexpected org list: %#v, err=%v", orgs, err)
	}

	if err := service.UpdateOrgSetting(ctx, 9, "2", UpdateOrgSettingRequest{"watermark_enabled": true}); err != nil {
		t.Fatalf("UpdateOrgSetting returned error: %v", err)
	}
	if err := service.UpdateOrgPortalLogo(ctx, 9, 33); err != nil {
		t.Fatalf("UpdateOrgPortalLogo returned error: %v", err)
	}
	if err := service.UpdateAlertPopupSetting(ctx, 9, UpdateAlertPopupSettingRequest{"enabled": true}); err != nil {
		t.Fatalf("UpdateAlertPopupSetting returned error: %v", err)
	}
}

func TestCourseCopyAndImportHelpersUseFrontendPayloads(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/course-copy/courses":
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"calc"}` {
				t.Fatalf("unexpected conditions query: %q", got)
			}
			if got := r.URL.Query().Get("fields"); got != "id,name,course_code" {
				t.Fatalf("unexpected fields query: %q", got)
			}
			_, _ = w.Write([]byte(`{"courses":[{"id":1,"name":"Calculus","course_code":"MATH100"}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/course-copy/copy":
			var body CopyCourseRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode copy body: %v", err)
			}
			if len(body.CourseIDs) != 2 || body.CourseIDs[0] != 8 || body.CourseIDs[1] != 9 {
				t.Fatalf("unexpected copy body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/12/moodle/import":
			var body MoodleImportRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode moodle import body: %v", err)
			}
			if body.UploadID != 77 {
				t.Fatalf("unexpected moodle import body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPost && r.URL.Path == "/api/course/mbz/import":
			mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
			if err != nil {
				t.Fatalf("parse content-type: %v", err)
			}
			if mediaType != "multipart/form-data" {
				t.Fatalf("unexpected content-type: %s", mediaType)
			}
			reader := multipart.NewReader(r.Body, params["boundary"])
			part, err := reader.NextPart()
			if err != nil {
				t.Fatalf("read multipart part: %v", err)
			}
			if part.FormName() != "upload_id" {
				t.Fatalf("unexpected form field: %s", part.FormName())
			}
			data, err := io.ReadAll(part)
			if err != nil {
				t.Fatalf("read multipart data: %v", err)
			}
			if strings.TrimSpace(string(data)) != "88" {
				t.Fatalf("unexpected upload_id payload: %q", string(data))
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/task/last":
			if got := r.URL.Query().Get("no-intercept"); got != "true" {
				t.Fatalf("unexpected no-intercept: %q", got)
			}
			if got := r.URL.Query().Get("type"); got != "import_course_from_mbz" {
				t.Fatalf("unexpected task type: %q", got)
			}
			_, _ = w.Write([]byte(`[{"status":"success","output":{"imported":{"course":{"id":99}}}}]`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	courses, err := service.ListCopyableCoursesWithQuery(ctx, &CopyableCoursesQuery{Keyword: "calc"})
	if err != nil {
		t.Fatalf("ListCopyableCoursesWithQuery returned error: %v", err)
	}
	if len(courses.Courses) != 1 || courses.Courses[0].CourseCode != "MATH100" {
		t.Fatalf("unexpected copyable courses: %#v", courses)
	}

	if err := service.CopyCourses(ctx, &CopyCourseRequest{CourseIDs: []int{8, 9}}); err != nil {
		t.Fatalf("CopyCourses returned error: %v", err)
	}
	if err := service.ImportCourseMoodlePackage(ctx, 12, &MoodleImportRequest{UploadID: 77}); err != nil {
		t.Fatalf("ImportCourseMoodlePackage returned error: %v", err)
	}
	if err := service.ImportMBZUpload(ctx, 88); err != nil {
		t.Fatalf("ImportMBZUpload returned error: %v", err)
	}

	tasks, err := service.GetLastTasks(ctx, "import_course_from_mbz")
	if err != nil {
		t.Fatalf("GetLastTasks returned error: %v", err)
	}
	if len(tasks) != 1 || tasks[0].Output == nil || tasks[0].Output.Imported == nil || tasks[0].Output.Imported.Course == nil || tasks[0].Output.Imported.Course.ID != 99 {
		t.Fatalf("unexpected tasks: %#v", tasks)
	}
}

func TestOrgHeaderHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/orgs/9/lang-settings":
			_, _ = w.Write([]byte(`{"lang_settings":["zh-CN","en-US"]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/orgs/9/lang-settings":
			_, _ = w.Write([]byte(`{"lang_settings":["zh-CN"]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/anonymous-api/orgs/9/login-settings":
			_, _ = w.Write([]byte(`{"login_settings":{"enable_wechat":true}}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	memberLangs, err := service.GetLangSettings(ctx, 9)
	if err != nil || len(memberLangs.LangSettings) != 2 {
		t.Fatalf("unexpected member lang settings: %#v, err=%v", memberLangs, err)
	}

	anonymousLangs, err := service.GetLangSettingsWithPrefix(ctx, 9, true)
	if err != nil || len(anonymousLangs.LangSettings) != 1 || anonymousLangs.LangSettings[0] != "zh-CN" {
		t.Fatalf("unexpected anonymous lang settings: %#v, err=%v", anonymousLangs, err)
	}

	loginSettings, err := service.GetLoginSettings(ctx, 9)
	if err != nil || loginSettings.LoginSettings["enable_wechat"] != true {
		t.Fatalf("unexpected login settings: %#v, err=%v", loginSettings, err)
	}
}

func TestPlanHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/org/7/org-plan-info":
			_, _ = w.Write([]byte(`{
				"org":{"storage_total":2048,"storage_used":512,"transfer_total":1024,"transfer_used":128},
				"org_plan":{"id":3,"start_date":"2026-03-01","end_date":"2026-03-31","expired":false,"trial":false,"plan":{"storage_total":2048,"transfer_per_month":1024}}
			}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/org/request/change-plan/9":
			var body RequestChangePlanRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode change-plan body: %v", err)
			}
			if body.Period != "yearly" {
				t.Fatalf("unexpected change-plan body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"ok":true}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/custom-operation":
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	info, err := service.GetOrgPlanInfo(ctx, 7)
	if err != nil {
		t.Fatalf("GetOrgPlanInfo returned error: %v", err)
	}
	if info.Org == nil || info.Org.StorageTotal != 2048 || info.OrgPlan == nil || info.OrgPlan.Plan == nil || info.OrgPlan.Plan.TransferPerMonth != 1024 {
		t.Fatalf("unexpected org plan info: %#v", info)
	}

	if _, err := service.RequestChangePlan(ctx, 9, &RequestChangePlanRequest{Period: "yearly"}); err != nil {
		t.Fatalf("RequestChangePlan returned error: %v", err)
	}
	if err := service.WGAdminOperateRequest(ctx, "/api/custom-operation"); err != nil {
		t.Fatalf("WGAdminOperateRequest returned error: %v", err)
	}
}

func TestDepartmentAndOutlineSettingHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/departments-for-user":
			_, _ = w.Write([]byte(`{"departments":[{"id":1,"name":"计算机学院"}]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/source-department-code-for-user":
			_, _ = w.Write([]byte(`{"department_code":"CS"}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/user/department":
			var body UpdateUserDepartmentRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode user/department body: %v", err)
			}
			if body.DepartmentID != 7 {
				t.Fatalf("unexpected user/department body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/outline-setting":
			_, _ = w.Write([]byte(`{"id":3,"formatted_options":[{"key":"course_objective","title":"课程目标","required":true}]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/outline-setting/3":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/outline-setting/3":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/outline-setting/3/sort":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodDelete && r.URL.Path == "/api/outline-setting/3/option/course_objective":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/outline-setting/toggle":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/outline-setting/3/required-options":
			var body SaveOutlineRequiredOptionsRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode required-options body: %v", err)
			}
			if len(body.RequiredOptions) != 2 || body.RequiredOptions[0] != "course_objective" {
				t.Fatalf("unexpected required-options body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	departments, err := service.ListDepartmentsForUser(ctx)
	if err != nil || len(departments.Departments) != 1 || departments.Departments[0].Name != "计算机学院" {
		t.Fatalf("unexpected departments-for-user: %#v, err=%v", departments, err)
	}

	sourceCode, err := service.GetSourceDepartmentCodeForUser(ctx)
	if err != nil || sourceCode.DepartmentCode != "CS" {
		t.Fatalf("unexpected source department code: %#v, err=%v", sourceCode, err)
	}

	if err := service.UpdateUserDepartment(ctx, &UpdateUserDepartmentRequest{DepartmentID: 7}); err != nil {
		t.Fatalf("UpdateUserDepartment returned error: %v", err)
	}

	outlineSetting, err := service.GetOutlineSetting(ctx)
	if err != nil || outlineSetting.ID != 3 || len(outlineSetting.FormattedOptions) != 1 || outlineSetting.FormattedOptions[0].Key != "course_objective" {
		t.Fatalf("unexpected outline setting: %#v, err=%v", outlineSetting, err)
	}

	if err := service.AddOutlineSetting(ctx, 3, map[string]any{"key": "course_objective"}); err != nil {
		t.Fatalf("AddOutlineSetting returned error: %v", err)
	}
	if err := service.UpdateOutlineSetting(ctx, 3, map[string]any{"title": "课程目标"}); err != nil {
		t.Fatalf("UpdateOutlineSetting returned error: %v", err)
	}
	if err := service.SortOutlineSetting(ctx, 3, map[string]any{"formatted_options": []string{"course_objective"}}); err != nil {
		t.Fatalf("SortOutlineSetting returned error: %v", err)
	}
	if err := service.DeleteOutlineSettingOption(ctx, 3, "course_objective"); err != nil {
		t.Fatalf("DeleteOutlineSettingOption returned error: %v", err)
	}
	if err := service.ToggleOutlineSetting(ctx, map[string]any{"enabled": true}); err != nil {
		t.Fatalf("ToggleOutlineSetting returned error: %v", err)
	}
	if err := service.SaveOutlineRequiredOptions(ctx, 3, &SaveOutlineRequiredOptionsRequest{
		RequiredOptions: []string{"course_objective", "teaching_schedule"},
	}); err != nil {
		t.Fatalf("SaveOutlineRequiredOptions returned error: %v", err)
	}
}
