package v1

import (
	"github.com/gin-gonic/gin"
)

type AddTagForm struct {
	TagName string `form:"tag_name" valid:"Required;MaxSize(100)"`
}

// @Summary Add article tag
// @Produce  json
// @Param tag_name body string true "TagName"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {

}
