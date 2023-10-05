package consulta

import (
	"time"

	consultaEntities "github.com/rof20004/healthy-api/application/domains/consulta/entities"
)

type CreateConsultaRequest struct {
	PacienteId     string    `json:"pacienteId"`
	ProfissionalId string    `json:"profissionalId"`
	Data           time.Time `json:"data"`
}

func (c CreateConsultaRequest) ToEntity() consultaEntities.Consulta {
	return consultaEntities.NewConsulta(c.PacienteId, c.ProfissionalId, c.Data)
}
