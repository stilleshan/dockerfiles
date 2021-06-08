package server

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/giuem/ga-proxy/ga"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func getUUID(c *gin.Context) string {
	uid, err := c.Cookie("uuid")
	if err != nil { // cookie no found
		uid = generateUUID(c.Request.UserAgent())
		c.SetCookie("uuid", uid, 2147483647, "/", "", false, false)
	}

	return uid
}

func generateUUID(name string) string {
	ns, err := uuid.NewV4()
	if err != nil {
		// error fallback
		unix32bits := uint32(time.Now().UTC().Unix())
		nameBytes := md5.Sum([]byte(name))

		return fmt.Sprintf("%x-%x-%x-%x-%x\n", unix32bits, nameBytes[0:2], nameBytes[2:4], nameBytes[4:6], nameBytes[6:12])
	}

	return uuid.NewV5(ns, name).String()
}

func getCommonData(c *gin.Context) ga.CommonData {
	return ga.CommonData{
		Version:          1,
		TrackingID:       c.Query("ga"),
		ClientID:         getUUID(c),
		UserIP:           c.ClientIP(),
		UserAgent:        c.Request.UserAgent(),
		DocumentReferer:  c.Query("dr"),
		ScreenResolution: c.Query("sr"),
		ViewportSize:     c.Query("vp"),
		DocumentEncoding: c.Query("de"),
		ScreenColors:     c.Query("sd"),
		UserLanguage:     c.Query("ul"),
		DocumentLink:     c.Request.Referer(),
		DocumentTitle:    c.Query("dt"),
	}
}

func getTimingData(c *gin.Context) ga.TimingData {
	return ga.TimingData{
		PageLoadedTime:     c.Query("plt"),
		DNSTime:            c.Query("dns"),
		PageDownloadedTime: c.Query("pdt"),
		RedirectTime:       c.Query("rrt"),
		TCPTime:            c.Query("tcp"),
		ServerResponseTime: c.Query("srt"),
		DomInteractiveTime: c.Query("dit"),
		ContentLoadedTime:  c.Query("clt"),
	}
}
