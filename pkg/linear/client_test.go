package linear

import (
	"context"
	"os"
	"testing"
)

func TestGetIssueByIdentifier(t *testing.T) {
	apiKey := os.Getenv("LINEAR_API_KEY")
	if apiKey == "" {
		t.Skip("LINEAR_API_KEY not set")
	}

	client := NewClient(apiKey)

	const query = `
	query {
	  issues(filter: { team: { key: { eq: "ETERNIS" } }, number: { eq: 3641 } }) {
	    nodes {
	      id identifier title description priority
	      state { id name type }
	      team  { id name key }
	      assignee { id name email }
	    }
	  }
	}`

	var resp struct {
		Issues IssueConnection `json:"issues"`
	}
	err := client.do(context.Background(), query, nil, &resp)
	if err != nil {
		t.Fatalf("query failed: %v", err)
	}

	if len(resp.Issues.Nodes) == 0 {
		t.Fatal("expected at least one issue, got none")
	}

	issue := resp.Issues.Nodes[0]
	if issue.Identifier != "ETERNIS-3641" {
		t.Fatalf("expected identifier ETERNIS-3641, got %s", issue.Identifier)
	}

	t.Logf("Issue: %s — %s", issue.Identifier, issue.Title)
	t.Logf("State: %s, Priority: %d", issue.State.Name, issue.Priority)
	if issue.Assignee != nil {
		t.Logf("Assignee: %s", issue.Assignee.Name)
	}
}
