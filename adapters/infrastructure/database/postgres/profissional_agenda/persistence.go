package profissional_agenda

import (
	"database/sql"
	"errors"

	profissionalAgendaDtos "github.com/rof20004/healthy-api/application/domains/profissional_agenda/dtos"
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
	dml := "INSERT INTO profissional_agenda(id, profissional_id, data, created_at) VALUES($1, $2, $3, $4)"

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
	dml := `SELECT id, profissional_id, data, created_at FROM profissional_agenda`

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

func (p PostgreSQLProfissionalAgendaPersistenceAdapter) FindAllByProfissionalId(profissionalId string) (profissionalAgendaDtos.AgendaDto, error) {
	dml := `SELECT DISTINCT
    				p.nome,
    				p.cpf,
    				p.email,
    				p.foto,
    				p.crp,
    				pa.data, 
    				pa.created_at 
			FROM profissional_agenda pa
			INNER JOIN profissionais p ON p.id = pa.profissional_id
			WHERE p.id = $1
			ORDER BY pa.data ASC`

	var (
		profissionalNome  sql.NullString
		profissionalCPF   sql.NullString
		profissionalEmail sql.NullString
		profissionalFoto  sql.NullString
		profissionalCRP   sql.NullString
		data              sql.NullTime
		createdAt         sql.NullTime
	)

	var profissionalAgendas profissionalAgendaDtos.AgendaDto

	rows, err := p.db.Query(dml, profissionalId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return profissionalAgendas, profissionalAgendaErrors.ErrProfissionalAgendasNaoEncontradas.WithRootCause(err)
		}

		return profissionalAgendas, err
	}

	for rows.Next() {
		if err := rows.Scan(&profissionalNome, &profissionalCPF, &profissionalEmail, &profissionalFoto, &profissionalCRP, &data, &createdAt); err != nil {
			return profissionalAgendas, err
		}

		if profissionalAgendas.Profissional.Nome == "" {
			profissionalAgendas.Profissional = profissionalAgendaDtos.ProfissionalDto{
				Id:    profissionalId,
				Nome:  profissionalNome.String,
				CPF:   profissionalCPF.String,
				Email: profissionalEmail.String,
				Foto:  profissionalFoto.String,
				CRP:   profissionalCRP.String,
			}
		}

		if len(profissionalAgendas.Datas) == 0 {
			profissionalAgendas.Datas = make([]profissionalAgendaDtos.DataDto, 0)
		}

		profissionalAgendas.Datas = append(profissionalAgendas.Datas, profissionalAgendaDtos.DataDto{
			Dia:  data.Time.Format("02/01/2006"),
			Hora: data.Time.Format("15:04"),
		})
	}

	return profissionalAgendas, nil
}
