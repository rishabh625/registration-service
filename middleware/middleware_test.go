package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLatencyMiddleware(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	router := gin.New()
	router.Use(LatencyMiddleware())

	router.GET("/test", func(c *gin.Context) {
		time.Sleep(100 * time.Millisecond) // Simulate processing time
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	logOutput := buf.String()
	expectedLog := "Method: GET, Path: /test, Status: 200, Latency:"
	if !containsLatencyLog(logOutput, expectedLog) {
		t.Errorf("Expected log to contain latency, got: %s", logOutput)
	}
}

func containsLatencyLog(logOutput, expectedLog string) bool {
	return strings.Contains(logOutput, expectedLog)
}
