package input

import "github.com/rof20004/healthy-api/application/domains/paciente/entities"

type GetAllPacientesInputPort interface {
	GetAllPacientes() ([]entities.Paciente, error)
}
