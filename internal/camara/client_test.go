package camara

import (
	"net/url"
	"testing"
)

func TestBuildURL(t *testing.T) {
	client := NewClient(DefaultConfig())

	params := url.Values{}
	params.Add("ano", "2026")
	params.Add("mes", "3")

	got, err := client.BuildURL(
		"deputados/123/despesas",
		params,
	)

	if err != nil {
		t.Fatalf("esperava nil, recebeu %v", err)
	}

	want := "https://dadosabertos.camara.leg.br/api/v2/deputados/123/despesas?ano=2026&mes=3"

	if got != want {
		t.Errorf("esperava %s, recebeu %s", want, got)
	}
}
