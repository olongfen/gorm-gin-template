package router

import (
	"fmt"
	"net/http"
	"project/src/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"project/src/pkg/log_project"
	"project/src/setting"
)

// 初始化路由
var Engine = gin.Default()

// init 初始化路由模块
func init() {
	if !setting.DevEnv {
		gin.SetMode(gin.ReleaseMode)
		Engine.Use(gin.Logger())
	}

	// 添加中间件
	Engine.Use(middleware.CORS())
	Engine.Use(middleware.GinLogFormatter())
	Engine.Use(gin.Recovery())
	// 没有路由请求
	Engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": fmt.Sprintf("%v ", http.StatusNotFound) + http.StatusText(http.StatusNotFound),
		})
	})
	// TODO 路由
	{
		api := Engine.Group("api/v1")
		api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.Use(middleware.Common())

		// 测试连接
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ping": "pong >>>>>>> update"})
		})

	}
	log_project.Infoln("router init success !")
}