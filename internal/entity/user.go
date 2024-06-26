package entity

import (
	"errors"
)

type User struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Email_verify bool   `json:"email_varify"`
}

func NewUser(name, password, email string) *User {
	return &User{
		Name:     name,
		Password: password,
		Email:    email,
	}
}

func (u *User) Validate() error {
	// Add validation logic here
	// For example, you can check if the name, password, and email are not empty
	if u.Name == "" {
		return errors.New("name cannot be empty")
	}
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	// Add more validation rules as needed
	return nil
}
