package admin

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
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

func TestGetLangSettingsUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/orgs/1/lang-settings" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"lang_settings":["zh-CN"]}`), nil
	})

	resp, err := svc.GetLangSettings(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetLangSettings error: %v", err)
	}
	if len(resp.LangSettings) != 1 {
		t.Fatalf("unexpected lang settings length: %d", len(resp.LangSettings))
	}
}

func TestGetOutlineSettingUsesTypedResponse(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/outline-setting" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"id":1,"org_id":1,"formatted_options":[],"formatted_default_options":[]}`), nil
	})

	resp, err := svc.GetOutlineSetting(context.Background())
	if err != nil {
		t.Fatalf("GetOutlineSetting error: %v", err)
	}
	if resp.OrgID != 1 {
		t.Fatalf("unexpected org id: %d", resp.OrgID)
	}
}

func TestListAcademicYearsUsesTypedWrapper(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/academic-years" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"academic_years":[]}`), nil
	})

	resp, err := svc.ListAcademicYears(context.Background())
	if err != nil {
		t.Fatalf("ListAcademicYears error: %v", err)
	}
	if resp.AcademicYears == nil {
		t.Fatal("expected academic_years slice")
	}
}

func TestListDepartmentsUsesTypedWrapper(t *testing.T) {
	svc := newTestService(t, func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", req.Method)
		}
		if req.URL.Path != "/api/departments" {
			t.Fatalf("unexpected path: %s", req.URL.Path)
		}
		return jsonResponse(`{"departments":[]}`), nil
	})

	resp, err := svc.ListDepartments(context.Background(), map[string]string{"fields": "id,name"})
	if err != nil {
		t.Fatalf("ListDepartments error: %v", err)
	}
	if resp.Departments == nil {
		t.Fatal("expected departments slice")
	}
}
