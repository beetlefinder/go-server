.PHONY: all get vet build run test

GOOS := $(GOOS)
GOARCH := $(GOARCH)

ifndef BIN_NAME
    BIN_NAME = beetle_finder_go_server
endif

SRC_PATH := .
SRC := $(SRC_PATH)/...
SRC_RUN := $(SRC_PATH)/main.go

ifndef CONN_STR
    CONN_STR := postgres://postgres:postgres@localhost:5432
endif
RECREATE_SQL := $(SRC_PATH)/db/sql/develop/recreate_develop.sql

all: build

get:
	@go get -v $(SRC)
	@go get -v golang.org/x/lint/golint

vet:
	@go vet $(SRC)

golint: vet
	@golint $(SRC)

gofmt: golint
	@gofmt -l $(SRC_PATH)

build: gofmt
	@go build -o $(BIN_NAME) -v $(SRC_PATH)

run:
	@go run $(SRC_RUN)

test:
	@go test -v --race --covermode=atomic --coverprofile=coverage.txt $(SRC)

recreate:
	@psql -f $(RECREATE_SQL) $(CONN_STR)
