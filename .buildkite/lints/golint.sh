#!/bin/sh
set -euo pipefail

echo "--- :golang::lint-roller: golinting"

make lint
