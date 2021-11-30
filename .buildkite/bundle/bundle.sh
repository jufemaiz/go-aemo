#!/bin/sh
set -euo pipefail

echo "--- :golang: bundle = build + zip"

make clean && make list && make zip/all
