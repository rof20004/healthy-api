package entities

import (
	"strings"
	"time"

	consultaErrors "github.com/rof20004/healthy-api/application/domains/consulta/errors"

	"github.com/google/uuid"
)

type Consulta struct {
	CreatedAt      time.Time `json:"createdAt"`
	Id             string    `json:"id"`
	PacienteId     string    `json:"pacienteId"`
	ProfissionalId string    `json:"profissionalId"`
	Data           time.Time `json:"data"`
}

func NewConsulta(pacienteId, profissionalId string, data time.Time) Consulta {
	return Consulta{
		Id:             uuid.NewString(),
		PacienteId:     pacienteId,
		ProfissionalId: profissionalId,
		Data:           data,
		CreatedAt:      time.Now(),
	}
}

func validateConsultaId(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return consultaErrors.ErrIdInvalido.WithRootCause(err)
	}

	return nil
}

func validatePacienteId(pacienteId string) error {
	if strings.TrimSpace(pacienteId) == "" {
		return consultaErrors.ErrPacienteIdObrigatorio
	}

	return nil
}

func validateProfissionalId(profissionalId string) error {
	if strings.TrimSpace(profissionalId) == "" {
		return consultaErrors.ErrProfissionalIdObrigatorio
	}

	return nil
}

func validateConsultaData(data time.Time) error {
	if data.IsZero() {
		return consultaErrors.ErrDataInvalida
	}

	return nil
}

func (u Consulta) ValidateNew() error {
	if err := validateConsultaId(u.Id); err != nil {
		return err
	}

	if err := validatePacienteId(u.PacienteId); err != nil {
		return err
	}

	if err := validateProfissionalId(u.ProfissionalId); err != nil {
		return err
	}

	if err := validateConsultaData(u.Data); err != nil {
		return err
	}

	return nil
}
