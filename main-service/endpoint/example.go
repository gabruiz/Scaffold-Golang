package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitExampleRestService(router *gin.Engine) {
	router.GET("/", homePage)
	router.GET("/ping", pong)
	router.POST("/echo/:echo", echo) // accept a pathparam named "echo"
}

func homePage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome to homepage")
}

func pong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Pong")
}

func echo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, c.Param("echo"))
}
