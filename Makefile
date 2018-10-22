# setting some defaults if those variables are empty
GO_PIPELINE_NAME?=$(shell basename ${PWD})
CLUSTER?=services
OWNER=vevo
APP_NAME=$(GO_PIPELINE_NAME)
IMAGE_NAME=$(OWNER)/$(APP_NAME)
AWS_REGION?=us-east-1
AWS_PROFILE?=default
AWS_ACCOUNT?=$(AWS_PROFILE)
AWSCONFIG_VERSION?=1.0.2
STRINGER_VERSION?=latest
GO_REVISION?=$(shell git rev-parse HEAD)
GO_TO_REVISION?=$(GO_REVISION)
GO_FROM_REVISION?=$(shell git rev-parse refs/remotes/origin/master)
GIT_TAG=$(IMAGE_NAME):$(GO_REVISION)
BUILD_VERSION?=1.0.$(GO_PIPELINE_COUNTER)
BUILD_TAG=$(IMAGE_NAME):$(BUILD_VERSION)
LATEST_TAG=$(IMAGE_NAME):latest
# DOCKERFILE is a path related to resources/docker-compose.yaml
DOCKERFILE?=../Dockerfile

DC=GO_REVISION=$(GO_REVISION) GO_TO_REVISION=$(GO_TO_REVISION) GO_FROM_REVISION=$(GO_FROM_REVISION) DOCKERFILE=$(DOCKERFILE) BUILD_VERSION=$(BUILD_VERSION) CLUSTER=$(CLUSTER) AWS_REGION=$(AWS_REGION) AWS_PROFILE=$(AWS_PROFILE) AWS_ACCOUNT=$(AWS_ACCOUNT) GO_PIPELINE_NAME=$(GO_PIPELINE_NAME) AWSCONFIG_VERSION=$(AWSCONFIG_VERSION) STRINGER_VERSION=$(STRINGER_VERSION) docker-compose -p $(GO_PIPELINE_NAME) -f resources/docker-compose.yaml

docker-lint:
	$(DC) run --rm dockerlint

docker-login:
	@docker login -u "$(DOCKER_USER)" -p "$(DOCKER_PASS)"

docker-build:
	docker build --build-arg BUILD_VERSION=$(BUILD_VERSION) -t $(GIT_TAG) .

docker-tag:
	docker tag $(GIT_TAG) $(BUILD_TAG)
	docker tag $(GIT_TAG) $(LATEST_TAG)

docker-push: docker-login
	docker push $(GIT_TAG)
	docker push $(BUILD_TAG)
	docker push $(LATEST_TAG)

# Cleans previous docker-compose resources
dc-clean:
	$(DC) down --rmi local --remove-orphans -v
	$(DC) rm -f -v

# Sets up aws credentials inside docker-compose for later use in stringer
dc-config: dc-clean
	$(DC) run --rm config

go-dep:
	@if [ -f "glide.yaml" ] ; then \
		go get github.com/Masterminds/glide \
		&& go install github.com/Masterminds/glide \
		&& glide install --strip-vendor; \
	elif [ -f "Godeps/Godeps.json" ] ; then \
		go get github.com/tools/godep \
		&& godep restore; \
	else \
		go get -d -t -v ./...; \
	fi

GOFILES_NOVENDOR=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

go-fmt:
	@[ $$(gofmt -l $(GOFILES_NOVENDOR) | wc -l) -gt 0 ] && echo "Code differs from gofmt's style" && exit 1 || true

go-lint: go-fmt
	@go get -u golang.org/x/lint/golint; \
	if [ -f "glide.yaml" ] ; then \
		golint -set_exit_status $$(glide novendor); \
		go vet -v $$(glide novendor); \
	else \
		golint -set_exit_status ./...; \
		go vet -v ./...; \
	fi

go-test:
	@if [ -f "glide.yaml" ] ; then \
		go test $$(glide novendor); \
	else \
		go test -v ./...; \
	fi

go-build: go-dep go-lint go-test
	# github.com/VEVO/nfs-go/logger.version is only necessary when using nfs-go
	@go build -v -a -ldflags "-X main.version=$(BUILD_VERSION) -X github.com/VEVO/nfs-go/logger.version=$(BUILD_VERSION)"

go-compile:
	@docker run --rm -v "$${PWD}":/go/src/github.com/VEVO/$(GO_PIPELINE_NAME) -w /go/src/github.com/VEVO/$(GO_PIPELINE_NAME) -e GOARCH=amd64 -e GOOS=linux -e CGO_ENABLED=0 -e BUILD_VERSION=$(BUILD_VERSION) golang:alpine make go-build

build: docker-lint docker-build docker-tag docker-push

# vim: ft=make
