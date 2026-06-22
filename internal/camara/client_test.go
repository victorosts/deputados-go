package camara

import (
	"context"
	"net/http"
	"net/http/httptest"
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

func TestGetDeputado(t *testing.T) {
	// Arrange
	server := httptest.NewServer(
		http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			w.WriteHeader(http.StatusOK)

			w.Write([]byte(`
			{
				"dados": {
					"id": 123,
					"nomeCivil": "José da Silva"
				}
			}
			`))
		}),
	)

	defer server.Close()

	client := NewClient(Config{
		BaseURL: server.URL,
	})

	// Act
	deputado, err := client.GetDeputado(
		context.Background(),
		123,
	)

	// Assert
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}

	if deputado.ID != 123 {
		t.Fatalf("esperava ID: 123, recebeu ID: %d", deputado.ID)
	}
}

func TestGetDeputado_ServerError(t *testing.T) {
	// Arrange
	server := httptest.NewServer(
		http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			http.Error(
				w,
				"erro interno",
				http.StatusInternalServerError,
			)
		}),
	)

	defer server.Close()

	client := NewClient(Config{
		BaseURL: server.URL,
	})

	// Act
	_, err := client.GetDeputado(
		context.Background(),
		123,
	)

	// Assert
	if err == nil {
		t.Fatal("esperava erro")
	}
}

func TestGetDeputado_Path(t *testing.T) {
	// Arrange
	var path string

	server := httptest.NewServer(
		http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			path = r.URL.Path

			w.Write([]byte(`
			{
				"dados": {
					"id": 123,
					"nomeCivil": "José da Silva"
				}
			}
			`))
		}),
	)

	defer server.Close()

	client := NewClient(Config{
		BaseURL: server.URL,
	})

	// Act
	_, _ = client.GetDeputado(
		context.Background(),
		123,
	)

	// Assert
	if path != "/deputados/123" {
		t.Fatalf(
			"path incorreto: %s",
			path,
		)
	}
}
