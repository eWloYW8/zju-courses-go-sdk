package interactions

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eWloYW8/zju-courses-go-sdk/internal/sdk"
)

func TestInteractionScoreAndVoteHelpersUseFrontendEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPut && r.URL.Path == "/api/courses/interactions/21/score":
			var body UpdateInteractionScoreRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode interaction score body: %v", err)
			}
			if body["score"] != float64(10) {
				t.Fatalf("unexpected interaction score body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodPut && r.URL.Path == "/api/courses/interactions/vote/22":
			var body UpdateInteractionVoteRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decode interaction vote body: %v", err)
			}
			if body["is_score_public"] != true {
				t.Fatalf("unexpected interaction vote body: %#v", body)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	service := New(sdk.NewClient(sdk.WithBaseURL(server.URL)))
	ctx := context.Background()

	if err := service.UpdateInteractionScore(ctx, 21, &UpdateInteractionScoreRequest{"score": 10}); err != nil {
		t.Fatalf("UpdateInteractionScore returned error: %v", err)
	}
	if err := service.VoteInteraction(ctx, 22, &UpdateInteractionVoteRequest{"is_score_public": true}); err != nil {
		t.Fatalf("VoteInteraction returned error: %v", err)
	}
}
