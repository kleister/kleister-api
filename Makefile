DIST := dist
BIN := bin
EXECUTABLE := kleister-api
SHA := $(shell git rev-parse --short HEAD)

LDFLAGS += -extldflags "-static" -X "github.com/kleister/kleister-api/config.VersionDev=$(SHA)"

RELEASES ?= windows/386 windows/amd64 darwin/386 darwin/amd64 linux/386 linux/amd64 linux/arm
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)

ifneq ($(DRONE_TAG),)
	VERSION ?= $(DRONE_TAG)
else
	ifneq ($(DRONE_BRANCH),)
		VERSION ?= $(DRONE_BRANCH)
	else
		VERSION ?= master
	endif
endif

all: clean deps vet lint test build

clean:
	go clean -i ./...
	rm -rf $(BIN) $(DIST)

deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/mitchellh/gox
	go get -u github.com/sanbornm/go-selfupdate

vendor:
	go get -u github.com/govend/govend
	govend -v

update:
	govend -vtlu

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

lint:
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

install: $(BIN)/$(EXECUTABLE)
	cp $< $(GOPATH)/bin/

build: $(BIN)/$(EXECUTABLE)

$(BIN)/$(EXECUTABLE): $(wildcard *.go)
	CGO_ENABLED=0 go build -ldflags '-s -w $(LDFLAGS)' -o $@

release: release-build release-copy release-check

release-build:
	gox -osarch='$(RELEASES)' -ldflags='-s -w $(LDFLAGS)' -output='$(BIN)/$(EXECUTABLE)-{{.OS}}-{{.Arch}}'

release-copy:
	mkdir -p $(DIST)/release
	$(foreach file,$(wildcard $(BIN)/$(EXECUTABLE)-*),cp $(file) $(DIST)/release/$(EXECUTABLE)-$(VERSION)-$(word 3,$(subst -, ,$(notdir $(file))))-$(subst .exe,,$(word 4,$(subst -, ,$(notdir $(file)))));)

release-check:
	cd $(DIST)/release; $(foreach file,$(wildcard $(DIST)/release/*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

latest: release-build latest-copy latest-check

latest-copy:
	mkdir -p $(DIST)/latest
	$(foreach file,$(wildcard $(BIN)/$(EXECUTABLE)-*),cp $(file) $(DIST)/latest/$(EXECUTABLE)-latest-$(word 3,$(subst -, ,$(notdir $(file))))-$(subst .exe,,$(word 4,$(subst -, ,$(notdir $(file)))));)

latest-check:
	cd $(DIST)/latest; $(foreach file,$(wildcard $(DIST)/latest/*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

updater: release-build updater-copy updater-push

updater-copy:
	mkdir -p $(DIST)/updater
	$(foreach file,$(wildcard $(BIN)/$(EXECUTABLE)-*),cp $(file) $(DIST)/updater/$(word 3,$(subst -, ,$(notdir $(file))))-$(subst .exe,,$(word 4,$(subst -, ,$(notdir $(file)))));)

updater-push:
	go-selfupdate -o $(DIST)/publish $(DIST)/updater $(VERSION)

publish: release latest updater

.PHONY: all clean deps vendor update fmt vet lint test build release latest updater publish
