package main

import (
	"github.com/apichef/blink-restful-api-go/database"
	"github.com/apichef/blink-restful-api-go/middlewares"
	"github.com/apichef/blink-restful-api-go/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/example/basic/docs"
	"os"
)

var (
	server = gin.New()
)

// @securityDefinitions.basic
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.Title = "Blink API"
	docs.SwaggerInfo.Description = "This is the GO implementation of Blink API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	registerMiddlewares()
	registerSwagger()
	registerRoutes()
	database.Migrate()
	runServer()
}

func registerSwagger() {
	server.StaticFile("/docs/swagger.json", "./docs/swagger.json")
	url := ginSwagger.URL("http://localhost:8080/docs/swagger.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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