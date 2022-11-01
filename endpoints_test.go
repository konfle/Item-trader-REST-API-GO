package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestAddItem(t *testing.T) {
	router := SetUpRouter()
	router.POST("/api/v1/items", addItem)
	itemID := xid.New().String()
	newItem := item{
		ID:         itemID,
		ItemName:   "Test Hearth of The Oak",
		SaleStatus: false,
	}
	jsonValue, _ := json.Marshal(newItem)
	req, _ := http.NewRequest("POST", "/api/v1/items", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
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

func TestHealthChek(t *testing.T) {
	mockResponse := "{\n    \"status\": \"healthy\"\n}"
	router := SetUpRouter()
	router.GET("api/v1/status", healthCheck)
	req, _ := http.NewRequest("GET", "/api/v1/status", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllItems(t *testing.T) {
	router := SetUpRouter()
	router.GET("/api/v1/items", getAllItems)
	req, _ := http.NewRequest("GET", "/api/v1/items", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var companies []item
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, companies)
}

func TestGetItem(t *testing.T) {
	router := SetUpRouter()
	router.GET("/api/v1/items/:id", getItem)
	testItem := item{
		ID:         "2",
		ItemName:   "Test Hearth of The Oak",
		SaleStatus: false,
	}
	jsonValue, _ := json.Marshal(testItem)
	reqFound, _ := http.NewRequest("GET", "/api/v1/items/"+testItem.ID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("GET", "/api/v1/items/13", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestToggleSaleStatus(t *testing.T) {
	router := SetUpRouter()
	router.PATCH("/api/v1/items/:id", toggleSaleStatus)
	testItem := item{
		ID:         "2",
		ItemName:   "Test Hearth of The Oak",
		SaleStatus: false,
	}
	jsonValue, _ := json.Marshal(testItem)
	reqFound, _ := http.NewRequest("PATCH", "/api/v1/items/"+testItem.ID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("PATCH", "/api/v1/items/13", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteItem(t *testing.T) {
	router := SetUpRouter()
	router.DELETE("/api/v1/items/:id", getItem)
	testItem := item{
		ID:         "2",
		ItemName:   "Test Hearth of The Oak",
		SaleStatus: false,
	}
	jsonValue, _ := json.Marshal(testItem)
	reqFound, _ := http.NewRequest("DELETE", "/api/v1/items/"+testItem.ID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("DELETE", "/api/v1/items/13", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
