.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/hkd-org/go-hkd/cmd/swarm
	go build -o bin/geth     github.com/hkd-org/go-hkd/cmd/geth
	go build -o bin/bootnode github.com/hkd-org/go-hkd/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/hkd-org/go-hkd/swarm/...
