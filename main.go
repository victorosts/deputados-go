package main

import (
	"context"
	"deputados-go/internal/camara"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	start := time.Now()
	fmt.Println("Programa Inicializado")

	client := camara.NewClient(camara.DefaultConfig())
	g, ctx := errgroup.WithContext(context.Background())
	id := 178937

	var (
		deputado *camara.DeputadoDetalhes
		despesas []camara.Despesa
	)

	g.Go(func() error {
		var err error
		deputado, err = client.GetDeputado(ctx, id)
		return err
	})

	g.Go(func() error {
		var err error
		despesas, err = client.GetDeputadoDespesas(ctx, id, 2026, 3)
		return err
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Deputado -> %s | ID -> %d\n", deputado.NomeCivil, deputado.ID)

	for _, despesa := range despesas {
		fmt.Printf("Despesa -> %s\n", despesa.TipoDespesa)
		fmt.Printf("Data da despesa -> %s\n", despesa.DataDocumento)
	}

	fmt.Println("Programa Finalizado em:", time.Since(start))
}
