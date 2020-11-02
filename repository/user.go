package repository

import "goPjt/model"

type UserRepository struct {}

func (UserRepository) Create(user *model.User) error {
	_, err := DbEngine.Insert(user)
	if err!= nil{
		return err
	}
	return nil
}
