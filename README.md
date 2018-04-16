# BeetleFinder/Go-Server

[![Documentation][godoc-badge]][godoc]
[![Build Status][travis-badge]][travis]
[![Test Coverage][codecov-badge]][codecov]
[![Code Quality][go-report-card-badge]][go-report-card]

Server for beetlefinder alert service.

## Requirements

* Go 1.10 (but 1.6+ should be ok)
* PostgreSQL 10 (but 9+ should be ok)

For developers it's recommended to create PostgreSQL user with login `postgres` and password `postgres` and make it able on `localhost:5432`. Or every time use `make recreate` with `CONN_STR=postgres://custom:12345678@localhost:5432`.

## Installation

`%GOPATH%` for Windows and `$GOPATH` for Unix-style OS (Linux, Mac).

```sh
go get github.com/beetlefinder/go-server
cd %GOPATH%/src/github.com/beetlefinder/go-server
make get
make recreate
make
```

`make` command depends on `make get` (one time at installation only), so don't skip `make get` command.

## Build

```sh
make
```

Or with parameters:

```sh
make BIN_NAME=main BIN_PATH=./
```

Parameters by default:

* BIN_NAME=beetle_finder_go_server
* BIN_PATH=./bin

## Run

```sh
make run
```

## Create or recreate dev DB

```sh
make recreate
```

Or with parameters:

```sh
make recreate CONN_STR=postgres://custom:12345678@localhost:5432
```

Parameters by default:

* CONN_STR=postgres://postgres:postgres@localhost:5432

[godoc]: https://godoc.org/github.com/beetlefinder/go-server
[travis]: https://travis-ci.org/beetlefinder/go-server
[codecov]: https://codecov.io/gh/beetlefinder/go-server
[go-report-card]: https://goreportcard.com/report/github.com/beetlefinder/go-server

[godoc-badge]: https://godoc.org/github.com/beetlefinder/go-server?status.svg
[travis-badge]: https://travis-ci.org/beetlefinder/go-server.svg?branch=develop
[codecov-badge]: https://codecov.io/gh/beetlefinder/go-server/branch/develop/graph/badge.svg
[go-report-card-badge]: https://goreportcard.com/badge/github.com/beetlefinder/go-server
