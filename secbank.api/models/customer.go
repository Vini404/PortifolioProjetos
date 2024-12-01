package models

import (
	"errors"
	"regexp"
	"time"
)

type Customer struct {
	ID               int       `db:"id"`
	FullName         string    `db:"fullname"`
	Phone            string    `db:"phone"`
	Email            string    `db:"email"`
	Password         string    `db:"password"`
	Document         string    `db:"document"`
	Birthday         time.Time `db:"birthday"`
	CreatedTimeStamp time.Time `db:"createdtimestamp"`
	UpdatedTimeStamp time.Time `db:"updatedtimestamp"`
}

func (c Customer) Validate() error {
	if c.FullName == "" || len(c.FullName) < 3 {
		return errors.New("o nome completo deve conter pelo menos 3 caracteres")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(c.Email) {
		return errors.New("o email não é válido")
	}

	docRegex := regexp.MustCompile(`^\d{11}$|^\d{14}$`)
	if !docRegex.MatchString(c.Document) {
		return errors.New("o documento deve ser um CPF (11 dígitos) ou CNPJ (14 dígitos)")
	}

	if len(c.Password) < 6 {
		return errors.New("a senha deve ter pelo menos 6 caracteres")
	}

	containsDigit := regexp.MustCompile(`[0-9]`)
	if !containsDigit.MatchString(c.Password) {
		return errors.New("a senha deve incluir pelo menos um número")
	}

	if c.Birthday.After(time.Now()) {
		return errors.New("a data de nascimento não pode ser no futuro")
	}

	if time.Since(c.Birthday) < time.Hour*24*365*18 {
		return errors.New("o usuário deve ter pelo menos 18 anos")
	}

	return nil
}
