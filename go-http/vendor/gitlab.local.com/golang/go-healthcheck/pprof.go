// Package healthcheck HTTP Server health check
// Created by chenguolin 2019-02-13
package healthcheck

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

const (
	// AllocsPath allocs
	allocsPath = "/debug/pprof/allocs"
	// BlockPath  block
	blockPath = "/debug/pprof/block"
	// CmdlinePath  cmdline
	cmdlinePath = "/debug/pprof/cmdline"
	// GoroutinePath  goroutine
	goroutinePath = "/debug/pprof/goroutine"
	// HeapPath  heap
	heapPath = "/debug/pprof/heap"
	// MutexPath  mutex
	mutexPath = "/debug/pprof/mutex"
	// ProfilePath  profile
	profilePath = "/debug/pprof/profile"
	// ThreadCreatePath  threadcreate
	threadCreatePath = "/debug/pprof/threadcreate"
	// TracePath   trace
	tracePath = "/debug/pprof/trace"
)

// addPProf add health router
func addPProf(e *gin.Engine) {
	e.GET(allocsPath, gin.WrapH(pprof.Handler("allocs")))
	e.GET(blockPath, gin.WrapH(pprof.Handler("block")))
	e.GET(cmdlinePath, gin.WrapH(pprof.Handler("cmdline")))
	e.GET(goroutinePath, gin.WrapH(pprof.Handler("goroutine")))
	e.GET(heapPath, gin.WrapH(pprof.Handler("heap")))
	e.GET(mutexPath, gin.WrapH(pprof.Handler("mutex")))
	e.GET(profilePath, gin.WrapH(pprof.Handler("profile")))
	e.GET(threadCreatePath, gin.WrapH(pprof.Handler("threadcreate")))
	e.GET(tracePath, gin.WrapH(pprof.Handler("trace")))
}
