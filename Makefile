PREFIX  := quay.io/roboll/etcd-metrics-proxy
VERSION := $(shell git describe --tags --abbrev=0 HEAD)

check:
	go vet .
.PHONY: check

test:
	go test -v . -cover -race -p=1
.PHONY: test

build:
	GOOS=linux go build -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo -o etcd-metrics-proxy .
.PHONY: build

container: build
	docker build -t $(PREFIX):$(VERSION) .
.PHONY: container

push: container
	docker push $(PREFIX):$(VERSION)
.PHONY: push
