package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/users", getUsers)
	server.POST("/users", createUser)
	// id ->> role_id
	server.POST("/roles", CreateRole)
	server.GET("/roles/:id", singleRole)
	server.GET("/rolespermission/:id", singleRolePermisssion)
	server.POST("/permissions", CreatePermission)
	server.POST("/roles/:role_id/permissions", AssignPermissionToRole)

	// create tenants
	server.POST("/addtenants", CreateTenants)
	server.GET("/tenants", GetAllTenant)
	server.POST("/login", login)

	// routes for clients
	server.POST("/addclient", CreateClient)
	server.GET("/client", GetAllClient)

	// create product
	server.POST("/addproduct", CreateProducts)
	server.POST("/calculateprice", CalculatePriceHandler)

	// create organization
	server.POST("/createOrganization", CreateOrganization)

}
