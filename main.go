package main

import (
	"login_api/handlers"
	"login_api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	handlers.InitDB() // Initialize the database connection

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.SetTrustedProxies([]string{"192.168.1.1"})

	// Your routes and handlers here
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	router.POST("/signup", handlers.SignUp)
	router.POST("/login", handlers.Login)
	router.POST("/generate-fetch", handlers.GenerateAndInsertData)
	router.POST("/read-file", handlers.ReadFile)
	router.GET("/getUPC-data", handlers.GetUPC)
	router.POST("/getDataBy-UPC-Country", handlers.GetDataByUPC0)
	router.POST("/getDataBy-UPC-State", handlers.GetDataByUPC1)
	router.POST("/getDataBy-UPC-City", handlers.GetDataByUPC2)
	router.GET("/addToCart-get", handlers.GetAddToCartData)
	router.POST("/addToCart-create", handlers.AddToCartCreate)
	router.PUT("/addToCart-update", handlers.UpdateAddToCart)
	router.DELETE("/addToCart-delete", handlers.DeleteAddToCart)

	router.POST("/upload-profile", handlers.SaveProfileImage)
	router.GET("/get-userInfo", handlers.GetUserInfo)
	router.PUT("/update-userInfo", handlers.UpdateUserInfo)
	router.Run(":8080")
}
