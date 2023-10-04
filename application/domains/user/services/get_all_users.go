package services

import (
	"github.com/rof20004/healthy-api/application/domains/user/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/user/ports/output"
)

type GetAllUsersService struct {
	userPersistenceOutputPort outputPorts.UserPersistenceOutputPort
}

func NewGetAllUsersService(userPersistenceOutputPort outputPorts.UserPersistenceOutputPort) GetAllUsersService {
	return GetAllUsersService{userPersistenceOutputPort}
}

func (s GetAllUsersService) GetAllUsers() ([]entities.User, error) {
	users, err := s.userPersistenceOutputPort.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}
