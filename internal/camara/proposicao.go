package camara

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) GetProposicoes(
	ctx context.Context,
	idDeputado int,
) ([]ProposicaoDTO, error) {
	var response ProposicaoResponse

	params := url.Values{
		"idDeputadoAutor": []string{strconv.Itoa(idDeputado)},
	}

	if err := c.ApiGet(ctx, "proposicoes", params, &response); err != nil {
		return nil, fmt.Errorf("Erro ao buscar as proposições do deputado %d: %w", idDeputado, err)
	}

	return response.Dados, nil
}
