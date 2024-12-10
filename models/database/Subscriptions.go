package database

import "time"

type Subscriptions struct {
	CreatedAt    time.Time `gorm:"autoCreateTime;column=created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;column=updated_at"`
	ID           string    `gorm:"type:char(26);primaryKey"`
	UserID       string    `gorm:"type:char(26);not null;column=user_id"`
	PurchaseName string    `gorm:"not null;column=purchase_name"`
	PurchaseDate time.Time `gorm:"not null;column=purchase_date"`
	User         *Users    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
