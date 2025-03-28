package routes

import (
	"myproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// function to get an events
func getUsers(c *gin.Context){
	users, err := models.GetAllUsers()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch User"})
		return 
	}
	c.JSON(http.StatusOK, users)
}


func createUser(c *gin.Context){
	var user models.User
	err := c.ShouldBindJSON(&user)
	if(err != nil){
		c.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse Data"})
		return 
	}
	err = user.Save()
	if(err != nil){
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create","details":err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"message":"User Created"})
}