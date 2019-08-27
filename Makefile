ifndef VERBOSE
.SILENT:
endif

override WORK_DIR ?= $(shell pwd)
override GO_PKG ?= github.com/Nerufa/go-blueprint
override GO_PATH ?= $(shell pwd)/../../../..

GOOS ?= "linux"
GOARCH ?= "amd64"

CGO_ENABLED ?= 0
GO_IMAGE ?= nerufa/docker-go
GO_IMAGE_TAG ?= 1.12

DOCKER_IMAGE = nerufa/go-blueprint
TAG ?= unknown
CACHE_TAG ?= unknown_cache

ifneq (, $(shell which go))
 GO_PKG = $(shell go list -e -f "{{ .ImportPath }}")
 GOOS = $(shell go env GOOS)
 GOARCH = $(shell go env GOARCH)
endif

define build_resources
 	find "$(WORK_DIR)/resources" -maxdepth 1 -mindepth 1 -exec cp -R -f {} $(WORK_DIR)/artifacts/${1} \;
endef

define go_docker
	docker run --rm \
		-v /${WORK_DIR}:/${WORK_DIR} \
		-v /${GO_PATH}/pkg/mod:/${GO_PATH}/pkg/mod \
		-w /${WORK_DIR} \
		-e GOPATH=/${GO_PATH} \
		${GO_IMAGE}:${GO_IMAGE_TAG} \
		sh -c 'GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} TAG=${TAG} $(subst ",,${1})'
endef

all: vendor generate build ## install cli tools, update vendor, generate code & build application
.PHONY: all

build: init ## build application
	if [ "${DIND}" = "1" ]; then \
		$(call go_docker,"make build") ;\
    else \
		$(call build_resources) ;\
        GO111MODULE=on GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} \
        go build -mod vendor -ldflags "-X $(GO_PKG)/cmd/version.appVersion=$(TAG)-$$(date -u +%Y%m%d%H%M)" -o "$(WORK_DIR)/artifacts/bin" main.go ;\
    fi;
.PHONY: build

