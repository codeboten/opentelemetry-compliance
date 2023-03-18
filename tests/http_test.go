package main

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
)

var opts = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	os.Exit(status)
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

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^That low cardinality route available is "([^"]*)"\.$`, thatLowCardinalityRouteAvailableIs)
	ctx.Step(`^The instrumentation creates a Server Span for a GET operation\.$`, theInstrumentationCreatesAServerSpanForAGETOperation)
	ctx.Step(`^The Server Span Name SHOULD be "([^"]*)"\.$`, theServerSpanNameSHOULDBe)
	ctx.Step(`^There is an HTTP server with a low cardinality route available\.$`, thereIsAnHTTPServerWithALowCardinalityRouteAvailable)
	ctx.Step(`^There is an HTTP server without a low cardinality route available\.$`, thereIsAnHTTPServerWithoutALowCardinalityRouteAvailable)
	ctx.Step(`^There is an OpenTelemetry HTTP instrumentation for that server\.$`, thereIsAnOpenTelemetryHTTPInstrumentationForThatServer)
}
