OUT := amidakuji
FUNC := Test
ASSET_TARGET := glossary/asset.go
ASSET_SOURCE_DIR := assets
VERSION := $(shell git describe --always --long)

all: run

${ASSET_TARGET}: 
	go-bindata -o "${ASSET_TARGET}" -pkg "glossary" -prefix "${ASSET_SOURCE_DIR}" ${ASSET_SOURCE_DIR}/emoji ${ASSET_SOURCE_DIR}/karaoke ${ASSET_SOURCE_DIR}/NanumBarunGothic.ttf

build: ${ASSET_TARGET}
	go build -i -v -o ${OUT}.dll -buildmode=c-shared -ldflags "-w -s -X main.version=${VERSION}"

run: build
	rundll32.exe ${OUT}.dll ${FUNC}

clean:
	-@rm ${ASSET_TARGET} ${OUT}.dll ${OUT}.h #${OUT}-*

.PHONY: build run clean
