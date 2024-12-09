package repositories

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UsersRepository interface {
	SaveUser(user *database.Users) error
	GetByUsername(username string) *database.Users
	GetByEmail(email string) *database.Users
	GetByUserId(id string) *database.Users
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

	if user.ID == "" {
		user.ID = ulid.Make().String()
	}

	if err := usersRepository.DB.Save(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to save user")
		return err
	}

	return nil
}

func (usersRepository *usersRepository) GetByUsername(username string) *database.Users {
	var user *database.Users
	if err := usersRepository.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}
	return user

}

func (usersRepository *usersRepository) GetByEmail(email string) *database.Users {
	var user *database.Users
	if err := usersRepository.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}
	return user
}

func (usersRepository *usersRepository) GetByUserId(id string) *database.Users {
	var user *database.Users
	if err := usersRepository.DB.Preload("Profile").Where("id = ?", id).First(&user).Error; err != nil {
		return nil
	}
	return user
}
