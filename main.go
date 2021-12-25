package main

import (
	"os"
    "gitub.com/gin-gonic/gin"
	"golang-restaurant-management/database"
	"golang-restaurant-management/routes"
	"golang-restaurant-management/middleware"
	"golang-restaurant-management/controllers"
	"go.mongodb.org"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client,"food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.FoodRoutes(router)
	router.MenuRoutes(router)
	router.TableRoutes(router)
	router.OrderRoutes(router)
	router.OrderItemRoutes(router)
	router.InvoiceRoutes(router)

	router.Run(":" + port)
}
