package linear

import (
	"encoding/json"
	"time"
)

type Issue struct {
	ID          string        `json:"id"`
	Identifier  string        `json:"identifier"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Priority    int           `json:"priority"`
	State       WorkflowState `json:"state"`
	Team        Team          `json:"team"`
	Assignee    *User         `json:"assignee"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type WorkflowState struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Color string `json:"color"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type IssueCreateInput struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	TeamID      string  `json:"teamId"`
	StateID     *string `json:"stateId,omitempty"`
	AssigneeID  *string `json:"assigneeId,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
}

type IssueUpdateInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	StateID     *string `json:"stateId,omitempty"`
	AssigneeID  *string `json:"assigneeId,omitempty"`
	Priority    *int    `json:"priority,omitempty"`
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	EndCursor   string `json:"endCursor"`
}

type IssueConnection struct {
	Nodes    []Issue  `json:"nodes"`
	PageInfo PageInfo `json:"pageInfo"`
}

type TeamConnection struct {
	Nodes    []Team   `json:"nodes"`
	PageInfo PageInfo `json:"pageInfo"`
}

type WorkflowStateConnection struct {
	Nodes    []WorkflowState `json:"nodes"`
	PageInfo PageInfo        `json:"pageInfo"`
}

type graphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

type graphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []graphQLError  `json:"errors,omitempty"`
}

type graphQLError struct {
	Message    string         `json:"message"`
	Extensions map[string]any `json:"extensions,omitempty"`
}
