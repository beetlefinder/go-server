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

// Alert is a alert controller - handleres served http request
type Alert struct{}

// Create handlerFunc for create alert.
func (Alert) Create(ctx goctx.Context) gin.HandlerFunc {
	alerts := manager.Alert{}

	return func(c *gin.Context) {
		alert := dto.Alert{}
		c.BindQuery(&alert)

		res, err := alerts.Create(ctx, alert.UserID, alert.Message)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, gin.H{"userID": res.UserID, "message": res.Message})
		return
	}
}

// Get handlerFunc for get alert by id.
func (Alert) Get(ctx goctx.Context) gin.HandlerFunc {
	alerts := manager.Alert{}

	return func(c *gin.Context) {
		alert := dto.User{}
		c.BindQuery(&alert)

		res, ok := alerts.GetByID(ctx, alert.ID)
		if !ok {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, gin.H{"userID": res.UserID, "message": res.Message})
		return
	}
}

// GetList handlerFunc for get alerts by user id.
func (Alert) GetList(ctx goctx.Context) gin.HandlerFunc {
	alerts := manager.Alert{}

	return func(c *gin.Context) {
		alert := dto.User{}
		c.BindQuery(&alert)

		res := alerts.GetList(ctx, alert.ID)

		c.JSON(http.StatusOK, res)
		return
	}
}
