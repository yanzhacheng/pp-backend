package article_service

import (
	"pp-backend/models"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	Author        string
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"author":          a.Author,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}
	if err := models.AddArticle(article); err != nil {
		return err
	}

	return nil
}

func (a *Article) Edit() error {
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	})
}

func (a *Article) Get() (*models.Article, error) {

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	var articles []*models.Article

	articles, err := models.GetArticles(a.getMaps())
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}

	return maps
}

func (a *Article) SearchArticles() ([]*models.Article, error) {
	var articles []*models.Article

	articles, err := models.SearchArticles(a.Title)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *Article) GetAllByTag() ([]*models.Article, error) {
	var articles []*models.Article

	articles, err := models.GetArticlesByTag(a.TagID)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
