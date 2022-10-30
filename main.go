package main

import (
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID         string
	ItemName   string
	SaleStatus bool
}

var items = []item{
	{ID: "1", ItemName: "Shako", SaleStatus: false},
	{ID: "2", ItemName: "Enigma", SaleStatus: false},
	{ID: "3", ItemName: "Death Web", SaleStatus: false},
}

func getAllItems(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, items)
}

func addItem(context *gin.Context) {
	var newItem item

	if err := context.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)

	context.IndentedJSON(http.StatusCreated, newItem)
}

func getItemById(id string) (*item, error) {
	for index, item := range items {
		if item.ID == id {
			return &items[index], nil
		}
	}

	return nil, errors.New("item not found")
}

func getItem(context *gin.Context) {
	id := context.Param("id")
	item, err := getItemById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, item)
}

func toggleSaleStatus(context *gin.Context) {
	id := context.Param("id")
	item, err := getItemById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	item.SaleStatus = !item.SaleStatus

	context.IndentedJSON(http.StatusOK, item)
}

func welcomeMsg(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to my REST API!"})
}

func healthCheck(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"status": "healthy"})
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.1"})
	router.GET("/api/v1", welcomeMsg)
	router.GET("api/v1/status", healthCheck)
	router.GET("/api/v1/items", getAllItems)
	router.GET("/api/v1/items/:id", getItem)
	router.PATCH("/api/v1/items/:id", toggleSaleStatus)
	router.POST("/api/v1/items", addItem)
	router.Run("localhost:8080")
}
