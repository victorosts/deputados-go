package camara

type DeputadoDTO struct {
	ID           int    `json:"id"`
	NomeCivil    string `json:"nomeCivil"`
	Nome         string `json:"nome"`
	SiglaPartido string `json:"siglaPartido"`
	SiglaUf      string `json:"siglaUf"`
	URLFoto      string `json:"urlFoto"`
	Email        string `json:"email"`
}

type DeputadosResponse struct {
	Dados []DeputadoDTO `json:"dados"`
}
