package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ahmadirfaan/match-nearby-app-rest/config/storage"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/oklog/ulid/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var redisClient = storage.InitRedis()

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

	saveToCache(user.ID, user, context.Background())
	return nil
}

func (usersRepository *usersRepository) GetByUsername(username string) *database.Users {
	var user *database.Users
	if err := usersRepository.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}
	saveToCache(user.ID, user, context.Background())
	return user

}

func (usersRepository *usersRepository) GetByEmail(email string) *database.Users {
	var user *database.Users
	if err := usersRepository.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}
	saveToCache(user.ID, user, context.Background())
	return user
}

func (usersRepository *usersRepository) GetByUserId(id string) *database.Users {
	var user *database.Users

	contextBackground := context.Background()
	result, err := redisClient.Get(contextBackground, fmt.Sprintf("USER_ID_%s", id)).Result()
	if errors.Is(err, redis.Nil) || err != nil {
		if err := usersRepository.DB.Preload("Profile").Where("id = ?", id).First(&user).Error; err != nil {
			logrus.Error("Failed to get user by id")
			return nil
		}
		saveToCache(id, user, contextBackground)
	} else {
		err = json.Unmarshal([]byte(result), &user)
		if err != nil {
			logrus.Error("Failed to unmarshal user")
			return nil
		}
		logrus.Infof("GET USER_ID ON REDIS: %s", id)
	}

	return user
}

func saveToCache(id string, user *database.Users, contextBackground context.Context) {
	userJSON, err := json.Marshal(user)
	if err != nil {
		logrus.Error("Failed to marshal user")
	}
	err = redisClient.Set(contextBackground, fmt.Sprintf("USER_ID_%s", id), userJSON, 0).Err()
	if err != nil {
		logrus.Fatalf("Failed  save to Redis: %v", id)
	}
	logrus.Infof("SAVE USER_ID ON REDIS: %s", id)
}
