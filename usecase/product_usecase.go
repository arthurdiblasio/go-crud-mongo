package usecase

import "github.com/arthurdiblasio/go-crud-mongo/model"

type ProductUsercase struct {
}

func NewProductUseCase() ProductUsercase {
	return ProductUsercase{}
}

func (pu *ProductUsercase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
