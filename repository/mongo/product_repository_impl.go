package mongo

import (
	"errors"
	"hexagonal-architecture-example/core/entity"
	"hexagonal-architecture-example/core/repository"
	"time"
)

var (
	ProductCollection="productCollection"
	NO_KEYS_FOUND_ERROR=errors.New("NO_KEYS_FOUND")
)
type productRepository struct {
	manager  *dmManager
	timeout  time.Duration
}

func (p productRepository) Store() error {
	panic("implement me")
}

func (p productRepository) Update() error {
	panic("implement me")
}

func (p productRepository) FindAll() []entity.Product {
	panic("implement me")
}

func (p productRepository) FindByCode(code string) entity.Product {
	panic("implement me")
}

func (p productRepository) Delete(code string) error {
	panic("implement me")
}

func NewProductRepository(mongoTimeout int) repository.ProductRepository {
	repo := &productRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
	}
	manager:=GetDmManager()
	repo.manager = manager
	return repo
}
