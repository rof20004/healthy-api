package valueobjects

import (
	"fmt"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type Email struct {
	Value string
}

func NewEmail(email string) Email {
	return Email{
		Value: email,
	}
}

func (e Email) IsValid() bool {
	return emailRegex.MatchString(e.Value)
}

func (e Email) String() string {
	return e.Value
}

func (e Email) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, e.Value)), nil
}
