package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type UserRole string

func (r *UserRole) Scan(value string) error {
	*r = UserRole([]byte(value))
	return nil
}

func (r UserRole) Value() (driver.Value, error) {
	return string(r), nil
}

const (
	AdminUserRole     UserRole = "admin"
	ModeratorUserRole UserRole = "moderator"
	UserUserRole      UserRole = "user"
)

type User struct {
	gorm.Model
	Name  string
	Email string   `gorm:"unique"`
	Role  UserRole `gorm:"type:user_role;default:'user'"`
}
