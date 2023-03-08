# OpenTelemetry Compliance

This repository includes code to test compliance of telemetry producers against the OpenTelemetry specification. The validator is implemented using the OpenTelemetry Collector with an OTLP receiver, and a validator exporter. This exporter validates all the data against semantic conventions and emits validation results in various formats.

## Getting started

### Requirements (TBD)

The initial requirements include Python to run the example:

* Python 3.7+

### Setup the validator

```bash
make ocb
make otel-validator
./bin/otel-validator --config ./validator/config.yaml
```

### Produce telemetry

Run code that produces telemetry and emits it using OTLP. The current example uses Python:

```bash
./start.sh
```

### Run with dlv

Run with `dlv` to debug the collector execution

```bash
dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient --log exec bin/otel-validator -- --config validator/config.yaml
```
