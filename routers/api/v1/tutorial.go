package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"pp-backend/pkg/app"
	"pp-backend/pkg/e"
	"pp-backend/service/tutorial_service"
)

// @Summary Get a single tutorial
// @Tags Tutorial
// @Produce  json
// @Param id query int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tutorial [get]
func GetTutorial(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tutorialService := tutorial_service.Tutorial{ID: id}
	exists, err := tutorialService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	tutorial, err := tutorialService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, tutorial)
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
// @Router /tutorial [post]
func AddTutorial(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddTutorialForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tutorialService := tutorial_service.Tutorial{
		Title:   form.Title,
		Content: form.Content,
		Type:    form.Type,
	}
	if err := tutorialService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
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
// @Param id body int true "ID"
// @Param title body string false "Title"
// @Param content body string false "Content"
// @Param type body int false "Type"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tutorial [put]
func EditTutorial(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form EditTutorialForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tutorialService := tutorial_service.Tutorial{
		ID:      form.ID,
		Title:   form.Title,
		Content: form.Content,
		Type:    form.Type,
	}

	exists, err := tutorialService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = tutorialService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Delete tutorial
// @Tags Tutorial
// @Produce  json
// @Param id query int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tutorial [delete]
func DeleteTutorial(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	tutorialService := tutorial_service.Tutorial{ID: id}
	exists, err := tutorialService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = tutorialService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type SearchTutorialForm struct {
	Keyword string `form:"keyword" valid:"Required"`
}

// @Summary Search multiple tutorials
// @Tags Tutorial
// @Produce  json
// @Param keyword query string true "Search keyword"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tutorials/search [get]
func SearchTutorials(c *gin.Context) {
	var appG = app.Gin{C: c}

	keyword := c.Query("keyword")

	tutorialService := tutorial_service.Tutorial{
		Title: keyword,
	}

	tutorials, err := tutorialService.SearchTutorials()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = tutorials

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary Get multiple tutorials
// @Tags Tutorial
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /tutorials [get]
func GetTutorials(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tutorialService := tutorial_service.Tutorial{}

	tutorials, err := tutorialService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = tutorials
	//data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
