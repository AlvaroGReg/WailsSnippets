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
	log.Printf("snippet_service::List() => result:: %+v", result)

	return result
}

// CREATE
func (s *SnippetService) Create(input domain.CreateSnippetInput) (domain.Snippet, error) {
	log.Printf("snippet_service::Create(input)::input %+v", input)

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
	log.Printf("snippet_service::Create(input) => res::res %+v", snippet)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.snippets = append(s.snippets, snippet)

	return snippet, nil
}

// DELETE
func (s *SnippetService) DeleteSnippet(id string) error {
	log.Printf("snippet_service::DeleteSnippet(id)::id %+v", id)

	for i, value := range s.snippets {
		if value.ID == id {
			log.Printf("snippet_service::DeleteSnippet(id)::snippetToDelete %+v", value)
			s.snippets = append(s.snippets[:i], s.snippets[i+1:]...)
			return nil
		}
	}
	return errors.New("No se pudo borrar el snippet")
}

// last
// func (s *SnippetService) UpdateSnippet(snippet domain.Snippet) (domain.Snippet, error)
