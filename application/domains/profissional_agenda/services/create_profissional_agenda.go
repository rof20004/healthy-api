package services

import (
	"github.com/rof20004/healthy-api/application/domains/profissional_agenda/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/profissional_agenda/ports/output"
)

type CreateProfissionalAgendaService struct {
	profissionalAgendaPersistenceOutputPort outputPorts.ProfissionalAgendaPersistenceOutputPort
}

func NewCreateProfissionalAgendaService(profissionalAgendaPersistenceOutputPort outputPorts.ProfissionalAgendaPersistenceOutputPort) CreateProfissionalAgendaService {
	return CreateProfissionalAgendaService{profissionalAgendaPersistenceOutputPort}
}

func (s CreateProfissionalAgendaService) CreateProfissionalAgenda(profissionalAgenda entities.ProfissionalAgenda) error {
	if err := profissionalAgenda.ValidateNew(); err != nil {
		return err
	}

	if err := s.profissionalAgendaPersistenceOutputPort.SaveProfissionalAgenda(profissionalAgenda); err != nil {
		return err
	}

	return nil
}
