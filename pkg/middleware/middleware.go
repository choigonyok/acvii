package middleware

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

var config = cors.DefaultConfig()

type Middleware struct {
	middleware []gin.HandlerFunc
}

func New() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Get() []gin.HandlerFunc {
	return m.middleware
}

func (m *Middleware) AllowOrigin(s ...string) {
	m.middleware = append(m.middleware,
		func() gin.HandlerFunc {
			if s[0] == "*" {
				config.AllowAllOrigins = true
			} else {
				config.AllowedOrigins = s
			}
			return cors.New(config)
		}())
}

func (m *Middleware) AllowMethod(s ...string) {
	m.middleware = append(m.middleware,
		func() gin.HandlerFunc {
			config.AllowedMethods = s
			return cors.New(config)
		}())
}

func (m *Middleware) AllowHeader(s ...string) {
	m.middleware = append(m.middleware,
		func() gin.HandlerFunc {
			config.AllowedHeaders = s
			return cors.New(config)
		}())
}

func (m *Middleware) AllowCredential() {
	m.middleware = append(m.middleware,
		func() gin.HandlerFunc {
			config.AllowCredentials = true
			return cors.New(config)
		}())
}
