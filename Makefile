DIST := dist
BIN := bin
EXECUTABLE := kleister-api
IMPORT := github.com/kleister/kleister-api
SHA := $(shell git rev-parse --short HEAD)

LDFLAGS += -X "github.com/kleister/kleister-api/config.VersionDev=$(SHA)"

TARGETS ?= linux/*,darwin/*,windows/*
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)

TAGS ?=

ifneq ($(DRONE_TAG),)
	VERSION ?= $(DRONE_TAG)
else
	ifneq ($(DRONE_BRANCH),)
		VERSION ?= $(DRONE_BRANCH)
	else
		VERSION ?= master
	endif
endif

all: clean test build

clean:
	go clean -i ./...
	rm -rf $(BIN) $(DIST)

generate:
	go generate $(PACKAGES)

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

lint:
	@which golint > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/golang/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

install: $(BIN)/$(EXECUTABLE)
	cp $< $(GOPATH)/bin/

build: $(BIN)/$(EXECUTABLE)

$(BIN)/$(EXECUTABLE): $(wildcard *.go)
	go build -tags '$(TAGS)' -ldflags '-s -w $(LDFLAGS)' -o $@

release: release-build release-copy release-check

release-build:
	@which xgo > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/karalabe/xgo; \
	fi
	xgo -dest $(BIN) -tags '$(TAGS)' -ldflags '-s -w $(LDFLAGS)' -targets '$(TARGETS)' -out $(EXECUTABLE)-$(VERSION) $(IMPORT)

release-copy:
	mkdir -p $(DIST)/release
	$(foreach file,$(wildcard $(BIN)/$(EXECUTABLE)-*),cp $(file) $(DIST)/release/$(notdir $(file));)

release-check:
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/release/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

latest: release-build latest-copy latest-check

latest-copy:
	mkdir -p $(DIST)/latest
	$(foreach file,$(wildcard $(BIN)/$(EXECUTABLE)-*),cp $(file) $(DIST)/latest/$(subst $(EXECUTABLE)-$(VERSION),$(EXECUTABLE)-latest,$(notdir $(file)));)

latest-check:
	cd $(DIST)/latest; $(foreach file,$(wildcard $(DIST)/latest/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

publish: release latest

.PHONY: all clean generate fmt vet lint test build release latest publish
