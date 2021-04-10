package logic

import (
	"hexagonal-architecture-example/core/dto"
	"hexagonal-architecture-example/core/repository"
	"hexagonal-architecture-example/core/service"
)

type productService struct {
	repo repository.ProductRepository
}

func (p productService) Store() error {
	panic("implement me")
}

func (p productService) Update() error {
	panic("implement me")
}

func (p productService) FindAll() []dto.Product {
	panic("implement me")
}

func (p productService) FindByCode(code string) dto.Product {
	panic("implement me")
}

func (p productService) Delete(code string) error {
	panic("implement me")
}

func NewProductService(repo repository.ProductRepository) service.ProductService {
	return &productService{
		repo,
	}
}
