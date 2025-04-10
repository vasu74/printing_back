package routes

import (
	"myproject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	if err := role.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create role"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func CreatePermission(c *gin.Context) {
	var permission models.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	if err := permission.PermissionSave(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create permission"})
		return
	}
	c.JSON(http.StatusOK, permission)
}

func AssignPermissionToRole(c *gin.Context) {
	roleID, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid role ID"})
		return
	}

	var rolePermission models.RolePermission
	if err := c.ShouldBindJSON(&rolePermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	// Validate role and permission existence before mapping
	err = models.MapPermissionToRole(roleID, rolePermission.PermissionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Permission assigned successfully"})
}

// get role by its ID
func singleRole(c *gin.Context) {
	roleId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid role ID format"})
		return
	}

	roleExists, err := models.RoleExists(roleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error checking if role exists"})
		return
	}
	if !roleExists {
		c.JSON(http.StatusNotFound, gin.H{"message": "role ID does not exist"})
		return
	}

	role, err := models.GetRoleById(roleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving role details"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func singleRolePermisssion(c *gin.Context) {
	roleId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid role ID format"})
		return
	}

	roleExists, err := models.RoleExists(roleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error checking if role exists"})
		return
	}
	if !roleExists {
		c.JSON(http.StatusNotFound, gin.H{"message": "role ID does not exist"})
		return
	}

	role, err := models.GetPermissionsByRole(roleId, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving role details"})
		return
	}
	c.JSON(http.StatusOK, role)
}
