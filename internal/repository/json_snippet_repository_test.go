package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"SnippetsDome/internal/domain"
)

func TestJSONSnippetRepositoryEnsureFile(t *testing.T) {
	tests := []struct {
		name          string
		setupFilePath func(t *testing.T) string
		initial       string
		wantErr       bool
	}{
		{
			name: "creates a new file at the selected path",
			setupFilePath: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "my-snippets.json")
			},
		},
		{
			name: "keeps an existing file unchanged",
			setupFilePath: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "existing.json")
			},
			initial: `[{"id":"existing"}]`,
		},
		{
			name: "rejects a missing parent directory",
			setupFilePath: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "missing", "snippets.json")
			},
			wantErr: true,
		},
		{
			name: "rejects a directory path",
			setupFilePath: func(t *testing.T) string {
				return t.TempDir()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := tt.setupFilePath(t)
			if tt.initial != "" {
				if err := os.WriteFile(filePath, []byte(tt.initial), 0o644); err != nil {
					t.Fatalf("preparing file: %v", err)
				}
			}

			repository := NewJSONSnippetRepository()
			repository.SetFilePath(filePath)
			err := repository.EnsureFile()
			if (err != nil) != tt.wantErr {
				t.Fatalf("EnsureFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			content, err := os.ReadFile(filePath)
			if err != nil {
				t.Fatalf("reading selected file: %v", err)
			}
			want := tt.initial
			if want == "" {
				want = "[]\n"
			}
			if string(content) != want {
				t.Errorf("file content = %q, want %q", content, want)
			}
		})
	}
}

func TestJSONSnippetRepositoryList(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "library.json")
	if err := os.WriteFile(filePath, []byte(`[{"id":"first"},{"id":"second"}]`), 0o644); err != nil {
		t.Fatalf("preparing file: %v", err)
	}
	repository := NewJSONSnippetRepository()
	repository.SetFilePath(filePath)

	snippets, err := repository.List()
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	want := []domain.Snippet{{ID: "first"}, {ID: "second"}}
	if !reflect.DeepEqual(snippets, want) {
		t.Errorf("List() = %#v, want %#v", snippets, want)
	}
}

func TestJSONSnippetRepositorySaveList(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "library.json")
	repository := NewJSONSnippetRepository()
	repository.SetFilePath(filePath)
	snippets := []domain.Snippet{{ID: "snippet-1", Title: "Example", Code: "fmt.Println()"}}

	if err := repository.SaveList(snippets); err != nil {
		t.Fatalf("SaveList() error = %v", err)
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("reading saved file: %v", err)
	}
	var persisted []domain.Snippet
	if err := json.Unmarshal(data, &persisted); err != nil {
		t.Fatalf("decoding saved file: %v", err)
	}
	if !reflect.DeepEqual(persisted, snippets) {
		t.Errorf("saved snippets = %#v, want %#v", persisted, snippets)
	}
}
