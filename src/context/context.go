// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package context

import (
	goctx "context"
	"log"

	"github.com/beetlefinder/go-server/src/db"
)

type key int

// Context keys.
const (
	KeyDB key = iota
)

// New initializes base context.
func New(db *db.DB) goctx.Context {
	ctx := goctx.Background()
	ctx = goctx.WithValue(ctx, KeyDB, db)

	return ctx
}

// DB returns db instance by ctx.
func DB(ctx goctx.Context) *db.DB {
	ctxDB, ok := ctx.Value(KeyDB).(*db.DB)
	if !ok {
		log.Fatalf("TODO: change string to internal error")
	}

	return ctxDB
}
