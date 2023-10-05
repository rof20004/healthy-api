package input

import "github.com/rof20004/healthy-api/application/domains/profissional/entities"

type CreateProfissionalInputPort interface {
	CreateProfissionalAgenda(profissional entities.Profissional) error
}
