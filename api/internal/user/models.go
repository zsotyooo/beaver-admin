package user

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
	UserRoleAdmin     UserRole = "admin"
	UserRoleModerator UserRole = "moderator"
	UserRoleUser      UserRole = "user"
)

type User struct {
	gorm.Model
	Name  string
	Email string   `gorm:"unique"`
	Role  UserRole `gorm:"type:user_role;default:'user'"`
}

func (u *User) HasRole(roles []UserRole) bool {
	for _, role := range roles {
		if u.Role == role {
			return true
		}
	}
	return false
}

func (u *User) IsSuperUser() bool {
	return u.HasRole([]UserRole{UserRoleAdmin, UserRoleModerator})
}
