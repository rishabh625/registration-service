package handlers

import (
	"codingquestions/registration/entities"
	"codingquestions/registration/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

type App struct {
	Handler *RegistrationHandler
}

type RegistrationHandler struct {
	service *service.RegistrationService
}

func NewRegistrationHandler(service *service.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{service: service}
}

func (s *RegistrationHandler) HandleRegistrationRequest(c *gin.Context) {
	var req entities.RegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.service.ValidateAndSaveRequest(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Details could not be processed " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User details processed successfully",
	})
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (s *RegistrationHandler) HandleGetUserRegistration(c *gin.Context) {
	id := c.Param("pan")
	data := s.service.FetchSaveRequest(strings.TrimSpace(id))
	if reflect.DeepEqual(data, entities.RegistrationRequest{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find User"})
		return
	}
	c.JSON(http.StatusOK, data)
}
