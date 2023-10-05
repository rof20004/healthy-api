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
	dml := "INSERT INTO consultas(id, paciente_id, profissional_id, data, created_at) VALUES($1, $2, $3, $4, $5)"

	var (
		id             = consulta.Id
		pacienteId     = consulta.PacienteId
		profissionalId = consulta.ProfissionalId
		data           = consulta.Data
		createdAt      = consulta.CreatedAt
	)

	if _, err := p.db.Exec(dml, id, pacienteId, profissionalId, data, createdAt); err != nil {
		return err
	}

	return nil
}

func (p PostgreSQLConsultaPersistenceAdapter) FindAllByPacienteId(pacienteId string) ([]consultaEntities.Consulta, error) {
	dml := `SELECT id, profissional_id, data, created_at FROM consultas WHERE paciente_id = $1`

	var (
		id             sql.NullString
		profissionalId sql.NullString
		data           sql.NullTime
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
		if err := rows.Scan(&id, &profissionalId, &data, &createdAt); err != nil {
			return consultas, err
		}

		consulta := consultaEntities.Consulta{
			Id:             id.String,
			PacienteId:     pacienteId,
			ProfissionalId: profissionalId.String,
			Data:           data.Time,
			CreatedAt:      createdAt.Time,
		}

		consultas = append(consultas, consulta)
	}

	return consultas, nil
}
