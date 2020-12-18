package main

import (
	"github.com/apichef/blink-restful-api-go/database"
	"github.com/apichef/blink-restful-api-go/middlewares"
	"github.com/apichef/blink-restful-api-go/routes"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	server = gin.New()
)

func main() {
	registerMiddlewares()
	registerRoutes()
	database.Migrate()
	runServer()
}

func registerMiddlewares() {
	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.Authenticate(),
	)
}

func registerRoutes() {
	rg := server.Group("/api/v1")
	routes.AddBookRoutes(rg)
}

func runServer() {
	err := server.Run(":" + getAppPort())
	if err != nil {
		panic(err.Error())
	}
}

func getAppPort() string {
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		return "8080"
	}

	return port
}