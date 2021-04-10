package service

import (
	"hexagonal-architecture-example/core/dto"
)

type ProductService interface {
	Store() error
	Update() error
	FindAll()[] dto.Product
	FindByCode(code string)dto.Product
	Delete(code string) error
}

