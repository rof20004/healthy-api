package output

import "github.com/rof20004/healthy-api/application/domains/consulta/entities"

type ConsultaPersistenceOutputPort interface {
	SaveConsulta(consulta entities.Consulta) error
	FindAll() ([]entities.Consulta, error)
	FindAllByPacienteId(pacienteId string) ([]entities.Consulta, error)
}
