package controller

import (
	"time"

	"github.com/unclebob/user"
	"github.com/unclebob/user/entities"
)

type UserController struct {
	User           user.Usercase
	contextTimeout time.Duration
}

func (u *UserController) Get(gr *entities.UserQuery) (*entities.User, error) {
	res, err := u.Get(gr)
	if err != nil {
		return nil, err
	}

	return res, nil
}
