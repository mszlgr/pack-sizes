# Exclude vendored and generated files from linting
APP_SRC = $(shell go list -f '{{range $$_, $$f := .GoFiles}}{{$$.Dir}}/{{$$f}}{{"\n"}}{{end}}' ./... | sed -e s,$$(pwd)/,, | grep -v "^gen/")
SRC = go.mod go.sum $(shell find . -name '*.go')

APP_NAME = pack-sizes
BUILD_CMD = go build

all: build

lint:
	@ FMT=$$(gofmt -l $(APP_SRC)); if test -n "$$FMT"; then echo "These files are not go fmt'd:"; echo "$$FMT" | sed 's/^/  - /'; exit 1; fi;

fmt:
	@ gofmt -l -w $(APP_SRC)

build: $(SRC)
	@ $(BUILD_CMD)

static: $(SRC)
	@ CGO_ENABLED=0 $(BUILD_CMD) 

_test:
	@ go test -race ./... --count 1

test: _test lint

clean:
	@ rm -f $(OUT) $(OUT_LINUX) $(OUT_MACOS) $(GEN_OUT)
	@ go clean ./...


go.sum: go.mod
	@ go mod download

docker-build:
	@ docker build . -t $(APP_NAME):$(version)

docker-build-dev:
	@ docker build . -t $(APP_NAME):dev

docker-run: docker-build-dev
	@ docker run -e APP_HOST=0.0.0.0 -p 12345:12345 $(APP_NAME):dev
