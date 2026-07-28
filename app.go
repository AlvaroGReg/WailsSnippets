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
	ctx              context.Context
	config           domain.AppConfig
	configRepository *repository.JSONConfigRepository
	snippets         *service.SnippetService
}

// NewApp creates a new App application struct
func NewApp() *App {
	configRepository := repository.NewJSONConfigRepository("WailsSnippets")
	config, err := configRepository.Load()
	if err != nil {
		log.Printf("unable to load snippets configuration: %v", err)
	}

	snippets := service.NewSnippetService()
	snippets.SetSnippetsDirectory(config.SnippetsDirectory)
	return &App{config: config, configRepository: configRepository, snippets: snippets}
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
	directory, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Chose snippets folder",
	})
	if err != nil {
		return "", err
	}
	if directory == "" { // The native dialog was cancelled.
		return a.snippets.SnippetsDirectory(), nil
	}

	previousDirectory := a.snippets.SnippetsDirectory()
	a.snippets.SetSnippetsDirectory(directory)
	if err := a.snippets.EnsureSnippetsFile(); err != nil {
		a.snippets.SetSnippetsDirectory(previousDirectory)
		return "", err
	}
	previousConfig := a.config
	a.config.SnippetsDirectory = directory
	if err := a.configRepository.SaveConfig(a.config); err != nil {
		a.config = previousConfig
		a.snippets.SetSnippetsDirectory(previousDirectory)
		return "", err
	}
	return directory, nil
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
