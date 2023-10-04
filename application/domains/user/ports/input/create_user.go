package input

import "github.com/rof20004/healthy-api/application/domains/user/entities"

type CreateUserInputPort interface {
	CreateUser(user entities.User) error
}
