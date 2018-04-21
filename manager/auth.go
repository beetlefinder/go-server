// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package manager

import (
	goctx "context"

	"github.com/beetlefinder/go-server/context"
	"github.com/beetlefinder/go-server/dto"
)

// Auth is a authentication manager.
type Auth struct{}

// Create .
//
// TODO: create method should return created dto.
func (a Auth) Create(ctx goctx.Context, login string, passHash string) error {
	auths := context.DB(ctx).Table("auth")

	// TODO: verification here or at user creation?
	auth := dto.Auth{
		Login:        login,
		PasswordHash: passHash,
	}

	auths.Create(&auth)

	return nil
}

// GetByLogin .
func (a Auth) GetByLogin(ctx goctx.Context, login string) (*dto.Auth, bool) {
	auth := dto.Auth{Login: login}
	context.DB(ctx).Table("auth").Where("login = ?", login).First(&auth)

	return &auth, auth.ID != 0
}
