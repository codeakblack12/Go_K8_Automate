package joincode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Create(joinCommand string) (*CreateJoinCodeResponse, error) {
	payload := CreateJoinCodeRequest{
		JoinCommand: strings.TrimSpace(joinCommand),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal create join-code request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/api/v1/join", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create join-code request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call join-code create API: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read join-code create response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			Body:       string(respBody),
		}
	}

	var result CreateJoinCodeResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to decode join-code create response: %w", err)
	}

	return &result, nil
}

func (c *Client) Resolve(joinCode string) (*ResolveJoinCodeResponse, error) {
	joinCode = strings.TrimSpace(joinCode)

	req, err := http.NewRequest(http.MethodGet, c.baseURL+"/api/v1/resolve/"+joinCode, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create resolve join-code request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call join-code resolve API: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read resolve join-code response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			Body:       string(respBody),
		}
	}

	var result ResolveJoinCodeResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to decode resolve join-code response: %w", err)
	}

	return &result, nil
}
