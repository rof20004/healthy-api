package entities

import (
	"strings"
	"time"

	profissionalErrors "github.com/rof20004/healthy-api/application/domains/profissional/errors"
	"github.com/rof20004/healthy-api/application/valueobjects"

	"github.com/google/uuid"
)

type Profissional struct {
	CreatedAt time.Time          `json:"createdAt"`
	Id        string             `json:"id"`
	Nome      string             `json:"nome"`
	CPF       string             `json:"cpf"`
	Email     valueobjects.Email `json:"email"`
	Foto      string             `json:"foto"`
	CRP       string             `json:"crp"`
	Senha     valueobjects.Senha `json:"senha"`
}

func NewProfissional(nome, cpf, foto, crp string, email valueobjects.Email, senha valueobjects.Senha) Profissional {
	return Profissional{
		Id:        uuid.NewString(),
		Nome:      nome,
		CPF:       cpf,
		Email:     email,
		Foto:      foto,
		CRP:       crp,
		Senha:     senha,
		CreatedAt: time.Now(),
	}
}

func validateProfissionalId(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return profissionalErrors.ErrIdInvalido.WithRootCause(err)
	}

	return nil
}

func validateProfissionalNome(nome string) error {
	if strings.TrimSpace(nome) == "" {
		return profissionalErrors.ErrNomeObrigatorio
	}

	return nil
}

func validateProfissionalCPF(cpf string) error {
	if strings.TrimSpace(cpf) == "" {
		return profissionalErrors.ErrCPFObrigatorio
	}

	return nil
}

func validateProfissionalEmail(email valueobjects.Email) error {
	if strings.TrimSpace(email.Value) == "" {
		return profissionalErrors.ErrEmailObrigatorio
	}

	if ok := email.IsValid(); !ok {
		return profissionalErrors.ErrEmailInvalido
	}

	return nil
}

func validateProfissionalFoto(foto string) error {
	if strings.TrimSpace(foto) == "" {
		return profissionalErrors.ErrFotoObrigatorio
	}

	return nil
}

func validateProfissionalCRP(crp string) error {
	if strings.TrimSpace(crp) == "" {
		return profissionalErrors.ErrCRPObrigatorio
	}

	return nil
}

func validateProfissionalSenha(senha valueobjects.Senha) error {
	if strings.TrimSpace(senha.Value) == "" {
		return profissionalErrors.ErrSenhaObrigatoria
	}

	return nil
}

func (u Profissional) ValidateNew() error {
	if err := validateProfissionalId(u.Id); err != nil {
		return err
	}

	if err := validateProfissionalNome(u.Nome); err != nil {
		return err
	}

	if err := validateProfissionalCPF(u.CPF); err != nil {
		return err
	}

	if err := validateProfissionalEmail(u.Email); err != nil {
		return err
	}

	if err := validateProfissionalFoto(u.Foto); err != nil {
		return err
	}

	if err := validateProfissionalCRP(u.CRP); err != nil {
		return err
	}

	if err := validateProfissionalSenha(u.Senha); err != nil {
		return err
	}

	return nil
}
