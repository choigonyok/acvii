package middleware

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

var config = cors.DefaultConfig()

type Middleware struct {
	middlewarePool []gin.HandlerFunc
}

// New creates new Middleware
func New() *Middleware {
	return &Middleware{}
}

// Get returns current middleware pool
func (m *Middleware) Get() []gin.HandlerFunc {
	return m.middlewarePool
}

// AllowOrigin adds cors origin configuration middleware in pool
func (m *Middleware) AllowOrigin(s ...string) {
	m.middlewarePool = append(m.middlewarePool,
		func() gin.HandlerFunc {
			if s[0] == "*" {
				config.AllowAllOrigins = true
			} else {
				config.AllowedOrigins = s
			}
			return cors.New(config)
		}())
}

// AllowOrigin adds cors method configuration middleware in pool
func (m *Middleware) AllowMethod(s ...string) {
	m.middlewarePool = append(m.middlewarePool,
		func() gin.HandlerFunc {
			config.AllowedMethods = s
			return cors.New(config)
		}())
}

// AllowOrigin adds cors header configuration middleware in pool
func (m *Middleware) AllowHeader(s ...string) {
	m.middlewarePool = append(m.middlewarePool,
		func() gin.HandlerFunc {
			config.AllowedHeaders = s
			return cors.New(config)
		}())
}

// AllowOrigin adds cors credential configuration middleware in pool
func (m *Middleware) AllowCredential() {
	m.middlewarePool = append(m.middlewarePool,
		func() gin.HandlerFunc {
			config.AllowCredentials = true
			return cors.New(config)
		}())
}
