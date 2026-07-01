package camara

type DespesasResponse struct {
	Dados []DespesaDTO `json:"dados"`
}

type DeputadosResponse struct {
	Dados []DeputadoDTO `json:"dados"`
}

type DeputadoDetalhesResponse struct {
	Dados DeputadoDTO `json:"dados"`
}

type DeputadoDTO struct {
	ID           int    `json:"id"`
	NomeCivil    string `json:"nomeCivil"`
	Nome         string `json:"nome"`
	SiglaPartido string `json:"siglaPartido"`
	SiglaUf      string `json:"siglaUf"`
	URLFoto      string `json:"urlFoto"`
	Email        string `json:"email"`
}

type DespesaDTO struct {
	Ano               int     `json:"ano"`
	Mes               int     `json:"mes"`
	TipoDespesa       string  `json:"tipoDespesa"`
	DataDocumento     string  `json:"dataDocumento"`
	NomeFornecedor    string  `json:"nomeFornecedor"`
	CnpjCpfFornecedor string  `json:"cnpjCpfFornecedor"`
	ValorDocumento    float64 `json:"valorDocumento"`
	ValorLiquido      float64 `json:"valorLiquido"`
	ValorGlosa        float64 `json:"valorGlosa"`
	UrlDocumento      string  `json:"urlDocumento"`
}
