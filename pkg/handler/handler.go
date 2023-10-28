package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var handlers = gin.HandlersChain{}

// Get assembles and returns handlers
func Get(router *gin.Engine) http.Handler {
	handlers = append(handlers, testHandlerFunc1)
	handlers = append(handlers, testHandlerFunc2)
	router.Any("/1", testHandlerFunc1)
	router.Any("/2", testHandlerFunc2)
	router.Any("/3", handlers...)
	return router.Handler()
}

// TEST
func testHandlerFunc1(c *gin.Context) {
	fmt.Println("HandlerFunc1")
}

// TEST
func testHandlerFunc2(c *gin.Context) {
	fmt.Println("HandlerFunc2")
}
