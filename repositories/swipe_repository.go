package repositories

import (
	"github.com/ahmadirfaan/match-nearby-app-rest/app"
	"github.com/ahmadirfaan/match-nearby-app-rest/models/database"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

var LayoutFormat = "2006-01-02 15:04:05.00"
var now = time.Now()
var yesterday = now.AddDate(0, 0, -1).Format(LayoutFormat)
var today = now.Format(LayoutFormat)

type SwipeRepository interface {
	SaveSwipe(model *database.Swipes) error
	GetSwipeNearby(userID string) []database.Profiles
	GetSwipeStatus(swiperID string, swipedID string) []database.Swipes
}

type swipeRepository struct {
	DB *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) SwipeRepository {
	return &swipeRepository{
		DB: db,
	}
}

func (sr *swipeRepository) GetSwipeStatus(swiperID string, swipedID string) []database.Swipes {
	var swipes []database.Swipes

	sr.DB.Debug().Table("swipes").Where("swiper_id = ? AND swiped_id = ? AND swiped_at BETWEEN ? AND ?", swiperID, swipedID, yesterday, today).
		Or("swiper_id = ? AND swiped_id = ? AND direction = ?", swiperID, swipedID, "Like").Find(&swipes)
	return swipes
}

func (sr *swipeRepository) GetSwipeNearby(userID string) []database.Profiles {

	var profiles []database.Profiles

	sr.DB.Table("profiles").
		Where("user_id NOT IN (?)",
			sr.DB.Table("swipes").
				Select("swiped_id").
				Where("swiper_id = ? AND DATE(swiped_at) BETWEEN ? AND ?", userID, yesterday, today).Or("swiper_id = ? and direction = 'Like'"),
		).
		Limit(app.Init().Config.MaximalTotalSwipe).
		Find(&profiles)

	return profiles
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
