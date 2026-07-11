---
name: linear
description: Use when the user wants to interact with Linear.app — reading, searching, or commenting on issues/tickets.
user-invocable: true
argument-hint: "[get <id> | search <query> | comment <id> <body>]"
---

# Linear Issue Management

Interact with Linear.app issues using the `linear` CLI.

## Prerequisites

- **`linear` CLI** must be installed (`linear --version`). If missing, install with:
  ```bash
  curl -fsSL https://raw.githubusercontent.com/lwlee2608/linear-cli/main/install.sh | bash
  ```
- **`LINEAR_API_KEY`** must be set. Check with `echo $LINEAR_API_KEY`. If unset, tell the user to export it (`export LINEAR_API_KEY="lin_api_..."`).

## Commands

### Get an issue

```bash
linear issue get ENG-123
linear issue get DEV-12 --download-images ./DEV-12-images
```

Returns: identifier, title, state, team, project, created date, priority, assignee, labels, description.

Use `--download-images <directory>` to download images embedded in the description. Existing files are not overwritten.

### Search issues

```bash
linear issue search "login bug"
linear issue search "login bug" --limit 50
```

Returns a table: ID, TITLE, STATE, ASSIGNEE. Default limit is 20.

### Add a comment

```bash
linear issue comment ENG-123 "comment body"
```

Accepts the issue identifier (e.g. `ENG-123`) or UUID. Prints the comment author and body on success.
