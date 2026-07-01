package camara

type DeputadoDetalhesDTO struct {
	ID                  int    `json:"id"`
	CPF                 string `json:"cpf"`
	DataFalecimento     string `json:"dataFalecimento"`
	DataNascimento      string `json:"dataNascimento"`
	Escolaridade        string `json:"escolaridade"`
	MunicipioNascimento string `json:"municipioNascimento"`
	NomeCivil           string `json:"nomeCivil"`
	Sexo                string `json:"sexo"`
	UFNascimento        string `json:"ufNascimento"`
}

type DeputadoDetalhesResponse struct {
	Dados DeputadoDetalhesDTO `json:"dados"`
}
