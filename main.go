// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/beetlefinder/go-server/config"
	"github.com/beetlefinder/go-server/context"
	"github.com/beetlefinder/go-server/controller"
	"github.com/beetlefinder/go-server/db"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Parse("./config/config.yml")
	start(cfg.Port, cfg.Driver, cfg.DBConnStr)
}

func start(port string, driver string, conn string) {
	db := db.Connect(driver, conn)
	ctx := context.New(db)

	app := gin.Default()
	controller.User{}.Route(ctx, app)

	app.Run(fmt.Sprintf(":%s", port))
}
