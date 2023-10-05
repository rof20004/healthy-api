package entities

import (
	"strings"
	"time"

	profissionalAgendaErrors "github.com/rof20004/healthy-api/application/domains/profissional_agenda/errors"

	"github.com/google/uuid"
)

type ProfissionalAgenda struct {
	CreatedAt      time.Time `json:"createdAt"`
	Id             string    `json:"id"`
	ProfissionalId string    `json:"profissionalId"`
	Data           time.Time `json:"data"`
}

func NewProfissionalAgenda(profissionalId string, data time.Time) ProfissionalAgenda {
	return ProfissionalAgenda{
		Id:             uuid.NewString(),
		ProfissionalId: profissionalId,
		Data:           data,
		CreatedAt:      time.Now(),
	}
}

func validateProfissionalAgendaId(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return profissionalAgendaErrors.ErrIdInvalido.WithRootCause(err)
	}

	return nil
}

func validateProfissionalId(profissionalId string) error {
	if strings.TrimSpace(profissionalId) == "" {
		return profissionalAgendaErrors.ErrProfissionalIdObrigatorio
	}

	return nil
}

func validateProfissionalData(data time.Time) error {
	if data.IsZero() {
		return profissionalAgendaErrors.ErrDataInvalida
	}

	return nil
}

func (u ProfissionalAgenda) ValidateNew() error {
	if err := validateProfissionalAgendaId(u.Id); err != nil {
		return err
	}

	if err := validateProfissionalId(u.Id); err != nil {
		return err
	}

	if err := validateProfissionalData(u.Data); err != nil {
		return err
	}

	return nil
}
