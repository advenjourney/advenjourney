SHELL := bash
NAME := api
IMPORT := github.com/advenjourney/advenjourney/$(NAME)
BIN := bin
DIST := dist
CONTAINER_PREFIX := advenjourney

ifeq ($(OS), Windows_NT)
	EXECUTABLE := $(NAME).exe
else
	EXECUTABLE := $(NAME)
endif

PACKAGES ?= $(shell go list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f)
GENERATE ?= $(PACKAGES)

ifndef DATE
	DATE := $(shell date -u '+%Y%m%d')
endif

ifndef VERSION
	VERSION ?= $(shell git rev-parse --short HEAD)
endif

ifndef REVISION
	REVISION ?= $(shell git rev-parse --short HEAD)
endif

.PHONY: all
all: assets build

.PHONY: web
web: 
	(cd web; yarn dev; cd ..)

.PHONY: clean
clean:
	rm -rf ./web/docs/.vuepress/dist
	rm -rf ./build

.PHONY: assets
assets:
	(rm -rf ./build/web; mkdir -p build/web)
	(cd web; yarn exec vuepress build docs; mv ./docs/.vuepress/dist ../build/web; cd ..)
	(cd api; make clean; cp -R ../build/web ./assets; make generate; cd ..)

.PHONY: build
build:
	(cd api; make build; mv bin/api ../build; cd ..)

.PHONY: test
test:
	echo "Acceptance tests ..."

.PHONY: release
release: release-dirs release-build release-checksums

.PHONY: release-dirs
release-dirs:
	mkdir -p $(DIST)

.PHONY: release-build
release-build:
	@which gox > /dev/null; if [ $$? -ne 0 ]; then \
		GO111MODULE=off  $(GO) get -u github.com/mitchellh/gox; \
	fi
	gox -arch="386 amd64 arm" -osarch '!darwin/386' -verbose -ldflags '-w $(LDFLAGS)' -output="$(DIST)/$(EXECUTABLE)-{{.OS}}-{{.Arch}}" ./cmd/$(NAME)

.PHONY: release-checksums
release-checksums:
	cd $(DIST); $(foreach file, $(wildcard $(DIST)/$(EXECUTABLE)-*), sha256sum $(notdir $(file)) > $(notdir $(file)).sha256;)

.PHONY: container
container:
	docker build -t $(CONTAINER_PREFIX)/$(NAME):$(VERSION) -f ./Dockerfile .

.PHONE: container-push
container-push: container
	docker push $(CONTAINER_PREFIX)/$(NAME):$(VERSION)
