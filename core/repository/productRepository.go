package repository

import (
	"hexagonal-architecture-example/core/entity"
)

type ProductRepository interface {
	Store() error
	Update() error
	FindAll()[] entity.Product
	FindByCode(code string)entity.Product
	Delete(code string) error
}
