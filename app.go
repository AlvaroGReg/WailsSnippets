package main

import (
	"context"

	"WailsSnippets/internal/domain"
	"WailsSnippets/internal/service"
)

// App struct
type App struct {
	ctx      context.Context
	snippets *service.SnippetService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		snippets: service.NewSnippetService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetSnippets() []domain.Snippet {
	return a.snippets.List()
}

func (a *App) CreateSnippet(input domain.CreateSnippetInput) (domain.Snippet, error) {
	return a.snippets.Create(input)
}
