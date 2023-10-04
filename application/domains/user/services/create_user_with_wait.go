package services

import (
	"log"
	"time"

	"github.com/rof20004/healthy-api/application/domains/user/entities"
	outputPorts "github.com/rof20004/healthy-api/application/domains/user/ports/output"
)

type CreateUserWithWaitService struct {
	userPersistenceOutputPort outputPorts.UserPersistenceOutputPort
}

func NewCreateUserWithWaitService(userPersistenceOutputPort outputPorts.UserPersistenceOutputPort) CreateUserWithWaitService {
	return CreateUserWithWaitService{userPersistenceOutputPort}
}

func (s CreateUserWithWaitService) CreateUser(user entities.User) error {
	log.Println("Waiting to save user...")

	time.Sleep(time.Minute * 2)

	if err := user.ValidateNew(); err != nil {
		return err
	}

	if err := s.userPersistenceOutputPort.SaveUser(user); err != nil {
		return err
	}

	return nil
}
