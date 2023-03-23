package service

import (
	"oneMarketing/internal/app/model"
)

type UserService interface {
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *model.User) error
}

//type userService struct {
//}
