package router

import (
	"codingquestions/registration/handlers"
	"codingquestions/registration/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(router *gin.Engine, app *handlers.App) {
	router.Use(middleware.LatencyMiddleware())

	router.GET("/ping", handlers.PingHandler)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", app.Handler.HandleRegistrationRequest)
		userRoutes.GET("/:pan", app.Handler.HandleGetUserRegistration)
	}
}
