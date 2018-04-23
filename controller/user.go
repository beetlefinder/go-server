// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	goctx "context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/beetlefinder/go-server/manager"
)

// User is a user controller - handleres served http request
// in context of user changing / getting.
type User struct{}

// Route implements Controller interface.
func (user User) Route(ctx goctx.Context, app *gin.Engine) {
	userURI := app.Group("/user")

	userURI.POST("/", user.Create(ctx))
	userURI.GET("/:id", user.Get(ctx))
}

// Create handlerFunc for create user.
//
// JSON struct as parameter.
func (User) Create(ctx goctx.Context) gin.HandlerFunc {
	users := manager.User{}

	// TODO: refactor struct hardcode.
	return func(c *gin.Context) {
		user := struct {
			Login string
			Pass  string
			Nick  string
		}{}
		c.BindQuery(&user)

		res, err := users.Create(ctx, user.Login, user.Pass, user.Nick)
		if err != nil {
			c.AbortWithStatus(http.StatusConflict)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": res})
		return
	}
}

// Get returns handler for user getting.
//
// One uint /:id parameter.
func (User) Get(ctx goctx.Context) gin.HandlerFunc {
	users := manager.User{}

	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		res, ok := users.GetByID(ctx, uint(id))
		if !ok {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": res})
		return
	}
}
