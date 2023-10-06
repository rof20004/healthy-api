package consulta

import (
	"database/sql"
	"errors"

	consultaEntities "github.com/rof20004/healthy-api/application/domains/consulta/entities"
	consultaErrors "github.com/rof20004/healthy-api/application/domains/consulta/errors"
)

type PostgreSQLConsultaPersistenceAdapter struct {
	db *sql.DB
}

func NewPostgreSQLConsultaPersistenceAdapter(db *sql.DB) PostgreSQLConsultaPersistenceAdapter {
	return PostgreSQLConsultaPersistenceAdapter{db}
}

func (p PostgreSQLConsultaPersistenceAdapter) SaveConsulta(consulta consultaEntities.Consulta) error {
	dml := "INSERT INTO consultas(id, paciente_id, profissional_id, data, link, created_at) VALUES($1, $2, $3, $4, $5, $6)"

	var (
		id             = consulta.Id
		pacienteId     = consulta.PacienteId
		profissionalId = consulta.ProfissionalId
		data           = consulta.Data
		link           = consulta.Link
		createdAt      = consulta.CreatedAt
	)

	if _, err := p.db.Exec(dml, id, pacienteId, profissionalId, data, link, createdAt); err != nil {
		return err
	}

	return nil
}

func (p PostgreSQLConsultaPersistenceAdapter) FindAll() ([]consultaEntities.Consulta, error) {
	dml := `SELECT id, paciente_id, profissional_id, data, link, created_at FROM consultas`

	var (
		id             sql.NullString
		pacienteId     sql.NullString
		profissionalId sql.NullString
		data           sql.NullTime
		link           sql.NullString
		createdAt      sql.NullTime
	)

	consultas := make([]consultaEntities.Consulta, 0)

	rows, err := p.db.Query(dml)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return consultas, consultaErrors.ErrConsultasNaoEncontradas.WithRootCause(err)
		}

		return consultas, err
	}

	for rows.Next() {
		if err := rows.Scan(&id, &pacienteId, &profissionalId, &data, &link, &createdAt); err != nil {
			return consultas, err
		}

		consulta := consultaEntities.Consulta{
			Id:             id.String,
			PacienteId:     pacienteId.String,
			ProfissionalId: profissionalId.String,
			Data:           data.Time,
			Link:           link.String,
			CreatedAt:      createdAt.Time,
		}

		consultas = append(consultas, consulta)
	}

	return consultas, nil
}

func (p PostgreSQLConsultaPersistenceAdapter) FindAllByPacienteId(pacienteId string) ([]consultaEntities.Consulta, error) {
	dml := `SELECT id, profissional_id, data, link, created_at FROM consultas WHERE paciente_id = $1`

	var (
		id             sql.NullString
		profissionalId sql.NullString
		data           sql.NullTime
		link           sql.NullString
		createdAt      sql.NullTime
	)

	consultas := make([]consultaEntities.Consulta, 0)

	rows, err := p.db.Query(dml, pacienteId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return consultas, consultaErrors.ErrConsultasNaoEncontradas.WithRootCause(err)
		}

		return consultas, err
	}

	for rows.Next() {
		if err := rows.Scan(&id, &profissionalId, &data, &link, &createdAt); err != nil {
			return consultas, err
		}

		consulta := consultaEntities.Consulta{
			Id:             id.String,
			PacienteId:     pacienteId,
			ProfissionalId: profissionalId.String,
			Data:           data.Time,
			Link:           link.String,
			CreatedAt:      createdAt.Time,
		}

		consultas = append(consultas, consulta)
	}

	return consultas, nil
}
