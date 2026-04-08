# linear-cli

A command-line interface for [Linear](https://linear.app).

## Installation

### Quick install

```bash
curl -fsSL https://raw.githubusercontent.com/lwlee2608/linear-cli/main/install.sh | bash
```

### From source

```bash
make build
```

The binary will be at `bin/linear`.

To install to `~/.local/bin`:

```bash
make install
```

## Configuration

Set your Linear API key as an environment variable:

```bash
export LINEAR_API_KEY=lin_api_xxxxxxxxxxxxx
```

You can generate an API key from Linear under **Settings > Security & access > API**.

## Usage

### Search issues

```bash
linear issue search <keywords>
```

Options:
- `--limit` — maximum number of results (default: 20)

Example:

```bash
linear issue search "login bug"
linear issue search auth --limit 5
```

### Get an issue

```bash
linear issue get <id>
```

Example:

```bash
linear issue get ENG-123
```
