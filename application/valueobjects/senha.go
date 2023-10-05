package valueobjects

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Senha struct {
	Value string
}

func NewSenha(senha string) Senha {
	if strings.TrimSpace(senha) == "" {
		return Senha{}
	}

	s, _ := gerarHashSenha(senha)

	return Senha{
		Value: s,
	}
}

func ValueOf(senha string) Senha {
	return Senha{
		Value: senha,
	}
}

func (e Senha) Verify(senha string) bool {
	return verificarSenha(e.Value, senha)
}

func (e Senha) String() string {
	return e.Value
}

func (e Senha) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, e.Value)), nil
}

func gerarHashSenha(senha string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func verificarSenha(senha string, hashSenha string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashSenha), []byte(senha))
	return err == nil
}
