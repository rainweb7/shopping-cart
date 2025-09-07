package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Order struct {
	ID    int    `json:"id"`
	Items []Item `json:"items"`
}

var items = []Item{
	{ID: 1, Name: "Item A", Price: 100},
	{ID: 2, Name: "Item B", Price: 200},
	{ID: 3, Name: "Item C", Price: 300},
}

var orders = []Order{}
var orderID = 1

func main() {
	r := gin.Default()

	r.POST("/users/login", func(c *gin.Context) {
		var user struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		c.BindJSON(&user)
		if user.Username == "user" && user.Password == "pass" {
			c.JSON(http.StatusOK, gin.H{"token": "dummy-token"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username/password"})
	})

	r.GET("/items", func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth != "Bearer dummy-token" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.JSON(http.StatusOK, items)
	})

	r.POST("/orders", func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth != "Bearer dummy-token" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		var newOrder struct {
			Items []Item `json:"items"`
		}
		c.BindJSON(&newOrder)
		order := Order{ID: orderID, Items: newOrder.Items}
		orderID++
		orders = append(orders, order)
		c.JSON(http.StatusOK, gin.H{"message": "Order placed"})
	})

	r.GET("/orders", func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth != "Bearer dummy-token" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.JSON(http.StatusOK, orders)
	})

	r.Run(":5000")
}
