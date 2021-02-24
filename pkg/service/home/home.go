package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Home Get API
// @Description description for home
// @Success 200 {string} string	"ok"
// @Router / [get]
func Get(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Home"})
}
