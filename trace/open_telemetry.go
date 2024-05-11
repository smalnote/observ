// Package trace handles distributed system tracing stuffs.
package trace

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

// Config OpenTelemetry resource config.
type Config struct {
	ServiceName    string
	ServiceVersion string
	Environment    string
}

// SetProvider set global OpenTelemetry trace provider with exporter.
func SetProvider(c Config, exporter sdktrace.SpanExporter) (func(context.Context) error, error) {
	batcher := sdktrace.NewBatchSpanProcessor(exporter)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithResource(c.resource()),
		sdktrace.WithSpanProcessor(batcher),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tp.Shutdown, nil
}

func (c *Config) resource() *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(c.ServiceName),
		semconv.ServiceVersion(c.ServiceVersion),
		semconv.DeploymentEnvironment(c.Environment),
	)
}

const name = "app"

type (
	Tracer = trace.Tracer
	Span   = trace.Span

	TracerOption = trace.TracerOption

	SpanStartOption = trace.SpanStartOption
	SpanEndOption   = trace.SpanEndOption

	// SpanOption are options that can be used at both the beginning and end of a span.
	SpanOption = trace.SpanOption

	// EventOption applies span event options to an EventConfig.
	EventOption = trace.EventOption

	// SpanEventOption are options that can be used with an event or a span.
	SpanEventOption = trace.SpanEventOption

	// SpanStartEventOption are options that can be used at the start of a span, or with an event.
	SpanStartEventOption = trace.SpanStartEventOption

	// SpanEndEventOption are options that can be used at the end of a span, or with an event.
	SpanEndEventOption = trace.SpanEndEventOption
)

// export alias
var (
	ContextPropagator propagation.TextMapPropagator = propagation.TraceContext{}

	WithTimestamp  func(t time.Time) SpanEventOption = trace.WithTimestamp
	WithStackTrace func(b bool) SpanEndEventOption   = trace.WithStackTrace
)

// DefaultTracer returns the default app tracer.
func DefaultTracer() trace.Tracer {
	return otel.Tracer(name)
}

// Start starts a new span with the default app tracer.
func Start(ctx context.Context, spanName string, opts ...SpanStartOption) (context.Context, Span) {
	return DefaultTracer().Start(ctx, spanName, opts...)
}

func End(span Span, options ...SpanEndOption) {
	span.End(options...)
}

// HTTPHeaderExtract extract trace context from HTTP header to ctx.
func HTTPHeaderExtract(ctx context.Context, header http.Header) context.Context {
	carrier := propagation.HeaderCarrier(header)
	return ContextPropagator.Extract(ctx, carrier)
}

// HTTPHeaderInject inject trace context from context to http header.
func HTTPHeaderInject(ctx context.Context, header http.Header) {
	carrier := propagation.HeaderCarrier(header)
	ContextPropagator.Inject(ctx, carrier)
}

// ContextTrace returns trace span in context.
func ContextTrace(ctx context.Context) propagation.MapCarrier {
	carrier := propagation.MapCarrier{}
	ContextPropagator.Inject(ctx, carrier)
	return carrier
}
