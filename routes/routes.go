// 在这个文件中注册 URL handler

package routes

import (
	"github.com/axiaoxin-com/pink-lady/webserver"
	"github.com/gin-gonic/gin"
)

// Routes 注册 API URL 路由
func Routes(app *gin.Engine) {
	cloud := app.Group("/cloud")
	cloud.Use(webserver.Session("hanyun"))
	cloud.POST("/add_user", add_user)
	cloud.POST("/login", user_login)
	cloud.GET("/captcha", func(c *gin.Context) {
		Captcha(c, 4)
	})
	// TODO: 在这里注册你的 gin API，如： app.GET("/", HandlerFunc)
}
