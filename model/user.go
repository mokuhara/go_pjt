package model

import "time"

type User struct {
	Id int64 `xorm:"pk autoincr int(64)" form:"id" json"id"`
	Email string `xorm:"varchar(40) not null unique" form:"email" json:"email"`
	Password string `xorm:"varchar(40)" form:"password" json:"password"`
	Type int64 `xorm:"int(64)" form:"type" json"type"`
	Version       uint64    `xorm:"version"`
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
	DeletedAt     time.Time `xorm:"deleted"`
}