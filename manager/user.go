// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package manager

import (
	goctx "context"
	"crypto/sha256"
	"fmt"

	"github.com/beetlefinder/go-server/context"
	"github.com/beetlefinder/go-server/dto"
)

// User is a user manager.
type User struct{}

// Create creates new user in DB by email and nick.
func (u User) Create(ctx goctx.Context, login string, pass string, nick string) (*dto.User, error) {
	db := context.DB(ctx)
	users := db.Table("user")
	auths := db.Table("auth")

	auth := Auth{}
	if auth.GetByLogin(ctx, login).Login != "" {
		// TODO: refactor to common error handling when realized.
		return nil, fmt.Errorf("user already exists")
	}

	// TODO: make data verification.
	users.Create(struct{ Nick string }{nick})
	auths.Create(struct {
		Login    string
		PassHash string
	}{
		login,
		fmt.Sprintf("%s", sha256.Sum256([]byte(pass))),
	})

	// TODO: return created user.
	return nil, nil
}

// GetByID gets user from DB by ID.
func (User) GetByID(ctx goctx.Context, id uint) (*dto.User, bool) {
	user := dto.User{}
	context.DB(ctx).Table("user").First(&user, id)

	return &user, user.ID != 0
}
