package main

// "go.opentelemetry.io/collector/component"

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/converter/expandconverter"
	"go.opentelemetry.io/collector/confmap/provider/envprovider"
	"go.opentelemetry.io/collector/confmap/provider/fileprovider"
	"go.opentelemetry.io/collector/confmap/provider/httpprovider"
	"go.opentelemetry.io/collector/confmap/provider/yamlprovider"
	"go.opentelemetry.io/collector/otelcol"
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

func startCollector(ctx context.Context, t *testing.T, col *Collector) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		require.NoError(t, col.Run(ctx))
	}()
	return wg
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	map_provider := make(map[string]confmap.Provider, 4)

	file_provider := fileprovider.New()
	env_provider := envprovider.New()
	yaml_provider := yamlprovider.New()
	http_provider := httpprovider.New()

	map_provider[file_provider.Scheme()] = file_provider
	map_provider[env_provider.Scheme()] = env_provider
	map_provider[yaml_provider.Scheme()] = yaml_provider
	map_provider[http_provider.Scheme()] = http_provider

	config_provider_settings := otelcol.ConfigProviderSettings{
		ResolverSettings: confmap.ResolverSettings{
			URIs:       []string{filepath.Join("testdata", "otelcol-nop.yaml")},
			Providers:  map_provider,
			Converters: []confmap.Converter{expandconverter.New()},
		},
	}

	cfgProvider, err := otelcol.NewConfigProvider(config_provider_settings)

	collector_settings := otelcol.CollectorSettings{
		BuildInfo:      component.NewDefaultBuildInfo(),
		Factories:      otelcol.Factories{},
		ConfigProvider: cfgProvider,
	}

	col, err := otelcol.NewCollector(collector_settings)

	fmt.Println(col)
	fmt.Println(err)

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
