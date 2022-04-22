package tutorial_service

import "pp-backend/models"

type Tutorial struct {
	ID      int
	Title   string
	Content string
	Type    int
}

func (t *Tutorial) Add() error {
	tutorial := map[string]interface{}{
		"title":   t.Title,
		"content": t.Content,
		"type":    t.Type,
	}
	if err := models.AddTutorial(tutorial); err != nil {
		return err
	}

	return nil
}

func (t *Tutorial) Edit() error {
	return models.EditTutorial(t.ID, map[string]interface{}{
		"title":   t.Title,
		"content": t.Content,
		"type":    t.Type,
	})
}

func (t *Tutorial) Get() (*models.Tutorial, error) {

	tutorial, err := models.GetTutorial(t.ID)
	if err != nil {
		return nil, err
	}

	return tutorial, nil
}

func (t *Tutorial) GetAll() ([]*models.Tutorial, error) {
	var tutorials []*models.Tutorial

	tutorials, err := models.GetTutorials(t.getMaps())
	if err != nil {
		return nil, err
	}

	return tutorials, nil
}

func (t *Tutorial) Delete() error {
	return models.DeleteTutorial(t.ID)
}

func (t *Tutorial) ExistByID() (bool, error) {
	return models.ExistTutorialByID(t.ID)
}

func (t *Tutorial) Count() (int, error) {
	return models.GetTutorialTotal(t.getMaps())
}

func (t *Tutorial) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	return maps
}

func (t *Tutorial) SearchTutorials() ([]*models.Tutorial, error) {
	var tutorials []*models.Tutorial

	tutorials, err := models.SearchTutorials(t.Title)
	if err != nil {
		return nil, err
	}

	return tutorials, nil
}
