package usecase

import "errors"

type Usecase struct {
}

func New() *Usecase {
	return &Usecase{}
}

func (u *Usecase) CreateUser(body []byte) error {
	// Implement the logic to create a user using the provided body
	return errors.New("not implemented")
}

func (u *Usecase) GetUser(userID string) ([]byte, error) {
	// Implement the logic to get a user by userID
	return nil, errors.New("not implemented")
}

func (u *Usecase) DeleteUser(userID string) error {
	// Implement the logic to delete a user by userID
	return errors.New("not implemented")
}

func (u *Usecase) ModifyUser(userID string, body []byte) error {
	// Implement the logic to modify a user using the provided userID and body
	return errors.New("not implemented")
}
