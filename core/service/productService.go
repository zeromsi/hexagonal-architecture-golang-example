package service

import (
	"hexagonal-architecture-example/core/dto"
	"hexagonal-architecture-example/core/entity"
)

type ProductService interface {
	Store( product entity.Product) error
	Update( product entity.Product)  error
	FindAll()[] dto.Product
	FindByCode(code string)dto.Product
	Delete(code string) error
}

