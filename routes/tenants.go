package routes

import (
	"myproject/models"
	"myproject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTenants(c *gin.Context) {
	var tenant models.Tenant

	// Bind JSON data to tenant struct
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	// Save the tenant to the database
	if err := tenant.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create tenant",
			"details": err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Tenant created successfully", "tenant_id": tenant.ID})
}

// function to get tenants
func GetAllTenant(c *gin.Context) {
	tenants, err := models.GetAllTenant()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse tenants details "})
		return
	}
	c.JSON(http.StatusOK, tenants)
}

// login
func login(c *gin.Context) {
	var tenant models.LoginTenant
	err := c.ShouldBindJSON(&tenant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data "})
		return
	}
	err = tenant.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not autheticate "})
		return
	}
	token, err := utils.GenerateToken(tenant.Name, tenant.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not autheticate user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully", "token": token})

}
