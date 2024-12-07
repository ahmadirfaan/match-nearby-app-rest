package database

import "time"

type Profiles struct {
	CreatedAt time.Time `gorm:"autoCreateTime;column=created_at" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column=updated_at" json:"-"`
	ID        string    `gorm:"type:char(26);primaryKey;not null;column=id"`
	UserID    string    `gorm:"type:char(26);unique;not null;column=user_id"` // Explicitly set to snake_case
	Name      string    `gorm:"not null;column=name"`
	Gender    string    `gorm:"type:gender_enum;column=gender"`
	Bio       string    `gorm:"column=bio"`
	PhotoURL  string    `gorm:"column=photo_url"`
	User      *Users    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Ensure foreignKey is explicitly linked
}
