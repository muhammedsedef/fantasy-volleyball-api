package middleware

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InjectHealthCheckMiddleware(engine *gin.Engine) {
	engine.GET("/healthcheck", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
}

func InjectSwaggerUi(ginEngine *gin.Engine) {
	// if configuration.Env() != "prod" {
	// Swagger injection
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Root path to SwaggerUI redirection
	ginEngine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, ctx.Request.URL.Host+"/swagger/index.html")
	})
	// }
}
