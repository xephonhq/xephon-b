PKG = github.com/at15/go.ice/example/github/pkg
VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/xb

.PHONY: test
test:
	go test -v -cover ./pkg/...

.PHONY: fmt
fmt:
	gofmt -d -l -w ./pkg ./cmd

.PHONY: generate
generate:
	gommon generate -v

.PHONY: loc
loc:
	cloc --exclude-dir=vendor,.idea,playground .

.PHONY: package
package: install
	cp $(shell which xb) .
	cp $(shell which xab) .
	zip xb-$(VERSION).zip xb
	zip xab-$(VERSION).zip xab
	rm xb
	rm xab