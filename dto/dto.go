// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package dto provides data containers (DTO pattern)
// to transfer information between server and DB.
package dto

import "github.com/jinzhu/gorm"

// User describes user data.
type User struct {
	gorm.Model
	Nick string
}

// Auth describes authentication data.
type Auth struct {
	gorm.Model
	Login    string
	PassHash string
}

// Alert describes alert information.
type Alert struct {
	gorm.Model
	UserID  uint
	Message string
}
