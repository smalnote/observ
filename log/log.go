package log

import (
	"github.com/gin-gonic/gin"
)

// Logger returns gin logger middleware.
func Logger() gin.HandlerFunc {
	return GinSlog()
}
