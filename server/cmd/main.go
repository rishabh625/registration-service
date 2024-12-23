package main

import (
	"codingquestions/registration/database"
	"codingquestions/registration/handlers"
	"codingquestions/registration/router"
	"codingquestions/registration/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	app := NewApp() // Initialize App with all dependencies
	router.RegisterRoutes(r, app)
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080

}

func NewApp() *handlers.App {
	db := database.NewInMemoryDatabase()
	svc := service.NewRegistrationService(db)
	handler := handlers.NewRegistrationHandler(svc)
	return &handlers.App{
		Handler: handler,
	}
}
