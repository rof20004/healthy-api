package consulta

import (
	consultaEntities "github.com/rof20004/healthy-api/application/domains/consulta/entities"
	"time"
)

type CreateConsultaRequest struct {
	PacienteId     string `json:"pacienteId"`
	ProfissionalId string `json:"profissionalId"`
	Data           string `json:"data" format:"01/02/2003 14:00"`
}

func (c CreateConsultaRequest) ToEntity() consultaEntities.Consulta {
	data, err := time.Parse("02/01/2006 15:04", c.Data)
	if err != nil {
		panic(err)
	}
	
	return consultaEntities.NewConsulta(c.PacienteId, c.ProfissionalId, data)
}
