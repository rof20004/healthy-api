package paciente

import (
	"database/sql"
	"errors"

	pacienteEntities "github.com/rof20004/healthy-api/application/domains/paciente/entities"
	pacienteErrors "github.com/rof20004/healthy-api/application/domains/paciente/errors"
)

type PostgreSQLPacientePersistenceAdapter struct {
	db *sql.DB
}

func NewPostgreSQLPacientePersistenceAdapter(db *sql.DB) PostgreSQLPacientePersistenceAdapter {
	return PostgreSQLPacientePersistenceAdapter{db}
}

func (p PostgreSQLPacientePersistenceAdapter) SavePaciente(paciente pacienteEntities.Paciente) error {
	dml := "INSERT INTO pacientes(id, nome, avatar, idade, created_at) VALUES($1, $2, $3, $4, $5)"

	var (
		id        = paciente.Id
		nome      = paciente.Nome
		avatar    = paciente.Avatar
		idade     = paciente.Idade
		createdAt = paciente.CreatedAt
	)

	if _, err := p.db.Exec(dml, id, nome, avatar, idade, createdAt); err != nil {
		return err
	}

	return nil
}

func (p PostgreSQLPacientePersistenceAdapter) FindAll() ([]pacienteEntities.Paciente, error) {
	dml := `SELECT id, nome, avatar, idade, created_at FROM pacientes`

	var (
		id        sql.NullString
		nome      sql.NullString
		avatar    sql.NullString
		idade     sql.NullInt64
		createdAt sql.NullTime
	)

	pacientes := make([]pacienteEntities.Paciente, 0)

	rows, err := p.db.Query(dml)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pacientes, pacienteErrors.ErrPacientesNaoEncontrados.WithRootCause(err)
		}

		return pacientes, err
	}

	for rows.Next() {
		if err := rows.Scan(&id, &nome, &avatar, &idade, &createdAt); err != nil {
			return pacientes, err
		}

		paciente := pacienteEntities.Paciente{
			Id:        id.String,
			Nome:      nome.String,
			Avatar:    avatar.String,
			Idade:     idade.Int64,
			CreatedAt: createdAt.Time,
		}

		pacientes = append(pacientes, paciente)
	}

	return pacientes, nil
}
