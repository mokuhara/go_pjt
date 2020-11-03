package repository

import (
	"fmt"
	"goPjt/model"
)
type ProfileRepository struct {}

func (ProfileRepository) Create(profile *model.Profile) error {
	_, err := DbEngine.Insert(profile)
	if err!= nil{
		return err
	}
	return nil
}

func (ProfileRepository) Get(userId int64) (*model.Profile, error) {
	profile := model.Profile{}
	has, err := DbEngine.Where("user_id = ?", userId).Desc("updated_at").Get(&profile)
	if err != nil{
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("user is not fuound")
	}
	return &profile, nil
}

func (ProfileRepository) Update(editProfile *model.Profile) error {
	affected, err := DbEngine.ID(editProfile.Id).Update(editProfile)
	if affected == 0 {
		return fmt.Errorf("can't update id: %d, editUser: %v", editProfile.Id, *editProfile)
	}
	if err != nil {
		return fmt.Errorf("faised update user id: %d, editUser: %v", editProfile.Id, *editProfile)
	}
	return nil
}

func (ProfileRepository) Delete(id int64) error {
	profile := new(model.Profile)
	_, err := DbEngine.Id(id).Delete(profile)
	if err != nil {
		return err
	}
	return nil
}
