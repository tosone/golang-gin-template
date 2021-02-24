package server

import (
	"github.com/gin-gonic/gin"

	serviceHello "github.com/tosone/golang-gin-template/pkg/service/hello"
	serviceHome "github.com/tosone/golang-gin-template/pkg/service/home"
)

func routers(app *gin.Engine) {
	app.GET("/", serviceHome.Get)

	var hello = app.Group("hello")
	{
		hello.GET("/", serviceHello.Get)
	}
}
