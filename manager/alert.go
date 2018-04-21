// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package manager

import (
	goctx "context"
	"fmt"

	"github.com/beetlefinder/go-server/context"
	"github.com/beetlefinder/go-server/dto"
)

// Alert is a alert manager.
type Alert struct{}

// Create creates new alert in db by userID and message.
//
// TODO: create method should return created dto.
func (a Alert) Create(ctx goctx.Context, userID uint, message string) error {
	db := context.DB(ctx).Table("alert")

	u := User{}
	// TODO: error handling.
	if _, ok := u.GetByID(ctx, userID); !ok {
		return fmt.Errorf("User dosen't exists")
	}

	// TODO: it will be fuck up here - string can't be nil.
	alert := dto.Alert{
		UserID:  userID,
		Message: message,
	}
	db.Create(&alert)

	return nil
}

// GetByID gets alert from db by ID.
func (Alert) GetByID(ctx goctx.Context, id uint) (*dto.Alert, bool) {
	alert := dto.Alert{}
	context.DB(ctx).Table("alert").First(&alert, id)

	return &alert, alert.ID != 0
}

// GetList gets alerts from db by userID.
func (Alert) GetList(ctx goctx.Context, userID uint) []dto.Alert {
	alerts := context.DB(ctx).Table("alert")

	var res []dto.Alert
	alerts.Where("user_id = ?", userID).Find(&res)

	return res
}
