package main

import (
	"gin_practice/controllers"
	"gin_practice/models"
	"gin_practice/repositories"
	"gin_practice/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{
			ID: 1, Name: "商品1", Price: 1000, Description: "説明", SoldOut: false,
		},
		{
			ID: 2, Name: "商品2", Price: 2000, Description: "説明", SoldOut: false,
		},
		{
			ID: 3, Name: "商品3", Price: 3000, Description: "説明", SoldOut: false,
		},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	ItemController := controllers.NewItemController((itemService))

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/items", ItemController.FindAll)

	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
