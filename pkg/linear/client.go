package linear

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const defaultEndpoint = "https://api.linear.app/graphql"

type Client struct {
	apiKey     string
	endpoint   string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		endpoint:   defaultEndpoint,
		httpClient: http.DefaultClient,
	}
}

type APIError struct {
	StatusCode int
	Errors     []graphQLError
}

func (e *APIError) Error() string {
	if len(e.Errors) > 0 {
		return fmt.Sprintf("linear API: %s", e.Errors[0].Message)
	}
	return fmt.Sprintf("linear API: HTTP %d", e.StatusCode)
}

func (c *Client) do(ctx context.Context, query string, vars map[string]any, dest any) error {
	body, err := json.Marshal(graphQLRequest{Query: query, Variables: vars})
	if err != nil {
		return fmt.Errorf("linear: marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("linear: create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("linear: send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("linear: read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{StatusCode: resp.StatusCode}
	}

	var gqlResp graphQLResponse
	if err := json.Unmarshal(respBody, &gqlResp); err != nil {
		return fmt.Errorf("linear: unmarshal response: %w", err)
	}

	if len(gqlResp.Errors) > 0 {
		return &APIError{Errors: gqlResp.Errors}
	}

	if dest != nil {
		if err := json.Unmarshal(gqlResp.Data, dest); err != nil {
			return fmt.Errorf("linear: unmarshal data: %w", err)
		}
	}
	return nil
}

func (c *Client) GetIssue(ctx context.Context, id string) (*Issue, error) {
	var resp struct {
		Issue Issue `json:"issue"`
	}
	if err := c.do(ctx, queryIssueGet, map[string]any{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.Issue, nil
}

func (c *Client) ListIssues(ctx context.Context, first int, after string, teamID string) (*IssueConnection, error) {
	vars := map[string]any{"first": first}
	if after != "" {
		vars["after"] = after
	}
	if teamID != "" {
		vars["filter"] = map[string]any{
			"team": map[string]any{"id": map[string]any{"eq": teamID}},
		}
	}

	var resp struct {
		Issues IssueConnection `json:"issues"`
	}
	if err := c.do(ctx, queryIssueList, vars, &resp); err != nil {
		return nil, err
	}
	return &resp.Issues, nil
}

func (c *Client) CreateIssue(ctx context.Context, input IssueCreateInput) (*Issue, error) {
	var resp struct {
		IssueCreate struct {
			Success bool  `json:"success"`
			Issue   Issue `json:"issue"`
		} `json:"issueCreate"`
	}
	if err := c.do(ctx, mutationIssueCreate, map[string]any{"input": input}, &resp); err != nil {
		return nil, err
	}
	if !resp.IssueCreate.Success {
		return nil, fmt.Errorf("linear: issueCreate returned success=false")
	}
	return &resp.IssueCreate.Issue, nil
}

func (c *Client) UpdateIssue(ctx context.Context, id string, input IssueUpdateInput) (*Issue, error) {
	var resp struct {
		IssueUpdate struct {
			Success bool  `json:"success"`
			Issue   Issue `json:"issue"`
		} `json:"issueUpdate"`
	}
	if err := c.do(ctx, mutationIssueUpdate, map[string]any{"id": id, "input": input}, &resp); err != nil {
		return nil, err
	}
	if !resp.IssueUpdate.Success {
		return nil, fmt.Errorf("linear: issueUpdate returned success=false")
	}
	return &resp.IssueUpdate.Issue, nil
}

func (c *Client) ListTeams(ctx context.Context, first int, after string) (*TeamConnection, error) {
	vars := map[string]any{"first": first}
	if after != "" {
		vars["after"] = after
	}

	var resp struct {
		Teams TeamConnection `json:"teams"`
	}
	if err := c.do(ctx, queryTeamList, vars, &resp); err != nil {
		return nil, err
	}
	return &resp.Teams, nil
}

func (c *Client) ListWorkflowStates(ctx context.Context, first int, after string, teamID string) (*WorkflowStateConnection, error) {
	vars := map[string]any{"first": first}
	if after != "" {
		vars["after"] = after
	}
	if teamID != "" {
		vars["filter"] = map[string]any{
			"team": map[string]any{"id": map[string]any{"eq": teamID}},
		}
	}

	var resp struct {
		WorkflowStates WorkflowStateConnection `json:"workflowStates"`
	}
	if err := c.do(ctx, queryWorkflowStateList, vars, &resp); err != nil {
		return nil, err
	}
	return &resp.WorkflowStates, nil
}
