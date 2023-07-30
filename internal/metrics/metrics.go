package metrics

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var environment = "local"
var service = "go-server-app"
var meter = otel.Meter(service)

var newResource = func() *resource.Resource {
	resource :=
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			semconv.DeploymentEnvironmentKey.String(environment),
		)
	return resource
}

func GetMeter() (*sdkmetric.MeterProvider, error) {
	ctx := context.Background()
	exporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint("collector:5555"),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	// labels/tags/resources that are common to all metrics.
	resourceVars := newResource()

	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resourceVars),
		sdkmetric.WithReader(
			// collects and exports metric data every 30 seconds.
			sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(30*time.Second)),
		),
	)

	otel.SetMeterProvider(mp)

	return mp, nil
}

func IncreaseEmployee(id int) {
	counter, _ := meter.
		Int64Counter(
			"employee_create",
			metric.WithDescription("how many ime create employee is called"),
		)
	counter.Add(
		context.Background(),
		1,
		// labels/tags
		metric.WithAttributes(attribute.Int("add employee with id", id)))
}

func GetEmployee(size int) {
	counter, _ := meter.
		Int64Counter(
			"employee_get",
			metric.WithDescription("how many times get employee is called"),
		)
	counter.Add(
		context.Background(),
		1,
		// labels/tags
		metric.WithAttributes(attribute.Int("get employee size", size)))
}
