package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/smalnote/observ/debug"
	"github.com/smalnote/observ/log"
	"github.com/smalnote/observ/log/otelslog"
	"github.com/smalnote/observ/trace"
)

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		},
	)
	otelHandler := otelslog.New(jsonHandler)
	slog.SetDefault(slog.New(otelHandler))

	ginEngine := gin.New()

	ginEngine.Use(log.GinSlog())
	ginEngine.Use(trace.GinTrace())

	debugRouter := ginEngine.Group("/debug")
	debug.RegisterGin(debugRouter)

	ginEngine.Run(":8080")
}
