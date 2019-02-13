// Package healthcheck gin register healthcheck router
// Created by chenguolin 2019-02-13
package healthcheck

import "github.com/gin-gonic/gin"

// RegisterHealthCheck gin register health check
func RegisterHealthCheck(e *gin.Engine) {
	// add health check
	addHealth(e)
	// add pprof
	addPProf(e)
}
