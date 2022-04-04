package models

import "github.com/jinzhu/gorm"

type Tutorial struct {
	Model

	Title   string `json:"title"`
	Type    int    `json:"type"`
	Content string `json:"content"`
}

// ExistTutorialByID checks if a tutorial exists based on ID
func ExistTutorialByID(id int) (bool, error) {
	var tutorial Tutorial
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&tutorial).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tutorial.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetTutorialTotal gets the total number of tutorials based on the constraints
func GetTutorialTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tutorial{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetTutorials gets a list of tutorials based on paging constraints
func GetTutorials(maps interface{}) ([]*Tutorial, error) {
	var tutorials []*Tutorial
	err := db.Where(maps).Find(&tutorials).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tutorials, nil
}

// GetTutorial Get a single tutorial based on ID
func GetTutorial(id int) (*Tutorial, error) {
	var tutorial Tutorial
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&tutorial).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &tutorial, nil
}

// EditTutorial modify a single tutorial
func EditTutorial(id int, data interface{}) error {
	if err := db.Model(&Tutorial{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddTutorial add a single Tutorial
func AddTutorial(data map[string]interface{}) error {
	tutorial := Tutorial{
		Title:   data["title"].(string),
		Content: data["content"].(string),
		Type:    data["type"].(int),
	}
	if err := db.Create(&tutorial).Error; err != nil {
		return err
	}

	return nil
}

// DeleteTutorial delete a single tutorial
func DeleteTutorial(id int) error {
	if err := db.Where("id = ?", id).Delete(Tutorial{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllTutorial clear all tutorial
func CleanAllTutorial() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tutorial{}).Error; err != nil {
		return err
	}

	return nil
}
