package main

// "go.opentelemetry.io/collector/component"

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	v1 "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	trace "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/grpc"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func thatLowCardinalityRouteAvailableIs(arg1 string) error {
	return godog.ErrPending
}

func theInstrumentationCreatesAServerSpanForAGETOperation() error {
	return godog.ErrPending
}

func theServerSpanNameSHOULDBe(arg1 string) error {
	return godog.ErrPending
}

func thereIsAnHTTPServerWithALowCardinalityRouteAvailable() error {
	return godog.ErrPending
}

func thereIsAnHTTPServerWithoutALowCardinalityRouteAvailable() error {
	return godog.ErrPending
}

func thereIsAnOpenTelemetryHTTPInstrumentationForThatServer() error {
	return godog.ErrPending
}

type server struct {
	v1.UnimplementedTraceServiceServer
}

func some_function() error {
	fmt.Println("some_function")
	return nil
}

var resource_spans = []*trace.ResourceSpans{}

func (s server) Export(ctx context.Context, req *v1.ExportTraceServiceRequest) (*v1.ExportTraceServiceResponse, error) {
	for _, entry := range req.ResourceSpans {
		resource_spans = append(resource_spans, entry)
	}
	fmt.Println(resource_spans)
	return &v1.ExportTraceServiceResponse{}, nil
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	listen, err := net.Listen("tcp", ":4317")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpc_server := grpc.NewServer()

	v1.RegisterTraceServiceServer(grpc_server, server{})

	log.Printf("server listening at %v", listen.Addr())

	go grpc_server.Serve(listen)
	time.Sleep(10000 * time.Millisecond)

	fmt.Println(resource_spans)

	//if err := grpc_server.Serve(listen); err != nil {
	//	log.Fatalf("Failed to serve: %v", err)
	//}

	status := godog.TestSuite{
		Name:                "HTTP Instrumentation",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	os.Exit(status)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^That low cardinality route available is "([^"]*)"\.$`, thatLowCardinalityRouteAvailableIs)
	ctx.Step(`^The instrumentation creates a Server Span for a GET operation\.$`, theInstrumentationCreatesAServerSpanForAGETOperation)
	ctx.Step(`^The Server Span Name SHOULD be "([^"]*)"\.$`, theServerSpanNameSHOULDBe)
	ctx.Step(`^There is an HTTP server with a low cardinality route available\.$`, thereIsAnHTTPServerWithALowCardinalityRouteAvailable)
	ctx.Step(`^There is an HTTP server without a low cardinality route available\.$`, thereIsAnHTTPServerWithoutALowCardinalityRouteAvailable)
	ctx.Step(`^There is an OpenTelemetry HTTP instrumentation for that server\.$`, thereIsAnOpenTelemetryHTTPInstrumentationForThatServer)
}
