package input

import "github.com/rof20004/healthy-api/application/domains/consulta/entities"

type GetAllConsultasInputPort interface {
	GetAllConsultasByPacienteId(pacienteId string) ([]entities.Consulta, error)
}
