package routers

import (
	"github.com/gin-gonic/gin"
	v1 "pp-backend/routers/api/v1"

	_ "pp-backend/docs"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	//apiv1.Use(jwt.JWT())
	{
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
	}

	return r
}
