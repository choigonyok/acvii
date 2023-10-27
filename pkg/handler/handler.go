package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var handlers = gin.HandlersChain{}

func Get(router *gin.Engine) http.Handler {
	handlers = append(handlers, testHandlerFunc1)
	handlers = append(handlers, testHandlerFunc2)
	router.Any("/1", testHandlerFunc1)
	router.Any("/2", testHandlerFunc2)
	router.Any("/3", handlers...)
	return router.Handler()
}

func testHandlerFunc1(c *gin.Context) {
	fmt.Println("HandlerFunc1")
}

func testHandlerFunc2(c *gin.Context) {
	fmt.Println("HandlerFunc2")
}
