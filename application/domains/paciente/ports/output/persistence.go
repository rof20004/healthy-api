package output

import "github.com/rof20004/healthy-api/application/domains/paciente/entities"

type PacientePersistenceOutputPort interface {
	SavePaciente(paciente entities.Paciente) error
	FindAll() ([]entities.Paciente, error)
}
