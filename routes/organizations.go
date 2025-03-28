package routes

import (
	"log"
	"myproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrganization(c *gin.Context) {
	var organization models.Organization
	// bind json
	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "could not parse data"})
		return
	}

	if err := organization.Save(); err != nil {
		log.Printf("Error creating organization: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create organization "})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"messgae": "Organization created"})
}
