package service

import (
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"WailsSnippets/internal/domain"
	"WailsSnippets/internal/repository"
)

type SnippetService struct {
	mu         sync.RWMutex
	repository *repository.JSONSnippetRepository
}

func NewSnippetService() *SnippetService {
	return &SnippetService{
		repository: repository.NewJSONSnippetRepository(),
	}
}

func (s *SnippetService) SetSnippetsDirectory(directory string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.repository.SetDirectory(directory)
}

func (s *SnippetService) SnippetsDirectory() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.repository.Directory()
}

func (s *SnippetService) EnsureSnippetsFile() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.repository.EnsureFile()
}

// GET
func (s *SnippetService) List() ([]domain.Snippet, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result, err := s.repository.List()
	if err != nil {
		return nil, err
	}
	log.Printf("snippet_service::List() => result:: %+v", result)

	return result, nil
}

// CREATE
func (s *SnippetService) CreateSnippet(input domain.CreateSnippetInput) (domain.Snippet, error) {
	log.Printf("snippet_service::CreateSnippet(input)::input %+v", input)

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

	log.Printf("snippet_service::CreateSnippet(input) => res::res %+v", snippet)

	s.mu.Lock()
	defer s.mu.Unlock()

	snippets, err := s.repository.List()
	if err != nil {
		return domain.Snippet{}, err
	}
	snippets = append(snippets, snippet)
	s.repository.SaveList(snippets)

	return snippet, nil
}

// DELETE
func (s *SnippetService) DeleteSnippet(id string) error {
	log.Printf("snippet_service::DeleteSnippet(id)::id %+v", id)
	snippets, err := s.List()
	if err != nil {
		return err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i, value := range snippets {
		if value.ID == id {
			log.Printf("snippet_service::DeleteSnippet(id)::snippetToDelete %+v", value)
			snippets = append(snippets[:i], snippets[i+1:]...)
			s.repository.SaveList(snippets)
			return nil
		}
	}
	return errors.New("unable to delete the snippet")
}

// UPDATE
func (s *SnippetService) UpdateSnippet(snippet domain.Snippet) (domain.Snippet, error) {
	log.Printf("snippet_service::UpdateSnippet(snippet)::newSnippet %+v", snippet)
	// normalize and validate snippet item
	snippet.Title = strings.TrimSpace(snippet.Title)
	snippet.Language = strings.TrimSpace(snippet.Language)
	snippet.Code = strings.TrimSpace(snippet.Code)

	if err := CheckValidSnippet(snippet); err != nil {
		return domain.Snippet{}, err
	}
	// get and check list
	snippets, err := s.List()
	if err != nil {
		return domain.Snippet{}, err
	}
	// rewrite item, then save
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, value := range snippets {
		if value.ID == snippet.ID {
			snippet.CreatedAt = value.CreatedAt
			snippets[i] = snippet
			s.repository.SaveList(snippets)
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
