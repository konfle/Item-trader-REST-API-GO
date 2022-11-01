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

func getItemIndexByID(id string) (*int, error) {
	for index, item := range items {
		if item.ID == id {
			return &index, nil
		}
	}

	return nil, errors.New("index not found")
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

func deleteItem(context *gin.Context) {
	id := context.Param("id")
	index, err := getItemIndexByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Item not found",
		})
		return
	}

	items = append(items[:*index], items[*index+1:]...)

	context.IndentedJSON(http.StatusOK, gin.H{
		"message": "Item successfully deleted!",
	})
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	v1 := router.Group("/api/v1")
	{
		v1.POST("/items", addItem)
		v1.GET("", welcomeMsg)
		v1.GET("/status", healthCheck)
		v1.GET("/items", getAllItems)
		v1.GET("/items/:id", getItem)
		v1.PATCH("/items/:id", toggleSaleStatus)
		v1.DELETE("/items/:id", deleteItem)
	}

	router.Run("localhost:8080")
}
