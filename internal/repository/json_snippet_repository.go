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
type JSONSnippetRepository struct {
	directory string
}

func NewJSONSnippetRepository() *JSONSnippetRepository {
	return &JSONSnippetRepository{}
}

func (r *JSONSnippetRepository) SetDirectory(directory string) {
	r.directory = directory
}

func (r *JSONSnippetRepository) Directory() string {
	return r.directory
}

// EnsureFile creates snippets.json only when the configured folder exists and
// the file has not been created yet.
func (r *JSONSnippetRepository) EnsureFile() error {
	if r.directory == "" {
		return errors.New("No route selected; select folder to save snippets")
	}
	info, err := os.Stat(r.directory)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.New("Route is not a valid folder")
	}

	path := filepath.Join(r.directory, snippetsFileName)
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
func (r *JSONSnippetRepository) List() ([]domain.Snippet, error) {
	if r.directory == "" {
		return nil, errors.New("No route selected; select folder to save snippets")
	}

	file, err := os.Open(filepath.Join(r.directory, snippetsFileName))
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

// Save file, rewritting if exists
func (r *JSONSnippetRepository) SaveList(snippets []domain.Snippet) error {
	file, err := os.Create(filepath.Join(r.directory, snippetsFileName))
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(snippets)
}
