package deputados

type Deputado struct {
	ID           int
	Nome         string
	SiglaPartido string
	SiglaUf      string
}

type Detalhes struct {
	ID                  int
	CPF                 string
	DataFalecimento     string
	DataNascimento      string
	Escolaridade        string
	MunicipioNascimento string
	NomeCivil           string
	Sexo                string
	UFNascimento        string
}

type Despesa struct {
	TipoDespesa    string
	TipoDocumento  string
	DataDocumento  string
	ValorDocumento float64
}
