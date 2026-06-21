package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type CamaraClient struct {
	client  *http.Client
	baseURL string
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
	camara := NewCamaraClient()

	ctx := context.Background()

	var deputados DeputadosResponse

	if err := camara.ApiGet(ctx, "deputados", nil, &deputados); err != nil {
		fmt.Printf("Falha na solicitação dos dados dos deputados, ERR: %s", err.Error())
		return
	}

	fmt.Println("Gerando amostragem dos deputados")
	for i, deputado := range deputados.Dados {
		if i == 5 {
			fmt.Println("Amostragem finalizada corretamente")
			break
		}

		detalhesDeputado, err := camara.GetDeputado(ctx, deputado.ID)
		if err != nil {
			fmt.Printf("Falha ao consultar os detalhes do deputado de ID: %d. Err: %s\n", deputado.ID, err)
			continue
		}
		fmt.Printf("ID: %d - Nome Civil: %s\n", detalhesDeputado.ID, detalhesDeputado.NomeCivil)

		despesas, err := camara.GetDeputadoDespesas(ctx, deputado.ID, 2026, 3)
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

func NewCamaraClient() *CamaraClient {
	return &CamaraClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://dadosabertos.camara.leg.br/api/v2/",
	}
}

func (c *CamaraClient) Do(
	ctx context.Context,
	method string,
	endpoint string,
	target any,
) error {
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		endpoint,
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("camara api retornou status: %d", resp.StatusCode)
	}

	// Caso seja uma request com resposta 204 esperada
	// e não seja enviado um target para atribuir valor
	if target == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *CamaraClient) BuildURL(
	endpoint string,
	params url.Values,
) (string, error) {
	requestURL, err := url.JoinPath(c.baseURL, endpoint)
	if err != nil {
		return "", err
	}

	u, err := url.Parse(requestURL)
	if err != nil {
		return "", err
	}

	u.RawQuery = params.Encode()

	return u.String(), nil
}

func (c *CamaraClient) ApiGet(
	ctx context.Context,
	endpoint string,
	params url.Values,
	target any,
) error {
	requestURL, err := c.BuildURL(endpoint, params)
	if err != nil {
		return err
	}

	if err := c.Do(
		ctx,
		http.MethodGet,
		requestURL,
		target,
	); err != nil {
		return err
	}

	return nil
}

func (c *CamaraClient) GetDeputado(
	ctx context.Context,
	id int,
) (*DeputadoDetalhes, error) {
	var response DeputadoDetalhesResponse

	endpoint := fmt.Sprintf("deputados/%d", id)

	if err := c.ApiGet(ctx, endpoint, nil, &response); err != nil {
		return nil, fmt.Errorf("Busca detalhes do deputado %d: %w", id, err)
	}

	return &response.Dados, nil
}

func (c *CamaraClient) GetDeputadoDespesas(
	ctx context.Context,
	id int,
	year int,
	month int,
) ([]Despesas, error) {
	var response DespesasResponse

	params := url.Values{
		"ano": []string{strconv.Itoa(year)},
		"mes": []string{strconv.Itoa(month)},
	}
	endpoint := fmt.Sprintf("deputados/%d/despesas", id)

	if err := c.ApiGet(ctx, endpoint, params, &response); err != nil {
		return nil, fmt.Errorf("Busca despesas do deputado %d: %w", id, err)
	}

	return response.Dados, nil
}
