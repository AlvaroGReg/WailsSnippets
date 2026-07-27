package service

import (
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"WailsSnippets/internal/domain"
)

type SnippetService struct {
	mu       sync.RWMutex
	snippets []domain.Snippet
}

func NewSnippetService() *SnippetService {
	return &SnippetService{
		snippets: []domain.Snippet{},
	}
}

// GET
func (s *SnippetService) List() []domain.Snippet {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Devolvemos una copia para no exponer el slice interno.
	result := make([]domain.Snippet, len(s.snippets))
	copy(result, s.snippets)

	return result
}

func (s *SnippetService) Create(input domain.CreateSnippetInput) (domain.Snippet, error) {
	log.Printf("Create snippet input: %+v", input)

	title := strings.TrimSpace(input.Title)
	code := strings.TrimSpace(input.Code)

	if title == "" {
		return domain.Snippet{}, errors.New("el título es obligatorio")
	}

	if code == "" {
		return domain.Snippet{}, errors.New("el código es obligatorio")
	}

	snippet := domain.Snippet{
		ID:        uuid.NewString(),
		Title:     title,
		Language:  strings.TrimSpace(input.Language),
		Code:      code,
		Tags:      input.Tags,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.snippets = append(s.snippets, snippet)

	return snippet, nil
}

// second phase
// func (s *SnippetService) DeleteSnippet(id string) error

// last
// func (s *SnippetService) UpdateSnippet(snippet domain.Snippet) (domain.Snippet, error)
