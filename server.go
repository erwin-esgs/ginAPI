package main

import (
	Concurency "ginAPI/controller/concurency"
	ProductController "ginAPI/controller/product"
	UserController "ginAPI/controller/user"

	// "ginAPI/database"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.GET("/user", UserController.Show)
	server.GET("/user/:id", UserController.Detail)
	server.POST("/user", UserController.Create)
	server.PATCH("/user", UserController.Edit)
	server.POST("/user/delete", UserController.Delete)

	server.GET("/product", ProductController.Show)
	server.GET("/product/:id", ProductController.Detail)
	server.POST("/product", ProductController.Create)
	server.PATCH("/product/:id", ProductController.Edit)
	server.POST("/product/delete", ProductController.Delete)

	server.GET("/go", Concurency.Show)

	// server.GET("/migrate", database.Migration)

	server.Run(":8123")
}
