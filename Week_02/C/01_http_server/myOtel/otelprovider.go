package myOtel

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// #######################################
// Provider
// #######################################
func newTraceProvider(ctx context.Context, appname string, exp sdktrace.SpanExporter) *sdktrace.TracerProvider {

	// The service.name attribute is required.
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(appname),
	)

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	newGlobalTextMapPropagator()

	return provider
}
