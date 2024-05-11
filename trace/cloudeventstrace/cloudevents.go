package cloudeventstrace

import (
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// HTTPTraceOption cloudevents HTTP client with OTel tracing.
var HTTPTraceOption = cloudevents.WithRoundTripper(
	otelhttp.NewTransport(http.DefaultTransport),
)
