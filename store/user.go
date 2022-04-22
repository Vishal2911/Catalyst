package store

import (
	"magantas/catalyst/model"
)

func (store Store) CreateUser(user *model.Users) error {

	result := store.db.Create(&user)
	return result.Error

}

func (store Store) GetUser(user *model.Users) error {

	result := store.db.Select(" * ").Table("users").Where("email_id = $1", user.EmailID)
	return result.Error

}

func (store Store) GetUsers(users *[]model.Users) error {

	result := store.db.Find(&users)
	return result.Error

}

func (store Store) UpdateUser(user *model.Users) error {

	result := store.db.Save(&user)
	return result.Error

}

func (store Store) DeleteUser(user *model.Users) error {

	result := store.db.Delete(&user)
	return result.Error

}

func (store Store) GetUserByEmailAndPassword(user *model.Users) (*model.Users, error) {
	newuser := model.Users{}
	result := store.db.Select(" * ").Table("users").Where("email_id = $1", user.EmailID).Where("password = $2",user.Password)
	result.Scan(&newuser)
	return &newuser, result.Error

}
