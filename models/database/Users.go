package database

import (
	"time"
)

type Users struct {
	CreatedAt     time.Time       `gorm:"autoCreateTime;column=created_at" json:"-"`
	UpdatedAt     time.Time       `json:"-" gorm:"column=updated_at"`
	ID            string          `gorm:"type:char(26);primaryKey;column=id"`
	Username      string          `gorm:"unique;not null;column=username"`
	Email         string          `gorm:"unique;not null;column=email"`
	Password      string          `gorm:"not null;column=password"`
	IsPremium     bool            `gorm:"default:false;column=is_premium"`
	PremiumExpiry *time.Time      `gorm:"column=premium_expiry"`
	Profile       Profiles        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	Subscriptions []Subscriptions `gorm:"foreignKey:UserID"`
	Swipes        []Swipes        `gorm:"foreignKey:SwiperID"`
	SwipedBy      []Swipes        `gorm:"foreignKey:SwipedID"`
}
