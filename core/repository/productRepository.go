package repository

import (
	"hexagonal-architecture-example/core"
)

type ProductRepository interface {
	Store( core.Product) error
	Update( core.Product) error
	FindAll()[] core.Product
	FindByCode(code string)core.Product
	Delete(code string) error
}
