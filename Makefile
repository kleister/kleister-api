DIST := dist
BIN := bin
EXECUTABLE := solder
VERSION := $(shell cat VERSION)

LDFLAGS += -X "main.version=$(VERSION)"

RELEASES ?= $(DIST)/$(EXECUTABLE)-linux-amd64 \
	$(DIST)/$(EXECUTABLE)-linux-386 \
	$(DIST)/$(EXECUTABLE)-linux-arm \
	$(DIST)/$(EXECUTABLE)-darwin-amd64

PACKAGES ?= $(shell GO15VENDOREXPERIMENT=1 go list ./... | grep -v /vendor/)

all: clean deps build test

clean:
	go clean -i ./...
	rm -rf $(BIN) $(DIST)

deps:
	GO15VENDOREXPERIMENT=1 go get -u github.com/jteeuwen/go-bindata/...
	GO15VENDOREXPERIMENT=1 go get -u github.com/govend/govend
	GO15VENDOREXPERIMENT=1 govend -v

vendor:
	GO15VENDOREXPERIMENT=1 govend -vtlu

generate:
	go generate $(PACKAGES)

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

build: $(BIN)/$(EXECUTABLE)

test:
	@for PKG in $(PACKAGES); do GO15VENDOREXPERIMENT=1 go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

release: $(RELEASES)

install: $(BIN)/$(EXECUTABLE)
	cp $< $(GOPATH)/bin/

$(BIN)/$(EXECUTABLE): $(wildcard *.go)
	GO15VENDOREXPERIMENT=1 CGO_ENABLED=1 go build -ldflags '-s -w $(LDFLAGS)' -o $@

$(BIN)/%/$(EXECUTABLE): GOOS=$(firstword $(subst -, ,$*))
$(BIN)/%/$(EXECUTABLE): GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
$(BIN)/%/$(EXECUTABLE):
	GO15VENDOREXPERIMENT=1 CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags '-s -w $(LDFLAGS)' -o $@

$(DIST)/$(EXECUTABLE)-%: GOOS=$(firstword $(subst -, ,$*))
$(DIST)/$(EXECUTABLE)-%: GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
$(DIST)/$(EXECUTABLE)-%: $(BIN)/%/$(EXECUTABLE)
	mkdir -p $(DIST)
	cp $(BIN)/$*/$(EXECUTABLE) $(DIST)/$(EXECUTABLE)-$(VERSION)-$(GOOS)-$(GOARCH)

.PHONY: all clean deps vendor generate fmt vet build test
.PRECIOUS: $(BIN)/%/$(EXECUTABLE)
