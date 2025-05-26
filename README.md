# Go Phone Directory

Simple CLI phone directory application.

## Commands

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

- `--action`: Required. Values: `add`, `search`, `list`, `delete`, `edit`
- `--name`: Required for all actions except `list`
- `--tel`: Required for `add` and `edit` actions
- `--file`: Optional. Custom JSON file path (default: `contacts.json`)
