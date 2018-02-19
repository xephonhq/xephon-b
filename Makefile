PKG = github.com/at15/go.ice/example/github/pkg
VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/xab
	go install -ldflags "$(FLAGS)" ./cmd/xb

.PHONY: fmt
fmt:
	gofmt -d -l -w ./pkg ./cmd