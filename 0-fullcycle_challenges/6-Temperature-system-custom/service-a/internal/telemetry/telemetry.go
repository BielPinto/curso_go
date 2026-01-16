package telemetry

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

// InitTracerProvider initializes the OTEL tracer provider
func InitTracerProvider(ctx context.Context, serviceName string) (*sdktrace.TracerProvider, error) {
	// Get OTEL endpoint from environment
	otlpEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if otlpEndpoint == "" {
		otlpEndpoint = "localhost:4318"
	}
	// Remove http:// prefix if present (otlptracehttp adds it automatically)
	if strings.HasPrefix(otlpEndpoint, "http://") {
		otlpEndpoint = strings.TrimPrefix(otlpEndpoint, "http://")
	}
	if strings.HasPrefix(otlpEndpoint, "https://") {
		otlpEndpoint = strings.TrimPrefix(otlpEndpoint, "https://")
	}

	// Create HTTP exporter
	exporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(otlpEndpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create exporter: %w", err)
	}

	// Create resource
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// Create tracer provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// Set global tracer provider
	otel.SetTracerProvider(tp)

	fmt.Printf("OTEL tracer initialized for %s with endpoint %s\n", serviceName, otlpEndpoint)

	return tp, nil
}

// Shutdown gracefully shuts down the tracer provider
func Shutdown(ctx context.Context, tp *sdktrace.TracerProvider) error {
	return tp.Shutdown(ctx)
}

// GetTracer returns a tracer instance
func GetTracer(name string) trace.Tracer {
	return otel.Tracer(name)
}
