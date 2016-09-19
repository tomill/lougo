help:
	cat Makefile

assets.go:
	go-assets-builder --package=lougo data/loudic.txt > assets.go

test:
	goimports -w *.go
	golint
	go vet
	go test
