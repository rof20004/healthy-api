package output

import "github.com/rof20004/healthy-api/application/domains/user/entities"

type UserPersistenceOutputPort interface {
	SaveUser(user entities.User) error
	FindAll() ([]entities.User, error)
}
