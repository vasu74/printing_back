package routes

import (
	"myproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {
	var client models.Client

	// bind json
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse client data"})
		return
	}

	if err := client.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messgae": "could not create client", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "client  created successfully", "tenant_id": client.ID})
}

func GetAllClient(c *gin.Context) {
	tenants, err := models.GetAllClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse client details "})
		return
	}
	c.JSON(http.StatusOK, tenants)
}
