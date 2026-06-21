package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const baseURL string = "https://dadosabertos.camara.leg.br/api/v2/"

var client = &http.Client{
	Timeout: 10 * time.Second,
}

type Deputado struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

type DeputadosResponse struct {
	Dados []Deputado `json:"dados"`
}

func main() {
	fmt.Println("Iniciando aplicação")

	body, err := CamaraApiGet("deputados")
	if err != nil {
		fmt.Printf("Falha na solicitação dos dados dos deputados, ERR: %s", err.Error())
		return
	}

	var deputados DeputadosResponse

	if err := json.Unmarshal(body, &deputados); err != nil {
		fmt.Printf("Falha ao converter JSON: %v\n", err)
		return
	}

	fmt.Println(deputados)
}

func CamaraApiGet(endpoint string) (body []byte, err error) {
	result, err := url.JoinPath(baseURL, endpoint)
	if err != nil {
		return
	}

	resp, err := client.Get(result)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Erro ao consultar API da Camara: %s", resp.Status)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return
}
