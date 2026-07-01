package camara

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) GetDeputados(
	ctx context.Context,
	filter DeputadoFilter,
) ([]DeputadoDTO, error) {
	var response DeputadosResponse

	params := url.Values{}

	if filter.UF != "" {
		params.Set("siglaUf", filter.UF)
	}

	if filter.Partido != "" {
		params.Set("siglaPartido", filter.Partido)
	}

	if filter.Pagina > 0 {
		params.Set("pagina", strconv.Itoa(filter.Pagina))
	}

	if err := c.ApiGet(ctx, "deputados", params, &response); err != nil {
		return nil, fmt.Errorf("Erro ao buscar deputados: %w", err)
	}

	return response.Dados, nil
}

func (c *Client) GetDeputado(
	ctx context.Context,
	id int,
) (*DeputadoDetalhesDTO, error) {
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
