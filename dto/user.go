package dto

import "github.com/jinzhu/gorm"

// User describes user information.
type User struct {
	gorm.Model
	Nick string
}
