package user

import (
	"github.com/unclebob/user/entities"
)

type Controller interface {
	Get(gr *entities.UserQuery) (res *entities.User, err error)
}
