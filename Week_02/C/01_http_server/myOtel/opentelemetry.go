package myOtel

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

/*
	A simple opentelemetry library function

	In this example we intentionally wrap the otel libraries to minimise later code changes throughout out code base

	Usage:

		At application startup

			var shutdownTracer observability.ShutdownTracerFn
			// currently does stdout and file tracing
			ctx, shutdownTracer, err = observability.ConfigureFileTracing(ctx)
			defer shutdownTracer()

		At the start of a process

			var span observability.Span
			ctx, span = observability.TraceStart(ctx)
			defer observability.TraceEnd(span)

		Optionally at the end of the process

			if err != nil {
				observability.TraceError(span, err)
			}
*/

// wrap otel objects so clients avoid additional imports
type Span trace.Span

// wrap otel objects so clients avoid additional imports
type Exporter *stdouttrace.Exporter

// wrap otel objects so clients avoid additional imports
type Tracer trace.Tracer

// #######################################
// HTTP Handler (add to the handler chain)
// #######################################

func NewTracedHttpHandler(ctx context.Context, appname string, handler http.Handler) (wrappedHandler http.Handler) {
	wrappedHandler = otelhttp.NewHandler(handler, appname)
	return wrappedHandler
}

// #######################################
// Configure tracing destination
// #######################################
type ShutdownTracerFn func()

func ConfigureStdoutTracing(ctx context.Context, appname string) (newCtx context.Context, shutdownFn ShutdownTracerFn, err error) {

	var exp sdktrace.SpanExporter

	// get a new Exporter
	exp, err = newExporterIoWriter(os.Stdout)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	// create a new traceprovider
	tp := newTraceProvider(ctx, appname, exp)

	// register as the global trace provider
	otel.SetTracerProvider(tp)

	return ctx, func() { _ = tp.Shutdown(ctx) }, err
}

// #######################################
// Span functions (start, end, error)
// #######################################

func TraceStart(ctx context.Context, appname, spanname string) (ctx2 context.Context, span Span) {
	ctx2, span = otel.Tracer(appname).Start(ctx, spanname)
	return ctx2, span
}

func TraceEnd(span Span) {
	if span != nil {
		span.End()
	}
}

func TraceError(span Span, err error) {
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
}
