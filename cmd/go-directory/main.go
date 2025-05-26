package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/LaulauChau/go-directory/api"
	"github.com/LaulauChau/go-directory/internal/service"
	"github.com/LaulauChau/go-directory/internal/storage"
)

const defaultDataFile = "contacts.json"

func main() {
	var (
		action  = flag.String("action", "", "Action to perform: add, delete, edit, search, list")
		name    = flag.String("name", "", "Contact name (firstname lastname)")
		tel     = flag.String("tel", "", "Phone number")
		file    = flag.String("file", defaultDataFile, "JSON file to store contacts")
		webMode = flag.Bool("web", false, "Run as web server")
		port    = flag.String("port", "8080", "Port for web server")
	)
	flag.Parse()

	if *webMode {
		startWebServer(*file, *port)
		return
	}

	if *action == "" {
		fmt.Println("Error: --action flag is required")
		printUsage()
		os.Exit(1)
	}

	dataFile, err := filepath.Abs(*file)
	if err != nil {
		fmt.Printf("Error: invalid file path: %v\n", err)
		os.Exit(1)
	}

	store := storage.NewJSONStorage(dataFile)
	directory, err := service.NewDirectory(store)
	if err != nil {
		fmt.Printf("Error initializing directory: %v\n", err)
		os.Exit(1)
	}

	switch *action {
	case "add":
		handleAdd(directory, *name, *tel)
	case "delete":
		handleDelete(directory, *name)
	case "edit":
		handleEdit(directory, *name, *tel)
	case "search":
		handleSearch(directory, *name)
	case "list":
		handleList(directory)
	default:
		fmt.Printf("Error: unknown action '%s'\n", *action)
		printUsage()
		os.Exit(1)
	}
}

func handleAdd(directory *service.Directory, name, phone string) {
	if name == "" || phone == "" {
		fmt.Println("Error: both --name and --tel are required for add action")
		os.Exit(1)
	}

	err := directory.AddContact(name, phone)
	if err != nil {
		fmt.Printf("Error adding contact: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Contact '%s' added successfully\n", name)
}

func handleDelete(directory *service.Directory, name string) {
	if name == "" {
		fmt.Println("Error: --name is required for delete action")
		os.Exit(1)
	}

	err := directory.DeleteContact(name)
	if err != nil {
		fmt.Printf("Error deleting contact: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Contact '%s' deleted successfully\n", name)
}

func handleEdit(directory *service.Directory, name, phone string) {
	if name == "" || phone == "" {
		fmt.Println("Error: both --name and --tel are required for edit action")
		os.Exit(1)
	}

	err := directory.EditContact(name, phone)
	if err != nil {
		fmt.Printf("Error editing contact: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Contact '%s' updated successfully\n", name)
}

func handleSearch(directory *service.Directory, name string) {
	if name == "" {
		fmt.Println("Error: --name is required for search action")
		os.Exit(1)
	}

	contact, err := directory.SearchContact(name)
	if err != nil {
		fmt.Printf("Error searching contact: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found contact:\nName: %s\nPhone: %s\n", contact.Name, contact.Phone)
}

func handleList(directory *service.Directory) {
	contacts := directory.ListContacts()
	if len(contacts) == 0 {
		fmt.Println("No contacts found")
		return
	}

	fmt.Printf("Found %d contact(s):\n", len(contacts))
	fmt.Println("-------------------")
	for _, contact := range contacts {
		fmt.Printf("Name: %s\nPhone: %s\n-------------------\n", contact.Name, contact.Phone)
	}
}

func startWebServer(file, port string) {
	dataFile, err := filepath.Abs(file)
	if err != nil {
		fmt.Printf("Error: invalid file path: %v\n", err)
		os.Exit(1)
	}

	store := storage.NewJSONStorage(dataFile)
	directory, err := service.NewDirectory(store)
	if err != nil {
		fmt.Printf("Error initializing directory: %v\n", err)
		os.Exit(1)
	}

	server := api.NewServer(directory, port)
	server.StartWithGracefulShutdown()
}

func printUsage() {
	fmt.Println("\nUsage:")
	fmt.Println("  go run ./cmd/go-directory/main.go --action <action> [options]")
	fmt.Println("  go run ./cmd/go-directory/main.go --web [--port <port>]")
	fmt.Println("\nActions:")
	fmt.Println("  add     Add a new contact (requires --name and --tel)")
	fmt.Println("  delete  Delete a contact (requires --name)")
	fmt.Println("  edit    Edit a contact's phone number (requires --name and --tel)")
	fmt.Println("  search  Search for a contact (requires --name)")
	fmt.Println("  list    List all contacts")
	fmt.Println("\nOptions:")
	fmt.Println("  --name    Contact name (firstname lastname)")
	fmt.Println("  --tel     Phone number")
	fmt.Println("  --file    JSON file to store contacts (default: contacts.json)")
	fmt.Println("  --web     Run as web server")
	fmt.Println("  --port    Port for web server (default: 8080)")
	fmt.Println("\nExamples:")
	fmt.Println("  go run ./cmd/go-directory/main.go --action add --name \"Charlie Brown\" --tel \"0000000000\"")
	fmt.Println("  go run ./cmd/go-directory/main.go --action search --name \"Alice\"")
	fmt.Println("  go run ./cmd/go-directory/main.go --action list")
	fmt.Println("  go run ./cmd/go-directory/main.go --web")
	fmt.Println("  go run ./cmd/go-directory/main.go --web --port 3000")
}
