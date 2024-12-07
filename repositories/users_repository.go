package repositories

import "gorm.io/gorm"

type UsersRepository interface {
}

type usersRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{
		DB: db,
	}
}
