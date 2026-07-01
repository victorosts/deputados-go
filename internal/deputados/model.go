package deputados

type Deputado struct {
	ID   int
	Nome string
}

type Despesa struct {
	TipoDespesa    string
	TipoDocumento  string
	DataDocumento  string
	ValorDocumento float64
}
