OUT := amidakuji
VERSION := $(shell git describe --always --long)
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

all: build build_windows

build:
	go build -i -v -o ${OUT} -ldflags "-w -s -X main.version=${VERSION}"

build_windows:
	go build -i -v -o ${OUT}.exe -ldflags "-w -s -X main.version=${VERSION} -H windowsgui"

run: build
	./${OUT}

test:
	@go test -short ${PKG_LIST}

vet:
	@go vet -copylocks=false ${PKG_LIST}

vet_annoying:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

#static: vet lint
#	go build -i -v -o ${OUT}-${VERSION} -ldflags "-extldflags \"-static\" -w -s -X main.version=${VERSION}"

#static_windows: vet lint
#	go build -i -v -o ${OUT}-${VERSION}.exe -ldflags "-extldflags \"-static\" -w -s -X main.version=${VERSION} -H windowsgui"

clean:
	-@rm ${OUT} ${OUT}.exe #${OUT}-*

.PHONY: build build_windows run vet vet_annoying lint clean
