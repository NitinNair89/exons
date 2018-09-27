package routes

import (
	"comm/actions"

	"github.com/labstack/echo"
)

//InitRoutes the routes for the server
func InitRoutes(server *echo.Echo) {
	//User routes
	server.GET("/api/community/get/:id", actions.CommunityGetOne)
	server.GET("/api/community/get/all", actions.CommunityGetAll)
	server.POST("/api/community/add", actions.CommunityAdd)
	server.DELETE("/api/community/update", actions.CommunityUpdate)
}
