package in_memory

import (
	"errors"
	"hexagonal-architecture-example/core"
	"hexagonal-architecture-example/core/repository"
)

type productRepository struct {
}

func (p productRepository) Store(product core.Product) error {
	products=append(products, product)
	return nil
}

func (p productRepository) Update(product core.Product) error {
	for i,each:=range products{
		if each.Code==product.Code{
			products[i]=product
			break;
		}
	}
	return nil
}

func (p productRepository) FindAll() []core.Product {
	return products
}

func (p productRepository) FindByCode(code string) core.Product {
	for _,each:=range products{
		if each.Code==code{
			return each
		}
	}
	return core.Product{}
}

func (p productRepository) Delete(code string) error {
	index:=-1
	for i,each:=range products{
		if each.Code==code{
			index=i
		}
	}
	if index==-1{
		return errors.New("No record found!")
	}
	products= append(products[:index], products[index+1:]...)
	return nil

}

func NewProductRepository() (repository.ProductRepository) {
	return &productRepository{
	}
}
