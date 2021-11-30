#!/bin/sh
set -euo pipefail

echo "--- :golang: build"

make clean && make list && make functions/all
