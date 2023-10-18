package main

import (
	"net/http"

	"github.com/choigonyok/goopt/internal/test"
	"github.com/choigonyok/goopt/pkg"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./assets", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	println(test.Test())

	pkg.InClusterConfig()

	router.Run(":8080")
}

func GoOptHandler(c *gin.Context) {
	println("NOTHING")
}
