package metrics

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func InitMetrics(ctx context.Context) {
	exporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint("localhost:4317"),
	)
	if err != nil {
		log.Panic("Error on create metric exporter: ", err)
	}

	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("classic_cc_problens"),
	)
	metricController := controller.New(
		processor.NewFactory(simple.NewWithInexpensiveDistribution(), exporter),
		controller.WithResource(resource),
		controller.WithExporter(exporter),
		controller.WithCollectPeriod(2*time.Second),
	)

	err = metricController.Start(ctx)
	if err != nil {
		log.Panic("Error on start metric controller", err)
	}

	global.SetMeterProvider(metricController)

}

func getMetric() metric.Meter {
	return global.MeterProvider().Meter("classicccproblems")
}

func AddCounter(ctx context.Context, name string, description string) {
	counter, _ := getMetric().SyncInt64().Counter(
		name,
		instrument.WithUnit("1"),
		instrument.WithDescription(description),
	)

	counter.Add(ctx, 1, attribute.String("fibonacci", "success"))
}
