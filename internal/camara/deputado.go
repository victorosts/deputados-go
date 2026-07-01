package camara

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) GetDeputados(
	ctx context.Context,
) ([]DeputadoDTO, error) {
	var response DeputadosResponse

	if err := c.ApiGet(ctx, "deputados", nil, &response); err != nil {
		return nil, fmt.Errorf("Erro ao buscar deputados: %w", err)
	}

	return response.Dados, nil
}

func (c *Client) GetDeputado(
	ctx context.Context,
	id int,
) (*DeputadoDTO, error) {
	var response DeputadoDetalhesResponse

	endpoint := fmt.Sprintf("deputados/%d", id)

	if err := c.ApiGet(ctx, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("Busca detalhes do deputado %d: %w", id, err)
	}

	return &response.Dados, nil
}

func (c *Client) GetDespesas(
	ctx context.Context,
	id int,
	year int,
	month int,
) ([]DespesaDTO, error) {
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
