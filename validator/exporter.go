package validator

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type validatorExporter struct {
}

// validateTraces validates each span against its semantic convention
func (ve *validatorExporter) validateTraces(ctx context.Context, td ptrace.Traces) error {
	resourceSpans := td.ResourceSpans()
	if resourceSpans.Len() == 0 {
		return nil
	}
	for i := 0; i < resourceSpans.Len(); i++ {
		fmt.Println(resourceSpans.At(i))
	}
	return nil
}

// validateMetrics validates each metric against its semantic convention
func (ve *validatorExporter) validateMetrics(ctx context.Context, md pmetric.Metrics) error {
	return nil
}

// validateLogs validates each log against its semantic convention.
func (ve *validatorExporter) validateLogs(ctx context.Context, ld plog.Logs) error {
	return nil
}

func createTracesExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Traces, error) {
	conf := cfg.(*Config)

	ve, err := createValidatorExporter()
	if err != nil {
		return nil, err
	}
	return exporterhelper.NewTracesExporter(
		ctx,
		set,
		conf,
		ve.validateTraces,
	)
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Metrics, error) {
	conf := cfg.(*Config)

	ve, err := createValidatorExporter()
	if err != nil {
		return nil, err
	}
	return exporterhelper.NewMetricsExporter(
		ctx,
		set,
		conf,
		ve.validateMetrics,
	)
}

func createLogsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Logs, error) {
	conf := cfg.(*Config)

	ve, err := createValidatorExporter()
	if err != nil {
		return nil, err
	}
	return exporterhelper.NewLogsExporter(
		ctx,
		set,
		conf,
		ve.validateLogs,
	)
}

func createValidatorExporter() (validatorExporter, error) {
	return validatorExporter{}, nil
}
