package services

import (
	"github.com/rof20004/healthy-api/application/domains/paciente/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/paciente/ports/output"
)

type GetAllPacientesService struct {
	pacientePersistenceOutputPort outputPorts.PacientePersistenceOutputPort
}

func NewGetAllPacientesService(pacientePersistenceOutputPort outputPorts.PacientePersistenceOutputPort) GetAllPacientesService {
	return GetAllPacientesService{pacientePersistenceOutputPort}
}

func (s GetAllPacientesService) GetAllPacientes() ([]entities.Paciente, error) {
	pacientes, err := s.pacientePersistenceOutputPort.FindAll()
	if err != nil {
		return pacientes, err
	}

	return pacientes, nil
}
