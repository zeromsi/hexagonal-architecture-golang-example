package service

import (
	"hexagonal-architecture-example/core"
)

type ProductService interface {
	Store( product core.Product) error
	Update( product core.ProductDto)  error
	FindAll()[] core.ProductDto
	FindByCode(code string)core.ProductDto
	Delete(code string) error
}

