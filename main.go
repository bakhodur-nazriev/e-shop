package main

import (
	"e-shop/controllers"
	"e-shop/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := gin.Default()

	models.ConnectToDatabase()

	r.GET("/", controllers.GetProducts)

	r.Run()

	//db, err := gorm.Open("postgres", "user=bakhodur password=1996 dbname=e-shop sslmode=disable")
	//if err != nil {
	//	panic("Failed to connect database")
	//}
	//defer db.Close()
	//
	///* Migrate the schema */
	//db.AutoMigrate(&models.Product{}, &models.Categories{}, &models.User{})
	//
	///* Create */
	//db.Create(&models.Product{
	//	Title: "Boots",
	//	Code:  "L1212",
	//	Size:  32,
	//	Price: 150,
	//})

	/* Read */
	//var product models.Products
	//db.First(&product, 1)                   //Find product with id 1
	//db.First(&product, "code = ?", "L1212") //Find product with code L1212
	//
	///* Update update product's price to 2000 */
	//db.Model(&product).Update("Price", 250)
	//
	///* Delete - delete product */
	//db.Delete(&product)
}
