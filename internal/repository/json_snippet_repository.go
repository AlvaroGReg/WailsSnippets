package repository

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"SnippetsDome/internal/domain"
)

// JSONSnippetRepository contains the basic JSON-file operations for snippets.
type JSONSnippetRepository struct {
	filePath string
}

func NewJSONSnippetRepository() *JSONSnippetRepository {
	return &JSONSnippetRepository{}
}

func (r *JSONSnippetRepository) SetFilePath(filePath string) {
	r.filePath = filePath
}

func (r *JSONSnippetRepository) FilePath() string {
	return r.filePath
}

// EnsureFile creates the configured file only when its parent directory exists.
func (r *JSONSnippetRepository) EnsureFile() error {
	if r.filePath == "" {
		return errors.New("no snippets file selected")
	}

	info, err := os.Stat(r.filePath)
	if os.IsNotExist(err) {
		parent, parentErr := os.Stat(filepath.Dir(r.filePath))
		if parentErr != nil {
			return parentErr
		}
		if !parent.IsDir() {
			return errors.New("snippets file parent is not a directory")
		}
		if err := os.WriteFile(r.filePath, []byte("[]\n"), 0o644); err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("snippets path is not a file")
	}
	return nil
}

// List only reads an existing snippets file.
func (r *JSONSnippetRepository) List() ([]domain.Snippet, error) {
	if r.filePath == "" {
		return nil, errors.New("no snippets file selected")
	}

	file, err := os.Open(r.filePath)
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
	if r.filePath == "" {
		return errors.New("no snippets file selected")
	}

	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(snippets)
}
