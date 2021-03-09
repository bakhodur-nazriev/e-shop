package controllers

import (
	"e-shop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get /products
// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// POST /products
// Create new products
func CreateProduct(c *gin.Context) {
	//	Validate input
	var input models.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//	Create book
	product := models.Product{
		Title:       input.Title,
		Code:        input.Code,
		Size:        input.Size,
		Price:       input.Price,
		Description: input.Description,
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func GetProduct(c *gin.Context) {
	//	Get model if exist
	var product models.Product

	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
