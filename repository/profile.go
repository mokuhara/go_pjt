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
	has, err := DbEngine.Where("user_id = ?", userId).Desc("id").Get(&profile)
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

func (ProfileRepository) GetAll() ([]map[string]string, error) {
	//profiles := make([]model.Profile, 0)
	//err := DbEngine.Asc("id").Find(&profiles)
	result, err := DbEngine.Query("SELECT * FROM (SELECT *, rank() over(partition by user_id order by id desc) AS rank FROM \"public\".\"profile\" LIMIT 100 ) AS a WHERE rank = 1")
	if err != nil {
		return nil, err
	}
	var arr []map[string]string
	for _, s := range result {
		raw := make(map[string]string, 0)
		for k, v := range s {
			raw[k] = string(v)
		}
		arr = append(arr, raw)
	}
	return arr, nil
}
