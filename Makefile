.PHONY: build

help:
	cat Makefile

setup:
	go get github.com/ikawaha/kagome/...
	go get github.com/jessevdk/go-assets-builder

assets.go:
	go-assets-builder --package=lougo data/loudic.txt > assets.go

test:
	goimports -w *.go cmd
	golint
	go vet
	go test

build:
	rm -f build/*
	
	GOOS=linux GOARCH=amd64 go build -o build/lougo cmd/lougo/main.go
	cd build; tar czf lougo-Linux-x86_64.tar.gz lougo
	
	GOOS=darwin GOARCH=amd64 go build -o build/lougo cmd/lougo/main.go
	cd build; zip lougo-Darwin-i386.zip lougo
