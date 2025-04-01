package usecase

import (
	"github.com/arthurdiblasio/go-crud-mongo/model"
	"github.com/arthurdiblasio/go-crud-mongo/repository"
)

type ProductUsercase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUsercase {
	return ProductUsercase{
		repository,
	}
}

func (pu *ProductUsercase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
