package repository

import (
	"gorm.io/gorm"
	"oneMarketing/internal/app/model"
)

type UserRepository interface {
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) UpdateUser(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *userRepository) DeleteUser(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}
