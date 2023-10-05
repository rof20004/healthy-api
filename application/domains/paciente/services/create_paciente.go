package services

import (
	"github.com/rof20004/healthy-api/application/domains/paciente/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/paciente/ports/output"
)

type CreatePacienteService struct {
	pacientePersistenceOutputPort outputPorts.PacientePersistenceOutputPort
}

func NewCreatePacienteService(pacientePersistenceOutputPort outputPorts.PacientePersistenceOutputPort) CreatePacienteService {
	return CreatePacienteService{pacientePersistenceOutputPort}
}

func (s CreatePacienteService) CreatePaciente(paciente entities.Paciente) error {
	if err := paciente.ValidateNew(); err != nil {
		return err
	}

	if err := s.pacientePersistenceOutputPort.SavePaciente(paciente); err != nil {
		return err
	}

	return nil
}
