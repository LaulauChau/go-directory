package service

import (
	"fmt"
	"strings"

	"github.com/LaulauChau/go-directory/internal/domain"
	"github.com/LaulauChau/go-directory/internal/storage"
)

type Directory struct {
	storage  storage.Storage
	contacts []domain.Contact
}

func NewDirectory(storage storage.Storage) (*Directory, error) {
	dir := &Directory{
		storage: storage,
	}

	contacts, err := storage.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load contacts: %w", err)
	}

	dir.contacts = contacts
	return dir, nil
}

func (d *Directory) AddContact(name, phone string) error {
	name = strings.TrimSpace(name)
	phone = strings.TrimSpace(phone)

	if d.contactExists(name) {
		return fmt.Errorf("contact with name '%s' already exists", name)
	}

	contact := domain.NewContact(name, phone)
	d.contacts = append(d.contacts, contact)
	return d.storage.Save(d.contacts)
}

func (d *Directory) DeleteContact(name string) error {
	name = strings.TrimSpace(name)
	for i, contact := range d.contacts {
		if strings.EqualFold(contact.Name, name) {
			d.contacts = append(d.contacts[:i], d.contacts[i+1:]...)
			return d.storage.Save(d.contacts)
		}
	}
	return fmt.Errorf("contact with name '%s' not found", name)
}

func (d *Directory) EditContact(name, newPhone string) error {
	name = strings.TrimSpace(name)
	newPhone = strings.TrimSpace(newPhone)

	for i, contact := range d.contacts {
		if strings.EqualFold(contact.Name, name) {
			d.contacts[i].Phone = newPhone
			return d.storage.Save(d.contacts)
		}
	}
	return fmt.Errorf("contact with name '%s' not found", name)
}

func (d *Directory) SearchContact(name string) (*domain.Contact, error) {
	name = strings.TrimSpace(name)
	for _, contact := range d.contacts {
		if strings.Contains(strings.ToLower(contact.Name), strings.ToLower(name)) {
			return &contact, nil
		}
	}
	return nil, fmt.Errorf("contact with name '%s' not found", name)
}

func (d *Directory) SearchContacts(name string) []domain.Contact {
	name = strings.TrimSpace(name)
	var matches []domain.Contact

	for _, contact := range d.contacts {
		if strings.Contains(strings.ToLower(contact.Name), strings.ToLower(name)) {
			matches = append(matches, contact)
		}
	}

	return matches
}

func (d *Directory) ListContacts() []domain.Contact {
	return d.contacts
}

func (d *Directory) contactExists(name string) bool {
	name = strings.TrimSpace(name)
	for _, contact := range d.contacts {
		if strings.EqualFold(contact.Name, name) {
			return true
		}
	}
	return false
}
