package service

import (
	"awesomeProject/entity"
	"errors"
)

var Users []entity.User

func InitializeUser() {
	Users = []entity.User{}
}

func GetAllUsers() []entity.User {
	return Users
}

func InsertUser(user entity.User) entity.User {
	user.ID = uint64(len(Users) + 1)
	Users = append(Users, user)
	return user
}

func GetUserByID(id uint64) (entity.User, error) {
	for i := range Users {
		if Users[i].ID == id {
			return Users[i], nil
		}
	}
	return entity.User{}, errors.New("user not found")
}

func UpdateUserByID(user entity.User) (entity.User, error) {
	for i := range Users {
		if Users[i].ID == user.ID {
			Users[i] = user
			return Users[i], nil
		}
	}
	return entity.User{}, errors.New("user not found")
}

func DeleteUserByID(id uint64) error {
	for i := range Users {
		if Users[i].ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
