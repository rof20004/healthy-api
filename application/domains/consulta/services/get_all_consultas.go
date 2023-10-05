package services

import (
	"github.com/rof20004/healthy-api/application/domains/consulta/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/consulta/ports/output"
)

type GetAllConsultasService struct {
	consultaPersistenceOutputPort outputPorts.ConsultaPersistenceOutputPort
}

func NewGetAllProfissionaisService(consultaPersistenceOutputPort outputPorts.ConsultaPersistenceOutputPort) GetAllConsultasService {
	return GetAllConsultasService{consultaPersistenceOutputPort}
}

func (s GetAllConsultasService) GetAllConsultasByPacienteId(pacienteId string) ([]entities.Consulta, error) {
	consultas, err := s.consultaPersistenceOutputPort.FindAllByPacienteId(pacienteId)
	if err != nil {
		return consultas, err
	}

	return consultas, nil
}
