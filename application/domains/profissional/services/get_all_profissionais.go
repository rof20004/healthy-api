package services

import (
	"github.com/rof20004/healthy-api/application/domains/profissional/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/profissional/ports/output"
)

type GetAllProfissionaisService struct {
	profissionalPersistenceOutputPort outputPorts.ProfissionalPersistenceOutputPort
}

func NewGetAllProfissionaisService(profissionalPersistenceOutputPort outputPorts.ProfissionalPersistenceOutputPort) GetAllProfissionaisService {
	return GetAllProfissionaisService{profissionalPersistenceOutputPort}
}

func (s GetAllProfissionaisService) GetAllProfissionais() ([]entities.Profissional, error) {
	profissionais, err := s.profissionalPersistenceOutputPort.FindAll()
	if err != nil {
		return profissionais, err
	}

	return profissionais, nil
}
