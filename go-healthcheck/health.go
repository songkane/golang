// Package healthcheck HTTP Server health check
// Created by chenguolin 2019-02-13
package healthcheck

import "github.com/gin-gonic/gin"

const (
	// HealthPath health
	healthPath = "/debug/health"
)

// addHealth add health router
func addHealth(e *gin.Engine) {
	e.GET(healthPath, checkHealth)
}

// checkHealth check health http handler
func checkHealth(c *gin.Context) {
	code := 200
	contentType := "application/json; charset=utf-8"
	body := "ok"

	c.Data(code, contentType, []byte(body))
}
