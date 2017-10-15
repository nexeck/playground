#!/usr/bin/env sh

set -euo pipefail

task_init() {
    go get -u github.com/google/pprof
    go get -u github.com/ahmetb/govvv
    #go get -u github.com/smartystreets/goconvey
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
}

task_build() {
    go build -ldflags="$(govvv -flags -pkg github.com/nexeck/playground/pkg/version)" main.go
}

task_run() {
    go run -ldflags="$(govvv -flags -pkg github.com/nexeck/playground/pkg/version)" main.go $@
}

cmd="$1"
shift
case "$cmd" in
    init)   task_init ;;
    build)	task_build ;;
    run)	task_run $@ ;;
esac
