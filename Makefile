.PHONY: all get vet build run test

GOOS := $(GOOS)
GOARCH := $(GOARCH)

BIN := go-server

SRC_PATH := .
SRC := $(SRC_PATH)/...

ifndef CONN_STR
    CONN_STR := postgres://postgres:postgres@localhost:5432
endif
RECREATE_SQL := $(SRC_PATH)/db/sql/develop/recreate_develop.sql

all: build

get:
	@go get -v $(SRC)
	@go get -v golang.org/x/lint/golint
	@go get -v github.com/githubnemo/CompileDaemon

vet:
	@go vet $(SRC)

golint:
	@golint $(SRC)

gofmt:
	@gofmt -l $(SRC_PATH)

format: vet golint gofmt

build:
	@go build -v $(SRC_PATH)

run:
	@CompileDaemon -include=Makefile -build="make" -command="./$(BIN)"

test:
	@go test -v --race --covermode=atomic --coverprofile=coverage.txt $(SRC)

recreate:
	@psql -f $(RECREATE_SQL) $(CONN_STR)
