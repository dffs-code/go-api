package main

import (
	"go-api-products/controller"
	"go-api-products/db"
	"go-api-products/repository"
	"go-api-products/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewPRoductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")
}
