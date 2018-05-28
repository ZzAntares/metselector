#!/usr/bin/env sh
command -v markys >/dev/null 2>&1 || {
    echo >&2 "I require golint but it's not installed.  Aborting.";
    echo >&2 "  See: https://github.com/golang/lint";
    exit 1;
}

golint $(go list ./... | grep -v /vendor/) &&
    go test $(go list ./... | grep -v /vendor/)
