package dto

import "hexagonal-architecture-example/core/entity"

type Product struct {
	Code      string `json:"code" msgpack:"code" xml:"code"`
	Name        string   `json:"name" msgpack:"name" xml:"name"`
	Price    int32 `json:"price" msgpack:"price" xml:"price"`
}


func(product Product) GetProduct() entity.Product{
	return entity.Product{
		Code:  product.Code,
		Name:  product.Name,
		Price: product.Price,
	}
}