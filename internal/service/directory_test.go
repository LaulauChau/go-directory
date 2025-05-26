package service

import (
	"testing"

	"github.com/LaulauChau/go-directory/internal/domain"
)

type mockStorage struct {
	contacts []domain.Contact
}

func (m *mockStorage) Load() ([]domain.Contact, error) {
	return m.contacts, nil
}

func (m *mockStorage) Save(contacts []domain.Contact) error {
	m.contacts = contacts
	return nil
}

func newMockStorage() *mockStorage {
	return &mockStorage{
		contacts: make([]domain.Contact, 0),
	}
}

func TestNewDirectory(t *testing.T) {
	storage := newMockStorage()
	dir, err := NewDirectory(storage)

	if err != nil {
		t.Errorf("Expected no error creating directory, got %v", err)
	}

	if len(dir.contacts) != 0 {
		t.Errorf("Expected empty contacts slice, got %d contacts", len(dir.contacts))
	}
}

func TestAddContact(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	err := dir.AddContact("John Doe", "1234567890")
	if err != nil {
		t.Errorf("Expected no error adding contact, got %v", err)
	}

	if len(dir.contacts) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(dir.contacts))
	}

	contact := dir.contacts[0]
	if contact.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", contact.Name)
	}

	if contact.Phone != "1234567890" {
		t.Errorf("Expected phone '1234567890', got '%s'", contact.Phone)
	}
}

func TestAddContact_Duplicate(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	err := dir.AddContact("John Doe", "1234567890")
	if err != nil {
		t.Fatalf("Failed to add first contact: %v", err)
	}

	err = dir.AddContact("John Doe", "0987654321")
	if err == nil {
		t.Error("Expected error when adding duplicate contact")
	}

	if len(dir.contacts) != 1 {
		t.Errorf("Expected 1 contact after duplicate attempt, got %d", len(dir.contacts))
	}
}

func TestAddContact_CaseInsensitive(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	err := dir.AddContact("John Doe", "1234567890")
	if err != nil {
		t.Fatalf("Failed to add first contact: %v", err)
	}

	err = dir.AddContact("john doe", "0987654321")
	if err == nil {
		t.Error("Expected error when adding duplicate contact with different case")
	}
}

func TestDeleteContact(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	err := dir.AddContact("John Doe", "1234567890")
	if err != nil {
		t.Fatalf("Failed to add first contact: %v", err)
	}

	err = dir.AddContact("Jane Smith", "0987654321")
	if err != nil {
		t.Fatalf("Failed to add second contact: %v", err)
	}

	err = dir.DeleteContact("John Doe")
	if err != nil {
		t.Errorf("Expected no error deleting contact, got %v", err)
	}

	if len(dir.contacts) != 1 {
		t.Errorf("Expected 1 contact after deletion, got %d", len(dir.contacts))
	}

	if dir.contacts[0].Name != "Jane Smith" {
		t.Errorf("Expected remaining contact to be 'Jane Smith', got '%s'", dir.contacts[0].Name)
	}
}

func TestDeleteContact_NotFound(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	err := dir.DeleteContact("Non Existent")
	if err == nil {
		t.Error("Expected error when deleting non-existent contact")
	}
}

func TestEditContact(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	err := dir.AddContact("John Doe", "1234567890")
	if err != nil {
		t.Fatalf("Failed to add contact: %v", err)
	}

	err = dir.EditContact("John Doe", "5555555555")
	if err != nil {
		t.Errorf("Expected no error editing contact, got %v", err)
	}

	if dir.contacts[0].Phone != "5555555555" {
		t.Errorf("Expected phone to be '5555555555', got '%s'", dir.contacts[0].Phone)
	}
}

func TestEditContact_NotFound(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	err := dir.EditContact("Non Existent", "1234567890")
	if err == nil {
		t.Error("Expected error when editing non-existent contact")
	}
}

func TestSearchContact(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	err := dir.AddContact("John Doe", "1234567890")
	if err != nil {
		t.Fatalf("Failed to add first contact: %v", err)
	}

	err = dir.AddContact("Jane Smith", "0987654321")
	if err != nil {
		t.Fatalf("Failed to add second contact: %v", err)
	}

	contact, err := dir.SearchContact("John Doe")
	if err != nil {
		t.Errorf("Expected no error searching contact, got %v", err)
	}

	if contact.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", contact.Name)
	}

	if contact.Phone != "1234567890" {
		t.Errorf("Expected phone '1234567890', got '%s'", contact.Phone)
	}
}

func TestSearchContact_NotFound(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	_, err := dir.SearchContact("Non Existent")
	if err == nil {
		t.Error("Expected error when searching for non-existent contact")
	}
}

func TestListContacts(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	contacts := dir.ListContacts()
	if len(contacts) != 0 {
		t.Errorf("Expected empty list, got %d contacts", len(contacts))
	}

	err := dir.AddContact("John Doe", "1234567890")
	if err != nil {
		t.Fatalf("Failed to add first contact: %v", err)
	}

	err = dir.AddContact("Jane Smith", "0987654321")
	if err != nil {
		t.Fatalf("Failed to add second contact: %v", err)
	}

	contacts = dir.ListContacts()
	if len(contacts) != 2 {
		t.Errorf("Expected 2 contacts, got %d", len(contacts))
	}
}

func TestTrimSpaces(t *testing.T) {
	storage := newMockStorage()
	dir, _ := NewDirectory(storage)

	err := dir.AddContact("  John Doe  ", "  1234567890  ")
	if err != nil {
		t.Errorf("Expected no error adding contact with spaces, got %v", err)
	}

	contact := dir.contacts[0]
	if contact.Name != "John Doe" {
		t.Errorf("Expected trimmed name 'John Doe', got '%s'", contact.Name)
	}

	if contact.Phone != "1234567890" {
		t.Errorf("Expected trimmed phone '1234567890', got '%s'", contact.Phone)
	}
}
