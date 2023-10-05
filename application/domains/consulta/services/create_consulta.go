package services

import (
	"github.com/rof20004/healthy-api/application/domains/consulta/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/consulta/ports/output"
)

type CreateConsultaService struct {
	consultaPersistenceOutputPort outputPorts.ConsultaPersistenceOutputPort
}

func NewCreateConsultaService(consultaPersistenceOutputPort outputPorts.ConsultaPersistenceOutputPort) CreateConsultaService {
	return CreateConsultaService{consultaPersistenceOutputPort}
}

func (s CreateConsultaService) CreateConsulta(consulta entities.Consulta) error {
	if err := consulta.ValidateNew(); err != nil {
		return err
	}

	if err := s.consultaPersistenceOutputPort.SaveConsulta(consulta); err != nil {
		return err
	}

	return nil
}
