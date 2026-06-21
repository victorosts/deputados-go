package camara

import (
	"context"
	"fmt"
)

type Deputado struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

type DeputadosResponse struct {
	Dados []Deputado `json:"dados"`
}

type DeputadoDetalhes struct {
	ID        int    `json:"id"`
	NomeCivil string `json:"nomeCivil"`
}

type DeputadoDetalhesResponse struct {
	Dados DeputadoDetalhes `json:"dados"`
}

func (c *Client) GetDeputados(
	ctx context.Context,
) ([]Deputado, error) {
	var response DeputadosResponse

	if err := c.ApiGet(ctx, "deputados", nil, &response); err != nil {
		return nil, fmt.Errorf("Erro ao buscar deputados: %w", err)
	}

	return response.Dados, nil
}

func (c *Client) GetDeputado(
	ctx context.Context,
	id int,
) (*DeputadoDetalhes, error) {
	var response DeputadoDetalhesResponse

	endpoint := fmt.Sprintf("deputados/%d", id)

	if err := c.ApiGet(ctx, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("Busca detalhes do deputado %d: %w", id, err)
	}

	return &response.Dados, nil
}
