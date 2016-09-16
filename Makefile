PREFIX  := quay.io/roboll/etcd-metrics-proxy
VERSION := 0.1

build:
	GOOS=linux go build -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo -o etcd-metrics-proxy .
.PHONY: build

container: build
	docker build -t $(PREFIX):$(VERSION) .
.PHONY: container

release: container
	docker push $(PREFIX):$(VERSION)
.PHONY: release
