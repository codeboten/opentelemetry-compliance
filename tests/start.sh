#!/bin/bash

ROOT=`git rev-parse --show-toplevel`
OTELCOL_BUILDER_DIR=${ROOT}/bin
OTELCOL_BUILDER=${OTELCOL_BUILDER_DIR}/ocb
COLLECTOR=${ROOT}/bin/otel-validator

function install_ocb {
    if [ ! -f  ${OTELCOL_BUILDER} ]; then
        make ocb
    fi
}

function build_collector {
    if [ ! -f  ${COLLECTOR} ]; then
        echo "Building validator"
        make otel-validator
    fi
}

install_ocb
build_collector

echo "Run validator"
# ${COLLECTOR} --config ${ROOT}/validator/config.yaml &
# VALIDATOR_PID=$!

function onshutdown {
    kill ${VALIDATOR_PID}
}
# trap onshutdown EXIT

echo "Emitting telemetry"
python3 -m venv .venv
source .venv/bin/activate
# pip install -r ../examples/python/requirements.txt
opentelemetry-instrument ../examples/python/test.py
deactivate
