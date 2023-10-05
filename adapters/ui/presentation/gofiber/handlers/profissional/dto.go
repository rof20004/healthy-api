package profissional

import (
	profissionalEntities "github.com/rof20004/healthy-api/application/domains/profissional/entities"
	"github.com/rof20004/healthy-api/application/valueobjects"
)

type CreateProfissionalRequest struct {
	Nome  string `json:"nome"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
	Foto  string `json:"foto"`
	CRP   string `json:"crp"`
	Senha string `json:"senha"`
}

func (c CreateProfissionalRequest) ToEntity() profissionalEntities.Profissional {
	email := valueobjects.NewEmail(c.Email)
	senha := valueobjects.NewSenha(c.Senha)
	return profissionalEntities.NewProfissional(c.Nome, c.CPF, c.Foto, c.CRP, email, senha)
}
