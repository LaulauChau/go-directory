package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/LaulauChau/go-directory/internal/domain"
)

func createTempFile(t *testing.T) string {
	tmpDir := t.TempDir()
	return filepath.Join(tmpDir, "test_contacts.json")
}

func TestNewJSONStorage(t *testing.T) {
	filePath := "test.json"
	storage := NewJSONStorage(filePath)

	if storage.filePath != filePath {
		t.Errorf("Expected FilePath to be %s, got %s", filePath, storage.filePath)
	}
}

func TestLoad_NonExistentFile(t *testing.T) {
	filePath := createTempFile(t)
	storage := NewJSONStorage(filePath)

	contacts, err := storage.Load()
	if err != nil {
		t.Errorf("Expected no error for non-existent file, got %v", err)
	}

	if len(contacts) != 0 {
		t.Errorf("Expected empty contacts for non-existent file, got %d contacts", len(contacts))
	}
}

func TestLoad_EmptyFile(t *testing.T) {
	filePath := createTempFile(t)

	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("Failed to close test file: %v", err)
	}

	storage := NewJSONStorage(filePath)
	contacts, err := storage.Load()
	if err != nil {
		t.Errorf("Expected no error for empty file, got %v", err)
	}

	if len(contacts) != 0 {
		t.Errorf("Expected empty contacts for empty file, got %d contacts", len(contacts))
	}
}

func TestSaveAndLoad(t *testing.T) {
	filePath := createTempFile(t)
	storage := NewJSONStorage(filePath)

	contacts := []domain.Contact{
		domain.NewContact("John Doe", "1234567890"),
		domain.NewContact("Jane Smith", "0987654321"),
	}

	err := storage.Save(contacts)
	if err != nil {
		t.Fatalf("Failed to save contacts: %v", err)
	}

	loadedContacts, err := storage.Load()
	if err != nil {
		t.Fatalf("Failed to load contacts: %v", err)
	}

	if len(loadedContacts) != len(contacts) {
		t.Errorf("Expected %d contacts, got %d", len(contacts), len(loadedContacts))
	}

	for i, contact := range contacts {
		if contact.Name != loadedContacts[i].Name {
			t.Errorf("Expected name %s, got %s", contact.Name, loadedContacts[i].Name)
		}
		if contact.Phone != loadedContacts[i].Phone {
			t.Errorf("Expected phone %s, got %s", contact.Phone, loadedContacts[i].Phone)
		}
	}
}
