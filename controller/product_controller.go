package controller

import (
	"net/http"

	"github.com/arthurdiblasio/go-crud-mongo/model"
	"github.com/arthurdiblasio/go-crud-mongo/usecase"
	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUsercase
}

func NewProductController(usecase usecase.ProductUsercase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
