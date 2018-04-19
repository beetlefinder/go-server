// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package manager

import (
	"container/list"
	goctx "context"
	"fmt"

	"github.com/beetlefinder/go-server/context"
	"github.com/beetlefinder/go-server/dto"
)

// Alert is a alert manager.
type Alert struct{}

// Create creates new alert in db by userID and message.
func (a Alert) Create(ctx goctx.Context, userID uint, message string) error {
	db := context.DB(ctx).Table("alert")

	u := User{}
	if _, ok := u.GetByID(ctx, userID); !ok {
		return fmt.Errorf("User dosen't exists")
	}

	db.Create(struct {
		UserID  uint
		Message string
	}{userID, message})
	return nil
}

// GetByID gets alert from db by ID.
func (Alert) GetByID(ctx goctx.Context, id uint) (*dto.Alert, bool) {
	alert := dto.Alert{}
	context.DB(ctx).Table("alert").First(&alert, id)

	return &alert, alert.ID != 0
}

// GetList gets all alerts from db.
func (Alert) GetList(ctx goctx.Context) (*list.List, error) {
	alerts := list.New()

	rows, err := context.DB(ctx).Table("alert").Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		alert := dto.Alert{}
		err = rows.Scan(alert.ID, alert.UserID, alert.Message, alert.CreatedAt, alert.UpdatedAt, alert.DeletedAt)
		if err != nil {
			return alerts, err
		}
		alerts.PushBack(alert)
	}

	return alerts, nil
}
