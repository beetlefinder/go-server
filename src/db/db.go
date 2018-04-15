// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package db provides sql querying.
package db

import (
	"log"

	"github.com/jinzhu/gorm"
	// Database driver initializing.
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB embedding used to extend standard methods.
type DB struct {
	*gorm.DB
}

// Connect connects to the database.
//
// This is a generic solution, so it's possible to connect to
// all SQL DB supported by `gorm`.
//
// Connection example:
// db.Connect(
// 	"postgres",
// 	"postgresql://postgres:postgres@localhost:5432/beetlefinder?sslmode=disable")
func Connect(driver string, conn string) *DB {
	db, err := gorm.Open(driver, conn)
	if err != nil {
		// TODO: change to internal error handling.
		log.Fatal(err)
	}

	return &DB{db}
}
