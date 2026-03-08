package meetings

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestListLessonRooms(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/lesson-rooms" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"room_code":"A101","room_name":"阶梯教室","app_id":"app-1"}]`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.ListLessonRooms(context.Background())
	if err != nil {
		t.Fatalf("ListLessonRooms returned error: %v", err)
	}
	if len(result) != 1 || result[0].RoomCode != "A101" || result[0].RoomName != "阶梯教室" {
		t.Fatalf("unexpected lesson rooms: %#v", result)
	}
}

func TestListRoomLocations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/course/18/room-locations" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"rooms":[{"id":1,"building":"紫金港东1A","room_name":"101","room_code":"E1A-101","seats":120}]}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.ListRoomLocations(context.Background(), 18)
	if err != nil {
		t.Fatalf("ListRoomLocations returned error: %v", err)
	}
	if len(result.Rooms) != 1 || result.Rooms[0].RoomCode != "E1A-101" {
		t.Fatalf("unexpected room locations: %#v", result.Rooms)
	}
}

func TestListEnabledRoomLocations(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/org/25/enable-room-locations" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("start_time"); got != "2026-03-09T08:00:00Z" {
			t.Fatalf("unexpected start_time: %s", got)
		}
		if got := r.URL.Query().Get("end_time"); got != "2026-03-09T10:00:00Z" {
			t.Fatalf("unexpected end_time: %s", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"rooms":[{"id":2,"building":"紫金港东2B","room_name":"202","room_code":"E2B-202","seats":80}]}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.ListEnabledRoomLocations(context.Background(), 25, "2026-03-09T08:00:00Z", "2026-03-09T10:00:00Z")
	if err != nil {
		t.Fatalf("ListEnabledRoomLocations returned error: %v", err)
	}
	if len(result.Rooms) != 1 || result.Rooms[0].Building != "紫金港东2B" {
		t.Fatalf("unexpected enabled room locations: %#v", result.Rooms)
	}
}

func TestGetZoomSettings(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/orgs/25/zoom-settings" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"org_zoom_settings":{"mode":"share","basic_default_recording_type":"local"}}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.GetZoomSettings(context.Background(), 25)
	if err != nil {
		t.Fatalf("GetZoomSettings returned error: %v", err)
	}
	if result.OrgZoomSettings == nil || result.OrgZoomSettings.Mode != "share" {
		t.Fatalf("unexpected zoom settings: %#v", result)
	}
}

func TestGetZoomUserInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/user/18/zoom-info" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"zoom_info":{"code":429,"message":"rate limit","type":1,"email":"teacher@zju.edu.cn"}}`))
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	result, err := service.GetZoomUserInfo(context.Background(), 18)
	if err != nil {
		t.Fatalf("GetZoomUserInfo returned error: %v", err)
	}
	if result.ZoomInfo == nil || result.ZoomInfo.Email != "teacher@zju.edu.cn" || result.ZoomInfo.Code != 429 {
		t.Fatalf("unexpected zoom user info: %#v", result)
	}
}
