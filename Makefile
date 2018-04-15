.PHONY: all get vet build run test

GOOS := $(GOOS)
GOARCH := $(GOARCH)

ifndef BIN_NAME
    BIN_NAME = beetle_finder_go_server
endif
ifndef BIN_PATH
    BIN_PATH := ./bin
endif
BIN := $(BIN_PATH)/$(BIN_NAME)

SRC_PATH := ./src
SRC := $(SRC_PATH)/...

ifndef CONN_STR
    CONN_STR := postgres://postgres:postgres@localhost:5432
endif
RECREATE_SQL := ./db/sql/develop/recreate_develop.sql

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
	@go build -o $(BIN) -v $(SRC_PATH)
	@echo Built into $(BIN)

run:
	@go run $(SRC_PATH)/main.go

test:
	@go test -v --race --covermode=atomic --coverprofile=coverage.txt $(SRC)

recreate:
	@psql -f $(RECREATE_SQL) $(CONN_STR)
