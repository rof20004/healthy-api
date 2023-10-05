package services

import (
	"github.com/rof20004/healthy-api/application/domains/profissional_agenda/dtos"
	"github.com/rof20004/healthy-api/application/domains/profissional_agenda/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/profissional_agenda/ports/output"
)

type GetAllProfissionalAgendaService struct {
	profissionalAgendaPersistenceOutputPort outputPorts.ProfissionalAgendaPersistenceOutputPort
}

func NewGetAllProfissionalAgendaService(profissionalAgendaPersistenceOutputPort outputPorts.ProfissionalAgendaPersistenceOutputPort) GetAllProfissionalAgendaService {
	return GetAllProfissionalAgendaService{profissionalAgendaPersistenceOutputPort}
}

func (s GetAllProfissionalAgendaService) GetAllProfissionalAgendas() ([]entities.ProfissionalAgenda, error) {
	profissionalAgendas, err := s.profissionalAgendaPersistenceOutputPort.FindAll()
	if err != nil {
		return profissionalAgendas, err
	}

	return profissionalAgendas, nil
}

func (s GetAllProfissionalAgendaService) GetAllProfissionalAgendasByProfissionalId(profissionalId string) (dtos.AgendaDto, error) {
	profissionalAgendas, err := s.profissionalAgendaPersistenceOutputPort.FindAllByProfissionalId(profissionalId)
	if err != nil {
		return profissionalAgendas, err
	}

	return profissionalAgendas, nil
}
