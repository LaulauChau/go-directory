package storage

import "github.com/LaulauChau/go-directory/internal/domain"

type Storage interface {
	Load() ([]domain.Contact, error)
	Save(contacts []domain.Contact) error
}
