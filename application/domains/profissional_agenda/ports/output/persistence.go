package output

import "github.com/rof20004/healthy-api/application/domains/profissional_agenda/entities"

type ProfissionalAgendaPersistenceOutputPort interface {
	SaveProfissionalAgenda(profissionalAgenda entities.ProfissionalAgenda) error
	FindAll() ([]entities.ProfissionalAgenda, error)
}
