package main

// https://circleci.com/blog/gin-gonic-testing/
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestWelcomeMsq(t *testing.T) {
	mockResponse := "{\n    \"message\": \"Welcome to my REST API!\"\n}"
	router := SetUpRouter()
	router.GET("/api/v1", welcomeMsg)
	req, _ := http.NewRequest("GET", "/api/v1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCompaniesHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/v1/items", getAllItems)
	req, _ := http.NewRequest("GET", "/api/v1/items", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []item
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, companies)
}
