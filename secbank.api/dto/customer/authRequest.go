package dto

import (
	"errors"
	"regexp"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a AuthRequest) Validate() error {
	// Validar o email
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(a.Email) {
		return errors.New("o email não é válido.")
	}

	// Validar a senha
	if len(a.Password) < 1 {
		return errors.New("A senha não foi preenchida.")
	}

	return nil
}
