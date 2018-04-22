.PHONY: all get vet build run test

GOOS := $(GOOS)
GOARCH := $(GOARCH)

ifndef BIN
    BIN = beetlefinder_go_server
    ifndef GOOS
        ifeq ($(OS), Windows_NT)
            BIN := $(BIN).exe
        endif
    endif
    ifeq ($(GOOS), windows)
        BIN := $(BIN).exe
    endif
endif

SRC_PATH := .
SRC := $(SRC_PATH)/...
SRC_RUN := $(SRC_PATH)/main.go $(SRC_PATH)/server.go

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

golint: vet
	@golint $(SRC)

gofmt: golint
	@gofmt -l $(SRC_PATH)

build: gofmt
	@go build -o $(BIN) -v $(SRC_PATH)

run:
	@go run $(SRC_RUN)

daemon:
	@CompileDaemon -include=Makefile -build="make BIN=$(BIN)" -command="./$(BIN)"

test:
	@go test -v --race --covermode=atomic --coverprofile=coverage.txt $(SRC)

recreate:
	@psql -f $(RECREATE_SQL) $(CONN_STR)
