package output

import "github.com/rof20004/healthy-api/application/domains/profissional/entities"

type ProfissionalPersistenceOutputPort interface {
	SaveProfissional(profissional entities.Profissional) error
	FindAll() ([]entities.Profissional, error)
}
