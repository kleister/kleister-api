DIST := dist
BIN := bin
EXECUTABLE := solder
VERSION := $(shell cat VERSION)

LDFLAGS += -X "main.version=$(VERSION)"

RELEASES ?= $(DIST)/$(EXECUTABLE)-linux-amd64 \
	$(DIST)/$(EXECUTABLE)-linux-386 \
	$(DIST)/$(EXECUTABLE)-linux-arm \
	$(DIST)/$(EXECUTABLE)-darwin-amd64 \
	$(DIST)/$(EXECUTABLE)-darwin-386 \
	$(DIST)/$(EXECUTABLE)-windows-amd64 \
	$(DIST)/$(EXECUTABLE)-windows-386

PACKAGES ?= $(shell go list ./... | grep -v /vendor/)

all: clean deps build test

clean:
	go clean -i ./...
	rm -rf $(BIN) $(DIST)

deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/govend/govend
	govend -v

vendor:
	govend -vtlu

generate:
	go generate $(PACKAGES)

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

build: $(BIN)/$(EXECUTABLE)

test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

release: $(RELEASES)

install: $(BIN)/$(EXECUTABLE)
	cp $< $(GOPATH)/bin/

$(BIN)/$(EXECUTABLE): $(wildcard *.go)
	CGO_ENABLED=0 go build -ldflags '-s -w $(LDFLAGS)' -o $@

$(BIN)/%/$(EXECUTABLE): GOOS=$(firstword $(subst -, ,$*))
$(BIN)/%/$(EXECUTABLE): GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
$(BIN)/%/$(EXECUTABLE):
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags '-s -w $(LDFLAGS)' -o $@

$(DIST)/$(EXECUTABLE)-%: GOOS=$(firstword $(subst -, ,$*))
$(DIST)/$(EXECUTABLE)-%: GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
$(DIST)/$(EXECUTABLE)-%: $(BIN)/%/$(EXECUTABLE)
	mkdir -p $(DIST)
	cp $(BIN)/$*/$(EXECUTABLE) $(DIST)/$(EXECUTABLE)-$(VERSION)-$(GOOS)-$(GOARCH)

.PHONY: all clean deps vendor generate fmt vet build test
.PRECIOUS: $(BIN)/%/$(EXECUTABLE)
