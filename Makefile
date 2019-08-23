GO_DIR ?= $(shell pwd)
GO_PKG ?= $(shell go list -e -f "{{ .ImportPath }}")

GOOS?=$(shell go env GOOS || echo linux)
GOARCH?=$(shell go env GOARCH || echo amd64)
CGO_ENABLED?=0

GO_IMAGE?=golang
GO_IMAGE_TAG?=1.12

DOCKER_IMAGE=unknown
TAG?=unknown
CACHE_TAG?=unknown_cache

define build_resources
 	find "$(GO_DIR)/resources" -maxdepth 1 -mindepth 1 -exec cp -R -f {} $(GO_DIR)/artifacts/${1} \;
endef

all: install vendor generate build ## install cli tools, update vendor, generate code & build application

build: init ## build application
	$(call build_resources) ;\
	GO111MODULE=on GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} \
	go build -mod vendor -ldflags "-X $(GO_PKG)/cmd/version.appVersion=$(TAG)-$$(date -u +%Y%m%d%H%M)" -o "$(GO_DIR)/artifacts/bin" main.go

clean: ## remove generated files, tidy vendor dependicies
	export GO111MODULE=on ;\
	rm -rf profile.out artifacts/* generated/* vendor ;\
	go mod tidy

dev-docker-compose-down: ## down dev env
	docker-compose -p blueprint -f docker/docker-compose.yml -f docker/docker-compose-local.yml down

dev-docker-compose-up: ## up dev env
	docker-compose -p blueprint -f docker/docker-compose.yml -f docker/docker-compose-local.yml up -d

dev-test: test lint ## test application in dev env with race and lint

docker-image: ## build docker image
	docker rmi ${DOCKER_IMAGE}:${TAG} || true ;\
	docker build --cache-from ${DOCKER_IMAGE}:${CACHE_TAG} -f "${GO_DIR}/docker/app/Dockerfile" -t ${DOCKER_IMAGE}:${TAG} ${GO_DIR}

docker-protoc-generate: ## generate proto, grpc client & server
	docker run --rm \
	 	-v ${GO_DIR}/resources:${GO_DIR}/resources \
	 	-v ${GO_DIR}/generated:${GO_DIR}/generated \
	 	-v ${GO_DIR}/prototool.yaml:${GO_DIR}/prototool.yaml \
	 	-w ${GO_DIR} \
	 	uber/prototool \
	 	prototool generate resources/proto

docker-push: ## push docker image to registry
	docker push ${DOCKER_IMAGE}:${TAG}

generate: gqlgen-generate docker-protoc-generate go-generate ## execute all generators

github-build: docker-image docker-push ## build application in GitLab CI

github-test: vendor test-with-coverage ## test application in GitLab CI

go-depends: ## view final versions that will be used in a build for all direct and indirect dependencies
	cd $(GO_DIR) ;\
	GO111MODULE=on \
	go list -m all

go-generate: ## go generate
	cd $(GO_DIR) ;\
	go generate $$(go list ./app/...) || exit 1 ;\
	$(MAKE) vendor

go-update-all: ## view available minor and patch upgrades for all direct and indirect
	cd $(GO_DIR) ;\
	GO111MODULE=on \
	go list -u -m all

gqlgen-generate: ## generate graphql server
	gqlgen -v

install: init ## install cli tools
	export GO111MODULE=off ;\
    go get -u github.com/smartystreets/goconvey ;\
    go get -u github.com/google/wire/cmd/wire ;\
    go get -u github.com/golangci/golangci-lint/cmd/golangci-lint ;\
    go get -u github.com/99designs/gqlgen ;\
    go get -u github.com/vektah/dataloaden ;\
	go get -u github.com/rubenv/sql-migrate/sql-migrate

lint: ## execute linter
	golangci-lint run --no-config --issues-exit-code=0 --deadline=30m \
      --disable-all --enable=deadcode  --enable=gocyclo --enable=golint --enable=varcheck \
      --enable=structcheck --enable=maligned --enable=errcheck --enable=dupl --enable=ineffassign \
      --enable=interfacer --enable=unconvert --enable=goconst --enable=gosec --enable=megacheck ./...

test-with-coverage: ## test application with race and total coverage
	GO111MODULE=on CGO_ENABLED=1 \
	go test -mod vendor -v -race -covermode atomic -coverprofile profile.out ./app/... || exit 1 ;\
    go tool cover -func=profile.out && rm -rf profile.out

test: ## test application with race
	GO111MODULE=on CGO_ENABLED=1 \
	go test -mod vendor  -race -v ./app/...

vendor: ## update vendor dependencies
	rm -rf $(GO_DIR)/vendor ;\
	GO111MODULE=on \
	go mod vendor

.PHONY: all build clean \
        dev-docker-compose-down dev-docker-compose-up dev-test \
        docker-image docker-protoc-generate docker-push \
        generate \
        github-build github-test \
        go-depends go-generate go-update-all gqlgen-generate install lint \
        test-with-coverage test vendor help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init:
	mkdir -p generated artifacts ;\
    rm -rf artifacts/*

.DEFAULT_GOAL := help