package linear

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdateComment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request graphQLRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if !strings.Contains(request.Query, "commentUpdate") {
			t.Errorf("query does not contain commentUpdate: %q", request.Query)
		}
		if request.Variables["id"] != "comment-id" {
			t.Errorf("id = %v, want comment-id", request.Variables["id"])
		}
		input, ok := request.Variables["input"].(map[string]any)
		if !ok || input["body"] != "updated body" {
			t.Errorf("input = %#v, want body updated body", request.Variables["input"])
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"commentUpdate":{"success":true,"comment":{"id":"comment-id","body":"updated body","user":{"id":"user-id","name":"Ada"}}}}}`))
	}))
	defer server.Close()

	client := NewClient("api-key")
	client.endpoint = server.URL
	comment, err := client.UpdateComment(context.Background(), "comment-id", CommentUpdateInput{Body: "updated body"})
	if err != nil {
		t.Fatalf("UpdateComment() error = %v", err)
	}
	if comment.ID != "comment-id" || comment.Body != "updated body" {
		t.Errorf("UpdateComment() = %#v", comment)
	}
	if comment.User == nil || comment.User.Name != "Ada" {
		t.Errorf("UpdateComment() user = %#v", comment.User)
	}
}

func TestUpdateCommentUnsuccessful(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"commentUpdate":{"success":false,"comment":null}}}`))
	}))
	defer server.Close()

	client := NewClient("api-key")
	client.endpoint = server.URL
	_, err := client.UpdateComment(context.Background(), "comment-id", CommentUpdateInput{Body: "updated body"})
	if err == nil || !strings.Contains(err.Error(), "success=false") {
		t.Fatalf("UpdateComment() error = %v, want success=false error", err)
	}
}
