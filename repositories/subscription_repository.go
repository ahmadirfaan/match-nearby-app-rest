package repositories

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SubscriptionsRepository interface {
	SaveSubscription(model *database.Subscriptions) error
}

type subscriptionsRepository struct {
	DB *gorm.DB
}

func NewSubscriptionsRepository(db *gorm.DB) SubscriptionsRepository {
	return &subscriptionsRepository{
		DB: db,
	}
}

func (sr subscriptionsRepository) SaveSubscription(model *database.Subscriptions) error {

	//generate ID
	if model.ID == "" {
		model.ID = ulid.Make().String()
	}

	if err := sr.DB.Save(&model).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to save subscription")
		return err
	}

	return nil
}
