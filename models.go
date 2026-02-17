package ollama_go_api_v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// List returns a list of installed models
func (c *Client) List(ctx context.Context) (*ListResponse, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		c.baseURL+"/api/tags",
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("request creation error: %w", err)
	}

	resp, err := c.client.Do(req)
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

	var result ListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &result, nil
}

// Pull downloads a model from the Ollama library
func (c *Client) Pull(ctx context.Context, modelName string) error {
	reqBody := PullRequest{
		Model: modelName,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/api/pull",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return fmt.Errorf("request creation error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("api error (%d): %s", resp.StatusCode, string(body))
	}

	return nil
}

// Delete removes a model
func (c *Client) Delete(ctx context.Context, modelName string) error {
	reqBody := map[string]string{"name": modelName}
	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(
		ctx,
		"DELETE",
		c.baseURL+"/api/delete",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return fmt.Errorf("request creation error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("api error (%d): %s", resp.StatusCode, string(body))
	}

	return nil
}

// Show returns detailed information about a model
func (c *Client) Show(ctx context.Context, modelName string) (*ShowResponse, error) {
	reqBody := ShowRequest{Model: modelName}
	jsonData, _ := json.Marshal(reqBody)

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		c.baseURL+"/api/show",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, fmt.Errorf("request creation error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
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

	var result ShowResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &result, nil
}
