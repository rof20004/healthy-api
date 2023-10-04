package user

import (
	"database/sql"
	"errors"

	userEntities "github.com/rof20004/healthy-api/application/domains/user/entities"
	userErrors "github.com/rof20004/healthy-api/application/domains/user/errors"
	"github.com/rof20004/healthy-api/application/valueobjects"
)

type PostgreSQLUserPersistenceAdapter struct {
	db *sql.DB
}

func NewPostgreSQLUserPersistenceAdapter(db *sql.DB) PostgreSQLUserPersistenceAdapter {
	return PostgreSQLUserPersistenceAdapter{db}
}

func (p PostgreSQLUserPersistenceAdapter) SaveUser(user userEntities.User) error {
	dml := "INSERT INTO users(id, name, age, email, created_at) VALUES($1, $2, $3, $4, $5)"

	var (
		id        = user.Id
		name      = user.Name
		age       = user.Age
		email     = user.Email
		createdAt = user.CreatedAt
	)

	if _, err := p.db.Exec(dml, id, name, age, email.Value, createdAt); err != nil {
		return err
	}

	return nil
}

func (p PostgreSQLUserPersistenceAdapter) FindAll() ([]userEntities.User, error) {
	dml := `SELECT id, name, age, email, created_at FROM users`

	var (
		id        sql.NullString
		name      sql.NullString
		age       sql.NullInt64
		email     sql.NullString
		createdAt sql.NullTime
	)

	users := make([]userEntities.User, 0)

	rows, err := p.db.Query(dml)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, userErrors.ErrUsersNotFound.WithRootCause(err)
		}

		return users, err
	}

	for rows.Next() {
		if err := rows.Scan(&id, &name, &age, &email, &createdAt); err != nil {
			return users, err
		}

		user := userEntities.User{
			Id:        id.String,
			Name:      name.String,
			Age:       age.Int64,
			Email:     valueobjects.NewEmail(email.String),
			CreatedAt: createdAt.Time,
		}

		users = append(users, user)
	}

	return users, nil
}
