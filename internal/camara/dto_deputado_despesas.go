package camara

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

type DespesasResponse struct {
	Dados []DespesaDTO `json:"dados"`
}
