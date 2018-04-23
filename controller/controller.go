// Copyright 2017 Vladimir Poliakov. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package controller provides http handling.
package controller

import (
	goctx "context"

	"github.com/gin-gonic/gin"
)

// Controller describes routable controller of http group methods.
type Controller interface {
	Route(ctx goctx.Context, app *gin.Engine)
}
