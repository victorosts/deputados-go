package main

import (
	"context"
	"deputados-go/internal/camara"
	"deputados-go/internal/deputados"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Programa Inicializado")

	client := camara.NewClient(camara.DefaultConfig())
	deputadoService := deputados.NewService(client)
	// id := 178937

	var data []deputados.Deputado
	data, err := deputadoService.ListarDeputados(
		context.Background(),
		camara.DeputadoFilter{
			UF: "MG",
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	for _, deputado := range data {
		fmt.Printf("ID: %d, Nome: %s, Partido: %s, UF: %s\n", deputado.ID, deputado.Nome, deputado.SiglaPartido, deputado.SiglaUf)
	}

	fmt.Println("Programa Finalizado em:", time.Since(start))
}
