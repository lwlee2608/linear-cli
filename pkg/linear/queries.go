package linear

const queryIssueGet = `
query IssueGet($id: String!) {
  issue(id: $id) {
    id identifier title description priority
    createdAt updatedAt
    state { id name type color }
    team  { id name key }
    assignee { id name email }
    project { id name }
    labels { nodes { id name color } }
  }
}`

const queryIssueList = `
query IssueList($first: Int, $after: String, $filter: IssueFilter) {
  issues(first: $first, after: $after, filter: $filter) {
    nodes {
      id identifier title priority
      createdAt updatedAt
      state { id name type color }
      team  { id name key }
      assignee { id name email }
    }
    pageInfo { hasNextPage endCursor }
  }
}`

const mutationIssueCreate = `
mutation IssueCreate($input: IssueCreateInput!) {
  issueCreate(input: $input) {
    success
    issue {
      id identifier title description priority
      createdAt updatedAt
      state { id name type color }
      team  { id name key }
    }
  }
}`

const mutationIssueUpdate = `
mutation IssueUpdate($id: String!, $input: IssueUpdateInput!) {
  issueUpdate(id: $id, input: $input) {
    success
    issue {
      id identifier title description priority
      createdAt updatedAt
      state { id name type color }
      team  { id name key }
    }
  }
}`

const queryTeamList = `
query TeamList($first: Int, $after: String) {
  teams(first: $first, after: $after) {
    nodes { id name key }
    pageInfo { hasNextPage endCursor }
  }
}`

const queryWorkflowStateList = `
query WorkflowStateList($first: Int, $after: String, $filter: WorkflowStateFilter) {
  workflowStates(first: $first, after: $after, filter: $filter) {
    nodes { id name type color }
    pageInfo { hasNextPage endCursor }
  }
}`

const queryIssueSearch = `
query IssueSearch($query: String!, $first: Int, $after: String) {
  searchIssues(query: $query, first: $first, after: $after) {
    nodes {
      id identifier title priority
      createdAt updatedAt
      state { id name type color }
      team  { id name key }
      assignee { id name email }
    }
    pageInfo { hasNextPage endCursor }
  }
}`
