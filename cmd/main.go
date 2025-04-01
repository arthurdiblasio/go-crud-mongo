package main

import (
	"github.com/arthurdiblasio/go-crud-mongo/controller"
	"github.com/arthurdiblasio/go-crud-mongo/db"
	"github.com/arthurdiblasio/go-crud-mongo/repository"
	"github.com/arthurdiblasio/go-crud-mongo/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsercase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsercase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
