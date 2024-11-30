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
	// Validar o nome completo
	if c.FullName == "" || len(c.FullName) < 3 {
		return errors.New("o nome completo deve conter pelo menos 3 caracteres")
	}

	// Validar o telefone (formato básico)
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	if !phoneRegex.MatchString(c.Phone) {
		return errors.New("o telefone não é válido. Deve incluir o código do país e conter apenas números")
	}

	// Validar o email
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(c.Email) {
		return errors.New("o email não é válido")
	}

	// Validar o documento (formato básico para CPF ou CNPJ)
	docRegex := regexp.MustCompile(`^\d{11}$|^\d{14}$`)
	if !docRegex.MatchString(c.Document) {
		return errors.New("o documento deve ser um CPF (11 dígitos) ou CNPJ (14 dígitos)")
	}

	// Validar a senha (mínimo 6 caracteres, incluindo um número)
	passwordRegex := regexp.MustCompile(`^(?=.*[0-9]).{6,}$`)
	if !passwordRegex.MatchString(c.Password) {
		return errors.New("a senha deve ter pelo menos 6 caracteres e incluir ao menos um número")
	}

	// Validar a data de nascimento
	if c.Birthday.After(time.Now()) {
		return errors.New("a data de nascimento não pode ser no futuro")
	}

	// Validar se o usuário tem pelo menos 18 anos
	if time.Since(c.Birthday) < time.Hour*24*365*18 {
		return errors.New("o usuário deve ter pelo menos 18 anos")
	}

	return nil
}
