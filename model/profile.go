package model

import "time"

type Profile struct {
	Id int64 `xorm:"pk autoincr int(64)" form:"id" json"id"`
	UserId int64 `xorm:"int(64)" form:"userId" json"userId"`
	Name string `xorm:"varchar(40)" form:"name" json:"name"`
	Description string `xorm:"varchar(511)" form:"description" json:"description"`
	Icon string `xorm:"varchar(255)" form:"icon" json:"icon"`
	//Version       uint64    `xorm:"version"`
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
	DeletedAt     time.Time `xorm:"deleted"`
}