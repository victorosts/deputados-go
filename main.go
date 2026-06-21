package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("start point")

	resp, err := http.Get("https://dadosabertos.camara.leg.br/api/v2/deputados")
	if err != nil {
		fmt.Printf("Application ended with error: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error while reading response body: %s", err)
		return
	}

	fmt.Println(string(body))
}
