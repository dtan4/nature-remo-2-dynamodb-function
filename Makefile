NAME     := nature-remo-2-dynamodb-function
VERSION  := v0.1.0
REVISION := $(shell git rev-parse --short HEAD)

SRCS     := $(shell find . -type f -name '*.go')
LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
NOVENDOR := $(shell go list ./... | grep -v vendor)

DIST_DIRS := find * -type d -exec

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): $(SRCS)
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: ci-test
ci-test:
	go test -coverpkg=./... -coverprofile=coverage.txt -v ./...

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u github.com/golang/dep/cmd/dep
endif

.PHONY: deps
deps: dep
	dep ensure -v

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: mockgen
mockgen:
	go get -v github.com/golang/mock/gomock
	go get -v github.com/golang/mock/mockgen
	mockgen -source vendor/github.com/aws/aws-sdk-go/service/s3/s3iface/interface.go -destination aws/mock/s3.go -package mock

.PHONY: test
test:
	go test -coverpkg=./... -v $(NOVENDOR)

.PHONY: update-deps
update-deps: dep
	dep ensure -update -v
