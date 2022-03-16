package models

import "github.com/jinzhu/gorm"

type Match struct {
	Model

	P1Name  string `json:"p1_name"`
	P1Score int    `json:"p1_score"`
	P2Name  string `json:"p2_name"`
	P2Score int    `json:"p2_score"`
}

// AddMatch Add a Match
func AddMatch(p1Name string, p1Score int, p2Name string, p2Score int) error {
	match := Match{
		P1Name:  p1Name,
		P1Score: p1Score,
		P2Name:  p2Name,
		P2Score: p2Score,
	}
	if err := db.Create(&match).Error; err != nil {
		return err
	}

	return nil
}

// GetMatches gets a list of matches based on paging and constraints
func GetMatches(pageNum int, pageSize int, maps interface{}) ([]Match, error) {
	var (
		matches []Match
		err     error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&matches).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&matches).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return matches, nil
}

// GetMatchTotal counts the total number of matches based on the constraint
func GetMatchTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Match{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistMatchByID determines whether a Match exists based on the ID
func ExistMatchByID(id int) (bool, error) {
	var match Match
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&match).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if match.ID > 0 {
		return true, nil
	}

	return false, nil
}

// DeleteMatch delete a match
func DeleteMatch(id int) error {
	if err := db.Where("id = ?", id).Delete(&Match{}).Error; err != nil {
		return err
	}

	return nil
}

// EditMatch modify a single match
func EditMatch(id int, data interface{}) error {
	if err := db.Model(&Match{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllMatch clear all match
func CleanAllMatch() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Match{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
