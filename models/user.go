package models

import "github.com/jinzhu/gorm"

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
}

// ExistUserByUsernameAndPassword checks if a user exists based on username and password
func ExistUserByUsernameAndPassword(username string, password string) (bool, error) {
	var user User
	err := db.Where("username = ? AND password = ? ", username, password).First(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// EditUser modify a single user
func EditUser(id int, data interface{}) error {
	if err := db.Model(&User{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddUser add a single User
func AddUser(data map[string]interface{}) error {
	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// GetUser Get user info based on ID
func GetUser(id int) (*User, error) {
	var user User
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

// DeleteUser delete a single user
func DeleteUser(id int) error {
	if err := db.Where("id = ?", id).Delete(User{}).Error; err != nil {
		return err
	}

	return nil
}

// ExistUserByID checks if a user exists based on ID
func ExistUserByID(id int) (bool, error) {
	var user User
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
