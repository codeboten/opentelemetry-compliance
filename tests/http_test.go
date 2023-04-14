package main

// "go.opentelemetry.io/collector/component"

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"

	v1 "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	trace "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/grpc"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

var resourceSpans = []*trace.ResourceSpans{}

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func thatLowCardinalityRouteAvailableIs(arg1 string) error {
	return nil
}

func theInstrumentationCreatesAServerSpanForAGETOperation() error {
	return nil
}

func theServerSpanNameSHOULDBe(expectedServerSpanName string) error {
	span := resourceSpans[0].ScopeSpans[0].Spans[0]

	err := assertExpectedActual(assert.Equal, span.Name, expectedServerSpanName, "%s should be equal to %s", span.Name, expectedServerSpanName)

	if err != nil {
		return err
	}

	return nil
}

func thereIsAnHTTPServerWithALowCardinalityRouteAvailable() error {
	return nil
}

func thereIsAnHTTPServerWithoutALowCardinalityRouteAvailable() error {
	return nil
}

func thereIsAnOpenTelemetryHTTPInstrumentationForThatServer() error {
	return nil
}

type server struct {
	v1.UnimplementedTraceServiceServer
}

func (s server) Export(ctx context.Context, req *v1.ExportTraceServiceRequest) (*v1.ExportTraceServiceResponse, error) {
	for _, entry := range req.ResourceSpans {
		resourceSpans = append(resourceSpans, entry)
	}
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

	go grpc_server.Serve(listen)
	command := exec.Command("./start.sh")
	err = command.Run()

	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(3000 * time.Millisecond)

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

type asserter struct {
	err error
}

type expectActualAssertion func(t assert.TestingT, expected, actual interface{}, messageArgs ...interface{}) bool

func assertExpectedActual(assertion expectActualAssertion, expected, actual interface{}, messageAndArgs ...interface{}) error {
	var t asserter
	assertion(&t, expected, actual, messageAndArgs...)
	return t.err
}

func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}
