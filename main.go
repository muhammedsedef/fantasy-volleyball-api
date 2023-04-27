package main

import (
	configuration "fantasy-volleyball-api/appconfig"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/controller"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/handler/user/create"
	repository "fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/repository/user"
	server "fantasy-volleyball-api/pkg"
	"fantasy-volleyball-api/pkg/couchbase"
	"github.com/gin-gonic/gin"
)

func main() {
	// bootstrap
	engine := gin.New()
	engine.Use(gin.Recovery())

	//couchbase
	couchbaseCluster := couchbase.ConnectCluster(
		configuration.CouchbaseHost,
		configuration.CouchbaseUsername,
		configuration.CouchbasePassword,
	)

	userRepository := repository.NewUserRepository(couchbaseCluster)
	userCreateCommandHandler := create.NewUserCreateCommandHandler(userRepository)
	userController := controller.NewUserController(userCreateCommandHandler)

	userGroup := engine.Group("/api/user")
	userGroup.POST("/create", userController.CreateUser)
	server.NewServer(engine).StartHttpServer()
}
