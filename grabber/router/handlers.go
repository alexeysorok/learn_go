package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func collectHandler(c *gin.Context) {
	category := c.Param("category")

	news.ColLect(category)
	c.String(http.StatusOK, "Search is initialized")
}

func resultHandler(c *gin.Context) {
	category := c.Param("caregory")
	topics := news.Result(category)

	c.JSON(http.StatusOK, topics)
}
