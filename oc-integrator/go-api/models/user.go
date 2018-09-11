package models

import (
	"openshift_exp/oc-integrator/go-api/common"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func (user *User) Insert() error {
	err := DB.Create(user).Error
	return err
}

func (user *User) Update() error {
	return DB.Save(user).Error
}

func CheckAuth(userId string, pwd string) (string, bool) {
	var user User
	var err error
	err = DB.First(&user, "username = ?", userId).Error

	if err == nil {
		if user.Password == common.Md5(user.Email+pwd) {
			return user.Username, true
		}
	}
	return user.Username, false
}

func DelUserByID(id string) error {
	var user User
	err := DB.Where("id = ?", id).Delete(&user).Error
	return err
}
