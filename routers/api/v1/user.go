package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"pp-backend/pkg/app"
	"pp-backend/pkg/e"
	"pp-backend/service/user_service"
)

type LoginForm struct {
	Username string `form:"username" valid:"Required;MaxSize(12)"`
	Password string `form:"password" valid:"Required;MaxSize(12)"`
}

// @Summary User login
// @Tags User
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /user/login [post]
func Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form LoginForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := user_service.User{
		Username: form.Username,
		Password: form.Password,
	}

	exist, err := userService.Login()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	if !exist {
		appG.Response(http.StatusBadRequest, e.LOGIN_ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Get user info
// @Tags User
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /user [get]
func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	userService := user_service.User{ID: id}
	exists, err := userService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	user, err := userService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, user)

}

type AddUserForm struct {
	Username string `form:"username" valid:"Required;MaxSize(12)"`
	Password string `form:"password" valid:"Required;MaxSize(12)"`
}

// @Summary Add User
// @Tags User
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /user [post]
func AddUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddUserForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := user_service.User{
		Username: form.Username,
		Password: form.Password,
	}
	if err := userService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditUserForm struct {
	ID       int    `form:"id" valid:"Required;Min(1)"`
	Username string `form:"username" valid:"Required;MaxSize(12)"`
	Password string `form:"password" valid:"Required;MaxSize(12)"`
}

// @Summary Update user info
// @Tags User
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /user [put]
func EditUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditUserForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := user_service.User{
		ID:       form.ID,
		Username: form.Username,
		Password: form.Password,
	}
	exists, err := userService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = userService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
