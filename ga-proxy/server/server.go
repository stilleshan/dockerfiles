package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Run starts a HTTP server
func Run(ip, port string) {
	addr := fmt.Sprintf("%v:%v", ip, port)

	r := gin.New()
	logger := gin.Logger()
	r.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/ping" {
			return
		}
		logger(c)
	})

	r.NoRoute(handleRedirect)
	// version < 1
	r.GET("/", handlePageView)
	// version >= 1
	r.GET("/p", handlePageView)
	r.GET("/t", handleTiming)

	r.GET("/ping", handlePing)
	r.HEAD("/ping", handlePing)

	r.Run(addr)
}
