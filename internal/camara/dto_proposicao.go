package camara

type ProposicaoResponse struct {
	Dados []ProposicaoDTO `json:"dados"`
}

type ProposicaoDTO struct {
	ID               int    `json:"id"`
	Ementa           string `json:"ementa"`
	DataApresentacao string `json:"dataApresentacao"`
}
