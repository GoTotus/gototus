package totus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Totus struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewTotus creates a new Totus client instance
func NewTotus(apiKey, endpoint, proxy string) (*Totus, error) {
	if apiKey == "" {
		apiKey = os.Getenv("TOTUS_KEY")
	}
	if apiKey == "" {
		return nil, ErrMissingAPIKey
	}
	if endpoint == "" {
		endpoint = "https://api.totus.cloud"
	}

	client := &http.Client{
		Timeout: 30 * time.Second, // Prevent hanging requests
	}
	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy URL: %w", err)
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	return &Totus{
		apiKey:  apiKey,
		baseURL: strings.TrimRight(endpoint, "/"),
		client:  client,
	}, nil
}

// makeRequest performs an HTTP request to the Totus API and decodes the response into the target
func (t *Totus) makeRequest(method, endpoint string, params url.Values, body io.Reader, target any) error {
	urlStr := t.baseURL + "/" + strings.TrimLeft(endpoint, "/")
	if len(params) > 0 {
		urlStr += "?" + params.Encode()
	}

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+t.apiKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := t.client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		message := string(bodyBytes)
		if message == "" {
			message = "unknown error"
		}
		switch {
		case resp.StatusCode == 401:
			return fmt.Errorf("%s: %w", message, ErrAuthentication)
		case resp.StatusCode == 404:
			return fmt.Errorf("%s: %w", message, ErrNotFound)
		case resp.StatusCode >= 400 && resp.StatusCode < 500:
			return fmt.Errorf("%s: %w", message, ErrClient)
		case resp.StatusCode >= 500:
			return fmt.Errorf("%s: %w", message, ErrServer)
		default:
			return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, message)
		}
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// Reference returns a Reference service instance
func (t *Totus) Reference() *Reference {
	return &Reference{t}
}

// Validate returns a Validate service instance
func (t *Totus) Validate() *Validate {
	return &Validate{t}
}

func MapStrWithDef(m map[string]any, key string, def string) string {
	if v, ok := m[key]; ok {
		return fmt.Sprintf("%v", v)
	} else {
		return def
	}
}
