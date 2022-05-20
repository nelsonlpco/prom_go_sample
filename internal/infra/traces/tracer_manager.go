package traces

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

const instrumentationName = "otelgrpc"

type TracerManager struct {
	tracerProvider *sdktrace.TracerProvider
	resource       *resource.Resource
	exporter       *otlptrace.Exporter
	tracer         trace.Tracer
}

func (t *TracerManager) WithTraceProvider() *TracerManager {
	t.tracerProvider = sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(t.exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(t.resource),
	)

	return t
}

func (t *TracerManager) WithResource() *TracerManager {
	t.resource = resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("classic_cc_problens"),
	)

	return t
}

func (t *TracerManager) withExporter(ctx context.Context) *TracerManager {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("localhost:4317"),
		otlptracegrpc.WithTimeout(5 * time.Second),
	}

	client := otlptracegrpc.NewClient(opts...)
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Panic(fmt.Sprintf("error on create otlp exporter : %v", err))
	}

	t.exporter = exporter

	return t
}

func (t *TracerManager) withTracer(name string) *TracerManager {
	t.tracer = t.tracerProvider.Tracer(name)

	return t
}

func (t *TracerManager) Start(ctx context.Context, description string) *Span {
	ctx, internal := t.tracer.Start(ctx, description)
	span := NewSpan(ctx, internal, description)

	return span
}

func New(ctx context.Context) *TracerManager {
	return new(TracerManager).
		WithResource().
		withExporter(ctx).
		WithTraceProvider().
		withTracer(instrumentationName)
}
