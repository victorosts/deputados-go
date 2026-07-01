package deputados

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

func (s *Service) ListarDeputados(
	ctx context.Context,
	filter camara.DeputadoFilter,
) ([]Deputado, error) {
	dtos, err := s.camara.GetDeputados(ctx, filter)

	if err != nil {
		return nil, err
	}

	deputados := make(
		[]Deputado,
		0,
		len(dtos),
	)

	for _, dto := range dtos {
		deputados = append(
			deputados,
			Deputado{
				ID:           dto.ID,
				Nome:         dto.Nome,
				SiglaPartido: dto.SiglaPartido,
				SiglaUf:      dto.SiglaUf,
			},
		)
	}

	return deputados, nil
}

func (s *Service) ListarDeputadoDetalhes(
	ctx context.Context,
	id int,
) (*Detalhes, error) {
	dto, err := s.camara.GetDeputado(ctx, id)

	if err != nil {
		return nil, err
	}

	return &Detalhes{
		ID:                  dto.ID,
		NomeCivil:           dto.NomeCivil,
		CPF:                 dto.CPF,
		Sexo:                dto.Sexo,
		DataNascimento:      dto.DataNascimento,
		DataFalecimento:     dto.DataFalecimento,
		UFNascimento:        dto.UFNascimento,
		Escolaridade:        dto.Escolaridade,
		MunicipioNascimento: dto.MunicipioNascimento,
	}, nil
}

func (s *Service) ListarDespesas(
	ctx context.Context,
	id int,
	year int,
	month int,
) ([]Despesa, error) {
	dtos, err := s.camara.GetDespesas(ctx, id, year, month)

	if err != nil {
		return nil, err
	}

	despesas := make(
		[]Despesa,
		0,
		len(dtos),
	)

	for _, dto := range dtos {
		despesas = append(despesas, Despesa{
			TipoDespesa:    dto.TipoDespesa,
			TipoDocumento:  dto.DataDocumento,
			DataDocumento:  dto.DataDocumento,
			ValorDocumento: dto.ValorDocumento,
		})
	}

	return despesas, nil
}
