package profissional

import (
	"database/sql"
	"errors"

	profissionalEntities "github.com/rof20004/healthy-api/application/domains/profissional/entities"
	profissionalErrors "github.com/rof20004/healthy-api/application/domains/profissional/errors"
	"github.com/rof20004/healthy-api/application/valueobjects"
)

type PostgreSQLProfissionalPersistenceAdapter struct {
	db *sql.DB
}

func NewPostgreSQLProfissionalPersistenceAdapter(db *sql.DB) PostgreSQLProfissionalPersistenceAdapter {
	return PostgreSQLProfissionalPersistenceAdapter{db}
}

func (p PostgreSQLProfissionalPersistenceAdapter) SaveProfissional(profissional profissionalEntities.Profissional) error {
	dml := "INSERT INTO profissionais(id, nome, cpf, email, foto, crp, senha, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"

	var (
		id        = profissional.Id
		nome      = profissional.Nome
		cpf       = profissional.CPF
		email     = profissional.Email
		foto      = profissional.Foto
		crp       = profissional.CRP
		senha     = profissional.Senha
		createdAt = profissional.CreatedAt
	)

	if _, err := p.db.Exec(dml, id, nome, cpf, email.Value, foto, crp, senha.Value, createdAt); err != nil {
		return err
	}

	return nil
}

func (p PostgreSQLProfissionalPersistenceAdapter) FindAll() ([]profissionalEntities.Profissional, error) {
	dml := `SELECT id, nome, cpf, email, foto, crp, created_at FROM profissionais`

	var (
		id        sql.NullString
		nome      sql.NullString
		cpf       sql.NullString
		email     sql.NullString
		foto      sql.NullString
		crp       sql.NullString
		createdAt sql.NullTime
	)

	profissionais := make([]profissionalEntities.Profissional, 0)

	rows, err := p.db.Query(dml)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return profissionais, profissionalErrors.ErrProfissionaisNaoEncontrados.WithRootCause(err)
		}

		return profissionais, err
	}

	for rows.Next() {
		if err := rows.Scan(&id, &nome, &cpf, &email, &foto, &crp, &createdAt); err != nil {
			return profissionais, err
		}

		profissional := profissionalEntities.Profissional{
			Id:        id.String,
			Nome:      nome.String,
			CPF:       cpf.String,
			Email:     valueobjects.NewEmail(email.String),
			Foto:      foto.String,
			CRP:       crp.String,
			CreatedAt: createdAt.Time,
		}

		profissionais = append(profissionais, profissional)
	}

	return profissionais, nil
}
