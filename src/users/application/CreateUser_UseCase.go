package application

import (
	"vitalPoint/src/users/domain"

	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	repo domain.IUser
}

func NewCreateUser(repo domain.IUser) *CreateUser {
	return &CreateUser{repo: repo}
}

func (cu *CreateUser) Execute(userName string, email string, password string, role string) error {
	// Generar hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Guardar usuario con el ESP32ID
	err = cu.repo.SaveUser(userName, email, string(hashedPassword), role)
	if err != nil {
		return err
	}

	return nil
}
