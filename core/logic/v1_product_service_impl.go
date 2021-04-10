package logic

import (
	"errors"
	"hexagonal-architecture-example/core"
	"hexagonal-architecture-example/core/repository"
	"hexagonal-architecture-example/core/service"
)

type productService struct {
	repo repository.ProductRepository
}

func (p productService) Store( product core.Product) error {
	existingOne:=p.repo.FindByCode(product.Code)
	if existingOne.Code!=""{
		return errors.New("Already Exists!")
	}
	err := p.repo.Store(product)
	if err != nil {
		return err
	}
	return nil
}

func (p productService) Update( dto core.ProductDto)  error {
	existingOne:=p.repo.FindByCode(dto.Code)
	if existingOne.Code==""{
		return errors.New("No Record Exists!")
	}

	err := p.repo.Update(dto.GetProduct())
	if err != nil {
		return err
	}
	return nil
}

func (p productService) FindAll() []core.ProductDto {
	products:=p.repo.FindAll()
	var dtos [] core.ProductDto
	for _,each:=range products{
		dtos= append(dtos, each.GetProductDto())
	}
	return dtos
}

func (p productService) FindByCode(code string) core.ProductDto {
	product:=p.repo.FindByCode(code)
	return product.GetProductDto()
}

func (p productService) Delete(code string) error {
	return p.repo.Delete(code)
}

func NewV1ProductService(repo repository.ProductRepository) service.ProductService {
	return &productService{
		repo,
	}
}
