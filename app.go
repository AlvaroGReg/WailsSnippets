package main

import (
	"context"
	"log"

	"WailsSnippets/internal/domain"
	"WailsSnippets/internal/repository"
	"WailsSnippets/internal/service"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	snippets *service.SnippetService
}

// NewApp creates a new App application struct
func NewApp() *App {
	configRepository := repository.NewJSONConfigRepository("WailsSnippets")
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
	if a.snippets.SnippetsDirectory() == "" {
		return
	}
	if err := a.snippets.EnsureSnippetsFile(); err != nil {
		log.Printf("unable to create snippets file: %v", err)
	}
}

func (a *App) GetSnippets() ([]domain.Snippet, error) {
	return a.snippets.List()
}

func (a *App) SelectSnippetsDirectory() (string, error) {
	// TODO: Getting route is reacts work, removing the need of context and runtime imports here.
	// move it while making the front and swap directory: string to prop instead of return
	directory, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Chose snippets folder",
	})
	if err != nil {
		return "", err
	}
	return a.snippets.SelectSnippetsDirectory(directory)
}

func (a *App) GetSnippetsStoragePath() string {
	return a.snippets.SnippetsDirectory()
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