clean: ## remove generated files, tidy vendor dependicies
	if [ "${DIND}" = "1" ]; then \
		$(call go_docker,"make clean") ;\
    else \
        export GO111MODULE=on ;\
    	rm -rf profile.out artifacts/* generated/* vendor ;\
    	go mod tidy ;\
    fi;
.PHONY: clean

dev-build-up: build docker-image-cache dev-docker-compose-up ## create new build and recreate containers in docker-compose
.PHONY: dev-build-up

dev-docker-compose-down: ## down dev env
	TAG=${TAG} docker-compose -p blueprint -f docker/docker-compose.yml -f docker/docker-compose-local.yml down
.PHONY: dev-docker-compose-down

dev-docker-compose-up: ## up dev env
	TAG=${TAG} docker-compose -p blueprint -f docker/docker-compose.yml -f docker/docker-compose-local.yml up -d
.PHONY: dev-docker-compose-up

dev-test: test lint ## test application in dev env with race and lint
.PHONY: dev-test

dind: ## useful for windows
	if [ "${DIND}" = "1" ]; then \
		echo "Already in DIND" ;\
    else \
    	if [ -z "${GOPATH}" ]; then \
    		echo "GOPATH should be present" && exit 1 ;\
        fi ;\
	    docker run -it --rm --name dind --privileged \
            -v //var/run/docker.sock://var/run/docker.sock \
            -v /${WORK_DIR}:/${WORK_DIR} \
            -v /${GO_PATH}/pkg/mod:/${GO_PATH}/pkg/mod \
            -w /${WORK_DIR} \
            nerufa/docker-dind:19 ;\
    fi;
.PHONY: dind

docker-clean: ## delete previous docker image build
	docker rmi ${DOCKER_IMAGE}:${TAG} || true ;\
.PHONY: docker-clean

docker-image-cache: ## build docker image and tagged as cache
	docker build --cache-from ${DOCKER_IMAGE}:${CACHE_TAG} -f "${WORK_DIR}/docker/app/Dockerfile" -t ${DOCKER_IMAGE}:${TAG} -t ${DOCKER_IMAGE}:${CACHE_TAG} ${WORK_DIR}
.PHONY: docker-image-cache

docker-image: ## build docker image
	docker build --cache-from ${DOCKER_IMAGE}:${CACHE_TAG} -f "${WORK_DIR}/docker/app/Dockerfile" -t ${DOCKER_IMAGE}:${TAG} ${WORK_DIR}
.PHONY: docker-image

docker-protoc-generate: ## generate proto, grpc client & server
	docker run --rm \
	 	-v /${WORK_DIR}/resources:/${WORK_DIR}/resources \
	 	-v /${WORK_DIR}/generated:/${WORK_DIR}/generated \
	 	-v /${WORK_DIR}/prototool.yaml:/${WORK_DIR}/prototool.yaml \
	 	-w /${WORK_DIR} \
	 	uber/prototool \
	 	prototool generate resources/proto
.PHONY: docker-protoc-generate

docker-push: ## push docker image to registry
	docker push ${DOCKER_IMAGE}:${TAG}
.PHONY: docker-push

generate: gqlgen-generate docker-protoc-generate go-generate ## execute all generators
.PHONY: generate

github-build: docker-image docker-push docker-clean ## build application in GitLab CI
.PHONY: github-build

github-test: vendor test-with-coverage ## test application in GitLab CI
.PHONY: github-test

go-depends: ## view final versions that will be used in a build for all direct and indirect dependencies
	if [ "${DIND}" = "1" ]; then \
		$(call go_docker,"make go-depends") ;\
    else \
        cd $(WORK_DIR) ;\
        GO111MODULE=on go list -m all ;\
    fi;
.PHONY: go-depends

go-generate: ## go generate
	if [ "${DIND}" = "1" ]; then \
		$(call go_docker,"make go-generate") ;\
    else \
        cd $(WORK_DIR) ;\
        go generate $$(go list ./app/...) || exit 1 ;\
        $(MAKE) vendor  ;\
    fi;
.PHONY: go-generate

go-update-all: ## view available minor and patch upgrades for all direct and indirect
	if [ "${DIND}" = "1" ]; then \
		$(call go_docker,"make go-update-all") ;\
    else \
        cd $(WORK_DIR) ;\
    	GO111MODULE=on go list -u -m all ;\
    fi;
.PHONY: go-update-all

gqlgen-generate: ## generate graphql server
	if [ "${DIND}" = "1" ]; then \
		$(call go_docker,"make gqlgen-generate") ;\
    else \
        gqlgen -v ;\
    fi;
.PHONY: gqlgen-generate

lint: ## execute linter
	if [ "${DIND}" = "1" ]; then \
		$(call go_docker,"make lint") ;\
    else \
        golangci-lint run --no-config --issues-exit-code=0 --deadline=30m \
          --disable-all --enable=deadcode  --enable=gocyclo --enable=golint --enable=varcheck \
          --enable=structcheck --enable=maligned --enable=errcheck --enable=dupl --enable=ineffassign \
          --enable=interfacer --enable=unconvert --enable=goconst --enable=gosec --enable=megacheck ./... ;\
    fi;
.PHONY: lint

test-with-coverage: ## test application with race and total coverage
	if [ "${DIND}" = "1" ]; then \
		$(call go_docker,"make test-with-coverage") ;\
    else \
        GO111MODULE=on CGO_ENABLED=1 \
        go test -mod vendor -v -race -covermode atomic -coverprofile profile.out ./app/... || exit 1 ;\
        go tool cover -func=profile.out && rm -rf profile.out ;\
    fi;
.PHONY: test-with-coverage

test: ## test application with race
	if [ "${DIND}" = "1" ]; then \
		$(call go_docker,"make test") ;\
    else \
        GO111MODULE=on CGO_ENABLED=1 \
        go test -mod vendor  -race -v ./app/... ;\
    fi;
.PHONY: test

vendor: ## update vendor dependencies
	if [ "${DIND}" = "1" ]; then \
		$(call go_docker,"make vendor") ;\
    else \
        rm -rf $(WORK_DIR)/vendor ;\
    	GO111MODULE=on \
    	go mod vendor ;\
    fi;
.PHONY: vendor

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

init:
	mkdir -p generated artifacts ;\
    rm -rf artifacts/*
.PHONY: init

.DEFAULT_GOAL := help