package todo

import (
	"api/internal/user"
	"errors"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string
	Done   bool
	UserID uint      `gorm:"not null;->;<-:create"`
	User   user.User `gorm:"foreignKey:UserID"`
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	if t.UserID == 0 {
		err = errors.New("UserID must be set")
	}
	return
}

func (t *Todo) ValidateAccess(authUser user.User) error {
	if authUser.IsSuperUser() {
		return nil
	}

	if t.UserID == authUser.ID {
		return nil
	}

	return errors.New("You are not authorized to view this users todos")
}
