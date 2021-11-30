#!/bin/sh
set -euo pipefail

if [[ "$BUILDKITE_BRANCH" != "main" ]]; then
  echo "Cannot release latest from branch $BUILDKITE_BRANCH. Must be 'main'."
  exit 1
fi

if [[ "$BUILDKITE_TAG" != "" ]]; then
  echo "Cannot release latest from tag $BUILDKITE_TAG. Must be empty."
  exit 1
fi

buildkite-agent artifact upload "zip/**" "s3://$BUILDKITE_ARTIFACT_S3_BUCKET/$BUILDKITE_PIPELINE_SLUG/latest"
