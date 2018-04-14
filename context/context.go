// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package context

import (
	goctx "context"
	"log"
)

type key int

// Context keys.
const (
	KeyDB key = iota
)

// TODO: change `interface{}` to DB struct when realized.

// New initializes base context.
func New(db interface{}) goctx.Context {
	ctx := goctx.Background()
	ctx = goctx.WithValue(ctx, KeyDB, db)

	return ctx
}

// DB returns db instance by ctx.
func DB(ctx goctx.Context) interface{} {
	db, ok := ctx.Value(KeyDB).(interface{})
	if !ok {
		log.Fatalf("context: has no db")
	}

	return db
}
