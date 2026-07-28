package repository

import (
	"WailsSnippets/internal/domain"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestJSONSnippetRepositoryEnsureFile(t *testing.T) {
	tests := []struct {
		name                 string
		setupDirectory       func(t *testing.T) string
		existingFile         bool
		initialFileContent   string
		wantFile             bool
		wantFileContent      string
		wantDirectoryMissing bool
		wantErr              bool
	}{
		{
			name: "valid directory creates snippets file",
			setupDirectory: func(t *testing.T) string {
				return t.TempDir()
			},
			wantFile:        true,
			wantFileContent: "[]\n",
		},
		{
			name: "empty directory returns an error",
			setupDirectory: func(t *testing.T) string {
				return ""
			},
			wantErr: true,
		},
		{
			name: "missing directory returns an error without creating it",
			setupDirectory: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "missing-directory")
			},
			wantDirectoryMissing: true,
			wantErr:              true,
		},
		{
			name: "path to a file returns an error",
			setupDirectory: func(t *testing.T) string {
				path := filepath.Join(t.TempDir(), "not-a-directory")
				if err := os.WriteFile(path, []byte("file, not directory"), 0o644); err != nil {
					t.Fatalf("preparing file path: %v", err)
				}
				return path
			},
			wantErr: true,
		},
		{
			name: "existing snippets file keeps its content",
			setupDirectory: func(t *testing.T) string {
				return t.TempDir()
			},
			existingFile:       true,
			initialFileContent: "[{\"id\":\"existing\"}]\n",
			wantFile:           true,
			wantFileContent:    "[{\"id\":\"existing\"}]\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the directory state required by this test case.
			directory := tt.setupDirectory(t)
			// Build the repository and point it at the directory under test.
			r := NewJSONSnippetRepository()
			r.SetDirectory(directory)

			// Keep the original bytes so an existing file can be compared later.
			var initialContent []byte
			if tt.existingFile {
				// Create snippets.json before calling EnsureFile for the existing-file case.
				path := filepath.Join(directory, snippetsFileName)
				if err := os.WriteFile(path, []byte(tt.initialFileContent), 0o644); err != nil {
					t.Fatalf("preparing snippets.json: %v", err)
				}
				// Read the file's initial state for the final non-overwrite assertion.
				var err error
				initialContent, err = os.ReadFile(path)
				if err != nil {
					t.Fatalf("reading initial snippets.json: %v", err)
				}
			}

			// Run the method being tested.
			err := r.EnsureFile()
			// Check whether its error result matches the expectation for this case.
			if (err != nil) != tt.wantErr {
				t.Fatalf("EnsureFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantDirectoryMissing {
				// Verify that EnsureFile did not create a missing configured directory.
				if _, err := os.Stat(directory); !os.IsNotExist(err) {
					t.Fatalf("directory exists or could not be checked; Stat() error = %v", err)
				}
			}

			if !tt.wantFile {
				// Error cases do not expect snippets.json to be present.
				return
			}

			// Read the file created or preserved by EnsureFile.
			path := filepath.Join(directory, snippetsFileName)
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("reading snippets.json: %v", err)
			}
			if len(data) == 0 {
				t.Fatal("empty snippets.json")
			}
			// Check the exact content expected for the successful case.
			if string(data) != tt.wantFileContent {
				t.Errorf("snippets.json content = %q, want %q", data, tt.wantFileContent)
			}
			// Existing files must keep exactly the bytes they had before EnsureFile.
			if tt.existingFile && string(data) != string(initialContent) {
				t.Errorf("existing snippets.json was modified: got %q, want original %q", data, initialContent)
			}
		})
	}
}

