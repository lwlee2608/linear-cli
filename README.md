# linear-cli

A command-line interface for [Linear](https://linear.app).

## Installation

```bash
make build
```

The binary will be at `bin/linear-cli`.

## Configuration

Set your Linear API key as an environment variable:

```bash
export LINEAR_API_KEY=lin_api_xxxxxxxxxxxxx
```

You can generate an API key from Linear under **Settings > Security & access > API**.

## Usage

### Search issues

```bash
linear-cli issue search <keywords>
```

Options:
- `--limit` — maximum number of results (default: 20)

Example:

```bash
linear-cli issue search "login bug"
linear-cli issue search auth --limit 5
```

### Get an issue

```bash
linear-cli issue get <id>
```

Example:

```bash
linear-cli issue get ENG-123
```
