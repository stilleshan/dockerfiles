package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/giuem/ga-proxy/ga"
)

func handlePageView(c *gin.Context) {
	if len(c.Request.Referer()) == 0 || len(c.Query("ga")) == 0 {
		handleRedirect(c)
		return
	}
	c.Status(http.StatusOK)
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	go ga.PageView(getCommonData(c))
}

func handleTiming(c *gin.Context) {
	if len(c.Request.Referer()) == 0 || len(c.Query("ga")) == 0 {
		handleRedirect(c)
		return
	}
	c.Status(http.StatusOK)
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")

	go ga.Timing(getCommonData(c), getTimingData(c))
}

func handlePing(c *gin.Context) {
	err := ga.Detect()
	if err != nil {
		if c.Request.Method == http.MethodHead {
			c.Status(http.StatusBadGateway)
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"msg": err.Error()})
		}
		return
	}

	if c.Request.Method == http.MethodHead {
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	}
}

func handleRedirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://github.com/giuem/ga-proxy")
}
