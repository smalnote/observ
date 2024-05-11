package trace

import (
	"go.opentelemetry.io/otel/exporters/zipkin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var _ sdktrace.SpanExporter = (*zipkin.Exporter)(nil)

// NewZipkinExporter returns a otel exporter of Zipkin.
func NewZipkinExporter(url string) (*zipkin.Exporter, error) {
	return zipkin.New(url)
}
