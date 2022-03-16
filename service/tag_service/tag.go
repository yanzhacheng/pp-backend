package tag_service

type Tag struct {
	ID      int
	TagName string

	PageNum  int
	PageSize int
}

//func (t *Tag) ExistByName() (bool, error) {
//	return models.ExistTagByName(t.TagName)
//}
//
//func (t *Tag) ExistByID() (bool, error) {
//	return models.ExistTagByID(t.ID)
//}
//
//func (t *Tag) Add() error {
//	return models.AddTag(t.TagName)
//}
