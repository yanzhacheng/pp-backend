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
	apiv1 := r.Group("")
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
		apiv1.GET("/article", v1.GetArticle)
		//新建文章
		apiv1.POST("/article", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/article", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/article", v1.DeleteArticle)
		//搜索文章
		apiv1.GET("/articles/search", v1.SearchArticles)

		//获取教程列表
		//apiv1.GET("/tutorials", v1.GetTutorials)
		//获取指定教程
		apiv1.GET("/tutorial", v1.GetTutorial)
		//获取所有教程
		apiv1.GET("/tutorials", v1.GetTutorials)
		//新建教程
		apiv1.POST("/tutorial", v1.AddTutorial)
		//更新指定教程
		apiv1.PUT("/tutorial", v1.EditTutorial)
		//删除指定教程
		apiv1.DELETE("/tutorial", v1.DeleteTutorial)
		//搜索教程
		apiv1.GET("/tutorials/search", v1.SearchTutorials)

		// 获取用户信息
		apiv1.GET("/user", v1.GetUser)
		// 注册
		apiv1.POST("/user", v1.AddUser)
		// 更新用户信息
		apiv1.PUT("/user", v1.EditUser)
		// 登录
		apiv1.POST("/user/login", v1.Login)
	}

	return r
}
