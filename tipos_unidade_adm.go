package sipope

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TipoUnidadeAdministrativa struct {
	ID           int64      `json:"id"`
	Nome         string     `json:"nome"`
	CriadoEm     time.Time  `json:"criado_em"`
	AtualizadoEm time.Time  `json:"atualizado_em"`
	ApagadoEm    *time.Time `json:"apagado_em"`
}

func (c *Client) ListTiposUnidadeAdministrativa(ctx context.Context) ([]*TipoUnidadeAdministrativa, error) {
	req, err := c.newRequest(ctx, http.MethodGet, c.url+"/tipos-unidade-administrativa", nil)
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
		var tipos []*TipoUnidadeAdministrativa
		if err := json.NewDecoder(res.Body).Decode(&tipos); err != nil {
			return nil, err
		}
		return tipos, nil
	case http.StatusUnauthorized:
		return nil, ErrUnauthorized
	default:
		return nil, fmt.Errorf("unexpected error response: %d", res.StatusCode)
	}
}

func (c *Client) GetTipoUnidadeAdministrativa(ctx context.Context, id int64) (*TipoUnidadeAdministrativa, error) {
	url := fmt.Sprintf("%s/tipos-unidade-administrativa/%d", c.url, id)
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
		var tipo TipoUnidadeAdministrativa
		if err := json.NewDecoder(res.Body).Decode(&tipo); err != nil {
			return nil, err
		}
		return &tipo, nil
	case http.StatusUnauthorized:
		return nil, ErrUnauthorized
	case http.StatusNotFound:
		return nil, ErrNotFound
	default:
		return nil, fmt.Errorf("unexpected error responde: %d", res.StatusCode)
	}
}
