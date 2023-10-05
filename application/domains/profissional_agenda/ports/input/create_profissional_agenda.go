package input

import "github.com/rof20004/healthy-api/application/domains/profissional_agenda/entities"

type CreateProfissionalAgendaInputPort interface {
	CreateProfissionalAgenda(profissional entities.ProfissionalAgenda) error
}
