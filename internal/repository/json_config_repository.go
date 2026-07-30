package repository

import (
	"encoding/json"
	"os"
	"path/filepath"

	"SnippetsDome/internal/domain"
)

const configFileName = "config.json"

// JSONConfigRepository persists the application configuration in the user's
// operating-system configuration directory.
type JSONConfigRepository struct {
	applicationName string
}

func NewJSONConfigRepository(applicationName string) *JSONConfigRepository {
	return &JSONConfigRepository{applicationName: applicationName}
}

// Load returns an empty configuration when the config file does not exist yet.
func (r *JSONConfigRepository) Load() (domain.AppConfig, error) {
	path, err := r.filePath()
	if err != nil {
		return domain.AppConfig{}, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return domain.AppConfig{}, nil
	}
	if err != nil {
		return domain.AppConfig{}, err
	}

	var config domain.AppConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return domain.AppConfig{}, err
	}
	return config, nil
}

func (r *JSONConfigRepository) SaveConfig(config domain.AppConfig) error {
	path, err := r.filePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func (r *JSONConfigRepository) filePath() (string, error) {
	directory, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(directory, r.applicationName, configFileName), nil
}
