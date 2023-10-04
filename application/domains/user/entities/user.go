package entities

import (
	"strings"
	"time"

	userErrors "github.com/rof20004/healthy-api/application/domains/user/errors"
	"github.com/rof20004/healthy-api/application/valueobjects"

	"github.com/google/uuid"
)

type User struct {
	CreatedAt time.Time          `json:"createdAt"`
	Id        string             `json:"id"`
	Name      string             `json:"name"`
	Email     valueobjects.Email `json:"email"`
	Age       int64              `json:"age"`
}

func NewUser(name string, age int64, email valueobjects.Email) User {
	return User{
		Id:        uuid.NewString(),
		Name:      name,
		Age:       age,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

func validateUserId(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return userErrors.ErrInvalidId.WithRootCause(err)
	}

	return nil
}

func validateUserName(name string) error {
	if strings.TrimSpace(name) == "" {
		return userErrors.ErrNameIsRequired
	}

	return nil
}

func validateUserAge(age int64) error {
	if age <= 0 {
		return userErrors.ErrInvalidAge
	}

	return nil
}

func validateUserEmail(email valueobjects.Email) error {
	if strings.TrimSpace(email.Value) == "" {
		return userErrors.ErrEmailIsRequired
	}

	if ok := email.IsValid(); !ok {
		return userErrors.ErrInvalidEmail
	}

	return nil
}

func (u User) ValidateNew() error {
	if err := validateUserId(u.Id); err != nil {
		return err
	}

	if err := validateUserName(u.Name); err != nil {
		return err
	}

	if err := validateUserAge(u.Age); err != nil {
		return err
	}

	if err := validateUserEmail(u.Email); err != nil {
		return err
	}

	return nil
}
