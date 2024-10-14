include .bingo/Variables.mk

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
golangci: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run ./...

.PHONY: staticcheck
staticcheck: $(STATICCHECK)
	$(STATICCHECK) -tags '$(TAGS)' $(PACKAGES)

.PHONY: lint
lint: $(REVIVE)
	for PKG in $(PACKAGES); do $(REVIVE) -config revive.toml -set_exit_status $$PKG || exit 1; done;

.PHONY: generate
generate: openapi mocks
	go generate $(PACKAGES)

.PHONY: openapi
openapi: $(OAPI_CODEGEN)
	$(OAPI_CODEGEN) --config=pkg/api/v1/config.yml openapi/v1.yml

.PHONY: mocks
mocks: \
	pkg/upload/mock.go pkg/store/mock.go \
	pkg/service/build_versions/mock.go \
	pkg/service/builds/mock.go \
	pkg/service/mods/mock.go \
	pkg/service/packs/mock.go \
	pkg/service/team_mods/mock.go \
	pkg/service/team_packs/mock.go \
	pkg/service/teams/mock.go \
	pkg/service/user_mods/mock.go \
	pkg/service/user_packs/mock.go \
	pkg/service/user_teams/mock.go \
	pkg/service/users/mock.go \
	pkg/service/versions/mock.go \
	pkg/service/minecraft/mock.go \
	pkg/service/forge/mock.go \
	pkg/service/neoforge/mock.go \
	pkg/service/quilt/mock.go \
	pkg/service/fabric/mock.go

pkg/upload/mock.go: pkg/upload/upload.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package upload

pkg/store/mock.go: pkg/store/store.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package store

pkg/service/build_versions/mock.go: pkg/service/build_versions/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package buildversions

pkg/service/builds/mock.go: pkg/service/builds/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package builds

pkg/service/mods/mock.go: pkg/service/mods/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package mods

pkg/service/packs/mock.go: pkg/service/packs/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package packs

pkg/service/team_mods/mock.go: pkg/service/team_mods/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package teammods

pkg/service/team_packs/mock.go: pkg/service/team_packs/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package teampacks

pkg/service/teams/mock.go: pkg/service/teams/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package teams

pkg/service/user_mods/mock.go: pkg/service/user_mods/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package usermods

pkg/service/user_packs/mock.go: pkg/service/user_packs/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package userpacks

pkg/service/user_teams/mock.go: pkg/service/user_teams/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package userteams

pkg/service/users/mock.go: pkg/service/users/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package users

pkg/service/versions/mock.go: pkg/service/versions/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package versions

pkg/service/minecraft/mock.go: pkg/service/minecraft/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package minecraft

pkg/service/forge/mock.go: pkg/service/forge/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package forge

pkg/service/neoforge/mock.go: pkg/service/neoforge/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package neoforge

pkg/service/quilt/mock.go: pkg/service/quilt/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package quilt

pkg/service/fabric/mock.go: pkg/service/fabric/service.go $(MOCKGEN)
	$(MOCKGEN) -source $< -destination $@ -package fabric

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
