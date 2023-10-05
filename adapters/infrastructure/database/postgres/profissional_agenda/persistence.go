package profissional_agenda

import (
	"database/sql"
	"errors"

	profissionalAgendaEntities "github.com/rof20004/healthy-api/application/domains/profissional_agenda/entities"
	profissionalAgendaErrors "github.com/rof20004/healthy-api/application/domains/profissional_agenda/errors"
)

type PostgreSQLProfissionalAgendaPersistenceAdapter struct {
	db *sql.DB
}

func NewPostgreSQLProfissionalAgendaPersistenceAdapter(db *sql.DB) PostgreSQLProfissionalAgendaPersistenceAdapter {
	return PostgreSQLProfissionalAgendaPersistenceAdapter{db}
}

func (p PostgreSQLProfissionalAgendaPersistenceAdapter) SaveProfissionalAgenda(profissionalAgenda profissionalAgendaEntities.ProfissionalAgenda) error {
	dml := "INSERT INTO profissionais(id, profissional_id, data, created_at) VALUES($1, $2, $3, $4)"

	var (
		id             = profissionalAgenda.Id
		profissionalId = profissionalAgenda.ProfissionalId
		data           = profissionalAgenda.Data
		createdAt      = profissionalAgenda.CreatedAt
	)

	if _, err := p.db.Exec(dml, id, profissionalId, data, createdAt); err != nil {
		return err
	}

	return nil
}

func (p PostgreSQLProfissionalAgendaPersistenceAdapter) FindAll() ([]profissionalAgendaEntities.ProfissionalAgenda, error) {
	dml := `SELECT id, profissional_id, data, created_at FROM profissionalAgendas`

	var (
		id             sql.NullString
		profissionalId sql.NullString
		data           sql.NullTime
		createdAt      sql.NullTime
	)

	profissionalAgendas := make([]profissionalAgendaEntities.ProfissionalAgenda, 0)

	rows, err := p.db.Query(dml)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return profissionalAgendas, profissionalAgendaErrors.ErrProfissionalAgendasNaoEncontradas.WithRootCause(err)
		}

		return profissionalAgendas, err
	}

	for rows.Next() {
		if err := rows.Scan(&id, &profissionalId, &data, &createdAt); err != nil {
			return profissionalAgendas, err
		}

		profissionalAgenda := profissionalAgendaEntities.ProfissionalAgenda{
			Id:             id.String,
			ProfissionalId: profissionalId.String,
			Data:           data.Time,
			CreatedAt:      createdAt.Time,
		}

		profissionalAgendas = append(profissionalAgendas, profissionalAgenda)
	}

	return profissionalAgendas, nil
}
