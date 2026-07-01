package proposicoes

import (
	"context"
	"deputados-go/internal/camara"
)

type Service struct {
	camara *camara.Client
}

func NewService(
	camara *camara.Client,
) *Service {
	return &Service{
		camara: camara,
	}
}

func (s *Service) ListarProposicoes(
	ctx context.Context,
	idDeputado int,
) ([]Proposicao, error) {
	dtos, err := s.camara.GetProposicoes(ctx, idDeputado)

	if err != nil {
		return nil, err
	}

	proposicoes := make(
		[]Proposicao,
		0,
		len(dtos),
	)

	for _, dto := range dtos {
		proposicoes = append(
			proposicoes,
			Proposicao{
				ID:               dto.ID,
				Ementa:           dto.Ementa,
				DataApresentacao: dto.DataApresentacao,
			},
		)
	}

	return proposicoes, nil
}
