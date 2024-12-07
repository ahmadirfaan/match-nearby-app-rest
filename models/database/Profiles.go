package database

import "time"

type Profiles struct {
	CreatedAt time.Time `gorm:"autoCreateTime;column=created_at" json:"-"`
	UpdatedAt time.Time `json:"-" gorm:"column=updated_at"`
	ID        string    `gorm:"type:char(26);primaryKey"`
	UserID    string    `gorm:"type:char(26);unique;not null;column=user_id"` // 1-to-1 relationship
	Name      string    `gorm:"not null;column=name"`
	Gender    string    `gorm:"column=gender;type:enum('MALE', 'FEMALE')"`
	Bio       string    `gorm:"column=bio"`
	PhotoURL  string    `gorm:"column=photo_url"`
	User      *Users    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
}
