package repository

import (
	"errors"
	"goPjt/model"
)

type UserRepository struct {}

func (UserRepository) Create(user *model.User) error {
	_, err := DbEngine.Insert(user)
	if err!= nil{
		return err
	}
	return nil
}

func (UserRepository) Get(email string) (*model.User, error) {
	user := model.User{}
	has, err := DbEngine.Where("email = ?", email).Get(&user)
	if err != nil{
		return nil, err
	}
	if !has {
		return nil, errors.New("user is not fuound")
	}
	return &user, nil
}
