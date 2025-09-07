package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "pong"})
	})

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Shopping Cart Backend 🚀")
	})

	// Products route
	r.GET("/api/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"products": []string{"Laptop", "Phone", "Shoes"},
		})
	})

	// ✅ Users routes
	r.POST("/users", func(c *gin.Context) {
		log.Println("POST /users")
		c.JSON(http.StatusCreated, gin.H{"message": "User created"})
	})

	r.GET("/users", func(c *gin.Context) {
		log.Println("GET /users")
		c.JSON(http.StatusOK, gin.H{"users": []string{"user1", "user2"}})
	})

	r.POST("/users/login", func(c *gin.Context) {
		log.Println("POST /users/login")
		c.JSON(http.StatusOK, gin.H{"token": "dummy-token"})
	})

	// ✅ Items routes
	r.POST("/items", func(c *gin.Context) {
		log.Println("POST /items")
		c.JSON(http.StatusCreated, gin.H{"message": "Item created"})
	})

	r.GET("/items", func(c *gin.Context) {
		log.Println("GET /items")
		c.JSON(http.StatusOK, gin.H{"items": []string{"Laptop", "Phone"}})
	})

	// ✅ Carts routes
	r.POST("/carts/", func(c *gin.Context) {
		log.Println("POST /carts/")
		c.JSON(http.StatusCreated, gin.H{"message": "Item added to cart"})
	})

	r.GET("/carts", func(c *gin.Context) {
		log.Println("GET /carts")
		c.JSON(http.StatusOK, gin.H{"carts": []string{"cart1"}})
	})

	// ✅ Orders routes
	r.POST("/orders", func(c *gin.Context) {
		log.Println("POST /orders")
		c.JSON(http.StatusCreated, gin.H{"message": "Order created"})
	})

	r.GET("/orders", func(c *gin.Context) {
		log.Println("GET /orders")
		c.JSON(http.StatusOK, gin.H{"orders": []string{"order1"}})
	})
}
