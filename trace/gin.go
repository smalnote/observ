package trace

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/codes"
)

// GinTrace returns a gin middleware that add span for every request
func GinTrace(stackTrace bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctxWithTrace := HTTPHeaderExtract(c.Request.Context(), c.Request.Header)
		ctx, span := Start(ctxWithTrace, c.Request.URL.Path)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		if c.Writer.Status() >= http.StatusBadRequest {
			span.SetStatus(codes.Error, "http code "+strconv.Itoa(c.Writer.Status()))
		}
		span.End(WithStackTrace(c.Writer.Status() >= http.StatusBadRequest))
	}
}
