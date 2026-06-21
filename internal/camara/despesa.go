package camara

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type Despesa struct {
	TipoDespesa    string  `json:"tipoDespesa"`
	TipoDocumento  string  `json:"tipoDocumento"`
	DataDocumento  string  `json:"dataDocumento"`
	ValorDocumento float64 `json:"valorDocumento"`
}

type DespesasResponse struct {
	Dados []Despesa `json:"dados"`
}

func (c *Client) GetDeputadoDespesas(
	ctx context.Context,
	id int,
	year int,
	month int,
) ([]Despesa, error) {
	var response DespesasResponse

	params := url.Values{
		"ano": []string{strconv.Itoa(year)},
		"mes": []string{strconv.Itoa(month)},
	}
	endpoint := fmt.Sprintf("deputados/%d/despesas", id)

	if err := c.ApiGet(ctx, endpoint, params, &response); err != nil {
		return nil, fmt.Errorf("Busca despesas do deputado %d: %w", id, err)
	}

	return response.Dados, nil
}
