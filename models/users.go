package models

import "time"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    *string
	Age      uint8
	Birthday *time.Time
}
