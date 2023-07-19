package server

import (
	"strconv"

	"skyzar-backend/constants"
	"skyzar-backend/logging"
	"skyzar-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()
}

func StartServer() {
	Router.Use(cors.Default())
	
	routes.Router = Router
	routes.InitRoutes()

	logging.Debug("Server starting on " + strconv.Itoa(constants.Port))
	Router.Run()
}