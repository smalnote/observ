package main

import (
	"log/slog"
	"os"

	"github.com/smalnote/observ/log/otelslog"
)

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	otelHandler := otelslog.New(jsonHandler)
	slog.SetDefault(slog.New(otelHandler))
}
