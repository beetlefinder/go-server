// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	goctx "context"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/beetlefinder/go-server/manager"
)

// Auth is a authorize controller - handleres served http request
type Auth struct{}

// Authorize handlerFunc check login and password.
func (Auth) Authorize(ctx goctx.Context) gin.HandlerFunc {
	auths := manager.Auth{}

	return func(c *gin.Context) {
		auth := struct {
			Login string
			Pass  string
		}{}
		c.BindQuery(&auth)

		res, ok := auths.GetByLogin(ctx, auth.Login)
		if !ok || bcrypt.CompareHashAndPassword([]byte(res.PasswordHash), []byte(auth.Pass)) != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": res})
		return
	}
}
