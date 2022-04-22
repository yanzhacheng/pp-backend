package user_service

import "pp-backend/models"

type User struct {
	ID       int
	Username string
	Password string
}

func (u *User) Add() error {
	user := map[string]interface{}{
		"username": u.Username,
		"password": u.Password,
	}
	if err := models.AddUser(user); err != nil {
		return err
	}

	return nil
}

func (u *User) Edit() error {
	return models.EditUser(u.ID, map[string]interface{}{
		"username": u.Username,
		"password": u.Password,
	})
}

func (u *User) Get() (*models.User, error) {

	user, err := models.GetUser(u.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Delete() error {
	return models.DeleteUser(u.ID)
}

func (u *User) Login() (bool, error) {
	return models.ExistUserByUsernameAndPassword(u.Username, u.Password)
}

func (u *User) ExistByID() (bool, error) {
	return models.ExistUserByID(u.ID)
}
