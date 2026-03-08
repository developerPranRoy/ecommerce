// ProductHandler.go

package handler

import (
	"net/http"	
	"github.com/gin-gonic/gin"
)

// Product represents the product model
type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Sample product storage
var products = map[string]Product{}

// GetProduct retrieves a single product by ID
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, exists := products[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct adds a new product
func CreateProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	products[newProduct.ID] = newProduct
	c.JSON(http.StatusCreated, newProduct)
}

// UpdateProduct modifies an existing product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, exists := products[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	products[id] = updatedProduct
	c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct removes a product by ID
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, exists := products[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	delete(products, id)
	c.JSON(http.StatusNoContent, nil)
}