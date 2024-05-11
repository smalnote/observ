package debug

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

// RegisterGin the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.IRouter. prefixOptions is a optional. If not prefixOptions,
// the default path prefix is used, otherwise first prefixOptions will be path prefix.
func RegisterGin(router gin.IRouter) {
	router.GET("/pprof", gin.WrapF(pprof.Index))
	router.GET("/cmdline", gin.WrapF(pprof.Cmdline))
	router.GET("/profile", gin.WrapF(pprof.Profile))
	router.POST("/symbol", gin.WrapF(pprof.Symbol))
	router.GET("/symbol", gin.WrapF(pprof.Symbol))
	router.GET("/trace", gin.WrapF(pprof.Trace))
	router.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
	router.GET("/block", gin.WrapH(pprof.Handler("block")))
	router.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
	router.GET("/heap", gin.WrapH(pprof.Handler("heap")))
	router.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
	router.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
}
