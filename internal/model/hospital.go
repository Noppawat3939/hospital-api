package model

import "time"

type Hospital struct {
	ID        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
