package input

import "github.com/rof20004/healthy-api/application/domains/paciente/entities"

type CreatePacienteInputPort interface {
	CreatePaciente(paciente entities.Paciente) error
}
