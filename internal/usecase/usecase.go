package usecase

import (
	"encoding/json"
	"errors"

	"github.com/antony-ramos/cop-orm/internal/controllers/database"
	"github.com/antony-ramos/cop-orm/internal/entity"
)

type Usecase struct {
	DB database.Database
}

func New(db database.Database) *Usecase {
	return &Usecase{DB: db}
}

func (u *Usecase) CreateUser(body []byte) error {
	user := entity.User{}
	err := json.Unmarshal(body, &user)
	if err != nil {
		return errors.New("unmarshal body into entity user: " + err.Error())
	}

	err = user.Validate()
	if err != nil {
		return errors.New("validate entity user: " + err.Error())
	}

	err = u.DB.CreateUser(user)
	if err != nil {
		return errors.New("create user: " + err.Error())
	}

	return nil
}

func (u *Usecase) GetUser(userID string) ([]byte, error) {
	user, err := u.DB.GetUser(userID)
	if err != nil {
		return nil, errors.New("get user: " + err.Error())
	}

	body, err := json.Marshal(user)
	if err != nil {
		return nil, errors.New("marshal user into body: " + err.Error())
	}

	return body, nil
}

func (u *Usecase) DeleteUser(userID string) error {
	err := u.DB.DeleteUser(userID)
	if err != nil {
		return errors.New("delete user: " + err.Error())
	}

	return nil
}

func (u *Usecase) ModifyUser(userID string, body []byte) error {
	user := entity.User{}
	err := json.Unmarshal(body, &user)
	if err != nil {
		return errors.New("unmarshal body into entity user: " + err.Error())
	}

	err = user.Validate()
	if err != nil {
		return errors.New("validate entity user: " + err.Error())
	}

	err = u.DB.ModifyUser(userID, user)
	if err != nil {
		return errors.New("modify user: " + err.Error())
	}

	return nil
}

func (u *Usecase) GetAllUsers() ([]byte, error) {
	users, err := u.DB.GetAllUsers()
	if err != nil {
		return nil, errors.New("get all users: " + err.Error())
	}

	body, err := json.Marshal(users)
	if err != nil {
		return nil, errors.New("marshal users into body: " + err.Error())
	}

	return body, nil
}

func (u *Usecase) CreateGroup(body []byte) error {
	group := entity.Group{}
	err := json.Unmarshal(body, &group)
	if err != nil {
		return errors.New("unmarshal body into entity group: " + err.Error())
	}
	err = group.Validate()
	if err != nil {
		return errors.New("validate entity group: " + err.Error())
	}
	err = u.DB.CreateGroup(group)
	if err != nil {
		return errors.New("create group: " + err.Error())
	}
	return nil
}
