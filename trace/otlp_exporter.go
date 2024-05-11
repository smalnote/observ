package trace

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var _ sdktrace.SpanExporter = (*otlptrace.Exporter)(nil)

// NewHTTPExporter returns a otel exporter of OTLP HTTP.
func NewHTTPExporter(ctx context.Context, url string) (*otlptrace.Exporter, error) {
	return otlptracehttp.New(ctx, otlptracehttp.WithEndpointURL(url))
}

// NewHTTPExporter returns a otel exporter of OTLP HTTP.
func NewGPRCExporter(ctx context.Context, endpoint string) (*otlptrace.Exporter, error) {
	return otlptracegrpc.New(ctx, otlptracegrpc.WithEndpointURL(endpoint))
}
