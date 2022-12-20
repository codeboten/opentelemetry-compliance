#!/bin/bash

OTELCOL_BUILDER_DIR=`git rev-parse --show-toplevel`/bin
OTELCOL_BUILDER=${OTELCOL_BUILDER_DIR}/ocb
COLLECTOR=`git rev-parse --show-toplevel`/bin/otel-validator

function install_ocb {
    if [ ! -f  $OTELCOL_BUILDER ]; then
        make ocb
    fi
}

function build_collector {
    if [ ! -f  $COLLECTOR ]; then
        echo "Building validator"
        make ocb
    fi
}

install_ocb
build_collector

# echo "Run validator"

echo "Emitting telemetry"
python3 -m venv .venv
source .venv/bin/activate
pip install -r ./examples/python/requirements.txt
opentelemetry-instrument ./examples/python/test.py
deactivate

# echo "Shutting down validator"