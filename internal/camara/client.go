package camara

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	client  *http.Client
	baseURL string
}

func NewClient(config Config) *Client {
	if config.BaseURL == "" {
		config.BaseURL = "https://dadosabertos.camara.leg.br/api/v2/"
	}

	if config.Timeout <= 0 {
		config.Timeout = 10 * time.Second
	}

	return &Client{
		client: &http.Client{
			Timeout: config.Timeout,
		},
		baseURL: config.BaseURL,
	}
}

func (c *Client) Do(
	ctx context.Context,
	method string,
	endpoint string,
	body io.Reader,
	target any,
) error {
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		endpoint,
		body,
	)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errBody, _ := io.ReadAll(resp.Body)

		return fmt.Errorf("camara api retornou status %d: %s", resp.StatusCode, string(errBody))
	}

	// Caso seja uma request com resposta 204 esperada
	// e não seja enviado um target para atribuir valor
	if target == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *Client) BuildURL(
	endpoint string,
	params url.Values,
) (string, error) {
	requestURL, err := url.JoinPath(c.baseURL, endpoint)
	if err != nil {
		return "", err
	}

	u, err := url.Parse(requestURL)
	if err != nil {
		return "", err
	}

	if params != nil {
		u.RawQuery = params.Encode()
	}

	return u.String(), nil
}

func (c *Client) ApiGet(
	ctx context.Context,
	endpoint string,
	params url.Values,
	target any,
) error {
	requestURL, err := c.BuildURL(endpoint, params)
	if err != nil {
		return err
	}

	return c.Do(
		ctx,
		http.MethodGet,
		requestURL,
		nil,
		target,
	)
}
