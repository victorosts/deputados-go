package main

import (
	"context"
	"deputados-go/internal/camara"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Programa Inicializado")

	var wg sync.WaitGroup
	ctx := context.Background()
	client := camara.NewClient(camara.DefaultConfig())
	id := 178937

	wg.Add(1)
	go func() {
		defer wg.Done()
		deputado, err := client.GetDeputado(ctx, id)
		if err == nil {
			fmt.Println(*deputado)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		despesas, err := client.GetDeputadoDespesas(ctx, id, 2026, 3)
		if err == nil {
			fmt.Println(despesas)
		}
	}()

	wg.Wait()

	fmt.Println("Programa Finalizado")
}
