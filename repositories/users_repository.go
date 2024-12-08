package repositories

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UsersRepository interface {
	SaveUser(user *database.Users) error
}

type usersRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{
		DB: db,
	}
}

func (usersRepository *usersRepository) SaveUser(user *database.Users) error {
	user.ID = ulid.Make().String()
	if err := usersRepository.DB.Create(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to save user")
		return err
	}

	return nil
}
