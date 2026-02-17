package ollama_go_api_v1

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Options for generation parameters
type Options struct {
	Temperature   float64 `json:"temperature,omitempty"`
	TopP          float64 `json:"top_p,omitempty"`
	TopK          int     `json:"top_k,omitempty"`
	RepeatPenalty float64 `json:"repeat_penalty,omitempty"`
	Seed          int     `json:"seed,omitempty"`
	MaxTokens     int     `json:"num_predict,omitempty"`
}

// ChatRequest represents a chat request
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
	Format   string    `json:"format,omitempty"`
	Options  *Options  `json:"options,omitempty"`
}

// ChatResponse represents a chat response
type ChatResponse struct {
	Model     string  `json:"model"`
	CreatedAt string  `json:"created_at"`
	Message   Message `json:"message"`
	Done      bool    `json:"done"`
}

// GenerateRequest represents a generate request
type GenerateRequest struct {
	Model   string   `json:"model"`
	Prompt  string   `json:"prompt"`
	Stream  bool     `json:"stream"`
	Format  string   `json:"format,omitempty"`
	Options *Options `json:"options,omitempty"`
}

// GenerateResponse represents a generate response
type GenerateResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

// ListResponse represents a list of models
type ListResponse struct {
	Models []ModelInfo `json:"models"`
}

// ModelInfo contains information about a model
type ModelInfo struct {
	Name       string `json:"name"`
	Model      string `json:"model"`
	Size       int64  `json:"size"`
	Digest     string `json:"digest"`
	ModifiedAt string `json:"modified_at"`
}

// PullRequest represents a pull request
type PullRequest struct {
	Model    string `json:"model"`
	Insecure bool   `json:"insecure,omitempty"`
}

// PullProgress represents pull progress
type PullProgress struct {
	Status    string `json:"status"`
	Digest    string `json:"digest,omitempty"`
	Total     int64  `json:"total,omitempty"`
	Completed int64  `json:"completed,omitempty"`
}

// ShowRequest represents a show request
type ShowRequest struct {
	Model string `json:"model"`
}

// ShowResponse represents detailed model information
type ShowResponse struct {
	License    string            `json:"license"`
	Modelfile  string            `json:"modelfile"`
	Parameters string            `json:"parameters"`
	Template   string            `json:"template"`
	System     string            `json:"system"`
	Details    map[string]string `json:"details"`
}
