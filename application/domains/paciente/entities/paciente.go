package entities

import (
	"strings"
	"time"

	pacienteErrors "github.com/rof20004/healthy-api/application/domains/paciente/errors"

	"github.com/google/uuid"
)

type Paciente struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	Nome      string    `json:"nome"`
	Avatar    string    `json:"avatar"`
	Idade     int64     `json:"idade"`
}

func NewPaciente(nome string, idade int64, avatar string) Paciente {
	return Paciente{
		Id:        uuid.NewString(),
		Nome:      nome,
		Avatar:    avatar,
		Idade:     idade,
		CreatedAt: time.Now(),
	}
}

func validatePacienteId(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return pacienteErrors.ErrIdInvalido.WithRootCause(err)
	}

	return nil
}

func validatePacienteNome(nome string) error {
	if strings.TrimSpace(nome) == "" {
		return pacienteErrors.ErrNomeObrigatorio
	}

	return nil
}

func validatePacienteIdade(idade int64) error {
	if idade <= 0 {
		return pacienteErrors.ErrIdadeInvalida
	}

	return nil
}

func (u Paciente) ValidateNew() error {
	if err := validatePacienteId(u.Id); err != nil {
		return err
	}

	if err := validatePacienteNome(u.Nome); err != nil {
		return err
	}

	if err := validatePacienteIdade(u.Idade); err != nil {
		return err
	}

	return nil
}
