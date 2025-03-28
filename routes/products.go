package routes

import (
	"myproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProducts(c *gin.Context) {
	var product models.Product

	// Bind JSON data to tenant struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	// Save the tenant to the database
	if err := product.SaveProduct(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create Product",
			"details": err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Tenant created successfully", "tenant_id": product.ID})
}

// for calculating price
func CalculatePriceHandler(c *gin.Context) {
	var request models.PriceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Calculate price dynamically
	finalPrice, err := request.CalculatePrice(request.ProductID, request.Parameters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"product_id": request.ProductID,
		"parameters": request.Parameters,
		"price":      finalPrice,
	})

}
