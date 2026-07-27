package repository

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"WailsSnippets/internal/domain"
)

const snippetsFileName = "snippets.json"

// JSONSnippetRepository contains the basic JSON-file operations for snippets.
type JSONSnippetRepository struct{}

func NewJSONSnippetRepository() *JSONSnippetRepository {
	return &JSONSnippetRepository{}
}

// EnsureFile creates snippets.json only when the configured folder exists and
// the file has not been created yet.
func (r *JSONSnippetRepository) EnsureFile(directory string) error {
	if directory == "" {
		return errors.New("No route selected; select folder to save snippets")
	}
	info, err := os.Stat(directory)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.New("Route is not a valid folder")
	}

	path := filepath.Join(directory, snippetsFileName)
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, []byte("[]\n"), 0o644); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

// List only reads an existing snippets.json file.
func (r *JSONSnippetRepository) List(directory string) ([]domain.Snippet, error) {
	if directory == "" {
		return nil, errors.New("No route selected; select folder to save snippets")
	}

	file, err := os.Open(filepath.Join(directory, snippetsFileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var snippets []domain.Snippet
	if err := json.NewDecoder(file).Decode(&snippets); err != nil {
		return nil, err
	}
	return snippets, nil
}

// Save is the complementary operation for when CREATE/UPDATE/DELETE are moved here.
func (r *JSONSnippetRepository) Save(directory string, snippets []domain.Snippet) error {
	file, err := os.Create(filepath.Join(directory, snippetsFileName))
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(snippets)
}
