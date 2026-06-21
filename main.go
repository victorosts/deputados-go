package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://dadosabertos.camara.leg.br/api/v2/"

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

type DeputadoDetalhes struct {
	ID        int    `json:"id"`
	NomeCivil string `json:"nomeCivil"`
}

type DeputadoDetalhesResponse struct {
	Dados DeputadoDetalhes `json:"dados"`
}

type Despesas struct {
	TipoDespesa    string  `json:"tipoDespesa"`
	TipoDocumento  string  `json:"tipoDocumento"`
	DataDocumento  string  `json:"dataDocumento"`
	ValorDocumento float64 `json:"valorDocumento"`
}

type DespesasResponse struct {
	Dados []Despesas `json:"dados"`
}

func main() {
	fmt.Println("Iniciando aplicação")

	var deputados DeputadosResponse

	if err := CamaraApiGet("deputados", nil, &deputados); err != nil {
		fmt.Printf("Falha na solicitação dos dados dos deputados, ERR: %s", err.Error())
		return
	}

	fmt.Println("Gerando amostragem dos deputados")
	for i, deputado := range deputados.Dados {
		if i == 5 {
			fmt.Println("Amostragem finalizada corretamente")
			break
		}

		detalhesDeputado, err := GetDeputado(deputado.ID)
		if err != nil {
			fmt.Printf("Falha ao consultar os detalhes do deputado de ID: %d. Err: %s\n", deputado.ID, err)
			continue
		}
		fmt.Printf("ID: %d - Nome Civil: %s\n", detalhesDeputado.ID, detalhesDeputado.NomeCivil)

		despesas, err := GetDeputadoDespesas(deputado.ID, 2026, 3)
		if err != nil {
			fmt.Printf("Não foi possível retornar as despesas do deputado %s. Err: %s\n", detalhesDeputado.NomeCivil, err)
			continue
		}

		for j, despesa := range despesas {
			if j == 2 {
				fmt.Println("Amostragem de despesas finalizada corretamente")
				break
			}

			fmt.Printf("Despesa: %s - Valor: %f\n", despesa.TipoDespesa, despesa.ValorDocumento)
		}
	}

}

func GetDeputado(id int) (*DeputadoDetalhes, error) {
	var response DeputadoDetalhesResponse

	endpoint := fmt.Sprintf("deputados/%d", id)

	if err := CamaraApiGet(endpoint, nil, &response); err != nil {
		return nil, err
	}

	return &response.Dados, nil
}

func GetDeputadoDespesas(id int, year int, month int) ([]Despesas, error) {
	var response DespesasResponse

	params := url.Values{
		"ano": []string{strconv.Itoa(year)},
		"mes": []string{strconv.Itoa(month)},
	}
	endpoint := fmt.Sprintf("deputados/%d/despesas", id)

	if err := CamaraApiGet(endpoint, params, &response); err != nil {
		return nil, err
	}

	return response.Dados, nil
}

func CamaraApiGet(endpoint string, params url.Values, target any) error {
	result, err := url.JoinPath(baseURL, endpoint)
	if err != nil {
		return err
	}

	if len(params) > 0 {
		result += "?" + params.Encode()
	}

	resp, err := client.Get(result)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("camara api retornou status: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
