PREFIX  := quay.io/roboll/etcd-metrics-proxy
VERSION := 0.1

check:
	go vet ${PKGS}
.PHONY: check

test:
	go test -v ${PKGS} -cover -race -p=1
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
