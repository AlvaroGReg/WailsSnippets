package service

import (
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"SnippetsDome/internal/domain"
	"SnippetsDome/internal/repository"
)

type SnippetService struct {
	mu               sync.RWMutex
	repository       *repository.JSONSnippetRepository
	config           domain.AppConfig
	configRepository *repository.JSONConfigRepository
}

func NewSnippetService(config domain.AppConfig, configRepository *repository.JSONConfigRepository) *SnippetService {
	snippetRepository := repository.NewJSONSnippetRepository()
	snippetRepository.SetDirectory(config.SnippetsDirectory)

	return &SnippetService{
		repository:       snippetRepository,
		config:           config,
		configRepository: configRepository,
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

// SelectSnippetsDirectory changes the snippets directory and persists the
// selection. If either operation fails, it restores the previous directory.
// An empty directory means the native selection dialog was cancelled.
func (s *SnippetService) SelectSnippetsDirectory(directory string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// TODO: directory is a must coming from frontend and "" will not be posible here
	if directory == "" {
		return s.repository.Directory(), nil
	}

	previousDirectory := s.repository.Directory()
	s.repository.SetDirectory(directory)
	if err := s.repository.EnsureFile(); err != nil {
		s.repository.SetDirectory(previousDirectory)
		return "", err
	}

	previousConfig := s.config
	s.config.SnippetsDirectory = directory
	if err := s.configRepository.SaveConfig(s.config); err != nil {
		s.config = previousConfig
		s.repository.SetDirectory(previousDirectory)
		return "", err
	}

	return directory, nil
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
