package models

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Code        string
	Size        uint
	Price       float64
	Description string
	Category    int `gorm:"references:ID"`
}

type Category struct {
	ID    int `gorm:"primaryKey"`
	Title string
}
