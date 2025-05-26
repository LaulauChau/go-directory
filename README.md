# Go Phone Directory

Phone directory application with CLI and web interface.

## Team Members

- [Minh-Phuoc, Laurent Chau (minh-phuoc.chau@efrei.net)](minh-phuoc.chau@efrei.net)

## Web Interface

Start the web server:

```bash
# Start web server on default port 8080
go run ./cmd/go-directory/main.go --web

# Start web server on custom port
go run ./cmd/go-directory/main.go --web --port 3000
```

## CLI Commands

```bash
# Add a contact
go run ./cmd/go-directory/main.go --action add --name "John Doe" --tel "1234567890"

# Search for a contact
go run ./cmd/go-directory/main.go --action search --name "John Doe"

# List all contacts
go run ./cmd/go-directory/main.go --action list

# Delete a contact
go run ./cmd/go-directory/main.go --action delete --name "John Doe"

# Edit a contact
go run ./cmd/go-directory/main.go --action edit --name "John Doe" --tel "0987654321"
```

## Flags

### CLI Mode

- `--action`: Required. Values: `add`, `search`, `list`, `delete`, `edit`
- `--name`: Required for all actions except `list`
- `--tel`: Required for `add` and `edit` actions
- `--file`: Optional. Custom JSON file path (default: `contacts.json`)

### Web Mode

- `--web`: Run as web server
- `--port`: Optional. Port for web server (default: `8080`)
- `--file`: Optional. Custom JSON file path (default: `contacts.json`)
