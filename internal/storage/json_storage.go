package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/LaulauChau/go-directory/internal/domain"
)

type JSONStorage struct {
	filePath string
}

func NewJSONStorage(filePath string) *JSONStorage {
	return &JSONStorage{
		filePath: filePath,
	}
}

func (s *JSONStorage) Load() ([]domain.Contact, error) {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return make([]domain.Contact, 0), nil
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(data) == 0 {
		return make([]domain.Contact, 0), nil
	}

	var contacts []domain.Contact
	err = json.Unmarshal(data, &contacts)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return contacts, nil
}

func (s *JSONStorage) Save(contacts []domain.Contact) error {
	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = os.WriteFile(s.filePath, data, 0600)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
