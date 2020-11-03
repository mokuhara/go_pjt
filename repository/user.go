package repository

import (
	"errors"
	"fmt"
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

func (UserRepository) GetAll() ([]model.User, error) {
	users := make([]model.User, 0)
	err := DbEngine.Asc("id").Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (UserRepository) Update(editUser *model.User) error {
	//なぜかUpdateが効かない、謎
	affected, err := DbEngine.ID(editUser.Id).Cols("type").Update(editUser.Type)
	if affected == 0 {
		return fmt.Errorf("can't update id: %d, editUser: %v", editUser.Id, *editUser)
	}
	if err != nil {
		return fmt.Errorf("faised update user id: %d, editUser: %v", editUser.Id, *editUser)
	}
	return nil
}

func (UserRepository) Delete(id int64) error {
	user := new(model.User)
	_, err := DbEngine.Id(id).Delete(user)
	if err != nil {
		return err
	}
	return nil
}
