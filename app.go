package main

import (
	"context"
	"log"

	"SnippetsDome/internal/domain"
	"SnippetsDome/internal/repository"
	"SnippetsDome/internal/service"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	snippets *service.SnippetService
}

// NewApp creates a new App application struct
func NewApp() *App {
	configRepository := repository.NewJSONConfigRepository("SnippetsDome")
	config, err := configRepository.Load()
	if err != nil {
		log.Printf("unable to load snippets configuration: %v", err)
	}

	snippets := service.NewSnippetService(config, configRepository)
	return &App{snippets: snippets}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	if a.snippets.SnippetsFilePath() == "" {
		return
	}
	if err := a.snippets.EnsureSnippetsFile(); err != nil {
		log.Printf("unable to create snippets file: %v", err)
	}
}

func (a *App) GetSnippets() ([]domain.Snippet, error) {
	return a.snippets.List()
}

// PickExistingSnippetsFile is a thin native-dialog bridge used by React.
func (a *App) PickExistingSnippetsFile() (string, error) {
	return runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:   "Choose snippets file",
		Filters: []runtime.FileFilter{{DisplayName: "JSON files", Pattern: "*.json"}},
	})
}

// CreateSnippetsFile is a thin native-dialog bridge used by React.
func (a *App) CreateSnippetsFile() (string, error) {
	return runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "Create snippets file",
		DefaultFilename: "snippets.json",
		Filters:         []runtime.FileFilter{{DisplayName: "JSON files", Pattern: "*.json"}},
	})
}

func (a *App) SetSnippetsStoragePath(filePath string) (string, error) {
	return a.snippets.SetSnippetsFile(filePath)
}

func (a *App) GetSnippetsStoragePath() string {
	return a.snippets.SnippetsFilePath()
}

func (a *App) CreateSnippet(input domain.CreateSnippetInput) (domain.Snippet, error) {
	return a.snippets.CreateSnippet(input)
}

func (a *App) UpdateSnippet(snippet domain.Snippet) (domain.Snippet, error) {
	return a.snippets.UpdateSnippet(snippet)
}

func (a *App) DeleteSnippet(id string) error {
	return a.snippets.DeleteSnippet(id)
}
