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
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		//apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)

		//获取教程列表
		//apiv1.GET("/tutorials", v1.GetTutorial)
		//获取指定教程
		apiv1.GET("/tutorials/:id", v1.GetTutorial)
		//新建教程
		apiv1.POST("/tutorials", v1.AddTutorial)
		//更新指定教程
		apiv1.PUT("/tutorials/:id", v1.EditTutorial)
		//删除指定教程
		apiv1.DELETE("/tutorials/:id", v1.DeleteTutorial)
	}

	return r
}
