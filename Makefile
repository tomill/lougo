MECABCONFIG:=mecab-config
LOU5DIST=https://github.com/tomill/Acme-Lou/archive/0.04.tar.gz

.PHONY: lou.dic

help:
	cat Makefile

setup: lou.dic
	export CGO_LDFLAGS="`$(MECABCONFIG) --libs`"; \
	export CGO_FLAGS="`$(MECABCONFIG) --inc-dir`"; \
		go get github.com/shogo82148/go-mecab

lou.dic:
	curl -sL $(LOU5DIST) | tar -xzO "*/share/utf8.dic" > lou.dic

test:
	goimports -w *.go
	golint
	go vet
	go test
