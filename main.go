package main

import (
	"context"
	"deputados-go/internal/camara"
	"deputados-go/internal/deputados"
	"deputados-go/internal/proposicoes"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	start := time.Now()
	fmt.Println("Programa Inicializado")

	client := camara.NewClient(camara.DefaultConfig())
	deputadoService := deputados.NewService(client)
	proposicoesService := proposicoes.NewService(client)
	g, ctx := errgroup.WithContext(context.Background())
	id := 178937

	var (
		deputado            *deputados.Deputado
		deputadoProposicoes []proposicoes.Proposicao
	)

	g.Go(func() error {
		var err error
		deputado, err = deputadoService.ListarDeputadoDetalhes(ctx, id)
		return err
	})

	g.Go(func() error {
		var err error
		deputadoProposicoes, err = proposicoesService.ListarProposicoes(ctx, id)
		return err
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Deputado -> %s | ID -> %d\n", deputado.Nome, deputado.ID)

	for _, proposicao := range deputadoProposicoes {
		fmt.Printf("ID da proposição -> %d\n", proposicao.ID)
		fmt.Println(proposicao.Ementa)
		fmt.Printf("Data da proposição -> %s\n\n", proposicao.DataApresentacao)
	}

	fmt.Println("Programa Finalizado em:", time.Since(start))
}
