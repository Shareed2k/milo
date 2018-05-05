PROJECT?=github.com/milo
APP?=milo

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CONTAINER_IMAGE?=shareed2k/${APP}
ITERATION?=alpha

GOOS?=$(shell go env GOOS)
GOARCH?=amd64

clean:
	rm -f ${APP}

bindata: clean ui
	rice embed-go -i ./internal/

ui:
	yarn --cwd ./ui build

grpc:
	protoc -I=./grpc/ --go_out=plugins=grpc:./internal/ ./grpc/*.proto

build: bindata
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/internal.Release=${RELEASE} \
		-X ${PROJECT}/internal.Commit=${COMMIT} -X ${PROJECT}/internal.BuildTime=${BUILD_TIME}" \
		-o ${APP} \
		./main.go

.PHONY: clean bindata build ui grpc