package redis

import (
	"hexagonal-architecture-example/core"
	"hexagonal-architecture-example/core/repository"
)

type productRepository struct {
}

func (p productRepository) Store(product core.Product) error {
	panic("implement me")
}

func (p productRepository) Update(product core.Product) error {
	panic("implement me")
}

func (p productRepository) FindAll() []core.Product {
	panic("implement me")
}

func (p productRepository) FindByCode(code string) core.Product {
	panic("implement me")
}

func (p productRepository) Delete(code string) error {
	panic("implement me")
}

func NewProductRepository() (repository.ProductRepository) {
	return &productRepository{
	}
}
