package main

import (
	"context"
	"deputados-go/internal/camara"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Iniciando aplicação")
	config := camara.DefaultConfig()
	camara := camara.NewClient(config)
	ctx := context.Background()

	deputados, err := camara.GetDeputados(ctx)
	if err != nil {
		fmt.Printf("Falha na solicitação dos dados dos deputados, ERR: %s", err.Error())
		return
	}

	fmt.Println("Gerando amostragem dos deputados")
	deputado := deputados[rand.Intn(15)]

	fmt.Printf("ID: %d - Nome: %s\n", deputado.ID, deputado.Nome)

	despesas, err := camara.GetDeputadoDespesas(ctx, deputado.ID, 2026, 3)
	if err != nil {
		fmt.Printf("Falha na solicitação das despesas do deputado %s: %s", deputado.Nome, err)
		return
	}

	despesa := despesas[0]
	fmt.Printf("Despesa: %s\n", despesa.TipoDespesa)
}
