package users

import (
	"learning-golang/api.example.com/pkg/db"
)

type Users []User

type User db.Users

func (u *User) TableName() string {
	return "users"
}
