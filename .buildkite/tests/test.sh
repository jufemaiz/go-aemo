#!/bin/sh
set -euo pipefail

echo "--- :golang: tests"

make test

echo "--- :codecov: upload coverage report"

bash -c "bash <(curl -s https://codecov.io/bash)"
