SHELL := bash
NAME := kleister-api
IMPORT := github.com/kleister/$(NAME)
BIN := bin
DIST := dist

ifeq ($(OS), Windows_NT)
	EXECUTABLE := $(NAME).exe
	UNAME := Windows
else
	EXECUTABLE := $(NAME)
	UNAME := $(shell uname -s)
endif

GOBUILD ?= CGO_ENABLED=0 go build
PACKAGES ?= $(shell go list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f -not -iname mock.go -not -path ./.devenv/\* -not -path ./.direnv/\*)

TAGS ?= netgo

ifndef OUTPUT
	ifeq ($(GITHUB_REF_TYPE), tag)
		OUTPUT ?= $(subst v,,$(GITHUB_REF_NAME))
	else
		OUTPUT ?= testing
	endif
endif

ifndef VERSION
	ifeq ($(GITHUB_REF_TYPE), tag)
		VERSION ?= $(subst v,,$(GITHUB_REF_NAME))
	else
		VERSION ?= $(shell git rev-parse --short HEAD)
	endif
endif

ifndef DATE
	DATE := $(shell date -u '+%Y%m%d')
endif

ifndef SHA
	SHA := $(shell git rev-parse --short HEAD)
endif

LDFLAGS += -s -w -extldflags "-static" -X "$(IMPORT)/pkg/version.String=$(VERSION)" -X "$(IMPORT)/pkg/version.Revision=$(SHA)" -X "$(IMPORT)/pkg/version.Date=$(DATE)"
GCFLAGS += all=-N -l

.PHONY: all
all: build

.PHONY: sync
sync:
	go mod download

.PHONY: clean
clean:
	go clean -i ./...
	rm -rf $(BIN) $(DIST)

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: golangci
golangci:
	go tool github.com/golangci/golangci-lint/cmd/golangci-lint run ./...

.PHONY: lint
lint:
	for PKG in $(PACKAGES); do go tool github.com/mgechev/revive -config revive.toml -set_exit_status $$PKG || exit 1; done;

.PHONY: generate
generate:
	go generate $(PACKAGES)

.PHONY: test
test: test
	go test -coverprofile coverage.out $(PACKAGES)

.PHONY: install
install: $(SOURCES)
	go install -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' ./cmd/$(NAME)

.PHONY: build
build: $(BIN)/$(EXECUTABLE)

$(BIN)/$(EXECUTABLE): $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./cmd/$(NAME)

$(BIN)/$(EXECUTABLE)-debug: $(SOURCES)
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -gcflags '$(GCFLAGS)' -o $@ ./cmd/$(NAME)

.PHONY: release
release: $(DIST) release-linux release-darwin release-windows release-checksum

$(DIST):
	mkdir -p $(DIST)

.PHONY: release-linux
release-linux: $(DIST) \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-386 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-amd64 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-5 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-6 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-7 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm64

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-5:
	GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./cmd/$(NAME)

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-6:
	GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./cmd/$(NAME)

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-arm-7:
	GOOS=linux GOARCH=arm GOARM=7 $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./cmd/$(NAME)

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-linux-%:
	GOOS=linux GOARCH=$* $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./cmd/$(NAME)

.PHONY: release-darwin
release-darwin: $(DIST) \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-darwin-amd64 \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-darwin-arm64

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-darwin-%:
	GOOS=darwin GOARCH=$* $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./cmd/$(NAME)

.PHONY: release-windows
release-windows: $(DIST) \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-windows-4.0-amd64.exe \
	$(DIST)/$(EXECUTABLE)-$(OUTPUT)-windows-4.0-arm64.exe \

$(DIST)/$(EXECUTABLE)-$(OUTPUT)-windows-4.0-%.exe:
	GOOS=windows GOARCH=$* $(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $@ ./cmd/$(NAME)

.PHONY: release-reduce
release-reduce:
	cd $(DIST); $(foreach file,$(wildcard $(DIST)/$(EXECUTABLE)-*),upx $(notdir $(file));)

.PHONY: release-checksum
release-checksum:
	cd $(DIST); $(foreach file,$(wildcard $(DIST)/$(EXECUTABLE)-*),sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

.PHONY: release-finish
release-finish: release-reduce release-checksum

.PHONY: watch
watch:
	go tool github.com/air-verse/air -c .air.toml

.PHONY: posting
posting:
	posting default --collection posting
