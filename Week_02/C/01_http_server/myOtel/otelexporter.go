package myOtel

import (
	"io"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// #######################################
// Exporter
//    func NewExporter(ctx context.Context) (someExporter.Exporter, error) {
//    	Your preferred exporter: console, jaeger, zipkin, OTLP, etc.
//    }
// #######################################

// Returns an exporter based on io.Writer, e.g. file os.stdout, os.file
func newExporterIoWriter(w io.Writer) (sdktrace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
	)
}
