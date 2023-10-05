package profissional_agenda

import (
	"time"

	profissionalAgendaEntities "github.com/rof20004/healthy-api/application/domains/profissional_agenda/entities"
)

type CreateProfissionalAgendaRequest struct {
	ProfissionalId string    `json:"profissionalId"`
	Data           time.Time `json:"data"`
}

func (c CreateProfissionalAgendaRequest) ToEntity() profissionalAgendaEntities.ProfissionalAgenda {
	return profissionalAgendaEntities.NewProfissionalAgenda(c.ProfissionalId, c.Data)
}
