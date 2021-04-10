package logic

import (
	"hexagonal-architecture-example/core/dto"
	"hexagonal-architecture-example/core/entity"
	"hexagonal-architecture-example/core/repository"
	"hexagonal-architecture-example/core/service"
)

type productService struct {
	repo repository.ProductRepository
}

func (p productService) Store( product entity.Product) error {
	err := p.repo.Store(product)
	if err != nil {
		return err
	}
	return nil
}

func (p productService) Update( product entity.Product)  error {
	err := p.repo.Update(product)
	if err != nil {
		return err
	}
	return nil
}

func (p productService) FindAll() []dto.Product {
	products:=p.repo.FindAll()
	var dtos [] dto.Product
	for _,each:=range products{
		dtos= append(dtos, each.GetProductDto())
	}
	return dtos
}

func (p productService) FindByCode(code string) dto.Product {
	product:=p.repo.FindByCode(code)
	return product.GetProductDto()
}

func (p productService) Delete(code string) error {
	panic("implement me")
}

func NewProductService(repo repository.ProductRepository) service.ProductService {
	return &productService{
		repo,
	}
}
