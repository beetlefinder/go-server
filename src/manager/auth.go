// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package manager

import (
	goctx "context"

	"github.com/beetlefinder/go-server/src/dto"
)

// Auth is a authentication manager.
type Auth struct{}

// GetByLogin .
func (a Auth) GetByLogin(ctx goctx.Context, login string) dto.Auth {
	return dto.Auth{}
}
