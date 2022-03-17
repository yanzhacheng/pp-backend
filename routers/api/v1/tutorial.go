package v1

import "github.com/gin-gonic/gin"

// @Summary Get a single tutorial
// @Tags Tutorial
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tutorials/{id} [get]
func GetTutorial(c *gin.Context) {
}

// @Summary Search multiple tutorials
// @Tags Tutorial
// @Produce  json
// @Param key_words body int true "Search key words"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tutorials [get]
func SearchTutorials(c *gin.Context) {
}

type AddTutorialForm struct {
	Title   string `form:"title" valid:"Required;MaxSize(100)"`
	Content string `form:"content" valid:"Required;MaxSize(65535)"`
	Type    int    `form:"type" valid:"Range(0,1)"`
}

// @Summary Add tutorial
// @Tags Tutorial
// @Produce  json
// @Param title body string true "Title"
// @Param content body string true "Content"
// @Param type body int false "Type"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tutorials [post]
func AddTutorial(c *gin.Context) {
}

type EditTutorialForm struct {
	ID      int    `form:"id" valid:"Required;Min(1)"`
	Title   string `form:"title" valid:"Required;MaxSize(100)"`
	Content string `form:"content" valid:"Required;MaxSize(65535)"`
	Type    int    `form:"type" valid:"Range(0,1)"`
}

// @Summary Update tutorial
// @Tags Tutorial
// @Produce  json
// @Param id path int true "ID"
// @Param title body string false "Title"
// @Param content body string false "Content"
// @Param type body int false "Type"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tutorials/{id} [put]
func EditTutorial(c *gin.Context) {
}

// @Summary Delete tutorial
// @Tags Tutorial
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tutorials/{id} [delete]
func DeleteTutorial(c *gin.Context) {
}
