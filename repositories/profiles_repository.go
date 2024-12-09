package repositories

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProfilesRepository interface {
	SaveProfile(user *database.Profiles) error
}

type profilesRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfilesRepository {
	return &profilesRepository{
		DB: db,
	}
}

func (profilesRepository *profilesRepository) SaveProfile(profile *database.Profiles) error {
	profile.ID = ulid.Make().String()
	if err := profilesRepository.DB.Save(&profile).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to save profile")
		return err
	}

	return nil
}
