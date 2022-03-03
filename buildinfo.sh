#!/bin/sh

set -e

BUILDSTAMP=$(date --utc --iso-8601=seconds)
COMMITSTAMP=$(git show --no-patch --format=%cI HEAD)
REF=$(git rev-parse HEAD)
SHORTREF=$(git rev-parse --short HEAD)
ACTION_URL="$GITHUB_SERVER_URL/$GITHUB_REPOSITORY/actions/runs/$GITHUB_RUN_ID"
ACTION_ID="$GITHUB_RUN_ID"

cat << HEREDOC
{
  "buildstamp":  "$BUILDSTAMP",
  "commitstamp": "$COMMITSTAMP",
  "ref":         "$REF",
  "shortref":    "$SHORTREF",
  "action_id":   "$ACTION_ID",
  "action_url":  "$ACTION_URL"
}
HEREDOC
