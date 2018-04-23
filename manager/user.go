// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package manager

import (
	goctx "context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/beetlefinder/go-server/context"
	"github.com/beetlefinder/go-server/dto"
)

// User is a user manager.
type User struct{}

// Create creates new user in DB by email and nick.
func (u User) Create(ctx goctx.Context, login string, pass string, nick string) (*dto.User, error) {
	db := context.DB(ctx)
	users := db.Table("user")

	if _, ok := new(Auth).GetByLogin(ctx, login); ok {
		// TODO: rewrite to common error handling when realized.
		return nil, fmt.Errorf("user already exists")
	}

	// TODO: make data verification.
	user := dto.User{Nick: nick}
	users.Create(&user)
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	if err = new(Auth).Create(ctx, user.ID, login, string(passHash)); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetByID gets user from DB by ID.
func (User) GetByID(ctx goctx.Context, id uint) (*dto.User, bool) {
	user := dto.User{}
	context.DB(ctx).Table("user").First(&user, id)

	return &user, user.ID != 0
}
