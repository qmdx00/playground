GOOS = linux
GOARCH = amd64
VERSION = $(shell git describe --tags --always)

build:
	mkdir -p output/ \
    && CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
	go build -ldflags "-X main.Version=${VERSION}" -o ./output/binary .
