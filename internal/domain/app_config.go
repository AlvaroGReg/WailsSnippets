package domain

// AppConfig contains the preferences that are persisted between application runs.
type AppConfig struct {
	SnippetsDirectory string `json:"snippetsDirectory"`
}
