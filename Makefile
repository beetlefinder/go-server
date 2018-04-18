.PHONY: all get vet build run test

GOOS := $(GOOS)
GOARCH := $(GOARCH)

ifndef BIN_NAME
    BIN_NAME = beetle_finder_go_server
endif

ifndef CONN_STR
    CONN_STR := postgres://postgres:postgres@localhost:5432
endif
RECREATE_SQL := ./db/sql/develop/recreate_develop.sql

all: build

get:
	@go get -v ./...
	@go get -v golang.org/x/lint/golint

vet:
	@go vet ./...

golint: vet
	@golint ./...

gofmt: golint
	@gofmt -l .

build: gofmt
	@go build -o $(BIN_NAME) -v .

run:
	@go run main.go

test:
	@go test -v --race --covermode=atomic --coverprofile=coverage.txt ./...

recreate:
	@psql -f $(RECREATE_SQL) $(CONN_STR)
