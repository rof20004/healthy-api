package user

import (
	"time"

	userEntities "github.com/rof20004/healthy-api/application/domains/user/entities"
	"github.com/rof20004/healthy-api/application/valueobjects"
)

type CreateUserRequest struct {
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int64     `json:"age"`
}

func (c CreateUserRequest) ToEntity() userEntities.User {
	email := valueobjects.NewEmail(c.Email)
	return userEntities.NewUser(c.Name, c.Age, email)
}
