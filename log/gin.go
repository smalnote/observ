// This is a copy of github.com/gin-contrib/zap with custom behaviour:
// #1 log request body when respone status is not 200
package log

import (
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GinSlog returns a gin.HandlerFunc (middleware) that logs requests using log/slog.
func GinSlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				slog.Error(e)
			}
			return
		}
		fields := []any{
			"path", path,
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"client_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"time", end.Format(time.RFC3339),
			"proto", c.Request.Proto,
			"latency", latency,
		}
		if la, ok := c.Request.Context().Value(http.LocalAddrContextKey).(net.Addr); ok {
			fields = append(fields, "local_addr", la.String())
		}
		if c.Writer.Status() >= http.StatusBadRequest {
			slog.ErrorContext(c.Request.Context(), "gin request errored", fields...)
			return
		}
		slog.DebugContext(c.Request.Context(), "gin request completed", fields...)
	}

}
