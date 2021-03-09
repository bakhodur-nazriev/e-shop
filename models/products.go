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

type CreateProductInput struct {
	Title       string  `json:"title" binding:"required"`
	Code        string  `json:"code" binding:"required"`
	Size        uint    `json:"size" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description" binding:"required"`
}