func TestJSONSnippetRepositoryList(t *testing.T) {
	tests := []struct {
		name            string
		emptyDirectory  bool
		existingFile    bool
		jsonContent     string
		expectedContent []domain.Snippet
		wantErr         bool
	}{
		{
			name:            "valid file with values on correct order",
			existingFile:    true,
			jsonContent:     `[{"id":"0_item"},{"id":"1_item"}]`,
			expectedContent: []domain.Snippet{{ID: "0_item"}, {ID: "1_item"}},
		},
		{
			name:            "valid empty file",
			existingFile:    true,
			jsonContent:     `[]`,
			expectedContent: []domain.Snippet{},
		},
		{
			name:           "empty directory returns an error",
			emptyDirectory: true,
			wantErr:        true,
		},
		{
			name:    "no file",
			wantErr: true,
		},
		{
			name:         "corrupted file",
			existingFile: true,
			jsonContent:  `a[]`,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the directory state required by this test case.
			directory := t.TempDir()
			if tt.emptyDirectory {
				// List must reject a repository that has no configured directory.
				directory = ""
			}
			// Build the repository and point it at the directory under test.
			r := NewJSONSnippetRepository()
			r.SetDirectory(directory)
			// Write the raw JSON fixture only for cases that require snippets.json.
			if tt.existingFile {
				path := filepath.Join(directory, snippetsFileName)
				if err := os.WriteFile(path, []byte(tt.jsonContent), 0o644); err != nil {
					t.Fatalf("creating snippets.json: %v", err)
				}
			}

			// Run the method under test.
			list, err := r.List()
			// Check whether the returned error matches this case's expectation.
			if (err != nil) != tt.wantErr {
				t.Fatalf("List() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			// Compare the decoded list, including its item order, with the expected value.
			if !reflect.DeepEqual(list, tt.expectedContent) {
				t.Errorf("List() = %#v, want %#v", list, tt.expectedContent)
			}
		})
	}
}

func TestJSONSnippetRepositorySaveList(t *testing.T) {
	tests := []struct {
		name           string
		setupDirectory func(t *testing.T) string
		initialContent string
		snippets       []domain.Snippet
		wantErr        bool
	}{
		{
			name: "saves all snippet fields",
			setupDirectory: func(t *testing.T) string {
				return t.TempDir()
			},
			snippets: []domain.Snippet{{
				ID: "snippet-1", Title: "Example", Language: "go", Code: "fmt.Println()",
				Tags: []string{"go", "example"}, CreatedAt: "2026-07-28T10:00:00Z",
			}},
		},
		{
			name: "saves an empty list",
			setupDirectory: func(t *testing.T) string {
				return t.TempDir()
			},
			snippets: []domain.Snippet{},
		},
		{
			name: "replaces an existing file",
			setupDirectory: func(t *testing.T) string {
				return t.TempDir()
			},
			initialContent: `[{"id":"old"}]`,
			snippets:       []domain.Snippet{{ID: "new", Title: "New", Code: "new()"}},
		},
		{
			name: "empty directory returns an error",
			setupDirectory: func(t *testing.T) string {
				return ""
			},
			wantErr: true,
		},
		{
			name: "missing directory returns an error",
			setupDirectory: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "missing-directory")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the directory state required by this test case.
			directory := tt.setupDirectory(t)
			r := NewJSONSnippetRepository()
			r.SetDirectory(directory)

			if tt.initialContent != "" {
				// Create an old file to verify that SaveList overwrites it.
				path := filepath.Join(directory, snippetsFileName)
				if err := os.WriteFile(path, []byte(tt.initialContent), 0o644); err != nil {
					t.Fatalf("creating initial snippets.json: %v", err)
				}
			}

			// Run the method under test.
			err := r.SaveList(tt.snippets)
			// Check whether the returned error matches this case's expectation.
			if (err != nil) != tt.wantErr {
				t.Fatalf("SaveList() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			// Read and decode the persisted JSON independently from List().
			data, err := os.ReadFile(filepath.Join(directory, snippetsFileName))
			if err != nil {
				t.Fatalf("reading snippets.json: %v", err)
			}
			var persisted []domain.Snippet
			if err := json.Unmarshal(data, &persisted); err != nil {
				t.Fatalf("decoding snippets.json: %v", err)
			}
			// Compare every persisted field with the input passed to SaveList.
			if !reflect.DeepEqual(persisted, tt.snippets) {
				t.Errorf("persisted snippets = %#v, want %#v", persisted, tt.snippets)
			}
		})
	}
}
