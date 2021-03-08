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
