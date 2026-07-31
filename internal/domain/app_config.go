package domain

// AppConfig contains the preferences that are persisted between application runs.
type AppConfig struct {
	SnippetsFilePath string `json:"snippetsFilePath"`
	// SnippetsDirectory preserves the previous configuration format long enough to
	// migrate existing installations to the default snippets.json file.
	SnippetsDirectory string `json:"snippetsDirectory,omitempty"`
}
