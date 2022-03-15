package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pp-backend/pkg/app"
	"pp-backend/pkg/e"
	"pp-backend/service/tag_service"
)

type AddTagForm struct {
	TagName string `form:"tag_name" valid:"Required;MaxSize(100)"`
}

//// @Summary Get multiple article tags
//// @Produce  json
//// @Param name query string false "Name"
//// @Param state query int false "State"
//// @Success 200 {object} app.Response
//// @Failure 500 {object} app.Response
//// @Router /api/v1/tags [get]
//func GetTags(c *gin.Context) {
//	appG := app.Gin{C: c}
//	name := c.Query("name")
//	state := -1
//	if arg := c.Query("state"); arg != "" {
//		state = com.StrTo(arg).MustInt()
//	}
//
//	tagService := tag_service.Tag{
//		Name:     name,
//		State:    state,
//		PageNum:  util.GetPage(c),
//		PageSize: setting.AppSetting.PageSize,
//	}
//	tags, err := tagService.GetAll()
//	if err != nil {
//		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
//		return
//	}
//
//	count, err := tagService.Count()
//	if err != nil {
//		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
//		return
//	}
//
//	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
//		"lists": tags,
//		"total": count,
//	})
//}

// @Summary Add article tag
// @Produce  json
// @Param tag_name body string true "TagName"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddTagForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{
		TagName: form.TagName,
	}
	exists, err := tagService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	err = tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
