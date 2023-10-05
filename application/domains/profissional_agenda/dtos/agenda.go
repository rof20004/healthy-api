package dtos

type AgendaDto struct {
	Profissional ProfissionalDto `json:"profissional"`
	Datas        []DataDto       `json:"datas"`
}

type ProfissionalDto struct {
	Id    string `json:"id"`
	Nome  string `json:"nome"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
	Foto  string `json:"foto"`
	CRP   string `json:"crp"`
}

type DataDto struct {
	Dia  string `json:"dia"`
	Hora string `json:"hora"`
}
