package models

type Categories struct {
	ID    uint `gorm:"primaryKey"`
	Title string
}
