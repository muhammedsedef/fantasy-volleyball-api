package main

import (
	configuration "fantasy-volleyball-api/appconfig"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/controller/user"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/handler/user/create"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/application/repository/user"
	server "fantasy-volleyball-api/pkg"
	"fantasy-volleyball-api/pkg/couchbase"
	middleware "fantasy-volleyball-api/pkg/middlewaer"
	"github.com/gin-gonic/gin"
)

func main() {
	// bootstrap
	engine := gin.New()
	engine.Use(gin.Recovery())

	middleware.InjectSwaggerUi(engine)
	middleware.InjectHealthCheckMiddleware(engine)

	//couchbase
	couchbaseCluster := couchbase.ConnectCluster(
		configuration.CouchbaseHost,
		configuration.CouchbaseUsername,
		configuration.CouchbasePassword,
	)

	userRepository := repository.NewUserRepository(couchbaseCluster)
	userCreateCommandHandler := create.NewUserCreateCommandHandler(userRepository)
	userController := user.NewUserController(userCreateCommandHandler)

	userGroup := engine.Group("/api/user")
	userGroup.POST("/create", userController.CreateUser)

	server.NewServer(engine).StartHttpServer()
}
