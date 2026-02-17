package ollama_go_api_v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents an Ollama API client
type Client struct {
	baseURL string
	client  *http.Client
}

// NewClient creates a new Ollama client
// baseURL should be like "http://localhost:11434" or "http://192.168.0.1:11434"
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 120 * time.Second,
		},
	}
}

// Chat sends a chat request to Ollama
func (c *Client) Chat(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/api/chat",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, fmt.Errorf("request creation error: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error (%d): %s", resp.StatusCode, string(body))
	}

	var result ChatResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &result, nil
}

// Generate sends a generate request to Ollama
func (c *Client) Generate(ctx context.Context, req *GenerateRequest) (*GenerateResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/api/generate",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, fmt.Errorf("request creation error: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error (%d): %s", resp.StatusCode, string(body))
	}

	var result GenerateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &result, nil
}

// Version gets the Ollama server version
func (c *Client) Version(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		c.baseURL+"/api/version",
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("request creation error: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("api error (%d): %s", resp.StatusCode, string(body))
	}

	var version struct {
		Version string `json:"version"`
	}
	if err := json.Unmarshal(body, &version); err != nil {
		return "", fmt.Errorf("unmarshal error: %w", err)
	}

	return version.Version, nil
}
