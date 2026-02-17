package ollama_go_api_v1

import (
	"encoding/json"
	"strings"
)

// ExtractJSON tries to extract JSON from text
func ExtractJSON(text string) string {
	start := strings.Index(text, "{")
	if start == -1 {
		start = strings.Index(text, "[")
	}
	if start == -1 {
		return ""
	}

	end := strings.LastIndex(text, "}")
	if end == -1 {
		end = strings.LastIndex(text, "]")
	}
	if end == -1 || end <= start {
		return ""
	}

	return text[start : end+1]
}

// ParseJSON parses JSON into a struct
func ParseJSON[T any](text string) (*T, error) {
	cleaned := ExtractJSON(text)
	if cleaned == "" {
		cleaned = text
	}

	var result T
	if err := json.Unmarshal([]byte(cleaned), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DefaultOptions returns default generation options
func DefaultOptions() *Options {
	return &Options{
		Temperature:   0.7,
		TopP:          0.9,
		TopK:          40,
		RepeatPenalty: 1.1,
	}
}

// StrictOptions returns strict options for precise output
func StrictOptions() *Options {
	return &Options{
		Temperature:   0.1,
		TopP:          0.9,
		TopK:          40,
		RepeatPenalty: 1.1,
	}
}
