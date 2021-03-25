TARGETS     ?= golang-gin-template
PKG          = github.com/tosone/golang-gin-template
CMD_DIR      = ./cmd
OUTPUT_DIR   = ./bin
BUILD_DIR    = ./build
BUILDSTAMP   = $(shell date '+%Y%m%d%H%M%S')
GITHASH      = $(shell git rev-parse HEAD)
VERSION      = $(shell git describe --abbrev=0 --tags --always)

.PHONY: build
build:
	@for target in $(TARGETS); do                                     \
	  go generate $(CMD_DIR)/$${target};                              \
	  go build -v -o $(OUTPUT_DIR)/$${target}                         \
	    -ldflags "-s -w -X main.Version=$(VERSION)                    \
	    -X main.BuildStamp=$(BUILDSTAMP) -X main.GitHash=$(GITHASH)"  \
	    $(CMD_DIR)/$${target};                                        \
	done

.PHONY: image
image:
	@for target in $(TARGETS); do                                     \
	  image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                 \
	  echo Building $${image}:$(VERSION);                             \
	  docker build -t $${image}:$(VERSION)                            \
	    --build-arg PKG=${PKG}                                        \
	    --build-arg GITCOMMIT=${COMMIT}                               \
	    --build-arg BUILDVERSION=${COMMIT}                            \
	    --build-arg BUILDDATE=${COMMIT}                               \
	    --build-arg TARGET=$${target}                                 \
	    -f $(BUILD_DIR)/$${target}/Dockerfile .;                      \
	done

.PHONY: lint
lint:
	golangci-lint run -v --timeout=5m

.PHONY: clean
clean:
	@$(RM) -r *.db *.db-journal
