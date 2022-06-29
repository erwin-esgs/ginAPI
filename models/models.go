package models

import "time"

type Username struct {
	// gorm.Model
	ID           int64 `gorm:"primaryKey"`
	NamaDepan    string
	NamaBelakang string
	Email        string
	Password     string
}

type Product struct {
	// gorm.Model
	ID          int64 `gorm:"primaryKey"`
	NamaProduct string
	Kuantity    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
