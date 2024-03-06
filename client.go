package sipope

import (
	"context"
	"io"
	"net/http"
)

const (
	prodUrl    = "https://sipope.seguranca.mg.gov.br/api"
	homologUrl = "http://homologacao.depenone.seguranca.mg.gov.br/api"
)

type Client struct {
	apiKey string
	url    string
}

type ClientOptFn func(c *Client)

// Usa a URL de homologação para fazer as requisições.
func WithHomologacaoURL() ClientOptFn {
	return func(c *Client) {
		c.url = homologUrl
	}
}

// Usa uma URL customizada para fazer as requisições. Ex: http://localhost:8000/api.
func WithCustomURL(customUrl string) ClientOptFn {
	return func(c *Client) {
		c.url = customUrl
	}
}

func NewClient(apiKey string, opts ...ClientOptFn) *Client {
	c := &Client{
		apiKey: apiKey,
		url:    prodUrl,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Cria uma nova requisição com os headers obrigatórios definidos.
func (c *Client) newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Api-Key", c.apiKey)
	return req, nil
}
