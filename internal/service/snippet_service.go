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

	snippet := domain.Snippet{
		ID:        uuid.NewString(),
		Title:     strings.TrimSpace(input.Title),
		Language:  strings.TrimSpace(input.Language),
		Code:      strings.TrimSpace(input.Code),
		Tags:      input.Tags,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	if err := CheckValidSnippet(snippet); err != nil {
		return domain.Snippet{}, err
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
	return errors.New("unable to delete the snippet")
}

// UPDATE
func (s *SnippetService) UpdateSnippet(snippet domain.Snippet) (domain.Snippet, error) {
	log.Printf("snippet_service::UpdateSnippet(snippet)::newSnippet %+v", snippet)

	snippet.Title = strings.TrimSpace(snippet.Title)
	snippet.Language = strings.TrimSpace(snippet.Language)
	snippet.Code = strings.TrimSpace(snippet.Code)

	if err := CheckValidSnippet(snippet); err != nil {
		return domain.Snippet{}, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for i, value := range s.snippets {
		if value.ID == snippet.ID {
			snippet.CreatedAt = value.CreatedAt
			s.snippets[i] = snippet
			return snippet, nil
		}
	}

	return domain.Snippet{}, errors.New("snippet not found")
}

// UTILS
func CheckValidSnippet(snippet domain.Snippet) error {
	if snippet.Title == "" {
		return errors.New("title is required")
	}

	if snippet.Code == "" {
		return errors.New("code is required")
	}

	return nil
}
