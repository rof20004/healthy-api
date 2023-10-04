package input

import "github.com/rof20004/healthy-api/application/domains/user/entities"

type GetAllUsersInputPort interface {
	GetAllUsers() ([]entities.User, error)
}
