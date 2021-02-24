app_name   = golang-gin-template

BuildStamp = $(shell date '+%Y%m%d%H%M%S')
GitHash    = $(shell git rev-parse HEAD)
Version    = $(shell git describe --abbrev=0 --tags --always)
Target     = ${app_name}

.PHONY: build
build: generate
	go build -v -o bin/${Target} -tags=jsoniter -ldflags \
	"-X main.BuildStamp=${BuildStamp} -X main.GitHash=${GitHash} -X main.Version=${Version}"

.PHONY: release
release: generate
	go build -v -o bin/${Target} -tags=jsoniter -ldflags \
	"-s -w -X main.BuildStamp=${BuildStamp} -X main.GitHash=${GitHash} -X main.Version=${Version}"

.PHONY: generate
generate:
	go generate

.PHONY: lint
lint:
	golangci-lint run -v --timeout=5m

.PHONY: image
image:
	docker build -t ${app_name}:$(Version) .

.PHONY: clean
clean:
	@$(RM) -r *.db *.db-journal
