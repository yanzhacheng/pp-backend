package tag_service

import (
	"pp-backend/models"
)

type Tag struct {
	ID        int
	TagName   string
	CreatedBy string
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.TagName)
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Add() error {
	return models.AddTag(t.TagName)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["tag_name"] = t.TagName

	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) GetAll() ([]models.Tag, error) {
	var tags []models.Tag

	tags, err := models.GetTags(t.getMaps())
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.TagName != "" {
		maps["name"] = t.TagName
	}

	return maps
}
