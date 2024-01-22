package user

import (
	"encoding/gob"
)

func RegisterSessions() {
	gob.Register(User{})
}
