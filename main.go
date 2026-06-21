package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("start point")

	GetDeputados()
}

func BaseRequest(endpoint string) (resp *http.Response, err error) {
	baseUrl := "https://dadosabertos.camara.leg.br/api/v2/"
	result, err := url.JoinPath(baseUrl, endpoint)
	if err != nil {
		panic(err)
	}

	return http.Get(result)

}

func GetDeputados() (string, error) {
	resp, err := BaseRequest("deputados")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	return string(body), nil
}
