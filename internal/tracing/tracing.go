package tracing

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gipv4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"google.golang.org/grpc"

	"github.com/houseme/yuncun-leping/utility/env"
)

const (
	tracerHostnameTagKey = "hostname"
)

// InitOtlpGrpc initializes and registers `otlpgrpc` to global TracerProvider.
//
// The output parameter `Shutdown` is used for waiting exported tracing spans to be uploaded,
// which is useful if your program is ending, and you do not want to lose recent spans.
func InitOtlpGrpc(serviceName, endpoint, traceToken, version, environment string) (*trace.TracerProvider, error) {
	// Try retrieving host ip for tracing info.
	var (
		intranetIPArray, err = gipv4.GetIntranetIpArray()
		hostIP               = "NoHostIpFound"
	)

	if err != nil {
		return nil, err
	}

	if len(intranetIPArray) == 0 {
		if intranetIPArray, err = gipv4.GetIpArray(); err != nil {
			return nil, err
		}
	}
	if len(intranetIPArray) > 0 {
		hostIP = intranetIPArray[0]
	}

	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint), // Replace the otel Agent Addr with the access point obtained in the prerequisite。
		otlptracegrpc.WithHeaders(map[string]string{"Authentication": traceToken}),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)
	ctx := context.Background()
	traceExp, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		return nil, err
	}
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			// The name of the service displayed on the traceback end。
			semconv.ServiceName(serviceName+"-"+environment),
			semconv.ServiceVersion(version),
			semconv.DeploymentEnvironment(environment),
			semconv.HostName(hostIP),
			attribute.String(tracerHostnameTagKey, hostIP),
		),
	)
	if err != nil {
		return nil, err
	}

	bsp := trace.NewBatchSpanProcessor(traceExp)
	tracerProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(res),
		trace.WithSpanProcessor(bsp),
	)

	// Set the global propagator to traceContext (not set by default).
	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tracerProvider)

	return tracerProvider, nil
}

// InitTracer initializes and registers jaeger to global TracerProvider.
func InitTracer(ctx context.Context, serviceName string) {
	if appEnv, err := env.New(ctx); err != nil {
		g.Log().Fatal(ctx, "InitTracer env new failed err:", err)
	} else {
		if _, err = InitOtlpGrpc(serviceName, appEnv.Endpoint(), appEnv.TraceToken(), appEnv.Version(), appEnv.Environment()); err != nil {
			g.Log().Fatal(ctx, "InitTracer InitOtlpGrpc failed err:", err)
		}
	}
}
