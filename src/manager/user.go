// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package manager

import (
	goctx "context"
	"fmt"

	"github.com/beetlefinder/go-server/src/context"
	"github.com/beetlefinder/go-server/src/dto"
)

// User is a user manager.
type User struct{}

// Create creates new user in db by email and nick.
func (u User) Create(ctx goctx.Context, login string, pass string, nick string) error {
	db := context.DB(ctx).Table("user")

	a := Auth{}
	if a.GetByLogin(ctx, login).Login != "" {
		return fmt.Errorf("user already exists")
	}

	db.Create(struct{ Nick string }{nick})
	return nil
}

// GetByID gets user from db by ID.
func (User) GetByID(ctx goctx.Context, id uint) (*dto.User, bool) {
	user := dto.User{}
	context.DB(ctx).Table("user").First(&user, id)

	return &user, user.ID != 0
}
