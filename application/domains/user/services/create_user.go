package services

import (
	"github.com/rof20004/healthy-api/application/domains/user/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/user/ports/output"
)

type CreateUserService struct {
	userPersistenceOutputPort outputPorts.UserPersistenceOutputPort
}

func NewCreateUserService(userPersistenceOutputPort outputPorts.UserPersistenceOutputPort) CreateUserService {
	return CreateUserService{userPersistenceOutputPort}
}

func (s CreateUserService) CreateUser(user entities.User) error {
	if err := user.ValidateNew(); err != nil {
		return err
	}

	if err := s.userPersistenceOutputPort.SaveUser(user); err != nil {
		return err
	}

	return nil
}
