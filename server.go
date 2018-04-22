// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	goctx "context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/beetlefinder/go-server/context"
	"github.com/beetlefinder/go-server/controller"
	"github.com/beetlefinder/go-server/db"
)

func start(port string, driver string, conn string) {
	database := db.Connect(driver, conn)
	ctx := context.New(database)

	app := gin.Default()
	api{}.route(ctx, app)

	fullPort := fmt.Sprintf(":%s", port)
	if err := app.Run(fullPort); err != nil {
		log.Fatal(err)
	}
}

type api struct{}

func (api) route(ctx goctx.Context, app *gin.Engine) {
	user := controller.User{}
	userURI := app.Group("/user")
	userURI.POST("/", user.Create(ctx))
	userURI.GET("/:id", user.Get(ctx))
}
