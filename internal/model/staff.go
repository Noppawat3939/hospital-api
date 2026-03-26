package model

import "time"

type Staff struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	HospitalID string    `gorm:"type:uuid;not null;index" json:"hospital_id"`
	Username   string    `gorm:"unique;not null" json:"username"`
	Password   string    `gorm:"not null" json:"-"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
