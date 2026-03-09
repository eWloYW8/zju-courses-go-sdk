package meetings

import (
	"context"
	"encoding/json"
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

func TestVTRSResourceHelpersUseFrontendEndpointsAndModels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/vtrses/9/resources/summary":
			_, _ = w.Write([]byte(`{"materials":3,"exercise_libs":2}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/stat/vtrses/9/resources":
			if got := r.URL.Query().Get("date_range"); got != "2026-03-01,2026-03-09" {
				t.Fatalf("unexpected date_range: %s", got)
			}
			_, _ = w.Write([]byte(`{"subject_lib":4,"video":2,"audio":1,"custom_type":7}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/vtrses/9/resources":
			if got := r.URL.Query().Get("parent_folder_id"); got != "5" {
				t.Fatalf("unexpected parent_folder_id: %s", got)
			}
			if got := r.URL.Query().Get("classification_id"); got != "8" {
				t.Fatalf("unexpected classification_id: %s", got)
			}
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected page: %s", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "20" {
				t.Fatalf("unexpected page_size: %s", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"线性代数"}` {
				t.Fatalf("unexpected conditions: %s", got)
			}
			_, _ = w.Write([]byte(`{
				"items":[
					{
						"id":44,
						"name":"期中复习资料",
						"created_by_id":12,
						"allow_download":true,
						"vtrs_name":"智慧教室A",
						"vtrs_code":"A101",
						"allow_copy":true,
						"cc_license_name":"CC BY-NC",
						"created_at":"2026-03-09T08:00:00Z",
						"created_by":{"id":12,"name":"张老师"},
						"upload":{
							"id":88,
							"name":"review.pdf",
							"size":1024,
							"type":"document",
							"status":"ready",
							"allow_download":true,
							"origin_allow_download":true,
							"reference_id":701,
							"created_at":"2026-03-09T07:59:00Z"
						}
					}
				],
				"page":2,
				"page_size":20,
				"pages":3,
				"total":41,
				"start":21,
				"end":40
			}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/vtrses/share-resources":
			if got := r.URL.Query().Get("ref_parent_type"); got != "course" {
				t.Fatalf("unexpected ref_parent_type: %s", got)
			}
			_, _ = w.Write([]byte(`{
				"items":[{"id":45,"name":"共享资源","upload":{"id":89,"reference_id":702}}],
				"page":1,
				"page_size":10,
				"pages":1,
				"total":1,
				"start":1,
				"end":1
			}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/vtrses/9/resources/folder":
			var body CreateVTRSResourceFolderRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode folder body: %v", err)
			}
			if body.Name != "复习资料" || body.ParentFolderID == nil || *body.ParentFolderID != 5 || body.ClassificationID == nil || *body.ClassificationID != 8 {
				t.Fatalf("unexpected folder body: %#v", body)
			}
			_, _ = w.Write([]byte(`{"id":51,"name":"复习资料"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/vtrses/9/resources/category/4/structure":
			_, _ = w.Write([]byte(`{
				"classifications":[
					{
						"id":4,
						"name":"课程资源",
						"sort":1,
						"has_sub_folder":true,
						"children":[
							{
								"id":6,
								"name":"讲义",
								"classification_id":4,
								"parent_id":0,
								"is_folder":true,
								"has_sub_folder":false,
								"children":[]
							}
						]
					}
				]
			}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/vtrses/9/resources/move":
			var body MoveVTRSResourcesRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode move body: %v", err)
			}
			if len(body.UploadReferenceIDs) != 2 || body.UploadReferenceIDs[0] != 701 || body.ParentFolderID == nil || *body.ParentFolderID != 6 {
				t.Fatalf("unexpected move body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/vtrses/resources/classifications/8/resource-operation-pre-check":
			if got := r.URL.Query().Get("upload_reference_ids"); got != "701,702" {
				t.Fatalf("unexpected upload_reference_ids: %s", got)
			}
			if got := r.URL.Query().Get("operation_type"); got != "move" {
				t.Fatalf("unexpected operation_type: %s", got)
			}
			_, _ = w.Write([]byte(`{"ok":true}`))
		case r.Method == http.MethodPut && r.URL.Path == "/api/vtrses/9/resource/44":
			var body UpdateVTRSResourceRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode update body: %v", err)
			}
			if body.Name != "期中资料更新版" {
				t.Fatalf("unexpected update body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	summary, err := service.GetVTRSResourcesSummary(context.Background(), 9)
	if err != nil {
		t.Fatalf("GetVTRSResourcesSummary returned error: %v", err)
	}
	if summary.All != 5 || summary.Materials != 3 || summary.ExerciseLibs != 2 {
		t.Fatalf("unexpected summary: %#v", summary)
	}

	stat, err := service.GetVTRSResourcesStat(context.Background(), 9, &ListVTRSResourcesStatParams{
		DateRange: "2026-03-01,2026-03-09",
	})
	if err != nil {
		t.Fatalf("GetVTRSResourcesStat returned error: %v", err)
	}
	if stat.SubjectLib != 4 || stat.Video != 2 || stat.Audio != 1 || stat.Other != 7 {
		t.Fatalf("unexpected stat: %#v", stat)
	}

	parentFolderID := 5
	classificationID := 8
	resources, err := service.ListVTRSResources(context.Background(), 9, &ListVTRSResourcesParams{
		ParentFolderID:   &parentFolderID,
		ClassificationID: &classificationID,
		Page:             2,
		PageSize:         20,
		Conditions:       `{"keyword":"线性代数"}`,
	})
	if err != nil {
		t.Fatalf("ListVTRSResources returned error: %v", err)
	}
	if len(resources.Items) != 1 || resources.Items[0].Upload == nil || resources.Items[0].Upload.ReferenceID != 701 {
		t.Fatalf("unexpected resources: %#v", resources.Items)
	}
	if resources.Items[0].CreatedBy == nil || resources.Items[0].CreatedBy.Name != "张老师" {
		t.Fatalf("created_by did not decode: %#v", resources.Items[0].CreatedBy)
	}

	shared, err := service.ListVTRSShareResources(context.Background(), &ListVTRSShareResourcesParams{
		RefParentType: "course",
		Page:          1,
		PageSize:      10,
	})
	if err != nil {
		t.Fatalf("ListVTRSShareResources returned error: %v", err)
	}
	if len(shared.Items) != 1 || shared.Items[0].Upload == nil || shared.Items[0].Upload.ReferenceID != 702 {
		t.Fatalf("unexpected shared resources: %#v", shared.Items)
	}

	if _, err := service.CreateVTRSResourceFolder(context.Background(), 9, &CreateVTRSResourceFolderRequest{
		Name:             "复习资料",
		ParentFolderID:   &parentFolderID,
		ClassificationID: &classificationID,
	}); err != nil {
		t.Fatalf("CreateVTRSResourceFolder returned error: %v", err)
	}

	structure, err := service.GetVTRSResourceCategoryStructure(context.Background(), 9, 4)
	if err != nil {
		t.Fatalf("GetVTRSResourceCategoryStructure returned error: %v", err)
	}
	if len(structure.Classifications) != 1 || len(structure.Classifications[0].Children) != 1 || !structure.Classifications[0].Children[0].IsFolder {
		t.Fatalf("unexpected structure: %#v", structure.Classifications)
	}

	moveFolderID := 6
	if err := service.MoveVTRSResources(context.Background(), 9, &MoveVTRSResourcesRequest{
		UploadReferenceIDs: []int{701, 702},
		ParentFolderID:     &moveFolderID,
		ClassificationID:   &classificationID,
	}); err != nil {
		t.Fatalf("MoveVTRSResources returned error: %v", err)
	}

	precheck, err := service.GetVTRSResourceOperationPreCheck(context.Background(), 8, []int{701, 702}, "move")
	if err != nil {
		t.Fatalf("GetVTRSResourceOperationPreCheck returned error: %v", err)
	}
	if ok, _ := precheck["ok"].(bool); !ok {
		t.Fatalf("unexpected precheck: %#v", precheck)
	}

	if err := service.UpdateVTRSResource(context.Background(), 9, 44, &UpdateVTRSResourceRequest{Name: "期中资料更新版"}); err != nil {
		t.Fatalf("UpdateVTRSResource returned error: %v", err)
	}
}

func TestVTRSMemberAndSubjectLibHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/api/vtrses/9/meetings":
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected meeting page query: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "20" {
				t.Fatalf("unexpected meeting page_size query: %q", got)
			}
			if got := r.URL.Query().Get("classification_id"); got != "8" {
				t.Fatalf("unexpected classification_id query: %q", got)
			}
			if got := r.URL.Query().Get("meeting_format"); got != "online" {
				t.Fatalf("unexpected meeting_format query: %q", got)
			}
			_, _ = w.Write([]byte(`{"items":[],"page":2,"page_size":20,"pages":1,"total":0}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/vtrses/9/members":
			if got := r.URL.Query().Get("keyword"); got != "张" {
				t.Fatalf("unexpected member keyword query: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "5" {
				t.Fatalf("unexpected member page_size query: %q", got)
			}
			if got := r.URL.Query().Get("fields"); got != "id,user" {
				t.Fatalf("unexpected member fields query: %q", got)
			}
			_, _ = w.Write([]byte(`{
				"members":[
					{"id":7,"user":{"id":12,"user_no":"2025001","name":"张老师","avatar_small_url":"avatar.png","email":"teacher@zju.edu.cn","mobile_phone":"13800000000","department":{"id":3,"name":"计算机学院"}}}
				],
				"page":1,
				"page_size":5,
				"pages":1,
				"total":1
			}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/api/vtrses/9/members/7":
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/vtrses/9/transfer-owner":
			var body TransferVTRSOwnerRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode transfer-owner body: %v", err)
			}
			if body.OwnerID != 12 {
				t.Fatalf("unexpected transfer-owner body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/vtrses/9/members":
			var body UserIDsRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode add-members body: %v", err)
			}
			if len(body.UserIDs) != 2 || body.UserIDs[0] != 12 || body.UserIDs[1] != 18 {
				t.Fatalf("unexpected add-members body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/api/vtrses/9/select-members":
			if got := r.URL.Query().Get("page"); got != "2" {
				t.Fatalf("unexpected select-members page query: %q", got)
			}
			if got := r.URL.Query().Get("page_size"); got != "8" {
				t.Fatalf("unexpected select-members page_size query: %q", got)
			}
			if got := r.URL.Query().Get("conditions"); got != `{"keyword":"李老师"}` {
				t.Fatalf("unexpected select-members conditions query: %q", got)
			}
			_, _ = w.Write([]byte(`{"items":[{"id":18,"name":"李老师"}]}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/api/vtrses/9/patrol":
			_, _ = w.Write([]byte(`{"status":"deleted"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/vtrses/9/subject-libs":
			if got := r.URL.Query().Get("keyword"); got != "题库" {
				t.Fatalf("unexpected subject-libs keyword query: %q", got)
			}
			if got := r.URL.Query().Get("parent_id"); got != "3" {
				t.Fatalf("unexpected subject-libs parent_id query: %q", got)
			}
			if got := r.URL.Query().Get("classification_id"); got != "8" {
				t.Fatalf("unexpected subject-libs classification_id query: %q", got)
			}
			if got := r.URL.Query().Get("predicate"); got != "title" {
				t.Fatalf("unexpected subject-libs predicate query: %q", got)
			}
			if got := r.URL.Query().Get("reverse"); got != "true" {
				t.Fatalf("unexpected subject-libs reverse query: %q", got)
			}
			_, _ = w.Write([]byte(`{
				"subject_libs":[{"id":5,"title":"题库A","parent_id":3,"classification_id":8,"is_folder":true,"is_shared":true,"nums":4,"type":"folder","created_at":"2026-03-09","updated_at":"2026-03-10"}],
				"page":1,
				"page_size":10,
				"pages":1,
				"total":1,
				"start":1,
				"end":1
			}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	classificationID := 8
	parentID := 3

	if _, err := service.ListVTRSMeetings(context.Background(), 9, &ListVTRSMeetingsParams{
		Page:             2,
		PageSize:         20,
		ClassificationID: &classificationID,
		MeetingFormat:    "online",
	}); err != nil {
		t.Fatalf("ListVTRSMeetings returned error: %v", err)
	}

	members, err := service.ListVTRSMembers(context.Background(), 9, &ListVTRSMembersParams{
		Keyword:  "张",
		Page:     1,
		PageSize: 5,
		Fields:   "id,user",
	})
	if err != nil || len(members.Items) != 1 || members.Items[0].Department != "计算机学院" || members.Items[0].UserID != 12 {
		t.Fatalf("unexpected VTRS members: %#v, err=%v", members, err)
	}

	if err := service.DeleteVTRSMember(context.Background(), 9, 7); err != nil {
		t.Fatalf("DeleteVTRSMember returned error: %v", err)
	}

	if err := service.TransferVTRSOwner(context.Background(), 9, 12); err != nil {
		t.Fatalf("TransferVTRSOwner returned error: %v", err)
	}

	if err := service.AddVTRSMembers(context.Background(), 9, []int{12, 18}); err != nil {
		t.Fatalf("AddVTRSMembers returned error: %v", err)
	}

	selected, err := service.SelectVTRSMembers(context.Background(), 9, &SelectVTRSMembersParams{
		Page:       2,
		PageSize:   8,
		Conditions: map[string]any{"keyword": "李老师"},
	})
	if err != nil || string(selected) != `{"items":[{"id":18,"name":"李老师"}]}` {
		t.Fatalf("unexpected select-members response: %s, err=%v", string(selected), err)
	}

	patrol, err := service.DeleteVTRSPatrol(context.Background(), 9)
	if err != nil || string(patrol) != `{"status":"deleted"}` {
		t.Fatalf("unexpected patrol response: %s, err=%v", string(patrol), err)
	}

	reverse := true
	subjectLibs, err := service.ListVTRSSubjectLibsWithParams(context.Background(), 9, &ListVTRSSubjectLibsParams{
		Keyword:          "题库",
		ParentID:         &parentID,
		ClassificationID: &classificationID,
		Page:             1,
		PageSize:         10,
		Predicate:        "title",
		Reverse:          &reverse,
	})
	if err != nil || len(subjectLibs.Items) != 1 || !subjectLibs.Items[0].IsShared || subjectLibs.Items[0].Nums != 4 {
		t.Fatalf("unexpected VTRS subject libs: %#v, err=%v", subjectLibs, err)
	}
}
