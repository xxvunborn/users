package user

import (
	"github.com/unclebob/user/entities"
)

// Usercase interface
type Usercase interface {
	Get(gr *entities.UserQuery) (res *entities.User, err error)
}
