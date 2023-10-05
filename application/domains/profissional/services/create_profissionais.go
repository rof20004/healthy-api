package services

import (
	"github.com/rof20004/healthy-api/application/domains/profissional/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/profissional/ports/output"
)

type CreateProfissionalService struct {
	profissionalPersistenceOutputPort outputPorts.ProfissionalPersistenceOutputPort
}

func NewCreateProfissionalService(profissionalPersistenceOutputPort outputPorts.ProfissionalPersistenceOutputPort) CreateProfissionalService {
	return CreateProfissionalService{profissionalPersistenceOutputPort}
}

func (s CreateProfissionalService) CreateProfissionalAgenda(profissional entities.Profissional) error {
	if err := profissional.ValidateNew(); err != nil {
		return err
	}

	if err := s.profissionalPersistenceOutputPort.SaveProfissional(profissional); err != nil {
		return err
	}

	return nil
}
