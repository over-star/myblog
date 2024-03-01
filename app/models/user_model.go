package models

import "myblog/pkg/database"

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
