// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	goctx "context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/beetlefinder/go-server/manager"
)

// User is a user controller - handleres served http request
// in context of user changing / getting.
type User struct{}

// Create .
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

	return nil
}
