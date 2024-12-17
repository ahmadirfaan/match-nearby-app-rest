package repositories

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SwipeRepository interface {
	SaveSwipe(model *database.Swipes) error
}

type swipeRepository struct {
	DB *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) SwipeRepository {
	return &swipeRepository{
		DB: db,
	}
}

func (sr *swipeRepository) SaveSwipe(model *database.Swipes) error {

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
