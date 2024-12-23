package handlers_test

import (
	"bytes"
	"codingquestions/registration/entities"
	"codingquestions/registration/handlers"
	"codingquestions/registration/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter(handler *handlers.RegistrationHandler) *gin.Engine {
	router := gin.Default()
	router.POST("/register", handler.HandleRegistrationRequest)
	router.GET("/user/:pan", handler.HandleGetUserRegistration)
	router.GET("/ping", handlers.PingHandler)
	return router
}

func TestHandleRegistrationRequestSuccess(t *testing.T) {
	mockService := service.NewRegistrationServiceMock()
	handler := handlers.NewRegistrationHandler(mockService)
	router := setupRouter(handler)

	validRequest := entities.RegistrationRequest{
		PAN:    "ABCDE1234F",
		Number: "9876543210",
		Name:   "John Doe",
		Email:  "john.doe@example.com",
	}
	body, _ := json.Marshal(validRequest)

	// Test valid registration request
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestHandleRegistrationRequestFailedInvalidPan(t *testing.T) {
	mockService := service.NewRegistrationServiceMock()
	handler := handlers.NewRegistrationHandler(mockService)
	router := setupRouter(handler)

	// Test invalid request (invalid PAN)
	invalidRequest := entities.RegistrationRequest{
		Name:   "John Doe",
		Email:  "jondoe@gmail.com",
		PAN:    "APMCL2889K1", // invalid pan 6 letters at start and ending in numeric
		Number: "9876543210",
	}
	body, _ := json.Marshal(invalidRequest)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestHandleRegistrationRequestFailedPANMissing(t *testing.T) {
	mockService := service.NewRegistrationServiceMock()
	handler := handlers.NewRegistrationHandler(mockService)
	router := setupRouter(handler)

	// Test invalid request (missing PAN)
	invalidRequest := entities.RegistrationRequest{
		Name:   "John Doe",
		Email:  "jondoe@gmail.com",
		Number: "9876543210"}
	body, _ := json.Marshal(invalidRequest)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestHandleRegistrationRequestFailedInvalidNumber(t *testing.T) {
	mockService := service.NewRegistrationServiceMock()
	handler := handlers.NewRegistrationHandler(mockService)
	router := setupRouter(handler)

	validRequest := entities.RegistrationRequest{
		PAN:    "ABCDE1234F",
		Number: "98765410AB",
		Name:   "John Doe",
		Email:  "john.doe@example.com",
	}
	body, _ := json.Marshal(validRequest)

	// Test invalid number
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestHandleRegistrationRequestFailedInvalidEmail(t *testing.T) {
	mockService := service.NewRegistrationServiceMock()
	handler := handlers.NewRegistrationHandler(mockService)
	router := setupRouter(handler)

	validRequest := entities.RegistrationRequest{
		PAN:    "ABCDE1234F",
		Number: "9876543210",
		Name:   "John Doe",
		Email:  "john.doe@example.com",
	}
	body, _ := json.Marshal(validRequest)

	// Test valid registration request
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	// Test invalid request (missing PAN)
	invalidRequest := entities.RegistrationRequest{
		Number: "9876543210",
	}
	body, _ = json.Marshal(invalidRequest)
	req = httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestHandleGetUserRegistration(t *testing.T) {
	mockService := service.NewRegistrationServiceMock()
	handler := handlers.NewRegistrationHandler(mockService)
	router := setupRouter(handler)
	validRequest := entities.RegistrationRequest{
		PAN:    "ABCDE1234F",
		Number: "9876543210",
		Name:   "John Doe",
		Email:  "john.doe@example.com",
	}
	body, _ := json.Marshal(validRequest)

	// Test valid registration request
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
	testPAN := "ABCDE1234F"

	// Test fetching valid user
	req = httptest.NewRequest(http.MethodGet, "/user/"+testPAN, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Test fetching non-existent user
	nonExistentPAN := "ZZZZZ9999Z"
	req = httptest.NewRequest(http.MethodGet, "/user/"+nonExistentPAN, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestPingHandler(t *testing.T) {
	router := gin.Default()
	router.GET("/ping", handlers.PingHandler)

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	expectedBody := `{"message":"pong"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}
