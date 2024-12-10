package database

import "time"

type Swipes struct {
	CreatedAt time.Time `gorm:"autoCreateTime;column=created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column=updated_at"`
	ID        string    `gorm:"type:char(26);primaryKey"`
	SwiperID  string    `gorm:"type:char(26);not null;column=swiper_id"`
	SwipedID  string    `gorm:"type:char(26);not null;column=swiped_id"`
	Direction string    `gorm:"not null;column=direction"`
	SwipedAt  time.Time `gorm:"autoCreateTime;column=swiped_at"`
	Swiper    *Users    `gorm:"foreignKey:SwiperID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Swiped    *Users    `gorm:"foreignKey:SwipedID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
