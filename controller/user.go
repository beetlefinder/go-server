// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	goctx "context"
	"net/http"

	"github.com/beetlefinder/go-server/dto"

	"github.com/gin-gonic/gin"

	"github.com/beetlefinder/go-server/manager"
)

// User is a user controller - handleres served http request
// in context of user changing / getting.
type User struct{}

// Create handlerFunc for create user.
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

// Get handlerFunc for get user by id.
func (User) Get(ctx goctx.Context) gin.HandlerFunc {
	users := manager.User{}

	return func(c *gin.Context) {
		user := dto.User{}
		c.BindQuery(&user)

		res, ok := users.GetByID(ctx, user.ID)
		if !ok {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, gin.H{"nick": res.Nick})
		return
	}
}
