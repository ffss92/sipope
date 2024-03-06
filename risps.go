package sipope

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Risp struct {
	ID           int        `json:"id"`
	Nome         string     `json:"nome"`
	Sigla        string     `json:"sigla"`
	CriadoEm     time.Time  `json:"criado_em"`
	AtualizadoEm time.Time  `json:"atualizado_em"`
	ApagadoEm    *time.Time `json:"apagado_em"`
}

func (r *Risp) HasApagadoEm() bool {
	return r.ApagadoEm != nil
}

func (c *Client) ListRisps(ctx context.Context) (*Paginated[*Risp], error) {
	req, err := c.newRequest(ctx, http.MethodGet, c.url+"/risps", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		var result Paginated[*Risp]
		if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	case http.StatusUnauthorized:
		return nil, ErrUnauthorized
	default:
		return nil, fmt.Errorf("unexpected error response: %d", res.StatusCode)
	}
}

func (c *Client) GetRisp(ctx context.Context, id int64) (*Risp, error) {
	url := fmt.Sprintf("%s/risps/%d", c.url, id)
	req, err := c.newRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		var risp Risp
		if err := json.NewDecoder(res.Body).Decode(&risp); err != nil {
			return nil, err
		}
		return &risp, nil
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusUnauthorized:
		return nil, ErrUnauthorized
	default:
		return nil, fmt.Errorf("unexpected error response: %d", res.StatusCode)
	}
}
