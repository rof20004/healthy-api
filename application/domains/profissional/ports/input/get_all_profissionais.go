package input

import "github.com/rof20004/healthy-api/application/domains/profissional/entities"

type GetAllProfissionaisInputPort interface {
	GetAllProfissionais() ([]entities.Profissional, error)
}
